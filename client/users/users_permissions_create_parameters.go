package users

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewUsersPermissionsCreateParams creates a new UsersPermissionsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersPermissionsCreateParams() *UsersPermissionsCreateParams {
	return &UsersPermissionsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPermissionsCreateParamsWithTimeout creates a new UsersPermissionsCreateParams object
// with the ability to set a timeout on a request.
func NewUsersPermissionsCreateParamsWithTimeout(timeout time.Duration) *UsersPermissionsCreateParams {
	return &UsersPermissionsCreateParams{
		timeout: timeout,
	}
}

// NewUsersPermissionsCreateParamsWithContext creates a new UsersPermissionsCreateParams object
// with the ability to set a context for a request.
func NewUsersPermissionsCreateParamsWithContext(ctx context.Context) *UsersPermissionsCreateParams {
	return &UsersPermissionsCreateParams{
		Context: ctx,
	}
}

// NewUsersPermissionsCreateParamsWithHTTPClient creates a new UsersPermissionsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersPermissionsCreateParamsWithHTTPClient(client *http.Client) *UsersPermissionsCreateParams {
	return &UsersPermissionsCreateParams{
		HTTPClient: client,
	}
}

/* UsersPermissionsCreateParams contains all the parameters to send to the API endpoint
   for the users permissions create operation.

   Typically these are written to a http.Request.
*/
type UsersPermissionsCreateParams struct {

	// Data.
	Data *models.WritableObjectPermission

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users permissions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsCreateParams) WithDefaults() *UsersPermissionsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users permissions create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersPermissionsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users permissions create params
func (o *UsersPermissionsCreateParams) WithTimeout(timeout time.Duration) *UsersPermissionsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users permissions create params
func (o *UsersPermissionsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users permissions create params
func (o *UsersPermissionsCreateParams) WithContext(ctx context.Context) *UsersPermissionsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users permissions create params
func (o *UsersPermissionsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users permissions create params
func (o *UsersPermissionsCreateParams) WithHTTPClient(client *http.Client) *UsersPermissionsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users permissions create params
func (o *UsersPermissionsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users permissions create params
func (o *UsersPermissionsCreateParams) WithData(data *models.WritableObjectPermission) *UsersPermissionsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users permissions create params
func (o *UsersPermissionsCreateParams) SetData(data *models.WritableObjectPermission) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPermissionsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
