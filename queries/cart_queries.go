package queries

const (
    // Cart queries
    GetCartByUserID = `
        SELECT id, user_id, created_at, updated_at
        FROM carts
        WHERE user_id = $1
    `

    CreateCart = `
        INSERT INTO carts (user_id, created_at, updated_at)
        VALUES ($1, NOW(), NOW())
        RETURNING id
    `

    // Cart Item queries
    GetCartItems = `
        SELECT ci.id, ci.cart_id, ci.product_id, ci.quantity,
               ci.created_at, ci.updated_at,
               p.id, p.name, p.description, p.price, p.category_id,
               p.sku, p.quantity_in_stock, p.image_url,
               p.created_at, p.updated_at
        FROM cart_items ci
        JOIN products p ON ci.product_id = p.id
        WHERE ci.cart_id = $1
        ORDER BY ci.created_at DESC
    `

    AddCartItem = `
        INSERT INTO cart_items (cart_id, product_id, quantity, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        ON CONFLICT (cart_id, product_id)
        DO UPDATE SET quantity = cart_items.quantity + EXCLUDED.quantity,
                      updated_at = NOW()
        RETURNING id
    `

    UpdateCartItemQuantity = `
        UPDATE cart_items
        SET quantity = $1, updated_at = NOW()
        WHERE id = $2 AND cart_id = $3
        RETURNING id
    `

    RemoveCartItem = `
        DELETE FROM cart_items
        WHERE id = $1 AND cart_id = $2
        RETURNING id
    `

    ClearCart = `
        DELETE FROM cart_items
        WHERE cart_id = $1
        RETURNING id
    `

    GetCartItemByID = `
        SELECT id, cart_id, product_id, quantity, created_at, updated_at
        FROM cart_items
        WHERE id = $1 AND cart_id = $2
    `

    DeleteCart = `
        DELETE FROM carts
        WHERE user_id = $1
    `
)