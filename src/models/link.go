package models

import "github.com/asaskevich/govalidator"

type Link struct {
	Slug   string `valid:"required"`
	Href   string `valid:"url,required"`
	QrCode string `valid:"optional"`
}

func NewLink(slug string, href string, qrCode string) (*Link, error) {
	link := &Link{
		Slug:   slug,
		Href:   href,
		QrCode: qrCode,
	}

	_, err := govalidator.ValidateStruct(link)
	if err != nil {
		return nil, err
	}

	return link, nil
}
