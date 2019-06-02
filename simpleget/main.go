package main

import (
	"net/http"
	"io/ioutil"
	"log"
)

func main () {
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