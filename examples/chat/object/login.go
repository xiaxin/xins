package object

type Login struct {
	Token string `json:"token"` // Token
}

func NewLogin(token string) *Login {
	return &Login{token}
}
