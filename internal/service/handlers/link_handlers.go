package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RofaBR/link-shortener-svc/internal/config"
	"github.com/RofaBR/link-shortener-svc/internal/service/requests"
	"github.com/RofaBR/link-shortener-svc/internal/storage"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"gitlab.com/distributed_lab/kit/kv"

	"github.com/teris-io/shortid"
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

	cfg := config.New(r.Context().Value("config").(kv.Getter))
	db := sqlx.NewDb(cfg.DB().RawDB(), "postgres")
	store := storage.New(db)

	shortURL, err := shortid.Generate()
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	link := &storage.Link{
		OriginalURL: request.OriginalURL,
		ShortURL:    shortURL,
	}

	err = store.CreateLink(link)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, link)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	cfg := config.New(r.Context().Value("config").(kv.Getter))
	db := sqlx.NewDb(cfg.DB().RawDB(), "postgres")
	store := storage.New(db)

	link, err := store.GetLinkByShortURL(shortURL)
	if err != nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}
