// client.go
package modbusclient

import (
	"modbus-client/protocol"
	"modbus-client/transport/tcp"
)

// Client wires protocol and transport together.
// It contains NO behavior beyond orchestration.
type Client struct {
	Transport *tcp.Client
}

// Do sends a single Modbus request and returns the raw response.
//
// Contract:
//   request → response | transport error
//
// No retries.
// No interpretation.
// No state.
func (c *Client) Do(req *protocol.Request) (*protocol.Response, error) {
	rawReq := req.EncodeTCP()

	rawResp, err := c.Transport.Send(rawReq)
	if err != nil {
		return nil, err
	}

	return protocol.DecodeTCP(rawResp)
}
