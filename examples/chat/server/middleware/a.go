package middleware

import (
	"xins/core"
)

func AMiddleware() core.RouteFunc {
	return func(ctx core.Context) {
		sess := ctx.Session()
		sess.Debug("a start")
		ctx.Abort()
		ctx.Next()
		sess.Debug("a end")
	}
}
