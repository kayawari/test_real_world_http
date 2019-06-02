package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
)

func simpleGet() {
	// 単純なget method
	resp, err := http.Get("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
	log.Println("status_code", resp.StatusCode)
	log.Println("headers:", resp.Header)
	log.Println("body:", string(body))
	log.Println("Content-Length", resp.Header.Get("Content-Length"))
}

func getWithQuery() {
	// クエリ付きget method
	values := url.Values{"query": {"hello world"}}
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

func simpleHead() {
	resp, err := http.Head("http://localhost:18888")
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
	log.Println("status_code:", resp.StatusCode)
}

func main () {
	simpleGet()
	log.Println("=====================")
	getWithQuery()
	log.Println("=====================")
	simpleHead()
}