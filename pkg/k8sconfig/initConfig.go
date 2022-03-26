package k8sconfig

import (
	"log"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func K8sRestConfig() *rest.Config{
	config, err := clientcmd.BuildConfigFromFlags("","./resources/config" )
	if err!=nil{
	   log.Fatal(err)
	}
	config.Insecure=true
	return config
}
