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

// NewVirtualizationClusterGroupsUpdateParams creates a new VirtualizationClusterGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClusterGroupsUpdateParams() *VirtualizationClusterGroupsUpdateParams {
	return &VirtualizationClusterGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterGroupsUpdateParamsWithTimeout creates a new VirtualizationClusterGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClusterGroupsUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClusterGroupsUpdateParams {
	return &VirtualizationClusterGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationClusterGroupsUpdateParamsWithContext creates a new VirtualizationClusterGroupsUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationClusterGroupsUpdateParamsWithContext(ctx context.Context) *VirtualizationClusterGroupsUpdateParams {
	return &VirtualizationClusterGroupsUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationClusterGroupsUpdateParamsWithHTTPClient creates a new VirtualizationClusterGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClusterGroupsUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClusterGroupsUpdateParams {
	return &VirtualizationClusterGroupsUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationClusterGroupsUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization cluster groups update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClusterGroupsUpdateParams struct {

	// Data.
	Data *models.ClusterGroup

	/* ID.

	   A UUID string identifying this cluster group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization cluster groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterGroupsUpdateParams) WithDefaults() *VirtualizationClusterGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization cluster groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClusterGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClusterGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithContext(ctx context.Context) *VirtualizationClusterGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClusterGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithData(data *models.ClusterGroup) *VirtualizationClusterGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetData(data *models.ClusterGroup) {
	o.Data = data
}

// WithID adds the id to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) WithID(id strfmt.UUID) *VirtualizationClusterGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster groups update params
func (o *VirtualizationClusterGroupsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
