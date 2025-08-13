package repository

import (
	"WoodCraft-API/models"
	"database/sql"
	"time"
)

type ServiceRepository struct {
	DB *sql.DB
}

// GetAllServices - fetch all services with created_at & updated_at
func (r *ServiceRepository) GetAllServices() ([]models.Service, error) {
	rows, err := r.DB.Query(`
		SELECT id, name, description, price_range, image_url, category, created_at, updated_at
		FROM services
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var services []models.Service
	for rows.Next() {
		var service models.Service
		if err := rows.Scan(
			&service.ID,
			&service.Name,
			&service.Description,
			&service.PriceRange,
			&service.ImageURL,
			&service.Category,
			&service.CreatedAt,
			&service.UpdatedAt,
		); err != nil {
			return nil, err
		}
		services = append(services, service)
	}

	return services, nil
}

// AddService - insert a new service
func (r *ServiceRepository) AddService(service models.Service) error {
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	_, err := r.DB.Exec(
		`INSERT INTO services (name, description, price_range, image_url, category, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		service.Name,
		service.Description,
		service.PriceRange,
		service.ImageURL,
		service.Category,
		service.CreatedAt,
		service.UpdatedAt,
	)
	return err
}
