// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// SendVideoNoteReader is a Reader for the SendVideoNote structure.
type SendVideoNoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendVideoNoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendVideoNoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendVideoNoteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendVideoNoteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendVideoNoteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendVideoNoteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendVideoNoteEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendVideoNoteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendVideoNoteOK creates a SendVideoNoteOK with default headers values
func NewSendVideoNoteOK() *SendVideoNoteOK {
	return &SendVideoNoteOK{}
}

/*SendVideoNoteOK handles this case with default header values.

SendVideoNoteOK send video note o k
*/
type SendVideoNoteOK struct {
	Payload *models.ResponseMessage
}

func (o *SendVideoNoteOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote][%d] sendVideoNoteOK  %+v", 200, o.Payload)
}

func (o *SendVideoNoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteBadRequest creates a SendVideoNoteBadRequest with default headers values
func NewSendVideoNoteBadRequest() *SendVideoNoteBadRequest {
	return &SendVideoNoteBadRequest{}
}

/*SendVideoNoteBadRequest handles this case with default header values.

Bad Request
*/
type SendVideoNoteBadRequest struct {
	Payload *models.Error
}

func (o *SendVideoNoteBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote][%d] sendVideoNoteBadRequest  %+v", 400, o.Payload)
}

func (o *SendVideoNoteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteUnauthorized creates a SendVideoNoteUnauthorized with default headers values
func NewSendVideoNoteUnauthorized() *SendVideoNoteUnauthorized {
	return &SendVideoNoteUnauthorized{}
}

/*SendVideoNoteUnauthorized handles this case with default header values.

Unauthorized
*/
type SendVideoNoteUnauthorized struct {
	Payload *models.Error
}

func (o *SendVideoNoteUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote][%d] sendVideoNoteUnauthorized  %+v", 401, o.Payload)
}

func (o *SendVideoNoteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteForbidden creates a SendVideoNoteForbidden with default headers values
func NewSendVideoNoteForbidden() *SendVideoNoteForbidden {
	return &SendVideoNoteForbidden{}
}

/*SendVideoNoteForbidden handles this case with default header values.

Forbidden
*/
type SendVideoNoteForbidden struct {
	Payload *models.Error
}

func (o *SendVideoNoteForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote][%d] sendVideoNoteForbidden  %+v", 403, o.Payload)
}

func (o *SendVideoNoteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteNotFound creates a SendVideoNoteNotFound with default headers values
func NewSendVideoNoteNotFound() *SendVideoNoteNotFound {
	return &SendVideoNoteNotFound{}
}

/*SendVideoNoteNotFound handles this case with default header values.

Not Found
*/
type SendVideoNoteNotFound struct {
	Payload *models.Error
}

func (o *SendVideoNoteNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote][%d] sendVideoNoteNotFound  %+v", 404, o.Payload)
}

func (o *SendVideoNoteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteEnhanceYourCalm creates a SendVideoNoteEnhanceYourCalm with default headers values
func NewSendVideoNoteEnhanceYourCalm() *SendVideoNoteEnhanceYourCalm {
	return &SendVideoNoteEnhanceYourCalm{}
}

/*SendVideoNoteEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendVideoNoteEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendVideoNoteEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote][%d] sendVideoNoteEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendVideoNoteEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendVideoNoteInternalServerError creates a SendVideoNoteInternalServerError with default headers values
func NewSendVideoNoteInternalServerError() *SendVideoNoteInternalServerError {
	return &SendVideoNoteInternalServerError{}
}

/*SendVideoNoteInternalServerError handles this case with default header values.

Internal
*/
type SendVideoNoteInternalServerError struct {
	Payload *models.Error
}

func (o *SendVideoNoteInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendVideoNote][%d] sendVideoNoteInternalServerError  %+v", 500, o.Payload)
}

func (o *SendVideoNoteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}