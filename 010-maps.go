/*
	Maps
	- Unordered collection of keys and values
	- Specify key and value data types
*/

package main

import "fmt"

func main() {

	// [CREATE]
	// - Create an empty map
	emptyMap := make(map[string]int)
	fmt.Println(emptyMap)			// Output: map[]

	// - Create and initialize a map
	exampleMap := map[string]int{
		"A": 12,
		"B": 10,
		"C": 90,
	}
	fmt.Println(exampleMap)			// Output: map[A:12 B:10 C:90]



	// [ADD]
	// - Add a new key-value pair
	exampleMap["D"] = 40
	fmt.Println(exampleMap)			// Output: map[A:12 B:10 C:90 D:40]



	// [ACCESS]
	// - Access existing key
	fmt.Println(exampleMap["B"])	// Output: 10

	// - Try to access key that does not exist. This will NOT throw an error
	fmt.Println(exampleMap["X"])	// Output: 0

	// - Check if key exists and print its value if it does
	key,status := exampleMap["C"]
	if status == true {
		fmt.Println(key)			// Output: 90
	}



	// [DELETE]
	// - Delete existing key
	delete(exampleMap, "C")			 
	fmt.Println(exampleMap)			// Output: map[A:12 B:10 D:40]

	// - Try to delete key that does not exist. NOTHING happens
	delete(exampleMap, "X")


}