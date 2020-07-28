package postgresql

import (
	"database/sql"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
)

// FinancialCategoryRepository struct
type FinancialCategoryRepository struct {
	db *sql.DB
}

// NewFinancialCategoryRepository sets the data source (e.g database)
func NewFinancialCategoryRepository(db *sql.DB) models.FinancialCategoryRepository {
	return &FinancialCategoryRepository{db: db}
}

// GetFinancialCategories gets financial categories
func (r *FinancialCategoryRepository) GetFinancialCategories(tx *sql.Tx) ([]models.FinancialCategory, error) {
	rows, err := tx.Query("SELECT * FROM categories;")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	categories := make([]models.FinancialCategory, 0)
	for rows.Next() {
		category := models.FinancialCategory{}
		err := rows.Scan(&category.ID, &category.Name, &category.Group)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// GetFinancialCategoryIDByPlaidID gets financial category id by plaid category id
func (r *FinancialCategoryRepository) GetFinancialCategoryIDByPlaidID(tx *sql.Tx, plaidCategoryID string) (int64, error) {

	if plaidCategoryID == "" {
		return 57, nil
	}

	var categoryID int64
	err := tx.QueryRow("SELECT category_id FROM category_mapping WHERE plaid_category_id=$1;",
		plaidCategoryID).Scan(&categoryID)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	} else {
		return categoryID, nil
	}

	err = tx.QueryRow("SELECT category_id FROM category_mapping WHERE plaid_category_id=$1;",
		plaidCategoryID[:len(plaidCategoryID)-3]).Scan(&categoryID)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	} else {
		return categoryID, nil
	}

	err = tx.QueryRow("SELECT category_id FROM category_mapping WHERE plaid_category_id=$1;",
		plaidCategoryID[:len(plaidCategoryID)-6]).Scan(&categoryID)
	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	} else {
		return categoryID, nil
	}
	// Return empty category (18)
	return 57, nil
}
