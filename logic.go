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
		fmt.Println("1")
		for i, _ := range dependons {
			fmt.Println("2")
			if new_dependons[i] == dependons[i] {
				fmt.Println("3")
				continue

			} else {
				fmt.Println("Dependons has been changed.")
				break
			}

		}
		fmt.Println("4")

	}
	fmt.Println("Restarted all")
	//	restart_all = true
}

func Workload_controller() {

}
