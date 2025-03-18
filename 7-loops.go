/*
	Loops
	- Note: an infinite loop will execute by default
*/

package main

import "fmt"

func main() {

	// Infinite loop: run forever; useful for certain web server tasks
	// Stop it with Ctrl + c. Comment it out to see the rest of the code
	for {
		fmt.Println("'[...] forever, Laura, forever.' - Michael Saylor")
	}
	

	// For loop
	for i := 1; i <= 10; i++ {
		fmt.Println("For loop:", i)

		if i == 5 {
			break	// Exit loop at 5
		}
	}


	// While loop, written with "for" keyword
	n := 1
	for n <= 10 {
		fmt.Println("While loop:", n)

		if n % 2 == 0 {
			n += 2
			continue	// Continue
		}
		n++
	}

}