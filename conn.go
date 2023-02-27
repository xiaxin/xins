package xins

import (
	"net"
)

// 链接
type Conn struct {
	id uint

	server *Server
	conn   net.Conn
}

func NewConn(server *Server, conn net.Conn) *Conn {
	return &Conn{
		server: server,
		conn:   conn,
	}
}

func (c *Conn) ID() uint {
	return c.id
}

func (c *Conn) SetID(id uint) {
	c.id = id
}

func (c *Conn) GetTCPConn() net.Conn {
	return c.conn
}

func (c *Conn) Close() error {
	return c.conn.Close()
}
