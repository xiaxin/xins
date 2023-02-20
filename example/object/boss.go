package object

type Boss struct {
	name string
	// 生命值
	hp int
	// 谁杀死的
	who uint
	// 是否被杀死
	dead bool
}

func NewBoss(name string, hp int) *Boss {
	return &Boss{
		name: name,
		hp:   hp,
		dead: false,
	}
}

func (b *Boss) SetWho(who uint) {
	b.who = who
}

func (b *Boss) GetWho() uint {
	return b.who
}

func (b *Boss) SetHP(hp int) {
	b.hp = hp
}

func (b *Boss) SubHP(hp int) {
	b.hp = b.hp - hp
}

func (b *Boss) GetHP() int {
	return b.hp
}

func (b *Boss) Dead() bool {
	return b.dead
}

func (b *Boss) Kill(who uint) {
	if b.dead {
		return
	}
	b.who = who
	b.dead = true
}
