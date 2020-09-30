// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// ThresholdTier Threshold Tiers
// swagger:model ThresholdTier
type ThresholdTier string

const (
	// ThresholdTierCritical captures enum value "Critical"
	ThresholdTierCritical ThresholdTier = "Critical"
	// ThresholdTierMajor captures enum value "Major"
	ThresholdTierMajor ThresholdTier = "Major"
	// ThresholdTierMinor captures enum value "Minor"
	ThresholdTierMinor ThresholdTier = "Minor"
	// ThresholdTierRefresh captures enum value "Refresh"
	ThresholdTierRefresh ThresholdTier = "Refresh"
	// ThresholdTierNone captures enum value "None"
	ThresholdTierNone ThresholdTier = "None"
)

// for schema
var thresholdTierEnum []interface{}

func init() {
	var res []ThresholdTier
	if err := json.Unmarshal([]byte(`["Critical","Major","Minor","Refresh","None"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		thresholdTierEnum = append(thresholdTierEnum, v)
	}
}

func (m ThresholdTier) validateThresholdTierEnum(path, location string, value ThresholdTier) error {
	if err := validate.Enum(path, location, value, thresholdTierEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this threshold tier
func (m ThresholdTier) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateThresholdTierEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}