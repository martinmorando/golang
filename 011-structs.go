/*
    Structs
    - Structs are defined outside main().
    - Functions associated with structs are defined outside structs.
*/

package main

import "fmt"


// Define struct Company outside main() with capital letter
type Company struct {
    // Create fields
    name string		// Default value: ""
    year int    	// Default value: 0
}

// Define function printName() associated with struct Company
func (c Company) printName() {
    fmt.Println(c.name)
}

func main() {

    // [CREATE INSTANCE OF A STRUCT]
    // - Alternative A: pass values in order
    mm := Company{"M Inc.", 2025}            // Use {}, not ()!
    fmt.Println(mm)                          // Output: {M Inc. 2025}

    // - Alternative B: pass some values only
    edison := Company{name: "Edison Inc."}
    fmt.Println(edison)                      // Output: {Edison Inc. 0}

    // Call function
    edison.printName()                       // Output: Edison Inc.

}