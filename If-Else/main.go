package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter your age : ")
	fmt.Scanf("%d", &age)

	if age < 18 {
		fmt.Println("You are a minor")
	} else if age < 30 {
		fmt.Println("You are young")
	} else if age < 50 {
		fmt.Println("You are middle-aged")
	} else {
		fmt.Println("You are old")
	}

	/*
		OR  -> ||
		AND -> &&
	*/

	fmt.Scanln() /* Clearing input buffer */

	var num int
	fmt.Print("Enter a number : ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	/* Updating variable and cheecking in one line */
	if num = 3; num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
	/* Above Syntax can be used when we are fetching value online and checking it using if-else in one line */
}
