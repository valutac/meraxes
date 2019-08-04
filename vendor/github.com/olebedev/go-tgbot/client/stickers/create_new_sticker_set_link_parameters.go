// Code generated by go-swagger; DO NOT EDIT.

package stickers

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

// NewCreateNewStickerSetLinkParams creates a new CreateNewStickerSetLinkParams object
// with the default values initialized.
func NewCreateNewStickerSetLinkParams() *CreateNewStickerSetLinkParams {
	var ()
	return &CreateNewStickerSetLinkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateNewStickerSetLinkParamsWithTimeout creates a new CreateNewStickerSetLinkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateNewStickerSetLinkParamsWithTimeout(timeout time.Duration) *CreateNewStickerSetLinkParams {
	var ()
	return &CreateNewStickerSetLinkParams{

		timeout: timeout,
	}
}

// NewCreateNewStickerSetLinkParamsWithContext creates a new CreateNewStickerSetLinkParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateNewStickerSetLinkParamsWithContext(ctx context.Context) *CreateNewStickerSetLinkParams {
	var ()
	return &CreateNewStickerSetLinkParams{

		Context: ctx,
	}
}

// NewCreateNewStickerSetLinkParamsWithHTTPClient creates a new CreateNewStickerSetLinkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateNewStickerSetLinkParamsWithHTTPClient(client *http.Client) *CreateNewStickerSetLinkParams {
	var ()
	return &CreateNewStickerSetLinkParams{
		HTTPClient: client,
	}
}

/*CreateNewStickerSetLinkParams contains all the parameters to send to the API endpoint
for the create new sticker set link operation typically these are written to a http.Request
*/
type CreateNewStickerSetLinkParams struct {

	/*Body*/
	Body *models.CreateNewStickerSetLinkBody
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) WithTimeout(timeout time.Duration) *CreateNewStickerSetLinkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) WithContext(ctx context.Context) *CreateNewStickerSetLinkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) WithHTTPClient(client *http.Client) *CreateNewStickerSetLinkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) WithBody(body *models.CreateNewStickerSetLinkBody) *CreateNewStickerSetLinkParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) SetBody(body *models.CreateNewStickerSetLinkBody) {
	o.Body = body
}

// WithToken adds the token to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) WithToken(token *string) *CreateNewStickerSetLinkParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the create new sticker set link params
func (o *CreateNewStickerSetLinkParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *CreateNewStickerSetLinkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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