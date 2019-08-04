// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CallbackQuery callback query
// swagger:model CallbackQuery
type CallbackQuery struct {

	// chat instance
	ChatInstance string `json:"chat_instance,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// from
	From *User `json:"from,omitempty"`

	// game short name
	GameShortName string `json:"game_short_name,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// inline message id
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// message
	Message *Message `json:"message,omitempty"`
}

// Validate validates this callback query
func (m *CallbackQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallbackQuery) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if m.From != nil {

		if err := m.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (m *CallbackQuery) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {

		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallbackQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallbackQuery) UnmarshalBinary(b []byte) error {
	var res CallbackQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}