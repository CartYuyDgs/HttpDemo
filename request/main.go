package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	context, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Printf("%s\n", context)
}

func requestByParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	params := make(url.Values)
	params.Add("name", "aaa")
	params.Add("age", "18")
	fmt.Println(params.Encode())
	request.URL.RawQuery = params.Encode()

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	printBody(r)
}

func byHeader() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("user-agent", "chrome")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}

	printBody(r)
}

func main() {
	//设置查询参数
	//定制请求头
	requestByParams()
	byHeader()
}
