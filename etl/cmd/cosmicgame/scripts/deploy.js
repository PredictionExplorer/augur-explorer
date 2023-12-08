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
  [owner, addr1, addr2, addr3, addr4, addr5,...addrs] = await ethers.getSigners();

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
  console.log("CharityWallet address:", charityWallet.address);

  const RaffleWallet = await hre.ethers.getContractFactory("RaffleWallet");
  const raffleWallet = await RaffleWallet.deploy(cosmicGame.address);
  raffleWallet.deployed();
  console.log("RaffleWallet address:", raffleWallet.address);

  const StakingWallet = await hre.ethers.getContractFactory("StakingWallet");
  stakingWallet = await StakingWallet.deploy(cosmicSignature.address,cosmicGame.address,charityWallet.address);
  await stakingWallet.deployed();
  console.log("StakingWallet address:", stakingWallet.address);

  const MarketingWallet = await hre.ethers.getContractFactory("MarketingWallet");
  marketingWallet = await MarketingWallet.deploy(cosmicToken.address);
  await marketingWallet.deployed();
  console.log("MarketingWallet address:", marketingWallet.address);

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
  await cosmicGame.setMarketingWallet(marketingWallet.address);
  await cosmicGame.setStakingWallet(stakingWallet.address);
  await cosmicGame.setRandomWalk(randomWalkNFT.address);
  await cosmicGame.setActivationTime(0);

  console.log("Addresses set");
  console.log("INSERT INTO cg_contracts VALUES('"+
	  cosmicGame.address+"','"+
	  cosmicSignature.address+"','"+
	  cosmicToken.address+"','"+
	  cosmicDAO.address+"','"+
	  charityWallet.address+"','"+
	  raffleWallet.address+"','"+
	  randomWalkNFT.address+"','"+
	  stakingWallet.address+"','"+
	  marketingWallet.address+
	  "')"
  );

  let donationAmount = hre.ethers.utils.parseEther('10');
  await cosmicGame.donate({value: donationAmount});

  let prizeTime = await cosmicGame.timeUntilPrize();
  console.log("Donation complete");

  const contractBalance = await ethers.provider.getBalance(cosmicGame.address);
  let bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 1",-1,{value: bidPrice.add(1000)}); // this works
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 1",-1,{value: bidPrice.add(1000)}); // this works

  let nanoSecondsExtra = await cosmicGame.nanoSecondsExtra();
  prizeTime = await cosmicGame.timeUntilPrize();

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 2",-1,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 2",-1,{value: bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
  let token_id = await mint_rwalk(owner);
  bidPrice = await cosmicGame.getBidPrice();
  await  cosmicGame.connect(owner).bid("bidWithRWLK",token_id,{value:bidPrice});

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr3).bid("bid 3",-1,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr3).bid("bid 3",-1,{value: bidPrice});

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr4).bid("",-1,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr4).bid("",-1,{value: bidPrice});

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr5).bid("",-1,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr5).bid("",-1,{value: bidPrice});

  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.sub(100).toNumber()]);
  await ethers.provider.send("evm_mine");

  await ethers.provider.send("evm_increaseTime", [100]);
  await ethers.provider.send("evm_mine");

  let prizeAmount = await cosmicGame.prizeAmount();
  let charityAmount = await cosmicGame.charityAmount();
  await cosmicGame.connect(addr5).claimPrize({gasLimit:30000000});
  let prizeAmount2 = await cosmicGame.prizeAmount();
  let expectedprizeAmount = prizeAmount.sub(charityAmount).div(2);

  let topic_sig = stakingWallet.interface.getEventTopic("StakeActionEvent");
  let max_ts = 0;
  let ts = await cosmicSignature.totalSupply();
  let rn = await cosmicGame.roundNum();
  for (let i =0; i<ts.toNumber(); i++) {
    let ownr = await cosmicSignature.ownerOf(i)
	let owner_signer = cosmicGame.provider.getSigner(ownr);
    await cosmicSignature.connect(owner_signer).setApprovalForAll(stakingWallet.address, true);
    tx = await stakingWallet.connect(owner_signer).stake(i);
    receipt = await tx.wait();
    log = receipt.logs.find(x=>x.topics.indexOf(topic_sig)>=0);
    parsed_log = stakingWallet.interface.parseLog(log);
    if (max_ts < parsed_log.args.unstakeTime.toNumber()) {
		max_ts = parsed_log.args.unstakeTime.toNumber();
		console.log("max ts changed to "+max_ts)
	}
  }
  console.log("max ts after loop = "+max_ts);

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 4",-1,{value: bidPrice});

  prizeTime = await cosmicGame.timeUntilPrize();

  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  await ethers.provider.send("evm_mine");

  prizeAmount = await cosmicGame.prizeAmount();
  charityAmount = await cosmicGame.charityAmount();
  await cosmicGame.connect(addr1).claimPrize({gasLimit:3000000});
  prizeAmount2 = await cosmicGame.prizeAmount();
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid("bid 5",-1,{value:bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  tx = await cosmicGame.connect(addr1).claimPrize({gasLimit:3000000});
  receipt = await tx.wait();
  topic_sig = cosmicSignature.interface.getEventTopic("MintEvent");
  log = receipt.logs.find(x=>x.topics.indexOf(topic_sig)>=0);
  parsed_log = cosmicSignature.interface.parseLog(log);
  token_id = parsed_log.args.tokenId;
  await cosmicSignature.connect(addr1).setTokenName(token_id,"name 0");
  await cosmicSignature.connect(addr1).setTokenName(token_id,"name after 0");

  await charityWallet.connect(addr1).send();

  tx =  {
     to: charityWallet.address,
     value: hre.ethers.utils.parseEther('4')
 };
  await addr2.sendTransaction(tx);
  await addr2.sendTransaction(tx);

  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr3).bid("bid 6",-1,{value:bidPrice});
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
	
  await randomWalkNFT.connect(addr1).setApprovalForAll(cosmicGame.address, true);
  await randomWalkNFT.connect(addr2).setApprovalForAll(cosmicGame.address, true);
  await randomWalkNFT.connect(addr3).setApprovalForAll(cosmicGame.address, true);
	
  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr1);
  await cosmicGame.connect(addr1).bidAndDonateNFT("donated token_id="+token_id,-1, randomWalkNFT.address, token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr2);
  await cosmicGame.connect(addr2).bidAndDonateNFT("me donated token_id="+token_id,-1, randomWalkNFT.address, token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr3);
  await cosmicGame.connect(addr3).bidAndDonateNFT("me donated token_id="+token_id,-1, randomWalkNFT.address, token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr3);
  await cosmicGame.connect(addr3).bidAndDonateNFT("me donated token_id="+token_id,-1, randomWalkNFT.address, token_id, { value: bidPrice });

  await cosmicGame.connect(addr3).bidWithCST("bid using ERC20 token");
  await ethers.provider.send("evm_mine");
  await cosmicGame.connect(addr3).bidWithCST("bid using ERC20 token");
  await ethers.provider.send("evm_mine");

  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  await ethers.provider.send("evm_mine");
  tx = await cosmicGame.connect(addr3).claimPrize({gasLimit:3000000});
  receipt = await tx.wait();
  await cosmicGame.connect(addr3).claimDonatedNFT(hre.ethers.BigNumber.from('0'))
  await cosmicGame.connect(addr3).claimDonatedNFT(hre.ethers.BigNumber.from('1'))
  topic_sig = raffleWallet.interface.getEventTopic("RaffleDepositEvent");
  deposit_logs = receipt.logs.filter(x=>x.topics.indexOf(topic_sig)>=0);
  let withdrawal_done = [];
  for (let i =0; i<deposit_logs.length; i++) {
	  let wlog = raffleWallet.interface.parseLog(deposit_logs[i]);
	  let winner_signer = raffleWallet.provider.getSigner(wlog.args.winner);
	  if (typeof withdrawal_done[wlog.args.winner] === 'undefined' ) {
		  await raffleWallet.connect(winner_signer).withdraw();
	      withdrawal_done[wlog.args.winner]=1;
	  } else {
			// skip
	  }
  }

  await marketingWallet.send(hre.ethers.utils.parseEther('7'),addr1.address);
  await marketingWallet.send(hre.ethers.utils.parseEther('5'),addr2.address);
 
  await ethers.provider.send("evm_setNextBlockTimestamp", [max_ts]);
  await ethers.provider.send("evm_mine");
  let num_actions;
  num_actions = await stakingWallet.numStakeActions();
  for (let i = 0; i < num_actions.toNumber(); i++) {
    let action_rec = await stakingWallet.stakedNFTs(i);
	let ownr = action_rec.owner;
	let owner_signer = cosmicGame.provider.getSigner(ownr);
	await stakingWallet.connect(owner_signer).unstake(i);
  }
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  num_actions  = await stakingWallet.numStakeActions();
  for (let i =0; i<num_actions.toNumber(); i++) {
    let action_rec = await stakingWallet.stakedNFTs(i);
	let ownr = action_rec.owner;
	let num_deposits = await stakingWallet.numETHDeposits();
	let owner_signer = stakingWallet.provider.getSigner(ownr);
    for (let j = 0; j < num_deposits.toNumber(); j++) {
		let deposit_rec = await stakingWallet.ETHDeposits(j);
        await stakingWallet.connect(owner_signer).claimReward(i,j);
      }
  } 
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
	
  // generate one deposit to charity and not to Staking Wallet
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr3).bid("bid 3",-1,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr3).bid("bid 3",-1,{value: bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  tx = await cosmicGame.connect(addr3).claimPrize({gasLimit:3000000});

  await ethers.provider.send("evm_mine");	// mine empty block as spacing


  for (let i =0; i<ts.toNumber(); i++) {
    let ownr = await cosmicSignature.ownerOf(i)
	let owner_signer = cosmicGame.provider.getSigner(ownr);
    await stakingWallet.connect(owner_signer).stake(i);
  }

  await cosmicToken.connect(addr1).approve(cosmicGame.address,hre.ethers.utils.parseEther('10000000'));
  await cosmicToken.connect(addr2).approve(cosmicGame.address,hre.ethers.utils.parseEther('10000000'));
  await cosmicToken.connect(addr3).approve(cosmicGame.address,hre.ethers.utils.parseEther('10000000'));
  await cosmicToken.connect(addr4).approve(cosmicGame.address,hre.ethers.utils.parseEther('10000000'));
  await cosmicGame.connect(addr1).bidWithCST("CST bid addr1")
  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [prizeTime.toNumber()]);
  await ethers.provider.send("evm_mine");
  tx = await cosmicGame.connect(addr1).claimPrize({gasLimit:3000000});

  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
