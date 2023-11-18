package handlers

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/codewithhoa/building-microservice-with-go/data"
)

var (
	ErrMarshalJSON      = errors.New("unable to marshal json")
	ErrUnMarshalJSON    = errors.New("unable to unmarshal json")
	ErrGetProducts      = errors.New("can not get list products")
	ErrMethodNotAllow   = errors.New("do not supported method")
	ErrAddProduct       = errors.New("can not add new product to list products")
	ErrInvalidProductId = errors.New("invalid product id")
)

// Products is a http.Handler
type Products struct {
	l *slog.Logger
}

var _ http.Handler = (*Products)(nil)

func NewProductsHandler(l *slog.Logger) *Products {
	return &Products{
		l: l,
	}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, rq *http.Request) {
	if rq.Method == http.MethodGet {
		p.getProducts(rw, rq)
		return
	}

	if rq.Method == http.MethodPost {
		p.addProduct(rw, rq)
		return
	}

	if rq.Method == http.MethodPut {
		// Get product id
		idString := strings.TrimPrefix(rq.URL.Path, "/products/")
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, ErrInvalidProductId.Error(), http.StatusBadRequest)
			return
		}
		p.l.Info("id of product that was needed to update", slog.Int("product id", id))
		p.updateProduct(id, rw, rq)
		return
	}

	http.Error(rw, ErrMethodNotAllow.Error(), http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, rq *http.Request) {
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

func (p *Products) addProduct(rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight post product handler")
	np := &data.Product{}

	err := np.FromJSON(rq.Body)
	if err != nil {
		http.Error(rw, ErrUnMarshalJSON.Error(), http.StatusBadRequest)
		return
	}

	// p.l.Debug(fmt.Sprintf("%#v", np))
	err = data.AddProduct(np)
	if err != nil {
		http.Error(rw, ErrAddProduct.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, rq *http.Request) {
	p.l.Info("insight put product handler")
	prod := &data.Product{}

	err := prod.FromJSON(rq.Body)
	if err != nil {
		http.Error(rw, ErrUnMarshalJSON.Error(), http.StatusBadRequest)
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
