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

// NewVirtualizationVirtualMachinesBulkPartialUpdateParams creates a new VirtualizationVirtualMachinesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationVirtualMachinesBulkPartialUpdateParams() *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	return &VirtualizationVirtualMachinesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationVirtualMachinesBulkPartialUpdateParamsWithTimeout creates a new VirtualizationVirtualMachinesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationVirtualMachinesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	return &VirtualizationVirtualMachinesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationVirtualMachinesBulkPartialUpdateParamsWithContext creates a new VirtualizationVirtualMachinesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationVirtualMachinesBulkPartialUpdateParamsWithContext(ctx context.Context) *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	return &VirtualizationVirtualMachinesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationVirtualMachinesBulkPartialUpdateParamsWithHTTPClient creates a new VirtualizationVirtualMachinesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationVirtualMachinesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	return &VirtualizationVirtualMachinesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationVirtualMachinesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization virtual machines bulk partial update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationVirtualMachinesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableVirtualMachineWithConfigContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization virtual machines bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) WithDefaults() *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization virtual machines bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) WithContext(ctx context.Context) *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) WithData(data *models.WritableVirtualMachineWithConfigContext) *VirtualizationVirtualMachinesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization virtual machines bulk partial update params
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) SetData(data *models.WritableVirtualMachineWithConfigContext) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationVirtualMachinesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
