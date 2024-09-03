package entities

// PaymentDetails struct
type PaymentDetails struct {
	ID            int64  `json:"id"`
	OrderID       int64  `json:"order_id"`
	PaymentMethod string `json:"payment_method"`
	TransactionID string `json:"transaction_id"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
}
