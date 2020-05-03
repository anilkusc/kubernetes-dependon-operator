package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const address := "https://kubernetes"

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