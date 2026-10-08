#!/usr/bin/env bash
#
# Same as deploy-v3-local.bash, but NEVER starts or kills a Hardhat node.
# It uses whatever node the `hardhat_on_localhost` network in hardhat.config.js points at
# (local or remote), and assumes that node is already running.
#
# Steps:
#   1. Checks that the configured RPC endpoint responds.
#   2. Deletes stale deploy/upgrade report files (the tasks refuse to run if they exist).
#   3. Deploys all Cosmic Signature contracts (the game proxy is born as V1 - by design,
#      only V1 has the full `initialize`).
#   4. Completes round 0 on V1: one ETH bid + time warp + main prize claim. This is required
#      before upgrading, because the V2+ reinitializers assert `roundNum > 0` and V2+ pricing
#      assumes `ethDutchAuctionBeginningBidPrice` was seeded by a first bid (Comment-202508094).
#      NOTE: this warps the node's block time ~1 hour ahead of real time (to reach
#      `mainPrizeTime`). Keep that in mind if anything else uses the same node.
#   5. Upgrades the proxy to CosmicSignatureGameV2, then to CosmicSignatureGameV3.
#   6. Sets `roundActivationTime` to "now" so round 1 is immediately playable on V3
#      (the claim in step 4 would otherwise delay activation by 30 minutes).
#
# The deployer/owner/bidder private key is taken from
# tasks/config/deploy-cosmic-signature-contracts-config-hardhat_on_localhost.json.
#
# Re-running this script performs a brand new deployment (new contract addresses) on the
# same node; the report files always describe the latest deployment.

set -u -o pipefail

ScriptDirPath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
RepoDirPath="$(dirname "${ScriptDirPath}")"

cd "${RepoDirPath}" || exit 1

PrintStep() {
	echo
	echo "==== ${1} ===="
	echo
}

# #region Step 1. Check that the configured RPC endpoint responds.

# Extract the URL from the `hardhat_on_localhost` network config.
# It may be overridden with the RPC_URL environment variable.
RpcUrl="${RPC_URL:-}"
if [ -z "${RpcUrl}" ]; then
	RpcUrl="$(awk '/hardhat_on_localhost: \{/,/\}/' 'hardhat.config.js' | grep -m 1 'url:' | sed -E 's/.*"([^"]+)".*/\1/')"
fi
if [ -z "${RpcUrl}" ]; then
	echo 'Error. Failed to extract the hardhat_on_localhost RPC URL from hardhat.config.js. Set the RPC_URL environment variable and rerun.' >&2
	exit 1
fi
echo "Using RPC endpoint: ${RpcUrl}"
if ! curl -s -X POST -H 'Content-Type: application/json' \
	--data '{"jsonrpc":"2.0","id":1,"method":"eth_chainId","params":[]}' \
	"${RpcUrl}" 2>/dev/null | grep -q '"result"'; then
	echo "Error. No RPC response from ${RpcUrl}. Is the Hardhat node running?" >&2
	exit 1
fi
echo 'The RPC endpoint responds.'

# #endregion
# #region Step 2. Delete stale report files.

rm -f \
	'tasks/output/deploy-cosmic-signature-contracts-report-hardhat_on_localhost.json' \
	'tasks/output/upgrade-cosmic-signature-game-report-hardhat_on_localhost-CosmicSignatureGameV2.json' \
	'tasks/output/upgrade-cosmic-signature-game-report-hardhat_on_localhost-CosmicSignatureGameV3.json'

# #endregion
# #region Environment for the tasks below. Same as tasks/runners/*.bash.

# The OpenZeppelin upgrades plugin stores its dev-network manifest under `$TMPDIR/openzeppelin-upgrades`.
# On a shared machine, `/tmp/openzeppelin-upgrades` may be owned by another user, causing
# an EACCES error. Point TMPDIR at a private folder to avoid that.
# Note: the manifest links the deployed proxy to the plugin. If you later run the upgrade tasks
# manually (outside this script), export the same TMPDIR first, or the plugin won't find the proxy.
export TMPDIR="${HOME}/.cosmic-signature-tmp"
mkdir -p "${TMPDIR}" || exit 1

export HARDHAT_MODE_CODE='2'
export ENABLE_HARDHAT_PREPROCESSOR='true'
export ENABLE_ASSERTS='true'
export ENABLE_SMTCHECKER='1'

# #endregion
# #region Step 3. Deploy all contracts (game proxy as V1).

PrintStep 'Deploying all Cosmic Signature contracts (game proxy as V1)'
(
	cd 'tasks/runners' &&
	'npx' 'hardhat' 'deploy-cosmic-signature-contracts' \
		'--deployconfigfilepath' '../config/deploy-cosmic-signature-contracts-config-hardhat_on_localhost.json' \
		'--network' 'hardhat_on_localhost'
) || { echo 'Error. Deployment failed.' >&2; exit 1; }

# #endregion
# #region Step 4. Complete round 0 on V1.

PrintStep 'Completing round 0 on V1 (bid + claim)'
'npx' 'hardhat' 'run' 'scripts/complete-round-zero.js' '--network' 'hardhat_on_localhost' ||
	{ echo 'Error. Completing round 0 failed.' >&2; exit 1; }

# #endregion
# #region Step 5. Upgrade the game proxy to V2, then V3.

PrintStep 'Upgrading the game proxy to CosmicSignatureGameV2'
(
	cd 'tasks/runners' &&
	'npx' 'hardhat' 'upgrade-cosmic-signature-game' \
		'--upgradeconfigfilepath' '../config/upgrade-cosmic-signature-game-config-hardhat_on_localhost-CosmicSignatureGameV2.json' \
		'--network' 'hardhat_on_localhost'
) || { echo 'Error. The V2 upgrade failed.' >&2; exit 1; }

PrintStep 'Upgrading the game proxy to CosmicSignatureGameV3'
(
	cd 'tasks/runners' &&
	'npx' 'hardhat' 'upgrade-cosmic-signature-game' \
		'--upgradeconfigfilepath' '../config/upgrade-cosmic-signature-game-config-hardhat_on_localhost-CosmicSignatureGameV3.json' \
		'--network' 'hardhat_on_localhost'
) || { echo 'Error. The V3 upgrade failed.' >&2; exit 1; }

# #endregion
# #region Step 6. Activate round 1.

PrintStep 'Activating round 1 on V3'
'npx' 'hardhat' 'run' 'scripts/activate-round.js' '--network' 'hardhat_on_localhost' ||
	{ echo 'Error. Round activation failed.' >&2; exit 1; }

# #endregion

PrintStep 'Success'
echo 'The game proxy is now running CosmicSignatureGameV3 on round 1, active and ready to play.'
echo 'Contract addresses: tasks/output/deploy-cosmic-signature-contracts-report-hardhat_on_localhost.json'
