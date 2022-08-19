package main

import "fmt"

func main() {
	greeter() /* Calling Function */

	fmt.Println(result(100, 94, 96), "%")
	fmt.Println(result2([]int{100, 94, 96}), "%")

	fmt.Println("")

	fmt.Println(adder(1, 2, 3, 4, 5))

	fmt.Println("")

	total, msg := adder2(1, 2, 3, 4, 5)
	fmt.Println(total, msg)
}

/* Function Declaration */
func greeter() {
	fmt.Println("Namaste")
}

/* Below function takes 3 values in argument and returns an integer value */
func result(marks1 int, marks2 int, marks3 int) int {
	return (marks1 + marks2 + marks3) / 3
}

/* Below function takes integer array in argument & returns float64 */
func result2(marks []int) float64 {
	var total int
	for _, value := range marks {
		total += value
	}
	return (float64(total)) / float64(len(marks))
}

/* Variadic Function */
/* In mathematics and in computer programming, a variadic function is a function of indefinite arity, i.e., one which accepts a variable number of arguments. */
func adder(values ...int) int {
	var total int
	for _, value := range values {
		total += value
	}
	return total
}

/* We can return multiple values through functions */
func adder2(values ...int) (int, string) {
	var total int
	for _, value := range values {
		total += value
	}
	return total, "\nTotal is " + fmt.Sprint(total) /* fmt.Sprint converts value to string */
}
