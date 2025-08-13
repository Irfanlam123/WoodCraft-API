package handlers

import (
	"WoodCraft-API/models"
	"WoodCraft-API/repository"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ServiceHandler struct {
	repo *repository.ServiceRepository
}

// Constructor
func NewServiceHandler(repo *repository.ServiceRepository) *ServiceHandler {
	return &ServiceHandler{repo: repo}
}

// GET /services
func (h *ServiceHandler) GetAllServices(c echo.Context) error {
	services, err := h.repo.GetAllServices()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch services"})
	}
	return c.JSON(http.StatusOK, services)
}

// POST /services
func (h *ServiceHandler) AddService(c echo.Context) error {
	var service models.Service

	// Bind JSON to struct
	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Set timestamps
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	// Add service to DB
	if err := h.repo.AddService(service); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to add service"})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Service added successfully"})
}
