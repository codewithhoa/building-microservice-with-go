package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
	"github.com/codewithhoa/building-microservice-with-go/product-api/internal/response"
)

// swagger:route PUT /products/{id} products putProduct
//
// # Updates an existing product by ID
//
// This operation will updates an existing product by ID.
//
//					Consumes:
//					- application/json
//
//					Produces:
//					- application/json
//
//					Schemes: http
//
//				  Deprecated: false
//
//			   	Responses:
//			     default: NoContentResponsesWrapper
//		       204: NoContentResponsesWrapper
//		       400: ErrorResponsesWrapper
//           404: ErrorNotFoundResponsesWrapper
//		       500: ErrorResponsesWrapper

// Put handle put request
func (p *Products) Put(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight put product handler")

	id, err := getProductID(rq)
	if err != nil {
		p.l.Error(err.Error())

		errRes := response.SimpleErrorResponse(ErrInvalidProductId, ErrInvalidProductIdCode)
		p.responseJSON(rw, http.StatusBadRequest, errRes)
		return
	}

	p.l.Debug("id of product that need to update", slog.Int("id", id))

	prod, ok := rq.Context().Value(productKey{}).(*data.Product)
	if !ok {
		p.l.Error(ErrInvalidProductInput.Error())

		errRes := response.SimpleErrorResponse(ErrInvalidProductInput, ErrInvalidProductInputCode)
		p.responseJSON(rw, http.StatusBadRequest, errRes)
		return
	}

	err = data.UpdateProductByID(id, prod)
	if err != nil {
		p.l.Error(err.Error())

		if errors.Is(err, data.ErrProductNotFound) {
			errRes := response.SimpleErrorResponse(ErrProductNotFound, ErrProductNotFoundCode)

			p.responseJSON(rw, http.StatusNotFound, errRes)
			return
		}

		// http.Error(rw, err.Error(), http.StatusInternalServerError)
		p.responseJSON(rw, http.StatusInternalServerError, response.SimpleErrorResponse(err, err.Error()))
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
