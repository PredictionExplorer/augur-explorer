package srvmonitor

import "sync"

// SharedRPCState holds the latest official-node block numbers so database
// and application monitors can compute their lag against the chain tip.
type SharedRPCState struct {
	mutex              sync.RWMutex
	officialMainnet    int64
	officialArbitrum   int64
	officialSepolia    int64
	officialSepoliaArb int64
}

// NewSharedRPCState creates a new shared RPC state.
func NewSharedRPCState() *SharedRPCState {
	return &SharedRPCState{}
}

// UpdateOfficial records the official block number for a chain id. Unknown
// chain ids are ignored.
func (s *SharedRPCState) UpdateOfficial(chainID string, blockNum int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	switch chainID {
	case chainIDMainnet:
		s.officialMainnet = blockNum
	case chainIDSepolia:
		s.officialSepolia = blockNum
	case chainIDArbitrum:
		s.officialArbitrum = blockNum
	case chainIDSepoliaArb:
		s.officialSepoliaArb = blockNum
	}
}

// Official returns the recorded official block number for a chain id, or 0
// when the chain id is unknown or nothing was recorded yet.
func (s *SharedRPCState) Official(chainID string) int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	switch chainID {
	case chainIDMainnet:
		return s.officialMainnet
	case chainIDSepolia:
		return s.officialSepolia
	case chainIDArbitrum:
		return s.officialArbitrum
	case chainIDSepoliaArb:
		return s.officialSepoliaArb
	default:
		return 0
	}
}

// Chain ids the monitors distinguish for official-node lag tracking.
const (
	chainIDMainnet    = "1"
	chainIDSepolia    = "11155111"
	chainIDArbitrum   = "42161"
	chainIDSepoliaArb = "421614"
)
