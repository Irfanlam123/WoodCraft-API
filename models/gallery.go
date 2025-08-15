package models

import "time"

type Gallery struct {
	ID             int       `json:"id" db:"id"`
	ProjectName    string    `json:"project_name" db:"project_name"`
	Description    string    `json:"description" db:"description"`
	ImageURL       string    `json:"image_url" db:"image_url"`
	CompletionDate time.Time `json:"completion_date" db:"completion_date"`
	MaterialUsed   string    `json:"material_used" db:"material_used"`
}
