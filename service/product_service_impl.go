package service

import (
	"e-commerce/model"
	"e-commerce/repository"
    "errors"  // Adicione esta linha
    "log"     // Adicione esta linha
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		productRepository: productRepository,
	}
}

func (service *ProductServiceImpl) GetAllProducts() ([]model.Product, error) {
	return service.productRepository.FindAll()
}

func (service *ProductServiceImpl) GetProductByID(id int64) (model.Product, error) {
	return service.productRepository.FindByID(id)
}

func (s *ProductServiceImpl) CreateProduct(product model.Product) (int64, error) {
	log.Printf("=== Service: Recebendo produto para criar ===")
	log.Printf("Service: Name=%s, Price=%f, CategoryID=%d",
		product.Name, product.Price, product.CategoryID)

	// Validar campos obrigatórios
	if product.Name == "" {
		err := errors.New("name is required")
		log.Printf("Erro de validação: %v", err)
		return 0, err
	}
	if product.Price <= 0 {
		err := errors.New("price must be greater than 0")
		log.Printf("Erro de validação: %v", err)
		return 0, err
	}
	if product.SKU == "" {
		err := errors.New("sku is required")
		log.Printf("Erro de validação: %v", err)
		return 0, err
	}

	log.Printf("Service: Validação passou, chamando repository.Save()")

	id, err := s.productRepository.Save(product)
	if err != nil {
		log.Printf("=== ERRO no repository: %v ===", err)
		return 0, err
	}

	log.Printf("=== Service: Produto criado com ID: %d ===", id)
	return id, nil
}

func (service *ProductServiceImpl) UpdateProduct(product model.Product) (int64, error) {
	return service.productRepository.Update(product)
}

func (service *ProductServiceImpl) DeleteProduct(id int64) (int64, error) {
	return service.productRepository.Delete(id)
}
