// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ReplyKeyboardRemove reply keyboard remove
// swagger:model ReplyKeyboardRemove
type ReplyKeyboardRemove struct {

	// remove keyboard
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`

	// selective
	Selective bool `json:"selective,omitempty"`
}

// Validate validates this reply keyboard remove
func (m *ReplyKeyboardRemove) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ReplyKeyboardRemove) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplyKeyboardRemove) UnmarshalBinary(b []byte) error {
	var res ReplyKeyboardRemove
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}