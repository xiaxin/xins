package cmd

import (
	"net"
	"xins/core"
)

var (
	protocol = core.NewProtocol()
)

func conn() (net.Conn, error) {
	conn, err := net.Dial("tcp", "localhost:9900")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
