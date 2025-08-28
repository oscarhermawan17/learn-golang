package main

import (
	"fmt"
	"math"
)

func main() {
    const inflationRate = 2.5

	var investmentAmount float64
    var years, expectedReturnRate float64

    fmt.Print("Investment amount: ")
    fmt.Scan(&investmentAmount)

    fmt.Print("Expected Return Rate: ")
    fmt.Scan(&expectedReturnRate)

    fmt.Print("Years: ")
    fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/100), years)
    futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

    // fmt.Println("Future Value:", futureValue)
    // fmt.Println("Future Value (adjusted for inflation):", futureRealValue)


    formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
    formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)

    fmt.Printf("Future Value: %v\nFuture Value (adjusted for inflation): %v\n", futureValue, futureRealValue)
    fmt.Printf("Future Value: %f\nFuture Value (adjusted for inflation): %f\n", futureValue, futureRealValue)

    fmt.Print(formattedFV, formattedRFV)
}