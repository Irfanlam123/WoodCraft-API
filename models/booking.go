package models

import "time"

type Booking struct {
	ID          int       `json:"id" db:"id" gorm:"primaryKey;autoIncrement"`
	Name        string    `json:"name" db:"name" binding:"required"`                 // Customer Name
	ContactInfo string    `json:"contact_info" db:"contact_info" binding:"required"` // Phone or Email
	Address     string    `json:"address" db:"address" binding:"required"`           // Customer Address
	ServiceID   int       `json:"service_id" db:"service_id" binding:"required"`     // Foreign key -> Services Table
	DateTime    time.Time `json:"date_time" db:"date_time" binding:"required"`       // Booking date & time
	Notes       string    `json:"notes,omitempty" db:"notes"`                        // Optional design notes
	Status      string    `json:"status" db:"status" gorm:"default:'Pending'"`       // Booking status
	CreatedAt   time.Time `json:"created_at" db:"created_at" gorm:"autoCreateTime"`  // Record created time
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at" gorm:"autoUpdateTime"`  // Last update time
}
