/**
 * Deploy two Samp ERC20 contracts for use with populate.js.
 *
 * Run from Cosmic-Signature (or any project that has Hardhat + Samp contract):
 *   NODE_PATH=$PWD/node_modules npx hardhat run /path/to/deploy-samp.js --network localhost
 *
 * Then set TSAMP1 and TSAMP2 to the printed addresses when running populate.js.
 */
const hre = require("hardhat");

async function main() {
    const [deployer] = await hre.ethers.getSigners();
    const deployerAddr = await deployer.getAddress();
    console.log("Deploying Samp contracts with account:", deployerAddr);

    const Samp = await hre.ethers.getContractFactory("Samp");

    const samp1 = await Samp.connect(deployer).deploy("ERC20 Token Sample1", "SAMP1");
    await samp1.waitForDeployment();
    const addr1 = await samp1.getAddress();

    const samp2 = await Samp.connect(deployer).deploy("ERC20 Token Sample2", "SAMP2");
    await samp2.waitForDeployment();
    const addr2 = await samp2.getAddress();

    console.log("Samp1 deployed to:", addr1);
    console.log("Samp2 deployed to:", addr2);
    console.log("");
    console.log("Set these for populate.js:");
    console.log("export TSAMP1=" + addr1);
    console.log("export TSAMP2=" + addr2);
}

main()
    .then(() => process.exit(0))
    .catch((err) => {
        console.error(err);
        process.exit(1);
    });
