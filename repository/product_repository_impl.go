package repository

import (
	"database/sql"
	"e-commerce/model"
	"e-commerce/queries"
	"log"
)

type ProductRepositoryImpl struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &ProductRepositoryImpl{
		DB: db,
	}
}

func (r *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	rows, err := r.DB.Query(queries.FindAllProducts)
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

func (r *ProductRepositoryImpl) FindByID(id int64) (model.Product, error) {
	var product model.Product
	err := r.DB.QueryRow(queries.FindProductByID, id).Scan(
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

func (r *ProductRepositoryImpl) Save(product model.Product) (int64, error) {
	log.Printf("=== Repository: Salvando produto ===")
	log.Printf("Repository: Name=%s, Price=%f, CategoryID=%d, SKU=%s",
		product.Name, product.Price, product.CategoryID, product.SKU)

	var id int64
	err := r.DB.QueryRow(
		queries.InsertProduct,
		product.Name,
		product.Description,
		product.Price,
		product.CategoryID,
		product.SKU,
		product.QuantityInStock,
		product.ImageURL,
	).Scan(&id)

	if err != nil {
		log.Printf("=== ERRO no PostgreSQL: %v ===", err)
		return 0, err
	}

	log.Printf("=== Repository: Produto salvo com ID: %d ===", id)
	return id, nil
}

func (r *ProductRepositoryImpl) Update(product model.Product) (int64, error) {
	result, err := r.DB.Exec(
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
	return result.RowsAffected()
}

func (r *ProductRepositoryImpl) Delete(id int64) (int64, error) {
	result, err := r.DB.Exec(queries.DeleteProduct, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}