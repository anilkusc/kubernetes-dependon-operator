package main

import (
	"fmt"
	"os"
)

const (
	deployment_url  string = "https://kubernetes/apis/apps/v1/namespaces/" + os.Getenv("NAMEPSACE") + "/deployments/"
	statefulset_url string = "https://kubernetes/apis/apps/v1/namespaces/" + os.Getenv("NAMEPSACE") + "/statefulsets/"
)

func Trigger(json string) {

	//TODO:trigger if change of the pods changed(start,stop,etc)
}

func stop_deployment(deployment_name string) {
	url := deployment_url + deployment_name
	data := `{"spec":{"replicas":0}}`
	MakeReqPatch(url, data)
	fmt.Println("Deployment " + deployment_name + "has been stopped.")
}

func start_deployment(json string) {

	//TODO:trigger if change of the pods changed(start,stop,etc)
}
func stop_statefulset(statefulset_name string) {
	url := statefulset_url + statefulset_name
	data := `{"spec":{"replicas":0}}`
	MakeReqPatch(url, data)
	fmt.Println("Statefulset " + statefulset_name + "has been stopped.")
}

func start_statefulset(json string) {

	//TODO:trigger if change of the pods changed(start,stop,etc)
}
