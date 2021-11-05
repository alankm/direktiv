// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateGlobalServiceBody CreateGlobalServiceBody create global service body
// Example: {"cmd":"","image":"direktiv/request:v12","minScale":"1","name":"fast-request","size":"small"}
//
// swagger:model CreateGlobalServiceBody
type CreateGlobalServiceBody struct {

	// cmd
	// Required: true
	Cmd *string `json:"cmd"`

	// Target image a service will use
	// Required: true
	Image *string `json:"image"`

	// Minimum amount of service pods to be live
	// Required: true
	MinScale *int64 `json:"minScale"`

	// Name of new service
	// Required: true
	Name *string `json:"name"`

	// Size of created service pods
	// Required: true
	// Enum: [[small medium large]]
	Size *string `json:"size"`
}

// Validate validates this create global service body
func (m *CreateGlobalServiceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCmd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinScale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateGlobalServiceBody) validateCmd(formats strfmt.Registry) error {

	if err := validate.Required("cmd", "body", m.Cmd); err != nil {
		return err
	}

	return nil
}

func (m *CreateGlobalServiceBody) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *CreateGlobalServiceBody) validateMinScale(formats strfmt.Registry) error {

	if err := validate.Required("minScale", "body", m.MinScale); err != nil {
		return err
	}

	return nil
}

func (m *CreateGlobalServiceBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var createGlobalServiceBodyTypeSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["[small medium large]"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createGlobalServiceBodyTypeSizePropEnum = append(createGlobalServiceBodyTypeSizePropEnum, v)
	}
}

const (

	// CreateGlobalServiceBodySizeSmallMediumLarge captures enum value "[small medium large]"
	CreateGlobalServiceBodySizeSmallMediumLarge string = "[small medium large]"
)

// prop value enum
func (m *CreateGlobalServiceBody) validateSizeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createGlobalServiceBodyTypeSizePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateGlobalServiceBody) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	// value enum
	if err := m.validateSizeEnum("size", "body", *m.Size); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create global service body based on context it is used
func (m *CreateGlobalServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateGlobalServiceBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateGlobalServiceBody) UnmarshalBinary(b []byte) error {
	var res CreateGlobalServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}