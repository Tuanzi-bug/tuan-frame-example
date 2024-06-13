package main

import (
	"fmt"
	tuanframe "github.com/Tuanzi-bug/tuan-frame"
)

func main() {
	engine := tuanframe.New()
	g := engine.Group("")
	g.Get("/hello", func(ctx *tuanframe.Context) {
		fmt.Fprintln(ctx.W, "hello tuan_frame")
	})
	engine.Run()
}
