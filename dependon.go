package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/tidwall/gjson"
)

const (
	deployment_url  string = "https://kubernetes/apis/apps/v1/namespaces/" + os.Getenv("NAMEPSACE") + "/deployments/"
	statefulset_url string = "https://kubernetes/apis/apps/v1/namespaces/" + os.Getenv("NAMEPSACE") + "/statefulsets/"
	dependons_url   string = "https://kubernetes/apis/anilkuscu.github.com/v1alpha1/namespaces/" + os.Getenv("NAMEPSACE") + "/dependons/"
)

func Trigger(json string) {

	workload_type := gjson.Get(json, "object.kind").String()

	if workload_type == "Deployment" {

	} else if workload_type == "StatefulSet" {

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

func Get_dependents() ([]string, []string) {

}

func Get_centrals() {

}

func Get_dependons() []string {
	var dependon_names []string
	dependons_info := MakeReqGet(dependons_url)
	dependons := gjson.Get(dependons_info, "items.#")
	dependons_count := int(dependons.Int())
	for i := 0; i < dependons_count; i++ {
		count := strconv.Itoa(i)
		dependon_names[i] = gjson.Get(dependons_info, "items."+count+".metadata.name").String()
	}
	return dependon_names
}
