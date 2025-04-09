package handler

import (
	"github.com/gin-gonic/gin"
	"inventory-service/domain"
	"inventory-service/usecase"
	"net/http"
	"strconv"
)

type Handler struct {
	service *usecase.Service
}

func NewHandler() *Handler {
	return &Handler{service: usecase.NewService()}
}

func (h *Handler) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAllProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error during getting all products: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID" + err.Error(),
		})
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Product not found" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	created, err := h.service.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create product: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, created)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID" + err.Error(),
		})
		return
	}

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}
	product.ID = uint(id)

	updated, err := h.service.UpdateProduct(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update product: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID" + err.Error(),
		})
		return
	}

	err = h.service.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
