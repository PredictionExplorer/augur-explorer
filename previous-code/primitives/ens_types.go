package primitives
import (

)
type ENSResolver struct {
	Aid						int64
	Node					string
}
type ENS_TextKeyValue struct {	// the data extracted from TextChanged event of ENS
	Key						string
	Value					string
}
type ENS_PubkeyChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	TxHash					string
	Node					string
	X						string
	Y						string
	DerivedAddr				string	// address calcualted from X/Y
}
type ENS_ContentHashChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	TxHash					string
	Node					string
	Hash					string
}
type ENS_NameChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	TxHash					string
	Node					string
	Name					string
}
type ENS_NodeAddr struct {
	Aid						int64
	AddressSetTs			int64
	BlockNum				int64
	Address					string
	AddressSetDate			string
	TxHash					string
}
type ENS_NodeTextInfo  struct {
	FQDN					string
	TextMetaInfo			[]ENS_TextKeyValue
}
type ENS_Info struct {
	LastOwnerAddr			string
	FirstRegisteredTs		int64
	NumTextKeyValuePairs	int64
	TsExpiration			int64
	CurAddrAid				int64
	DeactivatedTimeStamp	int64
	Active					bool
	FirstRegisteredDate		string
	ENS_Name				string
	Label					string
	Node					string
	FQDN					string
	ContentHash				string
	PublicKey_X				string
	PublicKey_Y				string
	CurAddr					string // current address (the last one, empty if non-existent)
	DeactivatedDate			string
	DeactivationTxHash		string
	AddressChangeHistory	[]ENS_NodeAddr
	OwnershipChangeHistory	[]ENS_NewOwner
	NameTextMetaInfo		ENS_NodeTextInfo
}
type EnsProcStatus struct {
	IniLoadBlockNumLimit	int64
	LastEvtId				int64
}
type ENS_Name1  struct {	// NameRegistered since 2019
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Expires					int64
	Contract				string
	Owner					string
	Label					string
	Node					string
	FQDN					string
	Name					string
	Cost					string
	TxHash					string
}
type ENS_Name2  struct {	// NameRegistered old version
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Expires					int64
	Contract				string
	Label					string
	Node					string
	FQDN					string
	Owner					string
	TxHash					string
}
type ENS_Name3  struct {	// NameRegistered old version
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	CreatedDate				int64
	Contract				string
	Caller					string
	Beneficiary				string
	Node					string
	Label					string
	Subdomain				string
	FQDN					string
	TxHash					string
}
type ENS_NameRenewed  struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Expires					int64
	Contract				string
	Label					string
	Node					string
	FQDN					string
	Name					string
	Cost					string
	TxHash					string
}
type ENS_NameMigrated  struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Expires					int64
	Contract				string
	Owner					string
	Label					string
	Node					string
	FQDN					string
	TxHash					string
}
type ENS_Transfer struct {
	From					string
	To						string
}
type ENS_NewOwner struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Expires					int64
	Contract				string
	Owner					string
	Label					string
	Node					string
	TxHash					string
	FQDN					string	// fully qualified domain name
	DateTime				string
	Name					string
}
type ENS_HashInvalidated struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	RegistrationDate		int64
	Contract				string
	Hash					string
	Name					string
	TxHash					string
	Value					string
}
type ENS_HashRegistered struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	RegistrationDate		int64
	Contract				string
	Hash					string
	Owner					string
	TxHash					string
	Value					string
}
type ENS_NewResolver struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	Node					string
	Address					string
	NameAddr				string
	TxHash					string
}
type ENS_RegistryTransfer struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	TxHash					string
	Owner					string
	Node					string
}
type ENS_RegistrarTransfer struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	TxHash					string
	From					string
	To						string
	TokenId					string
	Label					string
	Node					string
	FQDN					string
}
type ENS_TextChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	TxHash					string
	Node					string
	Key						string
	Value					string
}
type ENS_Rec struct {
	Word					string
	Label					string
	FQDN					string	// aka 'node'
}
type ENS_AddrChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	TxHash					string
	Contract				string
	FQDN					string
	Address					string
}
type ENS_AddressChanged struct {
	EvtId					int64
	BlockNum				int64
	TxId					int64
	TimeStamp				int64
	Contract				string
	TxHash					string
	CoinType				int
	FQDN					string
	Address					string
}
type ENS_NodeShort struct {
	CurOwnerAid				int64
	Id						int64
	Label					string
	Node					string
	FQDN					string
	FQDN_Words				string
}
type ENS_MultiAddress struct { // all possible addresses of a Node
	OwnerAddrTs				int64
	AddrChgTs				int64	// AddrChanged event
	NewResAddrTs			int64
	OwnerAddr				string
	AddrChgAddr				string
	NewResAddr				string
}
type UserAddrChange struct {
	TimeStamp				int64
	CoinType				int
	DateTime				string
	FQDN					string
	FQDN_Words				string
}
type UserOwnershipChange struct {
	TimeStamp				int64
	BlockNum				int64
	DateTime				string
	Label					string
	Node					string
	FQDN					string
	FQDN_Words				string
	TxHash					string
}
type UserENS struct {
	TsNameAcquired			int64
	NumTextKeyValuePairs	int64
	TsExpiration			int64
	CurAddrAid				int64
	DateNameAcquired		string
	ENS_Name				string
	NodeHash				string
	ContentHash				string
	PublicKey_X				string
	PublicKey_Y				string
	PublicKey_Addr			string
	CurAddr					string // current address (the last one, empty if non-existent)
	NodeAddressHistory		[]ENS_NodeAddr
}
