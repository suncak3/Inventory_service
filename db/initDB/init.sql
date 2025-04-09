-- Drop tables if they exist
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;

-- Create Category table
CREATE TABLE categories (
                            id SERIAL PRIMARY KEY,
                            name VARCHAR(255) NOT NULL
);

-- Create Product table
CREATE TABLE products (
                          id SERIAL PRIMARY KEY,
                          name VARCHAR(255) NOT NULL,
                          description TEXT,
                          price NUMERIC(10, 2) NOT NULL,
                          stock_level INTEGER NOT NULL,
                          category_id INTEGER REFERENCES categories(id) ON DELETE SET NULL
);

-- Insert sample categories
INSERT INTO categories (name) VALUES
                                  ('Electronics'),
                                  ('Books'),
                                  ('Clothing');

-- Insert sample products
INSERT INTO products (name, description, price, stock_level, category_id) VALUES
                                                                              ('Smartphone', 'Latest model with 128GB storage', 699.99, 50, 1),
                                                                              ('Novel Book', 'Bestseller fiction', 19.99, 200, 2),
                                                                              ('T-shirt', 'Cotton t-shirt in various sizes', 12.50, 150, 3);
