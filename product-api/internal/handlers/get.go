package handlers

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
	"github.com/codewithhoa/building-microservice-with-go/product-api/internal/response"
)

// swagger:route GET /products products listProducts
//
// # Returns all products
//
// Return a list of product from the database.
// Note: You can get the products that are out of stock.
//
//    Consumes:
//      - application/json
//
//    Produces:
//      - application/json
//
//    Schemes: http
//
//    Deprecated: false
//
//    Responses:
//      default: SuccessListProductsResponsesWrapper
//      200: SuccessListProductsResponsesWrapper
//      500: ErrorResponsesWrapper

// GetAll handles GET requests and returns all current products.
func (p *Products) GetAll(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight get products handler")
	lp, err := data.GetProducts()
	if err != nil {
		p.l.Error(err.Error())

		errRes := response.SimpleErrorResponse(ErrGetProducts, ErrGetProductsCode)
		p.responseJSON(rw, http.StatusInternalServerError, errRes)
		return
	}

	// TODO: paging the return value
	res := response.SuccessMultiResponse{
		Data: lp,
		Metadata: response.MetaData{
			TotalRecords: len(lp),
		},
	}

	p.responseJSON(rw, http.StatusOK, res)
}

// swagger:route GET /products/{id} products listSingleProduct
//
// # Finds a product by ID
//
// Return a product by ID from request.
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
//      default: SuccessSingleProductResponsesWrapper
//      200: SuccessSingleProductResponsesWrapper
//      400: ErrorInvalidIDSuppliedResponsesWrapper
//      404: ErrorNotFoundResponsesWrapper
//      500: ErrorResponsesWrapper

// GetByID handles GET requests and returns product by ID.
func (p *Products) GetByID(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight get one product handler")

	id, err := getProductID(rq)
	if err != nil {
		p.l.Error(err.Error())

		errRes := response.SimpleErrorResponse(ErrInvalidProductId, ErrInvalidProductIdCode)
		p.responseJSON(rw, http.StatusBadRequest, errRes)
		return
	}

	p.l.Debug("id of product that want to retrieve", slog.Int("id", id))

	prod, err := data.GetProductByID(id)
	if err != nil {
		if errors.Is(err, data.ErrProductNotFound) {
			p.l.Error(err.Error())

			errRes := response.SimpleErrorResponse(ErrProductNotFound, ErrProductNotFoundCode)

			p.responseJSON(rw, http.StatusNotFound, errRes)
			return
		}

		p.responseJSON(rw, http.StatusInternalServerError, response.SimpleErrorResponse(err, err.Error()))
		return
	}

	res := response.SuccessSingleResponse{
		Data: prod,
	}
	p.responseJSON(rw, http.StatusOK, res)
}
