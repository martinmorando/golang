/*
    Variables
    
    - Unnamed values (literal values): 123, "string"
    - Named values: constants, variables
*/

package main

import "fmt"

func main() {

	// [Constants]. Must be declared inmediately; only use "", not ''
	const imNotAVariable = "I'm a constant"


	// [Variables]. By default strings are assigned "" and numbers 0
	// Alternative A: one per line with its explicit type		
	var name string            // Declare
	name = "Satoshi"           // ... and assign
	
	var age int = 21           // Declare and assign value in one line

	var isCypherpunk bool = true


	// Alternative B: declare several variables in one line, same type explicit 
	var name2, name3 string


	// Alternative C: we know the value. The compiler infers the type
	name4 := "name4value"
	x := 111111111


	// Alternative D: we know their value, different types. The compiler infers the type
	name5, x2 := "name5value", 222222222


	// If a variable is declared and not used, throws an error
	// Print different data types in one line (to avoid error)
	fmt.Println(age, isCypherpunk, name, name2, name3, name4, name5, x, x2)

}