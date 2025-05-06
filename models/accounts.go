package models

import (
	"SimpleBank/db"
	"context"
)

type Accounts struct {
	Owner    string `json:"owner" binding:"required"`
	Balance  *int   `json:"balance" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

type AddMoney struct {
	AccountID int    `json:"AccountID" binding:"required"`
	Amount    int    `json:"Amount" binding:"required"`
	Currency  string `json:"Currency" binding:"required"`
}

func (account Accounts) CreateAccount() error {
	query := `INSERT INTO accounts(owner, balance, currency)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	var newAccountID int
	err := db.Conn.QueryRow(context.Background(), query, account.Owner, account.Balance, account.Currency).Scan(&newAccountID)
	return err
}

func GetAccount(accountID int64) (*Accounts, error) {
	query := `SELECT owner, balance, currency FROM accounts WHERE id = $1`

	var accountDetails Accounts
	err := db.Conn.QueryRow(context.Background(), query, accountID).Scan(&accountDetails.Owner, &accountDetails.Balance, &accountDetails.Currency)
	if err != nil {
		return nil, err
	}

	return &accountDetails, nil
}

func AddMoneyToAccount(accountID int, amount int) error {
	query := `UPDATE accounts
		SET balance = $1
		WHERE id = $2
		RETURNING id
	`
	var newAccountID int
	err := db.Conn.QueryRow(context.Background(), query, amount, accountID).Scan(&newAccountID)
	return err
}
