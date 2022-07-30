package main

import "reflect"

type MyType struct {
	Name string
}

func printMyType(t *MyType) {
	println(t.Name)
}

type MyTagType struct {
	Name string `json:"name"`
}

type ServiceType string

const (
	ServiceTypeClusterIPServiceType    = "ClusterIP"
	ServiceTypeNodePortServiceType     = "NodePort"
	ServiceTypeLoadBalancerServiceType = "LoadBalancer"
	ServiceTypeExternalNameServiceType = "ExternalName"
)

func main() {
	t := MyType{Name: "test"}
	printMyType(&t)

	mt := MyTagType{Name: "test"}
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	println(tag)
}
