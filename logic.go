package main

import (
	"fmt"
)

func Dependons_controller() {
	var new_dependons []string
	dependons := Get_dependons()
	fmt.Println("Dependons are:")
	fmt.Println(dependons)

	for {
		new_dependons = Get_dependons()
		for i, _ := range dependons {
			if new_dependons[i] == dependons[i] {
				continue
			} else {
				fmt.Println("Dependons has been changed.")
				break
			}

		}

	}
	fmt.Println("Restarted all")
	//	restart_all = true
}

func Workload_controller() {

}
