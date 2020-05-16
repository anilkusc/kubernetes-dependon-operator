package main

import (
	"io/ioutil"
	"os"
	"sync"
)

var restart_all bool = false

func main() {
	namespace, _ := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/namespace")
	os.Setenv("NAMESPACE", string(namespace))

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

}
