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

// NewIpamVlansBulkPartialUpdateParams creates a new IpamVlansBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansBulkPartialUpdateParams() *IpamVlansBulkPartialUpdateParams {
	return &IpamVlansBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansBulkPartialUpdateParamsWithTimeout creates a new IpamVlansBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVlansBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamVlansBulkPartialUpdateParams {
	return &IpamVlansBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVlansBulkPartialUpdateParamsWithContext creates a new IpamVlansBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamVlansBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamVlansBulkPartialUpdateParams {
	return &IpamVlansBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamVlansBulkPartialUpdateParamsWithHTTPClient creates a new IpamVlansBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamVlansBulkPartialUpdateParams {
	return &IpamVlansBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVlansBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vlans bulk partial update operation.

   Typically these are written to a http.Request.
*/
type IpamVlansBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableVLAN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansBulkPartialUpdateParams) WithDefaults() *IpamVlansBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamVlansBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamVlansBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamVlansBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) WithData(data *models.WritableVLAN) *IpamVlansBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlans bulk partial update params
func (o *IpamVlansBulkPartialUpdateParams) SetData(data *models.WritableVLAN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
