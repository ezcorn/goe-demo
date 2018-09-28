package listen

import (
	"github.com/ezcorn/goe"
)

func JumpToHelloWorldListen() *goe.Listen {
	return goe.NewListen("jumpToHelloWorld", "跳转到HelloWorld", func(in goe.In) goe.Program {
		return func(in goe.In, out goe.Out) {
			out.Echo("<script>location.href='/helloWorld'</script>")
		}
	})
}
