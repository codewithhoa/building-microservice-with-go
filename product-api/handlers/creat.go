package handlers

import (
	"net/http"

	"github.com/codewithhoa/building-microservice-with-go/product-api/data"
)

func (p *Products) Create(rw http.ResponseWriter, rq *http.Request) {
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
