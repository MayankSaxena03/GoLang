package main

import "fmt"

func main() {
	/* Defer statements are executed at end of function. If there are multiple defer statements LIFO will be followed */
	Planet()
	fmt.Println("")
	PrintReverse()
}

func Planet() {
	defer fmt.Println("Earth")
	fmt.Println("Sun")
	defer fmt.Println("Mercury")
	defer fmt.Println("Venus")
	fmt.Println("Mars")
	/* Output of above will be -> Sun Mars Venus Mercury Earth ;;; Sun and Mars will be printed first as they are not defer statements. Then defer statements willl run bottom to top in main() or they will run in order of LIFO. */
}

func PrintReverse() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i, " ")
	}
}
