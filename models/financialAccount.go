package models

import (
	"github.com/plaid/plaid-go/plaid"
)

// FinancialAccount ...
type FinancialAccount struct {
	ID               int64   `json:"id"`
	UserID           int64   `json:"user_id"`
	ItemID           int64   `json:"item_id"`
	PlaidAccountID   string  `json:"account_id_plaid"`
	CurrentBalance   float64 `json:"current_balance"`
	AvailableBalance float64 `json:"available_balance"`
	AccountName      string  `json:"account_name"`
	OfficialName     string  `json:"official_name"`
	AccountType      string  `json:"account_type"`
	AccountSubType   string  `json:"account_subtype"`
	AccountMask      string  `json:"account_mask"`
	Selected         bool    `json:"selected"`
}

// NewFinancialAccountFromPlaid creates a new financial account from a plaid account
func NewFinancialAccountFromPlaid(userID int64, itemID int64, plaidAccount plaid.Account) FinancialAccount {
	account := FinancialAccount{}
	account.UserID = userID
	account.ItemID = itemID
	account.PlaidAccountID = plaidAccount.AccountID
	account.AccountName = plaidAccount.Name
	account.OfficialName = plaidAccount.OfficialName
	account.AccountType = plaidAccount.Type
	account.AccountSubType = plaidAccount.Subtype
	account.AccountMask = plaidAccount.Mask
	account.CurrentBalance = plaidAccount.Balances.Current
	account.AvailableBalance = plaidAccount.Balances.Available
	account.Selected = false
	return account
}

// UpdateAccountFromPlaid updates this financial account with new values from plaid account
func (a *FinancialAccount) UpdateAccountFromPlaid(plaidAccount *plaid.Account) {
	a.PlaidAccountID = plaidAccount.AccountID
	a.CurrentBalance = plaidAccount.Balances.Current
	a.AvailableBalance = plaidAccount.Balances.Available
	a.AccountName = plaidAccount.Name
	a.OfficialName = plaidAccount.OfficialName
	a.AccountType = plaidAccount.Type
	a.AccountSubType = plaidAccount.Subtype
	a.AccountMask = plaidAccount.Mask
}

// SetSelected sets the selected value for this account
func (a *FinancialAccount) SetSelected(selected bool) {
	a.Selected = selected
}

// GetAccountID get the account's ID field
func (a *FinancialAccount) GetAccountID() int64 {
	return a.ID
}

// FinancialAccountRepository interface
type FinancialAccountRepository interface {
	AddAccount(account *FinancialAccount) error
	UpdateAccount(userID int64, accountID int64, account *FinancialAccount) error
	GetAccountByID(userID int64, accountID int64) (*FinancialAccount, error)
	GetAccountByPlaidID(userID int64, plaidAccountID string) (*FinancialAccount, error)
	GetItemAccounts(userID int64, itemID int64) ([]FinancialAccount, error)
	GetUserAccounts(userID int64) ([]FinancialAccount, error)
	RemoveItemAccounts(userID int64, itemID int64) error
	RemoveUserAccounts(userID int64) error
}
