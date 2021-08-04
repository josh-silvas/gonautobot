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

// NewIpamRolesReadParams creates a new IpamRolesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRolesReadParams() *IpamRolesReadParams {
	return &IpamRolesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRolesReadParamsWithTimeout creates a new IpamRolesReadParams object
// with the ability to set a timeout on a request.
func NewIpamRolesReadParamsWithTimeout(timeout time.Duration) *IpamRolesReadParams {
	return &IpamRolesReadParams{
		timeout: timeout,
	}
}

// NewIpamRolesReadParamsWithContext creates a new IpamRolesReadParams object
// with the ability to set a context for a request.
func NewIpamRolesReadParamsWithContext(ctx context.Context) *IpamRolesReadParams {
	return &IpamRolesReadParams{
		Context: ctx,
	}
}

// NewIpamRolesReadParamsWithHTTPClient creates a new IpamRolesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRolesReadParamsWithHTTPClient(client *http.Client) *IpamRolesReadParams {
	return &IpamRolesReadParams{
		HTTPClient: client,
	}
}

/* IpamRolesReadParams contains all the parameters to send to the API endpoint
   for the ipam roles read operation.

   Typically these are written to a http.Request.
*/
type IpamRolesReadParams struct {

	/* ID.

	   A UUID string identifying this role.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam roles read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesReadParams) WithDefaults() *IpamRolesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam roles read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRolesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam roles read params
func (o *IpamRolesReadParams) WithTimeout(timeout time.Duration) *IpamRolesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam roles read params
func (o *IpamRolesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam roles read params
func (o *IpamRolesReadParams) WithContext(ctx context.Context) *IpamRolesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam roles read params
func (o *IpamRolesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam roles read params
func (o *IpamRolesReadParams) WithHTTPClient(client *http.Client) *IpamRolesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam roles read params
func (o *IpamRolesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam roles read params
func (o *IpamRolesReadParams) WithID(id strfmt.UUID) *IpamRolesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam roles read params
func (o *IpamRolesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRolesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
