/*
	1) Compile file into executable (create executable binary): `go build main.go` 
	2) Run binary: `./main`

	or

	1) Run with: `go run main.go` (useful for development, no executable binary is saved)
*/

package main // This file is an executable program, not a library

import "fmt" // Import the package that has the I/O functions. Comes from "format"? 

func main() { // Declare a function
	fmt.Println("Teacher, leave them kids alone")
}