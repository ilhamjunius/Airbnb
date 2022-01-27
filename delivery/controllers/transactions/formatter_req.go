package transactions

type UpdateTransactionsRequestFormat struct {
	InvoiceID string `json:"invoice_id"`
	Status    string `json:"status"`
}

type PayloadRequestFormat struct {
	OrderId           string `json:"order_id"`
	TransactionStatus string `json:"transaction_status"`
}
