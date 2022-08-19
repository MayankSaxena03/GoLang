package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dice := rand.Intn(10) + 1

	/* Syntax */
	switch dice {
	case 1:
		fmt.Println("Dice rolled 1")
	case 2:
		fmt.Println("Dice rolled 2")
	case 3:
		fmt.Println("Dice rolled 3")
	case 4:
		fmt.Println("Dice rolled 4")
		fallthrough /* If case 4 is true then statement of case 4 will run along with steatement below it i.e. case 5 */
	case 5:
		fmt.Println("Dice rolled 5")
	case 6:
		fmt.Println("Dice rolled 6")
	default:
		fmt.Println("Dice rolled something else")
	}
}
