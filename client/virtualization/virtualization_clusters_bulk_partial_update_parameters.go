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

// NewVirtualizationClustersBulkPartialUpdateParams creates a new VirtualizationClustersBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClustersBulkPartialUpdateParams() *VirtualizationClustersBulkPartialUpdateParams {
	return &VirtualizationClustersBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersBulkPartialUpdateParamsWithTimeout creates a new VirtualizationClustersBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClustersBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *VirtualizationClustersBulkPartialUpdateParams {
	return &VirtualizationClustersBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewVirtualizationClustersBulkPartialUpdateParamsWithContext creates a new VirtualizationClustersBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewVirtualizationClustersBulkPartialUpdateParamsWithContext(ctx context.Context) *VirtualizationClustersBulkPartialUpdateParams {
	return &VirtualizationClustersBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewVirtualizationClustersBulkPartialUpdateParamsWithHTTPClient creates a new VirtualizationClustersBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClustersBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *VirtualizationClustersBulkPartialUpdateParams {
	return &VirtualizationClustersBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* VirtualizationClustersBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the virtualization clusters bulk partial update operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClustersBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableCluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization clusters bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersBulkPartialUpdateParams) WithDefaults() *VirtualizationClustersBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization clusters bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *VirtualizationClustersBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) WithContext(ctx context.Context) *VirtualizationClustersBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *VirtualizationClustersBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) WithData(data *models.WritableCluster) *VirtualizationClustersBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization clusters bulk partial update params
func (o *VirtualizationClustersBulkPartialUpdateParams) SetData(data *models.WritableCluster) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
