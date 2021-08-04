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

// NewVirtualizationVirtualMachinesReadParams creates a new VirtualizationVirtualMachinesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualMachinesReadParams() *VirtualizationVirtualMachinesReadParams {
	return &VirtualizationVirtualMachinesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithTimeout creates a new VirtualizationVirtualMachinesReadParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualMachinesReadParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesReadParams {
	return &VirtualizationVirtualMachinesReadParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithContext creates a new VirtualizationVirtualMachinesReadParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualMachinesReadParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesReadParams {
	return &VirtualizationVirtualMachinesReadParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesReadParamsWithHTTPClient creates a new VirtualizationVirtualMachinesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualMachinesReadParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesReadParams {
	return &VirtualizationVirtualMachinesReadParams{
		HTTPClient: client,
	}
}

/* VirtualizationVirtualMachinesReadParams contains all the parameters to send to the API endpoint
   for the virtualization virtual machines read operation.

   Typically these are written to a http.Request.
*/
type VirtualizationVirtualMachinesReadParams struct {

	/* ID.

	   A UUID string identifying this virtual machine.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual machines read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesReadParams) WithDefaults() *VirtualizationVirtualMachinesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual machines read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) WithID(id strfmt.UUID) *VirtualizationVirtualMachinesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines read params
func (o *VirtualizationVirtualMachinesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
