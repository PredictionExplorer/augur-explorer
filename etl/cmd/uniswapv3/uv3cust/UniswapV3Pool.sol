 // SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.7.6;

import 'interfaces/IERC20Minimal.sol';
import 'interfaces/IUniswapV3Pool.sol';
import 'interfaces/IUniswapV3PoolDeployer.sol';
import 'libraries/Tick.sol';
import 'libraries/TickBitmap.sol';
import 'libraries/Position.sol';
import 'libraries/Oracle.sol';
import 'libraries/SqrtPriceMath.sol';
import 'libraries/UnsafeMath.sol';

import 'interfaces/callback/IUniswapV3MintCallback.sol';

contract UniswapV3Pool is IUniswapV3Pool {
    using LowGasSafeMath for uint256;
    using LowGasSafeMath for int256;
    using SafeCast for uint256;
    using SafeCast for int256;
    using Tick for mapping(int24 => Tick.Info);
    using TickBitmap for mapping(int16 => uint256);
    using Position for mapping(bytes32 => Position.Info);
    using Position for Position.Info;
    using Oracle for Oracle.Observation[65535];

	address public immutable override factory;
	address public immutable override token0;
	address public immutable override token1;
	uint24 public immutable override fee;
	int24 public immutable override tickSpacing;
	uint128 public immutable override maxLiquidityPerTick;
	struct Slot0 {
		// the current price
		uint160 sqrtPriceX96;
		// the current tick
		int24 tick;
		// the most-recently updated index of the observations array
		uint16 observationIndex;
		// the current maximum number of observations that are being stored
		uint16 observationCardinality;
		// the next maximum number of observations to store, triggered in observations.write
		uint16 observationCardinalityNext;
		// the current protocol fee as a percentage of the swap fee taken on withdrawal
		// represented as an integer denominator (1/x)%
		uint8 feeProtocol;
		// whether the pool is locked
		bool unlocked;
	}
	Slot0 public slot0;
	uint256 public feeGrowthGlobal0X128;
	uint256 public feeGrowthGlobal1X128;
	struct ProtocolFees {
		uint128 token0;
		uint128 token1;
	}
	ProtocolFees public protocolFees;
	uint128 public liquidity;
	mapping(int24 => Tick.Info) public ticks;
	mapping(int16 => uint256) public tickBitmap;
	mapping(bytes32 => Position.Info) public positions;
	Oracle.Observation[65535] public observations;

	constructor() {
		int24 _tickSpacing;
		(factory, token0, token1, fee, _tickSpacing) = IUniswapV3PoolDeployer(msg.sender).parameters();
		tickSpacing = _tickSpacing;

		maxLiquidityPerTick = Tick.tickSpacingToMaxLiquidityPerTick(_tickSpacing);
	}
	function balance0() private view returns (uint256) {
		(bool success, bytes memory data) =
			token0.staticcall(abi.encodeWithSelector(IERC20Minimal.balanceOf.selector, address(this)));
		require(success && data.length >= 32);
		return abi.decode(data, (uint256));
	}
	function balance1() private view returns (uint256) {
		(bool success, bytes memory data) =
			token1.staticcall(abi.encodeWithSelector(IERC20Minimal.balanceOf.selector, address(this)));
		require(success && data.length >= 32);
		return abi.decode(data, (uint256));
	}
	function _blockTimestamp() internal view virtual returns (uint32) {
		return uint32(block.timestamp); // truncation is desired
	}
	function checkTicks(int24 tickLower, int24 tickUpper) private pure {
		require(tickLower < tickUpper, 'TLU');
		require(tickLower >= TickMath.MIN_TICK, 'TLM');
		require(tickUpper <= TickMath.MAX_TICK, 'TUM');
	}
	function initialize(uint160 sqrtPriceX96) external {
		require(slot0.sqrtPriceX96 == 0, 'AI');

		int24 tick = TickMath.getTickAtSqrtRatio(sqrtPriceX96);

		(uint16 cardinality, uint16 cardinalityNext) = observations.initialize(_blockTimestamp());

		slot0 = Slot0({
			sqrtPriceX96: sqrtPriceX96,
			tick: tick,
			observationIndex: 0,
			observationCardinality: cardinality,
			observationCardinalityNext: cardinalityNext,
			feeProtocol: 0,
			unlocked: true
		});

		emit Initialize(sqrtPriceX96, tick);
	}
	struct ModifyPositionParams {
		// the address that owns the position
		address owner;
		// the lower and upper tick of the position
		int24 tickLower;
		int24 tickUpper;
		// any change in liquidity
		int128 liquidityDelta;
	}
	function _modifyPosition(ModifyPositionParams memory params) private
		returns (
			Position.Info storage position,
			int256 amount0,
			int256 amount1
		)
	{
		checkTicks(params.tickLower, params.tickUpper);

		Slot0 memory _slot0 = slot0; // SLOAD for gas optimization

		position = _updatePosition(
			params.owner,
			params.tickLower,
			params.tickUpper,
			params.liquidityDelta,
			_slot0.tick
		);

		if (params.liquidityDelta != 0) {
			if (_slot0.tick < params.tickLower) {
				// current tick is below the passed range; liquidity can only become in range by crossing from left to
				// right, when we'll need _more_ token0 (it's becoming more valuable) so user must provide it
				amount0 = SqrtPriceMath.getAmount0Delta(
					TickMath.getSqrtRatioAtTick(params.tickLower),
					TickMath.getSqrtRatioAtTick(params.tickUpper),
					params.liquidityDelta
				);
			} else if (_slot0.tick < params.tickUpper) {
				// current tick is inside the passed range
				uint128 liquidityBefore = liquidity; // SLOAD for gas optimization

				// write an oracle entry
				(slot0.observationIndex, slot0.observationCardinality) = observations.write(
					_slot0.observationIndex,
					_blockTimestamp(),
					_slot0.tick,
					liquidityBefore,
					_slot0.observationCardinality,
					_slot0.observationCardinalityNext
				);

				amount0 = SqrtPriceMath.getAmount0Delta(
					_slot0.sqrtPriceX96,
					TickMath.getSqrtRatioAtTick(params.tickUpper),
					params.liquidityDelta
				);
				amount1 = SqrtPriceMath.getAmount1Delta(
					TickMath.getSqrtRatioAtTick(params.tickLower),
					_slot0.sqrtPriceX96,
					params.liquidityDelta
				);

				liquidity = LiquidityMath.addDelta(liquidityBefore, params.liquidityDelta);
			} else {
				// current tick is above the passed range; liquidity can only become in range by crossing from right to
				// left, when we'll need _more_ token1 (it's becoming more valuable) so user must provide it
				amount1 = SqrtPriceMath.getAmount1Delta(
					TickMath.getSqrtRatioAtTick(params.tickLower),
					TickMath.getSqrtRatioAtTick(params.tickUpper),
					params.liquidityDelta
				);
			}
		}
	}
	function _updatePosition(
		address owner,
		int24 tickLower,
		int24 tickUpper,
		int128 liquidityDelta,
		int24 tick
	) private returns (Position.Info storage position) {
		position = positions.get(owner, tickLower, tickUpper);

		uint256 _feeGrowthGlobal0X128 = feeGrowthGlobal0X128; // SLOAD for gas optimization
		uint256 _feeGrowthGlobal1X128 = feeGrowthGlobal1X128; // SLOAD for gas optimization

		// if we need to update the ticks, do it
		bool flippedLower;
		bool flippedUpper;
		if (liquidityDelta != 0) {
			uint32 time = _blockTimestamp();
			(int56 tickCumulative, uint160 secondsPerLiquidityCumulativeX128) =
				observations.observeSingle(
					time,
					0,
					slot0.tick,
					slot0.observationIndex,
					liquidity,
					slot0.observationCardinality
				);

			flippedLower = ticks.update(
				tickLower,
				tick,
				liquidityDelta,
				_feeGrowthGlobal0X128,
				_feeGrowthGlobal1X128,
				secondsPerLiquidityCumulativeX128,
				tickCumulative,
				time,
				false,
				maxLiquidityPerTick
			);
			flippedUpper = ticks.update(
				tickUpper,
				tick,
				liquidityDelta,
				_feeGrowthGlobal0X128,
				_feeGrowthGlobal1X128,
				secondsPerLiquidityCumulativeX128,
				tickCumulative,
				time,
				true,
				maxLiquidityPerTick
			);

			if (flippedLower) {
				tickBitmap.flipTick(tickLower, tickSpacing);
			}
			if (flippedUpper) {
				tickBitmap.flipTick(tickUpper, tickSpacing);
			}
		}

		(uint256 feeGrowthInside0X128, uint256 feeGrowthInside1X128) =
			ticks.getFeeGrowthInside(tickLower, tickUpper, tick, _feeGrowthGlobal0X128, _feeGrowthGlobal1X128);

		position.update(liquidityDelta, feeGrowthInside0X128, feeGrowthInside1X128);

		// clear any tick data that is no longer needed
		if (liquidityDelta < 0) {
			if (flippedLower) {
				ticks.clear(tickLower);
			}
			if (flippedUpper) {
				ticks.clear(tickUpper);
			}
		}
	}
	function mint(
		address recipient,
		int24 tickLower,
		int24 tickUpper,
		uint128 amount,
		bytes calldata data
	) external returns (uint256 amount0, uint256 amount1) {
		require(amount > 0);
		(, int256 amount0Int, int256 amount1Int) =
			_modifyPosition(
				ModifyPositionParams({
					owner: recipient,
					tickLower: tickLower,
					tickUpper: tickUpper,
					liquidityDelta: int256(amount).toInt128()
				})
			);

		amount0 = uint256(amount0Int);
		amount1 = uint256(amount1Int);

		uint256 balance0Before;
		uint256 balance1Before;
		if (amount0 > 0) balance0Before = balance0();
		if (amount1 > 0) balance1Before = balance1();
		IUniswapV3MintCallback(msg.sender).uniswapV3MintCallback(amount0, amount1, data);
		if (amount0 > 0) require(balance0Before.add(amount0) <= balance0(), 'M0');
		if (amount1 > 0) require(balance1Before.add(amount1) <= balance1(), 'M1');

		emit Mint(msg.sender, recipient, tickLower, tickUpper, amount, amount0, amount1);
	}
}
