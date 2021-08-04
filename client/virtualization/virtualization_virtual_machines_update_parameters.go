package virtualization

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewVirtualizationVirtualMachinesUpdateParams creates a new VirtualizationVirtualMachinesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualMachinesUpdateParams() *VirtualizationVirtualMachinesUpdateParams {
	return &VirtualizationVirtualMachinesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesUpdateParamsWithTimeout creates a new VirtualizationVirtualMachinesUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualMachinesUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesUpdateParams {
	return &VirtualizationVirtualMachinesUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesUpdateParamsWithContext creates a new VirtualizationVirtualMachinesUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualMachinesUpdateParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesUpdateParams {
	return &VirtualizationVirtualMachinesUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesUpdateParamsWithHTTPClient creates a new VirtualizationVirtualMachinesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualMachinesUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesUpdateParams {
	return &VirtualizationVirtualMachinesUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationVirtualMachinesUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization virtual machines update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationVirtualMachinesUpdateParams struct {

	// Data.
	Data *models.WritableVirtualMachineWithConfigContext

	/* ID.

	   A UUID string identifying this virtual machine.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual machines update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesUpdateParams) WithDefaults() *VirtualizationVirtualMachinesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual machines update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithData(data *models.WritableVirtualMachineWithConfigContext) *VirtualizationVirtualMachinesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetData(data *models.WritableVirtualMachineWithConfigContext) {
	o.Data = data
}

// WithID adds the id to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) WithID(id strfmt.UUID) *VirtualizationVirtualMachinesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization virtual machines update params
func (o *VirtualizationVirtualMachinesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
