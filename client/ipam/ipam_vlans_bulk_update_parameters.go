package ipam

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

// NewIpamVlansBulkUpdateParams creates a new IpamVlansBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansBulkUpdateParams() *IpamVlansBulkUpdateParams {
	return &IpamVlansBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansBulkUpdateParamsWithTimeout creates a new IpamVlansBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVlansBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamVlansBulkUpdateParams {
	return &IpamVlansBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVlansBulkUpdateParamsWithContext creates a new IpamVlansBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamVlansBulkUpdateParamsWithContext(ctx context.Context) *IpamVlansBulkUpdateParams {
	return &IpamVlansBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamVlansBulkUpdateParamsWithHTTPClient creates a new IpamVlansBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamVlansBulkUpdateParams {
	return &IpamVlansBulkUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVlansBulkUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vlans bulk update operation.

   Typically these are written to a http.Request.
*/
type IpamVlansBulkUpdateParams struct {

	// Data.
	Data *models.WritableVLAN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansBulkUpdateParams) WithDefaults() *IpamVlansBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamVlansBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) WithContext(ctx context.Context) *IpamVlansBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamVlansBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) WithData(data *models.WritableVLAN) *IpamVlansBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlans bulk update params
func (o *IpamVlansBulkUpdateParams) SetData(data *models.WritableVLAN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
