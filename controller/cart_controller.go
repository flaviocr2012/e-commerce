package service

import (
	"database/sql"
	"e-commerce/model"
	"e-commerce/repository"
	"errors"
	"log"
)

type CartServiceImpl struct {
	cartRepository    repository.CartRepository
	productRepository repository.ProductRepository
}

func NewCartService(
	cartRepo repository.CartRepository,
	productRepo repository.ProductRepository,
) *CartServiceImpl {
	return &CartServiceImpl{
		cartRepository:    cartRepo,
		productRepository: productRepo,
	}
}

func (s *CartServiceImpl) GetCart(userID int64) (*model.Cart, error) {
	log.Printf("Getting cart for user %d", userID)

	// Get or create cart
	cart, err := s.getOrCreateCart(userID)
	if err != nil {
		return nil, err
	}

	// Get cart items
	items, err := s.cartRepository.GetCartItems(cart.ID)
	if err != nil {
		return nil, err
	}

	cart.Items = items

	// Calculate total
	var total float64
	for _, item := range items {
		total += item.Subtotal
	}
	cart.Total = total

	return cart, nil
}

func (s *CartServiceImpl) AddToCart(userID int64, req *model.AddToCartRequest) error {
	log.Printf("Adding item to cart for user %d: product %d, quantity %d",
		userID, req.ProductID, req.Quantity)

	// Validate product exists and has stock
	product, err := s.productRepository.FindByID(req.ProductID)
	if err != nil {
		return errors.New("product not found")
	}
	if product.QuantityInStock < req.Quantity {
		return errors.New("insufficient stock")
	}

	// Get or create cart
	cart, err := s.getOrCreateCart(userID)
	if err != nil {
		return err
	}

	// Add item to cart
	_, err = s.cartRepository.AddCartItem(cart.ID, req.ProductID, req.Quantity)
	if err != nil {
		log.Printf("Error adding item to cart: %v", err)
		return err
	}

	return nil
}

func (s *CartServiceImpl) UpdateCartItem(userID int64, itemID int64, req *model.UpdateCartItemRequest) error {
	log.Printf("Updating cart item %d for user %d to quantity %d",
		itemID, userID, req.Quantity)

	if req.Quantity < 1 {
		return errors.New("quantity must be at least 1")
	}

	// Get user's cart
	cart, err := s.getOrCreateCart(userID)
	if err != nil {
		return err
	}

	// Verify item belongs to user's cart
	item, err := s.cartRepository.GetCartItemByID(itemID, cart.ID)
	if err != nil {
		return err
	}
	if item == nil {
		return errors.New("item not found in cart")
	}

	// Update quantity
	err = s.cartRepository.UpdateCartItemQuantity(itemID, cart.ID, req.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("item not found in cart")
		}
		return err
	}

	return nil
}

func (s *CartServiceImpl) RemoveFromCart(userID int64, itemID int64) error {
	log.Printf("Removing item %d from cart for user %d", itemID, userID)

	// Get user's cart
	cart, err := s.getOrCreateCart(userID)
	if err != nil {
		return err
	}

	// Verify item belongs to user's cart
	item, err := s.cartRepository.GetCartItemByID(itemID, cart.ID)
	if err != nil {
		return err
	}
	if item == nil {
		return errors.New("item not found in cart")
	}

	// Remove item
	err = s.cartRepository.RemoveCartItem(itemID, cart.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("item not found in cart")
		}
		return err
	}

	return nil
}

func (s *CartServiceImpl) ClearCart(userID int64) error {
	log.Printf("Clearing cart for user %d", userID)

	// Get user's cart
	cart, err := s.getOrCreateCart(userID)
	if err != nil {
		return err
	}

	// Clear cart items
	err = s.cartRepository.ClearCart(cart.ID)
	if err != nil {
		return err
	}

	return nil
}

// Helper method to get or create a cart for a user
func (s *CartServiceImpl) getOrCreateCart(userID int64) (*model.Cart, error) {
	cart, err := s.cartRepository.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	if cart == nil {
		log.Printf("Creating new cart for user %d", userID)
		cartID, err := s.cartRepository.CreateCart(userID)
		if err != nil {
			return nil, err
		}
		cart = &model.Cart{
			ID:     cartID,
			UserID: userID,
		}
	}

	return cart, nil
}