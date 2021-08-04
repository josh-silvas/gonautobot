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

// NewVirtualizationClusterGroupsDeleteParams creates a new VirtualizationClusterGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClusterGroupsDeleteParams() *VirtualizationClusterGroupsDeleteParams {
	return &VirtualizationClusterGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsDeleteParamsWithTimeout creates a new VirtualizationClusterGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClusterGroupsDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsDeleteParams {
	return &VirtualizationClusterGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsDeleteParamsWithContext creates a new VirtualizationClusterGroupsDeleteParams object
// with the ability to set a context for a request.
func NewVirtualizationClusterGroupsDeleteParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsDeleteParams {
	return &VirtualizationClusterGroupsDeleteParams{
		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsDeleteParamsWithHTTPClient creates a new VirtualizationClusterGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClusterGroupsDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsDeleteParams {
	return &VirtualizationClusterGroupsDeleteParams{
		HTTPClient: client,
	}
}

/* VirtualizationClusterGroupsDeleteParams contains all the parameters to send to the API endpoint
   for the virtualization cluster groups delete operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClusterGroupsDeleteParams struct {

	/* ID.

	   A UUID string identifying this cluster group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization cluster groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterGroupsDeleteParams) WithDefaults() *VirtualizationClusterGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization cluster groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) WithID(id strfmt.UUID) *VirtualizationClusterGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster groups delete params
func (o *VirtualizationClusterGroupsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
