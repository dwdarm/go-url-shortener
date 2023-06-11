package services

import (
	"context"
	"encoding/base64"

	"github.com/dwdarm/go-url-shortener/src/models"
	"github.com/dwdarm/go-url-shortener/src/repositories"
	"github.com/gosimple/slug"
	"github.com/skip2/go-qrcode"
	"go.step.sm/crypto/randutil"
)

type LinkService interface {
	GetLink(ctx context.Context, sg string) (*models.Link, error)
	CreateLink(ctx context.Context, sg string, href string) (*models.Link, error)
}

type LinkServiceImp struct {
	repo repositories.LinkRepository
}

func NewLinkService(repo repositories.LinkRepository) LinkService {
	return &LinkServiceImp{
		repo: repo,
	}
}

func (s *LinkServiceImp) GetLink(ctx context.Context, sg string) (*models.Link, error) {
	link, err := s.repo.FindBySlug(ctx, sg)

	return link, err
}

func generateQrCode(data string) string {
	var img []byte
	img, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		return ""
	}

	return "data: image/gif;base64, " + base64.StdEncoding.EncodeToString(img)
}

func (s *LinkServiceImp) createUniqueLink(ctx context.Context, href string) (*models.Link, error) {
	for {

		sg, err := randutil.Alphanumeric(6)
		if err != nil {
			return nil, err
		}

		exist, err := s.GetLink(ctx, sg)
		if err != nil {
			return nil, err
		}

		if exist == nil {
			in, err := models.NewLink(sg, href, generateQrCode(href))
			if err != nil {
				return nil, err
			}

			link, err := s.repo.Save(ctx, in)
			if err != nil {
				return nil, err
			}

			return link, err
		}
	}
}

func (s *LinkServiceImp) createCustomLink(ctx context.Context, sg string, href string) (*models.Link, error) {
	in, err := models.NewLink(slug.Make(sg), href, generateQrCode(href))
	if err != nil {
		return nil, err
	}

	link, err := s.repo.Save(ctx, in)

	return link, err
}

func (s *LinkServiceImp) CreateLink(ctx context.Context, sg string, href string) (*models.Link, error) {
	if len(sg) < 6 {
		return s.createUniqueLink(ctx, href)
	} else {
		return s.createCustomLink(ctx, sg, href)
	}
}
