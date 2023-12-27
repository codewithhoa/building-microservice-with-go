package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		ID:    1,
		Name:  "aaa",
		Price: 10,
		SKU:   "dddd-dddd-dddd",
	}

	validation := NewValidation()

	err := validation.Validate(p)
	if err != nil {
		t.Fatal(err)
	}
}
