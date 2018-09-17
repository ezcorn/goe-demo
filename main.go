package main

import "github.com/ezcorn/goe"
import "./action"

func main() {
	goe.ServerRoute["/"] = action.HelloAction
	goe.InitServer(9339, "https://github.com/ezcorn/goe-vendor.git")
}
