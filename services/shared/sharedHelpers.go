package shared

import "github.com/neelchoudhary/budgetwallet-api-server/models"

// DataToAccountPb convert models Financial Account to pb Financial Account
func DataToAccountPb(data models.FinancialAccount) *FinancialAccount {
	return &FinancialAccount{
		Id:             data.ID,
		UserId:         data.UserID,
		ItemId:         data.ItemID,
		PlaidAccountId: data.PlaidAccountID,
		CurrentBalance: data.CurrentBalance,
		AccountName:    data.AccountName,
		OfficialName:   data.OfficialName,
		AccountType:    data.AccountType,
		AccountSubtype: data.AccountSubType,
		AccountMask:    data.AccountMask,
		Selected:       data.Selected,
	}
}
