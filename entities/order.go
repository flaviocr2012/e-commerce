package entities

// Order struct
type Order struct {
	ID              int64          `json:"id"`
	UserID          int64          `json:"user_id"`
	OrderItems      []OrderItem    `json:"order_items"`
	TotalAmount     float64        `json:"total_amount"`
	Status          string         `json:"status"`
	ShippingAddress Address        `json:"shipping_address"`
	BillingAddress  Address        `json:"billing_address"`
	PaymentDetails  PaymentDetails `json:"payment_details"`
	CreatedAt       string         `json:"created_at"`
	UpdatedAt       string         `json:"updated_at"`
}
