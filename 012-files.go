/*
    Files
    - os.Getwd(): get working directory
*/

package main

import (
    "os"    // Import required
    "fmt"
)

func main() {

    // Get working directory
    wd, err := os.Getwd()

    // Handle possible errors
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(wd)

}