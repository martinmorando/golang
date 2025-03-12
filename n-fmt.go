package main

import "fmt"

func main() {

	// fmt.Print(): prints arguments next to each other, no line break
	fmt.Print("A", "B", "C")            // ABC


	// fmt.Println(): prints arguments with a space between them and prints a line break at the end
	fmt.Println("A", "B", "C")          // A B C


	// fmt.Printf(): print with interpolated strings, no line break
	// A) Using %v: general
	x := "A"
	y := "B"
	fmt.Printf("%v is not %v.", x, y)    // A is not B
	

	// B) Using %d: integers
	x2 := 4
	fmt.Printf("I have %d apples.", x2)  // I have 4 apples


	// C) Using %f: limit precision on doubles
	x3 := 8.53
	fmt.Printf("You have $%f.", x3)      // You have $8.530000
	fmt.Printf("You have $%.2f.", x3)    // You have $8.53
	fmt.Printf("You have $%.1f", x3)     // You have $8.5


	// D) Using %T: print the type of the argument
	x4 := "Bitcoin"
	fmt.Printf("I'm a %T", x4)           // I'm a string

}