/*
    Functions
*/

package main

import "fmt"

// This function expects 1 integer argument and returns 1 integer
// The data type comes after the argument name, similar to variable declarations
func basicFunction(x int) int {
	return x*2
}

// This function expects 2 integers and returns 1 integer
func basicFunction2(a, b int) int {
	return a*b
}

// This function expects 1 integer and returns 1 string and 1 integer
func basicFunction3(a int) (string, int) {
	return "Number:", a
}

// Example using defer
// The statement is executed just before the function returns
func basicFunction4(x int) int {
	defer fmt.Println("This will get printed just before the return")
	fmt.Println("This will get printed before")
	return x
}

func main() {

	y := basicFunction(1)
	fmt.Println(y)

	z := basicFunction2(2, 3)
	fmt.Println(z)

	q, r := basicFunction3(4)
	fmt.Println(q, r)

	w := basicFunction4(5)
	fmt.Println(w)

}