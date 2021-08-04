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

// NewIpamAggregatesBulkPartialUpdateParams creates a new IpamAggregatesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesBulkPartialUpdateParams() *IpamAggregatesBulkPartialUpdateParams {
	return &IpamAggregatesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesBulkPartialUpdateParamsWithTimeout creates a new IpamAggregatesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamAggregatesBulkPartialUpdateParams {
	return &IpamAggregatesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesBulkPartialUpdateParamsWithContext creates a new IpamAggregatesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamAggregatesBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamAggregatesBulkPartialUpdateParams {
	return &IpamAggregatesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamAggregatesBulkPartialUpdateParamsWithHTTPClient creates a new IpamAggregatesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamAggregatesBulkPartialUpdateParams {
	return &IpamAggregatesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamAggregatesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam aggregates bulk partial update operation.

   Typically these are written to a http.Request.
*/
type IpamAggregatesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableAggregate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesBulkPartialUpdateParams) WithDefaults() *IpamAggregatesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamAggregatesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamAggregatesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamAggregatesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) WithData(data *models.WritableAggregate) *IpamAggregatesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam aggregates bulk partial update params
func (o *IpamAggregatesBulkPartialUpdateParams) SetData(data *models.WritableAggregate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
