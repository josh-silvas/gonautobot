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

// NewUsersUsersBulkDeleteParams creates a new UsersUsersBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersUsersBulkDeleteParams() *UsersUsersBulkDeleteParams {
	return &UsersUsersBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersUsersBulkDeleteParamsWithTimeout creates a new UsersUsersBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewUsersUsersBulkDeleteParamsWithTimeout(timeout time.Duration) *UsersUsersBulkDeleteParams {
	return &UsersUsersBulkDeleteParams{
		timeout: timeout,
	}
}

// NewUsersUsersBulkDeleteParamsWithContext creates a new UsersUsersBulkDeleteParams object
// with the ability to set a context for a request.
func NewUsersUsersBulkDeleteParamsWithContext(ctx context.Context) *UsersUsersBulkDeleteParams {
	return &UsersUsersBulkDeleteParams{
		Context: ctx,
	}
}

// NewUsersUsersBulkDeleteParamsWithHTTPClient creates a new UsersUsersBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersUsersBulkDeleteParamsWithHTTPClient(client *http.Client) *UsersUsersBulkDeleteParams {
	return &UsersUsersBulkDeleteParams{
		HTTPClient: client,
	}
}

/* UsersUsersBulkDeleteParams contains all the parameters to send to the API endpoint
   for the users users bulk delete operation.

   Typically these are written to a http.Request.
*/
type UsersUsersBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users users bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersBulkDeleteParams) WithDefaults() *UsersUsersBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users users bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersUsersBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users users bulk delete params
func (o *UsersUsersBulkDeleteParams) WithTimeout(timeout time.Duration) *UsersUsersBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users users bulk delete params
func (o *UsersUsersBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users users bulk delete params
func (o *UsersUsersBulkDeleteParams) WithContext(ctx context.Context) *UsersUsersBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users users bulk delete params
func (o *UsersUsersBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users users bulk delete params
func (o *UsersUsersBulkDeleteParams) WithHTTPClient(client *http.Client) *UsersUsersBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users users bulk delete params
func (o *UsersUsersBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UsersUsersBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
