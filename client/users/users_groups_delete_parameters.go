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

// NewUsersGroupsDeleteParams creates a new UsersGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersGroupsDeleteParams() *UsersGroupsDeleteParams {
	return &UsersGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersGroupsDeleteParamsWithTimeout creates a new UsersGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewUsersGroupsDeleteParamsWithTimeout(timeout time.Duration) *UsersGroupsDeleteParams {
	return &UsersGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewUsersGroupsDeleteParamsWithContext creates a new UsersGroupsDeleteParams object
// with the ability to set a context for a request.
func NewUsersGroupsDeleteParamsWithContext(ctx context.Context) *UsersGroupsDeleteParams {
	return &UsersGroupsDeleteParams{
		Context: ctx,
	}
}

// NewUsersGroupsDeleteParamsWithHTTPClient creates a new UsersGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersGroupsDeleteParamsWithHTTPClient(client *http.Client) *UsersGroupsDeleteParams {
	return &UsersGroupsDeleteParams{
		HTTPClient: client,
	}
}

/* UsersGroupsDeleteParams contains all the parameters to send to the API endpoint
   for the users groups delete operation.

   Typically these are written to a http.Request.
*/
type UsersGroupsDeleteParams struct {

	/* ID.

	   A unique integer value identifying this group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsDeleteParams) WithDefaults() *UsersGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users groups delete params
func (o *UsersGroupsDeleteParams) WithTimeout(timeout time.Duration) *UsersGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users groups delete params
func (o *UsersGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users groups delete params
func (o *UsersGroupsDeleteParams) WithContext(ctx context.Context) *UsersGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users groups delete params
func (o *UsersGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users groups delete params
func (o *UsersGroupsDeleteParams) WithHTTPClient(client *http.Client) *UsersGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users groups delete params
func (o *UsersGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users groups delete params
func (o *UsersGroupsDeleteParams) WithID(id int64) *UsersGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users groups delete params
func (o *UsersGroupsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
