package object

import "xins"

var (
	DefaultUserManager  *UserManager
	DefaultGroupManager *GroupManager

	codc = &xins.JsonCodec{}

	tokens = map[string]string{
		"token-a": "user-a",
	}
)

func init() {
	DefaultUserManager = NewUserManager()
	DefaultGroupManager = NewGroupManager()

	DefaultGroupManager.Add(NewGroup("1"))
}

func GetUIDByToken(token string) (string, bool) {
	val, ok := tokens[token]
	return val, ok
}
