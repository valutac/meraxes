// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserProfilePhotosParams creates a new GetUserProfilePhotosParams object
// with the default values initialized.
func NewGetUserProfilePhotosParams() *GetUserProfilePhotosParams {
	var ()
	return &GetUserProfilePhotosParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserProfilePhotosParamsWithTimeout creates a new GetUserProfilePhotosParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserProfilePhotosParamsWithTimeout(timeout time.Duration) *GetUserProfilePhotosParams {
	var ()
	return &GetUserProfilePhotosParams{

		timeout: timeout,
	}
}

// NewGetUserProfilePhotosParamsWithContext creates a new GetUserProfilePhotosParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserProfilePhotosParamsWithContext(ctx context.Context) *GetUserProfilePhotosParams {
	var ()
	return &GetUserProfilePhotosParams{

		Context: ctx,
	}
}

// NewGetUserProfilePhotosParamsWithHTTPClient creates a new GetUserProfilePhotosParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserProfilePhotosParamsWithHTTPClient(client *http.Client) *GetUserProfilePhotosParams {
	var ()
	return &GetUserProfilePhotosParams{
		HTTPClient: client,
	}
}

/*GetUserProfilePhotosParams contains all the parameters to send to the API endpoint
for the get user profile photos operation typically these are written to a http.Request
*/
type GetUserProfilePhotosParams struct {

	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*Token
	  bot's token to authorize the request

	*/
	Token *string
	/*UserID*/
	UserID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user profile photos params
func (o *GetUserProfilePhotosParams) WithTimeout(timeout time.Duration) *GetUserProfilePhotosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user profile photos params
func (o *GetUserProfilePhotosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user profile photos params
func (o *GetUserProfilePhotosParams) WithContext(ctx context.Context) *GetUserProfilePhotosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user profile photos params
func (o *GetUserProfilePhotosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user profile photos params
func (o *GetUserProfilePhotosParams) WithHTTPClient(client *http.Client) *GetUserProfilePhotosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user profile photos params
func (o *GetUserProfilePhotosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get user profile photos params
func (o *GetUserProfilePhotosParams) WithLimit(limit *int64) *GetUserProfilePhotosParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get user profile photos params
func (o *GetUserProfilePhotosParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get user profile photos params
func (o *GetUserProfilePhotosParams) WithOffset(offset *int64) *GetUserProfilePhotosParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get user profile photos params
func (o *GetUserProfilePhotosParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithToken adds the token to the get user profile photos params
func (o *GetUserProfilePhotosParams) WithToken(token *string) *GetUserProfilePhotosParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get user profile photos params
func (o *GetUserProfilePhotosParams) SetToken(token *string) {
	o.Token = token
}

// WithUserID adds the userID to the get user profile photos params
func (o *GetUserProfilePhotosParams) WithUserID(userID int64) *GetUserProfilePhotosParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get user profile photos params
func (o *GetUserProfilePhotosParams) SetUserID(userID int64) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserProfilePhotosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Token != nil {

		// path param token
		if err := r.SetPathParam("token", *o.Token); err != nil {
			return err
		}

	}

	// query param user_id
	qrUserID := o.UserID
	qUserID := swag.FormatInt64(qrUserID)
	if qUserID != "" {
		if err := r.SetQueryParam("user_id", qUserID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}