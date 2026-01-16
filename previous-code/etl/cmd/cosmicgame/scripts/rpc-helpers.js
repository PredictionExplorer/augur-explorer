async function callContractMethodWithRetry(contractInstance, methodName, maxRetries = Infinity, delay = 5000, ...methodArgs) {
    let tx; 
    let attempt = 0;

    while (true) {
        try {
            let caddr = await contractInstance.getAddress();
            console.log(`Attempting to call ${methodName} on contract ${caddr}... (Attempt ${attempt + 1})`);

            // Dynamically call the method on the contract instance
            tx = await contractInstance[methodName](...methodArgs);
            await tx.wait(); // Wait for transaction confirmation

            console.log(`${methodName} transaction successful!`);
            break; // Exit loop if successful
        } catch (error) {
            console.error(`${methodName} transaction failed: ${error.message}`);

            attempt++;
            if (attempt >= maxRetries) {
                console.error("Max retries reached. Aborting transaction.");
                throw error; // Stop execution after max retries
            }

            console.log(`Retrying in ${delay / 1000} seconds...`);
            await new Promise(resolve => setTimeout(resolve, delay)); // Wait before retrying
        }
    }

    return tx; // Return the successful transaction
}

async function deployWithRetry(factory, maxRetries = 10, delay = 5000, ...deployArgs) {
    let contractInstance;
    let attempt = 0;

    while (true) {
        try {
            console.log(`Attempting deployment... (Attempt ${attempt + 1})`);
            console.log(deployArgs);
            contractInstance = await factory.deploy(...deployArgs); // Spread deploy arguments
            await contractInstance.waitForDeployment(); // Ensure deployment is complete
            console.log("Deployment successful!");
            break; // Exit loop if deployment succeeds
        } catch (error) {
            console.error(`Deployment failed: ${error.message}`);

            attempt++;
            if (attempt >= maxRetries) {
                console.error("Max retries reached. Aborting deployment.");
                throw error; // Rethrow if max retries exceeded
            }

            console.log(`Retrying in ${delay / 1000} seconds...`);
            await new Promise(resolve => setTimeout(resolve, delay)); // Wait before retrying
        }
    }

    return contractInstance;
}
module.exports = {deployWithRetry,callContractMethodWithRetry};
