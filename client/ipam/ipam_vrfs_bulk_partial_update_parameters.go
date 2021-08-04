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

// NewIpamVrfsBulkPartialUpdateParams creates a new IpamVrfsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsBulkPartialUpdateParams() *IpamVrfsBulkPartialUpdateParams {
	return &IpamVrfsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsBulkPartialUpdateParamsWithTimeout creates a new IpamVrfsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamVrfsBulkPartialUpdateParams {
	return &IpamVrfsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamVrfsBulkPartialUpdateParamsWithContext creates a new IpamVrfsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamVrfsBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamVrfsBulkPartialUpdateParams {
	return &IpamVrfsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamVrfsBulkPartialUpdateParamsWithHTTPClient creates a new IpamVrfsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamVrfsBulkPartialUpdateParams {
	return &IpamVrfsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamVrfsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam vrfs bulk partial update operation.

   Typically these are written to a http.Request.
*/
type IpamVrfsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableVRF

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vrfs bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsBulkPartialUpdateParams) WithDefaults() *IpamVrfsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamVrfsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamVrfsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamVrfsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) WithData(data *models.WritableVRF) *IpamVrfsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vrfs bulk partial update params
func (o *IpamVrfsBulkPartialUpdateParams) SetData(data *models.WritableVRF) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
