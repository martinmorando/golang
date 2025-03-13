package main

import(
    "fmt"
    "time"
)

func main() {

    /*
        Print current time
    */

    // [Example 1]: print current time in the default format
    fmt.Println( time.Now() )

    // [Example 2]: print current date as dd/mm/yyyy hh:mm:ss
    // Only accepts this specific date of 2006 when formatting 
    // Source: https://stackoverflow.com/questions/42217308/go-time-format-how-to-understand-meaning-of-2006-01-02-layout
    fmt.Println( time.Now().Format("02/01/2006 15:04:05") )



    /*
        Load time zone location
        - about time.LoadLocation: https://pkg.go.dev/time#LoadLocation
        - "error" is a predeclared type in Go, that's why some recommend using "err"
        - about nil and error: https://stackoverflow.com/a/35983594    
    */

    // [Example 1]: load time zone location as NY, USA
    location, err := time.LoadLocation("America/New_York")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Loaded location:", location)


    // [Example 2]: set time zone as UTC and ignore error
    // "_" is a "blank identifier", used to inform the compiler to ignore the value intentionally
    location2, _ := time.LoadLocation("UTC")
        fmt.Println("Loaded location:", location2)


    /*
        Why might the time zone not be available? At least 2 reasons:
        - missing IANA Time Zone Database
        - invalid time zone name
        Source: https://cloudscopic.io/blog/timezonehandling/
    */

}