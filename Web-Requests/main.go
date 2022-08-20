package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://help.websiteos.com/websiteos/example_of_a_simple_html_page.htm" /* A simple HTML page */

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() /* Always close response body */

	fmt.Printf("Type of response: %T\n", response)
	fmt.Printf("Status code: %d\n", response.StatusCode)

	fmt.Println("\n\n\n")

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of data: %T\n\n", data)
	fmt.Printf("Data: %s\n", data)
}
