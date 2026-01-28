// protocol/request.go
package protocol

import "encoding/binary"

// Request represents a raw Modbus request.
// It is a wire description, not an interpreted command.
type Request struct {
	TransactionID uint16
	ProtocolID    uint16 // always 0 for Modbus TCP, but preserved as data
	UnitID        uint8
	Function      FunctionCode
	Payload       []byte // function-specific data, uninterpreted
}

// EncodeTCP encodes the request into a Modbus TCP ADU.
//
// MBAP Header (7 bytes):
//   Transaction ID (2)
//   Protocol ID    (2)
//   Length         (2)  -> UnitID + PDU
//   Unit ID        (1)
//
// PDU:
//   Function Code  (1)
//   Payload        (N)
//
// This function performs NO validation.
func (r *Request) EncodeTCP() []byte {
	pduLen := 1 + len(r.Payload)
	length := uint16(1 + pduLen) // UnitID + PDU

	out := make([]byte, 7+pduLen)

	binary.BigEndian.PutUint16(out[0:2], r.TransactionID)
	binary.BigEndian.PutUint16(out[2:4], r.ProtocolID)
	binary.BigEndian.PutUint16(out[4:6], length)

	out[6] = r.UnitID
	out[7] = uint8(r.Function)

	copy(out[8:], r.Payload)

	return out
}
