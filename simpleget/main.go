package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
	"os"
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
	if err != nil {panic(err)}
	log.Println("status:", resp.Status)
	log.Println("status_code:", resp.StatusCode)
}

// x-www-form-urlencodedによるフォーム送信
func getByForm() {
	values := url.Values {
		"test": {"value"},
		"test2": {"value2"},
		"test3": {"value3=&=&"},
		// &とか=もRFC1866形式でエンコードしてくれる。
	}
	resp, err := http.PostForm("http://localhost:18888", values)
	if err != nil {panic(err)}
	log.Println("status:", resp.Status)
}

func postByMultipleForm() {
	file, err := os.Open("main.go")
	if err != nil {panic(err)}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {panic(err)}
	log.Println("status:", resp.Status)
	log.Println("contentn_length:", resp.ContentLength)
}

func main () {
	simpleGet()
	log.Println("=====================")
	getWithQuery()
	log.Println("=====================")
	simpleHead()
	log.Println("=====================")
	getByForm()
	log.Println("=====================")
	postByMultipleForm()
}