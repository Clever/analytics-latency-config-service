// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AnalyticsLatencyConfigs analytics latency configs
//
// swagger:model AnalyticsLatencyConfigs
type AnalyticsLatencyConfigs struct {

	// rds external
	// Required: true
	RdsExternal []*SchemaConfig `json:"rdsExternal"`

	// redshift fast
	// Required: true
	RedshiftFast []*SchemaConfig `json:"redshiftFast"`

	// snowflake
	// Required: true
	Snowflake []*SchemaConfig `json:"snowflake"`
}

// Validate validates this analytics latency configs
func (m *AnalyticsLatencyConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRdsExternal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedshiftFast(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnowflake(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsLatencyConfigs) validateRdsExternal(formats strfmt.Registry) error {

	if err := validate.Required("rdsExternal", "body", m.RdsExternal); err != nil {
		return err
	}

	for i := 0; i < len(m.RdsExternal); i++ {
		if swag.IsZero(m.RdsExternal[i]) { // not required
			continue
		}

		if m.RdsExternal[i] != nil {
			if err := m.RdsExternal[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rdsExternal" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsLatencyConfigs) validateRedshiftFast(formats strfmt.Registry) error {

	if err := validate.Required("redshiftFast", "body", m.RedshiftFast); err != nil {
		return err
	}

	for i := 0; i < len(m.RedshiftFast); i++ {
		if swag.IsZero(m.RedshiftFast[i]) { // not required
			continue
		}

		if m.RedshiftFast[i] != nil {
			if err := m.RedshiftFast[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("redshiftFast" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AnalyticsLatencyConfigs) validateSnowflake(formats strfmt.Registry) error {

	if err := validate.Required("snowflake", "body", m.Snowflake); err != nil {
		return err
	}

	for i := 0; i < len(m.Snowflake); i++ {
		if swag.IsZero(m.Snowflake[i]) { // not required
			continue
		}

		if m.Snowflake[i] != nil {
			if err := m.Snowflake[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("snowflake" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsLatencyConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsLatencyConfigs) UnmarshalBinary(b []byte) error {
	var res AnalyticsLatencyConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
