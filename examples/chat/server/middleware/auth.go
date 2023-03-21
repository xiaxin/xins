package middleware

import (
	"xins/core"
)

func AuthMiddleware(next core.RouteFunc) core.RouteFunc {
	return func(request core.Context) {
		// TODO 验证状态
		next(request)
	}
}
