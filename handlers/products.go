package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/codewithhoa/building-microservice-with-go/data"
	"github.com/gorilla/mux"
)

var (
	ErrMarshalJSON         = errors.New("unable to marshal json")
	ErrUnMarshalJSON       = errors.New("unable to unmarshal json")
	ErrGetProducts         = errors.New("can not get list products")
	ErrMethodNotAllow      = errors.New("do not supported method")
	ErrAddProduct          = errors.New("can not add new product to list products")
	ErrInvalidProductId    = errors.New("invalid product id")
	ErrInvalidProductInput = errors.New("invalid product struct")
)

// Products is a http.Handler
type Products struct {
	l *slog.Logger
}

func NewProductsHandler(l *slog.Logger) *Products {
	return &Products{
		l: l,
	}
}

func (p *Products) GetProducts(rw http.ResponseWriter, rq *http.Request) {
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

func (p *Products) AddProduct(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight post product handler")

	prod, ok := rq.Context().Value(productKey{}).(*data.Product)
	if !ok {
		http.Error(rw, ErrInvalidProductInput.Error(), http.StatusBadRequest)
		return
	}

	// p.l.Debug(fmt.Sprintf("%#v", np))
	err := data.AddProduct(prod)
	if err != nil {
		http.Error(rw, ErrAddProduct.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight put product handler")
	vars := mux.Vars(rq)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, ErrInvalidProductId.Error(), http.StatusBadRequest)
		return
	}

	p.l.Debug("id of product that need to update", slog.Int("id", id))

	prod, ok := rq.Context().Value(productKey{}).(*data.Product)
	if !ok {
		http.Error(rw, ErrInvalidProductInput.Error(), http.StatusBadRequest)
		return
	}

	if err := data.UpdateProductById(id, prod); err != nil {
		if errors.Is(err, data.ErrProductNotFound) {
			http.Error(rw, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

type productKey struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Info("insight middleware")

		// parse data from the request
		prod := &data.Product{}
		err := prod.FromJSON(r.Body)

		p.l.Debug(fmt.Sprintf("type of a is %T\n", prod))
		if err != nil {
			p.l.Error(ErrInvalidProductId.Error(), slog.String("err", err.Error()))
			http.Error(w, ErrUnMarshalJSON.Error(), http.StatusBadRequest)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), productKey{}, prod)
		r = r.WithContext(ctx)

		// call next handler
		next.ServeHTTP(w, r)
	})
}
