import { ethers } from 'ethers'
import { Pool } from '@uniswap/v3-sdk'
import { CurrencyAmount, Token, TradeType } from '@uniswap/sdk-core'
import { abi as IUniswapV3PoolABI } from '@uniswap/v3-core/artifacts/contracts/interfaces/IUniswapV3Pool.sol/IUniswapV3Pool.json'
import { Route } from '@uniswap/v3-sdk'
import { Trade } from '@uniswap/v3-sdk'
import { abi as QuoterABI } from '@uniswap/v3-periphery/artifacts/contracts/lens/Quoter.sol/Quoter.json'

const pkey = process.env["PRIVATE_KEY"]
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545')
const signer = new ethers.Wallet(pkey, provider);
const chain_id = 1234

const poolAddress = '0xAe9a1Df527E36DE6EBa251eA4FBAfC897e1D7E9A'
const poolContract = new ethers.Contract(poolAddress, IUniswapV3PoolABI, signer)
const quoterAddress = '0xB8da5FA6c6F9b55F1c9fA09E26A24D3675Dbc36E'
const quoterContract = new ethers.Contract(quoterAddress, QuoterABI, signer)

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

	const TokenA = new Token(chain_id, immutables.token0, 6, 'USDC', 'USD Coin')
	const TokenB = new Token(chain_id, immutables.token1, 18, 'WETH', 'Wrapped Ether')
	console.log("Token A",TokenA)
	console.log("Token b",TokenB)
	// create an instance of the pool object for the given pool
	const poolExample = new Pool(
		TokenA,
		TokenB,
		immutables.fee,
		state.sqrtPriceX96.toString(), 
		state.liquidity.toString(),
		state.tick
	)

	const amountIn = 500

	// call the quoter contract to determine the amount out of a swap, given an amount in
	const quotedAmountOut = await quoterContract.callStatic.quoteExactInputSingle(
		immutables.token0,
		immutables.token1,
		immutables.fee,
		amountIn.toString(),
		0
	)

	const swapRoute = new Route([poolExample], TokenA, TokenB)
	console.log("route:")
	console.log(swapRoute)

	const uncheckedTradeExample = await Trade.createUncheckedTrade({
		route: swapRoute,
		inputAmount: CurrencyAmount.fromRawAmount(TokenA, amountIn.toString()),
		outputAmount: CurrencyAmount.fromRawAmount(TokenB, quotedAmountOut.toString()),
		tradeType: TradeType.EXACT_INPUT,
	})

	console.log('The quoted amount out is', quotedAmountOut.toString())
	console.log('The unchecked trade object is', uncheckedTradeExample)
}

main()
