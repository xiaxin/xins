package core

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ Session = &session{}

// TODO
type Session interface {
	DebugLogger
	ID() any
	Conn() net.Conn
	Protocol() Protocol

	SetID(id any) // SetID 设置 Session ID
	Close()       // Close 关闭 Session

	NewContext() Context

	Send(ctx Context) bool
}

type session struct {
	id       any
	conn     net.Conn
	protocol Protocol

	closed chan struct{}

	ctxPool  sync.Pool // router context pool
	ctxQueue chan Context

	logger *zap.SugaredLogger
}

func NewSession(conn net.Conn, protocol Protocol) *session {
	sess := &session{
		id:       uuid.NewString(),
		conn:     conn,
		protocol: protocol,

		closed: make(chan struct{}),
		// TODO
		ctxQueue: make(chan Context, 10),
		ctxPool:  sync.Pool{New: func() interface{} { return NewContext() }},
	}

	sess.logger = logger.WithOptions(zap.AddCallerSkip(1))

	return sess
}

func (s *session) ID() any {
	return s.id
}

func (s *session) SetID(id any) {
	s.id = id
}

func (s *session) Conn() net.Conn {
	return s.conn
}
func (s *session) Protocol() Protocol {
	return s.protocol
}

// TODO
func (s *session) read() {
	s.Debug("read start")

	for {
		select {
		case <-s.closed:
			return
		default:
		}
		err := s.protocol.Handle(s)

		if nil != err {
			if errors.Is(err, io.EOF) {
				break
			}
			logger.Errorf("[session %s] read error, %s", s.id, err.Error())
			continue
		}
	}

	s.Close()
	s.Debug("read exit because of error")
}

// TODO
func (s *session) write() {

	for {
		var context Context
		select {
		case <-s.closed:
			return
		case context = <-s.ctxQueue:
		}

		writeBytes, err := s.protocol.Pack(context.Message())

		if nil != err {
			logger.Errorf("session %s pack outbound message err: %s", s.id, err)
			continue
		}
		if writeBytes == nil {
			continue
		}

		if _, err = s.writeBytes(writeBytes); err != nil {
			logger.Errorf("session %s conn write err: %s", s.id, err)
			break
		}
	}
	s.Close()
	logger.Debugf("session %s writeOutbound exit because of error", s.id)
}

func (s *session) Close() {
	close(s.closed)
}

func (s *session) Send(ctx Context) bool {
	select {
	case <-ctx.Done():
		return false
	case <-s.closed:
		return false
	case s.ctxQueue <- ctx:
		return true
	}
}

func (s *session) writeBytes(data []byte) (n int, err error) {
	return s.conn.Write(data)
}

func (s *session) Debug(message string) {
	// TODO
	s.logger.Debugf("[sid] [%s] %s", s.ID(), message)
}

func (s *session) Debugf(template string, args ...interface{}) {
	// TOD
	s.logger.Debugf("[sid] [%s] %s", s.ID(), fmt.Sprintf(template, args...))
}

func (s *session) NewContext() Context {
	c := s.ctxPool.Get().(*context)
	c.reset()
	c.SetSession(s)
	return c
}
