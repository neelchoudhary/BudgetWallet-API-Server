package dataprocessing

import (
	"bytes"
	context "context"
	"encoding/json"
	fmt "fmt"
	"io/ioutil"
	math "math"
	"net/http"
	"strings"
	"time"

	"github.com/neelchoudhary/budgetwallet-api-server/models"
	"github.com/neelchoudhary/budgetwallet-api-server/postgresql"
	"github.com/neelchoudhary/budgetwallet-api-server/utils"
	log "github.com/sirupsen/logrus"
)

var logger = func(methodName string, err error) *log.Entry {
	if err != nil {
		return log.WithFields(log.Fields{"service": "DataProcessingService", "method": methodName, "error": err.Error()})
	}
	return log.WithFields(log.Fields{"service": "DataProcessingService", "method": methodName})
}

// Service DataProcessingService struct
type Service struct {
	txRepo                   postgresql.TxRepository
	financialAccountRepo     models.FinancialAccountRepository
	financialTransactionRepo models.FinancialTransactionRepository
}

// NewDataProcessingServer contructor to assign repo
func NewDataProcessingServer(txRepo *postgresql.TxRepository, accountRepo *models.FinancialAccountRepository, transactionRepo *models.FinancialTransactionRepository) DataProcessingServiceServer {
	return &Service{txRepo: *txRepo, financialAccountRepo: *accountRepo, financialTransactionRepo: *transactionRepo}
}

// GetAccountDailySnapshots get daily snapshots for an account
func (s *Service) GetAccountDailySnapshots(ctx context.Context, req *GetAccountDailySnapshotsRequest) (*GetAccountDailySnapshotsResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("GetAccountDailySnapshots", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	userID, err := utils.GetUserIDMetadata(ctx)
	if err != nil {
		logger("GetAccountDailySnapshots", err).Error(fmt.Sprintf("GetUserIDMetadata failed"))
		return nil, utils.InternalServerError
	}

	transactions, err := s.financialTransactionRepo.GetAccountTransactions(tx, userID, req.GetAccountId())
	if err != nil {
		logger("GetAccountDailySnapshots", err).Error(fmt.Sprintf("Repo call to GetAccountTransactions failed"))
		return nil, utils.InternalServerError
	}

	if len(transactions) == 0 {
		res := &GetAccountDailySnapshotsResponse{
			AccountDailySnapshots: nil,
		}
		return res, nil
	}

	// Get oldest date
	oldestDate := transactions[len(transactions)-1].Date

	// Find current date
	currentDate := time.Now().Format("2006-01-02")

	// Get account to find current balance
	account, err := s.financialAccountRepo.GetAccountByID(tx, userID, req.GetAccountId())
	if err != nil {
		logger("GetAccountDailySnapshots", err).Error(fmt.Sprintf("Repo call to GetAccountByID failed"))
		return nil, utils.InternalServerError
	}
	availableBalance := account.AvailableBalance

	if account.AccountType == "credit" {
		availableBalance = account.CurrentBalance
	}

	accountDailySnapshots := make([]*AccountDailySnapshot, 0)

	// Loop through dates starting from current date to oldest date
	for date := currentDate; DateComparator(date, oldestDate); date = DateDecrementer(date) {
		// Get all transactions on that date => list of transactions on date
		transactionsOnDate := make([]models.FinancialTransaction, 0)
		for _, transaction := range transactions {
			if transaction.Date == date {
				transactionsOnDate = append(transactionsOnDate, transaction)
			}
		}

		// Loop through list of transactions on date and find balance, cash in, and cash out
		dailyCashOut := 0.0
		dailyCashIn := 0.0
		for _, transaction := range transactionsOnDate {
			if transaction.Amount > 0 {
				dailyCashOut += transaction.Amount
			} else {
				dailyCashIn += (transaction.Amount * -1)
			}
		}
		endBalance := availableBalance

		if account.AccountType == "credit" {
			availableBalance = availableBalance - dailyCashOut + dailyCashIn
		} else {
			availableBalance = availableBalance + dailyCashOut - dailyCashIn
		}

		// Insert new daily_account record with the above info for the date.
		accountDailySnapshot := &AccountDailySnapshot{
			Date:            date,
			StartDayBalance: math.Round(availableBalance*100) / 100,
			EndDayBalance:   math.Round(endBalance*100) / 100,
			CashOut:         math.Round(dailyCashOut*100) / 100,
			CashIn:          math.Round(dailyCashIn*100) / 100,
		}
		accountDailySnapshots = append(accountDailySnapshots, accountDailySnapshot)
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("GetAccountDailySnapshots", err).Error(utils.CommitTxErrorMsg)
		return nil, err
	}

	res := &GetAccountDailySnapshotsResponse{
		AccountDailySnapshots: accountDailySnapshots,
	}
	return res, nil
}

// GetAccountMonthlySnapshots get monthly snapshots for an account
func (s *Service) GetAccountMonthlySnapshots(ctx context.Context, req *GetAccountMonthlySnapshotsRequest) (*GetAccountMonthlySnapshotsResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("GetAccountMonthlySnapshots", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	userID, err := utils.GetUserIDMetadata(ctx)
	if err != nil {
		logger("GetAccountMonthlySnapshots", err).Error(fmt.Sprintf("GetUserIDMetadata failed"))
		return nil, utils.InternalServerError
	}

	transactions, err := s.financialTransactionRepo.GetAccountTransactions(tx, userID, req.GetAccountId())
	if err != nil {
		logger("GetAccountMonthlySnapshots", err).Error(fmt.Sprintf("Repo call to GetAccountTransactions failed"))
		return nil, utils.InternalServerError
	}

	if len(transactions) == 0 {
		res := &GetAccountMonthlySnapshotsResponse{
			AccountMonthlySnapshots: nil,
		}
		return res, nil
	}

	// Get oldest date
	oldestDate := transactions[len(transactions)-1].Date
	oldestDateTime, _ := time.Parse("2006-01-02", oldestDate)
	oldestDate = oldestDateTime.AddDate(0, 0, 1+(-1*oldestDateTime.Day())).Format("2006-01-02")

	// Find current date
	currentDate := time.Now().AddDate(0, 0, 1+(-1*time.Now().Day())).Format("2006-01-02")

	// Get account to find current balance
	account, err := s.financialAccountRepo.GetAccountByID(tx, userID, req.GetAccountId())
	if err != nil {
		logger("GetAccountMonthlySnapshots", err).Error(fmt.Sprintf("Repo call to GetAccountByID failed"))
		return nil, utils.InternalServerError
	}
	availableBalance := account.AvailableBalance

	if account.AccountType == "credit" {
		availableBalance = account.CurrentBalance
	}
	accountMonthlySnapshots := make([]*AccountMonthlySnapshot, 0)

	// Loop through dates starting from current date to oldest date
	for date := currentDate; DateComparator(date, oldestDate); date = MonthDecrementer(date) {
		// Get all transactions on that date => list of transactions on date
		transactionsInMonth := make([]models.FinancialTransaction, 0)
		for _, transaction := range transactions {
			if WithinMonth(transaction.Date, date) {
				transactionsInMonth = append(transactionsInMonth, transaction)
			}
		}

		// Loop through list of transactions on date and find balance, cash in, and cash out
		monthlyCashOut := 0.0
		monthlyCashIn := 0.0
		for _, transaction := range transactionsInMonth {
			if transaction.Amount > 0 {
				monthlyCashOut += transaction.Amount
			} else {
				monthlyCashIn += (transaction.Amount * -1)
			}
		}
		endBalance := availableBalance
		if account.AccountType == "credit" {
			availableBalance = availableBalance - monthlyCashOut + monthlyCashIn
		} else {
			availableBalance = availableBalance + monthlyCashOut - monthlyCashIn
		}

		// Insert new daily_account record with the above info for the date.
		accountMonthlySnapshot := &AccountMonthlySnapshot{
			Date:              date,
			StartMonthBalance: math.Round(availableBalance*100) / 100,
			EndMonthBalance:   math.Round(endBalance*100) / 100,
			CashOut:           math.Round(monthlyCashOut*100) / 100,
			CashIn:            math.Round(monthlyCashIn*100) / 100,
		}
		accountMonthlySnapshots = append(accountMonthlySnapshots, accountMonthlySnapshot)
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("GetAccountMonthlySnapshots", err).Error(utils.CommitTxErrorMsg)
		return nil, err
	}

	res := &GetAccountMonthlySnapshotsResponse{
		AccountMonthlySnapshots: accountMonthlySnapshots,
	}
	return res, nil
}

// GetCategoryMonthlySnapshots get monthly snapshots for a category
func (s *Service) GetCategoryMonthlySnapshots(ctx context.Context, req *GetCategoryMonthlySnapshotsRequest) (*GetCategoryMonthlySnapshotsResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("GetCategoryMonthlySnapshots", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	userID, err := utils.GetUserIDMetadata(ctx)
	if err != nil {
		logger("GetCategoryMonthlySnapshots", err).Error(fmt.Sprintf("GetUserIDMetadata failed"))
		return nil, utils.InternalServerError
	}

	transactions, err := s.financialTransactionRepo.GetUserTransactions(tx, userID)
	if err != nil {
		logger("GetCategoryMonthlySnapshots", err).Error(fmt.Sprintf("Repo call to GetUserTransactions failed"))
		return nil, utils.InternalServerError
	}

	if len(transactions) == 0 {
		res := &GetCategoryMonthlySnapshotsResponse{
			AccountMonthlySnapshots: nil,
		}
		return res, nil
	}

	// Get oldest date
	oldestDate := transactions[len(transactions)-1].Date
	oldestDateTime, _ := time.Parse("2006-01-02", oldestDate)
	oldestDate = oldestDateTime.AddDate(0, 0, 1+(-1*oldestDateTime.Day())).Format("2006-01-02")

	// Find current date
	currentDate := time.Now().AddDate(0, 0, 1+(-1*time.Now().Day())).Format("2006-01-02")

	accountMonthlySnapshots := make([]*AccountMonthlySnapshot, 0)

	// Loop through dates starting from current date to oldest date
	for date := currentDate; DateComparator(date, oldestDate); date = MonthDecrementer(date) {
		// Get all transactions on that date => list of transactions on date
		transactionsInMonth := make([]models.FinancialTransaction, 0)
		for _, transaction := range transactions {
			if WithinMonth(transaction.Date, date) {
				if req.GetCategoryId() == transaction.CategoryID {
					transactionsInMonth = append(transactionsInMonth, transaction)
				}
			}
		}

		// Loop through list of transactions on date and find balance, cash in, and cash out
		monthlyCashOut := 0.0
		monthlyCashIn := 0.0
		for _, transaction := range transactionsInMonth {
			if transaction.Amount > 0 {
				monthlyCashOut += transaction.Amount
			} else {
				monthlyCashIn += (transaction.Amount * -1)
			}
		}

		// Insert new daily_account record with the above info for the date.
		accountMonthlySnapshot := &AccountMonthlySnapshot{
			Date:    date,
			CashOut: math.Round(monthlyCashOut*100) / 100,
			CashIn:  math.Round(monthlyCashIn*100) / 100,
		}
		accountMonthlySnapshots = append(accountMonthlySnapshots, accountMonthlySnapshot)
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("GetCategoryMonthlySnapshots", err).Error(utils.CommitTxErrorMsg)
		return nil, err
	}

	res := &GetCategoryMonthlySnapshotsResponse{
		AccountMonthlySnapshots: accountMonthlySnapshots,
	}
	return res, nil
}

// DateComparator returns true if date1 is greater than or equal to date2
func DateComparator(date1 string, date2 string) bool {
	// Date format: YYYY-MM-DD
	date1Slice := strings.Split(date1, "-")
	date2Slice := strings.Split(date2, "-")
	// Year comparison
	if date1Slice[0] > date2Slice[0] {
		return true
	} else if date1Slice[0] < date2Slice[0] {
		return false
	}
	// Month comparison
	if date1Slice[1] > date2Slice[1] {
		return true
	} else if date1Slice[1] < date2Slice[1] {
		return false
	}

	// Day comparison
	if date1Slice[2] >= date2Slice[2] {
		return true
	} else if date1Slice[2] < date2Slice[2] {
		return false
	}
	return false
}

// DateDecrementer decrements the given date and returns it
func DateDecrementer(date string) string {
	// Date format: YYYY-MM-DD
	d, _ := time.Parse("2006-01-02", date)
	decrementedDate := d.AddDate(0, 0, -1)
	return decrementedDate.Format("2006-01-02")
}

// WithinMonth returns true if date is within the given month and year
func WithinMonth(date string, monthYear string) bool {
	// Date format: YYYY-MM-DD
	date1Slice := strings.Split(date, "-")
	date2Slice := strings.Split(monthYear, "-")
	// Year comparison
	if date1Slice[0] != date2Slice[0] {
		return false
	}
	// Month comparison
	if date1Slice[1] != date2Slice[1] {
		return false
	}
	return true
}

// MonthDecrementer decrements the given date by a month and returns it
func MonthDecrementer(date string) string {
	// Date format: YYYY-MM-DD
	d, _ := time.Parse("2006-01-02", date)
	decrementedDate := d.AddDate(0, -1, 0)
	return decrementedDate.Format("2006-01-02")
}

// TransactionForPython ...
type TransactionForPython struct {
	ID                 int64   `json:"id"`
	CategoryID         int64   `json:"category_id"`
	PlaidCategoryID    string  `json:"plaid_category_id"`
	PlaidTransactionID string  `json:"plaid_transaction_id"`
	Name               string  `json:"transaction_name"`
	Amount             float64 `json:"amount"`
	Date               string  `json:"date"`
}

// FindRecurringTransactions add recurring transactions to db
func (s *Service) FindRecurringTransactions(ctx context.Context, req *Empty) (*FindRecurringTransactionsResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	userID, err := utils.GetUserIDMetadata(ctx)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("GetUserIDMetadata failed"))
		return nil, utils.InternalServerError
	}

	transactions, err := s.financialTransactionRepo.GetUserTransactions(tx, userID)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Repo call to GetUserTransactions failed"))
		return nil, utils.InternalServerError
	}

	if len(transactions) == 0 {
		res := &FindRecurringTransactionsResponse{
			Success: true,
		}
		return res, nil
	}

	var pythonTransactions []*TransactionForPython
	for _, transaction := range transactions[0:200] {
		pythonTransactions = append(pythonTransactions, dataToTransactionForPython(transaction))
	}

	var jsonData []byte
	jsonData, err = json.Marshal(pythonTransactions)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to marshal json request"))
		return nil, utils.InternalServerError
	}

	response, err := http.Post("http://127.0.0.1:5001/recurring", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to make post call to dataprocessing server"))
		return nil, utils.InternalServerError
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to read response"))
		return nil, utils.InternalServerError
	}

	var recurringTransactions []*RecurringTransaction
	err = json.Unmarshal(responseData, &recurringTransactions)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to unmarshal json response"))
		return nil, utils.InternalServerError
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(utils.CommitTxErrorMsg)
		return nil, err
	}

	res := &FindRecurringTransactionsResponse{
		Success: true,
	}
	return res, nil
}

// GetRecurringTransactions get monthly snapshots for a category
func (s *Service) GetRecurringTransactions(ctx context.Context, req *Empty) (*GetRecurringTransactionsResponse, error) {
	tx, err := s.txRepo.StartTx(ctx)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}

	userID, err := utils.GetUserIDMetadata(ctx)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("GetUserIDMetadata failed"))
		return nil, utils.InternalServerError
	}

	transactions, err := s.financialTransactionRepo.GetUserTransactions(tx, userID)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Repo call to GetUserTransactions failed"))
		return nil, utils.InternalServerError
	}

	if len(transactions) == 0 {
		res := &GetRecurringTransactionsResponse{
			RecurringTransactions: nil,
		}
		return res, nil
	}

	var pythonTransactions []*TransactionForPython
	for _, transaction := range transactions[0:200] {
		pythonTransactions = append(pythonTransactions, dataToTransactionForPython(transaction))
	}

	var jsonData []byte
	jsonData, err = json.Marshal(pythonTransactions)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to marshal json request"))
		return nil, utils.InternalServerError
	}
	// fmt.Println(string(jsonData))

	response, err := http.Post("http://127.0.0.1:5001/recurring", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to make post call to dataprocessing server"))
		return nil, utils.InternalServerError
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to read response"))
		return nil, utils.InternalServerError
	}

	var recurringTransactions []*RecurringTransaction
	err = json.Unmarshal(responseData, &recurringTransactions)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(fmt.Sprintf("Failed to unmarshal json response"))
		return nil, utils.InternalServerError
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("GetRecurringTransactions", err).Error(utils.CommitTxErrorMsg)
		return nil, err
	}

	res := &GetRecurringTransactionsResponse{
		RecurringTransactions: recurringTransactions,
	}
	return res, nil
}

func dataToTransactionForPython(data models.FinancialTransaction) *TransactionForPython {
	return &TransactionForPython{
		ID:                 data.ID,
		CategoryID:         data.CategoryID,
		PlaidCategoryID:    data.PlaidCategoryID,
		PlaidTransactionID: data.PlaidTransactionID,
		Name:               data.Name,
		Amount:             data.Amount,
		Date:               data.Date,
	}
}
