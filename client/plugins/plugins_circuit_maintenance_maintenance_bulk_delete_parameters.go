package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParams creates a new PluginsCircuitMaintenanceMaintenanceBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParams() *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParamsWithTimeout creates a new PluginsCircuitMaintenanceMaintenanceBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParamsWithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkDeleteParams{
		timeout: timeout,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParamsWithContext creates a new PluginsCircuitMaintenanceMaintenanceBulkDeleteParams object
// with the ability to set a context for a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParamsWithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkDeleteParams{
		Context: ctx,
	}
}

// NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParamsWithHTTPClient creates a new PluginsCircuitMaintenanceMaintenanceBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsCircuitMaintenanceMaintenanceBulkDeleteParamsWithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	return &PluginsCircuitMaintenanceMaintenanceBulkDeleteParams{
		HTTPClient: client,
	}
}

/* PluginsCircuitMaintenanceMaintenanceBulkDeleteParams contains all the parameters to send to the API endpoint
   for the plugins circuit maintenance maintenance bulk delete operation.

   Typically these are written to a http.Request.
*/
type PluginsCircuitMaintenanceMaintenanceBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins circuit maintenance maintenance bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) WithDefaults() *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins circuit maintenance maintenance bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins circuit maintenance maintenance bulk delete params
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) WithTimeout(timeout time.Duration) *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins circuit maintenance maintenance bulk delete params
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins circuit maintenance maintenance bulk delete params
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) WithContext(ctx context.Context) *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins circuit maintenance maintenance bulk delete params
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance bulk delete params
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) WithHTTPClient(client *http.Client) *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins circuit maintenance maintenance bulk delete params
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsCircuitMaintenanceMaintenanceBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
