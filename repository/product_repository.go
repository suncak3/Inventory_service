package repository

import (
	"github.com/jmoiron/sqlx"
	"inventory-service/db"
	"inventory-service/domain"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository() *Repository {
	return &Repository{db: db.GetConnection()}
}

func (r *Repository) GetAllProducts() ([]domain.Product, error) {

	var products []domain.Product

	query := `
		SELECT 
			p.id, p.name, p.description, p.price, p.stock_level, p.category_id,
			c.id AS "category.id",
			c.name AS "category.name"
		FROM products p
		JOIN categories c ON c.id = p.category_id
`

	err := r.db.Select(&products, query)

	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *Repository) GetProductByID(id uint) (*domain.Product, error) {
	query := `SELECT 
		p.id, p.name, p.description, p.price, p.stock_level, p.category_id,
		c.id AS "category.id",
		c.name AS "category.name"
	FROM products p
	JOIN categories c ON c.id = p.category_id
	WHERE p.id = $1`

	var product domain.Product

	err := r.db.Get(&product, query, id)

	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) CreateProduct(product domain.Product) (*domain.Product, error) {
	query := `
			INSERT INTO products (name, description, price, stock_level, category_id)
			values (:name, :description, :price, :stock_level, :category_id)
			RETURNING id`

	rows, err := r.db.NamedQuery(query, &product)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&product.ID)
		if err != nil {
			return nil, err
		}
	}

	return r.GetProductByID(product.ID)
}

func (r *Repository) UpdateProduct(product domain.Product) (*domain.Product, error) {
	query := `update products set name = :name,
                    description = :description,
                    price = :price,
                    stock_level = :stock_level,
                    category_id = :category_id
                    where id = :id`

	_, err := r.db.NamedExec(query, &product)
	if err != nil {
		return nil, err
	}

	return r.GetProductByID(product.ID)
}

func (r *Repository) DeleteProduct(id uint) error {
	query := `delete from products where id = $1`
	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}
	return nil
}
