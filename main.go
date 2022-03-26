package main

import (
	"dierbei/operator-one/pkg/sysinit"
	"fmt"
	"github.com/valyala/fasthttp"
	"github.com/yeqown/log"
)

//func ProxyHandler(ctx *fasthttp.RequestCtx){
//	//代表匹配到了 path
//	 if getProxy:=sysinit.GetRoute(ctx.Request);getProxy!=nil{
//		 filters.ProxyFilters(getProxy.RequestFilters).Do(ctx) //过滤
//		 getProxy.Proxy.ServeHTTP(ctx) //反代
//		 filters.ProxyFilters(getProxy.ResponseFilters).Do(ctx) //过滤
//	 }else{
//		 ctx.Response.SetStatusCode(404)
//		 ctx.Response.SetBodyString("404...")
//	 }
//
//	// jtthink.ServeHTTP(ctx)
//
//}
//var jtthink=proxy.NewReverseProxy("www.jtthink.com",)
func main() {
	sysinit.InitConfig()
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%d", sysinit.SysConfig.Server.Port),ProxyHandler))
}