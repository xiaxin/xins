package xins

import (
	"errors"
	"io"

	"github.com/google/uuid"
)

type Session struct {
	id           string
	conn         *Conn
	closed       chan struct{}
	requestQueue chan Request
	// packer       Packer
	protocol Protocol
}

func NewSession(conn *Conn, protocol Protocol) *Session {
	return &Session{
		id:           uuid.NewString(),
		conn:         conn,
		closed:       make(chan struct{}),
		requestQueue: make(chan Request, 10),
		// packer:       packer,
		protocol: protocol,
	}
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) SetID(id string) {
	s.id = id
}

func (s *Session) Conn() *Conn {
	return s.conn
}

func (s *Session) read() {
	// TODO
	logger.Debug("[conn serve start]")

	// TODO
	for {

		err := s.protocol.Handle(s)

		if nil != err {
			if errors.Is(err, io.EOF) {
				break
			}
			logger.Errorf("[recv error] [read head] %s", err.Error())
			continue
		}
	}

	s.close()
	logger.Debugf("session %s read exit because of error", s.id)
}

func (s *Session) write() {

	tcpConn := s.conn.GetTCPConn()

	for {
		var request Request
		select {
		case <-s.closed:
			return
		case request = <-s.requestQueue:
		}

		writeBytes, err := s.protocol.Pack(request.Message())

		if nil != err {
			logger.Errorf("session %s pack outbound message err: %s", s.id, err)
			continue
		}
		if writeBytes == nil {
			continue
		}

		if _, err = tcpConn.Write(writeBytes); err != nil {
			logger.Errorf("session %s conn write err: %s", s.id, err)
			break
		}
	}
	s.close()
	logger.Debugf("session %s writeOutbound exit because of error", s.id)
}

func (s *Session) Send(request Request) (ok bool) {
	select {
	case <-s.closed:
		return false
	case s.requestQueue <- request:
		return true
	}
}

func (s *Session) close() {
	close(s.closed)
}
