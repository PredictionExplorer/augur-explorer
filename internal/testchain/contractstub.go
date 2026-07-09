package testchain

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// MethodHandler computes the outputs of one stubbed contract method from its
// decoded inputs. args and the returned values use the Go types produced and
// consumed by go-ethereum's ABI codec (*big.Int for uint256, common.Address
// for address, uint8 for uint8, string, ...).
type MethodHandler func(args []any) ([]any, error)

// ContractStub turns one or more contract ABIs into a CallHandler: incoming
// eth_call calldata is matched by 4-byte selector, and the configured return
// values (static via Return, dynamic via Handle) are ABI-encoded back.
// Methods that were never stubbed answer with an error, which the client
// surfaces exactly like a revert on a real node.
//
// The zero value is not usable; create instances with NewContractStub or
// MustContractStub. All methods are safe for concurrent use, so tests may
// re-stub values while a server is answering calls.
type ContractStub struct {
	mu      sync.RWMutex
	methods map[[4]byte]*stubMethod
	byName  map[string]*stubMethod
}

type stubMethod struct {
	method  abi.Method
	returns []any
	handler MethodHandler
}

// NewContractStub parses the given ABI JSON documents into one method table.
// When several ABIs define the same method (same selector), the first one
// wins; this lets a single address serve, for example, both a token's own
// ABI and the generic ERC-20 ABI.
func NewContractStub(abiJSONs ...string) (*ContractStub, error) {
	s := &ContractStub{
		methods: make(map[[4]byte]*stubMethod),
		byName:  make(map[string]*stubMethod),
	}
	for _, doc := range abiJSONs {
		parsed, err := abi.JSON(strings.NewReader(doc))
		if err != nil {
			return nil, fmt.Errorf("contractstub: parsing ABI: %w", err)
		}
		for _, method := range parsed.Methods {
			var sel [4]byte
			copy(sel[:], method.ID)
			if _, exists := s.methods[sel]; exists {
				continue
			}
			m := &stubMethod{method: method}
			s.methods[sel] = m
			s.byName[method.Name] = m
		}
	}
	return s, nil
}

// MustContractStub is NewContractStub for test setup: it panics on a bad ABI.
func MustContractStub(abiJSONs ...string) *ContractStub {
	s, err := NewContractStub(abiJSONs...)
	if err != nil {
		panic(err)
	}
	return s
}

// Return stubs a method (by its Solidity name, e.g. "roundNum") with static
// output values, replacing any previous stub. It panics when the method does
// not exist in the parsed ABIs or when the value count does not match the
// method's outputs, so typos fail loudly at test setup.
func (s *ContractStub) Return(name string, values ...any) *ContractStub {
	m := s.lookup(name)
	if len(values) != len(m.method.Outputs) {
		panic(fmt.Sprintf("contractstub: method %s returns %d values, stubbed with %d",
			name, len(m.method.Outputs), len(values)))
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	m.returns = values
	m.handler = nil
	return s
}

// Handle stubs a method with a dynamic handler that receives the decoded call
// arguments; it replaces any previous stub. Panics on unknown method names.
func (s *ContractStub) Handle(name string, handler MethodHandler) *ContractStub {
	m := s.lookup(name)
	s.mu.Lock()
	defer s.mu.Unlock()
	m.handler = handler
	m.returns = nil
	return s
}

func (s *ContractStub) lookup(name string) *stubMethod {
	s.mu.RLock()
	defer s.mu.RUnlock()
	m, ok := s.byName[name]
	if !ok {
		panic(fmt.Sprintf("contractstub: method %q not found in the parsed ABIs", name))
	}
	return m
}

// Handler adapts the stub to the Chain.RegisterCall signature.
func (s *ContractStub) Handler() CallHandler {
	return func(input []byte) ([]byte, error) {
		if len(input) < 4 {
			return nil, fmt.Errorf("contractstub: calldata shorter than a selector (%d bytes)", len(input))
		}
		var sel [4]byte
		copy(sel[:], input[:4])

		s.mu.RLock()
		m, ok := s.methods[sel]
		var (
			returns []any
			handler MethodHandler
		)
		if ok {
			returns = m.returns
			handler = m.handler
		}
		s.mu.RUnlock()

		if !ok {
			return nil, fmt.Errorf("contractstub: no ABI method for selector 0x%x", sel)
		}
		if handler != nil {
			args, err := m.method.Inputs.Unpack(input[4:])
			if err != nil {
				return nil, fmt.Errorf("contractstub: unpacking %s inputs: %w", m.method.Name, err)
			}
			outs, err := handler(args)
			if err != nil {
				return nil, err
			}
			return s.pack(m, outs)
		}
		if returns != nil {
			return s.pack(m, returns)
		}
		return nil, fmt.Errorf("contractstub: method %s has no stubbed return", m.method.Name)
	}
}

func (s *ContractStub) pack(m *stubMethod, values []any) ([]byte, error) {
	out, err := m.method.Outputs.Pack(values...)
	if err != nil {
		return nil, fmt.Errorf("contractstub: packing %s outputs: %w", m.method.Name, err)
	}
	return out, nil
}
