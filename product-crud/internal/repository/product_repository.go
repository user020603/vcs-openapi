package repository

import (
	"database/sql"
	"product-crud/internal/model"
	"time"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *model.Product) (int, error) {
	query := `
		INSERT INTO products (name, description, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`

	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	var id int
	err := r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.CreatedAt,
		product.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	query := `SELECT id, name, description, price, created_at, updated_at FROM products WHERE id = $1`

	product := &model.Product{}
	err := r.db.QueryRow(query, id).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Return nil if product not found
		}
		return nil, err
	}

	return product, nil
}

func (r *ProductRepository) GetAll() ([]*model.Product, error) {
	query := `SELECT id, name, description, price, created_at, updated_at FROM products ORDER BY id`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []*model.Product{}
	for rows.Next() {
		product := &model.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) Update(id int, product *model.Product) error {
	query := `
		UPDATE products
		SET name = $1, description = $2, price = $3, updated_at = $4
		WHERE id = $5`

	product.UpdatedAt = time.Now()

	_, err := r.db.Exec(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.UpdatedAt,
		id,
	)

	return err
}

func (r *ProductRepository) Delete(id int) error {
	query := `DELETE FROM products WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}