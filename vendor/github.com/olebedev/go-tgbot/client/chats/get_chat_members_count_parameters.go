// Code generated by go-swagger; DO NOT EDIT.

package chats

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
)

// NewGetChatMembersCountParams creates a new GetChatMembersCountParams object
// with the default values initialized.
func NewGetChatMembersCountParams() *GetChatMembersCountParams {
	var ()
	return &GetChatMembersCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetChatMembersCountParamsWithTimeout creates a new GetChatMembersCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetChatMembersCountParamsWithTimeout(timeout time.Duration) *GetChatMembersCountParams {
	var ()
	return &GetChatMembersCountParams{

		timeout: timeout,
	}
}

// NewGetChatMembersCountParamsWithContext creates a new GetChatMembersCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetChatMembersCountParamsWithContext(ctx context.Context) *GetChatMembersCountParams {
	var ()
	return &GetChatMembersCountParams{

		Context: ctx,
	}
}

// NewGetChatMembersCountParamsWithHTTPClient creates a new GetChatMembersCountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetChatMembersCountParamsWithHTTPClient(client *http.Client) *GetChatMembersCountParams {
	var ()
	return &GetChatMembersCountParams{
		HTTPClient: client,
	}
}

/*GetChatMembersCountParams contains all the parameters to send to the API endpoint
for the get chat members count operation typically these are written to a http.Request
*/
type GetChatMembersCountParams struct {

	/*ChatID*/
	ChatID string
	/*Token
	  bot's token to authorize the request

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get chat members count params
func (o *GetChatMembersCountParams) WithTimeout(timeout time.Duration) *GetChatMembersCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get chat members count params
func (o *GetChatMembersCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get chat members count params
func (o *GetChatMembersCountParams) WithContext(ctx context.Context) *GetChatMembersCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get chat members count params
func (o *GetChatMembersCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get chat members count params
func (o *GetChatMembersCountParams) WithHTTPClient(client *http.Client) *GetChatMembersCountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get chat members count params
func (o *GetChatMembersCountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChatID adds the chatID to the get chat members count params
func (o *GetChatMembersCountParams) WithChatID(chatID string) *GetChatMembersCountParams {
	o.SetChatID(chatID)
	return o
}

// SetChatID adds the chatId to the get chat members count params
func (o *GetChatMembersCountParams) SetChatID(chatID string) {
	o.ChatID = chatID
}

// WithToken adds the token to the get chat members count params
func (o *GetChatMembersCountParams) WithToken(token *string) *GetChatMembersCountParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get chat members count params
func (o *GetChatMembersCountParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetChatMembersCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param chat_id
	qrChatID := o.ChatID
	qChatID := qrChatID
	if qChatID != "" {
		if err := r.SetQueryParam("chat_id", qChatID); err != nil {
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