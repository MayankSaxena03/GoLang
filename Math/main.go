package main

import (
	"fmt"
	"math/rand"
	"time"
	// "crypto/rand"
)

func main() {
	var number1 int = 10
	var number2 float64 = 4.17
	/* fmt.Println("Sum = ", number1+number2) -> This statement will be invalid as number1 is of type int and number2 is of type float64 */
	fmt.Println("Sum = ", float64(number1)+number2, "\n") /* This will convert number1 in type float and add it to number 2. We can convert number2 to int and get integer answer. */

	/* There are different math operations available like math.floor(), math.ceil(), math.Sqrt(), math.Pow(), etc. */

	/* Random Number using math/rand */
	fmt.Println("Random Number = ", rand.Intn(100)) /* This will generate a random number between 0 and 100. It will return same number if the function is called again. */
	/* If we change the seed, value we are getting constantly will also change */
	rand.Seed(42)
	fmt.Println("Random Number = ", rand.Intn(100)) /* We will get different number now as we passed different value to the seed. */
	/* We cannot change Seed value manually again & again every time we run the program. So, we pass time to seed which is never constant and will return different random number every time */
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Random Number = ", rand.Intn(100)) /* We will get different number. Even if we re-run, the value will not be same as the time is changed */

	/* Random number using crypto/rand */
	/* To run this code comment "Random Number using math/rand" section and math/rand import. Uncomment crypto/rand import */
	// randomNumber, _ := rand.Int(rand.Reader, big.NewInt(100)) /* This will generate a random number between 0 and 100. */
	// fmt.Println("Random Number = ", randomNumber)
}
