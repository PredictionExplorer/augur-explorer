//const hre = require("hardhat");
const {
    basicDeployment
} = require("./Deploy.js");
const bidParamsEncoding = {
    type: "tuple(string,int256)",
    name: "bidparams",
    components: [{
            name: "msg",
            type: "string"
        },
        {
            name: "rwalk",
            type: "int256"
        },
    ],
};
async function main() {
    async function mint_rwalk(a) {
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
	async function stake_available_nfts() {
		let tscst = await cosmicSignature.totalSupply();
		for (let i = 0; i < tscst; i++) {
			let ownr = await cosmicSignature.ownerOf(i);
			if (ownr == (await stakingWalletCosmicSignatureNft.getAddress())) {
				continue; // already staked
			}
			let owner_signer = await hre.ethers.getSigner(ownr);
			try {
	            cosmicSignature.connect(owner_signer).setApprovalForAll(await stakingWalletCosmicSignatureNft.getAddress(), true);
			} catch (e) {
				//console.log("stake_available_nfts(): err at approval:",e);
			}
			try {
				await stakingWalletCosmicSignatureNft.connect(owner_signer).stake(i);
			} catch (e) {
			//	console.log("stake_available_nfts() stake() error: ",e);
			}
		}

	}
        [owner, addr1, addr2, addr3, addr4, addr5, ...addrs] =
        await ethers.getSigners();
    const {
        cosmicGameProxy,
        cosmicToken,
        cosmicSignature,
        charityWallet,
        cosmicDAO,
        raffleWallet,
        randomWalkNFT,
        stakingWalletCosmicSignatureNft,
        stakingWalletRandomWalkNft,
        marketingWallet,
        implementationAddr,
    } = await basicDeployment(owner, "", 0, "", false, true);
    console.log("Addresses set");
    console.log(
        "INSERT INTO cg_contracts VALUES('" +
        (await cosmicGameProxy.getAddress()) +
        "','" +
        (await cosmicSignature.getAddress()) +
        "','" +
        (await cosmicToken.getAddress()) +
        "','" +
        (await cosmicDAO.getAddress()) +
        "','" +
        (await charityWallet.getAddress()) +
        "','" +
        (await raffleWallet.getAddress()) +
        "','" +
        (await randomWalkNFT.getAddress()) +
        "','" +
        (await stakingWalletCosmicSignatureNft.getAddress()) +
        "','" +
        (await stakingWalletRandomWalkNft.getAddress()) +
        "','" +
        (await marketingWallet.getAddress()) +
        "','" +
        implementationAddr +
        "')"
    );
    let donationAmount = hre.ethers.parseEther("10");
    await cosmicGameProxy.donate({
        value: donationAmount
    });
    let donationData =
        '{"version":1,"title":"Hardhat donation","message":"Donation from HardHat","url":"http://hardhat.org"}';
    await cosmicGameProxy.donateWithInfo(donationData, {
        value: hre.ethers.parseEther("6"),
    });
    donationData =
        '{"version":1,"title":"ArtBlocks donation","message":"ArtBlocks offers a platform for creators, buyers and sellers of digital assets and any non-digital products, services and/or benefits to be furnished by or on behalf of sellers in connection with such sales","url":"https://www.artblocks.io"}';

    await cosmicGameProxy
        .connect(addr2)
        .donateWithInfo(donationData, {
            value: hre.ethers.parseEther("6")
        });

	let numStakeActions = 5
    for (let i = 0; i < numStakeActions; i++) {
        let token_id = await mint_rwalk(addr1);
        await randomWalkNFT
            .connect(addr1)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
        await stakingWalletRandomWalkNft.connect(addr1).stake(token_id);
    }
    for (let i = 0; i < 5; i++) {
        let token_id = await mint_rwalk(addr2);
        await randomWalkNFT
            .connect(addr2)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
        await stakingWalletRandomWalkNft.connect(addr2).stake(token_id);
    }
    for (let i = 0; i < 50; i++) {
        let token_id = await mint_rwalk(addr3);
        await randomWalkNFT
            .connect(addr3)
            .setApprovalForAll(await stakingWalletRandomWalkNft.getAddress(), true);
        await stakingWalletRandomWalkNft.connect(addr3).stake(token_id);
    }

    let prizeTime = await cosmicGameProxy.timeUntilPrize();
    console.log("Donation complete");

    var bidParams = {
        msg: "bid 1",
        rwalk: -1
    };
    let params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    const contractBalance = await ethers.provider.getBalance(
        await cosmicGameProxy.getAddress()
    );
    let bidPrice = await cosmicGameProxy.getBidPrice();
    await cosmicGameProxy.connect(addr1).bid(params, {
        value: bidPrice + 1000n
    }); // this works
    bidPrice = await cosmicGameProxy.getBidPrice();
    bdParams = {
        msg: "bid 1",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr2).bid(params, {
        value: bidPrice + 1000n
    }); // this works

    let nanoSecondsExtra = await cosmicGameProxy.nanoSecondsExtra();
    prizeTime = await cosmicGameProxy.timeUntilPrize();

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 2",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr1).bid(params, {
        value: bidPrice
    });
    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 2",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr1).bid(params, {
        value: bidPrice
    });
    prizeTime = await cosmicGameProxy.timeUntilPrize();
    let token_id = await mint_rwalk(owner);
    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bidWithRWLK",
        rwalk: Number(token_id)
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(owner).bid(params, {
        value: bidPrice
    });

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 3",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr3).bid(params, {
        value: bidPrice
    });
    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 3",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr3).bid(params, {
        value: bidPrice
    });

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr4).bid(params, {
        value: bidPrice
    });
    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr4).bid(params, {
        value: bidPrice
    });

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr5).bid(params, {
        value: bidPrice
    });
    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr5).bid(params, {
        value: bidPrice
    });

    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime) - 100]);
    await ethers.provider.send("evm_mine");

    await ethers.provider.send("evm_increaseTime", [100]);
    await ethers.provider.send("evm_mine");

    let prizeAmount = await cosmicGameProxy.prizeAmount();
    let charityAmount = await cosmicGameProxy.charityAmount();
    await cosmicGameProxy.connect(addr5).claimPrize({
        gasLimit: 30000000
    });
    let prizeAmount2 = await cosmicGameProxy.prizeAmount();
    let expectedprizeAmount = (prizeAmount - charityAmount) / 2n;

	stake_available_nfts()

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 4",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr1).bid(params, {
        value: bidPrice
    });

    prizeTime = await cosmicGameProxy.timeUntilPrize();

    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");

    prizeAmount = await cosmicGameProxy.prizeAmount();
    charityAmount = await cosmicGameProxy.charityAmount();
    await cosmicGameProxy.connect(addr1).claimPrize({
        gasLimit: 3000000
    });
    prizeAmount2 = await cosmicGameProxy.prizeAmount();
	stake_available_nfts();
    let ts = await cosmicSignature.totalSupply();
    let rn = await cosmicGameProxy.roundNum();
	/* discontinued, removal pending
    let topic_sig =
        stakingWalletCosmicSignatureNft.interface.getEvent("NftStaked").topicHash;
    for (let i = 0; i < Number(ts); i++) {
        let tx;
        let ownr = await cosmicSignature.ownerOf(i);
        let owner_signer = await hre.ethers.getSigner(ownr);
        await cosmicSignature
            .connect(owner_signer)
            .setApprovalForAll(await stakingWalletCosmicSignatureNft.getAddress(), true);
        tx = await stakingWalletCosmicSignatureNft.connect(owner_signer).stake(i);
    }
	*/
    let oldTotalSupply = ts;

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 5",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr1).bid(params, {
        value: bidPrice
    });
    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    tx = await cosmicGameProxy.connect(addr1).claimPrize({
        gasLimit: 5000000
    });
    receipt = await tx.wait();
    topic_sig = cosmicSignature.interface.getEvent("MintEvent").topicHash;
    log = receipt.logs.find((x) => x.topics.indexOf(topic_sig) >= 0);
    parsed_log = cosmicSignature.interface.parseLog(log);
    let args = parsed_log.args.toObject();
    token_id = args.nftId;
    await cosmicSignature.connect(addr1).setTokenName(token_id, "name 0");
    await cosmicSignature.connect(addr1).setTokenName(token_id, "name after 0");

	stake_available_nfts();
    await charityWallet.connect(addr1).send();

    tx = {
        to: await charityWallet.getAddress(),
        value: hre.ethers.parseEther("4"),
    };
    await addr2.sendTransaction(tx);
    await addr2.sendTransaction(tx);

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 6",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr3).bid(params, {
        value: bidPrice
    });
    await ethers.provider.send("evm_mine"); // mine empty block as spacing

    await randomWalkNFT
        .connect(addr1)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);
    await randomWalkNFT
        .connect(addr2)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);
    await randomWalkNFT
        .connect(addr3)
        .setApprovalForAll(await cosmicGameProxy.getAddress(), true);

    bidPrice = await cosmicGameProxy.getBidPrice();
    token_id = await mint_rwalk(addr1);
    bidParams = {
        msg: "donated token_id=" + token_id,
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy
        .connect(addr1)
        .bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });

    bidPrice = await cosmicGameProxy.getBidPrice();
    token_id = await mint_rwalk(addr2);
    bidParams = {
        msg: "me donated token_id=" + token_id,
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy
        .connect(addr2)
        .bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });

    bidPrice = await cosmicGameProxy.getBidPrice();
    token_id = await mint_rwalk(addr3);
    bidParams = {
        msg: "me donated token_id=" + token_id,
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy
        .connect(addr3)
        .bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });

    bidPrice = await cosmicGameProxy.getBidPrice();
    token_id = await mint_rwalk(addr3);
    bidParams = {
        msg: "me donated token_id=" + token_id,
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy
        .connect(addr3)
        .bidAndDonateNFT(params, await randomWalkNFT.getAddress(), token_id, {
            value: bidPrice,
        });

    await cosmicGameProxy.connect(addr3).bidWithCST("bid using ERC20 token");
    await ethers.provider.send("evm_mine");
    await cosmicGameProxy.connect(addr3).bidWithCST("bid using ERC20 token");
    await ethers.provider.send("evm_mine");

    await cosmicGameProxy.connect(owner).prepareMaintenance();

    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimPrize({
        gasLimit: 3000000
    });
    receipt = await tx.wait();
	stake_available_nfts();

    await cosmicGameProxy
        .connect(owner)
        .setCharity(await charityWallet.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setRandomWalkNft(await randomWalkNFT.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setRaffleWallet(await raffleWallet.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletCosmicSignatureNft(await stakingWalletCosmicSignatureNft.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletRandomWalkNft(await stakingWalletRandomWalkNft.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setMarketingWallet(await marketingWallet.getAddress());
    await cosmicGameProxy.connect(owner).setNumRaffleETHWinnersBidding(4);
    await cosmicGameProxy.connect(owner).setNumRaffleNFTWinnersBidding(6);
    await cosmicGameProxy.connect(owner).setNumRaffleNFTWinnersStakingRWalk(3);
    await cosmicGameProxy.connect(owner).setPrizePercentage(30);
    await cosmicGameProxy.connect(owner).setCharityPercentage(5);
    await cosmicGameProxy.connect(owner).setRafflePercentage(6);
    await cosmicGameProxy.connect(owner).setCharity(addr3.address);
    await cosmicGameProxy
        .connect(owner)
        .setCharity(await charityWallet.getAddress());
    await cosmicGameProxy.connect(owner).setStakingPercentage(9);
    await cosmicGameProxy
        .connect(owner)
        .setTokenContract(await cosmicToken.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setNftContract(await cosmicSignature.getAddress());
    await cosmicGameProxy.connect(owner).setRoundStartCstAuctionLength(13 * 3600);
    await cosmicGameProxy.connect(owner).setErc20RewardMultiplier(5);
    await cosmicGameProxy
        .connect(owner)
        .setStartingBidPriceCSTMinLimit(150000000000000000000n);
    await cosmicGameProxy
        .connect(owner)
        .setMarketingReward(120000000000000000000n);
    await cosmicGameProxy.connect(owner).setTokenReward(130000000000000000000n);
    await cosmicGameProxy.connect(owner).setMaxMessageLength(199);
    await cosmicSignature.connect(owner).setTokenGenerationScriptURL("ipfs://");
    await cosmicSignature.connect(owner).setBaseURI("nttp://");
    let iAddrBytes = await ethers.provider.getStorage(
        await cosmicGameProxy.getAddress(),
        "0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc"
    );
    let iAddr = await ethers.AbiCoder.defaultAbiCoder()
        .decode(["address"], iAddrBytes)
        .toString();
    await cosmicGameProxy.connect(owner).upgradeTo(iAddr);
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletCosmicSignatureNft(await stakingWalletCosmicSignatureNft.getAddress());
    await cosmicGameProxy
        .connect(owner)
        .setStakingWalletRandomWalkNft(await stakingWalletRandomWalkNft.getAddress());
    await cosmicGameProxy.connect(owner).setErc20RewardMultiplier(999);
    let tmp = await cosmicGameProxy.timeIncrease();
    await cosmicGameProxy.connect(owner).setTimeIncrease(tmp);
    tmp = await cosmicGameProxy.connect(owner).timeoutClaimPrize();
    await cosmicGameProxy.connect(owner).setTimeoutClaimPrize(tmp);
    tmp = await cosmicGameProxy.priceIncrease();
    await cosmicGameProxy.connect(owner).setPriceIncrease(tmp);
    tmp = await cosmicGameProxy.nanoSecondsExtra();
    await cosmicGameProxy.connect(owner).setNanoSecondsExtra(tmp);
    tmp = await cosmicGameProxy.initialSecondsUntilPrize();
    await cosmicGameProxy.connect(owner).setInitialSecondsUntilPrize(tmp);
    tmp = await cosmicGameProxy.initialBidAmountFraction();
    await cosmicGameProxy.connect(owner).updateInitialBidAmountFraction(tmp);
    tmp = await cosmicGameProxy.activationTime();
    await cosmicGameProxy.connect(owner).setActivationTime(tmp);

    await cosmicGameProxy.connect(owner).setRuntimeMode();

    await cosmicGameProxy.connect(addr3).claimDonatedNFT(0n);
    await cosmicGameProxy.connect(addr3).claimDonatedNFT(1n);
    topic_sig = raffleWallet.interface.getEvent("RaffleDepositEvent").topicHash;
    deposit_logs = receipt.logs.filter((x) => x.topics.indexOf(topic_sig) >= 0);
    let withdrawal_done = [];
    for (let i = 0; i < deposit_logs.length; i++) {
        let wlog = raffleWallet.interface.parseLog(deposit_logs[i]);
        let winner_signer = await hre.ethers.getSigner(wlog.args.winner);
        if (typeof withdrawal_done[wlog.args.winner] === "undefined") {
            await raffleWallet.connect(winner_signer).withdraw();
            withdrawal_done[wlog.args.winner] = 1;
        } else {
            // skip
        }
    }

    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {  msg: "", rwalk: -1 };
    params = ethers.AbiCoder.defaultAbiCoder().encode([bidParamsEncoding],[bidParams]);
    await cosmicGameProxy.connect(addr1).bid(params, { value: bidPrice  });
    bidPrice = await cosmicGameProxy.getBidPrice();
    await cosmicGameProxy.connect(addr1).bid(params, { value: bidPrice  });
    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    await cosmicGameProxy.connect(addr1).claimPrize({
        gasLimit: 30000000
    });

    await ethers.provider.send("evm_mine");
	stake_available_nfts()

    donationData =
        '{"version":1,"title":"EF donation","message":"Ethereum Foundation is a non-profit and part of a community of organizations and people working to fund protocol development, grow the ecosystem, and advocate for Ethereum.","url":"http://ethereum.org/en"}';
    await cosmicGameProxy.donateWithInfo(donationData, {
        value: hre.ethers.parseEther("9"),
    });
    await cosmicGameProxy.donateWithInfo(donationData, {
        value: hre.ethers.parseEther("8"),
    });

    await marketingWallet.send(hre.ethers.parseEther("7"), addr1.address);
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.send(hre.ethers.parseEther("7"), addr3.address);
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.send(hre.ethers.parseEther("2"), addr2.address);
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.send(hre.ethers.parseEther("6"), addr1.address);
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.send(hre.ethers.parseEther("5"), addr2.address);
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.send(hre.ethers.parseEther("5"), addr2.address);
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await marketingWallet.send(hre.ethers.parseEther("5"), addr3.address);
    await marketingWallet.send(hre.ethers.parseEther("1"), addr4.address);
    await marketingWallet.send(hre.ethers.parseEther("11"), addr1.address);
    for (let i = 1; i <= 20; i++) {
        let action_rec = (await stakingWalletCosmicSignatureNft.stakeActions(i)).toObject();
        let ownr = action_rec.nftOwnerAddress;
		if (ownr == "0x0000000000000000000000000000000000000000") { continue; }
        let owner_signer = await hre.ethers.getSigner(ownr);
        await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,2);
    }
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
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
    await ethers.provider.send("evm_mine"); // mine empty block as spacing

    // generate one deposit to charity and not to Staking Wallet
    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 3",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr3).bid(params, {
        value: bidPrice
    });
    bidPrice = await cosmicGameProxy.getBidPrice();
    bidParams = {
        msg: "bid 3",
        rwalk: -1
    };
    params = ethers.AbiCoder.defaultAbiCoder().encode(
        [bidParamsEncoding],
        [bidParams]
    );
    await cosmicGameProxy.connect(addr3).bid(params, {
        value: bidPrice
    });
    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    tx = await cosmicGameProxy.connect(addr3).claimPrize({
        gasLimit: 3000000
    });

    await ethers.provider.send("evm_mine"); // mine empty block as spacing

    ts = await cosmicSignature.totalSupply();

    for (let i = Number(oldTotalSupply); i < Number(ts); i++) {
        let ownr = await cosmicSignature.ownerOf(i);
        if (ownr == (await stakingWalletCosmicSignatureNft.getAddress())) {
            continue;
        }
        let owner_signer = await hre.ethers.getSigner(ownr);
        try {
            await stakingWalletCosmicSignatureNft.connect(owner_signer).stake(i);
        } catch (e) {
            console.log("ignoring stake() error for token " + i + ", owner " + ownr);
        }
    }

    await cosmicToken
        .connect(addr1)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
    await cosmicToken
        .connect(addr2)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
    await cosmicToken
        .connect(addr3)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
    await cosmicToken
        .connect(addr4)
        .approve(
            await cosmicGameProxy.getAddress(),
            hre.ethers.parseEther("10000000")
        );
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await cosmicGameProxy.connect(addr1).bidWithCST("CST bid addr1");
    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr1).claimPrize({
        gasLimit: 3000000
    });

    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await cosmicGameProxy.connect(addr3).bidWithCST("CST bid addr1");
    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimPrize({
        gasLimit: 3000000
    });
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await cosmicGameProxy.connect(addr3).bidWithCST("CST bid addr1");
    prizeTime = await cosmicGameProxy.timeUntilPrize();
    await ethers.provider.send("evm_increaseTime", [Number(prizeTime)]);
    await ethers.provider.send("evm_mine");
    tx = await cosmicGameProxy.connect(addr3).claimPrize({
        gasLimit: 3000000
    });
    await ethers.provider.send("evm_mine"); // mine empty block as spacing
    await ethers.provider.send("evm_mine"); // mine empty block as spacing

    for (let i = 1; i <= 5; i++) {
        await stakingWalletRandomWalkNft.connect(addr1).unstake(i);
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
	        await stakingWalletCosmicSignatureNft.connect(owner_signer).unstake(i,1);
	        await stakingWalletCosmicSignatureNft.connect(owner_signer).payReward(i,1);
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
