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

// NewIpamVlanGroupsBulkDeleteParams creates a new IpamVlanGroupsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsBulkDeleteParams() *IpamVlanGroupsBulkDeleteParams {
	return &IpamVlanGroupsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsBulkDeleteParamsWithTimeout creates a new IpamVlanGroupsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsBulkDeleteParams {
	return &IpamVlanGroupsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsBulkDeleteParamsWithContext creates a new IpamVlanGroupsBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsBulkDeleteParamsWithContext(ctx context.Context) *IpamVlanGroupsBulkDeleteParams {
	return &IpamVlanGroupsBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsBulkDeleteParamsWithHTTPClient creates a new IpamVlanGroupsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsBulkDeleteParams {
	return &IpamVlanGroupsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* IpamVlanGroupsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the ipam vlan groups bulk delete operation.

   Typically these are written to a http.Request.
*/
type IpamVlanGroupsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsBulkDeleteParams) WithDefaults() *IpamVlanGroupsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups bulk delete params
func (o *IpamVlanGroupsBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups bulk delete params
func (o *IpamVlanGroupsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups bulk delete params
func (o *IpamVlanGroupsBulkDeleteParams) WithContext(ctx context.Context) *IpamVlanGroupsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups bulk delete params
func (o *IpamVlanGroupsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups bulk delete params
func (o *IpamVlanGroupsBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups bulk delete params
func (o *IpamVlanGroupsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
