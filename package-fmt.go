/*
	Package fmt
*/

package main

import "fmt"

func main() {

    /*
        [Part A]: Printing
            A) fmt.Print()
            B) fmt.Println()
            C) fmt.Printf()
    */

    // A) fmt.Print(): prints arguments next to each other, no line break
    fmt.Print("A", "B", "C")            // ABC


    // B) fmt.Println(): prints arguments with a space between them and prints a line break at the end
    fmt.Println("A", "B", "C")          // A B C


    // C) fmt.Printf(): print with interpolated strings, no line break
        // C.1) Using %v: general
        x := "A"
        y := "B"
        fmt.Printf("%v is not %v.", x, y)    // A is not B

        // C.2) Using %d: integers
        x2 := 4
        fmt.Printf("I have %d apples.", x2)  // I have 4 apples

        // C.3) Using %f: limit precision on doubles
        x3 := 8.53
        fmt.Printf("You have $%f.", x3)      // You have $8.530000
        fmt.Printf("You have $%.2f.", x3)    // You have $8.53
        fmt.Printf("You have $%.1f", x3)     // You have $8.5

        // C.4) Using %T: print the type of the argument
        x4 := "Bitcoin"
        fmt.Printf("I'm a %T", x4)           // I'm a string



    /*
        [Part B]: Formatting (similar to respetice fmt.Print but without printing)
           A) fmt.Sprint
           B) fmt.Sprintln
           C) fmt.Sprintf
    */

    // A) fmt.Sprint: returns arguments next to each other, no spaces, no break line
    var nChocolates int = 2
    var messageA string
    messageA = fmt.Sprint("I have ", nChocolates, " chocolates.")
    fmt.Print(messageA)

    // B) fmt.Sprintln: returns spaces between arguments and a break line
    var nBeers int = 3
    var messageB string
    messageB = fmt.Sprintln("I have", nBeers, "beers.")
    fmt.Print(messageB)

    // C) fmt.Sprintf: returns with interpolated strings, no break line
    var messageC string
    passThis := "Jesus"
    messageC = fmt.Sprintf("My name is %v", passThis)
    fmt.Print(messageC)

}