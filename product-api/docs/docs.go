// Package handler Product API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//				Schemes: http
//				Host: localhost:9090/api
//				BasePath: /v1
//				License: MIT http://opensource.org/licenses/MIT
//				Contact: Hai Hoa<vukieuhaihoa@gmail.com>
//
//				Consumes:
//				- application/json
//	      - multipart/form-data
//
//				Produces:
//				- application/json
//
// swagger:meta
package docs

// MetaData defines the structure for meta data
// swagger:model
type MetaData struct {
	// The PageSize is total items of one page
	//
	// min: 1
	// Example: 10
	PageSize int `json:"pageSize"`

	// The Page is page which user want retrieve
	//
	// min: 1
	// Example: 2
	Page int `json:"page"`

	// The TotalPages is total of page
	//
	// Example: 100
	TotalPages int `json:"totalPages"`

	// The TotalRecords is total records
	//
	// Example: 1000
	TotalRecords int `json:"totalRecords"`

	// The sort
	//
	// Example: ["+createdAt", "-name"]
	Sort []string `json:"sort"`

	// The HasNext is true if have result in the next page
	//
	// Example: true
	HasNext bool `json:"hasNext"`
}

// swagger:model
type ProductsResponse struct {
	// ID of product in system
	//
	// Example: 123
	ID int `json:"id"`

	// Name of product
	//
	// Example: Cappuccino coffee
	Name string `json:"name"`

	// Description for product
	//
	// Example: Awesome coffee for your new day.
	Description string `json:"description"`

	// Price of product in dollar
	//
	// Example: 12
	Price float64 `json:"price"`

	// SKU of product
	// following format: xxx-xxx-xxx
	//
	// Example: 123-123-123
	SKU string `json:"sku"`
}

type SuccessListProductsResponses struct {
	Data []ProductsResponse `json:"data"`
	Meta MetaData           `json:"metadata"`
}

// successful operation
// swagger:response SuccessListProductsResponsesWrapper
// A SuccessListProductsResponsesWrapper is used to return list of products.
type SuccessListProductsResponsesWrapper struct {
	// in: body
	Body SuccessListProductsResponses `json:"body"`
}

type SuccessSingleProductResponses struct {
	Data ProductsResponse `json:"data"`
}

// successful operation
// swagger:response SuccessSingleProductResponsesWrapper
// A SuccessSingleProductResponsesWrapper is used to return a product.
type SuccessSingleProductResponsesWrapper struct {
	// in: body
	Body SuccessSingleProductResponses `json:"body"`
}

// swagger:model
type Error struct {
	// The Field is name of the field that make an error
	//
	// Example: "field1"
	Field string `json:"field"`

	// The Error is error name
	//
	// Example: "required"
	Error string `json:"error"`
}

// swagger:model
type ErrorResponses struct {
	// The Error is general error
	//
	// Example: "field1: required, field2: greater than 10"
	Error string `json:"error"`

	// The Code is code name of the error
	//
	// Example: "CREATE_BODY_INVALID"
	Code string `json:"code"`

	// The Errors is list of errors
	//
	// Example: [{"field": "field1", "error": "required"}]
	Errors []Error `json:"errors"`

	// The TraceID for tracking the error
	//
	// Example: "u8357577235577jdrf9083227"
	TraceID string `json:"traceID"`
}

// internal error occur in server
// swagger:response ErrorResponsesWrapper
// An ErrorResponsesWrapper is an response that is used when the request fail
type ErrorResponsesWrapper struct {
	// in: body
	Body ErrorResponses `json:"body"`
}

// invalid ID supplied
// swagger:response ErrorInvalidIDSuppliedResponsesWrapper
// An ErrorInvalidIDSuppliedResponsesWrapper is an response that is used when the request fail
type ErrorInvalidIDSuppliedResponsesWrapper struct {
	// in: body
	Body ErrorResponses `json:"body"`
}

// not found error
// swagger:response ErrorNotFoundResponsesWrapper
// An ErrorNotFoundResponsesWrapper is an response that is used when the request fail
type ErrorNotFoundResponsesWrapper struct {
	// in: body
	Body ErrorResponses `json:"body"`
}

// successful operation and no data to return.
// swagger:response NoContentResponsesWrapper
// An NoContentResponsesWrapper is an response that is used when the request success and no content to return.
type NoContentResponsesWrapper struct {
}

// swagger:parameters listSingleProduct deleteProduct putProduct
type ProductIDParamsWrapper struct {
	// The id of the product that you want to do somethings on it.
	// in: path
	// required: true
	// example: 1
	ID int `json:"id"`
}

// swagger:model
type ProductParams struct {
	// Name of product
	//
	// required: true
	// max length: 100
	// Example: Cappuccino coffee
	Name string `json:"name"`

	// Description for product
	//
	// required: true
	// max length: 255
	// Example: Awesome coffee for your new day.
	Description string `json:"description"`

	// Price of product in dollar
	//
	// required: true
	// min: 1
	// Example: 12
	Price float64 `json:"price"`

	// SKU of product
	// following format: xxx-xxx-xxx
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	// Example: aaa-vvv-aaa
	SKU string `json:"sku"`
}

// swagger:parameters postProduct putProduct
type ProductParamsWrapper struct {
	// Product object that needs to be added to the database
	// in: body
	// required: true
	Body ProductParams `json:"body"`
}

// swagger:parameters listProducts
type ProductQueriesWrapper struct {
	// The PageSize is total items of one page
	//
	// min: 1
	// in: query
	// Example: 10
	PageSize int `json:"pageSize"`

	// The Page is page which user want retrieve
	//
	// min: 1
	// in: query
	// Example: 2
	Page int `json:"page"`

	// The sort
	//
	// in: query
	// Example: ["+createdAt", "-name"]
	Sort []string `json:"sort"`
}
