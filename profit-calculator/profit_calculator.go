package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue : ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses : ")
	fmt.Scan(&expenses)

	fmt.Print("Enter taxRate : ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate)
	ratio := ebt / profit

	fmt.Printf("Earnings Before Tax : %.2f\n", ebt)
	fmt.Printf("Profit : %.2f\n", profit)
	fmt.Printf("Ratio EBT/Profit : %.2f\n", ratio)

}
