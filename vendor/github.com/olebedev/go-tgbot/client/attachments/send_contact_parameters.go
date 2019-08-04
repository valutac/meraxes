// Code generated by go-swagger; DO NOT EDIT.

package attachments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/olebedev/go-tgbot/models"
)

// NewSendContactParams creates a new SendContactParams object
// with the default values initialized.
func NewSendContactParams() *SendContactParams {
	var ()
	return &SendContactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendContactParamsWithTimeout creates a new SendContactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendContactParamsWithTimeout(timeout time.Duration) *SendContactParams {
	var ()
	return &SendContactParams{

		timeout: timeout,
	}
}

// NewSendContactParamsWithContext creates a new SendContactParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendContactParamsWithContext(ctx context.Context) *SendContactParams {
	var ()
	return &SendContactParams{

		Context: ctx,
	}
}

// NewSendContactParamsWithHTTPClient creates a new SendContactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendContactParamsWithHTTPClient(client *http.Client) *SendContactParams {
	var ()
	return &SendContactParams{
		HTTPClient: client,
	}
}

/*SendContactParams contains all the parameters to send to the API endpoint
for the send contact operation typically these are written to a http.Request
*/
type SendContactParams struct {

	/*Body*/
	Body *models.SendContactBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send contact params
func (o *SendContactParams) WithTimeout(timeout time.Duration) *SendContactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send contact params
func (o *SendContactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send contact params
func (o *SendContactParams) WithContext(ctx context.Context) *SendContactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send contact params
func (o *SendContactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send contact params
func (o *SendContactParams) WithHTTPClient(client *http.Client) *SendContactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send contact params
func (o *SendContactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send contact params
func (o *SendContactParams) WithBody(body *models.SendContactBody) *SendContactParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send contact params
func (o *SendContactParams) SetBody(body *models.SendContactBody) {
	o.Body = body
}

// WithToken adds the token to the send contact params
func (o *SendContactParams) WithToken(token *string) *SendContactParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the send contact params
func (o *SendContactParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *SendContactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
