package object

import (
	"xins/core"
)

type User struct {
	id   string
	name string

	session core.Session
}

func NewUser(session core.Session, id string, name string) *User {
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

	bytes, err := codc.Marshal(message)
	if nil != err {
		return err
	}

	return u.sendRequest(1, bytes)
}

// TODO
func (u *User) SendGroupMessage(message *GroupMessage) error {
	bytes, err := codc.Marshal(message)
	if nil != err {
		return err
	}

	return u.sendRequest(2, bytes)
}

func (u *User) sendRequest(id uint32, data []byte) error {

	u.session.Debugf("%d %s", id, data)
	message := core.NewMessage(id, data)

	bytes, err := u.session.Protocol().Pack(message)

	if nil != err {
		return err
	}

	ctx := u.session.NewContext()
	ctx.SetMessage(core.NewMessage(id, bytes))

	u.session.Send(ctx)

	return nil
}
