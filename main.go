package main

import (
	"fmt"

	"dierbei/operator-one/pkg/sysinit"

	"github.com/valyala/fasthttp"
	"github.com/yeqown/fasthttp-reverse-proxy/v2"
)

var jtthink = proxy.NewReverseProxy("www.jtthink.com")

// ProxyHandler 处理器
func ProxyHandler(ctx *fasthttp.RequestCtx) {
	if getProxy := sysinit.GetRoute(ctx.Request); getProxy != nil {
		getProxy.ServeHTTP(ctx)
	} else {
		ctx.Response.SetStatusCode(404)
		ctx.Response.SetBodyString("404...")
	}
}

func main() {
	sysinit.InitConfig()


	fasthttp.ListenAndServe(fmt.Sprintf(":%d", sysinit.SysConfig.Server.Port), ProxyHandler)
}
