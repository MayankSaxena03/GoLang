package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)
	fmt.Println("")

	/* for loop */
	/* No pre-increment in GoLang */
	fmt.Println("For Loop")
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("")

	/* for loop with range */
	fmt.Println("For Loop with range")
	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("")

	/* for each loop with range */
	fmt.Println("For Each type Loop with range")
	for i, day := range days {
		fmt.Println(i, day)
	}
	fmt.Println("")

	/* Continue, Break, Goto */
	rogueValue := 0
	fmt.Println("Continue, Break, Goto")
	for rogueValue < 10 {
		rogueValue++
		if rogueValue == 5 {
			continue
		}
		if rogueValue == 7 {
			break
		}
		fmt.Println(rogueValue)
	}

	fmt.Println("")

	goto welcome /* Will goto provided label */

welcome:
	fmt.Println("Welcome to GoLang")
}
