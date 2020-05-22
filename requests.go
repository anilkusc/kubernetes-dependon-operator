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

var token, _ = ioutil.ReadFile("/run/secrets/kubernetes.io/serviceaccount/token")

//var token, _ = ioutil.ReadFile("token")
var str = string(token)

func MakeReqGet(url string) string {

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
	defer resp.Body.Close()
	return string(body)
}

func MakeReqPatch(url string, data string) {

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

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}

	defer resp.Body.Close()
}

func MakeReqStream(url string, ch chan string) {

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
		line, _ := reader.ReadString('\n')
		ch <- string(line)
	}
	defer resp.Body.Close()

}

/*
func MakeReqStream(url string) {

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
		//		line, _ := reader.ReadString()
		log.Println(string(line))
		//TODO: Trigger appropriate functions

	}
	defer resp.Body.Close()
},*/

/*
func test(url string, ch chan string) {

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
	defer resp.Body.Close()
	fmt.Println("taking channel...")
	ch <- string(body)
}
*/
