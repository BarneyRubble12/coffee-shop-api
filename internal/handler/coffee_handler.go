package handler

import (
	"net/http"
	"strconv"

	"coffee-shop-api/internal/model"
	"coffee-shop-api/internal/service"

	"github.com/gin-gonic/gin"
)

// CoffeeHandler handles HTTP requests for coffee operations
type CoffeeHandler struct {
	service *service.CoffeeService
}

// NewCoffeeHandler creates a new instance of CoffeeHandler
func NewCoffeeHandler(service *service.CoffeeService) *CoffeeHandler {
	return &CoffeeHandler{
		service: service,
	}
}

// @Summary Get all coffees
// @Description Get a list of all coffees
// @Tags coffees
// @Accept json
// @Produce json
// @Success 200 {array} model.Coffee
// @Router /coffees [get]
func (h *CoffeeHandler) GetAllCoffees(c *gin.Context) {
	coffees := h.service.GetAllCoffees()
	c.JSON(http.StatusOK, coffees)
}

// @Summary Get a coffee by ID
// @Description Get a coffee by its ID
// @Tags coffees
// @Accept json
// @Produce json
// @Param id path int true "Coffee ID"
// @Success 200 {object} model.Coffee
// @Failure 404 {object} map[string]string
// @Router /coffees/{id} [get]
func (h *CoffeeHandler) GetCoffeeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	coffee, err := h.service.GetCoffeeByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coffee)
}

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

	createdCoffee := h.service.CreateCoffee(coffee)
	c.JSON(http.StatusCreated, createdCoffee)
}

// @Summary Update a coffee
// @Description Update an existing coffee by ID
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var coffee model.Coffee
	if err := c.ShouldBindJSON(&coffee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCoffee, err := h.service.UpdateCoffee(id, coffee)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCoffee)
}

// @Summary Delete a coffee
// @Description Delete a coffee by ID
// @Tags coffees
// @Accept json
// @Produce json
// @Param id path int true "Coffee ID"
// @Success 204 "No Content"
// @Failure 404 {object} map[string]string
// @Router /coffees/{id} [delete]
func (h *CoffeeHandler) DeleteCoffee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.DeleteCoffee(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Search coffees by origin
// @Description Get all coffees from a specific origin
// @Tags coffees
// @Accept json
// @Produce json
// @Param origin query string true "Coffee origin"
// @Success 200 {array} model.Coffee
// @Failure 400 {object} map[string]string
// @Router /coffees/search [get]
func (h *CoffeeHandler) SearchCoffeesByOrigin(c *gin.Context) {
	origin := c.Query("origin")
	if origin == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "origin parameter is required"})
		return
	}

	coffees := h.service.GetAllCoffees()
	var filteredCoffees []model.Coffee
	for _, coffee := range coffees {
		if coffee.Origin == origin {
			filteredCoffees = append(filteredCoffees, coffee)
		}
	}

	c.JSON(http.StatusOK, filteredCoffees)
}
