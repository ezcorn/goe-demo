package action

import (
	"github.com/ezcorn/goe"
	"net/http"
)

func HelloAction() *goe.Action {
	route := "/hello"
	comment := "Print word hello"
	return goe.NewAction(route, comment, []string{
		http.MethodPost,
		http.MethodGet,
	}, func(in goe.In, out goe.Out) {
		out.Echo("hello")
	})
}
