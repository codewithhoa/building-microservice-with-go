package main

import (
	"fmt"
	"testing"

	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/sdk/client"
	"github.com/codewithhoa/building-microservice-with-go/product-api/pkg/sdk/client/products"
)

func TestGetListProducts(t *testing.T) {
	c := client.Default
	params := products.NewListProductsParams()
	resp, err := c.Products.ListProducts(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.GetPayload().Data[0])
	// t.Fail()
}

func TestGetProductByID(t *testing.T) {
	c := client.Default
	params := products.NewListSingleProductParams()
	params.SetID(2)
	resp, err := c.Products.ListSingleProduct(params)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", resp.GetPayload().Data)
	t.Fail()
}
