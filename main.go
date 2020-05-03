package main

import (
	"fmt"
)

func main() {

	fmt.Println(MakeReqGet("/api/v1/namespaces/default/pods/"))
}
