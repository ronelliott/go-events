package events

import (
	"bufio"
	"io"
	"net"

	log "github.com/sirupsen/logrus"
)

// An event server interface
type Server interface {
	// Close any open resources
	Close() error

	// Listen on the given address
	Listen(string, chan<- error)
}

// The default server implementation
type ServerImpl struct{}

// Create a new server
func NewServer() Server {
	return &ServerImpl{}
}

// Close any open resources
func (server *ServerImpl) Close() error {
	return nil
}

// Handle the given connection
func (server *ServerImpl) Handle(raw net.Conn) error {
	remoteAddr := raw.RemoteAddr()
	log.Debug("Client connected: ", remoteAddr)
	defer log.Debug("Client disconnected: ", remoteAddr)
	conn := bufio.NewReadWriter(bufio.NewReader(raw), bufio.NewWriter(raw))
	defer raw.Close()

	for {
		line, err := conn.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				return nil
			}

			return err
		}

		log.Debug("Received: ", line, " from ", remoteAddr)
	}

	return nil
}

// Listen on the given address
func (server *ServerImpl) Listen(address string, errs chan<- error) {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		errs <- err
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			errs <- err
			continue
		}

		go func() {
			err = server.Handle(conn)

			if err != nil {
				errs <- err
			}
		}()
	}
}
