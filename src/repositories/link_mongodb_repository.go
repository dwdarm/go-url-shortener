package repositories

import (
	"context"
	"time"

	"github.com/dwdarm/go-url-shortener/src/errors"
	"github.com/dwdarm/go-url-shortener/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LinkMongodbInstance struct {
	Slug      string
	Href      string
	QrCode    string             `bson:"qr_code,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty"`
}

type LinkMongodbRepository struct {
	ctx context.Context
	db  *mongo.Database
}

func NewLinkMongodbRepository(ctx context.Context, db *mongo.Database) *LinkMongodbRepository {
	return &LinkMongodbRepository{
		ctx: ctx,
		db:  db,
	}
}

func (repo *LinkMongodbRepository) FindBySlug(slug string) (*models.Link, error) {
	var result LinkMongodbInstance

	err := repo.db.Collection("link").FindOne(repo.ctx, bson.D{{"slug", slug}}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return &models.Link{
		Slug:   result.Slug,
		Href:   result.Href,
		QrCode: result.QrCode,
	}, nil
}

func (repo *LinkMongodbRepository) Save(link *models.Link) (*models.Link, error) {
	newLink := LinkMongodbInstance{
		Slug:      link.Slug,
		Href:      link.Href,
		QrCode:    link.QrCode,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	exist, err := repo.FindBySlug(link.Slug)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return nil, errors.NewErrDuplicate("slug is already exist")
	}

	_, err = repo.db.Collection("link").InsertOne(repo.ctx, newLink)
	if err != nil {

		if mongo.IsDuplicateKeyError(err) {
			return nil, errors.NewErrDuplicate("slug is already exist")
		}

		return nil, err
	}

	return repo.FindBySlug(link.Slug)
}
