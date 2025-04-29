package rest

import (
	"context"
	"net/http"
	"product-crud/internal/model"
	"product-crud/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided information
// @Tags products
// @Accept json
// @Produce json
// @Param product body model.CreateProductRequest true "Product information"
// @Success 201 {object} model.ProductResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var req model.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.Create(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} model.ProductResponse
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.service.GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GetProducts godoc
// @Summary Get all products
// @Description Get a list of all products
// @Tags products
// @Accept json
// @Produce json
// @Success 200 {array} model.ProductResponse
// @Failure 500 {string} string "Internal Server Error"
// @Router /products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.service.GetAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product with the provided information
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body model.UpdateProductRequest true "Product information"
// @Success 200 {object} model.ProductResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var req model.UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := h.service.Update(context.Background(), id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = h.service.Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}