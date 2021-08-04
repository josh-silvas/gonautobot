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

// NewVirtualizationClustersDeleteParams creates a new VirtualizationClustersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClustersDeleteParams() *VirtualizationClustersDeleteParams {
	return &VirtualizationClustersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersDeleteParamsWithTimeout creates a new VirtualizationClustersDeleteParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClustersDeleteParamsWithTimeout(timeout time.Duration) *VirtualizationClustersDeleteParams {
	return &VirtualizationClustersDeleteParams{
		timeout: timeout,
	}
}

// NewVirtualizationClustersDeleteParamsWithContext creates a new VirtualizationClustersDeleteParams object
// with the ability to set a context for a request.
func NewVirtualizationClustersDeleteParamsWithContext(ctx context.Context) *VirtualizationClustersDeleteParams {
	return &VirtualizationClustersDeleteParams{
		Context: ctx,
	}
}

// NewVirtualizationClustersDeleteParamsWithHTTPClient creates a new VirtualizationClustersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClustersDeleteParamsWithHTTPClient(client *http.Client) *VirtualizationClustersDeleteParams {
	return &VirtualizationClustersDeleteParams{
		HTTPClient: client,
	}
}

/* VirtualizationClustersDeleteParams contains all the parameters to send to the API endpoint
   for the virtualization clusters delete operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClustersDeleteParams struct {

	/* ID.

	   A UUID string identifying this cluster.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization clusters delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersDeleteParams) WithDefaults() *VirtualizationClustersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization clusters delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithTimeout(timeout time.Duration) *VirtualizationClustersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithContext(ctx context.Context) *VirtualizationClustersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithHTTPClient(client *http.Client) *VirtualizationClustersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) WithID(id strfmt.UUID) *VirtualizationClustersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization clusters delete params
func (o *VirtualizationClustersDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
