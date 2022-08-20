package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformGetRequest()
	fmt.Println("")
	PerformPostJSONRequest()
	fmt.Println("")
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:1111/get"

	resp, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("Content Length:", resp.ContentLength)

	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Content:", string(content))
}

func PerformPostJSONRequest() {
	const myUrl = "http://localhost:1111/post"

	requestBody := strings.NewReader(`{
		"name":"John Doe",
		"age":30,
		"city":"New York"
		}`)
	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("Content Length:", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content:", string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:1111/post"

	data := url.Values{}
	data.Add("name", "John Doe")
	data.Add("age", "30")
	data.Add("city", "New York")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content:", string(content))
}
