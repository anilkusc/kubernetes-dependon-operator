package main

import (
	"fmt"
	"os"

	"github.com/tidwall/gjson"
)

const (
	deployment_url  string = "https://kubernetes/apis/apps/v1/namespaces/" + os.Getenv("NAMEPSACE") + "/deployments/"
	statefulset_url string = "https://kubernetes/apis/apps/v1/namespaces/" + os.Getenv("NAMEPSACE") + "/statefulsets/"
)

func Trigger(json string) {

	workload_type := gjson.Get(json, "object.kind").String()
	if workload_type == "Deployment" {

	} else if workload_type == "Statefulset" {

	}

}

func Stop_deployment(deployment_name string) {
	url := deployment_url + deployment_name
	data := `{"spec":{"replicas":0}}`
	MakeReqPatch(url, data)
	fmt.Println("Deployment " + deployment_name + "has been stopped.")
}

func Start_deployment(deployment_name string) {
	url := deployment_url + deployment_name
	data := `{"spec":{"replicas":1}}`
	MakeReqPatch(url, data)
	fmt.Println("Deployment " + deployment_name + "has been started.")
}

func Stop_statefulset(statefulset_name string) {
	url := statefulset_url + statefulset_name
	data := `{"spec":{"replicas":0}}`
	MakeReqPatch(url, data)
	fmt.Println("Statefulset " + statefulset_name + "has been stopped.")
}

func Start_statefulset(statefulset_name string) {
	url := statefulset_url + statefulset_name
	data := `{"spec":{"replicas":1}}`
	MakeReqPatch(url, data)
	fmt.Println("Statefulset " + statefulset_name + "has been started.")
}
