/*
    Package math/rand
    - Import math/rand and time
*/

package main

import (
    "fmt"
    "time"
    "math/rand"	
)

func main() {

    // Seed the generator with the current time in nanoseconds
    rand.Seed(time.Now().UnixNano())

    // Generate a random integer between 0 and 9 included
    n := rand.Intn(10)

    fmt.Println(n)

}