package protocol

import (
	"fmt"
	"xins"
)

// 路由
type Router struct {
	table map[uint32]Route
}

// 路由接口
type Route interface {
	Handle(request xins.Request)
}

func NewRouter() *Router {
	return &Router{
		table: make(map[uint32]Route),
	}
}

func (rg *Router) Add(id uint32, route Route) {
	rg.table[id] = route
}

func (rg *Router) Del(id uint32) {
	delete(rg.table, id)
}

func (rg *Router) Get(id uint32) (Route, error) {
	if router, ok := rg.table[id]; ok {
		return router, nil
	}
	// TODO
	return nil, fmt.Errorf("[%d] is not exists", id)
}
