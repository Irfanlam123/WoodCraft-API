// # Handlers for booking endpoints
package handlers

import (
	"WoodCraft-API/models"
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type BookingHandler struct {
	DB *sql.DB
}

// CreateBooking - POST /bookings
func (h *BookingHandler) CreateBooking(c echo.Context) error {
	var booking models.Booking

	// Bind JSON body to struct
	if err := c.Bind(&booking); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	// Set default values
	booking.Status = "Pending"
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	// Insert into database
	query := `
		INSERT INTO bookings (name, contact_info, address, service_id, date_time, notes, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
	`
	err := h.DB.QueryRow(
		query,
		booking.Name,
		booking.ContactInfo,
		booking.Address,
		booking.ServiceID,
		booking.DateTime,
		booking.Notes,
		booking.Status,
		booking.CreatedAt,
		booking.UpdatedAt,
	).Scan(&booking.ID)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create booking"})
	}

	return c.JSON(http.StatusCreated, booking)
}

// GetAllBookings - GET /bookings
func (h *BookingHandler) GetAllBookings(c echo.Context) error {
	rows, err := h.DB.Query("SELECT id, name, contact_info, address, service_id, date_time, notes, status, created_at, updated_at FROM bookings")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch bookings"})
	}
	defer rows.Close()

	var bookings []models.Booking

	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(
			&b.ID, &b.Name, &b.ContactInfo, &b.Address, &b.ServiceID,
			&b.DateTime, &b.Notes, &b.Status, &b.CreatedAt, &b.UpdatedAt,
		); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to parse bookings"})
		}
		bookings = append(bookings, b)
	}

	return c.JSON(http.StatusOK, bookings)
}
