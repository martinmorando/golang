/*
	Maps
	- Unordered collection of keys and values
	- Specify key and value data types
*/

package main

import "fmt"

func main() {

	// Create an empty map
	emptyMap := make(map[string]int)
	fmt.Println(emptyMap)			// map[]


	// Create and initialize a map
	exampleMap := map[string]int{
		"A": 12,
		"B": 10,
		"C": 90,
	}
	fmt.Println(exampleMap)			// map[A:12 B:10 C:90]

}