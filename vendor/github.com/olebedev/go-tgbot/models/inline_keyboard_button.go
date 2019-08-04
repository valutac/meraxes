// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InlineKeyboardButton inline keyboard button
// swagger:model InlineKeyboardButton
type InlineKeyboardButton struct {

	// callback data
	CallbackData string `json:"callback_data,omitempty"`

	// callback game
	CallbackGame CallbackGame `json:"callback_game,omitempty"`

	// pay
	Pay bool `json:"pay,omitempty"`

	// switch inline query
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`

	// switch inline query current chat
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`

	// text
	Text string `json:"text,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this inline keyboard button
func (m *InlineKeyboardButton) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InlineKeyboardButton) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineKeyboardButton) UnmarshalBinary(b []byte) error {
	var res InlineKeyboardButton
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
