package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {
	s := g.Server()

	// 注册接口
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello")
	})

	// 域名绑定
	s.Domain("localhost").BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello2!")
	})

	// 静态服务
	s.SetIndexFolder(true)
	s.SetServerRoot("/Users/jebbsli/jebbsli/me/projects/LetsGo")

	// 设置多个监听端口
	s.SetPort(8080, 8081, 8082)

	//gfile.PutBytes()

	s.Run()
}
