package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		ID:    1,
		Name:  "aaa",
		Price: 10,
		SKU:   "dddd-dddd-dddd",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
