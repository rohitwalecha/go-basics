package main

import "fmt"

// Control structures demonstration of user fetching account balance and depositing amount in the account

func main() {
	var accountBalance float64 = 10000
	var depositedAmount float64
	var choice int

	fmt.Println("Welcom to Go Bank.")
	fmt.Println("Press 1 -> Check Balance.")
	fmt.Println("Press 2 -> Deposit Amount")

	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Printf("Your account balance is %.2f\n", accountBalance)
	} else if choice == 2 {
		fmt.Print("Enter Amount to Deposit : ")
		fmt.Scan(&depositedAmount)
		accountBalance += depositedAmount
		fmt.Println("Thanks for making the transaction.")
		fmt.Printf("Your account balance is %.2f\n", accountBalance)
	}
}
