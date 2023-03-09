package main

import (
	. "github.com/PredictionExplorer/augur-explorer/contracts"
)
const (

	PRIZE_CLAIM_EVENT		= "27bc828c399c2947fea27bca8a75ced2e94ff2651d607271f051e39db52286ce"
	BID_EVENT				= "22664806da305c71fc17312c8fafba247a01ea6a4ee8a2e97f06cd21d979eefe"
	DONATION_EVENT			= "8b7fe5be5699654fd637d2250cb0d47e88205730710745e78e9d8bcaf8aad8f1"
	DONATION_RECEIVED_EVENT	= "46ff3d75d4645bdbbae4cd6109ba42e6e1b80ea25e69d10472b357b733300368"
	DONATION_SENT_EVENT		= "44d398d152fa0735a428b13ebc78f79fe4cb1b4722292cd233e278552fa36d32"
	CHARITY_UPDATED			= "a0bd6b2fdbf082ae2356710c23fc8d76d56d418cecb4514d119c77a8617b4ffe"
	TOKEN_NAME_EVENT		= "8ad5e159ff95649c8a9f323ac5a457e741897cf44ce07dfce0e98b84ef9d5f12"
	MINT_EVENT				= "af162acd8d98cd428a10ad5028b47e3d3e50b4089880be8bf474aa921fed6b2e"
)
var (
	evt_prize_claim_event,_ = hex.DecodeString(PRIZE_CLAIM_EVENT)
	evt_bid_event,_			= hex.DecodeString(BID_EVENT)
	evt_donation_event,_	= hex.DecodeString(DONATION_EVENT)
	evt_donation_received_event,_=hex.DecodeString(DONATION_RECEIVED_EVENT)
	evt_donation_seent_event,_= hex.DecodeString(DONATION_SENT_EVENT)
	evt_charity_updated,_	= hex.DecodeString(CHARITY_UPDATED)
	evt_token_name_event,_	= hex.DecodeString(TOKEN_NAME_EVENT)
	evt_mint_event,_		= hex.DecodeString(MINT_EVENT)

	biddingwar_abi			*abi.ABI
	cosmic_signature_abi	*abi.ABI
	cosmic_token_abi		*abi.ABI
	charity_wallet_abi		*abi.ABI

	biddingwar_addr			common.Address
	cosmic_signature_addr	common.Address
	cosmic_token_addr		common.Address
	charity_wallet_addr		common.Address
)

