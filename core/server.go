package core

import "net"

type Server struct {
	listener net.Listener

	// TODO
	options *Options

	stopped chan struct{}

	logger Logger
}

func NewServer(opts ...Option) *Server {
	options := newOptions(opts...)

	return &Server{
		stopped: make(chan struct{}),
		options: options,
	}
}

// func (s *Server) NewConn(conn net.Conn) *Conn {
// 	return s.connManager.NewConn(s, conn)
// }

// func (s *Server) AddConn(conn *Conn) {
// 	s.connManager.AddConn(conn)
// }

// func (s *Server) DelConn(conn *Conn) {
// 	s.connManager.DelConn(conn)
// }

func (s *Server) serve(addr string) error {
	var err error

	// 	if s.options.protocol == nil {
	// 		return ErrorProtocolIsNil
	// 	}

	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	// 	logger.Debug("[listener accept]")

	for {
		if s.isStopped() {
			s.logger.Debug("server accept loop stopped")
			return ErrorServerStopped
		}

		tcpConn, err := s.listener.Accept()
		if err != nil {
			if s.isStopped() {
				s.logger.Debug("server accept loop stopped")
				return ErrorServerStopped
			}
			s.logger.Debugf("accept error %s", err)
			continue
		}

		go s.handleConn(tcpConn)
	}

}

func (s *Server) handleConn(tcpConn net.Conn) {

	// 	conn := s.NewConn(tcpConn)

	// defer conn.Close()

	session := NewSession(tcpConn, s.options.Protocol())

	// 	s.AddConn(conn)
	// 	defer s.DelConn(conn)

	if s.options.beforeSession != nil {
		s.options.beforeSession(session)
	}

	go session.read()
	go session.write()

	select {
	// TODO
	case <-session.closed:
	case <-s.stopped:
	}

	if s.options.afterSession != nil {
		s.options.afterSession(session)
	}

}

func (s *Server) Run(addr string) error {
	// s.options.protocol.PrintRoutes(addr)

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

func (s *Server) Options() *Options {
	return s.options
}
