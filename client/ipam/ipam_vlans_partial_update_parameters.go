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

// NewIpamVlansPartialUpdateParams creates a new IpamVlansPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansPartialUpdateParams() *IpamVlansPartialUpdateParams {
	return &IpamVlansPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansPartialUpdateParamsWithTimeout creates a new IpamVlansPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVlansPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamVlansPartialUpdateParams {
	return &IpamVlansPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVlansPartialUpdateParamsWithContext creates a new IpamVlansPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamVlansPartialUpdateParamsWithContext(ctx context.Context) *IpamVlansPartialUpdateParams {
	return &IpamVlansPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamVlansPartialUpdateParamsWithHTTPClient creates a new IpamVlansPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamVlansPartialUpdateParams {
	return &IpamVlansPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVlansPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vlans partial update operation.

   Typically these are written to a http.Request.
*/
type IpamVlansPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the ipam vlans partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansPartialUpdateParams) WithDefaults() *IpamVlansPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamVlansPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) WithContext(ctx context.Context) *IpamVlansPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamVlansPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) WithData(data *models.WritableVLAN) *IpamVlansPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) SetData(data *models.WritableVLAN) {
	o.Data = data
}

// WithID adds the id to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) WithID(id strfmt.UUID) *IpamVlansPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlans partial update params
func (o *IpamVlansPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
