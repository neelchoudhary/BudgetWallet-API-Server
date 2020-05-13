package postgresql

import (
	"database/sql"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
)

// FinancialAccountRepository struct
type FinancialAccountRepository struct {
	db *sql.DB
}

// NewFinancialAccountRepository sets the data source (e.g database)
func NewFinancialAccountRepository(db *sql.DB) models.FinancialAccountRepository {
	return &FinancialAccountRepository{db: db}
}

// AddAccount add given account to the DB
func (r *FinancialAccountRepository) AddAccount(tx *sql.Tx, account *models.FinancialAccount) error {
	var accountID int64
	err := tx.QueryRow("INSERT INTO accounts (USER_ID, ITEM_ID, PLAID_ACCOUNT_ID, CURRENT_BALANCE, AVAILABLE_BALANCE, ACCOUNT_NAME, OFFICIAL_NAME, ACCOUNT_TYPE, ACCOUNT_SUBTYPE, ACCOUNT_MASK, SELECTED) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id",
		account.UserID, account.ItemID, account.PlaidAccountID, account.CurrentBalance, account.AvailableBalance, account.AccountName,
		account.OfficialName, account.AccountType, account.AccountSubType, account.AccountMask, account.Selected).Scan(&accountID)
	if err != nil {
		return err
	}
	account.ID = accountID
	return nil
}

// UpdateAccount update the given accountID's account with the new account in the DB
func (r *FinancialAccountRepository) UpdateAccount(tx *sql.Tx, userID int64, accountID int64, account *models.FinancialAccount) error {
	_, err := tx.Exec("UPDATE accounts SET CURRENT_BALANCE=$3, AVAILABLE_BALANCE=$4, ACCOUNT_NAME=$5, OFFICIAL_NAME=$6, ACCOUNT_TYPE=$7, ACCOUNT_SUBTYPE=$8, ACCOUNT_MASK=$9, SELECTED=$10 WHERE user_id=$1 AND id=$2",
		userID, accountID, account.CurrentBalance, account.AvailableBalance, account.AccountName, account.OfficialName, account.AccountType, account.AccountSubType, account.AccountMask, account.Selected)
	return err
}

// GetAccountByID get account by userID and accountID from the DB
func (r *FinancialAccountRepository) GetAccountByID(tx *sql.Tx, userID int64, accountID int64) (*models.FinancialAccount, error) {
	account := models.FinancialAccount{}
	err := tx.QueryRow("SELECT * FROM accounts WHERE user_id=$1 AND id=$2;", userID, accountID).Scan(
		&account.ID, &account.UserID, &account.ItemID, &account.PlaidAccountID, &account.CurrentBalance, &account.AvailableBalance,
		&account.AccountName, &account.OfficialName, &account.AccountType, &account.AccountSubType, &account.AccountMask, &account.Selected)
	return &account, err
}

// GetAccountByPlaidID get account by userID and plaidAccountID from the DB
func (r *FinancialAccountRepository) GetAccountByPlaidID(tx *sql.Tx, userID int64, plaidAccountID string) (*models.FinancialAccount, error) {
	account := models.FinancialAccount{}
	err := tx.QueryRow("SELECT * FROM accounts WHERE user_id=$1 AND plaid_account_id=$2;", userID, plaidAccountID).Scan(
		&account.ID, &account.UserID, &account.ItemID, &account.PlaidAccountID, &account.CurrentBalance, &account.AvailableBalance,
		&account.AccountName, &account.OfficialName, &account.AccountType, &account.AccountSubType, &account.AccountMask, &account.Selected)
	return &account, err
}

// GetItemAccounts get all accounts for the given userID and itemID from the DB
func (r *FinancialAccountRepository) GetItemAccounts(tx *sql.Tx, userID int64, itemID int64) ([]models.FinancialAccount, error) {
	rows, err := tx.Query("SELECT * FROM accounts WHERE user_id=$1 AND item_id=$2;", userID, itemID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := make([]models.FinancialAccount, 0)
	for rows.Next() {
		account := models.FinancialAccount{}
		err := rows.Scan(&account.ID, &account.UserID, &account.ItemID, &account.PlaidAccountID, &account.CurrentBalance, &account.AvailableBalance,
			&account.AccountName, &account.OfficialName, &account.AccountType, &account.AccountSubType, &account.AccountMask, &account.Selected)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// GetUserAccounts get all accounts for the given userID from the DB
func (r *FinancialAccountRepository) GetUserAccounts(tx *sql.Tx, userID int64) ([]models.FinancialAccount, error) {
	rows, err := tx.Query("SELECT * FROM accounts WHERE user_id=$1;", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	accounts := make([]models.FinancialAccount, 0)
	for rows.Next() {
		account := models.FinancialAccount{}
		err := rows.Scan(&account.ID, &account.UserID, &account.ItemID, &account.PlaidAccountID, &account.CurrentBalance, &account.AvailableBalance,
			&account.AccountName, &account.OfficialName, &account.AccountType, &account.AccountSubType, &account.AccountMask, &account.Selected)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

// RemoveItemAccounts remove all accounts for the given userID and itemID from the DB
func (r *FinancialAccountRepository) RemoveItemAccounts(tx *sql.Tx, userID int64, itemID int64) error {
	_, err := tx.Exec("DELETE FROM accounts WHERE user_id=$1 AND item_id=$2;", userID, itemID)
	return err
}

// RemoveUserAccounts remove all accounts for the given userID from the DB
func (r *FinancialAccountRepository) RemoveUserAccounts(tx *sql.Tx, userID int64) error {
	_, err := tx.Exec("DELETE FROM accounts WHERE user_id=$1;", userID)
	return err
}
