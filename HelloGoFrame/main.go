package main

import (
	_ "HelloGoFrame/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	"HelloGoFrame/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
