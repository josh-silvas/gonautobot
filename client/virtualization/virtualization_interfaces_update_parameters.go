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

// NewVirtualizationInterfacesUpdateParams creates a new VirtualizationInterfacesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationInterfacesUpdateParams() *VirtualizationInterfacesUpdateParams {
	return &VirtualizationInterfacesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesUpdateParamsWithTimeout creates a new VirtualizationInterfacesUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationInterfacesUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesUpdateParams {
	return &VirtualizationInterfacesUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationInterfacesUpdateParamsWithContext creates a new VirtualizationInterfacesUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationInterfacesUpdateParamsWithContext(ctx context.Context) *VirtualizationInterfacesUpdateParams {
	return &VirtualizationInterfacesUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationInterfacesUpdateParamsWithHTTPClient creates a new VirtualizationInterfacesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationInterfacesUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesUpdateParams {
	return &VirtualizationInterfacesUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationInterfacesUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization interfaces update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationInterfacesUpdateParams struct {

	// Data.
	Data *models.WritableVMInterface

	/* ID.

	   A UUID string identifying this interface.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization interfaces update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesUpdateParams) WithDefaults() *VirtualizationInterfacesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization interfaces update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) WithContext(ctx context.Context) *VirtualizationInterfacesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) WithData(data *models.WritableVMInterface) *VirtualizationInterfacesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) SetData(data *models.WritableVMInterface) {
	o.Data = data
}

// WithID adds the id to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) WithID(id strfmt.UUID) *VirtualizationInterfacesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization interfaces update params
func (o *VirtualizationInterfacesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
