package v1

import (
	"clothes-shop-backend/internal/models"
	"clothes-shop-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type V1Handlers struct {
	productService productService
	userService    userService
	cartService    cartService
}

func NewV1Handlers(productService productService, userService userService, cartService cartService) *V1Handlers {
	return &V1Handlers{
		productService: productService,
		userService:    userService,
		cartService:    cartService,
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

func (h *V1Handlers) GetUserByPhone(c *gin.Context) {
	phone := c.Param("phone")

	user, err := h.userService.GetUserByPhone(phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	tokenString, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("token", tokenString, 0, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"user": user, "token": tokenString})
}

func (h *V1Handlers) CreateUser(c *gin.Context) {
	var req models.User
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	user, err := h.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tokenString, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user, "token": tokenString})
}

func (h *V1Handlers) AddOneToCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "User ID not found in context"})
		return
	}

	productID := c.Param("product_id")

	err := h.cartService.AddOneToCart(userID.(string), productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart"})
}

func (h *V1Handlers) UpdateCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "User ID not found in context"})
		return
	}

	var req []*models.CartItemPrototype
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	err := h.cartService.UpdateCart(userID.(string), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated"})
}

func (h *V1Handlers) GetCartItems(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusForbidden, gin.H{"error": "User ID not found in context"})
		return
	}

	cartItems, err := h.cartService.GetCartItemsByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cartItems)
}

func (h *V1Handlers) Checkout(c *gin.Context) {
	userID := c.Param("user_id")

	err := h.cartService.Checkout(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Checkout successful"})
}
