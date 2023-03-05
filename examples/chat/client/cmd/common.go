package cmd

import (
	"net"
	xinsProtocol "xins/protocol/xins"
)

var (
	protocol = xinsProtocol.NewDefaultProtocol()
)

func conn() (net.Conn, error) {
	conn, err := net.Dial("tcp", "localhost:9900")
	if err != nil {
		return nil, err
	}

	return conn, nil
}
