/*
	Numbers	
*/

package main

import "fmt"

func main() {

	// There are many numeric data types to optimize memory usage based on specific needs
	// https://go.dev/ref/spec#Numeric_types
	var x1 uint8       // unsigned  8-bit integers (0 to 255)
	var x2 uint16      // unsigned 16-bit integers (0 to 65535)
	var x3 uint32      // unsigned 32-bit integers (0 to 4294967295)
	var x4 uint64      // unsigned 64-bit integers (0 to 18446744073709551615)

	var x5 int8        // signed  8-bit integers (-128 to 127)
	var x6 int16       // signed 16-bit integers (-32768 to 32767)
	var x7 int32       // signed 32-bit integers (-2147483648 to 2147483647)
	var x8 int64       // signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	var x9 float32     // IEEE 754 32-bit floating-point numbers
	var x10 float64    // IEEE 754 64-bit floating-point numbers

	var x11 complex64  // complex numbers with float32 real and imaginary parts
	var x12 complex128 // complex numbers with float64 real and imaginary parts

	var x13 byte       // alias for uint8
	var x14 rune       // alias for int32

	fmt.Println(x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14)

}