package main

import (
	"fmt"
	"io/ioutil"
)

/*
const (
	address string = "https://kubernetes"
)
*/
func main() {

	namespace, _ := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/namespace")
	fmt.Print("Namespace is :", string(namespace))
	//fmt.Println(MakeReqGet("/api/v1/namespaces/default/pods/"))

	//MakeReqStream("/api/v1/watch/namespaces/default/pods/management-748d988bf8-h8s6m")

}
