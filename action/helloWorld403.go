package action

import (
	"github.com/ezcorn/goe"
	"net/http"
)

func HelloWorld403Action() *goe.Action {
	return goe.NewAction("/helloWorld403", "打印helloWorld403", []string{
		http.MethodPost, http.MethodGet,
	}, func(in goe.In, out goe.Out, libs goe.Libs) {
		out.Echo("helloWorld403")
	})
}
