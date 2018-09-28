package action

import (
	"github.com/ezcorn/goe"
	"net/http"
)

func HelloWorldAction() *goe.Action {
	return goe.NewAction("/helloWorld", "打印helloWorld", []string{
		http.MethodPost, http.MethodGet,
	}, func(in goe.In, out goe.Out) {
		out.Echo("helloWorld")
	})
}
