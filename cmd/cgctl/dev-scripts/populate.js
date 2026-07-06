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

    // Build address → signer lookup map
    const signerMap = {};
    for (const signer of signers) {
        const address = await signer.getAddress();
        signerMap[address.toLowerCase()] = signer;
    }

    return { signers, signerMap };
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
async function getERC20SampleContracts(deployerSigner, sampContractName = "Samp") {
	// gets address of dummy token contract (for donation testing)
    let sampContract1Addr = process.env.TSAMP1;
    let sampContract2Addr = process.env.TSAMP2;
    const haveEnv = typeof sampContract1Addr === "string" && sampContract1Addr.length === 42 &&
        typeof sampContract2Addr === "string" && sampContract2Addr.length === 42;

    if (!haveEnv) {
        try {
            const signer = deployerSigner || (await hre.ethers.getSigners())[0];
            const SampFactory = await hre.ethers.getContractFactory(sampContractName);
            const samp1 = await SampFactory.connect(signer).deploy("ERC20 Token Sample1", "SAMP1");
            await samp1.waitForDeployment();
            const samp2 = await SampFactory.connect(signer).deploy("ERC20 Token Sample2", "SAMP2");
            await samp2.waitForDeployment();
            sampContract1Addr = await samp1.getAddress();
            sampContract2Addr = await samp2.getAddress();
            console.log("Deployed Samp tokens (TSAMP1/TSAMP2 unset):", sampContract1Addr, sampContract2Addr);
        } catch (e) {
            console.log("TSAMP1 and TSAMP2 must be set to two ERC20 contract addresses (42-char 0x...), or add a Samp contract to your project.");
            console.error(e.message);
            process.exit(1);
        }
    }

    const sampContract1 = await hre.ethers.getContractAt(sampContractName, sampContract1Addr);
    const sampContract2 = await hre.ethers.getContractAt(sampContractName, sampContract2Addr);

    return [sampContract1,sampContract2];
}
async function getRandomWalkNft(game) {

	
}
async function waitUntilPrizeTimeZero(ctrct) {
	let t = Number(await ctrct.getDurationUntilMainPrize());
	if (t <= 0) {
		console.log("time to claim prize = ", t, "(already claimable)");
		return;
	}
	console.log("time to claim prize = ", t);
	// Prefer Hardhat/Anvil time travel to avoid long real-time waits
	try {
		await ethers.provider.send("evm_increaseTime", [t]);
		await ethers.provider.send("evm_mine", []);
		console.log("evm_increaseTime(", t, ") + evm_mine() applied");
		return;
	} catch (e) {
		console.log("evm_increaseTime not available, waiting in real time (poll every 5s):", e.message);
	}
	do {
		t = Number(await ctrct.getDurationUntilMainPrize());
		console.log("time to claim prize = ", t);
		if (t <= 0) break;
		await sleep(5000);
	} while (t > 0);
}

function sleep(ms) {
	return new Promise(resolve => setTimeout(resolve, ms));
}

/** After a claim, the next round has delay before activation. Advance time so next bids succeed. */
async function advanceToNextRoundActive(ctrct) {
	try {
		const delay = Number(await ctrct.delayDurationBeforeRoundActivation());
		if (delay <= 0) return;
		await ethers.provider.send("evm_increaseTime", [delay + 1]);
		await ethers.provider.send("evm_mine", []);
	} catch (e) {
		// Real RPC (geth): wait until round activation time in real time
		const activationTime = Number(await ctrct.roundActivationTime());
		do {
			const block = await ethers.provider.getBlock("latest");
			const now = Number(block.timestamp);
			if (now >= activationTime) return;
			const waitSec = Math.min(5, activationTime - now);
			await sleep(waitSec * 1000);
		} while (true);
	}
}

/** Wait until current round is active (block.timestamp >= roundActivationTime). Use during deploy delay window. */
async function waitUntilRoundActive(ctrct) {
	const activationTime = Number(await ctrct.roundActivationTime());
	let block = await ethers.provider.getBlock("latest");
	let now = Number(block.timestamp);
	if (now >= activationTime) return;
	const needSec = activationTime - now + 1;
	try {
		await ethers.provider.send("evm_increaseTime", [needSec]);
		await ethers.provider.send("evm_mine", []);
		console.log("evm_increaseTime(" + needSec + ") + evm_mine() applied (round active)");
	} catch (e) {
		console.log("Waiting " + needSec + "s for round activation (real time)...");
		await sleep(needSec * 1000);
	}
}

/** Dump SQL INSERT for cg_contracts table (same format as populate-old.js printContractAddresses). */
async function printCgContractsInsert(cosmicGameProxy, cosmicSignatureNft, cosmicToken, charityWallet, prizesWallet, randomWalkNFT, stakingWalletCst, stakingWalletRWalk, marketingWallet) {
	const EIP1967_IMPL_SLOT = "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc";
	let implementationAddr = "0x0000000000000000000000000000000000000000";
	try {
		const implBytes = await ethers.provider.getStorage(await cosmicGameProxy.getAddress(), EIP1967_IMPL_SLOT);
		implementationAddr = (await ethers.AbiCoder.defaultAbiCoder().decode(["address"], implBytes))[0];
	} catch (e) {}
	let cosmicDaoAddr = "0x0000000000000000000000000000000000000000";
	try {
		cosmicDaoAddr = await cosmicGameProxy.cosmicDAO();
	} catch (e) {
		try {
			cosmicDaoAddr = await cosmicGameProxy.dao();
		} catch (e2) {}
	}
	const marketingAddr = marketingWallet ? await marketingWallet.getAddress() : "0x0000000000000000000000000000000000000000";
	const cosmicGameAddr = await cosmicGameProxy.getAddress();
	console.log("");
	console.log("CosmicGame main contract:", cosmicGameAddr);
	console.log("-- SQL INSERT for cg_contracts (cosmic_game_addr, cosmic_signature_addr, cosmic_token_addr, cosmic_dao_addr, charity_wallet_addr, prizes_wallet_addr, random_walk_addr, staking_wallet_cst_addr, staking_wallet_rwalk_addr, marketing_wallet_addr, implementation_addr):");
	console.log(
		"INSERT INTO cg_contracts VALUES(" +
		`'${cosmicGameAddr}',` +
		`'${await cosmicSignatureNft.getAddress()}',` +
		`'${await cosmicToken.getAddress()}',` +
		`'${cosmicDaoAddr}',` +
		`'${await charityWallet.getAddress()}',` +
		`'${await prizesWallet.getAddress()}',` +
		`'${await randomWalkNFT.getAddress()}',` +
		`'${await stakingWalletCst.getAddress()}',` +
		`'${await stakingWalletRWalk.getAddress()}',` +
		`'${marketingAddr}',` +
		`'${implementationAddr}')`
	);
	console.log("");
}

/** Wait until getNextCstBidPrice() <= bidder's CST balance so a CST bid won't revert. */
async function waitUntilCstPriceAffordable(ctrct, cosmicToken, bidderAddress) {
	for (;;) {
		const price = await ctrct.getNextCstBidPrice();
		const balance = await cosmicToken.balanceOf(bidderAddress);
		if (price <= balance) return;
		try {
			await ethers.provider.send("evm_increaseTime", [60]);
			await ethers.provider.send("evm_mine", []);
		} catch (e) {
			await sleep(5000);
		}
	}
}

/** Advance time so the CST Dutch auction price drops (price decreases over time). */
async function advanceTimeForCstDutchAuction(ctrct) {
	const [durationMicro, elapsed] = await ctrct.getCstDutchAuctionDurations();
	const durationSec = Number(durationMicro) / 1e6;
	const elapsedSec = Number(elapsed) > 0 ? Number(elapsed) : 0;
	const remainingSec = Math.max(0, durationSec - elapsedSec);
	if (remainingSec <= 0) return;
	try {
		const advance = Math.ceil(remainingSec) + 2;
		await ethers.provider.send("evm_increaseTime", [advance]);
		await ethers.provider.send("evm_mine", []);
	} catch (e) {
		const waitMs = Math.ceil(remainingSec * 1000);
		let waited = 0;
		while (waited < waitMs) {
			await sleep(1000);
			waited += 1000;
		}
	}
}

async function main() {
    async function mint_rwalk(randomWalkNFT,a) {
            tokenPrice = await randomWalkNFT.getMintPrice();
            let tx = await randomWalkNFT.connect(a).mint({
				value: tokenPrice,gasLimit: 1000000
            });
            let receipt = await tx.wait();
            let topic_sig = randomWalkNFT.interface.getEvent("MintEvent").topicHash;
            let log = receipt.logs.find((x) => x.topics.indexOf(topic_sig) >= 0);
            let parsed_log = randomWalkNFT.interface.parseLog(log);
            let token_id = parsed_log.args[0];
            return token_id;
 	}       
	async function stake_available_nfts(csig,stakingw) {
		let { signerMap } = await customGetSigners()
		let tscst = await csig.totalSupply();
		for (let i = 0; i < tscst; i++) {
			let ownr = await csig.ownerOf(i);
			ownr = ownr.toLowerCase();
			if (ownr == (await stakingw.getAddress())) {
				continue; // already staked
			}
			let owner_signer = signerMap[ownr];
			if (owner_signer === undefined) {
			} else {
			}
			try {
	            let tx = await csig.connect(owner_signer).setApprovalForAll(await stakingw.getAddress(), true,{gasLimit: 1000000});
				await tx.wait()
			} catch (e) {
			}
			try {
				let tx = await stakingw.connect(owner_signer).stake(i,{gasLimit:1000000});
				console.log("tx for staking:",tx)
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
			let tx =await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,{gasLimit:1000000});
			await tx.wait()
			num_unstaked=num_unstaked+1;
		}
	}
	let tx;
	const [owner, addr1, addr2, addr3, addr4, addr5] = (await customGetSigners()).signers;
    const cosmicGameProxy = await getCosmicSignatureGameContract()
	const cosmicGameAddr = await cosmicGameProxy.getAddress()
	console.log("CosmicGame address: "+cosmicGameAddr);
	const [ samp1,samp2 ] = await getERC20SampleContracts(owner);
	const cosmicSignatureNftAddr = await cosmicGameProxy.nft();
	const cosmicSignatureNft = await ethers.getContractAt("CosmicSignatureNft",cosmicSignatureNftAddr)
	const rwalkAddr = await cosmicGameProxy.randomWalkNft();
	const randomWalkNFT = await ethers.getContractAt("RandomWalkNFT",rwalkAddr);
	const stakingWalletCstAddr = await cosmicGameProxy.stakingWalletCosmicSignatureNft();
	const stakingWalletCst = await ethers.getContractAt("StakingWalletCosmicSignatureNft",stakingWalletCstAddr);
	const stakingWalletCosmicSignatureNft = stakingWalletCst;
	const stakingWalletRWalkAddr = await cosmicGameProxy.stakingWalletRandomWalkNft();
	const stakingWalletRWalk= await ethers.getContractAt("StakingWalletRandomWalkNft",stakingWalletRWalkAddr);
	const charityWalletAddr = await cosmicGameProxy.charityAddress();
	const charityWallet = await ethers.getContractAt("CharityWallet",charityWalletAddr);
	const prizesWalletAddr = await cosmicGameProxy.prizesWallet();
	const prizesWallet = await ethers.getContractAt("PrizesWallet",prizesWalletAddr);
	const cosmicTokenAddr = await cosmicGameProxy.token();
	const cosmicToken = await ethers.getContractAt("CosmicSignatureToken", cosmicTokenAddr);
	let marketingWallet;
	try {
		const marketingWalletAddr = await cosmicGameProxy.marketingWallet();
		marketingWallet = await ethers.getContractAt("MarketingWallet", marketingWalletAddr);
	} catch (e) {
		marketingWallet = null;
	}
	const cosmicSignature = cosmicSignatureNft;

	let latestBlock = await ethers.provider.getBlock("latest");
	let rActTime = await cosmicGameProxy.roundActivationTime();
	const roundNotActiveByTime = (Number(latestBlock.timestamp) - Number(rActTime)) < 0;
	// Admin setters (setTimeoutDurationToClaimMainPrize, setCharityAddress, etc.) only succeed when the contract
	// considers the round inactive. The contract can revert with "The current bidding round is already active"
	// even when our (timestamp - roundActivationTime) < 0, so we skip the config block and rely on deploy.
	if (roundNotActiveByTime) {
		console.log("Activation time not met by timestamp; skipping admin config (contract may still treat round as active).");
		try {
			await waitUntilRoundActive(cosmicGameProxy);
			console.log("Round active after wait.");
		} catch (e) {
			console.log("waitUntilRoundActive:", e.message || e);
		}
	}
	console.log("Transacting with CosmicGame contract " + cosmicGameAddr);
	try {
		await cosmicGameProxy.connect(owner).setDelayDurationBeforeRoundActivation(5, { gasLimit: 1000000 });
		console.log("Adjusted activation delay to 5 seconds");
	} catch (err) {
		console.warn("setDelayDurationBeforeRoundActivation(5) skipped (round may be active):", err.message || err);
	}
	let rn = await cosmicGameProxy.roundNum();
	console.log("Round num = "+Number(rn))
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
	latestBlock = await ethers.provider.getBlock("latest");
	console.log("current chain timestamp = "+latestBlock.timestamp);
 	let tdiff = Number(actTime) - Number(latestBlock.timestamp);
	if (tdiff > 0) {
		tx = await cosmicGameProxy.setRoundActivationTime(latestBlock.timestamp);
		await tx.wait();
	}
	console.log("time to activation = "+tdiff);
	console.log("Starting bid transactions for round 0")
    const contractBalance = await ethers.provider.getBalance(
        await cosmicGameProxy.getAddress()
    );
	console.log("Contract balance = ",contractBalance)
    let bidPrice = await cosmicGameProxy.getNextEthBidPrice();
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
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	console.log("Next bidPrice = "+ethers.formatEther(bidPrice))
    tx = await cosmicGameProxy.connect(addr2).bidWithEth(-1,"bid 1", { value: bidPrice });
	console.log("second bid tx hash = "+tx.hash)
	await tx.wait()


//    let nanoSecondsExtra = await cosmicGameProxy.nanoSecondsExtra();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	console.log("Next bidPrice = "+ethers.formatEther(bidPrice))
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 2", { value: bidPrice, gasLimit: 1000000 });
	await tx.wait()
	console.log("Bid tx completed (bidPrice="+bidPrice+")")
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1, "bid 2", { value: bidPrice , gasLimit: 1000000 });
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
    token_id = await mint_rwalk(randomWalkNFT,owner);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(owner).bidWithEth(Number(token_id),"bidWithRWlk", {value: bidPrice , gasLimit: 1000000 });
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice, gasLimit: 1000000 
    });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {
        value: bidPrice, gasLimit: 1000000 
    });
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice , gasLimit: 1000000 });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr4).bidWithEth(-1,"", { value: bidPrice , gasLimit: 1000000 });
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr5).bidWithEth(-1,"", {value: bidPrice , gasLimit: 1000000 });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr5).bidWithEth(-1,"", {value: bidPrice , gasLimit: 1000000 });
	await tx.wait()

	console.log("Finished bid transactions for round 0")
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();

    let prizeAmount = await cosmicGameProxy.getMainEthPrizeAmount();
    let charityAmount = await cosmicGameProxy.getCharityEthDonationAmount();
	let lastBidder = await cosmicGameProxy.lastBidderAddress()
	console.log("lastBidder = "+lastBidder)
	console.log("addr5 address = ",addr5,addr5.address)
	await waitUntilPrizeTimeZero(cosmicGameProxy)
	console.log("Sending claimPrize() tx for round "+Number(rn))
    tx = await cosmicGameProxy.connect(addr5).claimMainPrize({
        gasLimit: 6000000
    });
	await tx.wait()
	await advanceToNextRoundActive(cosmicGameProxy);
	console.log("Claimed prize for "+prizeAmount)
    let prizeAmount2 = await cosmicGameProxy.getMainEthPrizeAmount();
    let expectedprizeAmount = (prizeAmount - charityAmount) / 2n;

	console.log("Making another series of bids\n")
///	await stake_available_nfts(cosmicSignatureNft,stakingWalletCst)
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 4", {
        value: bidPrice, gasLimit: 1000000 
    });
	await tx.wait()
	console.log("Made first bid\n")

    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	console.log("Duration until main prize =",prizeTime)

    prizeAmount = await cosmicGameProxy.getMainEthPrizeAmount();
    charityAmount = await cosmicGameProxy.getCharityEthDonationAmount();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()
	await advanceToNextRoundActive(cosmicGameProxy);
    prizeAmount2 = await cosmicGameProxy.getMainEthPrizeAmount();
	//await stake_available_nfts(cosmicSignatureNft,stakingWalletCst);
    let ts = await cosmicSignatureNft.totalSupply();
    rn = await cosmicGameProxy.roundNum();
    let oldTotalSupply = ts;

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"bid 5", { value: bidPrice , gasLimit: 1000000 });
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 5000000
    });
    receipt = await tx.wait();
	await advanceToNextRoundActive(cosmicGameProxy);
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
	// Round where addr3 wins so they receive CST for later CST bids
	bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"addr3 wins for CST", { value: bidPrice, gasLimit: 1000000 });
	await tx.wait();
	await waitUntilPrizeTimeZero(cosmicGameProxy);
	tx = await cosmicGameProxy.connect(addr3).claimMainPrize({ gasLimit: 3000000 });
	await tx.wait();
	await advanceToNextRoundActive(cosmicGameProxy);
	//await stake_available_nfts(cosmicSignatureNft,stakingWalletCst);
    tx =await charityWallet.connect(addr1).send();
	await tx.wait()

    tx = {
        to: await charityWallet.getAddress(),
        value: hre.ethers.parseEther("4"),
	};
    tx = await addr2.sendTransaction(tx);
	await tx.wait()

    rn = await cosmicGameProxy.roundNum();
	tx = await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("9999999999999999"))
	await tx.wait()
	tx = await samp1.approve(prizesWalletAddr,hre.ethers.parseEther("9999999999999999"),{gasLimit: 1000000});
	await tx.wait()
	tx = await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("999999999999999999999999"))
	await tx.wait()
	tx = await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("999999999999999999999999"))
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	console.log("bid with eth and donate ERC20, tx1")
	tx = await cosmicGameProxy.connect(owner).bidWithEthAndDonateToken(-1,"bid&donateerc20 1",await samp1.getAddress(),10000000000000000000n,{value:bidPrice, gasLimit: 1000000 });
	await tx.wait()
	tx = await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("999999999999999999999999"))
	await tx.wait()
	tx = await samp2.approve(prizesWalletAddr,hre.ethers.parseEther("999999999999999999999999"));
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	console.log("bid with eth and donate ERC20, tx2")
	tx = await cosmicGameProxy.connect(owner).bidWithEthAndDonateToken(-1,"bid&donateerc20 2",await samp2.getAddress(),10000000000000000000n,{value:bidPrice, gasLimit: 1000000 });
	await tx.wait()

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 6", { value: bidPrice , gasLimit: 1000000 });
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

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(randomWalkNFT,addr1);
	tx = await randomWalkNFT.connect(addr1).setApprovalForAll(prizesWalletAddr, true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr1)
        .bidWithEthAndDonateNft(-1,"donated token_id="+token_id,await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice, gasLimit: 1000000 
        });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(randomWalkNFT,addr2);
	tx = await randomWalkNFT.connect(addr2).setApprovalForAll(prizesWalletAddr, true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr2)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice, gasLimit: 1000000 
        });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	tx = await randomWalkNFT.connect(addr3).setApprovalForAll(prizesWalletAddr, true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice, gasLimit: 1000000 
        });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	tx = await randomWalkNFT.connect(addr3).setApprovalForAll(prizesWalletAddr, true);
	await tx.wait()
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithEthAndDonateNft(-1,"me donated token_id="+token_id, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice, gasLimit: 1000000 
        });
	await tx.wait()
	await waitUntilCstPriceAffordable(cosmicGameProxy, cosmicToken, addr3.address);
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	let cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate1", await randomWalkNFT.getAddress(), token_id, {gasLimit: 1000000});
	await tx.wait()
	await waitUntilCstPriceAffordable(cosmicGameProxy, cosmicToken, addr3.address);
    token_id = await mint_rwalk(randomWalkNFT,addr3);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    tx = await cosmicGameProxy
        .connect(addr3)
        .bidWithCstAndDonateNft(cstPrice,"cst bid + donate2", await randomWalkNFT.getAddress(), token_id, {gasLimit: 1000000 });
	await tx.wait()
	await waitUntilCstPriceAffordable(cosmicGameProxy, cosmicToken, addr3.address);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"bid using ERC20 token", {gasLimit: 1000000 });
	await tx.wait()
	await waitUntilCstPriceAffordable(cosmicGameProxy, cosmicToken, addr3.address);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"bid using ERC20 token",{ gasLimit: 1000000 } );
	await tx.wait()


    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 6000000
    });
    receipt = await tx.wait();
	await advanceToNextRoundActive(cosmicGameProxy);
	const roundForDonatedToken = (await cosmicGameProxy.roundNum()) - 1n;
	const donatedAmountRound = 10000000000000000000n; // 10 ether, matches bid&donateerc20 1 & 2
	tx = await prizesWallet.connect(addr3).claimDonatedToken(roundForDonatedToken, await samp1.getAddress(), donatedAmountRound);
	await tx.wait()
	tx = await prizesWallet.connect(addr3).claimDonatedToken(roundForDonatedToken, await samp2.getAddress(), donatedAmountRound);
	await tx.wait()
	//await stake_available_nfts(cosmicSignatureNft,stakingWalletCst,{gasLimit: 1000000});

	let tmp = await prizesWallet.timeoutDurationToWithdrawPrizes();
	tx = await prizesWallet.connect(owner).setTimeoutDurationToWithdrawPrizes(Number(tmp)/2,{gasLimit: 1000000});
	await tx.wait()

    tx = await prizesWallet.connect(addr3).claimDonatedNft(0n,{gasLimit: 1000000});
	await tx.wait()
    tx = await prizesWallet.connect(addr3).claimDonatedNft(1n,{gasLimit: 1000000});
	await tx.wait()
    topic_sig = prizesWallet.interface.getEvent("EthReceived").topicHash;
    deposit_logs = receipt.logs.filter((x) => x.topics.indexOf(topic_sig) >= 0);
    let withdrawal_done = [];
    for (let i = 0; i < deposit_logs.length; i++) {
        let wlog = prizesWallet.interface.parseLog(deposit_logs[i]);
        let winner_signer = await hre.ethers.getSigner(wlog.args.prizeWinnerAddress);
        if (typeof withdrawal_done[wlog.args.winner] === "undefined") {
            tx = await prizesWallet.connect(winner_signer).getFunction("withdrawEth()")({ gasLimit: 1000000 });
			await tx.wait()
            withdrawal_done[wlog.args.winner] = 1;
        } else {
            // skip
        }
    }

    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"", { value: bidPrice , gasLimit: 1000000  });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr1).bidWithEth(-1,"", { value: bidPrice , gasLimit: 1000000  });
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 6000000 
    });
	await tx.wait()
	await advanceToNextRoundActive(cosmicGameProxy);

	//await stake_available_nfts(cosmicSignatureNft,stakingWalletCst)

    donationData =
        '{"version":1,"title":"EF donation","message":"Ethereum Foundation is a non-profit and part of a community of organizations and people working to fund protocol development, grow the ecosystem, and advocate for Ethereum.","url":"http://ethereum.org/en"}';
    tx = await cosmicGameProxy.donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("9"), gasLimit: 1000000 
    });
	await tx.wait()
    tx = await cosmicGameProxy.donateEthWithInfo(donationData, {
        value: hre.ethers.parseEther("8"), gasLimit: 1000000 
    });
	await tx.wait()

    if (marketingWallet) {
		tx = await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("7"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr3.address,hre.ethers.parseEther("7"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("2"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("6"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("5"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr2.address,hre.ethers.parseEther("5"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr3.address,hre.ethers.parseEther("5"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr4.address,hre.ethers.parseEther("1"));
		await tx.wait();
		tx = await marketingWallet.payReward(addr1.address,hre.ethers.parseEther("11"));
		await tx.wait();
	}
//	await unstake_all_nfts()
	tx = await stakingWalletCosmicSignatureNft.tryPerformMaintenance(owner.address, { gasLimit: 1000000 });
	await tx.wait()
//	await stake_available_nfts(cosmicSignatureNft,stakingWalletCst)
	

	tx = await samp1.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("999999999999999999999999"))
	await tx.wait()
	tx = await samp1.approve(prizesWalletAddr,hre.ethers.parseEther("999999999999999999999999"));
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	console.log("bid with eth and donate ERC20, tx3")
	tx = await cosmicGameProxy.connect(owner).bidWithEthAndDonateToken(-1,"bid&donateerc20 3",await samp1.getAddress(),11000000000000000000n,{value:bidPrice, gasLimit: 1000000 });
	await tx.wait()
	tx = await samp2.approve(await cosmicGameProxy.getAddress(),hre.ethers.parseEther("999999999999999999999999"))
	await tx.wait()
	tx = await samp2.approve(prizesWalletAddr,hre.ethers.parseEther("999999999999999999999999"));
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
	console.log("bid with eth and donate ERC20, tx4")
	tx = await cosmicGameProxy.connect(owner).bidWithEthAndDonateToken(-1,"bid&donateerc20 4",await samp2.getAddress(),11000000000000000000n,{value:bidPrice, gasLimit: 1000000 });
	await tx.wait()

    // generate one deposit to charity and not to Staking Wallet
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3",{value: bidPrice, gasLimit: 1000000 });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid 3", {value: bidPrice, gasLimit: 1000000  });
	await tx.wait()
	rn = cosmicGameProxy.roundNum();
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 5000000
    });
	await tx.wait()
	await advanceToNextRoundActive(cosmicGameProxy);

	tx = await prizesWallet.connect(addr3).claimDonatedToken(rn, await samp1.getAddress(), 11000000000000000000n, {gasLimit: 1000000});	// only claim one of the tokens (samp1, not samp2)
	await tx.wait()

    ts = await cosmicSignatureNft.totalSupply();
	// Skip auto-staking newly minted CST NFTs: many revert (e.g. already staked / invalid state),
	// and even staticCall causes Hardhat to log the revert. Re-enable the loop below if you need it.
	/*
	const stakingCstAddr = (await stakingWalletCosmicSignatureNft.getAddress()).toLowerCase();
	const prizesAddr = prizesWalletAddr.toLowerCase();
    for (let i = Number(oldTotalSupply); i < Number(ts); i++) {
        let ownr = (await cosmicSignatureNft.ownerOf(i)).toLowerCase();
        if (ownr === stakingCstAddr) continue;
        if (ownr === prizesAddr) continue;
        try {
            let owner_signer = await hre.ethers.getSigner(ownr);
            const connected = stakingWalletCosmicSignatureNft.connect(owner_signer);
            await connected.stake.staticCall(i, { gasLimit: 1000000 });
            tx = await connected.stake(i, { gasLimit: 1000000 });
            await tx.wait();
        } catch (e) {}
    }
	*/

    tx = await cosmicToken
        .connect(addr1)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000"),
			{gasLimit: 1000000}
        );
	await tx.wait()
    tx = await cosmicToken
        .connect(addr2)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000"),
			{gasLimit: 1000000}
        );
	await tx.wait()
    tx = await cosmicToken
        .connect(addr3)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000"),
			{gasLimit: 1000000}
        );
	await tx.wait()
    tx = await cosmicToken
        .connect(addr4)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000"),
			{gasLimit: 1000000}
        );
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice, gasLimit: 1000000});
	await tx.wait()
	await waitUntilCstPriceAffordable(cosmicGameProxy, cosmicToken, addr1.address);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    tx = await cosmicGameProxy.connect(addr1).bidWithCst(cstPrice,"CST bid addr1",{gasLimit: 1000000});
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr1).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()
	await advanceToNextRoundActive(cosmicGameProxy);
    donationAmount = hre.ethers.parseEther("500");
    tx = await cosmicGameProxy.connect(addr3).donateEth({
        value: donationAmount, gasLimit: 1000000
    });
	await tx.wait()
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice, gasLimit: 1000000});
	await tx.wait()
	await waitUntilCstPriceAffordable(cosmicGameProxy, cosmicToken, addr3.address);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"CST bid addr1",{ gasLimit: 1000000});
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()
	await advanceToNextRoundActive(cosmicGameProxy);
    bidPrice = await cosmicGameProxy.getNextEthBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithEth(-1,"bid eth", {value: bidPrice, gasLimit: 1000000 });
	await tx.wait()
	await waitUntilCstPriceAffordable(cosmicGameProxy, cosmicToken, addr3.address);
	cstPrice = await cosmicGameProxy.getNextCstBidPrice();
    tx = await cosmicGameProxy.connect(addr3).bidWithCst(cstPrice,"CST bid addr1",{ gasLimit: 1000000});
	await tx.wait()
    prizeTime = await cosmicGameProxy.getDurationUntilMainPrize();
	await waitUntilPrizeTimeZero(cosmicGameProxy)
    tx = await cosmicGameProxy.connect(addr3).claimMainPrize({
        gasLimit: 3000000
    });
	await tx.wait()
	await advanceToNextRoundActive(cosmicGameProxy);

	// RWalk unstake: use stake action IDs from 0 to actionCounter-1 (0-based). Skip if no stakes (e.g. RWalk staking block commented out).
	let numRwalkActions = Number(await stakingWalletRWalk.actionCounter());
	for (let i = 0; i < numRwalkActions; i++) {
		try {
			tx = await stakingWalletRWalk.connect(addr1).unstake(i, { gasLimit: 1000000 });
			await tx.wait();
		} catch (e) {
			// unstake(i) may revert if this action is not owned by addr1 or already unstaked
			try {
				tx = await stakingWalletRWalk.connect(addr2).unstake(i, { gasLimit: 1000000 });
				await tx.wait();
			} catch (e2) {
				try {
					tx = await stakingWalletRWalk.connect(addr3).unstake(i, { gasLimit: 1000000 });
					await tx.wait();
				} catch (e3) {
					console.warn("RWalk unstake(" + i + ") skipped:", (e3.message || e3).slice(0, 80));
				}
			}
		}
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
	        tx = await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,{gasLimit: 1000000});
			await tx.wait()
		} catch (e) {
		//	console.log("unstake() error: ",e);
		}
    }

	await printCgContractsInsert(cosmicGameProxy, cosmicSignatureNft, cosmicToken, charityWallet, prizesWallet, randomWalkNFT, stakingWalletCst, stakingWalletRWalk, marketingWallet);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
