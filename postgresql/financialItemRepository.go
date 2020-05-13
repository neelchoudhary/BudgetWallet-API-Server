package postgresql

import (
	"database/sql"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
)

// FinancialItemRepository struct
type FinancialItemRepository struct {
	db *sql.DB
}

// NewFinancialItemRepository sets the db and plaid client
func NewFinancialItemRepository(db *sql.DB) models.FinancialItemRepository {
	return &FinancialItemRepository{db: db}
}

// AddItem add given item to the DB
func (r *FinancialItemRepository) AddItem(tx *sql.Tx, item *models.FinancialItem) error {
	var itemID int64
	err := tx.QueryRow("INSERT INTO items (PLAID_ITEM_ID, PLAID_ACCESS_TOKEN, USER_ID, PLAID_INSTITUTION_ID, INSTITUTION_NAME, INSTITUTION_COLOR, INSTITUTION_LOGO, ERROR_CODE, ERROR_DEV_MSG, ERROR_USER_MSG) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id",
		item.PlaidItemID, item.PlaidAccessToken, item.UserID, item.PlaidInstitutionID,
		item.InstitutionName, item.InstitutionColor, item.InstitutionLogo,
		item.ErrorCode, item.ErrorDevMessage, item.ErrorUserMessage).Scan(&itemID)
	if err != nil {
		return err
	}
	item.ID = itemID
	return nil
}

// UpdateItem update the given itemID's item with the new item in the DB
func (r *FinancialItemRepository) UpdateItem(tx *sql.Tx, userID int64, itemID int64, item *models.FinancialItem) error {
	_, err := tx.Exec("UPDATE items SET PLAID_ITEM_ID=$3, PLAID_ACCESS_TOKEN=$4, PLAID_INSTITUTION_ID=$5, INSTITUTION_NAME=$6, INSTITUTION_COLOR=$7, INSTITUTION_LOGO=$8, ERROR_CODE=$9, ERROR_DEV_MSG=$10, ERROR_USER_MSG=$11 WHERE id=$1 AND user_id=$2",
		itemID, userID, item.PlaidItemID, item.PlaidAccessToken, item.PlaidInstitutionID, item.InstitutionName,
		item.InstitutionColor, item.InstitutionLogo, item.ErrorCode, item.ErrorDevMessage, item.ErrorUserMessage)
	return err
}

// GetItemByID get item by itemID and userID from the DB
func (r *FinancialItemRepository) GetItemByID(tx *sql.Tx, userID int64, itemID int64) (*models.FinancialItem, error) {
	item := models.FinancialItem{}
	err := tx.QueryRow("SELECT * FROM items WHERE user_id=$1 AND id=$2", userID, itemID).Scan(
		&item.ID, &item.PlaidItemID, &item.PlaidAccessToken, &item.UserID, &item.PlaidInstitutionID,
		&item.InstitutionName, &item.InstitutionColor, &item.InstitutionLogo,
		&item.ErrorCode, &item.ErrorDevMessage, &item.ErrorUserMessage)
	return &item, err
}

// GetItemByPlaidID get item by plaidItemID and userID from the DB
func (r *FinancialItemRepository) GetItemByPlaidID(tx *sql.Tx, userID int64, plaidItemID string) (*models.FinancialItem, error) {
	item := models.FinancialItem{}
	err := tx.QueryRow("SELECT * FROM items WHERE user_id=$1 AND plaid_item_id=$2", userID, plaidItemID).Scan(
		&item.ID, &item.PlaidItemID, &item.PlaidAccessToken, &item.UserID, &item.PlaidInstitutionID,
		&item.InstitutionName, &item.InstitutionColor, &item.InstitutionLogo,
		&item.ErrorCode, &item.ErrorDevMessage, &item.ErrorUserMessage)
	return &item, err
}

// GetUserItems get all items for the given userID from the DB
func (r *FinancialItemRepository) GetUserItems(tx *sql.Tx, userID int64) ([]models.FinancialItem, error) {
	rows, err := tx.Query("SELECT * FROM items WHERE user_id=$1;", userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]models.FinancialItem, 0)
	for rows.Next() {
		item := models.FinancialItem{}
		err := rows.Scan(&item.ID,
			&item.PlaidItemID, &item.PlaidAccessToken, &item.UserID, &item.PlaidInstitutionID,
			&item.InstitutionName, &item.InstitutionColor, &item.InstitutionLogo,
			&item.ErrorCode, &item.ErrorDevMessage, &item.ErrorUserMessage)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// RemoveItem remove item with the given userID and itemID from the DB
func (r *FinancialItemRepository) RemoveItem(tx *sql.Tx, userID int64, itemID int64) error {
	_, err := tx.Exec("DELETE FROM items WHERE user_id=$1 AND id=$2;", userID, itemID)
	return err
}

// RemoveUserItems remove all items with the given userID from the DB
func (r *FinancialItemRepository) RemoveUserItems(tx *sql.Tx, userID int64) error {
	_, err := tx.Exec("DELETE FROM items WHERE user_id=$1;", userID)
	return err
}
