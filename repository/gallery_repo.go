package repository

import (
	"WoodCraft-API/models"
	"database/sql"
)

type GalleryRepository struct {
	DB *sql.DB
}

func (r *GalleryRepository) GetAllGalleryItems() ([]models.Gallery, error) {
	rows, err := r.DB.Query("SELECT id, project_name, description, image_url, completion_date, material_used FROM gallery")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var galleries []models.Gallery
	for rows.Next() {
		var g models.Gallery
		if err := rows.Scan(&g.ID, &g.ProjectName, &g.Description, &g.ImageURL, &g.CompletionDate, &g.MaterialUsed); err != nil {
			return nil, err
		}
		galleries = append(galleries, g)
	}
	return galleries, nil
}

func InsertGallery(db *sql.DB, g models.Gallery) error {
	_, err := db.Exec(`
        INSERT INTO gallery (project_name, description, image_url, completion_date, material_used)
        VALUES ($1, $2, $3, $4, $5)`,
		g.ProjectName, g.Description, g.ImageURL, g.CompletionDate, g.MaterialUsed,
	)
	return err
}
