// Code generated by go-swagger; DO NOT EDIT.

package chats

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// PinChatMessageReader is a Reader for the PinChatMessage structure.
type PinChatMessageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PinChatMessageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPinChatMessageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPinChatMessageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPinChatMessageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPinChatMessageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPinChatMessageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewPinChatMessageEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPinChatMessageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPinChatMessageOK creates a PinChatMessageOK with default headers values
func NewPinChatMessageOK() *PinChatMessageOK {
	return &PinChatMessageOK{}
}

/*PinChatMessageOK handles this case with default header values.

PinChatMessageOK pin chat message o k
*/
type PinChatMessageOK struct {
	Payload *models.ResponseBool
}

func (o *PinChatMessageOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/pinChatMessage][%d] pinChatMessageOK  %+v", 200, o.Payload)
}

func (o *PinChatMessageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseBool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinChatMessageBadRequest creates a PinChatMessageBadRequest with default headers values
func NewPinChatMessageBadRequest() *PinChatMessageBadRequest {
	return &PinChatMessageBadRequest{}
}

/*PinChatMessageBadRequest handles this case with default header values.

Bad Request
*/
type PinChatMessageBadRequest struct {
	Payload *models.Error
}

func (o *PinChatMessageBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/pinChatMessage][%d] pinChatMessageBadRequest  %+v", 400, o.Payload)
}

func (o *PinChatMessageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinChatMessageUnauthorized creates a PinChatMessageUnauthorized with default headers values
func NewPinChatMessageUnauthorized() *PinChatMessageUnauthorized {
	return &PinChatMessageUnauthorized{}
}

/*PinChatMessageUnauthorized handles this case with default header values.

Unauthorized
*/
type PinChatMessageUnauthorized struct {
	Payload *models.Error
}

func (o *PinChatMessageUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/pinChatMessage][%d] pinChatMessageUnauthorized  %+v", 401, o.Payload)
}

func (o *PinChatMessageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinChatMessageForbidden creates a PinChatMessageForbidden with default headers values
func NewPinChatMessageForbidden() *PinChatMessageForbidden {
	return &PinChatMessageForbidden{}
}

/*PinChatMessageForbidden handles this case with default header values.

Forbidden
*/
type PinChatMessageForbidden struct {
	Payload *models.Error
}

func (o *PinChatMessageForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/pinChatMessage][%d] pinChatMessageForbidden  %+v", 403, o.Payload)
}

func (o *PinChatMessageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinChatMessageNotFound creates a PinChatMessageNotFound with default headers values
func NewPinChatMessageNotFound() *PinChatMessageNotFound {
	return &PinChatMessageNotFound{}
}

/*PinChatMessageNotFound handles this case with default header values.

Not Found
*/
type PinChatMessageNotFound struct {
	Payload *models.Error
}

func (o *PinChatMessageNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/pinChatMessage][%d] pinChatMessageNotFound  %+v", 404, o.Payload)
}

func (o *PinChatMessageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinChatMessageEnhanceYourCalm creates a PinChatMessageEnhanceYourCalm with default headers values
func NewPinChatMessageEnhanceYourCalm() *PinChatMessageEnhanceYourCalm {
	return &PinChatMessageEnhanceYourCalm{}
}

/*PinChatMessageEnhanceYourCalm handles this case with default header values.

Flood
*/
type PinChatMessageEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *PinChatMessageEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/pinChatMessage][%d] pinChatMessageEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PinChatMessageEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPinChatMessageInternalServerError creates a PinChatMessageInternalServerError with default headers values
func NewPinChatMessageInternalServerError() *PinChatMessageInternalServerError {
	return &PinChatMessageInternalServerError{}
}

/*PinChatMessageInternalServerError handles this case with default header values.

Internal
*/
type PinChatMessageInternalServerError struct {
	Payload *models.Error
}

func (o *PinChatMessageInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/pinChatMessage][%d] pinChatMessageInternalServerError  %+v", 500, o.Payload)
}

func (o *PinChatMessageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
