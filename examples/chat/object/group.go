package object

import (
	"errors"
)

type Group struct {
	id string

	table map[string]*User
}

func NewGroup(id string) *Group {
	return &Group{
		id:    id,
		table: make(map[string]*User),
	}
}

func (g *Group) ID() string {
	return g.id
}

func (g *Group) AddUser(user *User) error {
	if _, ok := g.table[user.ID()]; ok {
		return errors.New("user is exists")
	}

	g.table[user.ID()] = user

	return nil
}

func (g *Group) DelUser(user *User) error {
	if _, ok := g.table[user.ID()]; !ok {
		return errors.New("user is not exists")
	}

	delete(g.table, user.ID())
	return nil
}

func (g *Group) SendMessage(message *GroupMessage) error {

	for _, user := range g.table {
		user.SendGroupMessage(message)
	}

	return nil
}
