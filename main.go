package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
	//	"sync"
)

const (
	apiAddress                 string = "https://kubernetes"
	kubernetesDependonWatchURL string = apiAddress + "/apis/anilkuscu.github.com/v1alpha1/watch/namespaces/"
	kubernetesBaseWatchURL     string = apiAddress + "/apis/apps/v1/watch/namespaces/"
	kubernetesDependonURL      string = apiAddress + "/apis/anilkuscu.github.com/v1alpha1/namespaces/"
	kubernetesBaseURL          string = apiAddress + "/apis/apps/v1/namespaces/"
)

func main() {
	namespace, _ := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/namespace")
	os.Setenv("NAMESPACE", string(namespace))
	if os.Getenv("NAMESPACE") == "" {
		os.Setenv("NAMESPACE", "default")
	}

	dependon_url := kubernetesDependonWatchURL + os.Getenv("NAMESPACE") + "/dependons/"
	ch := make(chan string)
	go MakeReqStream(dependon_url, ch)

	for {
		dependonJson := <-ch
		operationType := gjson.Get(dependonJson, "type").String()
		if operationType == "ADDED" {
			go Dependons_controller(dependonJson)
		}
		if operationType == "MODIFIED" {
			fmt.Println("Dependon has been modified.The new thread is starting...")
			go Dependons_controller(dependonJson)
		}
		if operationType == "DELETED" {
			URL := kubernetesDependonURL + os.Getenv("NAMESPACE") + "/dependons"
			dependonList := MakeReqGet(URL)
			dependonList = gjson.Get(dependonList, "items.#.metadata.name").String()
			fmt.Println("This dependon has been deleted: " + gjson.Get(dependonJson, "object.metadata.name").String())
			fmt.Println("Dependon List : " + dependonList)
		} else if operationType == "" {
			fmt.Println("Unauthorized operation.")
			fmt.Println(dependonJson)
			break
		}

	}
	/*
		for {

			restart_all = false
			for restart_all == false {
				var wg sync.WaitGroup
				wg.Add(1)
				go func() {
					Dependons_controller()
					wg.Done()
				}()
				wg.Wait()
			}

		}
	*/
}
