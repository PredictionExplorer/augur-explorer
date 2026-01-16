
const hre = require("hardhat");

async function main() {
  const contract_addr = "0x5FbDB2315678afecb367f032d93F642f64180aa3";
  const token_addr= "0x59b670e9fa9d0a427751af201d676719a970857b";
  const token_id_str = "13000002";
  const acct_idx = 0;

  const accounts = await hre.ethers.getSigners()
  const cosmicGame = await hre.ethers.getContractAt("CosmicGame",contract_addr);
  const dummyArtBlock = await hre.ethers.getContractAt("ERC721",token_addr);
  const tokenId = hre.ethers.BigNumber.from(token_id_str);

  await dummyArtBlock.setApprovalForAll(cosmicGame.address,true);
  let bid_price = await cosmicGame.getBidPrice();
  const signer = accounts[acct_idx];
  const msg = "ArtBlocks token "+token_id_str;
  await cosmicGame.connect(signer).bidAndDonateNFT(msg,hre.ethers.utils.getAddress(token_addr),tokenId,{value:bid_price});
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
