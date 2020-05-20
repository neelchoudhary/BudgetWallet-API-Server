package postgresql

import (
	"database/sql"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
)

type financialTransactionRepository struct {
	db *sql.DB
}

// NewFinancialTransactionRepository returns a new instance of a postgresql financial transaction repository.
func NewFinancialTransactionRepository(db *sql.DB) models.FinancialTransactionRepository {
	return &financialTransactionRepository{db: db}
}

// AddTransaction add given transaction to the DB
func (r *financialTransactionRepository) AddTransaction(tx *sql.Tx, transaction *models.FinancialTransaction) error {
	var transactionID int64
	err := tx.QueryRow("INSERT INTO transactions (USER_ID, ITEM_ID, ACCOUNT_ID, CATEGORY_ID, PLAID_CATEGORY_ID, PLAID_TRANSACTION_ID, NAME, AMOUNT, DATE, PENDING, PLAID_ACCOUNT_ID) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id;",
		transaction.UserID, transaction.ItemID, transaction.AccountID, transaction.CategoryID, transaction.PlaidCategoryID,
		transaction.PlaidTransactionID, transaction.Name, transaction.Amount, transaction.Date, transaction.Pending,
		transaction.PlaidAccountID).Scan(&transactionID)
	if err != nil {
		return err
	}
	transaction.ID = transactionID
	return nil
}

// UpdateTransaction update transaction with new values from given transaction in the DB
func (r *financialTransactionRepository) UpdateTransaction(tx *sql.Tx, userID int64, transactionID int64, transaction *models.FinancialTransaction) error {
	_, err := tx.Exec("UPDATE transactions SET CATEGORY_ID=$3, NAME=$4, AMOUNT=$5, DATE=$6, PENDING=$7 WHERE user_id=$1 AND transaction_id=$2",
		userID, transactionID, transaction.CategoryID, transaction.Name, transaction.Amount, transaction.Date, transaction.Pending)
	return err
}

// DoesTransactionExist get transaction by userID and transactionID from the DB
func (r *financialTransactionRepository) DoesTransactionExist(tx *sql.Tx, userID int64, plaidTransactionID string) (bool, error) {
	var exists bool
	err := tx.QueryRow("SELECT EXISTS(SELECT 1 FROM transactions WHERE user_id=$1 AND plaid_transaction_id=$2);", userID, plaidTransactionID).Scan(&exists)
	return exists, err
}

// GetTransactionByID get transaction by userID and transactionID from the DB
func (r *financialTransactionRepository) GetTransactionByID(tx *sql.Tx, userID int64, transactionID int64) (*models.FinancialTransaction, error) {
	transaction := models.FinancialTransaction{}
	err := tx.QueryRow("SELECT * FROM transactions WHERE user_id=$1 AND id=$2;", userID, transactionID).Scan(
		&transaction.ID, &transaction.UserID, &transaction.ItemID, &transaction.AccountID,
		&transaction.CategoryID, &transaction.DailyAccountSnapshotID, &transaction.MonthlyAccountSnapshotID,
		&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name,
		&transaction.Amount, &transaction.Date, &transaction.Pending, &transaction.PlaidAccountID)
	return &transaction, err
}

// GetTransactionByPlaidID get transaction by userID and plaidTransactionID from the DB
func (r *financialTransactionRepository) GetTransactionByPlaidID(tx *sql.Tx, userID int64, plaidTransactionID string) (*models.FinancialTransaction, error) {
	transaction := models.FinancialTransaction{}
	err := tx.QueryRow("SELECT * FROM transactions WHERE user_id=$1 AND plaid_transaction_id=$2;", userID, plaidTransactionID).Scan(
		&transaction.ID, &transaction.UserID, &transaction.ItemID, &transaction.AccountID,
		&transaction.CategoryID, &transaction.DailyAccountSnapshotID, &transaction.MonthlyAccountSnapshotID,
		&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name,
		&transaction.Amount, &transaction.Date, &transaction.Pending, &transaction.PlaidAccountID)
	return &transaction, err
}

// GetAccountTransactions get transactions by userID and accountID from the DB
func (r *financialTransactionRepository) GetAccountTransactions(tx *sql.Tx, userID int64, accountID int64) ([]models.FinancialTransaction, error) {
	rows, err := tx.Query("SELECT * FROM transactions WHERE user_id=$1 AND account_id=$2 ORDER BY date DESC;", userID, accountID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := make([]models.FinancialTransaction, 0)
	for rows.Next() {
		transaction := models.FinancialTransaction{}
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.ItemID, &transaction.AccountID,
			&transaction.CategoryID, &transaction.DailyAccountSnapshotID, &transaction.MonthlyAccountSnapshotID,
			&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name,
			&transaction.Amount, &transaction.Date, &transaction.Pending, &transaction.PlaidAccountID)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// GetItemTransactions get transactions by userID and itemID from the DB
func (r *financialTransactionRepository) GetItemTransactions(tx *sql.Tx, userID int64, itemID int64) ([]models.FinancialTransaction, error) {
	rows, err := tx.Query("SELECT * FROM transactions WHERE user_id=$1 AND item_id=$2 ORDER BY date DESC;", userID, itemID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := make([]models.FinancialTransaction, 0)
	for rows.Next() {
		transaction := models.FinancialTransaction{}
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.ItemID, &transaction.AccountID,
			&transaction.CategoryID, &transaction.DailyAccountSnapshotID, &transaction.MonthlyAccountSnapshotID,
			&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name,
			&transaction.Amount, &transaction.Date, &transaction.Pending, &transaction.PlaidAccountID)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// GetUserTransactions get transactions by userID from the DB
func (r *financialTransactionRepository) GetUserTransactions(tx *sql.Tx, userID int64) ([]models.FinancialTransaction, error) {
	rows, err := tx.Query("SELECT * FROM transactions WHERE user_id=$1 ORDER BY date DESC;", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	transactions := make([]models.FinancialTransaction, 0)
	for rows.Next() {
		transaction := models.FinancialTransaction{}
		err := rows.Scan(
			&transaction.ID, &transaction.UserID, &transaction.ItemID, &transaction.AccountID,
			&transaction.CategoryID, &transaction.DailyAccountSnapshotID, &transaction.MonthlyAccountSnapshotID,
			&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name,
			&transaction.Amount, &transaction.Date, &transaction.Pending, &transaction.PlaidAccountID)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// RemoveItemTransactions remove all transactions for the given userID and itemID from the DB
func (r *financialTransactionRepository) RemoveItemTransactions(tx *sql.Tx, userID int64, itemID int64) error {
	_, err := tx.Exec("DELETE FROM transactions WHERE user_id=$1 AND item_id=$2;", userID, itemID)
	return err
}

// RemoveUserTransactions remove all transactions for the given userID from the DB
func (r *financialTransactionRepository) RemoveUserTransactions(tx *sql.Tx, userID int64) error {
	_, err := tx.Exec("DELETE FROM transactions WHERE user_id=$1;", userID)
	return err
}

// RemoveTransactionByID remove a transaction for the given userID and transactionID from the DB
func (r *financialTransactionRepository) RemoveTransactionByID(tx *sql.Tx, userID int64, transactionID int64) error {
	_, err := tx.Exec("DELETE FROM transactions WHERE user_id=$1 AND id=$2;", userID, transactionID)
	return err
}
