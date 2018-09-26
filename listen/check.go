package listen

import (
	"github.com/ezcorn/goe"
)

func CheckListen() *goe.Listen {
	return goe.NewListen("check", "Print word world", func(in goe.In) goe.Program {
		return func(in goe.In, out goe.Out) {
			out.Status(403)
		}
	})
}
