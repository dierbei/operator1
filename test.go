package main

import (
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/source"

	"log"
	"os"

	"dierbei/operator-one/pkg/filters"
	"dierbei/operator-one/pkg/k8sconfig"
	"dierbei/operator-one/pkg/sysinit"

	"github.com/valyala/fasthttp"
	v1 "k8s.io/api/networking/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
)

func ProxyHandler(ctx *fasthttp.RequestCtx) {
	//代表匹配到了 path
	if getProxy := sysinit.GetRoute(ctx.Request); getProxy != nil {
		filters.ProxyFilters(getProxy.RequestFilters).Do(ctx)  //过滤
		getProxy.Proxy.ServeHTTP(ctx)                          //反代
		filters.ProxyFilters(getProxy.ResponseFilters).Do(ctx) //过滤
	} else {
		ctx.Response.SetStatusCode(404)
		ctx.Response.SetBodyString("404...")
	}
}

func main() {
	logf.SetLogger(zap.New())
	var mylog = logf.Log.WithName("xltproxy")
	mgr, err := manager.New(k8sconfig.K8sRestConfig(), manager.Options{})
	if err != nil {
		mylog.Error(err, "unable to set up manager")
		os.Exit(1)
	}

	// crd
	err = k8sconfig.SchemeBuilder.AddToScheme(mgr.GetScheme())
	if err != nil {
		mylog.Error(err, "unable add schema")
		os.Exit(1)
	}
	//err=builder.ControllerManagedBy(mgr).
	//	For(&k8sconfig.Route{}).Complete(k8sconfig.NewXltProxyController())
	//if err != nil {
	//	mylog.Error(err, "unable to create manager")
	//	os.Exit(1)
	//}

	proxyCtl:=k8sconfig.NewXltProxyController()
	err=builder.ControllerManagedBy(mgr).
		For(&v1.Ingress{}).
		Watches(&source.Kind{
			Type: &v1.Ingress{},
		},
			handler.Funcs{DeleteFunc: proxyCtl.OnDelete}).
		Complete(proxyCtl)

	if err = builder.ControllerManagedBy(mgr).
		For(&v1.Ingress{}).Complete(k8sconfig.NewXltProxyController()); err != nil {
		mylog.Error(err, "unable to create manager")
		os.Exit(1)
	}

	sysinit.InitConfig() //初始化  业务系统配置
	errCh := make(chan error)
	go func() {
		//启动控制器管理器
		if err = mgr.Start(signals.SetupSignalHandler()); err != nil {
			errCh <- err
		}
	}()
	go func() {
		// 启动网关
		if err = fasthttp.ListenAndServe(fmt.Sprintf(":%d", sysinit.SysConfig.Server.Port), ProxyHandler); err != nil {
			errCh <- err
		}
	}()
	getError := <-errCh
	log.Println(getError.Error())
}
