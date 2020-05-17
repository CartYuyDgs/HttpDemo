package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func body(r *http.Response) {
	context, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s\n", context)
}

func status(r *http.Response) {
	fmt.Println(r.Status)
	fmt.Println(r.StatusCode)
}

func header(r *http.Response) {
	fmt.Println(r.Header.Get("Content-type"))
}

func encoding() {

}

func main() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	defer func() { _ = r.Body.Close() }()

	body(r)
	status(r)
	header(r)
}
