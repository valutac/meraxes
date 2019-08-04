// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// GetUpdatesBody get updates body
// swagger:model GetUpdatesBody
type GetUpdatesBody struct {

	// allowed updates
	AllowedUpdates []AllowedUpdate `json:"allowed_updates"`

	// limit
	Limit int64 `json:"limit,omitempty"`

	// offset
	Offset int64 `json:"offset,omitempty"`

	// timeout
	Timeout int64 `json:"timeout,omitempty"`
}

// Validate validates this get updates body
func (m *GetUpdatesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowedUpdates(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUpdatesBody) validateAllowedUpdates(formats strfmt.Registry) error {

	if swag.IsZero(m.AllowedUpdates) { // not required
		return nil
	}

	for i := 0; i < len(m.AllowedUpdates); i++ {

		if err := m.AllowedUpdates[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowed_updates" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUpdatesBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUpdatesBody) UnmarshalBinary(b []byte) error {
	var res GetUpdatesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
