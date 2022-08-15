package main

import "fmt"

func main() {
	var ptr *int /* Integer Pointer declaration */
	fmt.Println("Default value of pointer is ", ptr)

	num := 10
	ptr = &num /* Passing address of num value in ptr */
	fmt.Printf("Value of num is %d\n", num)
	fmt.Printf("Value stored in ptr is %d\n", *ptr)

	*ptr = 20 /* This will change the value of num to 20 */
	fmt.Printf("Value of num is %d\n", num)
	fmt.Printf("Value stored in ptr is %d\n", *ptr)
	/*
		& -> Reference Operator
		* -> Dereference Operator
	*/
}
