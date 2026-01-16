import { ethers } from 'ethers'
import { Pool,Position,NonfungiblePositionManager } from '@uniswap/v3-sdk'
import { Token,Percent } from '@uniswap/sdk-core'
import { abi as IUniswapV3PoolABI } from '@uniswap/v3-core/artifacts/contracts/interfaces/IUniswapV3Pool.sol/IUniswapV3Pool.json'
import { nearestUsableTick, TickMath } from '@uniswap/v3-sdk'
import { INonfungiblePositionManager as pmabi } from './INonfungiblePositionManager.json'

const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545')
const poolAddress = '0xAe9a1Df527E36DE6EBa251eA4FBAfC897e1D7E9A'
const poolContract = new ethers.Contract(poolAddress, IUniswapV3PoolABI, provider)

interface Immutables {
  factory: string
  token0: string
  token1: string
  fee: number
  tickSpacing: number
  maxLiquidityPerTick: ethers.BigNumber
}

interface State {
  liquidity: ethers.BigNumber
  sqrtPriceX96: ethers.BigNumber
  tick: number
  observationIndex: number
  observationCardinality: number
  observationCardinalityNext: number
  feeProtocol: number
  unlocked: boolean
}

async function getPoolImmutables() {
  const [factory, token0, token1, fee, tickSpacing, maxLiquidityPerTick] = await Promise.all([
    poolContract.factory(),
    poolContract.token0(),
    poolContract.token1(),
    poolContract.fee(),
    poolContract.tickSpacing(),
    poolContract.maxLiquidityPerTick(),
  ])

  const immutables: Immutables = {
    factory,
    token0,
    token1,
    fee,
    tickSpacing,
    maxLiquidityPerTick,
  }
  return immutables
}

async function getPoolState() {
  const [liquidity, slot] = await Promise.all([poolContract.liquidity(), poolContract.slot0()])

  const PoolState: State = {
    liquidity,
    sqrtPriceX96: slot[0],
    tick: slot[1],
    observationIndex: slot[2],
    observationCardinality: slot[3],
    observationCardinalityNext: slot[4],
    feeProtocol: slot[5],
    unlocked: slot[6],
  }

  return PoolState
}

async function main() {
  const [immutables, state] = await Promise.all([getPoolImmutables(), getPoolState()])


  const TokenB = new Token(3, immutables.token0, 18, 'WETH', 'Wrapped Ether')
  const TokenA = new Token(3, immutables.token1, 6, 'USDC', 'USD Coin')

  const poolExample = new Pool(
    TokenA,
    TokenB,
    immutables.fee,
    state.sqrtPriceX96.toString(),
    state.liquidity.toString(),
    state.tick
  )
  console.log(poolExample)

	console.log("pool.tickspacing:")
	console.log(poolExample.tickSpacing)

	const position = new Position({
		pool: poolExample,
		liquidity: 20000,
		tickLower: nearestUsableTick(state.tick, immutables.tickSpacing) - immutables.tickSpacing * 2,
		tickUpper: nearestUsableTick(state.tick, immutables.tickSpacing) + immutables.tickSpacing * 2,
	})
	console.log("tickLower:")
	console.log(position.tickLower)
	console.log("tickUpper:")
	console.log(position.tickUpper)
	const deadline = 1719099018
	const { calldata, value } = NonfungiblePositionManager.addCallParameters(position, {
		slippageTolerance: new Percent(50, 10_000),
		recipient: "0x913dA4198E6bE1D5f5E4a40D0667f70C0B5430Eb",
		deadline: deadline,
	})
}

main()
