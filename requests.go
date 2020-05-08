package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	address string = "https://kubernetes"
)

func MakeReqGet(url string) string {

	url = address + url
	token, err := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		fmt.Print(err)
	}

	str := string(token)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	TOKEN := fmt.Sprintf("%s %s ", "Bearer", str)
	req.Header.Add("Authorization", TOKEN)

	if err != nil {
		log.Fatalln(err)
	}
	resp, err2 := client.Do(req)

	if err2 != nil {
		log.Fatal("Error reading response. ", err)

	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	return string(body)
}

func MakeReqPatch(url string, data string) {

	token, err := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		fmt.Print(err)
	}

	str := string(token)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	var jsonData = []byte(data)
	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
	TOKEN := fmt.Sprintf("%s %s ", "Bearer", str)

	req.Header.Add("Authorization", TOKEN)
	req.Header.Add("Content-Type", "application/strategic-merge-patch+json")

	if err != nil {
		log.Fatalln(err)
	}
	resp, err2 := client.Do(req)

	if err2 != nil {
		log.Fatal("Error reading response. ", err)

	}

	// Read body from response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	fmt.Println("Following data is sended to api:")
	fmt.Println(string(body))
}

func MakeReqStream(url string) {

	url = address + url
	token, err := ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		fmt.Print(err)
	}

	str := string(token)
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	TOKEN := fmt.Sprintf("%s %s ", "Bearer", str)
	req.Header.Add("Authorization", TOKEN)

	if err != nil {
		log.Fatalln(err)
	}

	resp, err2 := client.Do(req)

	if err2 != nil {
		log.Fatal("Error reading response. ", err)

	}

	reader := bufio.NewReader(resp.Body)
	for {
		line, _ := reader.ReadBytes('\n')
		log.Println(string(line))
		//TODO: Trigger appropriate functions
	}

}
