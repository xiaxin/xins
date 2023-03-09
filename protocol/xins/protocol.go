package protocol

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"xins"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cast"
)

type protocol struct {
	packer Packer
	codec  xins.Codec
	router *Router
}

func NewProtocol() *protocol {
	return &protocol{
		packer: NewDefaultPacker(),
		codec:  &xins.JsonCodec{},
		router: NewRouter(),
	}
}

func (dp *protocol) NewMessage(id uint32, data interface{}) (*Message, error) {

	bytes, err := dp.codec.Encode(data)

	if nil != err {
		return nil, err
	}

	return NewMessage(id, bytes), nil
}

func (dp *protocol) Encode(v interface{}) ([]byte, error) {
	return dp.codec.Encode(v)
}

func (dp *protocol) Decode(data []byte, v interface{}) error {
	return dp.codec.Decode(data, v)
}

func (dp *protocol) Pack(message interface{}) ([]byte, error) {
	return dp.packer.Pack(message.(*Message))
}

func (dp *protocol) Unpack(reader io.Reader) (interface{}, error) {
	return dp.packer.Unpack(reader)
}

func (dp *protocol) AddRoute(id uint32, route RouteFunc, middlewares ...MiddlewareFunc) {
	dp.router.Add(id, route, middlewares...)
}

func (dp *protocol) AddMiddleware(middlewares ...MiddlewareFunc) {
	dp.router.AddMiddleware(middlewares...)
}

func (dp *protocol) Handle(session *xins.Session) error {

	message, err := dp.Unpack(session.Conn().GetTCPConn())

	if nil != err {
		return err
	}

	request := xins.NewRequest(session, message)

	go dp.router.HandleRequest(request)

	return nil
}

func (dp *protocol) PrintRoutes(addr string) {

	var w io.Writer = os.Stdout

	_, _ = fmt.Fprintf(w, "\n[XINS] Message-Route Table:\n")
	table := tablewriter.NewWriter(w)
	table.SetHeader([]string{"Message ID", "Route Handler", "Middleware"})
	table.SetAutoFormatHeaders(false) // don't uppercase the header
	table.SetAutoWrapText(false)      // respect the "\n" of cell content
	table.SetRowLine(true)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_LEFT})

	router := dp.router
	// sort ids
	ids := make([]uint32, 0, len(router.routes))
	for id := range router.routes {
		ids = append(ids, id)
	}

	sort.Slice(ids, func(i, j int) bool {
		a, b := cast.ToString(ids[i]), cast.ToString(ids[j])
		return a < b
	})

	// add table row
	for _, id := range ids {
		// route handler
		h := router.routes[id]
		handlerName := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()

		middlewareNames := make([]string, 0, len(router.middlewares)+len(router.routeMiddlewares[id]))
		// global middleware
		for _, m := range router.middlewares {
			middlewareName := fmt.Sprintf("%s(g)", runtime.FuncForPC(reflect.ValueOf(m).Pointer()).Name())
			middlewareNames = append(middlewareNames, middlewareName)
		}

		// route middleware
		for _, m := range router.routeMiddlewares[id] {
			middlewareName := runtime.FuncForPC(reflect.ValueOf(m).Pointer()).Name()
			middlewareNames = append(middlewareNames, middlewareName)
		}

		table.Append([]string{fmt.Sprintf("%v", id), handlerName, strings.Join(middlewareNames, "\n")})
	}

	table.Render()
	// todo
	_, _ = fmt.Fprintf(w, "[XINS] Serving at: %s\n\n", addr)

}
