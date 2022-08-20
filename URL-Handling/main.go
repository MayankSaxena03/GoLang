package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://some-website.com:3000/test?unique=try&gateway=payu" /* A simple HTML page */

func main() {
	fmt.Println(myurl)

	/* Parse URL */
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("")

	fmt.Println("Result : ", result)
	fmt.Println("Scheme : ", result.Scheme)
	fmt.Println("Host : ", result.Host)
	fmt.Println("Path : ", result.Path)
	fmt.Println("Port : ", result.Port())
	fmt.Println("RawQuery : ", result.RawQuery)

	fmt.Println("")

	/* Parse Query */
	query := result.Query()
	fmt.Printf("Type of query : %T\n", query)
	fmt.Println("Query : ", query)
	fmt.Println("unique : ", query["unique"])

	for k, v := range query {
		fmt.Println(k, ":", v)
	}

	fmt.Println("")

	/* Build URL */
	newURL := &url.URL{
		Scheme:   "https",
		Host:     "www.google.com",
		Path:     "/search",
		RawQuery: "q=go+lang",
	}

	anotherURL := newURL.String()
	fmt.Println("New URL : ", anotherURL)
}
