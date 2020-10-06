package postgresql

import (
	"database/sql"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
)

type recurringTransactionRepository struct {
	db *sql.DB
}

// NewRecurringTransactionRepository returns a new instance of a postgresql financial transaction repository.
func NewRecurringTransactionRepository(db *sql.DB) models.RecurringTransactionRepository {
	return &recurringTransactionRepository{db: db}
}

// AddTransaction add given transaction to the DB
func (r *recurringTransactionRepository) AddRecurringTransaction(tx *sql.Tx, recurringTransaction *models.RecurringTransaction) error {
	var transactionID int64
	err := tx.QueryRow("INSERT INTO recurring_transactions (USER_ID, NAME, CATEGORY_ID, RECURRING_COUNT, RECURRING_SCORE, IS_RECURRING, RECURRING_PLAID_IDS) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;",
		recurringTransaction.UserID, recurringTransaction.Name, recurringTransaction.CategoryID, recurringTransaction.RecurringCount,
		recurringTransaction.RecurringScore, recurringTransaction.IsRecurring, recurringTransaction.RecurringPlaidIDs).Scan(&transactionID)
	if err != nil {
		return err
	}
	recurringTransaction.ID = transactionID
	return nil
}

// GetRecurringTransaction get transactions by userID from the DB
func (r *recurringTransactionRepository) GetRecurringTransactions(tx *sql.Tx, userID int64) ([]models.RecurringTransaction, error) {
	rows, err := tx.Query("SELECT * FROM recurring_transactions WHERE user_id=$1;", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	recurringTransactions := make([]models.RecurringTransaction, 0)
	for rows.Next() {
		recurringTransaction := models.RecurringTransaction{}
		err := rows.Scan(
			&recurringTransaction.ID, &recurringTransaction.UserID, &recurringTransaction.Name,
			&recurringTransaction.CategoryID, &recurringTransaction.RecurringCount, &recurringTransaction.RecurringScore,
			&recurringTransaction.IsRecurring, &recurringTransaction.RecurringPlaidIDs)
		if err != nil {
			return nil, err
		}
		recurringTransactions = append(recurringTransactions, recurringTransaction)
	}

	return recurringTransactions, nil
}

// RemoveUserTransactions remove all transactions for the given userID from the DB
func (r *recurringTransactionRepository) RemoveRecurringTransactions(tx *sql.Tx, userID int64) error {
	_, err := tx.Exec("DELETE FROM recurring_transactions WHERE user_id=$1;", userID)
	return err
}
