// Code generated by go-swagger; DO NOT EDIT.

package messages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// EditMessageTextReader is a Reader for the EditMessageText structure.
type EditMessageTextReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditMessageTextReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditMessageTextOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditMessageTextBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEditMessageTextUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEditMessageTextForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEditMessageTextNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewEditMessageTextEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEditMessageTextInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditMessageTextOK creates a EditMessageTextOK with default headers values
func NewEditMessageTextOK() *EditMessageTextOK {
	return &EditMessageTextOK{}
}

/*EditMessageTextOK handles this case with default header values.

EditMessageTextOK edit message text o k
*/
type EditMessageTextOK struct {
	Payload *models.EditMessageTextOKBody
}

func (o *EditMessageTextOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextOK  %+v", 200, o.Payload)
}

func (o *EditMessageTextOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EditMessageTextOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageTextBadRequest creates a EditMessageTextBadRequest with default headers values
func NewEditMessageTextBadRequest() *EditMessageTextBadRequest {
	return &EditMessageTextBadRequest{}
}

/*EditMessageTextBadRequest handles this case with default header values.

Bad Request
*/
type EditMessageTextBadRequest struct {
	Payload *models.Error
}

func (o *EditMessageTextBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextBadRequest  %+v", 400, o.Payload)
}

func (o *EditMessageTextBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageTextUnauthorized creates a EditMessageTextUnauthorized with default headers values
func NewEditMessageTextUnauthorized() *EditMessageTextUnauthorized {
	return &EditMessageTextUnauthorized{}
}

/*EditMessageTextUnauthorized handles this case with default header values.

Unauthorized
*/
type EditMessageTextUnauthorized struct {
	Payload *models.Error
}

func (o *EditMessageTextUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextUnauthorized  %+v", 401, o.Payload)
}

func (o *EditMessageTextUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageTextForbidden creates a EditMessageTextForbidden with default headers values
func NewEditMessageTextForbidden() *EditMessageTextForbidden {
	return &EditMessageTextForbidden{}
}

/*EditMessageTextForbidden handles this case with default header values.

Forbidden
*/
type EditMessageTextForbidden struct {
	Payload *models.Error
}

func (o *EditMessageTextForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextForbidden  %+v", 403, o.Payload)
}

func (o *EditMessageTextForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageTextNotFound creates a EditMessageTextNotFound with default headers values
func NewEditMessageTextNotFound() *EditMessageTextNotFound {
	return &EditMessageTextNotFound{}
}

/*EditMessageTextNotFound handles this case with default header values.

Not Found
*/
type EditMessageTextNotFound struct {
	Payload *models.Error
}

func (o *EditMessageTextNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextNotFound  %+v", 404, o.Payload)
}

func (o *EditMessageTextNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageTextEnhanceYourCalm creates a EditMessageTextEnhanceYourCalm with default headers values
func NewEditMessageTextEnhanceYourCalm() *EditMessageTextEnhanceYourCalm {
	return &EditMessageTextEnhanceYourCalm{}
}

/*EditMessageTextEnhanceYourCalm handles this case with default header values.

Flood
*/
type EditMessageTextEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *EditMessageTextEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *EditMessageTextEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageTextInternalServerError creates a EditMessageTextInternalServerError with default headers values
func NewEditMessageTextInternalServerError() *EditMessageTextInternalServerError {
	return &EditMessageTextInternalServerError{}
}

/*EditMessageTextInternalServerError handles this case with default header values.

Internal
*/
type EditMessageTextInternalServerError struct {
	Payload *models.Error
}

func (o *EditMessageTextInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageText][%d] editMessageTextInternalServerError  %+v", 500, o.Payload)
}

func (o *EditMessageTextInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
