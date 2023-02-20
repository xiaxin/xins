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
	packer       Packer
}

func NewSession(conn *Conn, packer Packer) *Session {
	return &Session{
		id:           uuid.NewString(),
		conn:         conn,
		closed:       make(chan struct{}),
		requestQueue: make(chan Request, 10),
		packer:       packer,
	}
}

func (s *Session) ID() string {
	return s.id
}

func (s *Session) SetID(id string) {
	s.id = id
}

func (s *Session) read(router *Router) {
	// TODO
	logger.Debug("[conn serve start]")

	tcpConn := s.conn.GetTCPConn()

	for {
		message, err := s.packer.Unpack(tcpConn)

		if nil != err {
			if errors.Is(err, io.EOF) {
				break
			}
			logger.Errorf("[recv error] [read head] %s", err.Error())
			continue
		}

		// // 读取头数据
		// headData := make([]byte, defaultDataHandle.HeadLen())
		// if _, err = io.ReadFull(tcpConn, headData); err != nil {
		// 	if errors.Is(err, io.EOF) {
		// 		break
		// 	}
		// 	logger.Errorf("[recv error] [read head] %s", err.Error())
		// 	continue
		// }

		// // 解析头数据
		// head, err := defaultDataHandle.UnpackHead(headData)

		// if err != nil {
		// 	fmt.Printf("[recv error] [unpack head] %s\n", err.Error())
		// 	continue
		// }

		// if head == nil {
		// 	continue
		// }

		// var (
		// 	bodyData []byte
		// 	body     *DataBody
		// )
		// if head.Len() > 0 {
		// 	bodyData = make([]byte, head.Len())

		// 	if _, err = io.ReadFull(tcpConn, bodyData); err != nil {
		// 		logger.Errorf("[recv error] [read body] %s", err.Error())
		// 		continue
		// 	}

		// 	body, err = defaultDataHandle.UnpackBody(bodyData)

		// 	if err != nil {
		// 		logger.Errorf("[recv error] [unpack body] %s", err.Error())
		// 		continue
		// 	}
		// }

		request := NewRequest(s, message)

		// TODO
		go router.Handle(request)
	}

	s.close()
	logger.Debugf("session %s read exit because of error\n", s.id)
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

		writeBytes, err := s.packer.Pack(request.Message())

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
