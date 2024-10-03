//const hre = require("hardhat");
const { basicDeployment } = require("./Deploy.js");
const bidParamsEncoding = { 
	type: 'tuple(string,int256)',
	name: 'bidparams',
	components: [
		{name: 'msg', type: 'string'},
		{name: 'rwalk',type: 'int256'},
	]
}; 
async function main() {
	 async function mint_rwalk(a) {
	  tokenPrice = await randomWalkNFT.getMintPrice();
	  let tx = await randomWalkNFT.connect(a).mint({value: tokenPrice});
	  let receipt = await tx.wait();
	  let topic_sig = randomWalkNFT.interface.getEvent("MintEvent").topicHash;
	  let log = receipt.logs.find(x=>x.topics.indexOf(topic_sig)>=0);
	  let parsed_log = randomWalkNFT.interface.parseLog(log);
	  let token_id = parsed_log.args[0]
	  return token_id;
	}
  [owner, addr1, addr2, addr3, addr4, addr5,...addrs] = await ethers.getSigners();
/*
  const CosmicGame = await hre.ethers.getContractFactory("CosmicGame");
  const cosmicGame = await CosmicGame.deploy();
  await cosmicGame.waitForDeployment();
  console.log("CosmicGame address:", cosmicGame.address);

  const CosmicToken = await hre.ethers.getContractFactory("CosmicToken");
  const cosmicToken = await CosmicToken.deploy();
  cosmicToken.waitForDeployment();
  await cosmicToken.transferOwnership(cosmicGame.address);
  console.log("CosmicToken address:", cosmicToken.address);

  const CosmicSignature = await hre.ethers.getContractFactory("CosmicSignature");
  const cosmicSignature = await CosmicSignature.deploy(cosmicGame.address);
  cosmicSignature.waitForDeployment();
  console.log("CosmicSignature address:", cosmicSignature.address);

  const CosmicDAO = await hre.ethers.getContractFactory("CosmicDAO");
  const cosmicDAO = await CosmicDAO.deploy(cosmicToken.address);
  await cosmicDAO.waitForDeployment();
  console.log("CosmicDAO address", cosmicDAO.address);

  const CharityWallet = await hre.ethers.getContractFactory("CharityWallet");
  const charityWallet = await CharityWallet.deploy();
  charityWallet.waitForDeployment();
  console.log("CharityWallet address:", charityWallet.address);

  const RaffleWallet = await hre.ethers.getContractFactory("RaffleWallet");
  const raffleWallet = await RaffleWallet.deploy(cosmicGame.address);
  raffleWallet.waitForDeployment();
  console.log("RaffleWallet address:", raffleWallet.address);

  const MarketingWallet = await hre.ethers.getContractFactory("MarketingWallet");
  marketingWallet = await MarketingWallet.deploy(cosmicToken.address);
  await marketingWallet.waitForDeployment();
  console.log("MarketingWallet address:", marketingWallet.address);

  const RandomWalkNFT = await hre.ethers.getContractFactory("RandomWalkNFT");
  const randomWalkNFT = await RandomWalkNFT.deploy();
  randomWalkNFT.waitForDeployment();
  console.log("randomWalkNFT address:", randomWalkNFT.address);

  const StakingWalletCST = await hre.ethers.getContractFactory("StakingWalletCST");
  stakingWalletCST = await StakingWalletCST.deploy(cosmicSignature.address,cosmicGame.address,charityWallet.address);
  await stakingWalletCST.waitForDeployment();
  console.log("StakingWalletCST address:", stakingWalletCST.address);

  const StakingWalletRWalk = await hre.ethers.getContractFactory("StakingWalletRWalk");
  stakingWalletRWalk = await StakingWalletRWalk.deploy(randomWalkNFT.address,cosmicGame.address);
  await stakingWalletRWalk.waitForDeployment();
  console.log("StakingWalletRWalk address:", stakingWalletRWalk.address);

  const RandomWalkNFT2 = await hre.ethers.getContractFactory("RandomWalkNFT");
  const randomWalkNFT2 = await RandomWalkNFT2.deploy();
  randomWalkNFT2.waitForDep();
  console.log("randomWalkNFT2 address:", randomWalkNFT2.address);

  const BusinessLogic = await ethers.getContractFactory("BusinessLogic");
  bLogic = await BusinessLogic.deploy();
  await bLogic.deployed();
  */
    const {
        cosmicGameProxy,
        cosmicToken,
        cosmicSignature,
        charityWallet,
        cosmicDAO,
        raffleWallet,
        randomWalkNFT,
        stakingWalletCST,
        stakingWalletRWalk,
        marketingWallet,
        cosmicGame,
    } = await basicDeployment(
        owner,
        '',
        0,
        '',
        false,
        true
    );  
/*
  await cosmicGame.setTokenContract(cosmicToken.address);
  await cosmicGame.setNftContract(cosmicSignature.address);
  await cosmicGame.setCharity(charityWallet.address);
  await cosmicGame.setBusinessLogicContract(bLogic.address);
  await cosmicGame.setRaffleWallet(raffleWallet.address);
  await cosmicGame.setMarketingWallet(marketingWallet.address);
  await cosmicGame.setStakingWalletCST(stakingWalletCST.address);
  await cosmicGame.setStakingWalletRWalk(stakingWalletRWalk.address);
  await cosmicGame.setRandomWalk(randomWalkNFT.address);
  await cosmicGame.setActivationTime(0);
  await cosmicGame.setRuntimeMode()
*/
  console.log("Addresses set");
  console.log("INSERT INTO cg_contracts VALUES('"+
	  await cosmicGame.getAddress()+"','"+
	  await cosmicSignature.getAddress()+"','"+
	  await cosmicToken.getAddress()+"','"+
	  await cosmicDAO.getAddress()+"','"+
	  await charityWallet.getAddress()+"','"+
	  await raffleWallet.getAddress()+"','"+
	  await randomWalkNFT.getAddress()+"','"+
	  await stakingWalletCST.getAddress()+"','"+
	  await stakingWalletRWalk.getAddress()+"','"+
	  await marketingWallet.getAddress()+"','"+
	  await cosmicGame.getAddress()+
	  "')"
  );

  let donationAmount = hre.ethers.parseEther('10');
  await cosmicGame.donate({value: donationAmount});
  let donationData = '{"version":1,"title":"Hardhat donation","message":"Donation from HardHat","url":"http://hardhat.org"}';
  await cosmicGame.donateWithInfo(donationData,{value: hre.ethers.parseEther('6')});
  donationData = '{"version":1,"title":"ArtBlocks donation","message":"ArtBlocks offers a platform for creators, buyers and sellers of digital assets and any non-digital products, services and/or benefits to be furnished by or on behalf of sellers in connection with such sales","url":"https://www.artblocks.io"}';

  await cosmicGame.connect(addr2).donateWithInfo(donationData,{value: hre.ethers.parseEther('6')});

  for (let i=0; i<5; i++) {
	let token_id = await mint_rwalk(addr1);
    await randomWalkNFT.connect(addr1).setApprovalForAll(await stakingWalletRWalk.getAddress(), true);
	await stakingWalletRWalk.connect(addr1).stake(token_id);
  }
  for (let i=0; i<5; i++) {
	let token_id = await mint_rwalk(addr2);
    await randomWalkNFT.connect(addr2).setApprovalForAll(await stakingWalletRWalk.getAddress(), true);
	await stakingWalletRWalk.connect(addr2).stake(token_id);
  }
  for (let i=0; i<50; i++) {
	let token_id = await mint_rwalk(addr3);
    await randomWalkNFT.connect(addr3).setApprovalForAll(await stakingWalletRWalk.getAddress(), true);
	await stakingWalletRWalk.connect(addr3).stake(token_id);
  }

  let prizeTime = await cosmicGame.timeUntilPrize();
  console.log("Donation complete");

  var bidParams = {msg:'bid 1',rwalk:-1};
  let params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  const contractBalance = await ethers.provider.getBalance(await cosmicGame.getAddress());
  let bidPrice = await cosmicGame.getBidPrice();
  await cosmicGame.connect(addr1).bid(params,{value: bidPrice + 1000n}); // this works
  bidPrice = await cosmicGame.getBidPrice();
  bdParams = {msg:'bid 1',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr1).bid(params,{value: bidPrice + 1000n}); // this works

  let nanoSecondsExtra = await cosmicGame.nanoSecondsExtra();
  prizeTime = await cosmicGame.timeUntilPrize();

  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 2',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr1).bid(params,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 2',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr1).bid(params,{value: bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
  let token_id = await mint_rwalk(owner);
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bidWithRWLK',rwalk:Number(token_id)};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await  cosmicGame.connect(owner).bid(params,{value:bidPrice});

  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 3',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr3).bid(params,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 3',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr3).bid(params,{value: bidPrice});

  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr4).bid(params,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr4).bid(params,{value: bidPrice});

  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr5).bid(params,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr5).bid(params,{value: bidPrice});

  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [Number(prizeTime) - 100]);
  await ethers.provider.send("evm_mine");

  await ethers.provider.send("evm_increaseTime", [100]);
  await ethers.provider.send("evm_mine");

  let prizeAmount = await cosmicGame.prizeAmount();
  let charityAmount = await cosmicGame.charityAmount();
  await cosmicGame.connect(addr5).claimPrize({gasLimit:30000000});
  let prizeAmount2 = await cosmicGame.prizeAmount();
  let expectedprizeAmount = (prizeAmount-charityAmount) / 2n;

  let topic_sig = stakingWalletCST.interface.getEvent("StakeActionEvent").topicHash;
  let ts = await cosmicSignature.totalSupply();
  let rn = await cosmicGame.roundNum();
  for (let i =0; i< Number(ts); i++) {
	let tx;
    let ownr = await cosmicSignature.ownerOf(i)
	let owner_signer = await hre.ethers.getSigner(ownr);
    await cosmicSignature.connect(owner_signer).setApprovalForAll(await stakingWalletCST.getAddress(), true);
	tx = await stakingWalletCST.connect(owner_signer).stake(i);
  }
  let oldTotalSupply = ts;
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 4',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr1).bid(params,{value: bidPrice});

  prizeTime = await cosmicGame.timeUntilPrize();

  await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
  await ethers.provider.send("evm_mine");

  prizeAmount = await cosmicGame.prizeAmount();
  charityAmount = await cosmicGame.charityAmount();
  await cosmicGame.connect(addr1).claimPrize({gasLimit:3000000});
  prizeAmount2 = await cosmicGame.prizeAmount();
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 5',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr1).bid(params,{value:bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
  tx = await cosmicGame.connect(addr1).claimPrize({gasLimit:5000000});
  receipt = await tx.wait();
  topic_sig = cosmicSignature.interface.getEvent("MintEvent").topicHash;
  log = receipt.logs.find(x=>x.topics.indexOf(topic_sig)>=0);
  parsed_log = cosmicSignature.interface.parseLog(log);
  token_id = parsed_log.args.tokenId;
  await cosmicSignature.connect(addr1).setTokenName(token_id,"name 0");
  await cosmicSignature.connect(addr1).setTokenName(token_id,"name after 0");

  await charityWallet.connect(addr1).send();

  tx =  {
     to: await charityWallet.getAddress(),
     value: hre.ethers.parseEther('4')
 };
  await addr2.sendTransaction(tx);
  await addr2.sendTransaction(tx);

  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 6',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr3).bid(params,{value:bidPrice});
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
	
  await randomWalkNFT.connect(addr1).setApprovalForAll(await cosmicGame.getAddress(), true);
  await randomWalkNFT.connect(addr2).setApprovalForAll(await cosmicGame.getAddress(), true);
  await randomWalkNFT.connect(addr3).setApprovalForAll(await cosmicGame.getAddress(), true);
	
  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr1);
  bidParams = {msg:'donated token_id='+token_id,rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr1).bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr2);
  bidParams = {msg:'me donated token_id='+token_id,rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr2).bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr3);
  bidParams = {msg:'me donated token_id='+token_id,rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr3).bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, { value: bidPrice });

  bidPrice = await cosmicGame.getBidPrice();
  token_id = await mint_rwalk(addr3);
  bidParams = {msg:'me donated token_id='+token_id,rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr3).bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, { value: bidPrice });

  await cosmicGame.connect(addr3).bidWithCST("bid using ERC20 token");
  await ethers.provider.send("evm_mine");
  await cosmicGame.connect(addr3).bidWithCST("bid using ERC20 token");
  await ethers.provider.send("evm_mine");

  await cosmicGame.connect(owner).prepareMaintenance();

  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
  await ethers.provider.send("evm_mine");
  tx = await cosmicGame.connect(addr3).claimPrize({gasLimit:3000000});
  receipt = await tx.wait();

  await cosmicGame.connect(owner).setCharity(await charityWallet.getAddress());
  await cosmicGame.connect(owner).setRandomWalk(await randomWalkNFT.getAddress());
  await cosmicGame.connect(owner).setRaffleWallet(await raffleWallet.getAddress());
  await cosmicGame.connect(owner).setStakingWalletCST(await stakingWalletCST.getAddress());
  await cosmicGame.connect(owner).setStakingWalletRWalk(await stakingWalletRWalk.getAddress());
  await cosmicGame.connect(owner).setMarketingWallet(await marketingWallet.getAddress());
  await cosmicGame.connect(owner).setNumRaffleETHWinnersBidding(4);
  await cosmicGame.connect(owner).setNumRaffleNFTWinnersBidding(6);
  await cosmicGame.connect(owner).setNumRaffleNFTWinnersStakingRWalk(3);
  await cosmicGame.connect(owner).setPrizePercentage(30)
  await cosmicGame.connect(owner).setCharityPercentage(5);
  await cosmicGame.connect(owner).setRafflePercentage(6);
  await cosmicGame.connect(owner).setCharity(addr3.address);
  await cosmicGame.connect(owner).setCharity(await charityWallet.getAddress());
  await cosmicGame.connect(owner).setStakingPercentage(9);
  await cosmicGame.connect(owner).setTokenContract(await cosmicToken.getAddress());
  await cosmicGame.connect(owner).setNftContract(await cosmicSignature.getAddress());
  //let implAddr = await cosmicGameProxy.runner.provider.getStorage(await cosmicGame.getAddress(),'0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc');
 // let implAddr = await ethers.provider.getStorage(await cosmicGame.getAddress(),'0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc');
 // console.log("addr = "+implAddr);
 // await cosmicGame.connect(owner).upgradeTo(implAddr);
  await cosmicGame.connect(owner).setStakingWalletCST(await stakingWalletCST.getAddress());
  await cosmicGame.connect(owner).setStakingWalletRWalk(await stakingWalletRWalk.getAddress());
  await cosmicGame.connect(owner).setErc20RewardMultiplier(999);
  let tmp = await cosmicGame.timeIncrease();
  await cosmicGame.connect(owner).setTimeIncrease(tmp);
  tmp = await cosmicGame.connect(owner).timeoutClaimPrize()
  await cosmicGame.connect(owner).setTimeoutClaimPrize(tmp);
  tmp = await cosmicGame.priceIncrease();
  await cosmicGame.connect(owner).setPriceIncrease(tmp);
  tmp = await cosmicGame.nanoSecondsExtra();
  await cosmicGame.connect(owner).setNanoSecondsExtra(tmp);
  tmp = await cosmicGame.initialSecondsUntilPrize();
  await cosmicGame.connect(owner).setInitialSecondsUntilPrize(tmp);
  tmp = await cosmicGame.initialBidAmountFraction();
  await cosmicGame.connect(owner).updateInitialBidAmountFraction(tmp);
  tmp = await cosmicGame.activationTime();
  await cosmicGame.connect(owner).setActivationTime(tmp);

  await cosmicGame.connect(owner).setRuntimeMode();

  await cosmicGame.connect(addr3).claimDonatedNFT(0n)
  await cosmicGame.connect(addr3).claimDonatedNFT(1n)
  topic_sig = raffleWallet.interface.getEvent("RaffleDepositEvent").topicHash;
  deposit_logs = receipt.logs.filter(x=>x.topics.indexOf(topic_sig)>=0);
  let withdrawal_done = [];
  for (let i =0; i<deposit_logs.length; i++) {
	  let wlog = raffleWallet.interface.parseLog(deposit_logs[i]);
	  let winner_signer = await hre.ethers.getSigner(wlog.args.winner);
	  if (typeof withdrawal_done[wlog.args.winner] === 'undefined' ) {
		  await raffleWallet.connect(winner_signer).withdraw();
	      withdrawal_done[wlog.args.winner]=1;
	  } else {
			// skip
	  }
  }
 
  await ethers.provider.send("evm_mine");

  donationData = '{"version":1,"title":"EF donation","message":"Ethereum Foundation is a non-profit and part of a community of organizations and people working to fund protocol development, grow the ecosystem, and advocate for Ethereum.","url":"http://ethereum.org/en"}';
  await cosmicGame.donateWithInfo(donationData,{value: hre.ethers.parseEther('9')});
  await cosmicGame.donateWithInfo(donationData,{value: hre.ethers.parseEther('8')});

  await marketingWallet.send(hre.ethers.parseEther('7'),addr1.address);
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await marketingWallet.send(hre.ethers.parseEther('7'),addr3.address);
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await marketingWallet.send(hre.ethers.parseEther('2'),addr2.address);
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await marketingWallet.send(hre.ethers.parseEther('6'),addr1.address);
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await marketingWallet.send(hre.ethers.parseEther('5'),addr2.address);
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await marketingWallet.send(hre.ethers.parseEther('5'),addr2.address);
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await marketingWallet.send(hre.ethers.parseEther('5'),addr3.address);
  await marketingWallet.send(hre.ethers.parseEther('1'),addr4.address);
  await marketingWallet.send(hre.ethers.parseEther('11'),addr1.address);
  let num_actions;
  num_actions = await stakingWalletCST.numStakeActions();
  for (let i = 0; i < Number(num_actions); i++) {
    let action_rec = (await stakingWalletCST.stakeActions(i)).toObject();
	let ownr = action_rec.nftOwner;
	let owner_signer = await hre.ethers.getSigner(ownr);
	await stakingWalletCST.connect(owner_signer).unstake(i);
  }
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  /* Pending for refactoring in solidity
  num_actions  = await stakingWalletCST.numStakeActions();
  for (let i =0; i<Number(num_actions); i++) {
    let action_rec = (await stakingWalletCST.stakeActions(i)).toObject();
	let ownr = action_rec.nftOwner;
	let num_deposits = await stakingWalletCST.numETHDeposits();
	let owner_signer = await hre.ethers.getSigner(ownr);
    for (let j = 0; j < Number(num_deposits); j++) {
		let deposit_rec = await stakingWalletCST.ETHDeposits(j);
        await stakingWalletCST.connect(owner_signer).claimReward(i,j);
      }
  } 
  */
  await ethers.provider.send("evm_mine");	// mine empty block as spacing
	
  // generate one deposit to charity and not to Staking Wallet
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 3',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr3).bid(params,{value: bidPrice});
  bidPrice = await cosmicGame.getBidPrice();
  bidParams = {msg:'bid 3',rwalk:-1};
  params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams])
  await cosmicGame.connect(addr3).bid(params,{value: bidPrice});
  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
  tx = await cosmicGame.connect(addr3).claimPrize({gasLimit:3000000});

  await ethers.provider.send("evm_mine");	// mine empty block as spacing

  ts = await cosmicSignature.totalSupply();

  for (let i = Number(oldTotalSupply); i<Number(ts); i++) {
    let ownr = await cosmicSignature.ownerOf(i)
	if (ownr == (await stakingWalletCST.getAddress())) {
		continue;
	}
	let owner_signer = await hre.ethers.getSigner(ownr);
	try {
	   await stakingWalletCST.connect(owner_signer).stake(i);
	} catch(e) {
		console.log("ignoring stake() error for token "+i+", owner "+ownr);
	}
  }

  await cosmicToken.connect(addr1).approve(await cosmicGame.getAddress(),hre.ethers.parseEther('10000000'));
  await cosmicToken.connect(addr2).approve(await cosmicGame.getAddress(),hre.ethers.parseEther('10000000'));
  await cosmicToken.connect(addr3).approve(await cosmicGame.getAddress(),hre.ethers.parseEther('10000000'));
  await cosmicToken.connect(addr4).approve(await cosmicGame.getAddress(),hre.ethers.parseEther('10000000'));
  await cosmicGame.connect(addr1).bidWithCST("CST bid addr1")
  prizeTime = await cosmicGame.timeUntilPrize();
  await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
  await ethers.provider.send("evm_mine");
  tx = await cosmicGame.connect(addr1).claimPrize({gasLimit:3000000});

  await ethers.provider.send("evm_mine");	// mine empty block as spacing
  await ethers.provider.send("evm_mine");	// mine empty block as spacing

  for (let i=0; i<5; i++) {
	await stakingWalletRWalk.connect(addr1).unstake(i);
  }
}

main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error(error);
    process.exit(1);
  });
