package object

import (
	"errors"
)

type GroupManager struct {
	table map[string]*Group
}

func NewGroupManager() *GroupManager {
	return &GroupManager{
		table: make(map[string]*Group),
	}
}

func (gm *GroupManager) Add(g *Group) {
	gm.table[g.ID()] = g
}

func (gm *GroupManager) Del(g *Group) {
	delete(gm.table, g.ID())
}
func (gm *GroupManager) Get(id string) (*Group, error) {
	if group, ok := gm.table[id]; ok {
		return group, nil
	}
	return nil, errors.New("group is not exist")
}

func (gm *GroupManager) SendMessage(message *GroupMessage) error {
	gid := message.GID

	group, ok := gm.table[gid]

	if !ok {
		return errors.New("group is not exist")
	}

	return group.SendMessage(message)
}
