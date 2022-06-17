//var math = require("/usr/local/lib/node_modules/mathjs")
import { ethers } from 'ethers'
import { BigNumber, BigNumberish, constants, Contract, ContractTransaction, utils, Wallet } from 'ethers'
import { Pool } from '@uniswap/v3-sdk'
import { Token } from '@uniswap/sdk-core'
import { IUniswapV3PoolABI as abi } from './IUniswapV3Pool.json'
import bn from 'bignumber.js'

const pkey = process.env["PRIVATE_KEY"]
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545')
const signer = new ethers.Wallet(pkey, provider);
const poolAddress = '0xaf517E20601Df8d8584035EB895C02713bC1f3A4'
const poolContract = new ethers.Contract(poolAddress, abi, signer)

export function encodePriceSqrt(reserve1: BigNumberish, reserve0: BigNumberish): BigNumber {
  return BigNumber.from(
    new bn(reserve1.toString())
      .div(reserve0.toString())
      .sqrt()
      .multipliedBy(new bn(2).pow(96))
      .integerValue(3)
      .toString()
  )
}

async function initPool(
	sqrtPrice
) {
	const init = poolContract.initialize(sqrtPrice)
}
var sqrtPrice = encodePriceSqrt(1,1)
//var sqrtPrice = math.sqrt(1000) * 2 ** 96
initPool(sqrtPrice).then((result) => {
  console.log('tx submitted')
  console.log(result)
})
