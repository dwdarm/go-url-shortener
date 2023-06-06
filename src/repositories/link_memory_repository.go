package repositories

import (
	"sync"

	"github.com/dwdarm/go-url-shortener/src/errors"
	"github.com/dwdarm/go-url-shortener/src/models"
)

type LinkMemoryInstance struct {
	Slug   string
	Href   string
	QrCode string
}

type LinkMemoryRepository struct {
	mtx   *sync.Mutex
	links map[string]*LinkMemoryInstance
}

func NewLinkMemoryRepository() *LinkMemoryRepository {
	return &LinkMemoryRepository{
		mtx:   &sync.Mutex{},
		links: map[string]*LinkMemoryInstance{},
	}
}

func (repo *LinkMemoryRepository) FindBySlug(slug string) (*models.Link, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	data, exist := repo.links[slug]
	if !exist {
		return nil, nil
	}

	return models.NewLink(data.Slug, data.Href, data.QrCode)
}

func (repo *LinkMemoryRepository) Save(link *models.Link) (*models.Link, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	if repo.links[link.Slug] != nil {
		return nil, errors.NewErrDuplicate("slug is already exist")
	}

	repo.links[link.Slug] = &LinkMemoryInstance{
		Slug:   link.Slug,
		Href:   link.Href,
		QrCode: link.QrCode,
	}

	return link, nil
}
