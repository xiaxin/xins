package core

import (
	"fmt"
	"runtime/debug"
)

// 路由
type Router struct {
	// TODO
	routes      map[uint32]*Route
	middlewares []RouteFunc
}

type RouteFunc func(ctx Context)

// type MiddlewareFunc func(next RouteFunc) RouteFunc

func NewRouter() *Router {
	return &Router{
		routes: make(map[uint32]*Route),
	}
}

// Add 添加路由 TODO
func (r *Router) Add(id uint32, route RouteFunc, middlewares ...RouteFunc) {

	size := len(r.middlewares) + len(middlewares) + 1

	fns := make([]RouteFunc, size)

	copy(fns, r.middlewares)
	copy(fns[len(r.middlewares):], middlewares)
	fns[size-1] = route

	r.routes[id] = NewRoute(id, fns)
}

func (r *Router) AddMiddleware(middlewares ...RouteFunc) {
	for _, m := range middlewares {
		if m != nil {
			r.middlewares = append(r.middlewares, m)
		}
	}
}

func (rg *Router) Del(id uint32) {
	delete(rg.routes, id)
}

func (rg *Router) Get(id uint32) (*Route, error) {
	if route, ok := rg.routes[id]; ok {
		return route, nil
	}
	// TODO
	return nil, fmt.Errorf("[%d] is not exists", id)
}

// TODO
func (r *Router) HandleRequest(ctx Context) {

	sess := ctx.Session()

	// todo
	defer func() {
		if r := recover(); r != nil {
			// TODO
			fmt.Printf("%s", fmt.Sprintf("PANIC | %s | %+v \n%s", r, "", debug.Stack()))
		}
	}()

	message := ctx.Message()
	id := message.(*Message).ID()

	route, err := r.Get(id)

	if nil != err {
		sess.Errorf("[error] [router handle] error:%s", err)
		return
	}

	ctx.SetHandles(route.Handles)

	ctx.Next()
}
