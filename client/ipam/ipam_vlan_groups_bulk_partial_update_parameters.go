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

// NewIpamVlanGroupsBulkPartialUpdateParams creates a new IpamVlanGroupsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsBulkPartialUpdateParams() *IpamVlanGroupsBulkPartialUpdateParams {
	return &IpamVlanGroupsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsBulkPartialUpdateParamsWithTimeout creates a new IpamVlanGroupsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsBulkPartialUpdateParams {
	return &IpamVlanGroupsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsBulkPartialUpdateParamsWithContext creates a new IpamVlanGroupsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamVlanGroupsBulkPartialUpdateParams {
	return &IpamVlanGroupsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsBulkPartialUpdateParamsWithHTTPClient creates a new IpamVlanGroupsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsBulkPartialUpdateParams {
	return &IpamVlanGroupsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVlanGroupsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vlan groups bulk partial update operation.

   Typically these are written to a http.Request.
*/
type IpamVlanGroupsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableVLANGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsBulkPartialUpdateParams) WithDefaults() *IpamVlanGroupsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamVlanGroupsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) WithData(data *models.WritableVLANGroup) *IpamVlanGroupsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlan groups bulk partial update params
func (o *IpamVlanGroupsBulkPartialUpdateParams) SetData(data *models.WritableVLANGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
