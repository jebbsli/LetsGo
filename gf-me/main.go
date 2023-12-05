package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"os"
)

func main() {
	fileServer := g.Server()

	fileServer.BindHandler("/test", func(r *ghttp.Request) {
		r.Response.Writef("test ....")
	})

	curDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fileServer.SetIndexFolder(true)
	fileServer.SetServerRoot(curDir)
	fileServer.SetPort(8264)
	fileServer.Run()
}
