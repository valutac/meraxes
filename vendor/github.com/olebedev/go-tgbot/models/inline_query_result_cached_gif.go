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

// InlineQueryResultCachedGif inline query result cached gif
// swagger:model InlineQueryResultCachedGif
type InlineQueryResultCachedGif struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// gif file id
	// Required: true
	GifFileID *string `json:"gif_file_id"`

	// id
	// Required: true
	ID *string `json:"id"`

	// input message content
	InputMessageContent interface{} `json:"input_message_content,omitempty"`

	// parse mode
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// reply markup
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	// Required: true
	Type InlineType `json:"type"`
}

// Validate validates this inline query result cached gif
func (m *InlineQueryResultCachedGif) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGifFileID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParseMode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyMarkup(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InlineQueryResultCachedGif) validateGifFileID(formats strfmt.Registry) error {

	if err := validate.Required("gif_file_id", "body", m.GifFileID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedGif) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *InlineQueryResultCachedGif) validateParseMode(formats strfmt.Registry) error {

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

func (m *InlineQueryResultCachedGif) validateReplyMarkup(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyMarkup) { // not required
		return nil
	}

	if m.ReplyMarkup != nil {

		if err := m.ReplyMarkup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply_markup")
			}
			return err
		}
	}

	return nil
}

func (m *InlineQueryResultCachedGif) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InlineQueryResultCachedGif) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InlineQueryResultCachedGif) UnmarshalBinary(b []byte) error {
	var res InlineQueryResultCachedGif
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
