package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin) /* Will create a new reader which reads from standard I/O */
	fmt.Print("Enter rating for our Pizza : ")
	//comma ok || error ok syntax
	input, err := reader.ReadString('\n') /* If input is successful, it is stored in input otherwise error is stored in err */
	fmt.Print("Thanks for rating us : ", input)
	fmt.Printf("Type of rating is %T\n", input)
	fmt.Printf("Type of error is %T\n", err) /* It will show nil if input is successful */
}
