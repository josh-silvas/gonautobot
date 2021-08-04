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

// NewUsersUsersReadParams creates a new UsersUsersReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersReadParams() *UsersUsersReadParams {
	return &UsersUsersReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersReadParamsWithTimeout creates a new UsersUsersReadParams object
// with the ability to set a timeout on a request.
func NewUsersUsersReadParamsWithTimeout(timeout time.Duration) *UsersUsersReadParams {
	return &UsersUsersReadParams{
		timeout: timeout,
	}
}

// NewUsersUsersReadParamsWithContext creates a new UsersUsersReadParams object
// with the ability to set a context for a request.
func NewUsersUsersReadParamsWithContext(ctx context.Context) *UsersUsersReadParams {
	return &UsersUsersReadParams{
		Context: ctx,
	}
}

// NewUsersUsersReadParamsWithHTTPClient creates a new UsersUsersReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersReadParamsWithHTTPClient(client *http.Client) *UsersUsersReadParams {
	return &UsersUsersReadParams{
		HTTPClient: client,
	}
}

/* UsersUsersReadParams contains all the parameters to send to the API endpoint
   for the users users read operation.

   Typically these are written to a http.Request.
*/
type UsersUsersReadParams struct {

	/* ID.

	   A UUID string identifying this user.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersReadParams) WithDefaults() *UsersUsersReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users read params
func (o *UsersUsersReadParams) WithTimeout(timeout time.Duration) *UsersUsersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users read params
func (o *UsersUsersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users read params
func (o *UsersUsersReadParams) WithContext(ctx context.Context) *UsersUsersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users read params
func (o *UsersUsersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users read params
func (o *UsersUsersReadParams) WithHTTPClient(client *http.Client) *UsersUsersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users read params
func (o *UsersUsersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users users read params
func (o *UsersUsersReadParams) WithID(id strfmt.UUID) *UsersUsersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users users read params
func (o *UsersUsersReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
