package main

import (
	"fmt"
	"sort"
)

func main() {
	/* Slices are a key data type in Go, giving a more flexible and powerful alternative to arrays. */
	/* Unlike arrays, slices are not allocated at compile time and can grow dynamically. */
	/* Slices can be created with the builtin make function; this is how you create dynamically-sized arrays. */

	var fruitList = []string{"Apple", "Orange", "Banana", "Grape"} /* Slice Declaration */
	fmt.Println(fruitList)
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	fruitList = append(fruitList, "Pineapple") /* Append to Slice */
	fmt.Println(fruitList)

	fruitList = append(fruitList[:2], fruitList[3:]...) /* Slice Splice -> This statement will remove the element at index 2 */
	fmt.Println(fruitList)

	/*
		We can also splice our list in following ways -
		fruitList = append(fruitList[1:3]) // This will make the fruitList contain only elements from index 1 & 2.
		fruitList = append(fruitList[:2]) // This will append elements from index 0 to 1.
		fruitList = append(fruitList[3:]) // This will append elements from index 3 to end.
	*/

	fmt.Println("\n")

	/* Slices can be created with the builtin make function; this is how you create dynamically-sized arrays. */
	var marks = make([]int, 4) /* Slice Declaration */
	fmt.Println(marks)         /* Default 0 values will be stored in integer slice */
	marks[0] = 90
	marks[1] = 86
	marks[2] = 85
	marks[3] = 87
	/* marks[4] = 50 -> This will give error as we have initialized the size already. */
	fmt.Println(marks)

	/* To add values we ccan use append. It will resize the slice */
	marks = append(marks, 81)
	fmt.Println(marks)
	fmt.Println("Marks are sorted :", sort.IntsAreSorted(marks)) /* Checks if the slice is sorted */

	/* Sort Slice */
	sort.Ints(marks)
	fmt.Println(marks)
	fmt.Println("Marks are sorted :", sort.IntsAreSorted(marks)) /* Checks if the slice is sorted */
}
