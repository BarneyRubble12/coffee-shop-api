package handler

import (
	"net/http"
	"strconv"

	"coffee-shop-api/internal/model"
	"coffee-shop-api/internal/repository"
	"coffee-shop-api/internal/service"

	"github.com/gin-gonic/gin"
)

// CoffeeHandler handles HTTP requests for coffee operations
type CoffeeHandler struct {
	service *service.CoffeeService
}

// NewCoffeeHandler creates a new coffee handler instance
func NewCoffeeHandler(service *service.CoffeeService) *CoffeeHandler {
	return &CoffeeHandler{service: service}
}

// GetAllCoffees handles GET /coffees
// @Summary Get all coffees
// @Description Get a list of all available coffees
// @Tags coffees
// @Accept json
// @Produce json
// @Success 200 {array} model.Coffee
// @Router /coffees [get]
func (h *CoffeeHandler) GetAllCoffees(c *gin.Context) {
	coffees, err := h.service.GetAllCoffees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffees)
}

// GetCoffeeByID handles GET /coffees/:id
// @Summary Get a coffee by ID
// @Description Get a specific coffee by its ID
// @Tags coffees
// @Accept json
// @Produce json
// @Param id path int true "Coffee ID"
// @Success 200 {object} model.Coffee
// @Failure 404 {object} map[string]string
// @Router /coffees/{id} [get]
func (h *CoffeeHandler) GetCoffeeByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	coffee, err := h.service.GetCoffeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffee)
}

// CreateCoffee handles POST /coffees
// @Summary Create a new coffee
// @Description Create a new coffee product
// @Tags coffees
// @Accept json
// @Produce json
// @Param coffee body model.Coffee true "Coffee object"
// @Success 201 {object} model.Coffee
// @Failure 400 {object} map[string]string
// @Router /coffees [post]
func (h *CoffeeHandler) CreateCoffee(c *gin.Context) {
	var coffee model.Coffee
	if err := c.ShouldBindJSON(&coffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateCoffee(&coffee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, coffee)
}

// UpdateCoffee handles PUT /coffees/:id
// @Summary Update a coffee
// @Description Update an existing coffee product
// @Tags coffees
// @Accept json
// @Produce json
// @Param id path int true "Coffee ID"
// @Param coffee body model.Coffee true "Coffee object"
// @Success 200 {object} model.Coffee
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /coffees/{id} [put]
func (h *CoffeeHandler) UpdateCoffee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var coffee model.Coffee
	if err := c.ShouldBindJSON(&coffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coffee.ID = id
	if err := h.service.UpdateCoffee(&coffee); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coffee)
}

// DeleteCoffee handles DELETE /coffees/:id
// @Summary Delete a coffee
// @Description Delete a coffee product by its ID
// @Tags coffees
// @Accept json
// @Produce json
// @Param id path int true "Coffee ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /coffees/{id} [delete]
func (h *CoffeeHandler) DeleteCoffee(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	if err := h.service.DeleteCoffee(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// GetRepository returns the repository instance for cleanup
func (h *CoffeeHandler) GetRepository() repository.CoffeeRepository {
	return h.service.GetRepository()
}
