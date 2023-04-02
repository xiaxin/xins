package middleware

import (
	"xins/core"
)

func BMiddleware() core.RouteFunc {
	return func(ctx core.Context) {
		sess := ctx.Session()
		sess.Debug("b start")
		ctx.Next()
		sess.Debug("b end")
	}
}
