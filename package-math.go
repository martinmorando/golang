/*
    Package math
*/

package main

import (       // Import multiple libraries (parenthesis, not brackets!)
    "fmt"
    "math"
)

func main() {

    fmt.Println( math.Round(1.3) )  // Round to nearest integer

    fmt.Println( math.Sqrt(9) )     // Square root
    fmt.Println( math.Pow(2, 3) )   // Exponentiation
    
    fmt.Println( math.Pi )          // Pi
    fmt.Println( math.E )           // E

}