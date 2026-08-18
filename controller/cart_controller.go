package controller

import (
	"e-commerce/model"
	"e-commerce/service"
	"e-commerce/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CartController struct {
	cartService service.CartService
}

func NewCartController(cartService service.CartService) *CartController {
	return &CartController{
		cartService: cartService,
	}
}

// GetCart handles GET /carts - Get user's cart
func (c *CartController) GetCart(w http.ResponseWriter, r *http.Request) {
	// TODO: Get user ID from JWT token (for now, use a hardcoded user ID)
	userID := int64(1)

	cart, err := c.cartService.GetCart(userID)
	if err != nil {
		log.Printf("Error getting cart: %v", err)
		utils.HandleError(w, http.StatusInternalServerError, "Failed to get cart")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cart)
}

// AddToCart handles POST /carts/items - Add item to cart
func (c *CartController) AddToCart(w http.ResponseWriter, r *http.Request) {
	var req model.AddToCartRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.ProductID <= 0 {
		utils.HandleError(w, http.StatusBadRequest, "Product ID is required")
		return
	}
	if req.Quantity <= 0 {
		utils.HandleError(w, http.StatusBadRequest, "Quantity must be greater than 0")
		return
	}

	userID := int64(1)

	err = c.cartService.AddToCart(userID, &req)
	if err != nil {
		log.Printf("Error adding to cart: %v", err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Item added to cart successfully",
	})
}

// UpdateCartItem handles PUT /carts/items/{id} - Update item quantity
func (c *CartController) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemIDStr, ok := vars["id"]
	if !ok {
		utils.HandleError(w, http.StatusBadRequest, "Item ID not provided")
		return
	}

	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid item ID format")
		return
	}

	var req model.UpdateCartItemRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if req.Quantity <= 0 {
		utils.HandleError(w, http.StatusBadRequest, "Quantity must be greater than 0")
		return
	}

	userID := int64(1)

	err = c.cartService.UpdateCartItem(userID, itemID, &req)
	if err != nil {
		log.Printf("Error updating cart item: %v", err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Cart item updated successfully",
	})
}

// RemoveFromCart handles DELETE /carts/items/{id} - Remove item from cart
func (c *CartController) RemoveFromCart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemIDStr, ok := vars["id"]
	if !ok {
		utils.HandleError(w, http.StatusBadRequest, "Item ID not provided")
		return
	}

	itemID, err := strconv.ParseInt(itemIDStr, 10, 64)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid item ID format")
		return
	}

	userID := int64(1)

	err = c.cartService.RemoveFromCart(userID, itemID)
	if err != nil {
		log.Printf("Error removing from cart: %v", err)
		utils.HandleError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ClearCart handles DELETE /carts/clear - Clear entire cart
func (c *CartController) ClearCart(w http.ResponseWriter, r *http.Request) {
	userID := int64(1)

	err := c.cartService.ClearCart(userID)
	if err != nil {
		log.Printf("Error clearing cart: %v", err)
		utils.HandleError(w, http.StatusInternalServerError, "Failed to clear cart")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}