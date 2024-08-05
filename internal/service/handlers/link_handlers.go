package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RofaBR/link-shortener-svc/internal/service/requests"
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateLink(w http.ResponseWriter, r *http.Request) {
	var request requests.CreateLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	if err := request.Validate(); err != nil {
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	storage := LinkStorageFromContext(r.Context())

	shortURL, err := storage.CreateLink(request.OriginalURL)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, map[string]string{"short_url": shortURL})
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")
	storage := LinkStorageFromContext(r.Context())

	originalURL, err := storage.GetLink(shortURL)
	if err != nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
