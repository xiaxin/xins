package object

import "xins"

var (
	DefaultUserManager  *UserManager
	DefaultGroupManager *GroupManager

	codc = &xins.JsonCodec{}
)

func init() {
	DefaultUserManager = NewUserManager()
	DefaultGroupManager = NewGroupManager()

	DefaultGroupManager.Add(NewGroup("1"))
}
