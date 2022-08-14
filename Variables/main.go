package main

import "fmt"

/* someNumber := 1000 -> This will be invalid as we cannot use walrus operator outside of function. */

const PRN = 1914110257 /* When the first letter of variable or constant name is capital, the variable will be declared as public. */

func main() {
	/* Basic syntax for string */
	var username string = "Mayank"
	/* If the value is not initialized for integer the default value is an empty string. */
	fmt.Println(username)
	fmt.Printf("Username is of type %T\n\n", username)

	/* Basic Syntax for Boolean */
	var flag bool = true
	/* If the value is not initialized for boolean the default value is false. */
	fmt.Printf("Value of flag is ")
	fmt.Print(flag)
	fmt.Printf(" and is of type %T\n\n", flag)

	/* Basic Syntax for integer */
	var smallVal uint8 = 255 /* uint8 takes unsigned integer value of 8 bit which ranges form 0-255. We can use uint16, uint32 & uint64 for different size of integers. In general we can use int. */
	/* If the value is not initialized for integer the default value is 0. */
	fmt.Printf("Value of smallVal is ")
	fmt.Print(smallVal)
	fmt.Printf(" and is of type %T\n\n", smallVal)

	/* Basic Syntax for float */
	var smallFloat float64 = 34532.97489274239847 /* float64 store 64-bit floating point number. It is more precise than float32 */
	/* If the value is not initialized for float the default value is 0. */
	fmt.Printf("Value of smallFloat is ")
	fmt.Print(smallFloat)
	fmt.Printf(" and is of type %T\n\n", smallFloat)

	/* Basic syntax for constant */
	const Pi = 3.14
	fmt.Printf("Value of Pi is ")
	fmt.Print(Pi)
	fmt.Printf(" and is of type %T\n\n", Pi)
	/* Pi=2.78 -> This will give an error as we cannot change the value of constant. */

	/* Implicit Declaration */
	var implicitVal = "Hello World" /* It will automatically detect and set the type of implicitVal to string. */
	/* implicitVal = 3; This statement is invalid because implicitVal is already set to string. */
	fmt.Printf("Value of implicitVal is ")
	fmt.Print(implicitVal)
	fmt.Printf(" and is of type %T\n\n", implicitVal)

	/* no var style */
	someInt := 5 /* We can use walrus operator ":=" and avoid using var. It can be used within function. All other rules are same as implicit declaration. */
	fmt.Printf("Value of someInt is ")
	fmt.Print(someInt)
	fmt.Printf(" and is of type %T\n\n", someInt)
}
