package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
	"github.com/codewithhoa/building-microservice-with-go/product-api/internal/response"
	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/logger"
	"github.com/gorilla/mux"
)

var (
	ErrMarshalJSON         = errors.New("unable to marshal json")
	ErrUnMarshalJSON       = errors.New("unable to unmarshal json")
	ErrGetProducts         = errors.New("can't get list products")
	ErrMethodNotAllow      = errors.New("can't supported method")
	ErrAddProduct          = errors.New("can't add new product to list products")
	ErrInvalidProductId    = errors.New("invalid product id")
	ErrProductNotFound     = errors.New("product not found")
	ErrInvalidProductInput = errors.New("invalid product struct")
)

const (
	ErrMarshalJSONCode         = "UNABLE_MARSHAL_JSON"
	ErrUnMarshalJSONCode       = "UNABLE_UNMARSHAL_JSON"
	ErrGetProductsCode         = "CAN_NOT_GET_PRODUCT"
	ErrMethodNotAllowCode      = "METHOD_NOT_ALLOW"
	ErrAddProductCode          = "CAN_NOT_ADD_PRODUCT"
	ErrInvalidProductIdCode    = "INVALID_PRODUCT_ID"
	ErrProductNotFoundCode     = "PRODUCT_NOT_FOUND"
	ErrInvalidProductInputCode = "VALIDATION_ERROR_PRODUCT"
)

// Products handler for getting and updating products
type Products struct {
	l logger.Logger
	v *data.Validation
}

// NewProductsHandler return a new products handler with the given logger
func NewProductsHandler(l logger.Logger, v *data.Validation) *Products {
	return &Products{
		l: l,
		v: v,
	}
}

// getProductID return the product ID from the given r(*http.Request).
//
//   - return -1, ErrInvalidProductId error if cannot convert the id into integer,
//   - return id, nil if success.
func getProductID(r *http.Request) (int, error) {
	// parse the product ID from the url
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return -1, ErrInvalidProductId
	}

	return id, nil
}

// responseJSON is helper function.
// It helps avoid repeated code to return JSON for request.
//
//   - input: ErrorResponse or Any struct that can marshal to JSON
//
// Example:
//
//	Instead we repeat the following code:
//		err = data.ToJSON(res, rw)
//		if err != nil {
//			p.l.Error(ErrMarshalJSON.Error(), slog.Any("detail error", err.Error()))
//
//			rw.WriteHeader(http.StatusInternalServerError)
//			data.ToJSON(response.SimpleErrorResponse(ErrMarshalJSON, ErrMarshalJSONCode), rw)
//		}
//	We only need to use like:
//		p.responseJSON(rw, http.StatusOK, res)
func (p *Products) responseJSON(rw http.ResponseWriter, statusCode int, input any) {
	rw.WriteHeader(statusCode)
	err := data.ToJSON(input, rw)
	if err != nil {
		p.l.Error(ErrMarshalJSON.Error(), slog.Any("detail error", err.Error()))

		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(response.SimpleErrorResponse(ErrMarshalJSON, ErrMarshalJSONCode), rw)
	}
}
