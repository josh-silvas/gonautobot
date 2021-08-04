package virtualization

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewVirtualizationVirtualMachinesBulkDeleteParams creates a new VirtualizationVirtualMachinesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualMachinesBulkDeleteParams() *VirtualizationVirtualMachinesBulkDeleteParams {
	return &VirtualizationVirtualMachinesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesBulkDeleteParamsWithTimeout creates a new VirtualizationVirtualMachinesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualMachinesBulkDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesBulkDeleteParams {
	return &VirtualizationVirtualMachinesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesBulkDeleteParamsWithContext creates a new VirtualizationVirtualMachinesBulkDeleteParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualMachinesBulkDeleteParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesBulkDeleteParams {
	return &VirtualizationVirtualMachinesBulkDeleteParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesBulkDeleteParamsWithHTTPClient creates a new VirtualizationVirtualMachinesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualMachinesBulkDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesBulkDeleteParams {
	return &VirtualizationVirtualMachinesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* VirtualizationVirtualMachinesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the virtualization virtual machines bulk delete operation.

   Typically these are written to a http.Request.
*/
type VirtualizationVirtualMachinesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual machines bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesBulkDeleteParams) WithDefaults() *VirtualizationVirtualMachinesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual machines bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual machines bulk delete params
func (o *VirtualizationVirtualMachinesBulkDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines bulk delete params
func (o *VirtualizationVirtualMachinesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines bulk delete params
func (o *VirtualizationVirtualMachinesBulkDeleteParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines bulk delete params
func (o *VirtualizationVirtualMachinesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines bulk delete params
func (o *VirtualizationVirtualMachinesBulkDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines bulk delete params
func (o *VirtualizationVirtualMachinesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
