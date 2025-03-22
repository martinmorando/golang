/*
	Conditionals
*/

package main

import "fmt"

func main() {

    var x int = 2

    // Conditional
    if (x > 2) {
    	fmt.Println("x > 2")
    } else if (x < 2) {
    	fmt.Println("x < 2")
    } else {
    	fmt.Println("x = 2")
    }


    // Parenthesis in conditional are optional in Go
    if x > 1 {
    	fmt.Println("x > 1")
    }


    // Equal
    if x == 2 {
    	fmt.Println("x is 2")
    }

    // Switch. Breaks are not required in Go
    switch x {
      case 0:
      	fmt.Println("x is 0")
      case 1:
      	fmt.Println("x is 1")
      case 2:
      	fmt.Println("x is 2")
      default:
      	fmt.Println("x is not 0, 1 nor 2")
    }

}