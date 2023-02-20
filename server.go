package xins

import (
	"net"
)

const ()

// 服务器
type Server struct {
	connManager *ConnManager
	router      *Router
	listener    net.Listener
	stopped     chan struct{}

	packer Packer
}

func NewServer() *Server {
	return &Server{
		connManager: NewConnManager(),
		router:      NewRouter(),

		stopped: make(chan struct{}),
		packer:  NewDefaultPacker(),
	}
}

func (s *Server) AddRoute(id uint32, route Route) {
	s.router.Add(id, route)
}

func (s *Server) NewConn(conn net.Conn) *Conn {
	return s.connManager.NewConn(s, conn)
}

func (s *Server) AddConn(conn *Conn) {
	s.connManager.AddConn(conn)
}

func (s *Server) DelConn(conn *Conn) {
	s.connManager.DelConn(conn)
}

func (s *Server) serve(addr string) error {
	var err error

	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	logger.Debug("[listener accept]")

	for {
		if s.isStopped() {
			logger.Debug("server accept loop stopped")
			return ErrorServerStopped
		}

		tcpConn, err := s.listener.Accept()
		if err != nil {
			if s.isStopped() {
				logger.Debug("server accept loop stopped")
				return ErrorServerStopped
			}
			logger.Debugf("accept error %s", err)
			continue
		}

		go s.handleConn(tcpConn)
	}

}

func (s *Server) handleConn(tcpConn net.Conn) {
	defer tcpConn.Close()

	conn := s.NewConn(tcpConn)

	session := NewSession(conn, s.packer)

	s.AddConn(conn)
	defer s.DelConn(conn)

	go session.read(s.router)
	go session.write()

	select {
	case <-session.closed: // wait for session finished.
	case <-s.stopped: // or the server is stopped.
	}
}

func (s *Server) Run(addr string) error {
	return s.serve(addr)
}

func (s *Server) Stop() error {
	close(s.stopped)
	return s.listener.Close()
}

func (s *Server) isStopped() bool {
	select {
	case <-s.stopped:
		return true
	default:
		return false
	}
}
