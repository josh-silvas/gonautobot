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

// NewIpamRirsBulkDeleteParams creates a new IpamRirsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRirsBulkDeleteParams() *IpamRirsBulkDeleteParams {
	return &IpamRirsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsBulkDeleteParamsWithTimeout creates a new IpamRirsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamRirsBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamRirsBulkDeleteParams {
	return &IpamRirsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamRirsBulkDeleteParamsWithContext creates a new IpamRirsBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamRirsBulkDeleteParamsWithContext(ctx context.Context) *IpamRirsBulkDeleteParams {
	return &IpamRirsBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamRirsBulkDeleteParamsWithHTTPClient creates a new IpamRirsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRirsBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamRirsBulkDeleteParams {
	return &IpamRirsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* IpamRirsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the ipam rirs bulk delete operation.

   Typically these are written to a http.Request.
*/
type IpamRirsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam rirs bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsBulkDeleteParams) WithDefaults() *IpamRirsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam rirs bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam rirs bulk delete params
func (o *IpamRirsBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamRirsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs bulk delete params
func (o *IpamRirsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs bulk delete params
func (o *IpamRirsBulkDeleteParams) WithContext(ctx context.Context) *IpamRirsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs bulk delete params
func (o *IpamRirsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs bulk delete params
func (o *IpamRirsBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamRirsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs bulk delete params
func (o *IpamRirsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
