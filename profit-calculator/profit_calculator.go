package main

import "fmt"

// Go profit calcultor using approach elaborating using of functions and return types
func main() {

	revenue := scanInput("Enter revenue : ")
	expenses := scanInput("Enter expenses : ")
	taxRate := scanInput("Enter taxRate : ")

	ebt, profit, ratio := calculateExpenses(revenue, expenses, taxRate)

	fmt.Printf("Earnings Before Tax : %.2f\n", ebt)
	fmt.Printf("Profit : %.2f\n", profit)
	fmt.Printf("Ratio EBT/Profit : %.2f\n", ratio)

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
