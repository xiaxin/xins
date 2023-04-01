package core

type Route struct {
	ID      uint32
	Handles []RouteFunc
}

func NewRoute(id uint32, fns []RouteFunc) *Route {
	return &Route{
		ID:      id,
		Handles: fns,
	}
}
