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

// EditMessageLiveLocationReader is a Reader for the EditMessageLiveLocation structure.
type EditMessageLiveLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EditMessageLiveLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewEditMessageLiveLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEditMessageLiveLocationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewEditMessageLiveLocationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewEditMessageLiveLocationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEditMessageLiveLocationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewEditMessageLiveLocationEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewEditMessageLiveLocationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEditMessageLiveLocationOK creates a EditMessageLiveLocationOK with default headers values
func NewEditMessageLiveLocationOK() *EditMessageLiveLocationOK {
	return &EditMessageLiveLocationOK{}
}

/*EditMessageLiveLocationOK handles this case with default header values.

EditMessageLiveLocationOK edit message live location o k
*/
type EditMessageLiveLocationOK struct {
	Payload *models.ResponseMessage
}

func (o *EditMessageLiveLocationOK) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageLiveLocation][%d] editMessageLiveLocationOK  %+v", 200, o.Payload)
}

func (o *EditMessageLiveLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageLiveLocationBadRequest creates a EditMessageLiveLocationBadRequest with default headers values
func NewEditMessageLiveLocationBadRequest() *EditMessageLiveLocationBadRequest {
	return &EditMessageLiveLocationBadRequest{}
}

/*EditMessageLiveLocationBadRequest handles this case with default header values.

Bad Request
*/
type EditMessageLiveLocationBadRequest struct {
	Payload *models.Error
}

func (o *EditMessageLiveLocationBadRequest) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageLiveLocation][%d] editMessageLiveLocationBadRequest  %+v", 400, o.Payload)
}

func (o *EditMessageLiveLocationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageLiveLocationUnauthorized creates a EditMessageLiveLocationUnauthorized with default headers values
func NewEditMessageLiveLocationUnauthorized() *EditMessageLiveLocationUnauthorized {
	return &EditMessageLiveLocationUnauthorized{}
}

/*EditMessageLiveLocationUnauthorized handles this case with default header values.

Unauthorized
*/
type EditMessageLiveLocationUnauthorized struct {
	Payload *models.Error
}

func (o *EditMessageLiveLocationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageLiveLocation][%d] editMessageLiveLocationUnauthorized  %+v", 401, o.Payload)
}

func (o *EditMessageLiveLocationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageLiveLocationForbidden creates a EditMessageLiveLocationForbidden with default headers values
func NewEditMessageLiveLocationForbidden() *EditMessageLiveLocationForbidden {
	return &EditMessageLiveLocationForbidden{}
}

/*EditMessageLiveLocationForbidden handles this case with default header values.

Forbidden
*/
type EditMessageLiveLocationForbidden struct {
	Payload *models.Error
}

func (o *EditMessageLiveLocationForbidden) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageLiveLocation][%d] editMessageLiveLocationForbidden  %+v", 403, o.Payload)
}

func (o *EditMessageLiveLocationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageLiveLocationNotFound creates a EditMessageLiveLocationNotFound with default headers values
func NewEditMessageLiveLocationNotFound() *EditMessageLiveLocationNotFound {
	return &EditMessageLiveLocationNotFound{}
}

/*EditMessageLiveLocationNotFound handles this case with default header values.

Not Found
*/
type EditMessageLiveLocationNotFound struct {
	Payload *models.Error
}

func (o *EditMessageLiveLocationNotFound) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageLiveLocation][%d] editMessageLiveLocationNotFound  %+v", 404, o.Payload)
}

func (o *EditMessageLiveLocationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageLiveLocationEnhanceYourCalm creates a EditMessageLiveLocationEnhanceYourCalm with default headers values
func NewEditMessageLiveLocationEnhanceYourCalm() *EditMessageLiveLocationEnhanceYourCalm {
	return &EditMessageLiveLocationEnhanceYourCalm{}
}

/*EditMessageLiveLocationEnhanceYourCalm handles this case with default header values.

Flood
*/
type EditMessageLiveLocationEnhanceYourCalm struct {
	Payload *models.Error
}

func (o *EditMessageLiveLocationEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageLiveLocation][%d] editMessageLiveLocationEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *EditMessageLiveLocationEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEditMessageLiveLocationInternalServerError creates a EditMessageLiveLocationInternalServerError with default headers values
func NewEditMessageLiveLocationInternalServerError() *EditMessageLiveLocationInternalServerError {
	return &EditMessageLiveLocationInternalServerError{}
}

/*EditMessageLiveLocationInternalServerError handles this case with default header values.

Internal
*/
type EditMessageLiveLocationInternalServerError struct {
	Payload *models.Error
}

func (o *EditMessageLiveLocationInternalServerError) Error() string {
	return fmt.Sprintf("[POST /bot{token}/editMessageLiveLocation][%d] editMessageLiveLocationInternalServerError  %+v", 500, o.Payload)
}

func (o *EditMessageLiveLocationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
