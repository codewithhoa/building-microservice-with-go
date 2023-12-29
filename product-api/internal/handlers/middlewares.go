package handlers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
	"github.com/codewithhoa/building-microservice-with-go/product-api/internal/response"
)

var (
	ErrDeserializingProduct = errors.New("error deserializing product")
)

type productKey struct{}

func (p *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p.l.Info("insight middleware")

		// parse data from the request
		prod := &data.Product{}
		err := data.FromJSON(prod, r.Body)
		// p.l.Debug(fmt.Sprintf("type of a is %T\n", prod))
		if err != nil {
			p.l.Error(ErrDeserializingProduct.Error(), slog.String("err", err.Error()))
			p.responseJSON(w, http.StatusBadRequest, response.SimpleErrorResponse(ErrUnMarshalJSON, ErrUnMarshalJSONCode))
			return
		}

		// validate product
		if errs := p.v.Validate(prod); len(errs) != 0 {
			fmt.Println("[ERROR] deserializing product", errs)
			p.l.Error(ErrInvalidProductInput.Error(), slog.Any("err", errs.Errors()))
			var errFieldList []response.ErrorField

			for _, v := range errs {
				errField := response.ErrorField{
					Error: v.Error(),
					Field: v.Field(),
				}
				errFieldList = append(errFieldList, errField)
			}
			errRes := response.FullErrorResponse(ErrInvalidProductInput, ErrInvalidProductInputCode, errFieldList)
			p.responseJSON(w, http.StatusBadRequest, errRes)
			return
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), productKey{}, prod)
		r = r.WithContext(ctx)

		// call next handler
		next.ServeHTTP(w, r)
	})
}

func (p *Products) MiddlewareContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
