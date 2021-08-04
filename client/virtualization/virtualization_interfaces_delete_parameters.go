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

// NewVirtualizationInterfacesDeleteParams creates a new VirtualizationInterfacesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationInterfacesDeleteParams() *VirtualizationInterfacesDeleteParams {
	return &VirtualizationInterfacesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationInterfacesDeleteParamsWithTimeout creates a new VirtualizationInterfacesDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualizationInterfacesDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationInterfacesDeleteParams {
	return &VirtualizationInterfacesDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualizationInterfacesDeleteParamsWithContext creates a new VirtualizationInterfacesDeleteParams object
// with the ability to set a context for a request.
func NewVirtualizationInterfacesDeleteParamsWithContext(ctx context.Context) *VirtualizationInterfacesDeleteParams {
	return &VirtualizationInterfacesDeleteParams{
		Context: ctx,
	}
}

// NewVirtualizationInterfacesDeleteParamsWithHTTPClient creates a new VirtualizationInterfacesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationInterfacesDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationInterfacesDeleteParams {
	return &VirtualizationInterfacesDeleteParams{
		HTTPClient: client,
	}
}

/* VirtualizationInterfacesDeleteParams contains all the parameters to send to the API endpoint
   for the virtualization interfaces delete operation.

   Typically these are written to a http.Request.
*/
type VirtualizationInterfacesDeleteParams struct {

	/* ID.

	   A UUID string identifying this interface.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization interfaces delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesDeleteParams) WithDefaults() *VirtualizationInterfacesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization interfaces delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationInterfacesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationInterfacesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithContext(ctx context.Context) *VirtualizationInterfacesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationInterfacesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) WithID(id strfmt.UUID) *VirtualizationInterfacesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization interfaces delete params
func (o *VirtualizationInterfacesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationInterfacesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
