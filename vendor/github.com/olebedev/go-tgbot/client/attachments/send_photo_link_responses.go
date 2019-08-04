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

// SendPhotoLinkReader is a Reader for the SendPhotoLink structure.
type SendPhotoLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendPhotoLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendPhotoLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewSendPhotoLinkBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewSendPhotoLinkUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendPhotoLinkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewSendPhotoLinkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewSendPhotoLinkEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendPhotoLinkInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendPhotoLinkOK creates a SendPhotoLinkOK with default headers values
func NewSendPhotoLinkOK() *SendPhotoLinkOK {
	return &SendPhotoLinkOK{}
}

/*SendPhotoLinkOK handles this case with default header values.

SendPhotoLinkOK send photo link o k
*/
type SendPhotoLinkOK struct {
	Payload *models.ResponseMessage
}

func (o *SendPhotoLinkOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendPhoto#link][%d] sendPhotoLinkOK  %+v", 200, o.Payload)
}

func (o *SendPhotoLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendPhotoLinkBadRequest creates a SendPhotoLinkBadRequest with default headers values
func NewSendPhotoLinkBadRequest() *SendPhotoLinkBadRequest {
	return &SendPhotoLinkBadRequest{}
}

/*SendPhotoLinkBadRequest handles this case with default header values.

Bad Request
*/
type SendPhotoLinkBadRequest struct {
	Payload *models.Error
}

func (o *SendPhotoLinkBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendPhoto#link][%d] sendPhotoLinkBadRequest  %+v", 400, o.Payload)
}

func (o *SendPhotoLinkBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendPhotoLinkUnauthorized creates a SendPhotoLinkUnauthorized with default headers values
func NewSendPhotoLinkUnauthorized() *SendPhotoLinkUnauthorized {
	return &SendPhotoLinkUnauthorized{}
}

/*SendPhotoLinkUnauthorized handles this case with default header values.

Unauthorized
*/
type SendPhotoLinkUnauthorized struct {
	Payload *models.Error
}

func (o *SendPhotoLinkUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendPhoto#link][%d] sendPhotoLinkUnauthorized  %+v", 401, o.Payload)
}

func (o *SendPhotoLinkUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendPhotoLinkForbidden creates a SendPhotoLinkForbidden with default headers values
func NewSendPhotoLinkForbidden() *SendPhotoLinkForbidden {
	return &SendPhotoLinkForbidden{}
}

/*SendPhotoLinkForbidden handles this case with default header values.

Forbidden
*/
type SendPhotoLinkForbidden struct {
	Payload *models.Error
}

func (o *SendPhotoLinkForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendPhoto#link][%d] sendPhotoLinkForbidden  %+v", 403, o.Payload)
}

func (o *SendPhotoLinkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendPhotoLinkNotFound creates a SendPhotoLinkNotFound with default headers values
func NewSendPhotoLinkNotFound() *SendPhotoLinkNotFound {
	return &SendPhotoLinkNotFound{}
}

/*SendPhotoLinkNotFound handles this case with default header values.

Not Found
*/
type SendPhotoLinkNotFound struct {
	Payload *models.Error
}

func (o *SendPhotoLinkNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendPhoto#link][%d] sendPhotoLinkNotFound  %+v", 404, o.Payload)
}

func (o *SendPhotoLinkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendPhotoLinkEnhanceYourCalm creates a SendPhotoLinkEnhanceYourCalm with default headers values
func NewSendPhotoLinkEnhanceYourCalm() *SendPhotoLinkEnhanceYourCalm {
	return &SendPhotoLinkEnhanceYourCalm{}
}

/*SendPhotoLinkEnhanceYourCalm handles this case with default header values.

Flood
*/
type SendPhotoLinkEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *SendPhotoLinkEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendPhoto#link][%d] sendPhotoLinkEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *SendPhotoLinkEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendPhotoLinkInternalServerError creates a SendPhotoLinkInternalServerError with default headers values
func NewSendPhotoLinkInternalServerError() *SendPhotoLinkInternalServerError {
	return &SendPhotoLinkInternalServerError{}
}

/*SendPhotoLinkInternalServerError handles this case with default header values.

Internal
*/
type SendPhotoLinkInternalServerError struct {
	Payload *models.Error
}

func (o *SendPhotoLinkInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/sendPhoto#link][%d] sendPhotoLinkInternalServerError  %+v", 500, o.Payload)
}

func (o *SendPhotoLinkInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
