package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Conversions in Golang")
	fmt.Print("Enter a number between 1 & 5 : ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	fmt.Print("Thanks for rating us : ", input)

	numRating, err := strconv.ParseFloat(input, 64) /* This statement won't convert input to float as there is line break associated to it */
	fmt.Println("Rating before trimming space : ", numRating)
	numRating, err = strconv.ParseFloat(strings.TrimSpace(input), 64) /* This statement will convert input to float as there is no line break associated to it */
	fmt.Println("Rating after trimming space : ", numRating)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Rating is : ", numRating)
	}
}
