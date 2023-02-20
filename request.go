package xins

// 请求
type Request struct {
	session *Session

	message interface{}
}

func NewRequest(session *Session, message interface{}) Request {
	return Request{session, message}
}

func (r *Request) Session() *Session {
	return r.session
}

func (r *Request) Message() interface{} {
	return r.message
}

func (r *Request) SessionID() string {
	return r.session.ID()
}

func (r *Request) ConnID() uint {
	return r.session.conn.ID()
}
