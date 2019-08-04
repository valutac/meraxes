// Code generated by go-swagger; DO NOT EDIT.

package stickers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// CreateNewStickerSetLinkReader is a Reader for the CreateNewStickerSetLink structure.
type CreateNewStickerSetLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNewStickerSetLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateNewStickerSetLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateNewStickerSetLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewCreateNewStickerSetLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateNewStickerSetLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateNewStickerSetLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewCreateNewStickerSetLinkEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewCreateNewStickerSetLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateNewStickerSetLinkOK creates a CreateNewStickerSetLinkOK with default headers values
func NewCreateNewStickerSetLinkOK() *CreateNewStickerSetLinkOK {
	return &CreateNewStickerSetLinkOK{}
}

/*CreateNewStickerSetLinkOK handles this case with default header values.

CreateNewStickerSetLinkOK create new sticker set link o k
*/
type CreateNewStickerSetLinkOK struct {
	Payload *models.File
}

func (o *CreateNewStickerSetLinkOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/createNewStickerSet#link][%d] createNewStickerSetLinkOK  %+v", 200, o.Payload)
}

func (o *CreateNewStickerSetLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNewStickerSetLinkBadRequest creates a CreateNewStickerSetLinkBadRequest with default headers values
func NewCreateNewStickerSetLinkBadRequest() *CreateNewStickerSetLinkBadRequest {
	return &CreateNewStickerSetLinkBadRequest{}
}

/*CreateNewStickerSetLinkBadRequest handles this case with default header values.

Bad Request
*/
type CreateNewStickerSetLinkBadRequest struct {
	Payload *models.Error
}

func (o *CreateNewStickerSetLinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/createNewStickerSet#link][%d] createNewStickerSetLinkBadRequest  %+v", 400, o.Payload)
}

func (o *CreateNewStickerSetLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNewStickerSetLinkUnauthorized creates a CreateNewStickerSetLinkUnauthorized with default headers values
func NewCreateNewStickerSetLinkUnauthorized() *CreateNewStickerSetLinkUnauthorized {
	return &CreateNewStickerSetLinkUnauthorized{}
}

/*CreateNewStickerSetLinkUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateNewStickerSetLinkUnauthorized struct {
	Payload *models.Error
}

func (o *CreateNewStickerSetLinkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/createNewStickerSet#link][%d] createNewStickerSetLinkUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateNewStickerSetLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNewStickerSetLinkForbidden creates a CreateNewStickerSetLinkForbidden with default headers values
func NewCreateNewStickerSetLinkForbidden() *CreateNewStickerSetLinkForbidden {
	return &CreateNewStickerSetLinkForbidden{}
}

/*CreateNewStickerSetLinkForbidden handles this case with default header values.

Forbidden
*/
type CreateNewStickerSetLinkForbidden struct {
	Payload *models.Error
}

func (o *CreateNewStickerSetLinkForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/createNewStickerSet#link][%d] createNewStickerSetLinkForbidden  %+v", 403, o.Payload)
}

func (o *CreateNewStickerSetLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNewStickerSetLinkNotFound creates a CreateNewStickerSetLinkNotFound with default headers values
func NewCreateNewStickerSetLinkNotFound() *CreateNewStickerSetLinkNotFound {
	return &CreateNewStickerSetLinkNotFound{}
}

/*CreateNewStickerSetLinkNotFound handles this case with default header values.

Not Found
*/
type CreateNewStickerSetLinkNotFound struct {
	Payload *models.Error
}

func (o *CreateNewStickerSetLinkNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/createNewStickerSet#link][%d] createNewStickerSetLinkNotFound  %+v", 404, o.Payload)
}

func (o *CreateNewStickerSetLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNewStickerSetLinkEnhanceYourCalm creates a CreateNewStickerSetLinkEnhanceYourCalm with default headers values
func NewCreateNewStickerSetLinkEnhanceYourCalm() *CreateNewStickerSetLinkEnhanceYourCalm {
	return &CreateNewStickerSetLinkEnhanceYourCalm{}
}

/*CreateNewStickerSetLinkEnhanceYourCalm handles this case with default header values.

Flood
*/
type CreateNewStickerSetLinkEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *CreateNewStickerSetLinkEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/createNewStickerSet#link][%d] createNewStickerSetLinkEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *CreateNewStickerSetLinkEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateNewStickerSetLinkInternalServerError creates a CreateNewStickerSetLinkInternalServerError with default headers values
func NewCreateNewStickerSetLinkInternalServerError() *CreateNewStickerSetLinkInternalServerError {
	return &CreateNewStickerSetLinkInternalServerError{}
}

/*CreateNewStickerSetLinkInternalServerError handles this case with default header values.

Internal
*/
type CreateNewStickerSetLinkInternalServerError struct {
	Payload *models.Error
}

func (o *CreateNewStickerSetLinkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/createNewStickerSet#link][%d] createNewStickerSetLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateNewStickerSetLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
