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

// NewIpamIPAddressesBulkDeleteParams creates a new IpamIPAddressesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPAddressesBulkDeleteParams() *IpamIPAddressesBulkDeleteParams {
	return &IpamIPAddressesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesBulkDeleteParamsWithTimeout creates a new IpamIPAddressesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamIPAddressesBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamIPAddressesBulkDeleteParams {
	return &IpamIPAddressesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamIPAddressesBulkDeleteParamsWithContext creates a new IpamIPAddressesBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamIPAddressesBulkDeleteParamsWithContext(ctx context.Context) *IpamIPAddressesBulkDeleteParams {
	return &IpamIPAddressesBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamIPAddressesBulkDeleteParamsWithHTTPClient creates a new IpamIPAddressesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPAddressesBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamIPAddressesBulkDeleteParams {
	return &IpamIPAddressesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* IpamIPAddressesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the ipam ip addresses bulk delete operation.

   Typically these are written to a http.Request.
*/
type IpamIPAddressesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip addresses bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesBulkDeleteParams) WithDefaults() *IpamIPAddressesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip addresses bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip addresses bulk delete params
func (o *IpamIPAddressesBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamIPAddressesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses bulk delete params
func (o *IpamIPAddressesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses bulk delete params
func (o *IpamIPAddressesBulkDeleteParams) WithContext(ctx context.Context) *IpamIPAddressesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses bulk delete params
func (o *IpamIPAddressesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses bulk delete params
func (o *IpamIPAddressesBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamIPAddressesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses bulk delete params
func (o *IpamIPAddressesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
