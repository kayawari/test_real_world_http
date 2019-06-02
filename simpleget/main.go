package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
	"bytes"
	"os"
	"io"
	"mime/multipart"
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

func postBody() {
	reader := strings.NewReader("(=^^=)(=^^=)(=^^=)")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {panic(err)}
	log.Println("status:", resp.Status)
	log.Println("contentn_length:", resp.ContentLength)
}

func postByMultipleFormData() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Takayuki Kayawari")
	fileWriter, err := writer.CreateFormFile("thumnail", "photo.jpg")
	if err != nil {panic(err)}
	readFile, err := os.Open("photo.jpg")
	if err != nil {panic(err)}
	defer readFile.Close()
	io.Copy(fileWriter, readFile)
	writer.Close()
	
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {panic(err)}
	log.Println("status:", resp.Status)
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
	postBody()
	log.Println("=====================")
	postByMultipleFormData()
}