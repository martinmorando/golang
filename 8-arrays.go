/*
	Arrays
	- Collection of data elements of the same type
	- 0-indexed, fixed-size
	- Elements can be modified
*/

package main

import "fmt"

func main() {

	// [Create]
	// - Create an empty array of integers
	var numbersA [5]int       	
	fmt.Println( numbersA )		// [0 0 0 0 0]

	// - Create an array with integer elements
		// - Alternative A
		var numbersB = [3]int{4, 5, 8}
		fmt.Println( numbersB )		// [4 5 8]

		// - Alternative B
		numbersC := [3]int{4, 5, 8}
		fmt.Println( numbersC )		// [4 5 8]

	// - Create an array of strings
	var strA = [2]string{"A", "B"}
	fmt.Println( strA )				// [A B]


	// [Access]
	fmt.Println( numbersB[1] )		// 5


	// [Modify]
	numbersB[2] = 9
	fmt.Println( numbersB )			// [4 5 9]

}