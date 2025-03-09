package main

import(
	"fmt"
	"time"
)

func main() {

	fmt.Println( time.Now() )

	// Print current date as dd/mm/yyyy hh:mm:ss
	// Only accepts this specific date of 2006 when formatting: https://stackoverflow.com/questions/42217308/go-time-format-how-to-understand-meaning-of-2006-01-02-layout
	fmt.Println( time.Now().Format("02/01/2006 15:04:05") )
	
}