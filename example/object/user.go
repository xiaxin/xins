package object

type User struct {
	id   uint
	name string
	agg  int
}

func NewUser(id uint, name string) *User {
	return &User{
		id: id, name: name,
	}
}

func (u *User) SetAgg(agg int) {
	u.agg = agg
}

func (u *User) GetAgg() int {
	return u.agg
}
