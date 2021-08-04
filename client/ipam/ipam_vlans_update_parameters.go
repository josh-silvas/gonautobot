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

// NewIpamVlansUpdateParams creates a new IpamVlansUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansUpdateParams() *IpamVlansUpdateParams {
	return &IpamVlansUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansUpdateParamsWithTimeout creates a new IpamVlansUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVlansUpdateParamsWithTimeout(timeout time.Duration) *IpamVlansUpdateParams {
	return &IpamVlansUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVlansUpdateParamsWithContext creates a new IpamVlansUpdateParams object
// with the ability to set a context for a request.
func NewIpamVlansUpdateParamsWithContext(ctx context.Context) *IpamVlansUpdateParams {
	return &IpamVlansUpdateParams{
		Context: ctx,
	}
}

// NewIpamVlansUpdateParamsWithHTTPClient creates a new IpamVlansUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansUpdateParamsWithHTTPClient(client *http.Client) *IpamVlansUpdateParams {
	return &IpamVlansUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVlansUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vlans update operation.

   Typically these are written to a http.Request.
*/
type IpamVlansUpdateParams struct {

	// Data.
	Data *models.WritableVLAN

	/* ID.

	   A UUID string identifying this VLAN.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansUpdateParams) WithDefaults() *IpamVlansUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans update params
func (o *IpamVlansUpdateParams) WithTimeout(timeout time.Duration) *IpamVlansUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans update params
func (o *IpamVlansUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans update params
func (o *IpamVlansUpdateParams) WithContext(ctx context.Context) *IpamVlansUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans update params
func (o *IpamVlansUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans update params
func (o *IpamVlansUpdateParams) WithHTTPClient(client *http.Client) *IpamVlansUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans update params
func (o *IpamVlansUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlans update params
func (o *IpamVlansUpdateParams) WithData(data *models.WritableVLAN) *IpamVlansUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlans update params
func (o *IpamVlansUpdateParams) SetData(data *models.WritableVLAN) {
	o.Data = data
}

// WithID adds the id to the ipam vlans update params
func (o *IpamVlansUpdateParams) WithID(id strfmt.UUID) *IpamVlansUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlans update params
func (o *IpamVlansUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
