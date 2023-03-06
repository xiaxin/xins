package xins

import "sync"

// 请求
type Request struct {
	session *Session
	message interface{}

	m      sync.RWMutex
	keyval map[string]interface{}
}

func NewRequest(session *Session, message interface{}) *Request {
	return &Request{
		session: session,
		message: message,
		keyval:  make(map[string]interface{}),
	}
}

func (r *Request) Message() interface{} {
	return r.message
}

func (r *Request) Session() *Session {
	return r.session
}

func (r *Request) Get(key string) (value interface{}, exists bool) {
	r.m.RLock()
	value, exists = r.keyval[key]
	r.m.RUnlock()
	return
}

func (r *Request) Set(key string, value interface{}) {
	r.m.Lock()
	defer r.m.Unlock()

	if r.keyval == nil {
		r.keyval = make(map[string]interface{})
	}
	r.keyval[key] = value

}

func (r *Request) Remove(key string) {
	r.m.Lock()
	defer r.m.Unlock()

	delete(r.keyval, key)
}
