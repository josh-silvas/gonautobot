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

// NewUsersPermissionsDeleteParams creates a new UsersPermissionsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsDeleteParams() *UsersPermissionsDeleteParams {
	return &UsersPermissionsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsDeleteParamsWithTimeout creates a new UsersPermissionsDeleteParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsDeleteParamsWithTimeout(timeout time.Duration) *UsersPermissionsDeleteParams {
	return &UsersPermissionsDeleteParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsDeleteParamsWithContext creates a new UsersPermissionsDeleteParams object
// with the ability to set a context for a request.
func NewUsersPermissionsDeleteParamsWithContext(ctx context.Context) *UsersPermissionsDeleteParams {
	return &UsersPermissionsDeleteParams{
		Context: ctx,
	}
}

// NewUsersPermissionsDeleteParamsWithHTTPClient creates a new UsersPermissionsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsDeleteParamsWithHTTPClient(client *http.Client) *UsersPermissionsDeleteParams {
	return &UsersPermissionsDeleteParams{
		HTTPClient: client,
	}
}

/* UsersPermissionsDeleteParams contains all the parameters to send to the API endpoint
   for the users permissions delete operation.

   Typically these are written to a http.Request.
*/
type UsersPermissionsDeleteParams struct {

	/* ID.

	   A UUID string identifying this permission.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsDeleteParams) WithDefaults() *UsersPermissionsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions delete params
func (o *UsersPermissionsDeleteParams) WithTimeout(timeout time.Duration) *UsersPermissionsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions delete params
func (o *UsersPermissionsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions delete params
func (o *UsersPermissionsDeleteParams) WithContext(ctx context.Context) *UsersPermissionsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions delete params
func (o *UsersPermissionsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions delete params
func (o *UsersPermissionsDeleteParams) WithHTTPClient(client *http.Client) *UsersPermissionsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions delete params
func (o *UsersPermissionsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users permissions delete params
func (o *UsersPermissionsDeleteParams) WithID(id strfmt.UUID) *UsersPermissionsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users permissions delete params
func (o *UsersPermissionsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
