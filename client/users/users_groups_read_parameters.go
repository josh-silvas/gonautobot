package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUsersGroupsReadParams creates a new UsersGroupsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsReadParams() *UsersGroupsReadParams {
	return &UsersGroupsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsReadParamsWithTimeout creates a new UsersGroupsReadParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsReadParamsWithTimeout(timeout time.Duration) *UsersGroupsReadParams {
	return &UsersGroupsReadParams{
		timeout: timeout,
	}
}

// NewUsersGroupsReadParamsWithContext creates a new UsersGroupsReadParams object
// with the ability to set a context for a request.
func NewUsersGroupsReadParamsWithContext(ctx context.Context) *UsersGroupsReadParams {
	return &UsersGroupsReadParams{
		Context: ctx,
	}
}

// NewUsersGroupsReadParamsWithHTTPClient creates a new UsersGroupsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsReadParamsWithHTTPClient(client *http.Client) *UsersGroupsReadParams {
	return &UsersGroupsReadParams{
		HTTPClient: client,
	}
}

/* UsersGroupsReadParams contains all the parameters to send to the API endpoint
   for the users groups read operation.

   Typically these are written to a http.Request.
*/
type UsersGroupsReadParams struct {

	/* ID.

	   A unique integer value identifying this group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsReadParams) WithDefaults() *UsersGroupsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups read params
func (o *UsersGroupsReadParams) WithTimeout(timeout time.Duration) *UsersGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups read params
func (o *UsersGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups read params
func (o *UsersGroupsReadParams) WithContext(ctx context.Context) *UsersGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups read params
func (o *UsersGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups read params
func (o *UsersGroupsReadParams) WithHTTPClient(client *http.Client) *UsersGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups read params
func (o *UsersGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users groups read params
func (o *UsersGroupsReadParams) WithID(id int64) *UsersGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users groups read params
func (o *UsersGroupsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
