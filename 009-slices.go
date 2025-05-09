/*
    Slices
    - Think of an array but this one can change its size.
    - Slices are pointers, and are passed by reference to functions, 
      meaning modifications to the slice affects the original slice.
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
    fmt.Println(sliceA)			 // [A B C]

    // - Alternative B
    sliceB := []string{"A", "B", "C"}
    fmt.Println(sliceB)			 // [A B C]



    // [ADD ELEMENTS]
    // - Reassign to affect the slice
    sliceB = append(sliceB, "D")
    fmt.Println(sliceB)			 // [A B C D]

}