package dataprocessing

import (
	context "context"
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
		logger("GetFinancialCategories", err).Error(utils.StartTxErrorMsg)
		return nil, utils.InternalServerError
	}
	transactions, err := s.financialTransactionRepo.GetAccountTransactions(tx, req.GetUserId(), req.GetAccountId())
	if err != nil {
		// todo
	}

	// Get oldest date
	oldestDate := transactions[len(transactions)-1].Date

	// Find current date
	currentDate := time.Now().Format("2006-01-02")

	// Get account to find current balance
	account, err := s.financialAccountRepo.GetAccountByID(tx, req.GetUserId(), req.GetAccountId())
	if err != nil {
		return nil, err
	}
	currentBalance := account.CurrentBalance

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
		monthlyCashOut := 0.0
		monthlyCashIn := 0.0
		for _, transaction := range transactionsOnDate {
			if transaction.Amount > 0 {
				dailyCashOut += transaction.Amount
				monthlyCashOut += transaction.Amount
			} else {
				dailyCashIn += (transaction.Amount * -1)
				monthlyCashIn += (transaction.Amount * -1)
			}
		}
		currentBalance = currentBalance + dailyCashOut - dailyCashIn

		// Insert new daily_account record with the above info for the date.
		accountDailySnapshot := &AccountDailySnapshot{
			ItemId:    req.GetItemId(),
			AccountId: req.GetAccountId(),
			Date:      date,
			Balance:   currentBalance,
			CashOut:   dailyCashOut,
			CashIn:    dailyCashIn,
		}
		accountDailySnapshots = append(accountDailySnapshots, accountDailySnapshot)
	}

	err = s.txRepo.CommitTx(tx)
	if err != nil {
		logger("GetFinancialCategories", err).Error(utils.CommitTxErrorMsg)
		return nil, err
	}

	res := &GetAccountDailySnapshotsResponse{
		AccountDailySnapshots: accountDailySnapshots,
	}
	return res, nil
}

// DateComparator returns true if date1 is greater than or equal to date2
func DateComparator(date1 string, date2 string) bool {
	// Date format: YYYY-MM-DD
	date1Slice := strings.Split(date1, "-")
	date2Slice := strings.Split(date2, "-")
	if date1Slice[0] > date2Slice[0] {
		return true
	}
	if date1Slice[1] > date2Slice[1] {
		return true
	}
	if date1Slice[2] >= date2Slice[2] {
		return true
	}
	return false
}

// DateDecrementer decrements the given date and returns it
func DateDecrementer(date string) string {
	// Date format: YYYY-MM-DD
	// TODO handle error
	d, _ := time.Parse("2006-01-02", date)
	decrementedDate := d.AddDate(0, 0, -1)
	return decrementedDate.Format("2006-01-02")
}
