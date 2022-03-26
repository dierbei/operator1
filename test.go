package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
)

func maintest() {
	r:=mux.NewRouter()

	// 添加两个路由
	r.NewRoute().Path("/").Methods("GET")
	r.NewRoute().Path("/users/{id:\\d+}").Methods("GET","POST","PUT","DELETE","OPTIONS")

	// 创建匹配对象
	match:=&mux.RouteMatch{}

	// 进行路由匹配
	req:=&http.Request{URL:&url.URL{Path: "/users/abc"},Method:"GET"}
	fmt.Println(r.Match(req,match))
}
