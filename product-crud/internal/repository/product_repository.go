package repository

import (
	"product-crud/internal/model"
	"time"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) Create(product *model.Product) (int, error) {
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	result := r.db.Create(product)
	if result.Error != nil {
		return 0, result.Error
	}

	return product.ID, nil
}

func (r *ProductRepository) GetByID(id int) (*model.Product, error) {
	product := &model.Product{}
	result := r.db.First(product, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return product, nil
}

func (r *ProductRepository) GetAll() ([]*model.Product, error) {
	var products []*model.Product
	result := r.db.Order("id").Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *ProductRepository) Update(id int, product *model.Product) error {
	product.UpdatedAt = time.Now()

	result := r.db.Model(&model.Product{ID: id}).Updates(map[string]interface{}{
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"updated_at":  product.UpdatedAt,
	})

	return result.Error
}

func (r *ProductRepository) Delete(id int) error {
	result := r.db.Delete(&model.Product{}, id)
	return result.Error
}
