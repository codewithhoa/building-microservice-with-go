package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/codewithhoa/building-microservice-with-go/data"
)

type products struct {
	l *slog.Logger
}

var _ http.Handler = (*products)(nil)

func NewProductsHandler(l *slog.Logger) *products {
	return &products{
		l: l,
	}
}

var (
	ErrMarshalJSON    = errors.New("unable to marshal json")
	ErrGetProducts    = errors.New("can not get list products")
	ErrMethodNotAllow = errors.New("do not supported method")
)

func (p *products) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodGet {
		p.getProducts(rw, rq)
		return
	}

	http.Error(rw, ErrMethodNotAllow.Error(), http.StatusMethodNotAllowed)
}

func (p *products) getProducts(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight get products handler")
	lp, err := data.GetProducts()
	if err != nil {
		http.Error(rw, ErrGetProducts.Error(), http.StatusInternalServerError)
	}

	err = lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, ErrGetProducts.Error(), http.StatusInternalServerError)
	}
}
