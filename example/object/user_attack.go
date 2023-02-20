package object

type UserAttack struct {
	BossID int `json:"boss_id"`
}

func NewUserAttack(id int) *UserAttack {
	return &UserAttack{
		BossID: id,
	}
}
