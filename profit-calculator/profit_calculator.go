package main

import (
	"fmt"

	"example.com/profit-calculator/utilities"
	"github.com/pallinder/go-randomdata"
)

// Go profit calcultor using approach elaborating using of functions and return types
func main() {

	revenue := scanInput("Enter revenue : ")
	expenses := scanInput("Enter expenses : ")
	taxRate := scanInput("Enter taxRate : ")

	//Read file from utility package
	detailsText, err := utilities.ReadDetailsFromFile("profit-details.txt")
	if err != nil {
		fmt.Print("Error While reading file but still continuing execution !!")
	}
	fmt.Printf("Details in the file : %s\n", detailsText)

	ebt, profit, ratio := calculateExpenses(revenue, expenses, taxRate)
	//Write to file via utility package
	utilities.WriteDetailsToFile(ebt, profit, ratio, "profit-details.txt")

	fmt.Printf("Faltu ka phone number %s\n", randomdata.PhoneNumber())

}

func calculateExpenses(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func scanInput(inputString string) (inputValue float64) {
	fmt.Print(inputString)
	fmt.Scan(&inputValue)
	return
}
