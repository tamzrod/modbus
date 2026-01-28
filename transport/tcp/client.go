// transport/tcp/client.go
package tcp

import (
	"net"
	"time"
)

// Client is a minimal Modbus TCP transport.
type Client struct {
	Conn    net.Conn
	Timeout time.Duration
}

// Send sends a request and reads a single response.
//
// This function performs:
// - exactly one write
// - exactly one read
//
// Any failure here is a TRANSPORT error.
func (c *Client) Send(req []byte) ([]byte, error) {
	if c.Timeout > 0 {
		_ = c.Conn.SetDeadline(time.Now().Add(c.Timeout))
	}

	_, err := c.Conn.Write(req)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 260) // max Modbus TCP ADU
	n, err := c.Conn.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf[:n], nil
}
