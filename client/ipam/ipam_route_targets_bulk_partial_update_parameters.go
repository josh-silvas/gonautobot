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

// NewIpamRouteTargetsBulkPartialUpdateParams creates a new IpamRouteTargetsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRouteTargetsBulkPartialUpdateParams() *IpamRouteTargetsBulkPartialUpdateParams {
	return &IpamRouteTargetsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRouteTargetsBulkPartialUpdateParamsWithTimeout creates a new IpamRouteTargetsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamRouteTargetsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamRouteTargetsBulkPartialUpdateParams {
	return &IpamRouteTargetsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamRouteTargetsBulkPartialUpdateParamsWithContext creates a new IpamRouteTargetsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamRouteTargetsBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamRouteTargetsBulkPartialUpdateParams {
	return &IpamRouteTargetsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamRouteTargetsBulkPartialUpdateParamsWithHTTPClient creates a new IpamRouteTargetsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRouteTargetsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamRouteTargetsBulkPartialUpdateParams {
	return &IpamRouteTargetsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamRouteTargetsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam route targets bulk partial update operation.

   Typically these are written to a http.Request.
*/
type IpamRouteTargetsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableRouteTarget

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam route targets bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsBulkPartialUpdateParams) WithDefaults() *IpamRouteTargetsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam route targets bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamRouteTargetsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamRouteTargetsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamRouteTargetsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) WithData(data *models.WritableRouteTarget) *IpamRouteTargetsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam route targets bulk partial update params
func (o *IpamRouteTargetsBulkPartialUpdateParams) SetData(data *models.WritableRouteTarget) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRouteTargetsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
