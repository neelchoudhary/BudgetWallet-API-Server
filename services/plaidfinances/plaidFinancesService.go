package plaidfinances

import (
	context "context"
	fmt "fmt"
	"time"

	"github.com/plaid/plaid-go/plaid"
	log "github.com/sirupsen/logrus"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
	"github.com/neelchoudhary/budgetwallet-api-server/postgresql"
	"github.com/neelchoudhary/budgetwallet-api-server/services/shared"
	"github.com/neelchoudhary/budgetwallet-api-server/utils"
)

var logger = func(methodName string, err error) *log.Entry {
	if err != nil {
		return log.WithFields(log.Fields{"service": "PlaidFinancesService", "method": methodName, "error": err.Error()})
	}
	return log.WithFields(log.Fields{"service": "PlaidFinancesService", "method": methodName})
}

// Service PlaidFinancesService struct
type Service struct {
	txRepo                   postgresql.TxRepository
	financialItemRepo        models.FinancialItemRepository
	financialAccountRepo     models.FinancialAccountRepository
	financialTransactionRepo models.FinancialTransactionRepository
	plaidClient              *plaid.Client
}

// NewPlaidFinancesServer contructor to assign repo
func NewPlaidFinancesServer(txRepo *postgresql.TxRepository, itemRepo *models.FinancialItemRepository, accountRepo *models.FinancialAccountRepository, transactionRepo *models.FinancialTransactionRepository, plaidClient *plaid.Client) PlaidFinancesServiceServer {
	return &Service{txRepo: *txRepo, financialItemRepo: *itemRepo, financialAccountRepo: *accountRepo, financialTransactionRepo: *transactionRepo, plaidClient: plaidClient}
}

// LinkFinancialInstitution link a new financial institution from Plaid and add item and accounts to DB
func (s *Service) LinkFinancialInstitution(ctx context.Context, req *LinkFinancialInstitutionRequest) (*LinkFinancialInstitutionResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("LinkFinancialInstitution", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	item, err := models.NewFinancialItemFromPlaid(req.GetUserId(), req.GetPublicToken(), req.GetPlaidInstitutionId(), s.plaidClient)
	if err != nil {
		logger("LinkFinancialInstitution", err).Error(fmt.Sprintf("Item call to NewFinancialItemFromPlaid failed"))
		return nil, utils.InternalServerError
	}
	err = s.financialItemRepo.AddItem(tx, item)
	if err != nil {
		logger("LinkFinancialInstitution", err).Error(fmt.Sprintf("Repo call to AddItem failed"))
		return nil, utils.InternalServerError
	}
	accounts, err := item.GetFinancialAccountsFromPlaid(req.GetUserId(), s.plaidClient)
	if err != nil {
		logger("LinkFinancialInstitution", err).Error(fmt.Sprintf("Item call to GetFinancialAccountsFromPlaid failed"))
		return nil, utils.InternalServerError
	}
	for _, account := range accounts {
		// Surround in db commit
		err := s.financialAccountRepo.AddAccount(tx, &account)
		if err != nil {
			logger("LinkFinancialInstitution", err).Error(fmt.Sprintf("Repo call to AddAccount failed"))
			return nil, utils.InternalServerError
		}
	}

	var pbAccounts []*shared.FinancialAccount
	for _, account := range accounts {
		pbAccounts = append(pbAccounts, shared.DataToAccountPb(account))
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("LinkFinancialInstitution", err).Error(utils.CommitTxErrorMsg)
		return nil, utils.InternalServerError
	}

	res := &LinkFinancialInstitutionResponse{
		FinancialAccounts: pbAccounts,
	}
	return res, nil
}

// UpdateFinancialInstitution update financial institution (item) from Plaid in DB
func (s *Service) UpdateFinancialInstitution(ctx context.Context, req *UpdateFinancialInstitutionRequest) (*UpdateFinancialInstitutionResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("UpdateFinancialInstitution", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	item, err := s.financialItemRepo.GetItemByID(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("UpdateFinancialInstitution", err).Error(fmt.Sprintf("Repo call to GetItemByID failed"))
		return nil, utils.InternalServerError
	}
	err = item.UpdateItemFromPlaid(s.plaidClient)
	if err != nil {
		logger("UpdateFinancialInstitution", err).Error(fmt.Sprintf("Item call to UpdateItemFromPlaid failed"))
		return nil, utils.InternalServerError
	}
	err = s.financialItemRepo.UpdateItem(tx, req.GetUserId(), req.GetItemId(), item)
	if err != nil {
		logger("UpdateFinancialInstitution", err).Error(fmt.Sprintf("Repo call to UpdateItem failed"))
		return nil, utils.InternalServerError
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("UpdateFinancialInstitution", err).Error(utils.CommitTxErrorMsg)
		return nil, utils.InternalServerError
	}

	res := &UpdateFinancialInstitutionResponse{
		Success: true,
	}
	return res, nil
}

// UpdateFinancialAccounts update financial accounts from Plaid in DB
func (s *Service) UpdateFinancialAccounts(ctx context.Context, req *UpdateFinancialAccountsRequest) (*UpdateFinancialAccountsResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("UpdateFinancialAccounts", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	item, err := s.financialItemRepo.GetItemByID(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("UpdateFinancialAccounts", err).Error(fmt.Sprintf("Repo call to GetItemByID failed"))
		return nil, utils.InternalServerError
	}
	plaidResponse, err := s.plaidClient.GetBalances(item.PlaidAccessToken)
	if err != nil {
		logger("UpdateFinancialAccounts", err).Error(fmt.Sprintf("Plaid call to GetBalances failed"))
		return nil, utils.InternalServerError
	}
	for _, plaidAccount := range plaidResponse.Accounts {
		account, err := s.financialAccountRepo.GetAccountByPlaidID(tx, req.GetUserId(), plaidAccount.AccountID)
		if err != nil {
			logger("UpdateFinancialAccounts", err).Error(fmt.Sprintf("Repo call to GetAccountByPlaidID failed"))
			return nil, utils.InternalServerError
		}
		account.UpdateAccountFromPlaid(&plaidAccount)
		err = s.financialAccountRepo.UpdateAccount(tx, req.GetUserId(), account.GetAccountID(), account)
		if err != nil {
			logger("UpdateFinancialAccounts", err).Error(fmt.Sprintf("Repo call to UpdateAccount failed"))
			return nil, utils.InternalServerError
		}
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("UpdateFinancialAccounts", err).Error(utils.CommitTxErrorMsg)
		return nil, utils.InternalServerError
	}

	res := &UpdateFinancialAccountsResponse{
		Success: true,
	}
	return res, nil
}

// RemoveFinancialInstitution remove financial institution (item) from Plaid and the DB
func (s *Service) RemoveFinancialInstitution(ctx context.Context, req *RemoveFinancialInstitutionRequest) (*RemoveFinancialInstitutionResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("RemoveFinancialInstitution", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	item, err := s.financialItemRepo.GetItemByID(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("RemoveFinancialInstitution", err).Error(fmt.Sprintf("Repo call to GetItemByID failed"))
		return nil, utils.InternalServerError
	}
	err = item.RemoveItemFromPlaid(s.plaidClient)
	if err != nil {
		logger("RemoveFinancialInstitution", err).Error(fmt.Sprintf("Item call to RemoveItemFromPlaid failed"))
		return nil, utils.InternalServerError
	}
	err = s.financialAccountRepo.RemoveItemAccounts(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("RemoveFinancialInstitution", err).Error(fmt.Sprintf("Repo call to RemoveItemAccounts failed"))
		return nil, utils.InternalServerError
	}
	err = s.financialItemRepo.RemoveItem(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("RemoveFinancialInstitution", err).Error(fmt.Sprintf("Repo call to RemoveItem failed"))
		return nil, utils.InternalServerError
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("RemoveFinancialInstitution", err).Error(utils.CommitTxErrorMsg)
		return nil, utils.InternalServerError
	}

	res := &RemoveFinancialInstitutionResponse{
		Success: true,
	}
	return res, nil
}

// AddHistoricalFinancialTransactions add all transactions since 2015 to the DB from plaid for a user's item
func (s *Service) AddHistoricalFinancialTransactions(ctx context.Context, req *AddHistoricalFinancialTransactionsRequest) (*AddHistoricalFinancialTransactionsResponse, error) {
	// Start new db transaction
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("AddHistoricalFinancialTransactions", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	// Get item by id
	item, err := s.financialItemRepo.GetItemByID(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("AddHistoricalFinancialTransactions", err).Error(fmt.Sprintf("Repo call to GetItemByID failed"))
		return nil, utils.InternalServerError
	}

	// Remove all transactions for this item
	err = s.financialTransactionRepo.RemoveItemTransactions(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("AddHistoricalFinancialTransactions", err).Error(fmt.Sprintf("Repo call to RemoveItemTransactions failed"))
		return nil, utils.InternalServerError
	}

	// Get financial transactions for the given item from Plaid datinig back to 2015.
	allTransactions, err := item.GetFinancialTransactionsFromPlaid("2015-01-01", s.plaidClient)
	if err != nil {
		logger("AddHistoricalFinancialTransactions", err).Error(fmt.Sprintf("Item call to GetFinancialTransactionsFromPlaid failed"))
		return nil, utils.InternalServerError
	}
	for _, transaction := range allTransactions {
		// Set the account id and category id for this transaction
		// ! todo category id!!
		account, err := s.financialAccountRepo.GetAccountByPlaidID(tx, req.GetUserId(), transaction.PlaidAccountID)
		if err != nil {
			logger("AddHistoricalFinancialTransactions", err).Error(fmt.Sprintf("Repo call to GetAccountByPlaidID failed"))
			return nil, utils.InternalServerError
		}
		transaction.AccountID = account.ID
		transaction.CategoryID = 1

		// Add transaction to db
		s.financialTransactionRepo.AddTransaction(tx, &transaction)
	}

	// If new transaction count differs from expected count, log
	if int64(len(allTransactions)) != req.GetExpectedCount() {
		logger("AddHistoricalFinancialTransactions", nil).Warn(fmt.Sprintf("Count mismatch. Expected: %d, Added: %d transactions",
			req.GetExpectedCount(), len(allTransactions)))
	}

	// Commit db changes
	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("AddHistoricalFinancialTransactions", err).Error(utils.CommitTxErrorMsg)
		return nil, utils.InternalServerError
	}

	// Return response
	res := &AddHistoricalFinancialTransactionsResponse{
		Success: true,
	}

	return res, nil
}

// AddFinancialTransactions add new transactions from the last 10 days to the DB from plaid for a user's item
func (s *Service) AddFinancialTransactions(ctx context.Context, req *AddFinancialTransactionsRequest) (*AddFinancialTransactionsResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("AddFinancialTransactions", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	item, err := s.financialItemRepo.GetItemByID(tx, req.GetUserId(), req.GetItemId())
	if err != nil {
		logger("AddFinancialTransactions", err).Error(fmt.Sprintf("Repo call to GetItemByID failed"))
		return nil, utils.InternalServerError
	}

	startDate := time.Now().Local().Add(time.Duration(240) * time.Hour * -1).Format("2006-01-02") // 10 days back
	allTransactions, err := item.GetFinancialTransactionsFromPlaid(startDate, s.plaidClient)
	if err != nil {
		logger("AddFinancialTransactions", err).Error(fmt.Sprintf("Item call to GetFinancialTransactionsFromPlaid failed"))
		return nil, utils.InternalServerError
	}
	filteredTransactions := models.FilterTransactions(allTransactions, func(t models.FinancialTransaction) bool {
		exists, err := s.financialTransactionRepo.DoesTransactionExist(tx, req.GetUserId(), t.PlaidTransactionID)
		if err != nil {
			logger("AddFinancialTransactions", err).Error(fmt.Sprintf("Repo call to DoesTransactionExist failed"))
			return false
		}
		return !exists
	})
	for _, transaction := range filteredTransactions {
		account, err := s.financialAccountRepo.GetAccountByPlaidID(tx, req.GetUserId(), transaction.PlaidAccountID)
		if err != nil {
			logger("AddFinancialTransactions", err).Error(fmt.Sprintf("Repo call to GetAccountByPlaidID failed"))
			return nil, utils.InternalServerError
		}
		transaction.AccountID = account.ID
		transaction.CategoryID = 2

		s.financialTransactionRepo.AddTransaction(tx, &transaction)
	}

	// != given transaction count (from req)
	if int64(len(filteredTransactions)) != req.GetExpectedCount() {
		logger("AddFinancialTransactions", nil).Warn(fmt.Sprintf("Count mismatch. Expected: %d, Added: %d transactions",
			req.GetExpectedCount(), len(filteredTransactions)))
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("AddFinancialTransactions", err).Error(utils.CommitTxErrorMsg)
		return nil, utils.InternalServerError
	}

	res := &AddFinancialTransactionsResponse{
		Success: true,
	}

	return res, nil
}

// AddPlaidCategories add all plaid categories to the DB
func (s *Service) AddPlaidCategories(context.Context, *Empty) (*AddPlaidCategoriesResponse, error) {
	return nil, nil
}

// RemovePlaidCategories remove all plaid categories from the DB
func (s *Service) RemovePlaidCategories(context.Context, *Empty) (*RemovePlaidCategoriesResponse, error) {
	return nil, nil
}

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
