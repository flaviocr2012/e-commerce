package repository

import (
	"database/sql"
	"e-commerce/model"
	"e-commerce/queries"
	"github.com/google/uuid"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (repository *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	rows, err := repository.DB.Query(queries.FindAllProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(
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
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repository *ProductRepositoryImpl) FindByID(id uuid.UUID) (model.Product, error) {
	var product model.Product
	err := repository.DB.QueryRow(queries.FindProductByID, id).Scan(
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
		if err == sql.ErrNoRows {
			return product, nil
		}
		return product, err
	}
	return product, nil
}

func (repository *ProductRepositoryImpl) Save(product model.Product) (uuid.UUID, error) {
	product.ID = uuid.New()
	err := repository.DB.QueryRow(
		queries.InsertProduct,
		product.ID,
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.SKU,
		product.QuantityInStock,
		product.ImageURL,
	).Scan(&product.ID)
	if err != nil {
		return uuid.Nil, err
	}
	return product.ID, nil
}

func (repository *ProductRepositoryImpl) Update(product model.Product) (int64, error) {
	result, err := repository.DB.Exec(
		queries.UpdateProduct,
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.SKU,
		product.QuantityInStock,
		product.ImageURL,
		product.ID,
	)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func (repository *ProductRepositoryImpl) Delete(id uuid.UUID) (int64, error) {
	result, err := repository.DB.Exec(queries.DeleteProduct, id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
