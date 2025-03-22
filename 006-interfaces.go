/*
    Interfaces
    - Note: whitespace indentation between interfaces and structs is not required, 
      maybe even not suggested (haven't seen others use it). It helps me understand 
      the logical relationship between elements
*/

package main

import (
    "fmt"
    "math" // To round the prices
)

// Constant used when converting USD to BTC (it's an example; the focus is on interfaces)
const bitcoinExchangeRate = 82971.22

// Define an interface named "Building"
// Any struct that uses the methods inside is implementing it implicitly
type Building interface {
    calculatePrice() float64
}

    // Hotel struct for storing Hotel data
    type Hotel struct {
        price float64 // Store the price (USD)
    }

        // Hotel struct has a method called "calculatePrice()"
        // that converts price in USD to BTC
        func (h Hotel) calculatePrice() float64 {

            // Convert USD to BTC
            bitcoinPrice := h.price / bitcoinExchangeRate

            // Round to 8 decimal places
            bitcoinPrice = math.Round(bitcoinPrice * 1e8) / 1e8

            return bitcoinPrice
        }

    // House struct for storing House data
    type House struct {
        squareMeters int
        pricePerSquareMeter float64
    }

        // House struct has a method called "calculatePrice()"
        // Notice the price of the house is calculated different from the price of hotel
        func (h House) calculatePrice() float64 {

            // Calculate price of the house 
            fiatPrice := float64(h.squareMeters) * h.pricePerSquareMeter

            // Convert USD to BTC            
            bitcoinPrice := fiatPrice / bitcoinExchangeRate

            // Round to 8 decimal places
            bitcoinPrice = math.Round(bitcoinPrice * 1e8) / 1e8
            return bitcoinPrice
        }


    // When this method is called on any building, it uses the 
    // method associated to the respective struct
    func printBuildingPrice (b Building) {
        fmt.Println( "The price is", b.calculatePrice(), "BTC" )
    }



func main() {

    // Declaring a value of the interface type...
    var Bylton Building
    // ...and assign a concrete building (a hotel)
    Bylton = Hotel{ price: 90000000 }

    // Use the House interface and assign a concrete building (a house)
    StevensHouse := House{ squareMeters: 545, pricePerSquareMeter: 2100 }

    printBuildingPrice(Bylton)
    printBuildingPrice(StevensHouse)

}