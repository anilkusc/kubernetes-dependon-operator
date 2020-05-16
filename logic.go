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
		if len(new_dependons) == len(dependons) {

			for i, _ := range dependons {
				if new_dependons[i] == dependons[i] {
					continue
				} else {
					fmt.Println("Dependons has been changed.")
					break
				}
			}
		} else {
			fmt.Println("Dependons has been changed.")
			break
		}

	}
	restart_all = true
}

func Workload_controller() {

}
