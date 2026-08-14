package controller

import (
	"e-commerce/model"
	"e-commerce/service"
	"e-commerce/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "strconv" // Adicione esta importação
    "github.com/gorilla/mux" // Adicione esta importação
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
	// Extrai o ID da URL usando mux
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		utils.HandleError(w, http.StatusBadRequest, "Product ID not provided")
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID format")
		return
	}

	product, err := c.productService.GetProductByID(id)
	if err != nil {
		utils.HandleError(w, http.StatusNotFound, "Product not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (c *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		utils.HandleError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	log.Printf("=== Recebido produto: Name=%s, Price=%f, CategoryID=%d, SKU=%s",
		product.Name, product.Price, product.CategoryID, product.SKU)

	id, err := c.productService.CreateProduct(product)
	if err != nil {
		log.Printf("=== ERRO ao criar produto: %v ===", err)
		utils.HandleError(w, http.StatusInternalServerError, "Failed to create product: "+err.Error())
		return
	}

	log.Printf("=== Produto criado com ID: %d ===", id)

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
	// Get ID from URL
	id, err := utils.ParseID(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
		return
	}

	// Decode the JSON request body into a Product struct
	var product model.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Set the ID from URL
	product.ID = id

	// Call the service method with the product struct
	_, err = c.productService.UpdateProduct(product)
	if err != nil {
		http.Error(w, "Failed to update product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return 200 OK for successful update
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Product updated successfully",
	})
}

func (c *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL
	id, err := utils.ParseID(r)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, "Invalid product ID")
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