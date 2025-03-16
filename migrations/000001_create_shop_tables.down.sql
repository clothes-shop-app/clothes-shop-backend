DROP TABLE IF EXISTS product_images;
DROP TABLE IF EXISTS product_attribute_values;
DROP TABLE IF EXISTS attribute_values;
DROP TABLE IF EXISTS attributes;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS users;

DROP INDEX IF EXISTS idx_users_login;
DROP INDEX IF EXISTS idx_categories_lft_rgt;
DROP INDEX IF EXISTS idx_products_category_id;
DROP INDEX IF EXISTS idx_products_price;
DROP INDEX IF EXISTS idx_products_created_at;
DROP INDEX IF EXISTS idx_attributes_name;
DROP INDEX IF EXISTS idx_attribute_values_attribute_id;
DROP INDEX IF EXISTS idx_attribute_values_value;
DROP INDEX IF EXISTS idx_product_attribute_values_product_id;
DROP INDEX IF EXISTS idx_product_attribute_values_attribute_value_id;
DROP INDEX IF EXISTS idx_product_images_product_id;
DROP INDEX IF EXISTS idx_product_images_position;

DROP EXTENSION IF EXISTS "uuid-ossp";
