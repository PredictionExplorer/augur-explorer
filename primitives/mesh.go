package primitives

import (
	"github.com/0xProject/0x-mesh/zeroex"
)

func Get_mesh_event_code(mesh_code zeroex.OrderEventEndState) MeshEvtCode {

	switch mesh_code {
	case zeroex.ESInvalid :
		return 1
	case zeroex.ESOrderAdded :
		return 2
	case zeroex.ESOrderFilled :
		return 3
	case zeroex.ESOrderFullyFilled :
		return 4
	case zeroex.ESOrderCancelled :
		return 5
	case zeroex.ESOrderExpired :
		return 6
	case zeroex.ESOrderUnexpired :
		return 7
	case zeroex.ESOrderBecameUnfunded :
		return 8
	case zeroex.ESOrderFillabilityIncreased :
		return 9
	case zeroex.ESStoppedWatching :
		return 10
	}

	return 0		// Order is obtained with GetOrders() RPC call
}

