package main

import (
	"fmt"
)

func main() {

	var payment float64

	payment = float64(loanCalc())
	fmt.Println("total monthly payment is ", payment)

}

func loanCalc() float32 {

	var amount, interestRate, months, monthlyInterestRate, payment float64

	fmt.Println("to calculate the loan payment in month please start by entering the loan amount")
	fmt.Scanln(&amount)
	fmt.Println("inter the interest rate")
	fmt.Scanln(&interestRate)
	fmt.Println("enter the loan period in months")
	fmt.Scanln(&months)

	monthlyInterestRate = interestRate / months / 100.00

	payment = amount * (monthlyInterestRate + (1.0 / float64(months)))

	return float32(payment)

}
