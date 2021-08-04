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

// NewVirtualizationInterfacesBulkPartialUpdateParams creates a new VirtualizationInterfacesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationInterfacesBulkPartialUpdateParams() *VirtualizationInterfacesBulkPartialUpdateParams {
	return &VirtualizationInterfacesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesBulkPartialUpdateParamsWithTimeout creates a new VirtualizationInterfacesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationInterfacesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesBulkPartialUpdateParams {
	return &VirtualizationInterfacesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationInterfacesBulkPartialUpdateParamsWithContext creates a new VirtualizationInterfacesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationInterfacesBulkPartialUpdateParamsWithContext(ctx context.Context) *VirtualizationInterfacesBulkPartialUpdateParams {
	return &VirtualizationInterfacesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationInterfacesBulkPartialUpdateParamsWithHTTPClient creates a new VirtualizationInterfacesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationInterfacesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesBulkPartialUpdateParams {
	return &VirtualizationInterfacesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationInterfacesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization interfaces bulk partial update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationInterfacesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableVMInterface

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization interfaces bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesBulkPartialUpdateParams) WithDefaults() *VirtualizationInterfacesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization interfaces bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) WithContext(ctx context.Context) *VirtualizationInterfacesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) WithData(data *models.WritableVMInterface) *VirtualizationInterfacesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization interfaces bulk partial update params
func (o *VirtualizationInterfacesBulkPartialUpdateParams) SetData(data *models.WritableVMInterface) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
