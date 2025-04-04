/*
    Files
    - os.Getwd(): get working directory
    - os.Mkdir(): create a directory
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
    if err != nil {				// If err is nil, no errors occurred
        fmt.Println(err)
        return
    }

    fmt.Println(wd)				// Print working directory



    // [CREATE A DIRECTORY]
    // - Create a directory in CWD and set permissions to 0777
    err = os.Mkdir("./newFolder", 0777)
    if err != nil {
    	fmt.Println(err)
    	return
    }



    // [CREATE A FILE]
    // - Create an empty file, or overwrite existing
    file, err := os.Create("cookies.txt")
    if err != nil {
    	fmt.Println(err)
    	return
    }
    defer file.Close()   		// Releases system resources

}