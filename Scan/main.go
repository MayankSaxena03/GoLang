package main

import (
	"fmt"
)

func main() {
	/* Scan Function */
	/* In Go language, fmt package implements formatted I/O with functions analogous to C’s printf() and scanf() function. The fmt.Scan() function in Go language scans the input texts which is given in the standard input, reads from there and stores the successive space-separated values into successive arguments */
	// Declaring some variables
	var name string
	var alphabet_count int

	// Calling Scan() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Print("Enter your name and number of alphabets in your name : ")
	fmt.Scan(&name, &alphabet_count) /* Will take all inputs seperated by space. EOL will be count as space character */
	// Printing the given texts
	fmt.Printf("The word %s containing %d number of alphabets.\n\n", name, alphabet_count)

	/* Scanln Function */
	/* In Go language, fmt package implements formatted I/O with functions analogous to C’s printf() and scanf() function. The fmt.Scanln() function in Go language scans the input texts which is given in the standard input, reads from there and stores the successive space-separated values into successive arguments. This function stops scanning at a newline and after the final item, there must be a newline or EOF. */

	// Calling Scanln() function for
	// scanning and reading the input
	// texts given in standard input
	var name2 string
	var alphabet_count2 int
	fmt.Print("Enter your name and number of alphabets in your name : ")
	fmt.Scanln()                         /* Will clear the previous EOL characters so next statement could work. */
	fmt.Scanln(&name2, &alphabet_count2) /* Will take input until EOL. EOL will be considered as end of scan and further remaining values will not be scanned. If there are extra values, they will be ignored. */
	// Printing the given values
	fmt.Printf("The word %s containing %d number of alphabets.\n\n", name2, alphabet_count2)

	/* Scanf */
	/* In Go language, fmt package implements formatted I/O with functions analogous to C’s printf() and scanf() function. The fmt.Scanf() function in Go language scans the input texts which is given in the standard input, reads from there and stores the successive space-separated values into successive arguments as determined by the format. */

	// Declaring some variables
	var name3 string
	var alphabet_count3 int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Print("Enter your name and number of alphabets in your name : ")
	fmt.Scanf("%s %d", &name3, &alphabet_count3) /* It works similar to scanf in C. */
	// Printing the given values
	fmt.Printf("The word %s containing %d number of alphabets.\n\n", name3, alphabet_count3)
}
