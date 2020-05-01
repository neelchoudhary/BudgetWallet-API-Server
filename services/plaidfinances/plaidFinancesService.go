package plaidfinances

import (
	context "context"

	"github.com/plaid/plaid-go/plaid"

	"github.com/neelchoudhary/budgetmanagergrpc/models"
	shared "github.com/neelchoudhary/budgetmanagergrpc/services/shared"
)

// Service PlaidFinancesService struct
type Service struct {
	financialItemRepo        models.FinancialItemRepository
	financialAccountRepo     models.FinancialAccountRepository
	financialTransactionRepo models.FinancialTransactionRepository
	plaidClient              *plaid.Client
}

// NewPlaidFinancesServer contructor to assign repo
func NewPlaidFinancesServer(itemRepo *models.FinancialItemRepository, accountRepo *models.FinancialAccountRepository, transactionRepo *models.FinancialTransactionRepository, plaidClient *plaid.Client) PlaidFinancesServiceServer {
	return &Service{financialItemRepo: *itemRepo, financialAccountRepo: *accountRepo, financialTransactionRepo: *transactionRepo, plaidClient: plaidClient}
}

// LinkFinancialInstitution link a new financial institution from Plaid and add item and accounts to DB
func (s *Service) LinkFinancialInstitution(ctx context.Context, req *LinkFinancialInstitutionRequest) (*LinkFinancialInstitutionResponse, error) {
	// Surround everythhinig in db commit
	item, err := models.NewFinancialItemFromPlaid(req.GetUserId(), req.GetPublicToken(), req.GetPlaidInstitutionId(), s.plaidClient)
	if err != nil {
		return nil, err
	}
	err = s.financialItemRepo.AddItem(item)
	if err != nil {
		return nil, err
	}
	accounts, err := item.GetFinancialAccountsFromPlaid(req.GetUserId(), s.plaidClient)
	if err != nil {
		return nil, err
	}
	for _, account := range accounts {
		// Surround in db commit
		err := s.financialAccountRepo.AddAccount(&account)
		if err != nil {
			return nil, err
		}
	}

	var pbAccounts []*shared.FinancialAccount
	for _, account := range accounts {
		pbAccounts = append(pbAccounts, shared.DataToAccountPb(account))
	}

	res := &LinkFinancialInstitutionResponse{
		FinancialAccounts: pbAccounts,
	}
	return res, nil
}

// UpdateFinancialInstitution update financial institution (item) from Plaid in DB
func (s *Service) UpdateFinancialInstitution(ctx context.Context, req *UpdateFinancialInstitutionRequest) (*UpdateFinancialInstitutionResponse, error) {
	item, err := s.financialItemRepo.GetItemByID(req.GetUserId(), req.GetItemId())
	if err != nil {
		return nil, err
	}
	err = item.UpdateItemFromPlaid(s.plaidClient)
	if err != nil {
		return nil, err
	}
	err = s.financialItemRepo.UpdateItem(req.GetUserId(), req.GetItemId(), item)
	if err != nil {
		return nil, err
	}
	res := &UpdateFinancialInstitutionResponse{
		Success: true,
	}
	return res, nil
}

// UpdateFinancialAccounts update financial accounts from Plaid in DB
func (s *Service) UpdateFinancialAccounts(ctx context.Context, req *UpdateFinancialAccountsRequest) (*UpdateFinancialAccountsResponse, error) {
	item, err := s.financialItemRepo.GetItemByID(req.GetUserId(), req.GetItemId())
	if err != nil {
		return nil, err
	}
	plaidResponse, err := s.plaidClient.GetBalances(item.GetAccessToken())
	if err != nil {
		return nil, err
	}
	for _, plaidAccount := range plaidResponse.Accounts {
		account, err := s.financialAccountRepo.GetAccountByPlaidID(req.GetUserId(), plaidAccount.AccountID)
		if err != nil {
			return nil, err
		}
		account.UpdateAccountFromPlaid(&plaidAccount)
		err = s.financialAccountRepo.UpdateAccount(req.GetUserId(), account.GetAccountID(), account)
		if err != nil {
			return nil, err
		}
	}

	res := &UpdateFinancialAccountsResponse{
		Success: true,
	}
	return res, nil
}

// RemoveFinancialInstitution remove financial institution (item) from Plaid and the DB
func (s *Service) RemoveFinancialInstitution(ctx context.Context, req *RemoveFinancialInstitutionRequest) (*RemoveFinancialInstitutionResponse, error) {
	item, err := s.financialItemRepo.GetItemByID(req.GetUserId(), req.GetItemId())
	if err != nil {
		return nil, err
	}
	err = item.RemoveItemFromPlaid(s.plaidClient)
	if err != nil {
		return nil, err
	}
	err = s.financialAccountRepo.RemoveItemAccounts(req.GetUserId(), req.GetItemId())
	if err != nil {
		return nil, err
	}
	err = s.financialItemRepo.RemoveItem(req.GetUserId(), req.GetItemId())
	if err != nil {
		return nil, err
	}
	res := &RemoveFinancialInstitutionResponse{
		Success: true,
	}
	return res, nil
}

// AddTransactions add transactions to the DB from plaid for a user's item
func (s *Service) AddTransactions(ctx context.Context, req *AddTransactionsRequest) (*AddTransactionsResponse, error) {
	return nil, nil
}

// AddPlaidCategories add all plaid categories to the DB
func (s *Service) AddPlaidCategories(context.Context, *Empty) (*AddPlaidCategoriesResponse, error) {
	return nil, nil
}

// RemovePlaidCategories remove all plaid categories from the DB
func (s *Service) RemovePlaidCategories(context.Context, *Empty) (*RemovePlaidCategoriesResponse, error) {
	return nil, nil
}

// Min Returns the min of two ints
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// func (s *Service) addTransactionsHelper(userID int64, itemID int64, accessToken string, offset int, count int) (int, error) {
// 	// Get transactions from Plaid (500 max at a time) with offset pagination
// 	res, err := s.plaidClient.GetTransactionsWithOptions(accessToken, plaid.GetTransactionsOptions{
// 		EndDate:   time.Now().Local().Format("2006-01-02"),
// 		StartDate: "2015-01-01",
// 		Count:     Min(count, 500),
// 		Offset:    offset,
// 	})
// 	if err != nil {
// 		return 0, err
// 	}

// 	uniqueTransactionsAdded := 0

// 	// Write new transactions to db
// 	for _, plaidTransaction := range res.Transactions {
// 		// Convert PlaidAccountID to AccountID
// 		account, err := s.financialAccountRepo.GetAccountByPlaidID(userID, plaidTransaction.AccountID)
// 		if err != nil {
// 			return 0, err
// 		}

// 		// Convert PlaidCategoryID to CategoryID
// 		categoryID, err := GetCategoryID(plaidTransaction.CategoryID)
// 		if err != nil {
// 			return 0, err
// 		}

// 		// Create new transaction to add to db
// 		transaction := models.NewFinancialTransactionFromPlaid(userID, itemID, account.GetAccountID(), categoryID, plaidTransaction)

// 		var exists bool
// 		row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM transactions2 WHERE plaid_transaction_id=$1);", transactions[i].ID)
// 		if err := row.Scan(&exists); err != nil {
// 			return 0, err
// 		} else if !exists {
// 			err := s.financialTransactionRepo.AddTransaction(transaction)
// 			if err != nil {
// 				return 0, err
// 			}
// 		}
// 	}

// 	if len(res.Transactions) < 500 {
// 		return 0, nil
// 	}

// 	return count - uniqueTransactionsAdded, nil
// }

// // RemoveTransactions Get transactions from Plaid given the pagination offset to be added to the db
// func RemoveTransactions(w http.ResponseWriter, r *http.Request, userID int64, itemID int64, plaidTransactionIDs []string) {
// 	ctx := context.Background()
// 	tx, err := db.BeginTx(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for i := 0; i < len(plaidTransactionIDs); i++ {
// 		_, err = db.Exec("DELETE FROM transactions2 WHERE user_id=$1 AND item_id=$2 AND plaid_transaction_id=$3;", userID, itemID, plaidTransactionIDs[i])
// 		if err != nil {
// 			tx.Rollback()
// 			http.Error(w, http.StatusText(500)+". Error deleting transaction. "+err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 	}

// 	err = tx.Commit()
// 	if err != nil {
// 		http.Error(w, http.StatusText(500)+". Failed to commit tx changes to db. "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// }
