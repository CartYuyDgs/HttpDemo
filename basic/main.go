package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Println("err:", err)
	}
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%s\n", content)
}

func put() {
	req, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	}

	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%s\n", content)
}

func post() {
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%s\n", content)
}

func delete() {
	req, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		fmt.Println("err:", err)
	}
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("err:", err)
	}

	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Printf("%s\n", content)
}

func main() {
	//get()
	//post()
	put()
	delete()
}
