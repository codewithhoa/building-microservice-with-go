package handlers

import (
	"log/slog"
	"net/http"
)

type root struct {
	l *slog.Logger
}

var _ http.Handler = (*root)(nil)

func NewRootHandler(l *slog.Logger) *root {
	return &root{
		l: l,
	}
}

func (r *root) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	r.l.Info("insight root handler")
}
