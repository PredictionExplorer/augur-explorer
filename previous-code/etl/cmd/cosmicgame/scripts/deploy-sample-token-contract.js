
const { deployWithRetry,callContractMethodWithRetry} = require("./rpc-helpers.js");

async function main() {

	const Samp = await hre.ethers.getContractFactory("Samp");
    const samp1 = await deployWithRetry(Samp,10,1000);
    let samp1Addr = await samp1.getAddress();
	console.log("sample 1 address = "+samp1Addr);
    const samp2 = await deployWithRetry(Samp,10,1000);
    let samp2Addr = await samp2.getAddress();
	console.log("sample 2 address = "+samp2Addr);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });


