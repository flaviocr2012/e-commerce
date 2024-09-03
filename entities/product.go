package entities

type Product struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	Price           float64 `json:"price"`
	CategoryID      int64   `json:"category_id"`
	SKU             string  `json:"sku"`
	QuantityInStock int64   `json:"quantity_in_stock"`
	ImageURL        string  `json:"image_url"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}
