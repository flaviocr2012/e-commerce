package controller

import (
	"e-commerce/model"
	"e-commerce/service"
	"e-commerce/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

func (c *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.productService.GetAllProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ParseID(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := c.productService.GetProductByID(id)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, "Product not found")
		return
	}

	response := utils.Response{
		Status: "success",
		Data:   product,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Erro ao decodificar o corpo da requisição: %v", err)
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	id, err := c.productService.CreateProduct(product)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, "Failed to create product")
		return
	}

	// Set Location header
	w.Header().Set("Location", fmt.Sprintf("/products/%d", id))
	w.WriteHeader(http.StatusCreated)

	response := utils.Response{
		Status: "success",
		Data:   id,
	}
	json.NewEncoder(w).Encode(response)
}

func (c *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a Product struct
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service method with the product struct
	_, err = c.productService.UpdateProduct(product)
	if err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return 200 OK for successful update
	w.WriteHeader(http.StatusOK)
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' from the query string
	idStr := r.URL.Query().Get("id")

	// Convert the 'id' from string to int64
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid product ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Call the service method with the converted id
	_, err = c.productService.DeleteProduct(id)
	if err != nil {
		http.Error(w, "Failed to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return 204 No Content for successful deletion
	w.WriteHeader(http.StatusNoContent)
}
