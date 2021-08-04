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

// NewVirtualizationVirtualMachinesDeleteParams creates a new VirtualizationVirtualMachinesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualMachinesDeleteParams() *VirtualizationVirtualMachinesDeleteParams {
	return &VirtualizationVirtualMachinesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesDeleteParamsWithTimeout creates a new VirtualizationVirtualMachinesDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualMachinesDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesDeleteParams {
	return &VirtualizationVirtualMachinesDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesDeleteParamsWithContext creates a new VirtualizationVirtualMachinesDeleteParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualMachinesDeleteParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesDeleteParams {
	return &VirtualizationVirtualMachinesDeleteParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesDeleteParamsWithHTTPClient creates a new VirtualizationVirtualMachinesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualMachinesDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesDeleteParams {
	return &VirtualizationVirtualMachinesDeleteParams{
		HTTPClient: client,
	}
}

/* VirtualizationVirtualMachinesDeleteParams contains all the parameters to send to the API endpoint
   for the virtualization virtual machines delete operation.

   Typically these are written to a http.Request.
*/
type VirtualizationVirtualMachinesDeleteParams struct {

	/* ID.

	   A UUID string identifying this virtual machine.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual machines delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesDeleteParams) WithDefaults() *VirtualizationVirtualMachinesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual machines delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) WithID(id strfmt.UUID) *VirtualizationVirtualMachinesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines delete params
func (o *VirtualizationVirtualMachinesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
