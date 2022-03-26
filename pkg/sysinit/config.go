package sysinit

import (
	"fmt"
	"io/ioutil"
	"log"

	"k8s.io/api/networking/v1"
	"sigs.k8s.io/yaml"
)

var SysConfig = new(SysConfigStruct)

type Server struct {
	Port int //代表是代理启动端口
}
type SysConfigStruct struct {
	Server  Server
	Ingress v1.IngressSpec
}

func InitConfig() {
	config, err := ioutil.ReadFile("./app.yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(config, SysConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(SysConfig.Ingress)
	ParseRule()
}
