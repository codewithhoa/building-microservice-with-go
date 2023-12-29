package handlers

import (
	"net/http"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
	"github.com/codewithhoa/building-microservice-with-go/product-api/internal/response"
)

// swagger:route POST /products products postProduct
//
// # Add a new product to DB
//
// This will add a product to database of system.
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
//		       201: NoContentResponsesWrapper
//		       400: ErrorResponsesWrapper
//		       500: ErrorResponsesWrapper

// Post handle post request
func (p *Products) Post(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight post product handler")

	prod, ok := rq.Context().Value(productKey{}).(*data.Product)
	if !ok {
		p.l.Error(ErrInvalidProductInput.Error())

		errRes := response.SimpleErrorResponse(ErrInvalidProductInput, ErrInvalidProductInputCode)
		p.responseJSON(rw, http.StatusBadRequest, errRes)
		return
	}

	err := data.AddProduct(prod)
	if err != nil {
		p.l.Error(err.Error())

		p.responseJSON(rw, http.StatusInternalServerError, response.SimpleErrorResponse(err, err.Error()))
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
