package entities

import (
	"github.com/google/uuid"
	"time"
)

// Order struct
type Order struct {
	ID              uuid.UUID      `json:"id"`
	UserID          uuid.UUID      `json:"user_id"`
	OrderItems      []OrderItem    `json:"order_items"`
	TotalAmount     float64        `json:"total_amount"`
	Status          string         `json:"status"`
	ShippingAddress Address        `json:"shipping_address"`
	BillingAddress  Address        `json:"billing_address"`
	PaymentDetails  PaymentDetails `json:"payment_details"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
