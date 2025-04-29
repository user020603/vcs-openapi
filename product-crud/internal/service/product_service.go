package service

import (
	"context"
	"fmt"
	"product-crud/internal/model"
	"product-crud/internal/repository"
	"product-crud/pkg/cache"
)

const (
	productKeyPrefix = "product:"
	allProductsKey   = "products:all"
)

type ProductService struct {
	repo  *repository.ProductRepository
	cache *cache.RedisCache
}

func NewProductService(repo *repository.ProductRepository, cache *cache.RedisCache) *ProductService {
	return &ProductService{
		repo:  repo,
		cache: cache,
	}
}

func (s *ProductService) Create(ctx context.Context, req *model.CreateProductRequest) (*model.ProductResponse, error) {
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

	response := &model.ProductResponse{
		ID:          createdProduct.ID,
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
		Price:       createdProduct.Price,
		CreatedAt:   createdProduct.CreatedAt,
		UpdatedAt:   createdProduct.UpdatedAt,
	}

	err = s.cache.Set(ctx, productKey(id), response)
	if err != nil {
		fmt.Printf("Failed to cache product %d: %v\n", id, err)
	}

	if err := s.cache.Delete(ctx, allProductsKey); err != nil {
		fmt.Printf("Failed to delete all products cache: %v\n", err)
	}

	return response, nil
}

func (s *ProductService) GetByID(ctx context.Context, id int) (*model.ProductResponse, error) {
	var product model.ProductResponse
	found, err := s.cache.Get(ctx, productKey(id), &product)
	if err != nil {
		fmt.Printf("Cache error: %v\n", err)
	}
	
	if found {
		return &product, nil
	}
	
	productFromDB, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	if productFromDB == nil {
		return nil, nil 
	}
	
	response := &model.ProductResponse{
		ID:          productFromDB.ID,
		Name:        productFromDB.Name,
		Description: productFromDB.Description,
		Price:       productFromDB.Price,
		CreatedAt:   productFromDB.CreatedAt,
		UpdatedAt:   productFromDB.UpdatedAt,
	}
	
	if err := s.cache.Set(ctx, productKey(id), response); err != nil {
		fmt.Printf("Error caching product: %v\n", err)
	}
	
	return response, nil
}

func (s *ProductService) GetAll(ctx context.Context) ([]*model.ProductResponse, error) {
	var cachedProducts []*model.ProductResponse
	found, err := s.cache.Get(ctx, allProductsKey, &cachedProducts)
	if err != nil {
		fmt.Printf("Cache error: %v\n", err)
	}
	
	if found && len(cachedProducts) > 0 {
		return cachedProducts, nil
	}
	
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	
	var response []*model.ProductResponse
	for _, product := range products {
		productResponse := &model.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		}
		response = append(response, productResponse)
		
		if err := s.cache.Set(ctx, productKey(product.ID), productResponse); err != nil {
			fmt.Printf("Error caching product: %v\n", err)
		}
	}
	
	if err := s.cache.Set(ctx, allProductsKey, response); err != nil {
		fmt.Printf("Error caching all products: %v\n", err)
	}
	
	return response, nil
}

func (s *ProductService) Update(ctx context.Context, id int, req *model.UpdateProductRequest) (*model.ProductResponse, error) {
	existingProduct, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	if existingProduct == nil {
		return nil, nil 
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
	
	response := &model.ProductResponse{
		ID:          updatedProduct.ID,
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		CreatedAt:   updatedProduct.CreatedAt,
		UpdatedAt:   updatedProduct.UpdatedAt,
	}
	
	if err := s.cache.Set(ctx, productKey(id), response); err != nil {
		fmt.Printf("Error updating product in cache: %v\n", err)
	}
	
	if err := s.cache.Delete(ctx, allProductsKey); err != nil {
		fmt.Printf("Error invalidating all products cache: %v\n", err)
	}
	
	return response, nil
}

func (s *ProductService) Delete(ctx context.Context, id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	
	if err := s.cache.Delete(ctx, productKey(id)); err != nil {
		fmt.Printf("Error removing product from cache: %v\n", err)
	}
	
	if err := s.cache.Delete(ctx, allProductsKey); err != nil {
		fmt.Printf("Error invalidating all products cache: %v\n", err)
	}
	
	return nil
}


func productKey(id int) string {
	return fmt.Sprintf("%s%d", productKeyPrefix, id)
}
