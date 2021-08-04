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

// NewUsersPermissionsReadParams creates a new UsersPermissionsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsReadParams() *UsersPermissionsReadParams {
	return &UsersPermissionsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsReadParamsWithTimeout creates a new UsersPermissionsReadParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsReadParamsWithTimeout(timeout time.Duration) *UsersPermissionsReadParams {
	return &UsersPermissionsReadParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsReadParamsWithContext creates a new UsersPermissionsReadParams object
// with the ability to set a context for a request.
func NewUsersPermissionsReadParamsWithContext(ctx context.Context) *UsersPermissionsReadParams {
	return &UsersPermissionsReadParams{
		Context: ctx,
	}
}

// NewUsersPermissionsReadParamsWithHTTPClient creates a new UsersPermissionsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsReadParamsWithHTTPClient(client *http.Client) *UsersPermissionsReadParams {
	return &UsersPermissionsReadParams{
		HTTPClient: client,
	}
}

/* UsersPermissionsReadParams contains all the parameters to send to the API endpoint
   for the users permissions read operation.

   Typically these are written to a http.Request.
*/
type UsersPermissionsReadParams struct {

	/* ID.

	   A UUID string identifying this permission.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsReadParams) WithDefaults() *UsersPermissionsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions read params
func (o *UsersPermissionsReadParams) WithTimeout(timeout time.Duration) *UsersPermissionsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions read params
func (o *UsersPermissionsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions read params
func (o *UsersPermissionsReadParams) WithContext(ctx context.Context) *UsersPermissionsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions read params
func (o *UsersPermissionsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions read params
func (o *UsersPermissionsReadParams) WithHTTPClient(client *http.Client) *UsersPermissionsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions read params
func (o *UsersPermissionsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the users permissions read params
func (o *UsersPermissionsReadParams) WithID(id strfmt.UUID) *UsersPermissionsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users permissions read params
func (o *UsersPermissionsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
