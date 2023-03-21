package core

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ Session = &session{}

// TODO
type Session interface {
	DebugLogger
	ErrorLogger

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

	closeOnce sync.Once
	closed    chan struct{}

	ctxPool  sync.Pool // router context pool
	ctxQueue chan Context

	logger *zap.SugaredLogger
}

func NewSessionByUUID(conn net.Conn, protocol Protocol) *session {
	return NewSession(uuid.NewString(), conn, protocol)
}

func NewSession(id any, conn net.Conn, protocol Protocol) *session {
	sess := &session{
		id:       id,
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
func (s *session) read(timeout time.Duration) {

	for {
		select {
		case <-s.closed:
			return
		default:
		}

		if timeout > 0 {
			if err := s.conn.SetDeadline(time.Now().Add(timeout)); err != nil {
				s.Errorf("set read deadline err: %s", err)
				break
			}
		}

		err := s.protocol.Handle(s)

		if nil != err {
			// TODO 需要区分正常 Close 和 Timeout 返回的错误
			s.Errorf("[session %s] read error, %s", s.id, err.Error())
			break
		}
	}

	s.Close()

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
	s.closeOnce.Do(func() { close(s.closed) })
}

func (s *session) Send(ctx Context) bool {
	// 发送逻辑，不能放在 select 中，否则会导致随机选取，导致 closed = true 时，函数返回 true。
	select {
	case <-s.closed:
		return false
	case <-ctx.Done():
		return false
	default:
	}

	s.ctxQueue <- ctx
	return true
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

func (s *session) Error(message string) {
	// TODO
	s.logger.Errorf("[sid] [%s] %s", s.ID(), message)
}

func (s *session) Errorf(template string, args ...interface{}) {
	// TODO
	s.logger.Errorf("[sid] [%s] %s", s.ID(), fmt.Sprintf(template, args...))
}

func (s *session) NewContext() Context {
	c := s.ctxPool.Get().(*context)
	c.reset()
	c.SetSession(s)
	return c
}

func (s *session) Run() chan struct{} {
	go s.read(10 * time.Second)
	go s.write()

	return s.closed
}
