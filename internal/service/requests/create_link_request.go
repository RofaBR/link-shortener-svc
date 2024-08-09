package requests

import (
	"errors"
	"net/url"
)

type CreateLinkRequest struct {
	OriginalURL string `json:"original_url"`
}

func (r CreateLinkRequest) Validate() error {
	if r.OriginalURL == "" {
		return errors.New("original_url is required")
	}

	parsedURL, err := url.ParseRequestURI(r.OriginalURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return errors.New("invalid original_url format")
	}

	return nil
}
