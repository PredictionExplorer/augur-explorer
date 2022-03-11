

func process_transaction(storage *SQLStorage,tx *AugurTx,rcpt *types.Receipt) {


	logs := rcpt.Logs
	for i:=0; i<len(logs); i++ {
		log := logs[i]
		process_event_log(&log)
	}
}
func process_pool_created(tx *AugurTx,log *types.Log) {

	var eth_evt BalancerV2WeightedPoolFactoryPoolCreated
	err := condtoken_abi.UnpackIntoInterface(&eth_evt,"PositionSplit",log.Data)

	var evt BalV2PoolCreated
	evt
}
func process_pool_balance_changed(log *types.Log) {

}
func process_pool_registered(log *types.Log) {

}
func process_external_balance_transfer(log *types.Log) {

}
func process_event_log(log *types.Log) {

	if len(log.Topics) > 0 { return }
	topic0 := log.Topics[0].bytes()
	if bytes.Equal(topic0,evt_pool_created) {
		process_pool_created(log)
	}
	if bytes.Equal(topic0,evt_pool_balance_changed) {
		process_pool_balance_changed(log)
	}
	if bytes.Equal(topic0,evt_pool_registered) {
		process_pool_registered(log)
	}
	if bytes.Equal(topic0,evt_external_balance_transfer) {
		process_external_balance_transfer(log)
	}
	if bytes.Equal(topic0,evt_internal_balance_changed) {
		process_internal_balance_changed(log)
	}
	if bytes.Equal(topic0,evt_topic_deregistered) {
		process_token_deregistered(log)
	}
	if bytes.Equal(topic0,evt_topic_registered) {
		process_token_registered(log)
	}
	if bytes.Equal(topic0,evt_swap) {
		process_process_swap(log)
	}
}
