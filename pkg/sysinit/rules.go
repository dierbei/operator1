package sysinit

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/valyala/fasthttp"
	"github.com/yeqown/fasthttp-reverse-proxy/v2"
)

var MyRouter *mux.Router

type ProxyHandler struct {
	Proxy *proxy.ReverseProxy // proxy对象
}

//空函数没啥用
func (h *ProxyHandler) ServeHTTP(http.ResponseWriter, *http.Request) {}

func init() {
	MyRouter = mux.NewRouter()
}

// ParseRule 解析配置文件中的rules
func ParseRule() {
	for _, rule := range SysConfig.Ingress.Rules {
		for _, path := range rule.HTTP.Paths {

			//构建反代对象
			rProxy := proxy.NewReverseProxy(fmt.Sprintf("%s:%d",
				path.Backend.Service.Name,        // 服务名
				path.Backend.Service.Port.Number, // 端口
			))

			// path绑定反代处理器
			MyRouter.NewRoute().Path(path.Path).Methods(
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"OPTIONS",
			).Handler(&ProxyHandler{Proxy: rProxy})
		}
	}
}

// GetRoute 获取路由（先匹配 请求path ，如果匹配到 ，会返回 对应的proxy 对象)
func GetRoute(req fasthttp.Request) *proxy.ReverseProxy {
	match := &mux.RouteMatch{}

	httpReq := &http.Request{
		URL:    &url.URL{Path: string(req.URI().Path())}, // 请求路径path
		Method: string(req.Header.Method()),              // 请求方法
	}

	// 匹配到之后返回反代处理器
	if MyRouter.Match(httpReq, match) {
		return match.Handler.(*ProxyHandler).Proxy
	}
	return nil
}
