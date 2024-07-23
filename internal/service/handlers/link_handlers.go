package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/RofaBR/link-shortener-svc/internal/config"
	/*
		"github.com/RofaBR/link-shortener-svc/internal/service/requests"
		"github.com/RofaBR/link-shortener-svc/internal/config"
		"gitlab.com/distributed_lab/ape"
		"gitlab.com/distributed_lab/logan/v3/errors"
		"github.com/jmoiron/sqlx"
		"github.com/teris-io/shortid"
	*/)

type Link struct {
	ID          int    `db:"id"`
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
	CreatedAt   string `db:"created_at"`
}

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
	db := cfg.DB().RawDB()

	shortURL, err := shortid.Generate()
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	link := Link{
		OriginalURL: request.OriginalURL,
		ShortURL:    shortURL,
	}

	_, err = db.NamedExec(`INSERT INTO links (original_url, short_url) VALUES (:original_url, :short_url)`, &link)
	if err != nil {
		ape.RenderErr(w, problems.InternalError())
		return
	}

	ape.Render(w, link)
}

func GetLink(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "shortURL")

	cfg := config.New(r.Context().Value("config").(kv.Getter))
	db := cfg.DB().RawDB()

	var link Link
	err := db.Get(&link, `SELECT * FROM links WHERE short_url=$1`, shortURL)
	if err != nil {
		ape.RenderErr(w, problems.NotFound())
		return
	}

	http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}
