package sysinit

import (
	"net/http"

	"github.com/gorilla/mux"
)

//MyRouter route构建器，build 方法必须要执行
var MyRouter *mux.Router

func init() {
	MyRouter = mux.NewRouter()
}

type RouteBuilder struct {
	route *mux.Route
}

func NewRouteBuilder() *RouteBuilder {
	return &RouteBuilder{route: MyRouter.NewRoute()}
}

func (r *RouteBuilder) SetPath(path string, exact bool) *RouteBuilder {
	if exact {
		r.route.Path(path)
	} else {
		r.route.PathPrefix(path)
	}
	return r
}

// SetHost 第二个参数是故意的，方便调用时 传入 条件，省的外面写 if else
func (r *RouteBuilder) SetHost(host string, set bool) *RouteBuilder {
	if set {
		r.route.Host(host)
	}
	return r
}

// Build 构建
func (r *RouteBuilder) Build(handler http.Handler) {
	r.route.
		Methods("GET", "POST", "PUT", "DELETE", "OPTIONS").
		Handler(handler)
}
