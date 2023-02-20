package router

import "xins"

type Panic struct {
}

func (p *Panic) Handle(request xins.Request) {
	panic("error message")
}
