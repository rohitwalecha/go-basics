// This package has been created to demonstrate use of packages in go
package utilities

import (
	"fmt"
	"os"
)

func WriteDetailsToFile(ebt float64, profit float64, ratio float64, fileName string) {
	ebtText := fmt.Sprintf("Earnings Before Tax : %.2f\n", ebt)
	profitText := fmt.Sprintf("Profit : %.2f\n", profit)
	ratioText := fmt.Sprintf("Ratio EBT/Profit : %.2f\n", ratio)

	detailsText := ebtText + profitText + ratioText

	os.WriteFile(fileName, []byte(detailsText), 0644)
}

func ReadDetailsFromFile(fileName string) (string, error) {
	// This ReadFile function returns 0, error combination so that both can be handled this is go flex
	details, err := os.ReadFile(fileName)
	if err != nil {
		return "", err // Return empty string and err if err returned by ReadFile is not nil
	}
	return string(details), err
}
