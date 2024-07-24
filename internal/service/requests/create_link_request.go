package requests

import "errors"

type CreateLinkRequest struct {
	OriginalURL string `json:"original_url"`
}

func (r CreateLinkRequest) Validate() error {
	if r.OriginalURL == "" {
		return errors.New("original_url is required")
	}
	return nil
}
