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

// NewIpamPrefixesBulkDeleteParams creates a new IpamPrefixesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesBulkDeleteParams() *IpamPrefixesBulkDeleteParams {
	return &IpamPrefixesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesBulkDeleteParamsWithTimeout creates a new IpamPrefixesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesBulkDeleteParamsWithTimeout(timeout time.Duration) *IpamPrefixesBulkDeleteParams {
	return &IpamPrefixesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesBulkDeleteParamsWithContext creates a new IpamPrefixesBulkDeleteParams object
// with the ability to set a context for a request.
func NewIpamPrefixesBulkDeleteParamsWithContext(ctx context.Context) *IpamPrefixesBulkDeleteParams {
	return &IpamPrefixesBulkDeleteParams{
		Context: ctx,
	}
}

// NewIpamPrefixesBulkDeleteParamsWithHTTPClient creates a new IpamPrefixesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesBulkDeleteParamsWithHTTPClient(client *http.Client) *IpamPrefixesBulkDeleteParams {
	return &IpamPrefixesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the ipam prefixes bulk delete operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesBulkDeleteParams) WithDefaults() *IpamPrefixesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes bulk delete params
func (o *IpamPrefixesBulkDeleteParams) WithTimeout(timeout time.Duration) *IpamPrefixesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes bulk delete params
func (o *IpamPrefixesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes bulk delete params
func (o *IpamPrefixesBulkDeleteParams) WithContext(ctx context.Context) *IpamPrefixesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes bulk delete params
func (o *IpamPrefixesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes bulk delete params
func (o *IpamPrefixesBulkDeleteParams) WithHTTPClient(client *http.Client) *IpamPrefixesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes bulk delete params
func (o *IpamPrefixesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
