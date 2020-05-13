package models

// PlaidWebhook ...
type PlaidWebhook struct {
	WebhookType         string   `json:"webhook_type"`
	WebhookCode         string   `json:"webhook_code"`
	ItemIDPlaid         string   `json:"item_id"`
	NewTransactionCount int      `json:"new_transactions"`
	RemovedTransactions []string `json:"removed_transactions"`
}
