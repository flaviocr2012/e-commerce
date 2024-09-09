package service

import (
	"e-commerce/model"
	"e-commerce/repository"
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

func (service *ProductServiceImpl) CreateProduct(product model.Product) (int64, error) {
	return service.productRepository.Save(product)
}

func (service *ProductServiceImpl) UpdateProduct(product model.Product) (int64, error) {
	return service.productRepository.Update(product)
}

func (service *ProductServiceImpl) DeleteProduct(id int64) (int64, error) {
	return service.productRepository.Delete(id)
}
