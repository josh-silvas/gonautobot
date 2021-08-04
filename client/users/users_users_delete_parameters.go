package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUsersUsersDeleteParams creates a new UsersUsersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersDeleteParams() *UsersUsersDeleteParams {
	return &UsersUsersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersDeleteParamsWithTimeout creates a new UsersUsersDeleteParams object
// with the ability to set a timeout on a request.
func NewUsersUsersDeleteParamsWithTimeout(timeout time.Duration) *UsersUsersDeleteParams {
	return &UsersUsersDeleteParams{
		timeout: timeout,
	}
}

// NewUsersUsersDeleteParamsWithContext creates a new UsersUsersDeleteParams object
// with the ability to set a context for a request.
func NewUsersUsersDeleteParamsWithContext(ctx context.Context) *UsersUsersDeleteParams {
	return &UsersUsersDeleteParams{
		Context: ctx,
	}
}

// NewUsersUsersDeleteParamsWithHTTPClient creates a new UsersUsersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersDeleteParamsWithHTTPClient(client *http.Client) *UsersUsersDeleteParams {
	return &UsersUsersDeleteParams{
		HTTPClient: client,
	}
}

/* UsersUsersDeleteParams contains all the parameters to send to the API endpoint
   for the users users delete operation.

   Typically these are written to a http.Request.
*/
type UsersUsersDeleteParams struct {

	/* ID.

	   A UUID string identifying this user.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersDeleteParams) WithDefaults() *UsersUsersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users delete params
func (o *UsersUsersDeleteParams) WithTimeout(timeout time.Duration) *UsersUsersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users delete params
func (o *UsersUsersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users delete params
func (o *UsersUsersDeleteParams) WithContext(ctx context.Context) *UsersUsersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users delete params
func (o *UsersUsersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users delete params
func (o *UsersUsersDeleteParams) WithHTTPClient(client *http.Client) *UsersUsersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users delete params
func (o *UsersUsersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users users delete params
func (o *UsersUsersDeleteParams) WithID(id strfmt.UUID) *UsersUsersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users delete params
func (o *UsersUsersDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
