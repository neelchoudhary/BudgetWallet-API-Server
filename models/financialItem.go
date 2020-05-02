package models

import (
	"errors"
	"time"

	"github.com/plaid/plaid-go/plaid"
)

// FinancialItem ...
type FinancialItem struct {
	ID                 int64  `json:"id"`
	PlaidItemID        string `json:"plaid_item_id"`
	PlaidAccessToken   string `json:"plaid_access_token"`
	UserID             int64  `json:"user_id"`
	PlaidInstitutionID string `json:"plaid_institution_id"`
	InstitutionName    string `json:"institution_name"`
	InstitutionColor   string `json:"institution_color"`
	InstitutionLogo    string `json:"institution_logo"`
	ErrorCode          string `json:"error_message"`
	ErrorDevMessage    string `json:"error_dev_msg"`
	ErrorUserMessage   string `json:"error_user_msg"`
}

// NewFinancialItemFromPlaid ...
func NewFinancialItemFromPlaid(userID int64, publicToken string, institutionPlaidID string, plaidClient *plaid.Client) (*FinancialItem, error) {
	// Exchange public token
	exchangeResponse, err := plaidClient.ExchangePublicToken(publicToken)
	if err != nil {
		return nil, err
	}

	// Get institution details
	institutionResponse, err := plaidClient.GetInstitutionByIDWithOptions(institutionPlaidID, plaid.GetInstitutionByIDOptions{
		IncludeOptionalMetadata: true,
	})
	if err != nil {
		return nil, err
	}

	// Construct new item
	item := &FinancialItem{
		UserID:             userID,
		PlaidItemID:        exchangeResponse.ItemID,
		PlaidAccessToken:   exchangeResponse.AccessToken,
		PlaidInstitutionID: institutionResponse.Institution.ID,
		InstitutionName:    institutionResponse.Institution.Name,
		InstitutionColor:   institutionResponse.Institution.PrimaryColor,
		InstitutionLogo:    institutionResponse.Institution.Logo,
	}
	return item, nil
}

// GetFinancialAccountsFromPlaid ...
func (i *FinancialItem) GetFinancialAccountsFromPlaid(userID int64, plaidClient *plaid.Client) ([]FinancialAccount, error) {
	// Get all new plaid accounts from this item
	accountsResponse, err := plaidClient.GetAccounts(i.PlaidAccessToken)
	if err != nil {
		return nil, err
	}

	// Convert plaid accounts to financial accounts
	plaidAccounts := accountsResponse.Accounts
	accounts := make([]FinancialAccount, 0)
	for _, plaidAccount := range plaidAccounts {
		account := NewFinancialAccountFromPlaid(userID, i.ID, plaidAccount)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// GetFinancialTransactionsFromPlaid ...
func (i *FinancialItem) GetFinancialTransactionsFromPlaid(startDate string, plaidClient *plaid.Client) ([]FinancialTransaction, error) {
	paginate := 0
	shouldLoop := true
	transactions := make([]FinancialTransaction, 0)
	for shouldLoop {
		var err error
		shouldLoop, err = i.getFinancialTransactionsFromPlaidWithPagination(paginate, startDate, &transactions, plaidClient)
		if err != nil {
			return nil, err
		}
		// Delay to avoid rate limits
		if shouldLoop {
			paginate += 500
			time.Sleep(100 * time.Millisecond)
		}
	}
	return transactions, nil
}

// GetFinancialTransactionsFromPlaid ...
func (i *FinancialItem) getFinancialTransactionsFromPlaidWithPagination(offset int, startDate string, transactions *[]FinancialTransaction, plaidClient *plaid.Client) (bool, error) {
	// Get transactions from Plaid (500 max at a time) with offset pagination
	res, err := plaidClient.GetTransactionsWithOptions(i.PlaidAccessToken, plaid.GetTransactionsOptions{
		EndDate:   time.Now().Local().Format("2006-01-02"),
		StartDate: startDate,
		Count:     500,
		Offset:    offset,
	})

	if err != nil {
		return false, err
	}

	// Convert plaid transactions to financial transactions
	for _, plaidTransaction := range res.Transactions {
		// ! need to provide proper account & category id
		transaction := NewFinancialTransactionFromPlaid(i.UserID, i.ID, plaidTransaction)
		*transactions = append(*transactions, transaction)
	}

	// If there are less than 500 transactions, we have reached the end
	if len(res.Transactions) < 500 {
		return false, nil
	}

	// Otherwise, continue looping
	return true, nil
}

// UpdateItemFromPlaid update item from plaid in the DB
func (i *FinancialItem) UpdateItemFromPlaid(plaidClient *plaid.Client) error {
	itemResponse, err := plaidClient.GetItem(i.PlaidAccessToken)
	if err != nil {
		return err
	}
	institutionResponse, err := plaidClient.GetInstitutionByIDWithOptions(itemResponse.Item.InstitutionID, plaid.GetInstitutionByIDOptions{
		IncludeOptionalMetadata: true,
	})
	if err != nil {
		return err
	}
	institutionID := itemResponse.Item.InstitutionID
	itemError := itemResponse.Item.Error
	i.PlaidInstitutionID = institutionID
	i.InstitutionName = institutionResponse.Institution.Name
	i.InstitutionColor = institutionResponse.Institution.PrimaryColor
	i.InstitutionLogo = institutionResponse.Institution.Logo
	i.ErrorCode = itemError.ErrorCode
	i.ErrorDevMessage = itemError.ErrorMessage
	i.ErrorUserMessage = itemError.DisplayMessage
	return nil
}

// RemoveItemFromPlaid calls Plaid api to remove item
func (i *FinancialItem) RemoveItemFromPlaid(plaidClient *plaid.Client) error {
	res, err := plaidClient.RemoveItem(i.PlaidAccessToken)
	if err != nil {
		return err
	}
	if !res.Removed {
		return errors.New("Remove Item: Failed to remove item from plaid")
	}
	return nil
}

// // GetAccessToken get the access token field
// func (i *FinancialItem) GetAccessToken() string {
// 	return i.PlaidAccessToken
// }

// // GetItemID get the item's ID field
// func (i *FinancialItem) GetItemID() int64 {
// 	return i.ID
// }

// FinancialItemRepository interface
type FinancialItemRepository interface {
	AddItem(item *FinancialItem) error
	UpdateItem(userID int64, itemID int64, item *FinancialItem) error
	GetItemByID(userID int64, itemID int64) (*FinancialItem, error)
	GetItemByPlaidID(userID int64, plaidItemID string) (*FinancialItem, error)
	GetUserItems(userID int64) ([]FinancialItem, error)
	RemoveItem(userID int64, itemID int64) error
	RemoveUserItems(userID int64) error
}
