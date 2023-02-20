package xins

func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		id:   id,
		data: data,
	}
}

type Message struct {
	id   uint32
	data []byte
}

func (m *Message) ID() uint32 {
	return m.id
}

func (m *Message) Data() []byte {
	return m.data
}
