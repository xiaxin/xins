package protocol

import (
	"fmt"
	"xins"
)

// 路由
type Router struct {
	// TODO
	routes           map[uint32]RouteFunc
	routeMiddlewares map[uint32][]MiddlewareFunc
	middlewares      []MiddlewareFunc
}

type RouteFunc func(request *xins.Request)
type MiddlewareFunc func(next RouteFunc) RouteFunc

func NewRouter() *Router {
	return &Router{
		routes:           make(map[uint32]RouteFunc),
		routeMiddlewares: make(map[uint32][]MiddlewareFunc),
	}
}

func (r *Router) Add(id uint32, route RouteFunc, fns ...MiddlewareFunc) {

	r.routes[id] = route

	ms := make([]MiddlewareFunc, 0, len(fns))

	for _, fn := range fns {
		if fn != nil {
			ms = append(ms, fn)
		}
	}
	if len(ms) != 0 {
		r.routeMiddlewares[id] = ms
	}
}

func (r *Router) AddMiddleware(middlewares ...MiddlewareFunc) {
	for _, m := range middlewares {
		if m != nil {
			r.middlewares = append(r.middlewares, m)
		}
	}
}

func (rg *Router) Del(id uint32) {
	delete(rg.routes, id)
}

func (rg *Router) Get(id uint32) (RouteFunc, error) {
	if route, ok := rg.routes[id]; ok {
		return route, nil
	}
	// TODO
	return nil, fmt.Errorf("[%d] is not exists", id)
}

func (r *Router) HandleRequest(request *xins.Request) {

}
