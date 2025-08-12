package handlers

import (
	"WoodCraft-API/models"
	"WoodCraft-API/repository"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	repo *repository.BookingRepository
}

func NewBookingHandler(repo *repository.BookingRepository) *BookingHandler {
	return &BookingHandler{repo: repo}
}

// CreateBooking - POST /bookings
func (h *BookingHandler) CreateBooking(c echo.Context) error {
	var booking models.Booking

	if err := c.Bind(&booking); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	booking.Status = "Pending"
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	err := h.repo.CreateBooking(&booking)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create booking"})
	}

	return c.JSON(http.StatusCreated, booking)
}

// GetAllBookings - GET /bookings
func (h *BookingHandler) GetAllBookings(c echo.Context) error {
	bookings, err := h.repo.GetAllBookings()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch bookings"})
	}
	return c.JSON(http.StatusOK, bookings)
}

// GetBookingByID - GET /bookings/:id
func (h *BookingHandler) GetBookingByID(c echo.Context) error {
	id := c.Param("id")

	booking, err := h.repo.GetBookingByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Booking not found"})
	}

	return c.JSON(http.StatusOK, booking)
}
