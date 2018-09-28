package main

import (
	"./action"
	"./listen"
	"github.com/ezcorn/goe"
)

func main() {
	goe.RegAction(action.HelloWorldAction, func(a *goe.Action) {
		a.Relate("check")
	})
	goe.RegAction(action.HelloWorld403Action, nil)
	goe.RegListen(listen.JumpToHelloWorldListen, func(l *goe.Listen) {
		l.Relate("/helloWorld403")
	})

	goe.InitServer(9339)
}
