package middleware

import (
	"xins"
	protocol "xins/protocol/xins"
)

func AuthMiddleware(next protocol.RouteFunc) protocol.RouteFunc {
	return func(request *xins.Request) {
		// TODO 验证状态
		next(request)
	}
}
