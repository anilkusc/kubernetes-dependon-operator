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

	dat, _ := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/namespace")
	fmt.Print(string(dat))
	//fmt.Println(MakeReqGet("/api/v1/namespaces/default/pods/"))

	//MakeReqStream("/api/v1/watch/namespaces/default/pods/management-748d988bf8-h8s6m")

}
