// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// VideoNote video note
// swagger:model VideoNote
type VideoNote struct {

	// duration
	Duration int64 `json:"duration,omitempty"`

	// file id
	FileID string `json:"file_id,omitempty"`

	// file size
	FileSize int64 `json:"file_size,omitempty"`

	// length
	Length int64 `json:"length,omitempty"`

	// thumb
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// Validate validates this video note
func (m *VideoNote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateThumb(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VideoNote) validateThumb(formats strfmt.Registry) error {

	if swag.IsZero(m.Thumb) { // not required
		return nil
	}

	if m.Thumb != nil {

		if err := m.Thumb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("thumb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VideoNote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VideoNote) UnmarshalBinary(b []byte) error {
	var res VideoNote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
