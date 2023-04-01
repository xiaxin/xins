package middleware

import (
	"xins/core"
)

func AuthMiddleware() core.RouteFunc {
	return func(ctx core.Context) {
		sess := ctx.Session()
		sess.Debug("auth start")
		ctx.Next()
		sess.Debug("auth end")
	}
}
