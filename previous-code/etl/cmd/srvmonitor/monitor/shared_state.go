package monitor

import "sync"

// SharedRPCState holds shared RPC information for lag calculations
type SharedRPCState struct {
	mutex              sync.RWMutex
	officialMainnet    int64
	officialArbitrum   int64
	officialSepolia    int64
	officialSepoliaArb int64
}

// NewSharedRPCState creates a new shared RPC state
func NewSharedRPCState() *SharedRPCState {
	return &SharedRPCState{}
}

// UpdateOfficialMainnet updates the official Mainnet block number
func (s *SharedRPCState) UpdateOfficialMainnet(blockNum int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.officialMainnet = blockNum
}

// UpdateOfficialArbitrum updates the official Arbitrum block number
func (s *SharedRPCState) UpdateOfficialArbitrum(blockNum int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.officialArbitrum = blockNum
}

// UpdateOfficialSepolia updates the official Sepolia block number
func (s *SharedRPCState) UpdateOfficialSepolia(blockNum int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.officialSepolia = blockNum
}

// UpdateOfficialSepoliaArb updates the official Sepolia Arbitrum block number
func (s *SharedRPCState) UpdateOfficialSepoliaArb(blockNum int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.officialSepoliaArb = blockNum
}

// GetOfficialMainnet returns the official Mainnet block number
func (s *SharedRPCState) GetOfficialMainnet() int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.officialMainnet
}

// GetOfficialArbitrum returns the official Arbitrum block number
func (s *SharedRPCState) GetOfficialArbitrum() int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.officialArbitrum
}

// GetOfficialSepolia returns the official Sepolia block number
func (s *SharedRPCState) GetOfficialSepolia() int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.officialSepolia
}

// GetOfficialSepoliaArb returns the official Sepolia Arbitrum block number
func (s *SharedRPCState) GetOfficialSepoliaArb() int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.officialSepoliaArb
}






