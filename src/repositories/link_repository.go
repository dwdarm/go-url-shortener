package repositories

import "github.com/dwdarm/go-url-shortener/src/models"

type LinkRepository interface {
	FindBySlug(slug string) (*models.Link, error)
	Save(link *models.Link) (*models.Link, error)
}
