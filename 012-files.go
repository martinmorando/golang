/*
    Files
    - os.Getwd(): get working directory.
    - os.Mkdir(dirName, permission bits): create a directory.
    - os.Chdir(pathNewDir): change the working directory.
    - os.Create(fileName): create a file.
    - [file].WriteString("some text"): writes the specified string to the file, where [file] 
      is a pointer to a file created by os.Create. This operation will overwrite any existing content in the file."
    - os.ReadFile(fileName): read a file.
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
    err = os.Mkdir("./music", 0777)
    if err != nil {
    	fmt.Println(err)
    	return
    }



    // [CHANGE CURRENT DIRECTORY]
    err = os.Chdir("./music")
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



    // [WRITE TO A FILE]: overwrites existing file (created in the previous step)
    numberBytesWritten, err := file.WriteString("All tyrants will be shot.\n")
    if err != nil {
    	fmt.Println(err)
    	return
    } else {
    	fmt.Println(numberBytesWritten)
    }



    // [READ A FILE]
    content, err := os.ReadFile("cookies.txt")
    if err != nil {
    	fmt.Println(err)
    	return
    } else {
    	fmt.Println(string(content)) // Convert bytes to string
    }

}