package handlers

import (
	"WoodCraft-API/models"
	"WoodCraft-API/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GalleryHandler struct {
	Repo *repository.GalleryRepository
}

func NewGalleryHandler(repo *repository.GalleryRepository) *GalleryHandler {
	return &GalleryHandler{Repo: repo}
}

func (h *GalleryHandler) GetGallery(c echo.Context) error {
	items, err := h.Repo.GetAllGalleryItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch gallery",
		})
	}
	return c.JSON(http.StatusOK, items)
}

func (h *GalleryHandler) CreateGallery(c echo.Context) error {
	var g models.Gallery
	if err := c.Bind(&g); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid input",
		})
	}

	err := repository.InsertGallery(h.Repo.DB, g)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to insert gallery item",
		})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "Gallery item created successfully",
	})
}
