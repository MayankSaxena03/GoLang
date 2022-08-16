package main

import "fmt"

func main() {
	var numberArray [5]int = [5]int{1, 2, 3, 4, 5} /* Array Declaration and initialization */
	fmt.Println(numberArray)
	numberArray[0] = 54
	fmt.Println(numberArray)
	fmt.Println("Length of array is:", len(numberArray))
}
