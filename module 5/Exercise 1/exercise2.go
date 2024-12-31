package main

import (
	"errors"
	"fmt"
)

const (
	DepositOption          = 1
	WithdrawOption         = 2
	ViewBalanceOption      = 3
	ViewTransactionsOption = 4
	ExitOption             = 5
)

type Account struct {
	ID                 int
	Name               string
	Balance            float64
	TransactionHistory []string
}

var accounts []Account

// FindAccountByID finds an account by ID
func FindAccountByID(id int) *Account {
	for i := range accounts {
		if accounts[i].ID == id {
			return &accounts[i]
		}
	}
	return nil
}

// Deposit function to deposit money into an account
func Deposit(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("deposit amount must be greater than zero")
	}
	account.Balance += amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Deposited: $%.2f", amount))
	return nil
}

// Withdraw function to withdraw money from an account
func Withdraw(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}
	if account.Balance < amount {
		return errors.New("insufficient balance")
	}
	account.Balance -= amount
	account.TransactionHistory = append(account.TransactionHistory, fmt.Sprintf("Withdrew: $%.2f", amount))
	return nil
}

// ViewTransactionHistory function to display transaction history
func ViewTransactionHistory(account *Account) {
	if len(account.TransactionHistory) == 0 {
		fmt.Println("No transactions found.")
	} else {
		for _, transaction := range account.TransactionHistory {
			fmt.Println(transaction)
		}
	}
}

// ViewBalance function to display account balance
func ViewBalance(account *Account) {
	fmt.Printf("Account Balance: Rs.%.2f\n", account.Balance)
}

func DisplayMenu() {
	fmt.Println("\n1. Deposit")
	fmt.Println("2. Withdraw")
	fmt.Println("3. View Balance")
	fmt.Println("4. View Transaction History")
	fmt.Println("5. Exit")
	fmt.Print("Choose an option: ")
}

func main() {

	accounts = append(accounts, Account{ID: 1, Name: "Ram", Balance: 1000.00, TransactionHistory: []string{}})
	accounts = append(accounts, Account{ID: 2, Name: "Sita", Balance: 500.00, TransactionHistory: []string{}})

	var accountID int
	var account *Account
	var choice int

	fmt.Print("Enter account ID: ")
	_, err := fmt.Scanf("%d", &accountID)
	fmt.Scanln()
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}

	// Find the account
	account = FindAccountByID(accountID)
	if account == nil {
		fmt.Println("Account not found.")
		return
	}

	for {
		DisplayMenu()
		_, err := fmt.Scanf("%d", &choice)
		fmt.Scanln()
		if err != nil {
			fmt.Println("Invalid input. Please try again.")
			continue
		}

		switch choice {
		case DepositOption:
			var amount float64
			fmt.Print("Enter deposit amount: Rs.")
			_, err := fmt.Scanf("%f", &amount)
			fmt.Scanln()
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid amount.")
				continue
			}
			err = Deposit(account, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Deposit successful!")
			}

		case WithdrawOption:
			var amount float64
			fmt.Print("Enter withdrawal amount: Rs.")
			_, err := fmt.Scanf("%f", &amount)
			fmt.Scanln()
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid amount.")
				continue
			}
			err = Withdraw(account, amount)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Withdrawal successful!")
			}

		case ViewBalanceOption:
			ViewBalance(account)

		case ViewTransactionsOption:
			ViewTransactionHistory(account)

		case ExitOption:
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
