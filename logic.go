package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func Dependons_controller(json string) {
	dependon_name := gjson.Get(json, "object.metadata.name").String()
	dependon_url := kubernetesDependonWatchURL + os.Getenv("NAMESPACE") + "/dependons/" + dependon_name
	ch := make(chan string)
	go MakeReqStream(dependon_url, ch)
	for {
		dependonJson := <-ch
		operationType := gjson.Get(dependonJson, "type").String()
		if operationType == "ADDED" {
			fmt.Println("-----------------------------------------------")
			fmt.Println("New Dependon Added and Now Watching To : " + dependon_name)
			go Workload_controller(dependonJson)
		} else if operationType == "DELETED" {
			fmt.Println("-----------------------------------------------")
			fmt.Println("Dependon has been deleted: " + dependon_name)
			break
		} else if operationType == "MODIFIED" {
			fmt.Println("-----------------------------------------------")
			fmt.Println("Dependon has been modified.Thread is stopping: " + dependon_name)
			break
		}

	}
}

func Workload_controller(json string) {
	centralStatefulsets := gjson.Get(json, "object.spec.centrals.statefulsets")
	centralDeployments := gjson.Get(json, "object.spec.centrals.deployments")
	dependentStatefulsets := gjson.Get(json, "object.spec.dependons.statefulsets")
	dependentDeployments := gjson.Get(json, "object.spec.dependons.deployments")
	//dependonName := gjson.Get(json, "object.spec.dependons.deployments").String()
	fmt.Println("Dependent Statefulsets: " + dependentStatefulsets.String())
	fmt.Println("Dependent Deployments: " + dependentDeployments.String())
	ch := make(chan string)
	for _, statefulset := range centralStatefulsets.Array() {

		URL := kubernetesBaseWatchURL + os.Getenv("NAMESPACE") + "/statefulsets/" + statefulset.String()
		if statefulset.String() != "" || statefulset.String() != "null" {
			go MakeReqStream(URL, ch)
			fmt.Println("Start watching central statefulset: " + statefulset.String())
		}
	}
	for _, deployment := range centralDeployments.Array() {

		URL := kubernetesBaseWatchURL + os.Getenv("NAMESPACE") + "/deployments/" + deployment.String()
		if deployment.String() != "" || deployment.String() != "null" {
			fmt.Println("Start watching central deployment: " + deployment.String())
			go MakeReqStream(URL, ch)
		}
	}
	for {

		operation := gjson.Get(<-ch, "type").String()
		if operation == "DELETED" || gjson.Get(<-ch, "object.status.readyReplicas").Int() == 0 {
			fmt.Println("Some of Central Pods are not running properly.Dependent workloads are stopping.")
			for _, dependentStatefulset := range dependentStatefulsets.Array() {
				Stop_statefulset(dependentStatefulset.String())
			}
			for _, dependentDeployment := range dependentDeployments.Array() {
				Stop_deployment(dependentDeployment.String())
			}
			URL := kubernetesBaseWatchURL + os.Getenv("NAMESPACE") + "/" + strings.ToLower(gjson.Get(<-ch, "object.kind").String()) + "s/" + gjson.Get(<-ch, "object.metadata.name").String()

			ch2 := make(chan string)
			go MakeReqStream(URL, ch2)
			for {
				operation = gjson.Get(<-ch2, "type").String()
				//fmt.Println(<-ch2)
				if operation != "DELETED" || gjson.Get(<-ch2, "object.status.readyReplicas").Int() >= 1 {
					fmt.Println("Central Pod is now healty.Dependent workloads are starting.")
					for _, dependentStatefulset := range dependentStatefulsets.Array() {
						Start_statefulset(dependentStatefulset.String())
					}
					for _, dependentDeployment := range dependentDeployments.Array() {
						Start_deployment(dependentDeployment.String())
					}
					break
				}
			}

		}

	}
}
