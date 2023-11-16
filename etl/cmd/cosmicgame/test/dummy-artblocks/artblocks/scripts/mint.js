const hre = require("hardhat");

async function main() {
  const contract_addr = "0x4EEEdD240C98Ef89Ffc25616bee0B7Ac8Cf82Dd1";
  //const owner_addr = '0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266';
  const acct_idx = 0;

  const accounts = await hre.ethers.getSigners()
  const dummyArtBlock = await hre.ethers.getContractAt("DummyArtBlocks",contract_addr);

  const owner = accounts[acct_idx];
 // const owner = await hre.ethers.provider.getSigner(await hre.ethers.utils.getAddress(owner_addr));
  var tokenId = await dummyArtBlock.curTokenId();
  await dummyArtBlock.connect(owner).mint(owner.address);
  var tokenUri = await dummyArtBlock.tokenURI(tokenId)
  var ownerof = await dummyArtBlock.ownerOf(tokenId);
  console.log("minted tokenId = "+tokenId.toString());
  console.log("owner = "+ownerof.toString());
  console.log("URI = "+tokenUri);
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
