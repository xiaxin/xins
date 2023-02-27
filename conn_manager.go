package xins

import (
	"net"
	"sync"
)

// 链接管理
type ConnManager struct {
	Table map[uint]*Conn

	lock sync.RWMutex

	pk uint
}

func NewConnManager() *ConnManager {
	return &ConnManager{
		Table: make(map[uint]*Conn),
		pk:    0,
	}
}

func (cm *ConnManager) NewConn(server *Server, conn net.Conn) *Conn {
	return NewConn(server, conn)
}

func (cm *ConnManager) AddConn(conn *Conn) {

	cm.lock.Lock()
	cm.pk++
	conn.SetID(cm.pk)

	cm.Table[conn.ID()] = conn
	cm.lock.Unlock()
}

func (cm *ConnManager) DelConn(conn *Conn) {

	cm.lock.Lock()
	delete(cm.Table, conn.ID())
	cm.lock.Unlock()
}
