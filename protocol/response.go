// protocol/response.go
package protocol

import "errors"

// Response represents a raw Modbus response PDU.
type Response struct {
	TransactionID uint16
	ProtocolID    uint16
	UnitID        uint8

	Function uint8 // raw function byte (may include exception flag)
	Payload  []byte

	// Exception is non-nil only if this is an exception response.
	Exception *ExceptionCode
}

// DecodeTCP decodes a Modbus TCP ADU into a Response.
//
// This function:
// - does NOT validate lengths
// - does NOT interpret payload
// - does NOT convert exceptions into errors
//
// Transport errors are represented ONLY by returned error.
func DecodeTCP(b []byte) (*Response, error) {
	if len(b) < 8 {
		return nil, errors.New("short modbus tcp frame")
	}

	r := &Response{
		TransactionID: uint16(b[0])<<8 | uint16(b[1]),
		ProtocolID:    uint16(b[2])<<8 | uint16(b[3]),
		UnitID:        b[6],
		Function:      b[7],
	}

	if len(b) > 8 {
		r.Payload = b[8:]
	} else {
		r.Payload = nil
	}

	if IsExceptionFunction(r.Function) {
		if len(r.Payload) < 1 {
			return nil, errors.New("exception response missing code")
		}
		ex := ExceptionCode(r.Payload[0])
		r.Exception = &ex
	}

	return r, nil
}
