package services

import (
	"context"
	"testing"

	"github.com/dwdarm/go-url-shortener/src/repositories"
)

func createService() LinkService {
	repo := repositories.NewLinkMemoryRepository()
	srv := NewLinkService(repo)

	return srv
}

func TestCreateRandomSlug(t *testing.T) {
	href := "http://example.com"
	slugLength := 6

	srv := createService()

	link, err := srv.CreateLink(context.TODO(), "", href)

	if err != nil {
		t.Errorf("Expected nil on error but got one")
	}

	if link != nil {
		if href != link.Href {
			t.Errorf("Expected '%s', but get '%s", href, link.Href)
		}

		if len(link.Slug) != slugLength {
			t.Errorf("Expected slug length '%d', but get '%d", slugLength, len(link.Slug))
		}
	}
}

func TestCreateCustomSlug(t *testing.T) {
	slug := "hello world"
	href := "http://example.com"

	srv := createService()

	link, err := srv.CreateLink(context.TODO(), slug, href)

	if err != nil {
		t.Errorf("Expected not nil on error but got one")
	}

	if link != nil {
		if href != link.Href {
			t.Errorf("Expected '%s', but get '%s", href, link.Href)
		}
	}
}

func TestFailOnExistingSlug(t *testing.T) {
	slug := "hello world"
	href := "http://example.com"

	srv := createService()

	srv.CreateLink(context.TODO(), slug, href)

	link, err := srv.CreateLink(context.TODO(), slug, href)

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if link != nil {
		t.Errorf("Expected nil on link but got one")
	}
}

func TestInvalidHref(t *testing.T) {
	slug := "hello world"
	href := "test"

	srv := createService()

	link, err := srv.CreateLink(context.TODO(), slug, href)

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if link != nil {
		t.Errorf("Expected nil on link but got one")
	}
}

func TestGetLink(t *testing.T) {
	href := "http://example.com"

	srv := createService()
	f, _ := srv.CreateLink(context.TODO(), "", href)

	link, err := srv.GetLink(context.TODO(), f.Slug)

	if err != nil {
		t.Errorf("Expected not nil on error but got one")
	}

	if link != nil {
		if link.Slug != f.Slug {
			t.Errorf("Expected '%s', but get '%s", f.Slug, link.Slug)
		}
		if link.Href != f.Href {
			t.Errorf("Expected '%s', but get '%s", f.Href, link.Href)
		}
	}
}

func TestNonExistingLink(t *testing.T) {
	srv := createService()
	srv.CreateLink(context.TODO(), "hello world", "http://example.com")

	link, err := srv.GetLink(context.TODO(), "hello world")

	if err != nil {
		t.Errorf("Expected nil on error but got one")
	}

	if link != nil {
		t.Errorf("Expected nil on link but got one")
	}
}
