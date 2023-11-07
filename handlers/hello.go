package handlers

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
)

type hello struct {
	l *slog.Logger
}

var _ http.Handler = (*hello)(nil)

func NewHelloHandler(l *slog.Logger) *hello {
	return &hello{
		l: l,
	}
}

func (h *hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Info("insight hello handler")
	body, err := io.ReadAll(r.Body)
	// err = errors.New("something went wrong")
	if err != nil {
		// w.Write([]byte(err.Error()))
		// w.WriteHeader(http.StatusInternalServerError)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(string(body))
	// fmt.Println(r.UserAgent())
	res := fmt.Sprintf("Hello %s", string(body))

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(res))
	return
}
