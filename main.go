package main

import (
	"os"
)

const (
	address string = "https://kubernetes"
)

func main() {

	//fmt.Println(MakeReqGet("/api/v1/namespaces/default/pods/"))

	MakeReqStream("/api/v1/watch/namespaces/default/pods/management-748d988bf8-h8s6m")

}

func get_namespace() string {
	//TODO: Get namespace automatically
	return os.Getenv("NAMESPACE")
}
