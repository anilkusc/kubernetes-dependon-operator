package main

import (
	"fmt"
	"os"

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

func Get_dependents(dependon_name string) ([]string, []string) {
	var statefulsets, deployments []string
	url := dependons_url + os.Getenv("NAMESPACE") + "/dependons/" + dependon_name
	dependon_info := MakeReqGet(url)
	dependent_deployments := gjson.Get(dependon_info, "spec.dependents.deployments.#")
	dependent_deployment_count := int(dependent_deployments.Int())
	for i := 0; i < dependent_deployment_count; i++ {
		count := strconv.Itoa(i)
		deployment_name := gjson.Get(dependon_info, "spec.dependents.deployments."+count).String()
		deployments = append(deployments, deployment_name)
	}
	dependent_statefulsets := gjson.Get(dependon_info, "spec.dependents.statefulsets.#")
	dependent_statefulsets_count := int(dependent_statefulsets.Int())
	for i := 0; i < dependent_statefulsets_count; i++ {
		count := strconv.Itoa(i)
		statefulset_name := gjson.Get(dependon_info, "spec.dependents.statefulsets."+count).String()
		statefulsets = append(statefulsets, statefulset_name)
	}
	return statefulsets, deployments
}

func Get_centrals(dependon_name string) ([]string, []string) {
	var statefulsets, deployments []string
	url := dependons_url + os.Getenv("NAMESPACE") + "/dependons/" + dependon_name
	dependon_info := MakeReqGet(url)
	central_deployments := gjson.Get(dependon_info, "spec.centrals.deployments.#")
	dependon_deployment_count := int(central_deployments.Int())
	for i := 0; i < dependon_deployment_count; i++ {
		count := strconv.Itoa(i)
		deployment_name := gjson.Get(dependon_info, "spec.centrals.deployments."+count).String()
		deployments = append(deployments, deployment_name)
	}
	dependon_statefulsets := gjson.Get(dependon_info, "spec.centrals.statefulsets.#")
	dependon_statefulsets_count := int(dependon_statefulsets.Int())
	for i := 0; i < dependon_statefulsets_count; i++ {
		count := strconv.Itoa(i)
		statefulset_name := gjson.Get(dependon_info, "spec.centrals.statefulsets."+count).String()
		statefulsets = append(statefulsets, statefulset_name)
	}
	return statefulsets, deployments
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
