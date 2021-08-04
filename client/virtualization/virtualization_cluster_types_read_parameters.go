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

// NewVirtualizationClusterTypesReadParams creates a new VirtualizationClusterTypesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClusterTypesReadParams() *VirtualizationClusterTypesReadParams {
	return &VirtualizationClusterTypesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterTypesReadParamsWithTimeout creates a new VirtualizationClusterTypesReadParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClusterTypesReadParamsWithTimeout(timeout time.Duration) *VirtualizationClusterTypesReadParams {
	return &VirtualizationClusterTypesReadParams{
		timeout: timeout,
	}
}

// NewVirtualizationClusterTypesReadParamsWithContext creates a new VirtualizationClusterTypesReadParams object
// with the ability to set a context for a request.
func NewVirtualizationClusterTypesReadParamsWithContext(ctx context.Context) *VirtualizationClusterTypesReadParams {
	return &VirtualizationClusterTypesReadParams{
		Context: ctx,
	}
}

// NewVirtualizationClusterTypesReadParamsWithHTTPClient creates a new VirtualizationClusterTypesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClusterTypesReadParamsWithHTTPClient(client *http.Client) *VirtualizationClusterTypesReadParams {
	return &VirtualizationClusterTypesReadParams{
		HTTPClient: client,
	}
}

/* VirtualizationClusterTypesReadParams contains all the parameters to send to the API endpoint
   for the virtualization cluster types read operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClusterTypesReadParams struct {

	/* ID.

	   A UUID string identifying this cluster type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization cluster types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterTypesReadParams) WithDefaults() *VirtualizationClusterTypesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization cluster types read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterTypesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithTimeout(timeout time.Duration) *VirtualizationClusterTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithContext(ctx context.Context) *VirtualizationClusterTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithHTTPClient(client *http.Client) *VirtualizationClusterTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithID(id strfmt.UUID) *VirtualizationClusterTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
