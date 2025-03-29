/*
    Files
    - os.Getwd(): get working directory
    - os.Create(): create a file
*/

package main

import (
    "os"    // Import required
    "fmt"
)

func main() {

    // [GET WORKING DIRECTORY]
    wd, err := os.Getwd()
    if err != nil {				// Handle possible errors
        fmt.Println(err)
        return
    }

    fmt.Println(wd)				// Print working directory



    // [CREATE A FILE]
    // - Create an empty file, or overwrite existing
    file, err := os.Create("cookies.txt")
    if err != nil {
    	fmt.Println(err)
    	return
    }
    defer file.Close()   		// Releases system resources

}