package main

import "fmt"

type InsufficientFundsError struct {
	Amount  float64
	Balance float64
	Message string
}

func (e *InsufficientFundsError) Error() string {
	return fmt.Sprintf("Attempted to withdraw %.2f with balance of %.2f: %s", e.Amount, e.Balance, e.Message)
}

func withdraw(balance, amount float64) error {
	if amount > balance {
		return &InsufficientFundsError{
			Amount:  amount,
			Balance: balance,
			Message: "insufficient funds",
		}
	}
	return nil
}

func main() {
	err := withdraw(100, 200)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
