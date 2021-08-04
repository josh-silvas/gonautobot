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

// NewIpamVlansBulkDeleteParams creates a new IpamVlansBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansBulkDeleteParams() *IpamVlansBulkDeleteParams {
	return &IpamVlansBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansBulkDeleteParamsWithTimeout creates a new IpamVlansBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamVlansBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamVlansBulkDeleteParams {
	return &IpamVlansBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamVlansBulkDeleteParamsWithContext creates a new IpamVlansBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamVlansBulkDeleteParamsWithContext(ctx context.Context) *IpamVlansBulkDeleteParams {
	return &IpamVlansBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamVlansBulkDeleteParamsWithHTTPClient creates a new IpamVlansBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamVlansBulkDeleteParams {
	return &IpamVlansBulkDeleteParams{
		HTTPClient: client,
	}
}

/* IpamVlansBulkDeleteParams contains all the parameters to send to the API endpoint
   for the ipam vlans bulk delete operation.

   Typically these are written to a http.Request.
*/
type IpamVlansBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansBulkDeleteParams) WithDefaults() *IpamVlansBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans bulk delete params
func (o *IpamVlansBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamVlansBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans bulk delete params
func (o *IpamVlansBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans bulk delete params
func (o *IpamVlansBulkDeleteParams) WithContext(ctx context.Context) *IpamVlansBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans bulk delete params
func (o *IpamVlansBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans bulk delete params
func (o *IpamVlansBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamVlansBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans bulk delete params
func (o *IpamVlansBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
