// DB queries for booking
package repository

import (
	"WoodCraft-API/models"
	"database/sql"
	"time"
)

type BookingRepository struct {
	DB *sql.DB
}

// CreateBooking inserts a new booking into the database
func (r *BookingRepository) CreateBooking(booking *models.Booking) error {
	booking.Status = "Pending"
	booking.CreatedAt = time.Now()
	booking.UpdatedAt = time.Now()

	query := `
		INSERT INTO bookings (name, contact_info, address, service_id, date_time, notes, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id
	`

	return r.DB.QueryRow(
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
}

// GetAllBookings fetches all bookings from the database
func (r *BookingRepository) GetAllBookings() ([]models.Booking, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, contact_info, address, service_id, date_time, notes, status, created_at, updated_at
		FROM bookings
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking

	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(
			&b.ID,
			&b.Name,
			&b.ContactInfo,
			&b.Address,
			&b.ServiceID,
			&b.DateTime,
			&b.Notes,
			&b.Status,
			&b.CreatedAt,
			&b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}

	return bookings, nil
}
