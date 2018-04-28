package events

import (
	"errors"
	"net"
)

// The client interface
type Client interface {
	// Close any open resources
	Close() error

	// Send the given data to the server
	Send([]byte) error
}

// The default client implementation
type ClientImpl struct {
	conn net.Conn
}

// Create a new client from the given URI
func NewClient(address string) (Client, error) {
	conn, err := net.Dial("tcp", address)

	if err != nil {
		return nil, err
	}

	return &ClientImpl{conn}, nil
}

// Close any open resources
func (client *ClientImpl) Close() error {
	if client.conn != nil {
		return client.conn.Close()
	}

	return nil
}

// Send the given data on the given channel
func (client *ClientImpl) Send(data []byte) error {
	if client.conn == nil {
		return errors.New("Not connected")
	}

	_, err := client.conn.Write(data)

	if err != nil {
		return err
	}

	_, err = client.conn.Write([]byte{'\n'})
	return err
}
