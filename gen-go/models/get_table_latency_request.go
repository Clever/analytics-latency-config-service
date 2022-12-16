// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetTableLatencyRequest get table latency request
//
// swagger:model GetTableLatencyRequest
type GetTableLatencyRequest struct {

	// database
	// Required: true
	Database AnalyticsDatabase `json:"database"`

	// schema
	// Required: true
	Schema *string `json:"schema"`

	// table
	// Required: true
	Table *string `json:"table"`
}

// Validate validates this get table latency request
func (m *GetTableLatencyRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTable(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTableLatencyRequest) validateDatabase(formats strfmt.Registry) error {

	if err := m.Database.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("database")
		}
		return err
	}

	return nil
}

func (m *GetTableLatencyRequest) validateSchema(formats strfmt.Registry) error {

	if err := validate.Required("schema", "body", m.Schema); err != nil {
		return err
	}

	return nil
}

func (m *GetTableLatencyRequest) validateTable(formats strfmt.Registry) error {

	if err := validate.Required("table", "body", m.Table); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetTableLatencyRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTableLatencyRequest) UnmarshalBinary(b []byte) error {
	var res GetTableLatencyRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
