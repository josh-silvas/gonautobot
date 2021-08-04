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

// NewVirtualizationInterfacesBulkUpdateParams creates a new VirtualizationInterfacesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationInterfacesBulkUpdateParams() *VirtualizationInterfacesBulkUpdateParams {
	return &VirtualizationInterfacesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesBulkUpdateParamsWithTimeout creates a new VirtualizationInterfacesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationInterfacesBulkUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesBulkUpdateParams {
	return &VirtualizationInterfacesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationInterfacesBulkUpdateParamsWithContext creates a new VirtualizationInterfacesBulkUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationInterfacesBulkUpdateParamsWithContext(ctx context.Context) *VirtualizationInterfacesBulkUpdateParams {
	return &VirtualizationInterfacesBulkUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationInterfacesBulkUpdateParamsWithHTTPClient creates a new VirtualizationInterfacesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationInterfacesBulkUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesBulkUpdateParams {
	return &VirtualizationInterfacesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationInterfacesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization interfaces bulk update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationInterfacesBulkUpdateParams struct {

	// Data.
	Data *models.WritableVMInterface

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization interfaces bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesBulkUpdateParams) WithDefaults() *VirtualizationInterfacesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization interfaces bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) WithContext(ctx context.Context) *VirtualizationInterfacesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) WithData(data *models.WritableVMInterface) *VirtualizationInterfacesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization interfaces bulk update params
func (o *VirtualizationInterfacesBulkUpdateParams) SetData(data *models.WritableVMInterface) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
