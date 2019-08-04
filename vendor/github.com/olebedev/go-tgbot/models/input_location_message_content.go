// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InputLocationMessageContent input location message content
// swagger:model InputLocationMessageContent
type InputLocationMessageContent struct {

	// latitude
	// Required: true
	Latitude *float64 `json:"latitude"`

	// live period
	// Maximum: 86400
	// Minimum: 60
	LivePeriod int64 `json:"live_period,omitempty"`

	// longitude
	// Required: true
	Longitude *float64 `json:"longitude"`
}

// Validate validates this input location message content
func (m *InputLocationMessageContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLatitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLivePeriod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLongitude(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InputLocationMessageContent) validateLatitude(formats strfmt.Registry) error {

	if err := validate.Required("latitude", "body", m.Latitude); err != nil {
		return err
	}

	return nil
}

func (m *InputLocationMessageContent) validateLivePeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.LivePeriod) { // not required
		return nil
	}

	if err := validate.MinimumInt("live_period", "body", int64(m.LivePeriod), 60, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("live_period", "body", int64(m.LivePeriod), 86400, false); err != nil {
		return err
	}

	return nil
}

func (m *InputLocationMessageContent) validateLongitude(formats strfmt.Registry) error {

	if err := validate.Required("longitude", "body", m.Longitude); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InputLocationMessageContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputLocationMessageContent) UnmarshalBinary(b []byte) error {
	var res InputLocationMessageContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}