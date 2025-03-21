/*
	Slices
	- Think of an array but this one can change its size
*/

package main

import "fmt"

func main() {


	// [CREATE AN EMPTY SLICE]
	// - Alternative A
	var emptySliceA []int 		// Similar with "string", "bool"
	fmt.Println(emptySliceA)	// []

	// - Alternative B
	emptySliceB := []int{}
	fmt.Println(emptySliceB)	// []



	// [CREATE A SLICE WITH ELEMENTS]
	// - Alternative A
	var sliceA = []string{"A", "B", "C"}
	fmt.Println(sliceA)			// [A B C]

	// - Alternative B
	sliceB := []string{"A", "B", "C"}
	fmt.Println(sliceB)			// [A B C]

}