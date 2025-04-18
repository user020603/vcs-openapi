package service

import (
	"product-crud/internal/model"
	"product-crud/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (s *ProductService) Create(req *model.CreateProductRequest) (*model.ProductResponse, error) {
	product := &model.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	id, err := s.repo.Create(product)
	if err != nil {
		return nil, err
	}

	createdProduct, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &model.ProductResponse{
		ID:          createdProduct.ID,
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
		CreatedAt:   createdProduct.CreatedAt,
		UpdatedAt:   createdProduct.UpdatedAt,
	}, nil
}

func (s *ProductService) GetByID(id int) (*model.ProductResponse, error) {
	product, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, nil // Product not found
	}

	return &model.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, nil
}

func (s *ProductService) GetAll() ([]*model.ProductResponse, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []*model.ProductResponse
	for _, product := range products {
		response = append(response, &model.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	return response, nil
}

func (s *ProductService) Update(id int, req *model.UpdateProductRequest) (*model.ProductResponse, error) {
	existingProduct, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if existingProduct == nil {
		return nil, nil // Product not found
	}

	if req.Name != "" {
		existingProduct.Name = req.Name
	}
	if req.Description != "" {
		existingProduct.Description = req.Description
	}
	if req.Price > 0 {
		existingProduct.Price = req.Price
	}

	err = s.repo.Update(id, existingProduct)
	if err != nil {
		return nil, err
	}

	updatedProduct, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &model.ProductResponse{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		CreatedAt:   updatedProduct.CreatedAt,
		UpdatedAt:   updatedProduct.UpdatedAt,
	}, nil
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}