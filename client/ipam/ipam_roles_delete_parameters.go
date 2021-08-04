package ipam

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIpamRolesDeleteParams creates a new IpamRolesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRolesDeleteParams() *IpamRolesDeleteParams {
	return &IpamRolesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRolesDeleteParamsWithTimeout creates a new IpamRolesDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamRolesDeleteParamsWithTimeout(timeout time.Duration) *IpamRolesDeleteParams {
	return &IpamRolesDeleteParams{
		timeout: timeout,
	}
}

// NewIpamRolesDeleteParamsWithContext creates a new IpamRolesDeleteParams object
// with the ability to set a context for a request.
func NewIpamRolesDeleteParamsWithContext(ctx context.Context) *IpamRolesDeleteParams {
	return &IpamRolesDeleteParams{
		Context: ctx,
	}
}

// NewIpamRolesDeleteParamsWithHTTPClient creates a new IpamRolesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRolesDeleteParamsWithHTTPClient(client *http.Client) *IpamRolesDeleteParams {
	return &IpamRolesDeleteParams{
		HTTPClient: client,
	}
}

/* IpamRolesDeleteParams contains all the parameters to send to the API endpoint
   for the ipam roles delete operation.

   Typically these are written to a http.Request.
*/
type IpamRolesDeleteParams struct {

	/* ID.

	   A UUID string identifying this role.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesDeleteParams) WithDefaults() *IpamRolesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam roles delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam roles delete params
func (o *IpamRolesDeleteParams) WithTimeout(timeout time.Duration) *IpamRolesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles delete params
func (o *IpamRolesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles delete params
func (o *IpamRolesDeleteParams) WithContext(ctx context.Context) *IpamRolesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles delete params
func (o *IpamRolesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles delete params
func (o *IpamRolesDeleteParams) WithHTTPClient(client *http.Client) *IpamRolesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles delete params
func (o *IpamRolesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam roles delete params
func (o *IpamRolesDeleteParams) WithID(id strfmt.UUID) *IpamRolesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam roles delete params
func (o *IpamRolesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRolesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
