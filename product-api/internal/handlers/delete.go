package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
	"github.com/codewithhoa/building-microservice-with-go/product-api/internal/response"
)

// swagger:route DELETE /products/{id} products deleteProduct
//
// # Delete an existing product by ID
//
//    Consumes:
//    - application/json
//
//    Produces:
//    - application/json
//
//    Schemes: http
//
//    Deprecated: false
//
//    Responses:
//      default: NoContentResponsesWrapper
//      204: NoContentResponsesWrapper
//      400: ErrorInvalidIDSuppliedResponsesWrapper
//      404: ErrorNotFoundResponsesWrapper
//      500: ErrorResponsesWrapper

// Delete handles DEL requests.
func (p *Products) Delete(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight delete delete product handler")

	id, err := getProductID(rq)
	if err != nil {
		p.l.Error(err.Error())

		errRes := response.SimpleErrorResponse(ErrInvalidProductId, ErrInvalidProductIdCode)
		p.responseJSON(rw, http.StatusBadRequest, errRes)
		return
	}

	p.l.Debug("id of product that need to delete", slog.Int("id", id))

	if err := data.DeleteProductByID(id); err != nil {
		p.l.Error(err.Error())

		if errors.Is(err, data.ErrProductNotFound) {
			errRes := response.SimpleErrorResponse(ErrProductNotFound, ErrProductNotFoundCode)

			p.responseJSON(rw, http.StatusNotFound, errRes)
			return
		}

		p.responseJSON(rw, http.StatusInternalServerError, response.SimpleErrorResponse(err, err.Error()))
		return
	}

	rw.WriteHeader(http.StatusNoContent)
}
