package k8sconfig

import (
	"context"
	"fmt"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type XltProxyController struct {
	client.Client
}

func NewXltProxyController() *XltProxyController {
	return &XltProxyController{}
}

func (r *XltProxyController) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	//obj := &v1.Ingress{}
	//if err := r.Get(ctx, req.NamespacedName, obj); err != nil {
	//	return reconcile.Result{}, err
	//}
	//if v, ok := obj.Annotations["kubernetes.io/ingress.class"]; ok && v == "jtthink" {
	//	fmt.Println(obj)
	//}
	//return reconcile.Result{}, nil

	obj:=&Route{}
	err:=r.Get(ctx,req.NamespacedName,obj)
	if err!=nil{
		return reconcile.Result{}, err
	}
	fmt.Println(obj)
	return reconcile.Result{}, nil
}

func (r *XltProxyController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}
