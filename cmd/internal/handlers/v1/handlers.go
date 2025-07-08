package v1

import (
	"clothes-shop-backend/cmd/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1Handlers struct {
	productService productService
}

func NewV1Handlers(productService productService) *V1Handlers {
	return &V1Handlers{
		productService: productService,
	}
}

func (h *V1Handlers) GetPaginatedProducts(c *gin.Context) {
	type Pagination struct {
		Page  int `form:"page" binding:"min=1"`
		Limit int `form:"limit" binding:"min=1,max=100"`
	}

	var p Pagination
	err := c.ShouldBindQuery(&p)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters. Page and limit must be positive integers, limit max 100"})
		return
	}

	if p.Page == 0 {
		p.Page = 1
	}
	if p.Limit == 0 {
		p.Limit = 10
	}

	products, err := h.productService.GetPaginatedProducts(p.Page, p.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *V1Handlers) UploadProduct(c *gin.Context) {
	var req models.UploadProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	product, err := h.productService.UploadProduct(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *V1Handlers) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	product, err := h.productService.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}
