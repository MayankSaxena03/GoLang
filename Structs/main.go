package main

import "fmt"

/* A structure or struct in Golang is a user-defined type that allows to group/combine items of possibly different types into a single type. Any real-world entity which has some set of properties/fields can be represented as a struct. This concept is generally compared with the classes in object-oriented programming. It can be termed as a lightweight class that does not support inheritance but supports composition. */
/* There is no inheritance in GoLang */

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func main() {
	mayank := User{"Mayank", 20, "mayank@gmail.com", true}
	fmt.Println(mayank)
	fmt.Printf("%+v\n", mayank)
	fmt.Printf("Name is %v\nEmail is %v\n", mayank.Name, mayank.Email)
	/*
		%v	the value in a default format
		when printing structs, the plus flag (%+v) adds field names
		%#v	a Go-syntax representation of the value
		%T	a Go-syntax representation of the type of the value
		%%	a literal percent sign; consumes no value
	*/
}
