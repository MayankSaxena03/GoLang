package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	file, err := os.Create("./myFile.txt") /* This will create a file in current directory */
	if err != nil {
		panic(err) /* Panic will stop execution of programs and show us the error */
	}

	/* If file already exist we can use os.OpenFile */

	defer file.Close() /* This will close the file after execution */

	io.WriteString(file, "Added this from Go Program") /* This will write the string to the file */

	data, err := ioutil.ReadFile("./myFile.txt") /* This will read the file */
	if err != nil {
		panic(err)
	}

	println(string(data)) /* This will print the data in the file */
}
