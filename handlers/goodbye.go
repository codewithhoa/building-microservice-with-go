package handlers

import (
	"log/slog"
	"net/http"
)

type goodbye struct {
	l *slog.Logger
}

var _ http.Handler = (*goodbye)(nil)

func NewGoodbyeHandler(l *slog.Logger) *goodbye {
	return &goodbye{
		l: l,
	}
}

func (g *goodbye) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	g.l.Info("insight goodbye handler")
	return
}
