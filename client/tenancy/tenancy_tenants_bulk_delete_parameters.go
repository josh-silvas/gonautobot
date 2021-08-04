package tenancy

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTenancyTenantsBulkDeleteParams creates a new TenancyTenantsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsBulkDeleteParams() *TenancyTenantsBulkDeleteParams {
	return &TenancyTenantsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsBulkDeleteParamsWithTimeout creates a new TenancyTenantsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsBulkDeleteParamsWithTimeout(timeout time.Duration) *TenancyTenantsBulkDeleteParams {
	return &TenancyTenantsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsBulkDeleteParamsWithContext creates a new TenancyTenantsBulkDeleteParams object
// with the ability to set a context for a request.
func NewTenancyTenantsBulkDeleteParamsWithContext(ctx context.Context) *TenancyTenantsBulkDeleteParams {
	return &TenancyTenantsBulkDeleteParams{
		Context: ctx,
	}
}

// NewTenancyTenantsBulkDeleteParamsWithHTTPClient creates a new TenancyTenantsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsBulkDeleteParamsWithHTTPClient(client *http.Client) *TenancyTenantsBulkDeleteParams {
	return &TenancyTenantsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the tenancy tenants bulk delete operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsBulkDeleteParams) WithDefaults() *TenancyTenantsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants bulk delete params
func (o *TenancyTenantsBulkDeleteParams) WithTimeout(timeout time.Duration) *TenancyTenantsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants bulk delete params
func (o *TenancyTenantsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants bulk delete params
func (o *TenancyTenantsBulkDeleteParams) WithContext(ctx context.Context) *TenancyTenantsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants bulk delete params
func (o *TenancyTenantsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants bulk delete params
func (o *TenancyTenantsBulkDeleteParams) WithHTTPClient(client *http.Client) *TenancyTenantsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants bulk delete params
func (o *TenancyTenantsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
