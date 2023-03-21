package object

import "xins/core"

var (
	DefaultUserManager  *UserManager
	DefaultGroupManager *GroupManager

	codc = &core.JsonCodec{}

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
