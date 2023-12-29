// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProductParams product params
//
// swagger:model ProductParams
type ProductParams struct {

	// Description for product
	// Example: Awesome coffee for your new day.
	// Required: true
	// Max Length: 255
	Description *string `json:"description"`

	// Name of product
	// Example: Cappuccino coffee
	// Required: true
	// Max Length: 100
	Name *string `json:"name"`

	// Price of product in dollar
	// Example: 12
	// Required: true
	// Minimum: 1
	Price *float64 `json:"price"`

	// SKU of product
	// following format: xxx-xxx-xxx
	// Example: aaa-vvv-aaa
	// Required: true
	// Pattern: [a-z]+-[a-z]+-[a-z]+
	SKU *string `json:"sku"`
}

// Validate validates this product params
func (m *ProductParams) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSKU(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductParams) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", *m.Description, 255); err != nil {
		return err
	}

	return nil
}

func (m *ProductParams) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *ProductParams) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	if err := validate.Minimum("price", "body", *m.Price, 1, false); err != nil {
		return err
	}

	return nil
}

func (m *ProductParams) validateSKU(formats strfmt.Registry) error {

	if err := validate.Required("sku", "body", m.SKU); err != nil {
		return err
	}

	if err := validate.Pattern("sku", "body", *m.SKU, `[a-z]+-[a-z]+-[a-z]+`); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this product params based on context it is used
func (m *ProductParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProductParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductParams) UnmarshalBinary(b []byte) error {
	var res ProductParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
