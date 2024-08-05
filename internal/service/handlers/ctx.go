package handlers

import (
	"context"
	"net/http"

	"github.com/RofaBR/link-shortener-svc/internal/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	CtxLinkStorage
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxWithLinkStorage(ctx context.Context, storage data.LinkStorage) context.Context {
	return context.WithValue(ctx, CtxLinkStorage, storage)
}

func LinkStorageFromContext(ctx context.Context) data.LinkStorage {
	return ctx.Value(CtxLinkStorage).(data.LinkStorage)
}
