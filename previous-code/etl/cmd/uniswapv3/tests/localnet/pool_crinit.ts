import { ethers } from 'ethers'
import { BigNumber, BigNumberish, constants, Contract, ContractTransaction, utils, Wallet } from 'ethers'
import { Pool } from '@uniswap/v3-sdk'
import { Token } from '@uniswap/sdk-core'
import { IUniswapV3PoolABI as abi } from './IUniswapV3Pool.json'
import bn from './bignumber.js'

const pkey = process.env["PRIVATE_KEY"]
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545')
const signer = new ethers.Wallet(pkey, provider);
const nftpmAddress = '0x7a8f8e48D4CC990BBD8E088FB027e850486e8e0C'
const poolContract = new ethers.Contract(poolAddress, abi, signer)

/*
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
*/
function encodePriceSqrt(reserve1: BigNumberish, reserve0: BigNumberish): string {
	console.log(reserve0)
	console.log(reserve1)

  	var num =  new bn(reserve1.toString())
	num.div(reserve0.toString())
	num = num.sqrt()
	var m = new bn(2).pow(96)
    num = num.multipliedBy(m)
    num = num.integerValue(3)
	var out = num.toFixed()
	console.log("sqrtpriceX96:")
	console.log(out)
	return out
}
async function initPool(
	sqrtPrice: BigNumber
) {
	const init = poolContract.initialize(sqrtPrice)
}
var sqrtPriceStr = encodePriceSqrt(1,1)
var p = BigNumber.from(sqrtPriceStr)
console.log("sqrtPrice sent in tx")
console.log(p)
initPool(p).then((result) => {
  console.log('tx submitted')
  console.log(result)
})

