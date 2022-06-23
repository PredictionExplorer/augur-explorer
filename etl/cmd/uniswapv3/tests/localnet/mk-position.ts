import { ethers } from 'ethers'
import { BigNumber, BigNumberish, constants, Contract, ContractTransaction, utils, Wallet } from 'ethers'
import { Pool,NonfungiblePositionManager,Position } from '@uniswap/v3-sdk'
import { Token,Percent } from '@uniswap/sdk-core'
import { nearestUsableTick, TickMath } from '@uniswap/v3-sdk'
import { INonfungiblePositionManager as pmabi } from './INonfungiblePositionManager.json'
import { IUniswapV3PoolABI as poolabi } from './IUniswapV3Pool.json'
import bn from './bignumber.js'
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
interface Immutables {
	factory: string
	token0: string
	token1: string
	fee: number
	tickSpacing: number
	maxLiquidityPerTick: ethers.BigNumber
}
const poolAddress = "0x0E12de19803D02Ba47f52d8BAf5600c0Cfaf1E52"
const pkey = process.env["PRIVATE_KEY"]
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545')
const signer = new ethers.Wallet(pkey, provider);
const poolContract = new ethers.Contract(poolAddress, poolabi, signer)
const token0_addr = "0xb03cf72bC5a9A344AAC43534D664917927367487"
const token1_addr = "0x6226649431c4180a390f810bfD604b50EB68d9c5"
//const pool_fee = 500
//const price_str = "79228162514264337593543950336"
//const liquidity_str "10000"
//const tick = -887272
const chain_id = 1234
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

	const Token0 = new Token(chain_id, token0_addr, 18, 'WETH', 'Wrapped Eth')
	const Token1 = new Token(chain_id, token1_addr, 6, 'USDC', 'US Dollar')

	console.log("liquidity: ")
	console.log(state.liquidity.toString())
	console.log("tick: ")
	console.log(state.tick)
	const pool_ctrct = new Pool(
		Token0,
		Token1,
		immutables.fee,
		state.sqrtPriceX96.toString(),
		state.liquidity.toString(),
		state.tick
	)
	const position = new Position({
		pool: pool_ctrct,
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
