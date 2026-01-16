
import { ethers } from 'ethers'
import { Pool } from '@uniswap/v3-sdk'
import { Token } from '@uniswap/sdk-core'
import { abi as IUniswapV3FactoryABI } from './IUniswapV3Factory.json'

const pkey = process.env["PRIVATE_KEY"]
const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545')
const signer = new ethers.Wallet(pkey, provider);
const factoryAddress = '0xaf517E20601Df8d8584035EB895C02713bC1f3A4'
const token1 = '0xb03cf72bC5a9A344AAC43534D664917927367487'
const token2 = '0x6226649431c4180a390f810bfD604b50EB68d9c5'
const fee = 500;
const factoryContract = new ethers.Contract(factoryAddress, abi, signer)

async function crPool(
	tokens: [string, string],
	feeAmount:number
) {
	const create = factoryContract.createPool(tokens[0], tokens[1], feeAmount)
}

crPool([token1,token2],fee).then((result) => {
  console.log('tx submitted')
})
