package model

import "time"

type Cart struct {
    ID        int64     `json:"id"`
    UserID    int64     `json:"user_id"`
    Items     []CartItem `json:"items,omitempty"`
    Total     float64   `json:"total,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type CartItem struct {
    ID        int64     `json:"id"`
    CartID    int64     `json:"cart_id"`
    ProductID int64     `json:"product_id"`
    Product   Product   `json:"product,omitempty"`
    Quantity  int       `json:"quantity"`
    Subtotal  float64   `json:"subtotal,omitempty"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

type AddToCartRequest struct {
    ProductID int64 `json:"product_id" validate:"required"`
    Quantity  int   `json:"quantity" validate:"required,min=1"`
}

type UpdateCartItemRequest struct {
    Quantity int `json:"quantity" validate:"required,min=1"`
}