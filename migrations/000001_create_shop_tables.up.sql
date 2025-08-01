CREATE TABLE users (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    phone VARCHAR(255) UNIQUE NOT NULL,
    name TEXT NOT NULL,
    avatar_url TEXT NULL,
    address TEXT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE categories (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name TEXT NOT NULL,
    lft INT NULL, -- Nested Set 
    rgt INT NULL -- Nested Set
);

CREATE INDEX idx_categories_lft_rgt ON categories(lft, rgt);

CREATE TABLE products (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name TEXT NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    category_id CHAR(36) NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_price ON products(price);
CREATE INDEX idx_products_created_at ON products(created_at);

CREATE TABLE attributes (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name TEXT NOT NULL
);

CREATE TABLE attribute_values (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    attribute_id CHAR(36),
    value TEXT NOT NULL,
    FOREIGN KEY (attribute_id) REFERENCES attributes(id) ON DELETE CASCADE
);

CREATE INDEX idx_attribute_values_attribute_id ON attribute_values(attribute_id);

-- Many-to-many 
CREATE TABLE product_attribute_values (
    product_id CHAR(36),
    attribute_value_id CHAR(36),
    PRIMARY KEY (product_id, attribute_value_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (attribute_value_id) REFERENCES attribute_values(id) ON DELETE CASCADE
);

CREATE INDEX idx_product_attribute_values_product_id ON product_attribute_values(product_id);
CREATE INDEX idx_product_attribute_values_attribute_value_id ON product_attribute_values(attribute_value_id);

CREATE TABLE product_images (
    id CHAR(36) PRIMARY KEY DEFAULT (UUID()),
    product_id CHAR(36),
    image_url TEXT NOT NULL,
    position INT NOT NULL DEFAULT 0, -- sort order
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
);
