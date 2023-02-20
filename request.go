package xins

// 请求
type Request struct {
	session *Session

	message *Message
}

func NewRequest(session *Session, message *Message) Request {
	return Request{session, message}
}

func (r *Request) Session() *Session {
	return r.session
}

func (r *Request) Message() *Message {
	return r.message
}

func (r *Request) ID() uint32 {
	return r.message.ID()
}

func (r *Request) SessionID() string {
	return r.session.ID()
}

func (r *Request) ConnID() uint {
	return r.session.conn.ID()
}
