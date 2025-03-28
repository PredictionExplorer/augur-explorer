// [Comment-202409255]
// Because "hardhat.config.js" imports us, an attempt to import "hardhat" here would throw an error.
// So we must do things differently here.
// [/Comment-202409255]
//const hre = require("hardhat");
//const { HardhatContext } = require("hardhat/internal/context");
//const { getHardhatRuntimeEnvironment } = require("hardhat/plugins");
//const hre = getHardhatRuntimeEnvironment();
const hre = require("hardhat");
const basicDeployment = async function (
	deployerAcct,
	randomWalkAddr,
	activationTime,
	charityAddr,
	transferOwnership,
	switchToRuntime = true,
) {
	return await basicDeploymentAdvanced(
		"CosmicSignatureGame",
		deployerAcct,
		randomWalkAddr,
		activationTime,
		charityAddr,
		transferOwnership,
		switchToRuntime,
	);
};
const basicDeploymentAdvanced = async function (
	cgpName,
	deployerAcct,
	randomWalkAddr,
	activationTime,
	charityAddr,
	transferOwnership,
	switchToRuntime,
) {
	if (switchToRuntime === undefined) {
		console.error("switchToRuntime is not set");
		process.exit(1);
	}

	// Comment-202409255 applies.
//	const hre = HardhatContext.getHardhatContext().environment;

	let cosmicGameProxy, cosmicToken, cosmicSignature, charityWallet, cosmicDAO, randomWalkNFT, prizesWallet;
	let CosmicGame = await hre.ethers.getContractFactory(cgpName);
	cosmicGameProxy = await hre.upgrades.deployProxy(
		CosmicGame,
		args = [deployerAcct.address],
		opts = {
			kind: "uups"
		}
	);
	const cosmicGameProxyAddr = await cosmicGameProxy.getAddress();
	let implementationAddrBytes = await cosmicGameProxy.runner.provider.getStorage(cosmicGameProxyAddr,'0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc');
    let implementationAddr = ethers.AbiCoder.defaultAbiCoder().decode(['address'],implementationAddrBytes);
	let cosmicGameAddr = await cosmicGameProxy.runner.provider.getStorage(cosmicGameProxyAddr,'0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc');
	let cosmicGame = await CosmicGame.attach(cosmicGameProxyAddr);
	let CosmicToken = await hre.ethers.getContractFactory("CosmicSignatureToken");
	cosmicToken = await CosmicToken.connect(deployerAcct).deploy(await cosmicGameProxy.getAddress());
	await cosmicToken.waitForDeployment();
	let cosmicTokenAddr = await cosmicToken.getAddress();

	let CosmicSignature = await hre.ethers.getContractFactory("CosmicSignatureNft");
	cosmicSignature = await CosmicSignature.connect(deployerAcct).deploy(cosmicGameProxyAddr);
	await cosmicSignature.waitForDeployment();
	let cosmicSignatureAddr = await cosmicSignature.getAddress();

	let CosmicDAO = await hre.ethers.getContractFactory("CosmicSignatureDao");
	cosmicDAO = await CosmicDAO.connect(deployerAcct).deploy(cosmicTokenAddr);
	await cosmicDAO.waitForDeployment();
	let cosmicDAOAddr = await cosmicDAO.getAddress();

	let CharityWallet = await hre.ethers.getContractFactory("CharityWallet");
	charityWallet = await CharityWallet.connect(deployerAcct).deploy();
	await charityWallet.waitForDeployment();
	let charityWalletAddr = await charityWallet.getAddress();
	if (charityAddr.length == 0) {
		const [owner, otherAccount] = await hre.ethers.getSigners();
		charityAddr = otherAccount.address;
	}
	await charityWallet.setCharityAddress(charityAddr);
	if (transferOwnership) {
		await charityWallet.connect(deployerAcct).transferOwnership(cosmicDAOAddr);
	}

    let PrizesWallet = await hre.ethers.getContractFactory("PrizesWallet");
    prizesWallet = await PrizesWallet.connect(deployerAcct).deploy(cosmicGameProxyAddr);
    await prizesWallet.waitForDeployment();
    let prizesWalletAddr = await prizesWallet.getAddress();

	let MarketingWallet = await hre.ethers.getContractFactory("MarketingWallet");
	const marketingWallet = await MarketingWallet.connect(deployerAcct).deploy(cosmicToken);
	await marketingWallet.waitForDeployment();
	let marketingWalletAddr = await marketingWallet.getAddress();

	let RandomWalkNFT = await hre.ethers.getContractFactory("RandomWalkNFT");
	if (randomWalkAddr.length === 0) {
		randomWalkNFT = await RandomWalkNFT.connect(deployerAcct).deploy();
		await randomWalkNFT.waitForDeployment();
		randomWalkAddr = await randomWalkNFT.getAddress();
	} else {
		randomWalkNFT = await hre.ethers.getContractAt("RandomWalkNFT", randomWalkAddr);
	}
	let randomWalkNFTAddr = await randomWalkNFT.getAddress();

	let StakingWalletCosmicSignatureNft = await hre.ethers.getContractFactory("StakingWalletCosmicSignatureNft");
	let stakingWalletCosmicSignatureNft = await StakingWalletCosmicSignatureNft.connect(deployerAcct).deploy(
		await cosmicSignature.getAddress(),
		cosmicGameProxyAddr
		// charityAddr
	);
	await stakingWalletCosmicSignatureNft.waitForDeployment();
	let stakingWalletCosmicSignatureNftAddr = await stakingWalletCosmicSignatureNft.getAddress();

	let StakingWalletRandomWalkNft = await hre.ethers.getContractFactory("StakingWalletRandomWalkNft");
	const stakingWalletRandomWalkNft = await StakingWalletRandomWalkNft.connect(deployerAcct).deploy(randomWalkAddr);
	await stakingWalletRandomWalkNft.waitForDeployment();
	let stakingWalletRandomWalkNftAddr = await stakingWalletRandomWalkNft.getAddress();

	await cosmicGameProxy.connect(deployerAcct).setCosmicSignatureToken(cosmicTokenAddr);
	await cosmicGameProxy.connect(deployerAcct).setCosmicSignatureNft(cosmicSignatureAddr);
	await cosmicGameProxy.connect(deployerAcct).setCharityAddress(charityWalletAddr);
	await cosmicGameProxy.connect(deployerAcct).setRandomWalkNft(randomWalkNFTAddr);
	await cosmicGameProxy.connect(deployerAcct).setPrizesWallet(prizesWalletAddr);
	await cosmicGameProxy.connect(deployerAcct).setStakingWalletCosmicSignatureNft(stakingWalletCosmicSignatureNftAddr);
	await cosmicGameProxy.connect(deployerAcct).setStakingWalletRandomWalkNft(stakingWalletRandomWalkNftAddr);
	await cosmicGameProxy.connect(deployerAcct).setMarketingWallet(marketingWalletAddr);
	await cosmicGameProxy.connect(deployerAcct).setDelayDurationBeforeRoundActivation(1);
	if (activationTime == 0) {
		let latestBlock = await hre.ethers.provider.getBlock("latest");
		await cosmicGameProxy.connect(deployerAcct).setRoundActivationTime(0);
	} else {
		await cosmicGameProxy.connect(deployerAcct).setRoundActivationTime(0);
	}
	/*
	if (switchToRuntime) {
		await cosmicGameProxy.connect(deployerAcct).setRuntimeMode();
	}*/
	return {
		cosmicGameProxy,
		cosmicToken,
		cosmicSignature,
		charityWallet,
		cosmicDAO,
		prizesWallet,
		randomWalkNFT,
		stakingWalletCosmicSignatureNft,
		stakingWalletRandomWalkNft,
		marketingWallet,
		implementationAddr,
	};
};
module.exports = { basicDeployment, basicDeploymentAdvanced };
