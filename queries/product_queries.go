package queries

const (
	FindAllProducts = `
		SELECT id, name, description, price, category_id, sku, quantity_in_stock, image_url, created_at, updated_at
		FROM products
		ORDER BY id
	`

	FindProductByID = `
		SELECT id, name, description, price, category_id, sku, quantity_in_stock, image_url, created_at, updated_at
		FROM products
		WHERE id = $1
	`

	InsertProduct = `
		INSERT INTO products (name, description, price, category_id, sku, quantity_in_stock, image_url, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW(), NOW())
		RETURNING id
	`

	UpdateProduct = `
		UPDATE products
		SET name = $1, description = $2, price = $3, category_id = $4,
		    sku = $5, quantity_in_stock = $6, image_url = $7, updated_at = NOW()
		WHERE id = $8
	`

	DeleteProduct = `
		DELETE FROM products
		WHERE id = $1
	`
)
