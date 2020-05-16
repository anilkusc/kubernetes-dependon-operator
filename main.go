package main

import (
	"io/ioutil"
	"os"
)

var restart_all bool = false

/*
const (
	address string = "https://kubernetes"
)
*/
func main() {
	namespace, _ := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/namespace")
	os.Setenv("NAMESPACE", string(namespace))

	for {

		restart_all = false
		for restart_all == false {
			go Dependons_controller()
			select {}
		}

	}

}
