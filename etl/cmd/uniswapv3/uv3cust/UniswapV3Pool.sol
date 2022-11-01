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
import 'libraries/TransferHelper.sol';
import 'libraries/SwapMath.sol';

import 'interfaces/callback/IUniswapV3MintCallback.sol';
import 'interfaces/callback/IUniswapV3SwapCallback.sol';
import 'interfaces/callback/IUniswapV3FlashCallback.sol';

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
	function burn(
		int24 tickLower,
		int24 tickUpper,
		uint128 amount
	) external returns (uint256 amount0, uint256 amount1) {
		(Position.Info storage position, int256 amount0Int, int256 amount1Int) =
			_modifyPosition(
				ModifyPositionParams({
					owner: msg.sender,
					tickLower: tickLower,
					tickUpper: tickUpper,
					liquidityDelta: -int256(amount).toInt128()
				})
			);

		amount0 = uint256(-amount0Int);
		amount1 = uint256(-amount1Int);

		if (amount0 > 0 || amount1 > 0) {
			(position.tokensOwed0, position.tokensOwed1) = (
				position.tokensOwed0 + uint128(amount0),
				position.tokensOwed1 + uint128(amount1)
			);
		}

		emit Burn(msg.sender, tickLower, tickUpper, amount, amount0, amount1);
	}
	struct SwapCache {
		// the protocol fee for the input token
		uint8 feeProtocol;
		// liquidity at the beginning of the swap
		uint128 liquidityStart;
		// the timestamp of the current block
		uint32 blockTimestamp;
		// the current value of the tick accumulator, computed only if we cross an initialized tick
		int56 tickCumulative;
		// the current value of seconds per liquidity accumulator, computed only if we cross an initialized tick
		uint160 secondsPerLiquidityCumulativeX128;
		// whether we've computed and cached the above two accumulators
		bool computedLatestObservation;
	}
	struct SwapState {
		// the amount remaining to be swapped in/out of the input/output asset
		int256 amountSpecifiedRemaining;
		// the amount already swapped out/in of the output/input asset
		int256 amountCalculated;
		// current sqrt(price)
		uint160 sqrtPriceX96;
		// the tick associated with the current price
		int24 tick;
		// the global fee growth of the input token
		uint256 feeGrowthGlobalX128;
		// amount of input token paid as protocol fee
		uint128 protocolFee;
		// the current liquidity in range
		uint128 liquidity;
	}
	struct StepComputations {
		// the price at the beginning of the step
		uint160 sqrtPriceStartX96;
		// the next tick to swap to from the current tick in the swap direction
		int24 tickNext;
		// whether tickNext is initialized or not
		bool initialized;
		// sqrt(price) for the next tick (1/0)
		uint160 sqrtPriceNextX96;
		// how much is being swapped in in this step
		uint256 amountIn;
		// how much is being swapped out
		uint256 amountOut;
		// how much fee is being paid in
		uint256 feeAmount;
	}
	function swap(
		address recipient,
		bool zeroForOne,
		int256 amountSpecified,
		uint160 sqrtPriceLimitX96,
		bytes calldata data
	) external returns (int256 amount0, int256 amount1) {
		require(amountSpecified != 0, 'AS');

		Slot0 memory slot0Start = slot0;

		require(slot0Start.unlocked, 'LOK');
		require(
			zeroForOne
				? sqrtPriceLimitX96 < slot0Start.sqrtPriceX96 && sqrtPriceLimitX96 > TickMath.MIN_SQRT_RATIO
				: sqrtPriceLimitX96 > slot0Start.sqrtPriceX96 && sqrtPriceLimitX96 < TickMath.MAX_SQRT_RATIO,
				'SPL'
		);

		slot0.unlocked = false;

		SwapCache memory cache =
			SwapCache({
				liquidityStart: liquidity,
				blockTimestamp: _blockTimestamp(),
				feeProtocol: zeroForOne ? (slot0Start.feeProtocol % 16) : (slot0Start.feeProtocol >> 4),
				secondsPerLiquidityCumulativeX128: 0,
				tickCumulative: 0,
				computedLatestObservation: false
			});

		bool exactInput = amountSpecified > 0;

	SwapState memory state =
			SwapState({
				amountSpecifiedRemaining: amountSpecified,
				amountCalculated: 0,
				sqrtPriceX96: slot0Start.sqrtPriceX96,
				tick: slot0Start.tick,
				feeGrowthGlobalX128: zeroForOne ? feeGrowthGlobal0X128 : feeGrowthGlobal1X128,
				protocolFee: 0,
				liquidity: cache.liquidityStart
			});

		// continue swapping as long as we haven't used the entire input/output and haven't reached the price limit
		while (state.amountSpecifiedRemaining != 0 && state.sqrtPriceX96 != sqrtPriceLimitX96) {
			StepComputations memory step;

			step.sqrtPriceStartX96 = state.sqrtPriceX96;

			(step.tickNext, step.initialized) = tickBitmap.nextInitializedTickWithinOneWord(
				state.tick,
				tickSpacing,
				zeroForOne
			);

			// ensure that we do not overshoot the min/max tick, as the tick bitmap is not aware of these bounds
			if (step.tickNext < TickMath.MIN_TICK) {
				step.tickNext = TickMath.MIN_TICK;
			} else if (step.tickNext > TickMath.MAX_TICK) {
				step.tickNext = TickMath.MAX_TICK;
			}

			// get the price for the next tick
			step.sqrtPriceNextX96 = TickMath.getSqrtRatioAtTick(step.tickNext);

			// compute values to swap to the target tick, price limit, or point where input/output amount is exhausted
			(state.sqrtPriceX96, step.amountIn, step.amountOut, step.feeAmount) = SwapMath.computeSwapStep(
				state.sqrtPriceX96,
				(zeroForOne ? step.sqrtPriceNextX96 < sqrtPriceLimitX96 : step.sqrtPriceNextX96 > sqrtPriceLimitX96)
				? sqrtPriceLimitX96
				: step.sqrtPriceNextX96,
				state.liquidity,
				state.amountSpecifiedRemaining,
				fee
			);

			if (exactInput) {
				state.amountSpecifiedRemaining -= (step.amountIn + step.feeAmount).toInt256();
				state.amountCalculated = state.amountCalculated.sub(step.amountOut.toInt256());
			} else {
				state.amountSpecifiedRemaining += step.amountOut.toInt256();
				state.amountCalculated = state.amountCalculated.add((step.amountIn + step.feeAmount).toInt256());
			}

			// if the protocol fee is on, calculate how much is owed, decrement feeAmount, and increment protocolFee
			if (cache.feeProtocol > 0) {
				uint256 delta = step.feeAmount / cache.feeProtocol;
				step.feeAmount -= delta;
				state.protocolFee += uint128(delta);
			}

			// update global fee tracker
			if (state.liquidity > 0)
				state.feeGrowthGlobalX128 += FullMath.mulDiv(step.feeAmount, FixedPoint128.Q128, state.liquidity);

			// shift tick if we reached the next price
			if (state.sqrtPriceX96 == step.sqrtPriceNextX96) {
				// if the tick is initialized, run the tick transition
				if (step.initialized) {
					// check for the placeholder value, which we replace with the actual value the first time the swap
					// crosses an initialized tick
					if (!cache.computedLatestObservation) {
						(cache.tickCumulative, cache.secondsPerLiquidityCumulativeX128) = observations.observeSingle(
							cache.blockTimestamp,
							0,
							slot0Start.tick,
							slot0Start.observationIndex,
							cache.liquidityStart,
							slot0Start.observationCardinality
						);
						cache.computedLatestObservation = true;
					}
					int128 liquidityNet =
						ticks.cross(
							step.tickNext,
							(zeroForOne ? state.feeGrowthGlobalX128 : feeGrowthGlobal0X128),
							(zeroForOne ? feeGrowthGlobal1X128 : state.feeGrowthGlobalX128),
							cache.secondsPerLiquidityCumulativeX128,
							cache.tickCumulative,
							cache.blockTimestamp
						);
					// if we're moving leftward, we interpret liquidityNet as the opposite sign
					// safe because liquidityNet cannot be type(int128).min
					if (zeroForOne) liquidityNet = -liquidityNet;

					state.liquidity = LiquidityMath.addDelta(state.liquidity, liquidityNet);
				}

				state.tick = zeroForOne ? step.tickNext - 1 : step.tickNext;
			} else if (state.sqrtPriceX96 != step.sqrtPriceStartX96) {
				// recompute unless we're on a lower tick boundary (i.e. already transitioned ticks), and haven't moved
				state.tick = TickMath.getTickAtSqrtRatio(state.sqrtPriceX96);
			}
		}

		// update tick and write an oracle entry if the tick change
		if (state.tick != slot0Start.tick) {
			(uint16 observationIndex, uint16 observationCardinality) =
				observations.write(
					slot0Start.observationIndex,
					cache.blockTimestamp,
					slot0Start.tick,
					cache.liquidityStart,
					slot0Start.observationCardinality,
					slot0Start.observationCardinalityNext
				);
			(slot0.sqrtPriceX96, slot0.tick, slot0.observationIndex, slot0.observationCardinality) = (
				state.sqrtPriceX96,
				state.tick,
				observationIndex,
				observationCardinality
			);
		} else {
			// otherwise just update the price
			slot0.sqrtPriceX96 = state.sqrtPriceX96;
		}

		// update liquidity if it changed
		if (cache.liquidityStart != state.liquidity) liquidity = state.liquidity;

		// update fee growth global and, if necessary, protocol fees
		// overflow is acceptable, protocol has to withdraw before it hits type(uint128).max fees
		if (zeroForOne) {
			feeGrowthGlobal0X128 = state.feeGrowthGlobalX128;
			if (state.protocolFee > 0) protocolFees.token0 += state.protocolFee;
		} else {
			feeGrowthGlobal1X128 = state.feeGrowthGlobalX128;
			if (state.protocolFee > 0) protocolFees.token1 += state.protocolFee;
		}

		(amount0, amount1) = zeroForOne == exactInput
			? (amountSpecified - state.amountSpecifiedRemaining, state.amountCalculated)
			: (state.amountCalculated, amountSpecified - state.amountSpecifiedRemaining);

		// do the transfers and collect payment
		if (zeroForOne) {
			if (amount1 < 0) TransferHelper.safeTransfer(token1, recipient, uint256(-amount1));

			uint256 balance0Before = balance0();
			IUniswapV3SwapCallback(msg.sender).uniswapV3SwapCallback(amount0, amount1, data);
			require(balance0Before.add(uint256(amount0)) <= balance0(), 'IIA');
		} else {
			if (amount0 < 0) TransferHelper.safeTransfer(token0, recipient, uint256(-amount0));

			uint256 balance1Before = balance1();
			IUniswapV3SwapCallback(msg.sender).uniswapV3SwapCallback(amount0, amount1, data);
			require(balance1Before.add(uint256(amount1)) <= balance1(), 'IIA');
		}

		emit Swap(msg.sender, recipient, amount0, amount1, state.sqrtPriceX96, state.liquidity, state.tick);
		slot0.unlocked = true;
	}
	function flash(
		address recipient,
		uint256 amount0,
		uint256 amount1,
		bytes calldata data
	) external {
		uint128 _liquidity = liquidity;
		require(_liquidity > 0, 'L');

		uint256 fee0 = FullMath.mulDivRoundingUp(amount0, fee, 1e6);
		uint256 fee1 = FullMath.mulDivRoundingUp(amount1, fee, 1e6);
		uint256 balance0Before = balance0();
		uint256 balance1Before = balance1();

		if (amount0 > 0) TransferHelper.safeTransfer(token0, recipient, amount0);
		if (amount1 > 0) TransferHelper.safeTransfer(token1, recipient, amount1);

		IUniswapV3FlashCallback(msg.sender).uniswapV3FlashCallback(fee0, fee1, data);

		uint256 balance0After = balance0();
		uint256 balance1After = balance1();

		require(balance0Before.add(fee0) <= balance0After, 'F0');
		require(balance1Before.add(fee1) <= balance1After, 'F1');

		// sub is safe because we know balanceAfter is gt balanceBefore by at least fee
		uint256 paid0 = balance0After - balance0Before;
		uint256 paid1 = balance1After - balance1Before;

		if (paid0 > 0) {
			uint8 feeProtocol0 = slot0.feeProtocol % 16;
			uint256 fees0 = feeProtocol0 == 0 ? 0 : paid0 / feeProtocol0;
			if (uint128(fees0) > 0) protocolFees.token0 += uint128(fees0);
			feeGrowthGlobal0X128 += FullMath.mulDiv(paid0 - fees0, FixedPoint128.Q128, _liquidity);
		}
		if (paid1 > 0) {
			uint8 feeProtocol1 = slot0.feeProtocol >> 4;
			uint256 fees1 = feeProtocol1 == 0 ? 0 : paid1 / feeProtocol1;
			if (uint128(fees1) > 0) protocolFees.token1 += uint128(fees1);
			feeGrowthGlobal1X128 += FullMath.mulDiv(paid1 - fees1, FixedPoint128.Q128, _liquidity);
		}

		emit Flash(msg.sender, recipient, amount0, amount1, paid0, paid1);
	}

	function setFeeProtocol(uint8 feeProtocol0, uint8 feeProtocol1) external {
		require(
			(feeProtocol0 == 0 || (feeProtocol0 >= 4 && feeProtocol0 <= 10)) &&
			(feeProtocol1 == 0 || (feeProtocol1 >= 4 && feeProtocol1 <= 10))
		);
		uint8 feeProtocolOld = slot0.feeProtocol;
		slot0.feeProtocol = feeProtocol0 + (feeProtocol1 << 4);
		emit SetFeeProtocol(feeProtocolOld % 16, feeProtocolOld >> 4, feeProtocol0, feeProtocol1);
	}

	function collectProtocol(
		address recipient,
		uint128 amount0Requested,
		uint128 amount1Requested
	) external returns (uint128 amount0, uint128 amount1) {
		amount0 = amount0Requested > protocolFees.token0 ? protocolFees.token0 : amount0Requested;
		amount1 = amount1Requested > protocolFees.token1 ? protocolFees.token1 : amount1Requested;

		if (amount0 > 0) {
			if (amount0 == protocolFees.token0) amount0--; // ensure that the slot is not cleared, for gas savings
			protocolFees.token0 -= amount0;
			TransferHelper.safeTransfer(token0, recipient, amount0);
		}
		if (amount1 > 0) {
			if (amount1 == protocolFees.token1) amount1--; // ensure that the slot is not cleared, for gas savings
			protocolFees.token1 -= amount1;
			TransferHelper.safeTransfer(token1, recipient, amount1);
		}

		emit CollectProtocol(msg.sender, recipient, amount0, amount1);
	}
}
