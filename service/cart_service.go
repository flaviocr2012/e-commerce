package service

import "e-commerce/model"

type CartService interface {
	GetCart(userID int64) (*model.Cart, error)
	AddToCart(userID int64, req *model.AddToCartRequest) error
	UpdateCartItem(userID int64, itemID int64, req *model.UpdateCartItemRequest) error
	RemoveFromCart(userID int64, itemID int64) error
	ClearCart(userID int64) error
}