package object

import (
	"xins"

	protocol "xins/protocol/default"
)

type User struct {
	id   string
	name string

	session *xins.Session
}

func NewUser(session *xins.Session, id string, name string) *User {
	return &User{id, name, session}
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

// TODO
func (u *User) SendUserMessage(message *UserMessage) error {

	bytes, err := codc.Encode(message)
	if nil != err {
		return err
	}

	return u.sendRequest(1, bytes)
}

// TODO
func (u *User) SendGroupMessage(message *GroupMessage) error {
	bytes, err := codc.Encode(message)
	if nil != err {
		return err
	}

	return u.sendRequest(2, bytes)
}

func (u *User) sendRequest(id uint32, data []byte) error {

	u.session.Debugf("%s %s", id, data)
	message := protocol.NewMessage(id, data)

	bytes, err := u.session.Protocol().Pack(message)

	if nil != err {
		return err
	}

	u.session.SendRequest(xins.NewRequest(u.session, protocol.NewMessage(id, bytes)))

	return nil
}
