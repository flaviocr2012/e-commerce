package repository

import (
	"database/sql"
	"e-commerce/model"
	"e-commerce/queries"
	"log"
)

type CartRepository interface {
	GetCartByUserID(userID int64) (*model.Cart, error)
	CreateCart(userID int64) (int64, error)
	GetCartItems(cartID int64) ([]model.CartItem, error)
	AddCartItem(cartID int64, productID int64, quantity int) (int64, error)
	UpdateCartItemQuantity(itemID, cartID int64, quantity int) error
	RemoveCartItem(itemID, cartID int64) error
	ClearCart(cartID int64) error
	GetCartItemByID(itemID, cartID int64) (*model.CartItem, error)
	DeleteCart(userID int64) error
}

type CartRepositoryImpl struct {
	DB *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &CartRepositoryImpl{
		DB: db,
	}
}

func (r *CartRepositoryImpl) GetCartByUserID(userID int64) (*model.Cart, error) {
	var cart model.Cart
	err := r.DB.QueryRow(queries.GetCartByUserID, userID).Scan(
		&cart.ID,
		&cart.UserID,
		&cart.CreatedAt,
		&cart.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error getting cart for user %d: %v", userID, err)
		return nil, err
	}
	return &cart, nil
}

func (r *CartRepositoryImpl) CreateCart(userID int64) (int64, error) {
	var id int64
	err := r.DB.QueryRow(queries.CreateCart, userID).Scan(&id)
	if err != nil {
		log.Printf("Error creating cart for user %d: %v", userID, err)
		return 0, err
	}
	return id, nil
}

func (r *CartRepositoryImpl) GetCartItems(cartID int64) ([]model.CartItem, error) {
	rows, err := r.DB.Query(queries.GetCartItems, cartID)
	if err != nil {
		log.Printf("Error getting cart items for cart %d: %v", cartID, err)
		return nil, err
	}
	defer rows.Close()

	var items []model.CartItem
	for rows.Next() {
		var item model.CartItem
		var product model.Product
		err := rows.Scan(
			&item.ID,
			&item.CartID,
			&item.ProductID,
			&item.Quantity,
			&item.CreatedAt,
			&item.UpdatedAt,
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CategoryID,
			&product.SKU,
			&product.QuantityInStock,
			&product.ImageURL,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning cart item: %v", err)
			return nil, err
		}
		item.Product = product
		item.Subtotal = product.Price * float64(item.Quantity)
		items = append(items, item)
	}
	return items, nil
}

func (r *CartRepositoryImpl) AddCartItem(cartID int64, productID int64, quantity int) (int64, error) {
	var id int64
	err := r.DB.QueryRow(queries.AddCartItem, cartID, productID, quantity).Scan(&id)
	if err != nil {
		log.Printf("Error adding item to cart %d: %v", cartID, err)
		return 0, err
	}
	return id, nil
}

func (r *CartRepositoryImpl) UpdateCartItemQuantity(itemID, cartID int64, quantity int) error {
	result, err := r.DB.Exec(queries.UpdateCartItemQuantity, quantity, itemID, cartID)
	if err != nil {
		log.Printf("Error updating cart item %d: %v", itemID, err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *CartRepositoryImpl) RemoveCartItem(itemID, cartID int64) error {
	result, err := r.DB.Exec(queries.RemoveCartItem, itemID, cartID)
	if err != nil {
		log.Printf("Error removing cart item %d: %v", itemID, err)
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *CartRepositoryImpl) ClearCart(cartID int64) error {
	_, err := r.DB.Exec(queries.ClearCart, cartID)
	if err != nil {
		log.Printf("Error clearing cart %d: %v", cartID, err)
		return err
	}
	return nil
}

func (r *CartRepositoryImpl) GetCartItemByID(itemID, cartID int64) (*model.CartItem, error) {
	var item model.CartItem
	err := r.DB.QueryRow(queries.GetCartItemByID, itemID, cartID).Scan(
		&item.ID,
		&item.CartID,
		&item.ProductID,
		&item.Quantity,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error getting cart item %d: %v", itemID, err)
		return nil, err
	}
	return &item, nil
}

func (r *CartRepositoryImpl) DeleteCart(userID int64) error {
	_, err := r.DB.Exec(queries.DeleteCart, userID)
	if err != nil {
		log.Printf("Error deleting cart for user %d: %v", userID, err)
		return err
	}
	return nil
}