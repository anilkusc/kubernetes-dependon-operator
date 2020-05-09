package main

import (
	"fmt"
	"os"

	//      "os"
	"strconv"

	"github.com/tidwall/gjson"
)

const (
	deployment_url  string = "https://kubernetes/apis/apps/v1/namespaces/"
	statefulset_url string = "https://kubernetes/apis/apps/v1/namespaces/"
	dependons_url   string = "https://kubernetes/apis/anilkuscu.github.com/v1alpha1/namespaces/"
)

func Trigger(json string) {

	workload_type := gjson.Get(json, "object.kind").String()

	if workload_type == "Deployment" {

	} else if workload_type == "StatefulSet" {

	}

}

func Stop_deployment(deployment_name string) {
	url := deployment_url + os.Getenv("NAMESPACE") + "/deployments/" + deployment_name
	data := `{"spec":{"replicas":0}}`
	MakeReqPatch(url, data)
	fmt.Println("Deployment " + deployment_name + "has been stopped.")
}

func Start_deployment(deployment_name string) {
	url := deployment_url + os.Getenv("NAMESPACE") + "/deployments/" + deployment_name
	data := `{"spec":{"replicas":1}}`
	MakeReqPatch(url, data)
	fmt.Println("Deployment " + deployment_name + "has been started.")
}

func Stop_statefulset(statefulset_name string) {
	url := statefulset_url + os.Getenv("NAMESPACE") + "/statefulset/" + statefulset_name
	data := `{"spec":{"replicas":0}}`
	MakeReqPatch(url, data)
	fmt.Println("Statefulset " + statefulset_name + "has been stopped.")
}

func Start_statefulset(statefulset_name string) {
	url := statefulset_url + os.Getenv("NAMESPACE") + "/statefulset/" + statefulset_name
	data := `{"spec":{"replicas":1}}`
	MakeReqPatch(url, data)
	fmt.Println("Statefulset " + statefulset_name + "has been started.")
}

func Get_dependents() ([]string, []string) {
	var sts, deploy []string
	return sts, deploy
}

func Get_centrals() {

}

func Get_dependons() []string {
	var dependon_names []string
	url := dependons_url + os.Getenv("NAMESPACE") + "/dependons/"
	dependons_info := MakeReqGet(url)
	dependons := gjson.Get(dependons_info, "items.#")
	dependons_count := int(dependons.Int())
	for i := 0; i < dependons_count; i++ {
		count := strconv.Itoa(i)
		dependon_name := gjson.Get(dependons_info, "items."+count+".metadata.name").String()
		dependon_names = append(dependon_names, dependon_name)
	}
	return dependon_names
}
