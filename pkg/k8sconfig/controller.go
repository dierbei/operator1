package k8sconfig

import (
	"context"
	"k8s.io/client-go/util/workqueue"
	"log"
	"sigs.k8s.io/controller-runtime/pkg/event"

	"dierbei/operator-one/pkg/sysinit"

	v1 "k8s.io/api/networking/v1"
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
	ingress := &v1.Ingress{}
	err := r.Get(ctx, req.NamespacedName, ingress)
	if err != nil {
		return reconcile.Result{}, err
	}
	if r.IsJtProxy(ingress.Annotations) {
		err = sysinit.ApplyConfig(ingress)
		if err != nil {
			return reconcile.Result{}, err
		}
	}
	return reconcile.Result{}, nil
}

//判断是否 是否 我们所需要处理的ingress
func (r *XltProxyController) IsJtProxy(annotations map[string]string) bool {
	if v, ok := annotations["kubernetes.io/ingress.class"]; ok && v == "jtthink" {
		return true
	}
	return false
}

func (r *XltProxyController) OnDelete(event event.DeleteEvent, limitingInterface workqueue.RateLimitingInterface) {
	if r.IsJtProxy(event.Object.GetAnnotations()) {
		if err := sysinit.DeleteConfig(event.Object.GetName(), event.Object.GetNamespace()); err != nil {
			log.Println(err)
		}
	}
}

func (r *XltProxyController) InjectClient(c client.Client) error {
	r.Client = c
	return nil
}
