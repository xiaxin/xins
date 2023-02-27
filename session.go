package xins

import (
	"errors"
	"io"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	id           string
	conn         *Conn
	closed       chan struct{}
	requestQueue chan Request

	protocol Protocol

	timeout time.Duration
}

func NewSession(conn *Conn, protocol Protocol) *Session {
	return &Session{
		id:           uuid.NewString(),
		conn:         conn,
		closed:       make(chan struct{}),
		requestQueue: make(chan Request, 10),

		protocol: protocol,

		timeout: 0,
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
	logger.Debugf("[session %s] read start", s.id)

	for {
		err := s.protocol.Handle(s)

		if nil != err {
			if errors.Is(err, io.EOF) {
				break
			}
			logger.Errorf("[session %s] read error, %s", s.id, err.Error())
			continue
		}
	}

	s.close()
	logger.Debugf("[session %s] read exit because of error", s.id)
}

func (s *Session) write() {

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

		if _, err = s.WriteBytes(writeBytes); err != nil {
			logger.Errorf("session %s conn write err: %s", s.id, err)
			break
		}
	}
	s.close()
	logger.Debugf("session %s writeOutbound exit because of error", s.id)
}

// TODO
func (s *Session) SendRequest(request Request) (ok bool) {
	select {
	case <-s.closed:
		return false
	case s.requestQueue <- request:
		return true
	}
}

func (s *Session) WriteBytes(data []byte) (n int, err error) {
	return s.conn.GetTCPConn().Write(data)
}

func (s *Session) close() {
	close(s.closed)
}
