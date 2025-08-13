package routes

import (
	"WoodCraft-API/handlers"
	"WoodCraft-API/repository"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	// =======================
	// Booking Setup
	// =======================
	bookingRepo := &repository.BookingRepository{DB: db}
	bookingHandler := handlers.NewBookingHandler(bookingRepo)

	// =======================
	// Services Setup
	// =======================
	serviceRepo := &repository.ServiceRepository{DB: db}
	serviceHandler := handlers.NewServiceHandler(serviceRepo)

	// =======================
	// Booking Routes
	// =======================
	e.POST("/api/bookings", bookingHandler.CreateBooking)     // Form submission
	e.GET("/api/bookings", bookingHandler.GetAllBookings)     // Fetch all bookings
	e.GET("/api/bookings/:id", bookingHandler.GetBookingByID) // Get data from ID

	// =======================
	// Services Routes
	// =======================
	e.GET("/api/services", serviceHandler.GetAllServices)
	e.POST("/api/services", serviceHandler.AddService)

	// =======================
	// About / Contact Routes (TODO)
	// =======================

	// =======================
	// Gallery Routes (TODO)
	// =======================
}
