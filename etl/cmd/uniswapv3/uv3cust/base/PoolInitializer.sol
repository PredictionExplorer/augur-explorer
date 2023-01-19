// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity =0.7.6;

import '../interfaces/IUniswapV3Factory.sol';
import '../interfaces/IUniswapV3Pool.sol';
import "../interfaces/pool/IUniswapV3PoolActions.sol";
import "../interfaces/pool/IUniswapV3PoolState.sol";

import '../base/PeripheryImmutableState.sol';
import '../interfaces/IPoolInitializer.sol';

/// @title Creates and initializes V3 Pools
abstract contract PoolInitializer is IPoolInitializer, PeripheryImmutableState {
    /// @inheritdoc IPoolInitializer
    function createAndInitializePoolIfNecessary(
        address token0,
        address token1,
        uint24 fee,
        uint160 sqrtPriceX96
    ) external payable override returns (address pool) {
        require(token0 < token1);
        pool = IUniswapV3Factory(factory).getPool(token0, token1, fee);
        if (pool == address(0)) {
            pool = IUniswapV3Factory(factory).createPool(token0, token1, fee);
            IUniswapV3PoolActions(pool).initialize(sqrtPriceX96);
        } else {
            (uint160 sqrtPriceX96Existing, , , , , , ) = IUniswapV3PoolState(pool).slot0();
            if (sqrtPriceX96Existing == 0) {
                IUniswapV3PoolActions(pool).initialize(sqrtPriceX96);
            }
        }
    }
}
