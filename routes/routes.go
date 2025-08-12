package routes

import (
	"WoodCraft-API/handlers"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	// Initialize Handlers
	bookingHandler := &handlers.BookingHandler{DB: db}
	// serviceHandler := &handlers.ServiceHandler{DB: db}
	// aboutHandler := &handlers.AboutHandler{DB: db}
	// galleryHandler := &handlers.GalleryHandler{DB: db}

	// =======================
	// Booking Routes
	// =======================
	e.POST("/api/bookings", bookingHandler.CreateBooking) // Form submission
	e.GET("/api/bookings", bookingHandler.GetAllBookings) // Fetch all bookings

	// =======================
	// Services Routes
	// =======================
	// e.POST("/api/services", serviceHandler.CreateService)
	// e.GET("/api/services", serviceHandler.GetAllServices)

	// =======================
	// About / Contact Routes
	// =======================
	// e.POST("/api/about", aboutHandler.CreateAbout)
	// e.GET("/api/about", aboutHandler.GetAbout)

	// =======================
	// Gallery Routes
	// =======================
	// e.POST("/api/gallery", galleryHandler.CreateGalleryItem)
	// e.GET("/api/gallery", galleryHandler.GetGalleryItems)
}
