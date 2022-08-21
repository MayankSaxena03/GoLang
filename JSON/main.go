package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"course_name"` /* Set alias to course_name */
	Price    int
	Platform string   `json:"language"`
	Password string   `json:"-"` /* This field will not be exported */
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	EncodeJSON()
	fmt.Println("\n")
	DecodeJSON()
}

func EncodeJSON() {
	courses := []course{
		{"Go Language", 100, "Go", "", []string{"Go", "Programming", "Language"}},
		{"Java Language", 200, "Java", "", []string{"Java", "Programming", "Language"}},
		{"Regex", 300, "Regex", "", nil},
	}

	json, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(json))
}

func DecodeJSON() {
	jsonData := []byte(`
	{
		"course_name": "Java Language",
        "Price": 200,
        "language": "Java",
        "tags": ["Java","Programming","Language"]
	}
	`)

	var course1 course
	checkValid := json.Valid(jsonData)
	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &course1)
		fmt.Printf("%+v\n", course1)
	} else {
		fmt.Println("Invalid JSON")
	}

	fmt.Println("\n\n")
	/* Without using a struct */
	var myData map[string]interface{}
	json.Unmarshal(jsonData, &myData)
	fmt.Printf("%+v\n", myData)
	for k, v := range myData {
		fmt.Printf("%s: %s -> %T\n", k, v, v)
	}
}
