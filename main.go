package main

func main() {

	//fmt.Println(MakeReqGet("/api/v1/namespaces/default/pods/"))
	MakeReqStream("/api/watch/v1/namespaces/default/pods/")

}
