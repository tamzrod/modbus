// protocol/frame.go
package protocol

// FunctionCode represents a Modbus function code.
type FunctionCode uint8

// Modbus function codes (partial; extend only when needed).
const (
	FCReadCoils              FunctionCode = 1
	FCReadDiscreteInputs     FunctionCode = 2
	FCReadHoldingRegisters   FunctionCode = 3
	FCReadInputRegisters     FunctionCode = 4
)

// ExceptionFlag is ORed with a function code in exception responses.
const ExceptionFlag uint8 = 0x80

// ExceptionCode represents a Modbus exception code.
// These are DATA, not errors.
type ExceptionCode uint8

// Modbus exception codes (truth-preserving).
const (
	ExIllegalFunction        ExceptionCode = 1
	ExIllegalDataAddress     ExceptionCode = 2
	ExIllegalDataValue       ExceptionCode = 3
	ExSlaveDeviceFailure     ExceptionCode = 4
	ExAcknowledge            ExceptionCode = 5
	ExSlaveDeviceBusy        ExceptionCode = 6
)

// IsExceptionFunction returns true if the function code indicates
// an exception response (FC | 0x80).
func IsExceptionFunction(fc uint8) bool {
	return fc&ExceptionFlag != 0
}
