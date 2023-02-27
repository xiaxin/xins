package object

import "errors"

type UserManager struct {
	table map[string]*User
}

func NewUserManager() *UserManager {
	return &UserManager{
		table: make(map[string]*User),
	}
}

func (um *UserManager) Add(u *User) {
	um.table[u.ID()] = u
}

func (um *UserManager) Get(id string) (*User, error) {
	if user, ok := um.table[id]; ok {
		return user, nil
	}
	return nil, errors.New("user is not exist")
}

func (um *UserManager) Del(u *User) {
	delete(um.table, u.ID())
}

func (um *UserManager) SendMessage(id string, message string) error {
	user, ok := um.table[id]

	if !ok {
		return errors.New("user is not exist")
	}

	return user.SendUserMessage(NewUserMessage(id, message))
}
