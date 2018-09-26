package action

import (
	"github.com/ezcorn/goe"
	"net/http"
)

func WorldAction() *goe.Action {
	return goe.NewAction("/world", "Print word world", []string{
		http.MethodPost,
	}, func(in goe.In, out goe.Out) {
		out.Echo("world")
	})
}
