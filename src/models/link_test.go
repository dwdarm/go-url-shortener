package models

import (
	"testing"
)

func TestValidLinkModel(t *testing.T) {
	slug := "test"
	href := "http://example.com"

	link, err := NewLink(slug, href, "")

	if err != nil {
		t.Errorf("Expected nil on error but get error")
	}

	if link != nil {
		if link.Slug != slug {
			t.Errorf("Expected '%s' but get '%s", slug, link.Slug)
		}

		if link.Href != href {
			t.Errorf("Expected '%s' but get '%s", slug, link.Slug)
		}
	}
}

func TestInvalidHref(t *testing.T) {
	slug := "test"
	href := "testtest"

	link, err := NewLink(slug, href, "")

	if err == nil {
		t.Errorf("Expected error but got nil")
	}

	if link != nil {
		t.Errorf("Expected nil on link but got not nil")
	}
}
