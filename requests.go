package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"bufio"	
)

const (
	address string = "https://kubernetes"
	deploment_uri string = "https://kubernetes"
	statefulset_uri string = "https://kubernetes"
)

func MakeReqGet(url string) string {

	url = address+url
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

func MakeReqStream(url string) {

	url = address+url
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