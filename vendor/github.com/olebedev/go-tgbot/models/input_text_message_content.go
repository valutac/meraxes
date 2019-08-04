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

// InputTextMessageContent input text message content
// swagger:model InputTextMessageContent
type InputTextMessageContent struct {

	// disable web page preview
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// message text
	// Required: true
	MessageText *string `json:"message_text"`

	// parse mode
	ParseMode ParseMode `json:"parse_mode,omitempty"`
}

// Validate validates this input text message content
func (m *InputTextMessageContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessageText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParseMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InputTextMessageContent) validateMessageText(formats strfmt.Registry) error {

	if err := validate.Required("message_text", "body", m.MessageText); err != nil {
		return err
	}

	return nil
}

func (m *InputTextMessageContent) validateParseMode(formats strfmt.Registry) error {

	if swag.IsZero(m.ParseMode) { // not required
		return nil
	}

	if err := m.ParseMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("parse_mode")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InputTextMessageContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InputTextMessageContent) UnmarshalBinary(b []byte) error {
	var res InputTextMessageContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}