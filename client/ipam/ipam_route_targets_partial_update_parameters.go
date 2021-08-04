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

// NewIpamRouteTargetsPartialUpdateParams creates a new IpamRouteTargetsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRouteTargetsPartialUpdateParams() *IpamRouteTargetsPartialUpdateParams {
	return &IpamRouteTargetsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRouteTargetsPartialUpdateParamsWithTimeout creates a new IpamRouteTargetsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRouteTargetsPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamRouteTargetsPartialUpdateParams {
	return &IpamRouteTargetsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRouteTargetsPartialUpdateParamsWithContext creates a new IpamRouteTargetsPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamRouteTargetsPartialUpdateParamsWithContext(ctx context.Context) *IpamRouteTargetsPartialUpdateParams {
	return &IpamRouteTargetsPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamRouteTargetsPartialUpdateParamsWithHTTPClient creates a new IpamRouteTargetsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRouteTargetsPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamRouteTargetsPartialUpdateParams {
	return &IpamRouteTargetsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamRouteTargetsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam route targets partial update operation.

   Typically these are written to a http.Request.
*/
type IpamRouteTargetsPartialUpdateParams struct {

	// Data.
	Data *models.WritableRouteTarget

	/* ID.

	   A UUID string identifying this route target.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam route targets partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsPartialUpdateParams) WithDefaults() *IpamRouteTargetsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam route targets partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamRouteTargetsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) WithContext(ctx context.Context) *IpamRouteTargetsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamRouteTargetsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) WithData(data *models.WritableRouteTarget) *IpamRouteTargetsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) SetData(data *models.WritableRouteTarget) {
	o.Data = data
}

// WithID adds the id to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) WithID(id strfmt.UUID) *IpamRouteTargetsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam route targets partial update params
func (o *IpamRouteTargetsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRouteTargetsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
