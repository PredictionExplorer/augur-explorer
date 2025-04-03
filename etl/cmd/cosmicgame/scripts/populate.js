const { ethers } = require("hardhat");
require("./rpc-helpers.js");
async function customGetSigners() {

	const privateKeys = [
		"0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80",
		"0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d",
		"0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a",
		"0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6",
		"0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a",
		"0x8b3a350cf5c34c9194ca85829a2df0ec3153be0318b5e2d3348e872092edffba"
	];
	const signers = privateKeys.map(key => new ethers.Wallet(key, ethers.provider));
//	[owner, addr1, addr2, addr3, addr4, addr5, ...addrs] = await ethers.getSigners();
	return [signers[0],signers[1],signers[2],signers[3],signers[4],signers[5]];

}
async function getCosmicSignatureGameContract(cosmicSignatureGameContractName = "CosmicSignatureGame") {
    const cosmicSignatureGameAddr = process.env.CADDR;
    if (typeof cosmicSignatureGameAddr === "undefined" || cosmicSignatureGameAddr.length !== 42) {
        console.log("CADDR environment variable does not contain contract address of CosmicGame contract.");
        process.exit(1);
    }

    const cosmicSignatureGame = await hre.ethers.getContractAt(cosmicSignatureGameContractName, cosmicSignatureGameAddr);

    return cosmicSignatureGame;
}
async function getERC20SampleContracts(sampContractName = "Samp") {
	// gets address of dummy token contract (for donation testing)
    const sampContract1Addr = process.env.TSAMP1;
    const sampContract2Addr = process.env.TSAMP2;
    if (typeof sampContract1Addr === "undefined" || sampContract1Addr.length !== 42) {
        console.log("TSAMP1 environment variable does not contain contract address of token sample contract.");
        process.exit(1);
    }
    if (typeof sampContract2Addr === "undefined" || sampContract2Addr.length !== 42) {
        console.log("TSAMP2 environment variable does not contain contract address of token sample contract.");
        process.exit(1);
    }

    const sampContract1 = await hre.ethers.getContractAt(sampContractName, sampContract1Addr);
    const sampContract2 = await hre.ethers.getContractAt(sampContractName, sampContract2Addr);

    return [sampContract1,sampContract2];
}
async function getRandomWalkNft(game) {

	
}

async function main() {
    async function mint_rwalk(randomWalkNFT,a) {
            tokenPrice = await randomWalkNFT.getMintPrice();
            let tx = await randomWalkNFT.connect(a).mint({
                value: tokenPrice
            });
            let receipt = await tx.wait();
            let topic_sig = randomWalkNFT.interface.getEvent("MintEvent").topicHash;
            let log = receipt.logs.find((x) => x.topics.indexOf(topic_sig) >= 0);
            let parsed_log = randomWalkNFT.interface.parseLog(log);
            let token_id = parsed_log.args[0];
            return token_id;
 	}       
	async function stake_available_nfts(csig,stakingw) {
		let tscst = await csig.totalSupply();
		for (let i = 0; i < tscst; i++) {
			let ownr = await csig.ownerOf(i);
			if (ownr == (await stakingw.getAddress())) {
				continue; // already staked
			}
			let owner_signer = await hre.ethers.getSigner(ownr);
			if (owner_signer === undefined) {
			} else {
			}
			try {
	            csig.connect(owner_signer).setApprovalForAll(await stakingw.getAddress(), true);
			} catch (e) {
			}
			try {
				let tx = await stakingw.connect(owner_signer).stake(i);
				await tx.wait()
			} catch (e) {
			}
		}
	}
	async function unstake_all_nfts() {
		let num_a = await stakingWalletCosmicSignatureNft.actionCounter();
		let num_unstaked = 0;
		for (let i = 1; i <= num_a; i++) {
			let action_rec = (await stakingWalletCosmicSignatureNft.stakeActions(i)).toObject();
			let ownr = action_rec.nftOwnerAddress;
			let num_s = await stakingWalletCosmicSignatureNft.numStakedNfts();
			if (ownr == "0x0000000000000000000000000000000000000000") {
				continue; 
			}
			let owner_signer = await hre.ethers.getSigner(ownr);
			let tx =await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,10);
			await tx.wait()
			num_unstaked=num_unstaked+1;
		}
	}
	[owner, addr1, addr2, addr3, addr4, addr5] = await customGetSigners();
    const cosmicGameProxy = await getCosmicSignatureGameContract()
	const cosmicGameAddr = await cosmicGameProxy.getAddress()
	const [ samp1,samp2 ] = await getERC20SampleContracts();
	const cosmicSignatureNftAddr = await cosmicGameProxy.nft();
	const cosmicSignatureNft = await ethers.getContractAt("CosmicSignatureNft",cosmicSignatureNftAddr)
	const rwalkAddr = await cosmicGameProxy.randomWalkNft();
	const randomWalkNFT = await ethers.getContractAt("RandomWalkNFT",rwalkAddr);
	const stakingWalletCstAddr = await cosmicGameProxy.stakingWalletCosmicSignatureNft();
	const stakingWalletCst = await ethers.getContractAt("StakingWalletCosmicSignatureNft",stakingWalletCstAddr);
	const stakingWalletRWalkAddr = await cosmicGameProxy.stakingWalletRandomWalkNft();
	const stakingWalletRWalk= await ethers.getContractAt("StakingWalletRandomWalkNft",stakingWalletRWalkAddr);
	console.log("Transacting with CosmicGame contract "+cosmicGameAddr)
	await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(5);
	console.log("Adjusted activation delay to 5 seconds")
    let rn = await cosmicGameProxy.roundNum();
	console.log("Round num = "+Number(rn))
	let tx;
	let token_id;
    let donationAmount = hre.ethers.parseEther("100");
    await cosmicGameProxy.connect(addr5).donateEth({
        value: donationAmount
    });
	console.log("Made donate() transaction")
    let donationData =
        '{"version":1,"title":"Hardhat donation","message":"Donation from HardHat","url":"http://hardhat.org"}';
    await cosmicGameProxy.connect(addr4).donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("60"),
    });
    donationData =
        '{"version":1,"title":"ArtBlocks donation","message":"ArtBlocks offers a platform for creators, buyers and sellers of digital assets and any non-digital products, services and/or benefits to be furnished by or on behalf of sellers in connection with such sales","url":"https://www.artblocks.io"}';
	console.log("Made donate-with-info transaction (ArtBlocks)")
    await cosmicGameProxy
        .connect(addr2)
        .donateEthWithInfo(donationData, {
            value: hre.ethers.parseEther("90")
        });
    console.log("Donation complete");

/*
	console.log("Starting RWalk staking transactions")
	console.log("Starting RWalk transactions for addr "+addr1.address)
	let numStakeActions = 5
    for (let i = 0; i < numStakeActions; i++) {
        let token_id = await mint_rwalk(randomWalkNFT,addr1);
        await randomWalkNFT
            .connect(addr1)
            .setApprovalForAll(await stakingWalletRWalk.getAddress(), true);
		let txdata = stakingWalletRWalk.interface.encodeFunctionData("stake",[token_id]);
		const tx = await addr1.sendTransaction({
			to: await stakingWalletRWalk.getAddress(),
			data: txdata,
			gasLimit: 800000
		});
		await tx.wait();
        //await stakingWalletRandomWalkNft.connect(addr1).stake(token_id);
    }
	console.log("Finished RWalk staking transactions for addr "+addr1.address)
	console.log("Starting RWalk transactions for addr "+addr2.address)
    for (let i = 0; i < 5; i++) {
        let token_id = await mint_rwalk(randomWalkNFT,addr2);
        await randomWalkNFT
            .connect(addr2)
            .setApprovalForAll(await stakingWalletRWalk.getAddress(), true);
		let txdata = stakingWalletRWalk.interface.encodeFunctionData("stake",[token_id]);
		const tx = await addr2.sendTransaction({
			to: await stakingWalletRWalk.getAddress(),
			data: txdata,
			gasLimit: 800000
		});
		await tx.wait();
        //await stakingWalletRandomWalkNft.connect(addr2).stake(token_id);
    }
	console.log("Finished RWalk staking transactions for addr "+addr2.address)
	console.log("Starting RWalk transactions for addr "+addr3.address)
    for (let i = 0; i < 15; i++) {
        let token_id = await mint_rwalk(randomWalkNFT,addr3);
        await randomWalkNFT
            .connect(addr3)
            .setApprovalForAll(await stakingWalletRWalk.getAddress(), true);
		let txdata = stakingWalletRWalk.interface.encodeFunctionData("stake",[token_id]);
		const tx = await addr3.sendTransaction({
			to: await stakingWalletRWalk.getAddress(),
			data: txdata,
			gasLimit: 800000
		});
		await tx.wait();
        //await stakingWalletRandomWalkNft.connect(addr3).stake(token_id);
    }
	console.log("Finished RWalk staking transactions for addr "+addr3.address)
	*/
    let prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

	let delayDur = await cosmicGameProxy.delayDurationBeforeRoundActivation()
	console.log("Delay duration before activatiopn = "+delayDur)
	let actTime = await cosmicGameProxy.roundActivationTime()
	console.log("activation time = "+actTime);
	const latestBlock = await ethers.provider.getBlock("latest");
	console.log("current chain timestamp = "+latestBlock.timestamp);
	console.log("Starting bid transactions for round 0")
    const contractBalance = await ethers.provider.getBalance(
        await cosmicGameProxy.getAddress()
    );
	console.log("Contract balance = ",contractBalance)
    let bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	let txdata = cosmicGameProxy.interface.encodeFunctionData("bidWithEth",[-1,"bid 1"]);
	tx = await addr1.sendTransaction({
		to: await cosmicGameProxy.getAddress(),
		data: txdata,
		gasLimit: 500000,
		value: bidPrice,
	});
	console.log("first bid tx hash = "+tx.hash)
	await tx.wait();
	console.log("Bid tx completed (bidPrice="+ethers.formatEther(bidPrice)+")")
    //await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 1", { value: bidPrice + 1000n }); // this works
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	console.log("Next bidPrice = "+ethers.formatEther(bidPrice))
    tx = await cosmicGameProxy.connect(addr2).bidWithEth(-1,"bid 1", { value: bidPrice });
	console.log("second bid tx hash = "+tx.hash)
	await tx.wait()


//    let nanoSecondsExtra = await cosmicGameProxy.nanoSecondsExtra();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	console.log("Next bidPrice = "+ethers.formatEther(bidPrice))
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 2", { value: bidPrice });
	await tx.wait()
	console.log("Bid tx completed (bidPrice="+bidPrice+")")
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1, "bid 2", { value: bidPrice });
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    token_id = await mint_rwalk(randomWalkNFT,owner);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(owner).bidWithEth(Number(token_id),"bidWithRWlk", {value: bidPrice });
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice
    });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice
    });
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice });
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr5).bidWithEth(-1,"", {value: bidPrice });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr5).bidWithEth(-1,"", {value: bidPrice });
	await tx.wait()

	console.log("Finished bid transactions for round 0")
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    let prizeAmount = await cosmicGameProxy.getMainEthPrizeAmount();
    let charityAmount = await cosmicGameProxy.getCharityEthDonationAmount();
	let lastBidder = await cosmicGameProxy.lastBidderAddress()
	console.log("lastBidder = "+lastBidder)
	console.log("addr5 address = ",addr5,addr5.address)
	console.log("Sending claimPrize() tx for round 0")
    tx = await cosmicGameProxy.connect(addr5).claimMainPrize({
        gasLimit: 5000000
    });
	await tx.wait()
	console.log("Claimed prize for "+prizeAmount)
    let prizeAmount2 = await cosmicGameProxy.getMainEthPrizeAmount();
    let expectedprizeAmount = (prizeAmount - charityAmount) / 2n;

	console.log("Making another series of bids\n")
	await stake_available_nfts(cosmicSignatureNft,stakingWalletCst)
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 4", {
        value: bidPrice
    });
	await tx.wait()
	console.log("Made first bid\n")

    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	console.log("Duration until main prize =",prizeTime)

    prizeAmount = await cosmicGameProxy.getMainEthPrizeAmount();
    charityAmount = await cosmicGameProxy.getCharityEthDonationAmount();
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()
    prizeAmount2 = await cosmicGameProxy.getMainEthPrizeAmount();
	await stake_available_nfts(cosmicSignatureNft,stakingWalletCst);
    let ts = await cosmicSignatureNft.totalSupply();
    rn = await cosmicGameProxy.roundNum();
    let oldTotalSupply = ts;

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 5", { value: bidPrice });
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 5000000
    });
    receipt = await tx.wait();
    topic_sig = cosmicSignatureNft.interface.getEvent("NftMinted").topicHash;
	let event_logs = receipt.logs.filter(log => log.topics[0] === topic_sig);
	let mint_found = false;
	for (let i=0; i< event_logs.length; i++) {
		let parsed_log = cosmicSignatureNft.interface.parseLog(event_logs[i]);
	    let args = parsed_log.args.toObject();
		if (args.nftOwnerAddress === addr1.address) {
		  	token_id = args.nftId;
	    	tx = await cosmicSignatureNft.connect(addr1).setNftName(token_id, "name 0");
			await tx.wait()
	    	tx = await cosmicSignatureNft.connect(addr1).setNftName(token_id, "name after 0");
			await tx.wait()
			mint_found = true;
		}
		if (mint_found) {
			break;
		}
	}
	if (!mint_found) {
		console.log("No NFT was minted for the winner. Bug!")
		return;
	}
	await stake_available_nfts(cosmicSignatureNft,stakingWalletCst);
    tx =await charityWallet.connect(addr1).send();
	await tx.wait()

    tx = {
        to: await charityWallet.getAddress(),
        value: hre.ethers.parseEther("4"),
    };
	await tx.wait()
    tx = await addr2.sendTransaction(tx);
	await tx.wait()
    tx = await addr2.sendTransaction(tx);
	await tx.wait()

    rn = await cosmicGameProxy.roundNum();
	tx = await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await tx.wait()
	tx = await samp1.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	tx = await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp1.getAddress(),10000000000000000000n,{value:bidPrice});
	await tx.wait()
	tx = await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await tx.wait()
	tx = await samp2.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	tx = await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp2.getAddress(),10000000000000000000n,{value:bidPrice});
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 6", { value: bidPrice });
	await tx.wait()

    tx = await randomWalkNFT
        .connect(addr1)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);
	await tx.wait()
    tx = await randomWalkNFT
        .connect(addr2)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);
	await tx.wait()
    tx = await randomWalkNFT
        .connect(addr3)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr1);
	tx = await randomWalkNFT.connect(addr1).setApprovalForAll(await prizesWallet.getAddress(), true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr1)
        .bidWithEthAndDonateNft(-1,"donated token_id="+token_id,await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr2);
	tx = await randomWalkNFT.connect(addr2).setApprovalForAll(await prizesWallet.getAddress(), true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr2)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	tx = await randomWalkNFT.connect(addr3).setApprovalForAll(await prizesWallet.getAddress(), true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	tx = await randomWalkNFT.connect(addr3).setApprovalForAll(await prizesWallet.getAddress(), true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });
	await tx.wait()
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	let cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate1", await randomWalkNFT.getAddress(), token_id);
	await tx.wait()
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate2", await randomWalkNFT.getAddress(), token_id);
	await tx.wait()
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"bid using ERC20 token");
	await tx.wait()
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"bid using ERC20 token");
	await tx.wait()


    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
    receipt = await tx.wait();
	tx = prizesWallet.connect(addr3).claimDonatedToken(rn,await samp1.getAddress());
	await tx.wait()
	tx = prizesWallet.connect(addr3).claimDonatedToken(rn,await samp2.getAddress());
	await tx.wait()
	await stake_available_nfts(cosmicSignatureNft,stakingWalletCst);

    tx= await cosmicGameProxy
        .connect(owner)
        .setCharityAddress(await charityWallet.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setRandomWalkNft(await randomWalkNFT.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setPrizesWallet(await prizesWallet.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setStakingWalletCst(await stakingWalletCst.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setStakingWalletRWalk(await stakingWalletRWalk.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setMarketingWallet(await marketingWallet.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setNumRaffleEthPrizesForBidders(4);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setNumRaffleCosmicSignatureNftsForBidders(6);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setNumRaffleCosmicSignatureNftsForRandomWalkNftStakers(3);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setMainEthPrizeAmountPercentage(30);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setCharityEthDonationAmountPercentage(5);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setRaffleTotalEthPrizeAmountForBiddersPercentage(6);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setCharityAddress(addr3.address);
	await tx.wait()
	tx = await cosmicGameProxy.connect(owner).setChronoWarriorEthPrizeAmountPercentage(8);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setCharityAddress(await charityWallet.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setCosmicSignatureNftStakingTotalEthRewardAmountPercentage(19);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setCosmicSignatureToken(await cosmicToken.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setCosmicSignatureNft(await cosmicSignature.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setCstDutchAuctionDurationDivisor(13 * 3600);
	await tx.wait()
	let tmpval = await cosmicGameProxy.ethDutchAuctionDurationDivisor();
	tmpval = tmpval + 11n;
	tx = await cosmicGameProxy.connect(owner).setEthDutchAuctionDurationDivisor(tmpval);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setCstRewardAmountMultiplier(5);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setCstDutchAuctionBeginningBidPriceMinLimit(150000000000000000000n);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setMarketingWalletCstContributionAmount(120000000000000000000n);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setCstRewardAmountForBidding(130000000000000000000n);
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setBidMessageLengthMaxLimit(199);
	await tx.wait()
    tx = await cosmicSignatureNft.connect(owner).setNftGenerationScriptUri("ipfs://");
	await tx.wait()
    await cosmicSignatureNft.connect(owner).setNftBaseUri("nttp://");
    let iAddrBytes = await ethers.provider.getStorage(
        await cosmicGameProxy.getAddress(),
        "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
    );
    let iAddr = await ethers.AbiCoder.defaultAbiCoder()
        .decode(["address"], iAddrBytes)
        .toString();
    tx = await cosmicGameProxy.connect(owner).upgradeTo(iAddr);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setStakingWalletCosmicSignatureNft(await stakingWalletCst.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(owner)
        .setStakingWalletRandomWalkNft(await stakingWalletRWalk.getAddress());
	await tx.wait()
    tx = await cosmicGameProxy.connect(owner).setCstRewardAmountMultiplier(999);
	await tx.wait()
    let tmp = await cosmicGameProxy.mainPrizeTimeIncrementIncreaseDivisor();
    tx = await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementIncreaseDivisor(tmp);
	await tx.wait()
    tmp = await cosmicGameProxy.connect(owner).timeoutDurationToClaimMainPrize();
    tx = await cosmicGameProxy.connect(owner).setTimeoutDurationToClaimMainPrize(tmp);
	await tx.wait()
    tmp = await cosmicGameProxy.nextEthBidPriceIncreaseDivisor();
    tx = await cosmicGameProxy.connect(owner).setNextEthBidPriceIncreaseDivisor(tmp);
	await tx.wait()
    tmp = await cosmicGameProxy.mainPrizeTimeIncrementInMicroSeconds();
    tx = await cosmicGameProxy.connect(owner).setMainPrizeTimeIncrementInMicroSeconds(tmp);
	await tx.wait()
    tmp = await cosmicGameProxy.initialDurationUntilMainPrizeDivisor();
    tx = await cosmicGameProxy.connect(owner).setInitialDurationUntilMainPrizeDivisor(tmp);
	await tx.wait()
    tmp = await cosmicGameProxy.delayDurationBeforeRoundActivation();
    tx = await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(tmp);
	await tx.wait()
	tmp = await prizesWallet.timeoutDurationToWithdrawPrizes();
	tx = await prizesWallet.connect(owner).setTimeoutDurationToWithdrawPrizes(Number(tmp)/2);
	await tx.wait()

    tx = await prizesWallet.connect(addr3).claimDonatedNft(0n);
	await tx.wait()
    tx = await prizesWallet.connect(addr3).claimDonatedNft(1n);
	await tx.wait()
    topic_sig = prizesWallet.interface.getEvent("EthReceived").topicHash;
    deposit_logs = receipt.logs.filter((x) => x.topics.indexOf(topic_sig) >= 0);
    let withdrawal_done = [];
    for (let i = 0; i < deposit_logs.length; i++) {
        let wlog = prizesWallet.interface.parseLog(deposit_logs[i]);
        let winner_signer = await hre.ethers.getSigner(wlog.args.prizeWinnerAddress);
        if (typeof withdrawal_done[wlog.args.winner] === "undefined") {
            tx = await prizesWallet.connect(winner_signer).withdrawEth();
			await tx.wait()
            withdrawal_done[wlog.args.winner] = 1;
        } else {
            // skip
        }
    }

    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"", { value: bidPrice  });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"", { value: bidPrice  });
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 30000000
    });
	await tx.wait()

	await stake_available_nfts(cosmicSignatureNft,stakingWalletCst)

    donationData =
        '{"version":1,"title":"EF donation","message":"Ethereum Foundation is a non-profit and part of a community of organizations and people working to fund protocol development, grow the ecosystem, and advocate for Ethereum.","url":"http://ethereum.org/en"}';
    tx = await cosmicGameProxy.donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("9"),
    });
	await tx.wait()
    tx = await cosmicGameProxy.donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("8"),
    });
	await tx.wait()

    tx = await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("7"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr3.address,hre.ethers.parseEther("7"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("2"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("6"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("5"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("5"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr3.address,hre.ethers.parseEther("5"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr4.address,hre.ethers.parseEther("1"));
	await tx.wait()
    tx = await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("11"));
	await tx.wait()
	await unstake_all_nfts()
	tx = await stakingWalletCosmicSignatureNft.tryPerformMaintenance(true,owner.address)
	await tx.wait()
	tx = await stake_available_nfts(cosmicSignatureNft,stakingWalletCst)
	await tx.wait()

	tx = await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await tx.wait()
	tx = await samp1.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	tx = await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp1.getAddress(),11000000000000000000n,{value:bidPrice});
	await tx.wait()
	tx = await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await tx.wait()
	tx = await samp2.approve(await prizesWallet.getAddress(),hre.ethers.parseEther("9999999999999999"));
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
	tx = await cosmicGameProxy.bidWithEthAndDonateToken(-1,"bid&donateerc20",await samp2.getAddress(),11000000000000000000n,{value:bidPrice});
	await tx.wait()

    // generate one deposit to charity and not to Staking Wallet
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3",{value: bidPrice});
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {value: bidPrice });
	await tx.wait()
	rn = cosmicGameProxy.roundNum();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()

	tx = prizesWallet.connect(addr3).claimDonatedToken(rn,await samp1.getAddress());	// only claim one of the tokens (samp1, not samp2)
	await tx.wait()

    ts = await cosmicSignatureNft.totalSupply();

    for (let i = Number(oldTotalSupply); i < Number(ts); i++) {
        let ownr = await cosmicSignatureNft.ownerOf(i);
        if (ownr == (await stakingWalletCosmicSignatureNft.getAddress())) {
            continue;
        }
        let owner_signer = await hre.ethers.getSigner(ownr);
        try {
            tx  = await stakingWalletCosmicSignatureNft.connect(owner_signer).stake(i);
			await tx.wait()
        } catch (e) {
            //console.log("ignoring stake() error for token " + i + ", owner " + ownr);
        }
    }

    tx = await cosmicToken
        .connect(addr1)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
	await tx.wait()
    tx = await cosmicToken
        .connect(addr2)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
	await tx.wait()
    tx = await cosmicToken
        .connect(addr3)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
	await tx.wait()
    tx = await cosmicToken
        .connect(addr4)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	await tx.wait()
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    tx = await cosmicGameProxy.connect(addr1).bidWithCst(cstPrice,"CST bid addr1");
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()
    donationAmount = hre.ethers.parseEther("500");
    tx = await cosmicGameProxy.connect(addr3).donateEth({
        value: donationAmount
    });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	await tx.wait()
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"CST bid addr1");
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice });
	await tx.wait()
	cstPrice = await cosmicGameProxy.getNextCstBidPrice(0);
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"CST bid addr1");
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()

    for (let i = 1; i <= 5; i++) {
        tx = await stakingWalletRWalk.connect(addr1).unstake(i);
		await tx.wait()
    }
	numStakeActions = await stakingWalletCosmicSignatureNft.actionCounter();
    for (let i = 1; i <= numStakeActions; i++) {
        let action_rec = (await stakingWalletCosmicSignatureNft.stakeActions(i)).toObject();
        let ownr = action_rec.nftOwnerAddress;
		if (ownr == "0x0000000000000000000000000000000000000000") { continue; }
		if (ownr == addr1.address) {
			// make some address keeping its tokens
			continue;
		} else {
		}
        let owner_signer = await hre.ethers.getSigner(ownr);
		try {
	        tx = await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,1);
			await tx.wait()
	        tx = await stakingWalletCosmicSignatureNft.connect(owner_signer).payReward(i,100);
			await tx.wait()
		} catch (e) {
		//	console.log("unstake() error: ",e);
		}
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
