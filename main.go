package main

import (
	"github.com/valyala/fasthttp"
	"github.com/yeqown/fasthttp-reverse-proxy/v2"
)

var jtthink=proxy.NewReverseProxy("www.jtthink.com")

// ProxyHandler 处理器
func ProxyHandler(ctx *fasthttp.RequestCtx){
	jtthink.ServeHTTP(ctx)

	// 在处理之后进行请求头的修改
	ctx.Response.Header.Add("myname","xiaolatiao")
}

func main() {
	fasthttp.ListenAndServe(":80",ProxyHandler)
}