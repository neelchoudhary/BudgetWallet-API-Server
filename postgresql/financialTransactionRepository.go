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
func (r *financialTransactionRepository) AddTransaction(transaction *models.FinancialTransaction) error {
	var transactionID int64
	err := r.db.QueryRow("INSERT INTO transactions (USER_ID, ITEM_ID, ACCOUNT_ID, CATEGORY_ID, PLAID_CATEGORY_ID, PLAID_TRANSACTION_ID, NAME, AMOUNT, DATE, PENDING) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id;",
		transaction.UserID, transaction.ItemID, transaction.AccountID, transaction.CategoryID, transaction.PlaidCategoryID,
		transaction.PlaidTransactionID, transaction.Name, transaction.Amount, transaction.Date, transaction.Pending).Scan(&transactionID)
	if err != nil {
		return err
	}
	transaction.ID = transactionID
	return nil
}

// UpdateTransaction update transaction with new values from given transaction in the DB
func (r *financialTransactionRepository) UpdateTransaction(userID int64, transactionID int64, transaction *models.FinancialTransaction) error {
	_, err := r.db.Exec("UPDATE transactions SET CATEGORY_ID=$3, NAME=$4, AMOUNT=$5, DATE=$6, PENDING=$7 WHERE user_id=$1 AND transaction_id=$2",
		userID, transactionID, transaction.CategoryID, transaction.Name, transaction.Amount, transaction.Date, transaction.Pending)
	return err
}

// GetTransactionByID get transaction by userID and transactionID from the DB
func (r *financialTransactionRepository) GetTransactionByID(userID int64, transactionID int64) (*models.FinancialTransaction, error) {
	transaction := models.FinancialTransaction{}
	err := r.db.QueryRow("SELECT * FROM transactions WHERE user_id=$1 AND id=$2;", userID, transactionID).Scan(
		&transaction.ID, &transaction.UserID, &transaction.ItemID, &transaction.AccountID,
		&transaction.CategoryID, &transaction.DailyAccountSnapshotID, &transaction.MonthlyAccountSnapshotID,
		&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name, &transaction.Amount,
		&transaction.Date, &transaction.Pending)
	return &transaction, err
}

// GetTransactionByPlaidID get transaction by userID and plaidTransactionID from the DB
func (r *financialTransactionRepository) GetTransactionByPlaidID(userID int64, plaidTransactionID string) (*models.FinancialTransaction, error) {
	transaction := models.FinancialTransaction{}
	err := r.db.QueryRow("SELECT * FROM transactions WHERE user_id=$1 AND plaid_transaction_id=$2;", userID, plaidTransactionID).Scan(
		&transaction.ID, &transaction.UserID, &transaction.ItemID, &transaction.AccountID,
		&transaction.CategoryID, &transaction.DailyAccountSnapshotID, &transaction.MonthlyAccountSnapshotID,
		&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name, &transaction.Amount,
		&transaction.Date, &transaction.Pending)
	return &transaction, err
}

// GetAccountTransactions get transactions by userID and accountID from the DB
func (r *financialTransactionRepository) GetAccountTransactions(userID int64, accountID int64) ([]models.FinancialTransaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions WHERE user_id=$1 AND account_id=$2;", userID, accountID)
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
			&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name, &transaction.Amount,
			&transaction.Date, &transaction.Pending)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// GetItemTransactions get transactions by userID and itemID from the DB
func (r *financialTransactionRepository) GetItemTransactions(userID int64, itemID int64) ([]models.FinancialTransaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions WHERE user_id=$1 AND item_id=$2;", userID, itemID)
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
			&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name, &transaction.Amount,
			&transaction.Date, &transaction.Pending)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// GetUserTransactions get transactions by userID from the DB
func (r *financialTransactionRepository) GetUserTransactions(userID int64) ([]models.FinancialTransaction, error) {
	rows, err := r.db.Query("SELECT * FROM transactions WHERE user_id=$1;", userID)
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
			&transaction.PlaidCategoryID, &transaction.PlaidTransactionID, &transaction.Name, &transaction.Amount,
			&transaction.Date, &transaction.Pending)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// RemoveItemTransactions remove all transactions for the given userID and itemID from the DB
func (r *financialTransactionRepository) RemoveItemTransactions(userID int64, itemID int64) error {
	_, err := r.db.Exec("DELETE FROM transactions WHERE user_id=$1 AND item_id=$2;", userID, itemID)
	return err
}

// RemoveUserTransactions remove all transactions for the given userID from the DB
func (r *financialTransactionRepository) RemoveUserTransactions(userID int64) error {
	_, err := r.db.Exec("DELETE FROM transactions WHERE user_id=$1;", userID)
	return err
}
