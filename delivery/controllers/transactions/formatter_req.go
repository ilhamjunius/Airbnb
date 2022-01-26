package transactions

type UpdateTransactionsRequestFormat struct {
	InvoiceID string `json:"invoice_id"`
	Status    string `json:"status"`
}
