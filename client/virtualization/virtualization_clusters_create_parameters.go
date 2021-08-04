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

// NewVirtualizationClustersCreateParams creates a new VirtualizationClustersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVirtualizationClustersCreateParams() *VirtualizationClustersCreateParams {
	return &VirtualizationClustersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClustersCreateParamsWithTimeout creates a new VirtualizationClustersCreateParams object
// with the ability to set a timeout on a request.
func NewVirtualizationClustersCreateParamsWithTimeout(timeout time.Duration) *VirtualizationClustersCreateParams {
	return &VirtualizationClustersCreateParams{
		timeout: timeout,
	}
}

// NewVirtualizationClustersCreateParamsWithContext creates a new VirtualizationClustersCreateParams object
// with the ability to set a context for a request.
func NewVirtualizationClustersCreateParamsWithContext(ctx context.Context) *VirtualizationClustersCreateParams {
	return &VirtualizationClustersCreateParams{
		Context: ctx,
	}
}

// NewVirtualizationClustersCreateParamsWithHTTPClient creates a new VirtualizationClustersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewVirtualizationClustersCreateParamsWithHTTPClient(client *http.Client) *VirtualizationClustersCreateParams {
	return &VirtualizationClustersCreateParams{
		HTTPClient: client,
	}
}

/* VirtualizationClustersCreateParams contains all the parameters to send to the API endpoint
   for the virtualization clusters create operation.

   Typically these are written to a http.Request.
*/
type VirtualizationClustersCreateParams struct {

	// Data.
	Data *models.WritableCluster

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the virtualization clusters create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersCreateParams) WithDefaults() *VirtualizationClustersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the virtualization clusters create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VirtualizationClustersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithTimeout(timeout time.Duration) *VirtualizationClustersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithContext(ctx context.Context) *VirtualizationClustersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithHTTPClient(client *http.Client) *VirtualizationClustersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) WithData(data *models.WritableCluster) *VirtualizationClustersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the virtualization clusters create params
func (o *VirtualizationClustersCreateParams) SetData(data *models.WritableCluster) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClustersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
