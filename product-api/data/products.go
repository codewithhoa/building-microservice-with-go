package data

import (
	"errors"
	"time"
)

var (
	// ErrProductNotFound is an error raised when a product can not be found in the database
	ErrProductNotFound = errors.New("product not found")
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// The id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product

	// The name of the product
	//
	// required: true
	// max length: 255
	Name string `json:"name" validate:"required"`

	// The description of the product
	//
	// required: true
	// max length: 1000
	Description string `json:"description"`

	// The price of the product
	//
	// required: true
	// min: 0.01
	Price float64 `json:"price" validate:"gt=0"`

	// The SKU of the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`

	// The time when product is created
	//
	// required: false
	// swagger:strfmt date
	CreatedOn time.Time `json:"createdOn"`

	// The time when product is updated
	//
	// required: false
	// swagger:strfmt date
	UpdatedOn time.Time `json:"updatedOn"`

	// The time when product is deleted
	//
	// required: false
	// swagger:strfmt date
	DeletedOn time.Time `json:"deletedOn"`
}

// Products defines a slice of Product(pointer of a Product)
type Products []*Product

// In-memory database for the system
// TODO: Please use real DB If you have time.
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

// GetProducts returns a list of products from the database
func GetProducts() (Products, error) {
	// return nil, ErrProductNotFound
	return productList, nil
}

// GetProductByID returns a single product which matches the id from the
// database.
// If a product is not found this func returns a ErrProductNotFound error
func GetProductByID(id int) (*Product, error) {
	pos := findProductByID(id)
	if pos == -1 {
		return nil, ErrProductNotFound
	}

	return productList[pos], nil
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

// UpdateProductByID updates a product by given id.
// If a product with the given id does not exist in the DB,
// this func return a ErrProductNotFound error.
func UpdateProductByID(id int, p *Product) error {
	pos := findProductByID(id)
	if pos == -1 {
		return ErrProductNotFound
	}

	now := time.Now().UTC()
	p.ID = id
	p.UpdatedOn = now
	productList[pos] = p
	return nil
}

// DeleteProductByID deletes a product by give id.
// If a product with the given id does not exist in the DB,
// this func return a ErrProductNotFound error.
func DeleteProductByID(id int) error {
	pos := findProductByID(id)
	if pos == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:pos], productList[pos+1:]...)
	return nil
}

// findProductByID return the index of a product in the database if matching ID
// return -1 when no product matching with an given id.
func findProductByID(id int) int {
	for i, product := range productList {
		if product.ID == id {
			return i
		}
	}

	return -1
}
