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

// func (rg *Router) Handle(request Request) {

// 	// todo
// 	defer func() {
// 		if r := recover(); r != nil {
// 			logger.Errorf("%s", fmt.Sprintf("PANIC | %s | %+v \n%s", r, request.SessionID(), debug.Stack()))
// 		}
// 	}()

// 	// todo
// 	message := request.Message()
// 	id := message.(*Message).ID()

// 	router, ok := rg.table[id]

// 	if !ok {
// 		fmt.Printf("[error] [router handle] id:%d is not found\n", id)
// 		return
// 	}
// 	router.Handle(request)
// }
