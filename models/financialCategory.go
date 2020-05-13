package models

import "database/sql"

// PlaidCategory ...
type PlaidCategory struct {
	PlaidCategoryID string `json:"plaid_category_id"`
	Category1       string `json:"category1"`
	Category2       string `json:"category2"`
	Category3       string `json:"category3"`
}

// FinancialCategory ...
type FinancialCategory struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Group string `json:"grouping"`
}

// FinancialCategoryRepository interface
type FinancialCategoryRepository interface {
	GetFinancialCategories(tx *sql.Tx) ([]FinancialCategory, error)
	GetFinancialCategoryIDByPlaidID(tx *sql.Tx, plaidCategoryID string) (int64, error)
}
