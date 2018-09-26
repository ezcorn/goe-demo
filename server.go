package main

import (
	"./action"
	"./listen"
	"github.com/ezcorn/goe"
)

func main() {
	goe.RegAction(action.HelloAction, func(a *goe.Action) {
		a.Relate("check")
	})
	goe.RegAction(action.WorldAction, nil)
	goe.RegListen(listen.CheckListen, func(l *goe.Listen) {
		l.Relate("/hello")
	})

	goe.InitServer(9339)
}
