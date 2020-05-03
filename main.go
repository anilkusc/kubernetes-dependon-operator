package main

func main() {

	//fmt.Println(MakeReqGet("/api/v1/namespaces/default/pods/"))
	MakeReqStream("/api/watch/v1/namespaces/default/pods/management-748d988bf8-h8s6m")

}
