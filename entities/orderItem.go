package entities

// OrderItem struct
type OrderItem struct {
	ID        int64   `json:"id"`
	OrderID   int64   `json:"order_id"`
	ProductID int64   `json:"product_id"`
	Quantity  int64   `json:"quantity"`
	Price     float64 `json:"price"`
}
