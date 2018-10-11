package main

import (
	"./action"
	"./listen"
	"github.com/ezcorn/goe"
)

func main() {
	serverName := "demo"
	goe.RegAction(action.HelloWorldAction)
	goe.RegAction(action.HelloWorld403Action)
	goe.RegListen(listen.JumpToHelloWorldListen)
	goe.RelateActionToListen("/helloWorld403", "jumpToHelloWorld")
	goe.InitServer(serverName)
}
