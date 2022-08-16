package main

import "fmt"

func main() {
	/* In Go language, a map is a powerful, ingenious, and versatile data structure. Golang Maps is a collection of unordered pairs of key-value. It is widely used because it provides fast lookups and values that can retrieve, update or delete with the help of keys */
	animals := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	} /* Map Declaration and initialization */
	fmt.Println(animals)
	fmt.Println(animals[90]) /* Accessing the value of the map */

	/* Updating the value of the map */
	animals[90] = "Lion"
	fmt.Println(animals[90])

	/* Deleting the value of the map */
	delete(animals, 90)

	/* Iterating over the map */
	for key, value := range animals {
		fmt.Println(key, value)
	}

	/* Map declaration using make() */
	animals2 := make(map[int]string)
	animals2[90] = "Lion"
	animals2[91] = "Cat"
	animals2[92] = "Cow"
	fmt.Println(animals2)

	/* Map declaration using make() with initial capacity */
	animals3 := make(map[int]string, 10)
	fmt.Println(animals3)
}
