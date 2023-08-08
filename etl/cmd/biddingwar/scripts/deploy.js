const hre = require("hardhat");

async function main() {
	 async function mint_rwalk(a) {
	  tokenPrice = await randomWalkNFT.getMintPrice();
	  let tx = await randomWalkNFT.connect(a).mint({value: tokenPrice});
	  let receipt = await tx.wait();
	  let topic_sig = randomWalkNFT.interface.getEventTopic("MintEvent");
	  let log = receipt.logs.find(x=>x.topics.indexOf(topic_sig)>=0);
	  let parsed_log = randomWalkNFT.interface.parseLog(log);
	  let token_id = parsed_log.args[0]
	  return token_id;
	}
  [owner, addr1, addr2, addr3,...addrs] = await ethers.getSigners();
  //const [deployer] = await hre.ethers.getSigners();

  //console.log("Deploying contracts with the account:", deployer.address);

  //console.log("Account balance:", (await deployer.getBalance()).toString());

  const CosmicGame = await hre.ethers.getContractFactory("CosmicGame");
  const cosmicGame = await CosmicGame.deploy();
  await cosmicGame.deployed();
  console.log("CosmicGame address:", cosmicGame.address);

  const CosmicToken = await hre.ethers.getContractFactory("CosmicToken");
  const cosmicToken = await CosmicToken.deploy();
  cosmicToken.deployed();
  await cosmicToken.transferOwnership(cosmicGame.address);
  console.log("CosmicToken address:", cosmicToken.address);

  const CosmicSignature = await hre.ethers.getContractFactory("CosmicSignature");
  const cosmicSignature = await CosmicSignature.deploy(cosmicGame.address);
  cosmicSignature.deployed();
  console.log("CosmicSignature address:", cosmicSignature.address);

  const CosmicDAO = await hre.ethers.getContractFactory("CosmicDAO");
  const cosmicDAO = await CosmicDAO.deploy(cosmicToken.address);
  await cosmicDAO.deployed();
  console.log("CosmicDAO address", cosmicDAO.address);

  const CharityWallet = await hre.ethers.getContractFactory("CharityWallet");
  const charityWallet = await CharityWallet.deploy();
  charityWallet.deployed();
  await charityWallet.transferOwnership(cosmicDAO.address);
  console.log("CharityWallet address:", charityWallet.address);

  const RaffleWallet = await hre.ethers.getContractFactory("RaffleWallet");
  const raffleWallet = await RaffleWallet.deploy();
  raffleWallet.deployed();
  console.log("RaffleWallet address:", raffleWallet.address);

  const RandomWalkNFT = await hre.ethers.getContractFactory("RandomWalkNFT");
  const randomWalkNFT = await RandomWalkNFT.deploy();
  randomWalkNFT.deployed();
  console.log("randomWalkNFT address:", randomWalkNFT.address);

  const RandomWalkNFT2 = await hre.ethers.getContractFactory("RandomWalkNFT");
  const randomWalkNFT2 = await RandomWalkNFT2.deploy();
  randomWalkNFT2.deployed();
  console.log("randomWalkNFT2 address:", randomWalkNFT2.address);

  await cosmicGame.setTokenContract(cosmicToken.address);
  await cosmicGame.setNftContract(cosmicSignature.address);
  await cosmicGame.setCharity(charityWallet.address);
  await cosmicGame.setRaffleWallet(raffleWallet.address);
  await cosmicGame.setRandomWalk(randomWalkNFT.address);
  await cosmicGame.setActivationTime(0);

  console.log("Addresses set");
  console.log("INSERT INTO bw_contracts VALUES('"+
	  cosmicGame.address+"','"+
	  cosmicSignature.address+"','"+
	  cosmicToken.address+"','"+
	  cosmicDAO.address+"','"+
	  charityWallet.address+"','"+
	  raffleWallet.address+"','"+
	  randomWalkNFT.address+
	  "')"
  );

  let donationAmount = hre.ethers.utils.parseEther('10');
  await cosmicGame.donate({value: donationAmount});

  let bidPrice = await cosmicGame.getBidPrice();
  let prizeTime = await cosmicGame.timeUntilPrize();
  console.log("Donation complete");

  await cosmicGame.connect(addr1).bid("bid 1",{value: bidPrice.add(1000)}); // this works
  const contractBalance = await ethers.provider.getBalance(cosmicGame.address);

  let nanoSecondsExtra = await cosmicGame.nanoSecondsExtra();
  prizeTime = await cosmicGame.timeUntilPrize();

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 2",{value: bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
	console.log("prizeTime = "+prizeTime);
  let token_id = await mint_rwalk(owner);
  await  cosmicGame.connect(owner).bidWithRWLK(token_id,"bidWithRWLK");
  console.log("bid with Rwalk, tx , token_id="+token_id)
//	console.log(tx)

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr2).bid("bid 3",{value: bidPrice});

  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.sub(100).toNumber()]);
  await ethers.provider.send("evm_mine");

  await ethers.provider.send("evm_increaseTime", [100]);
  await ethers.provider.send("evm_mine");

  let prizeAmount = await cosmicGame.prizeAmount();
  let charityAmount = await cosmicGame.charityAmount();
  await cosmicGame.connect(addr2).claimPrize({gasLimit:1000000});
  let prizeAmount2 = await cosmicGame.prizeAmount();
  let expectedprizeAmount = prizeAmount.sub(charityAmount).div(2);

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 4",{value: bidPrice});

  prizeTime = await cosmicGame.timeUntilPrize();

  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  await ethers.provider.send("evm_mine");

  prizeAmount = await cosmicGame.prizeAmount();
  charityAmount = await cosmicGame.charityAmount();
  await cosmicGame.connect(addr1).claimPrize({gasLimit:1000000});
  prizeAmount2 = await cosmicGame.prizeAmount();
	console.log("prizeAmount2="+prizeAmount2)
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 5",{value:bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  tx = await cosmicGame.connect(addr1).claimPrize({gasLimit:1000000});
  receipt = await tx.wait();
  let topic_sig = cosmicSignature.interface.getEventTopic("MintEvent");
  log = receipt.logs.find(x=>x.topics.indexOf(topic_sig)>=0);
  parsed_log = cosmicSignature.interface.parseLog(log);
  token_id = parsed_log.args.tokenId;
  await cosmicSignature.connect(addr1).setTokenName(token_id,"name 0");

  await charityWallet.connect(addr1).send();

  tx =  {
     to: charityWallet.address,
     value: hre.ethers.utils.parseEther('4')
 };
  await addr2.sendTransaction(tx);
  await addr2.sendTransaction(tx);

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr3).bid("bid 6",{value:bidPrice});

  await ethers.provider.send("evm_mine");	// mine empty block as spacing
	
  await randomWalkNFT.connect(addr1).setApprovalForAll(cosmicGame.address, true);
  await randomWalkNFT.connect(addr2).setApprovalForAll(cosmicGame.address, true);
  await randomWalkNFT.connect(addr3).setApprovalForAll(cosmicGame.address, true);
	
  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr1);
  await cosmicGame.connect(addr1).bidAndDonateNFT("donated token_id="+token_id, randomWalkNFT.address, token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr2);
  await cosmicGame.connect(addr2).bidAndDonateNFT("me donated token_id="+token_id, randomWalkNFT.address, token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr3);
  await cosmicGame.connect(addr3).bidAndDonateNFT("me donated token_id="+token_id, randomWalkNFT.address, token_id, { value: bidPrice });

  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  await ethers.provider.send("evm_mine");
  tx = await cosmicGame.connect(addr3).claimPrize({gasLimit:1000000});
  topic_sig = raffleWallet.interface.getEventTopic("RaffleDepositEvent");

  await cosmicGame.connect(addr1).claimRaffleNFT();
  await cosmicGame.connect(addr1).claimRaffleNFT();
  await cosmicGame.connect(addr3).claimRaffleNFT();
  await cosmicGame.connect(addr3).claimRaffleNFT();
  await cosmicGame.connect(addr3).claimDonatedNFT(hre.ethers.BigNumber.from('0'))
  await cosmicGame.connect(addr3).claimDonatedNFT(hre.ethers.BigNumber.from('1'))

  receipt = await tx.wait();
  deposit_logs = receipt.logs.filter(x=>x.topics.indexOf(topic_sig)>=0);
  for (let i =0; i<deposit_logs.length; i++) {
	  let wlog = raffleWallet.interface.parseLog(deposit_logs[i]);
	  let winner_signer = raffleWallet.provider.getSigner(wlog.args.winner);
	  let deposit_id = wlog.args.depositId;
	  await raffleWallet.connect(owner).withdraw(deposit_id);
  }

  await ethers.provider.send("evm_mine");	// mine empty block as spacing

  await cosmicSignature.connect(addr2).setApprovalForAll(owner.address,true)
  await cosmicSignature.connect(owner).transferFrom(addr2.address,owner.address,hre.ethers.BigNumber.from('0'));
  await cosmicSignature.connect(addr1).setApprovalForAll(owner.address,true)
  await cosmicSignature.connect(owner).transferFrom(addr1.address,owner.address,hre.ethers.BigNumber.from('1'));

  await ethers.provider.send("evm_mine");	// mine empty block as spacing
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
