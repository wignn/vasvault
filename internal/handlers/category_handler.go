package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"vasvault/internal/services"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

type CreateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}

// POST /categories
func (h *CategoryHandler) Create(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")

	category, err := h.service.Create(req.Name, req.Color, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// GET /categories
func (h *CategoryHandler) List(c *gin.Context) {
	userID := c.GetUint("userID")
	search := c.Query("search")

	categories, err := h.service.List(userID, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GET /categories/:id
func (h *CategoryHandler) Detail(c *gin.Context) {
	userID := c.GetUint("userID")
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	category, err := h.service.Detail(userID, uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

//update category struct
type UpdateCategoryRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color"`
}

func (h *CategoryHandler) Update(c *gin.Context) {
	var req UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID") // PENTING: sama dengan middleware

	idParam := c.Param("id")
	categoryID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	category, err := h.service.Update(
		userID,
		uint(categoryID),
		req.Name,
		req.Color,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

//delete endpoint
func (h *CategoryHandler) Delete(c *gin.Context) {
	userID := c.GetUint("userID")

	idParam := c.Param("id")
	categoryID, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category ID"})
		return
	}

	err = h.service.Delete(userID, uint(categoryID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "category deleted successfully",
	})
}

