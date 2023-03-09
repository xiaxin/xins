package protocol

type Route struct {
	ID     uint32
	Handle RouteFunc
}
