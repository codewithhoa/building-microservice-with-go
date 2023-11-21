package data

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	ErrProductNotFound = errors.New("product not found")
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Price       float64   `json:"price" validate:"gt=0"`
	SKU         string    `json:"sku" validate:"required,sku"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
}

func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)

	return validate.Struct(p)
}

// validateSKU implement validator.Func
func validateSKU(fl validator.FieldLevel) bool {
	// *: This line may be suddenly interrupt app because it can panic
	re := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)

	// find all match string
	matches := re.FindAllString(fl.Field().String(), -1)
	return len(matches) == 1
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "def456",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
}

// GetProducts returns a list of products
func GetProducts() (Products, error) {
	return productList, nil
}

// AddProduct add a new product to product list
func AddProduct(p *Product) error {
	p.ID = getNextID()
	now := time.Now().UTC()
	p.CreatedOn = now
	productList = append(productList, p)

	return nil
}

// getNextID return the next id in product list
func getNextID() int {
	if len(productList) == 0 {
		return 1
	}

	lp := productList[len(productList)-1]
	return lp.ID + 1
}

func UpdateProductById(id int, p *Product) error {
	_, pos, err := findProductById(id)
	if err != nil {
		return err
	}

	now := time.Now().UTC()
	p.ID = id
	p.UpdatedOn = now
	productList[pos] = p
	return nil
}

// findProductById
func findProductById(id int) (*Product, int, error) {
	for i, product := range productList {
		if product.ID == id {
			return product, i, nil
		}
	}

	return nil, 0, ErrProductNotFound
}
