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

// NewIpamVlansReadParams creates a new IpamVlansReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansReadParams() *IpamVlansReadParams {
	return &IpamVlansReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansReadParamsWithTimeout creates a new IpamVlansReadParams object
// with the ability to set a timeout on a request.
func NewIpamVlansReadParamsWithTimeout(timeout time.Duration) *IpamVlansReadParams {
	return &IpamVlansReadParams{
		timeout: timeout,
	}
}

// NewIpamVlansReadParamsWithContext creates a new IpamVlansReadParams object
// with the ability to set a context for a request.
func NewIpamVlansReadParamsWithContext(ctx context.Context) *IpamVlansReadParams {
	return &IpamVlansReadParams{
		Context: ctx,
	}
}

// NewIpamVlansReadParamsWithHTTPClient creates a new IpamVlansReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansReadParamsWithHTTPClient(client *http.Client) *IpamVlansReadParams {
	return &IpamVlansReadParams{
		HTTPClient: client,
	}
}

/* IpamVlansReadParams contains all the parameters to send to the API endpoint
   for the ipam vlans read operation.

   Typically these are written to a http.Request.
*/
type IpamVlansReadParams struct {

	/* ID.

	   A UUID string identifying this VLAN.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansReadParams) WithDefaults() *IpamVlansReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans read params
func (o *IpamVlansReadParams) WithTimeout(timeout time.Duration) *IpamVlansReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans read params
func (o *IpamVlansReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans read params
func (o *IpamVlansReadParams) WithContext(ctx context.Context) *IpamVlansReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans read params
func (o *IpamVlansReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans read params
func (o *IpamVlansReadParams) WithHTTPClient(client *http.Client) *IpamVlansReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans read params
func (o *IpamVlansReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam vlans read params
func (o *IpamVlansReadParams) WithID(id strfmt.UUID) *IpamVlansReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlans read params
func (o *IpamVlansReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
