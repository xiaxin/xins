package core

import (
	"net"
)

// TODO
type Session interface {
	ID() interface{}
	Conn() net.Conn

	read()
	write()
}

type session struct {
	conn net.Conn
}

func NewSession(conn net.Conn, protocol Protocol, onstart func(session Session), onstop func(session Session)) Session {
	return &session{
		conn: conn,
	}
}

func (s *session) ID() any {
	return nil
}

func (s *session) Conn() net.Conn {
	return s.conn
}

func (s *session) read() {
}
func (s *session) write() {

}
