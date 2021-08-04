package dcim

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

// NewDcimPowerFeedsPartialUpdateParams creates a new DcimPowerFeedsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsPartialUpdateParams() *DcimPowerFeedsPartialUpdateParams {
	return &DcimPowerFeedsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsPartialUpdateParamsWithTimeout creates a new DcimPowerFeedsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsPartialUpdateParams {
	return &DcimPowerFeedsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsPartialUpdateParamsWithContext creates a new DcimPowerFeedsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsPartialUpdateParamsWithContext(ctx context.Context) *DcimPowerFeedsPartialUpdateParams {
	return &DcimPowerFeedsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsPartialUpdateParamsWithHTTPClient creates a new DcimPowerFeedsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsPartialUpdateParams {
	return &DcimPowerFeedsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerFeedsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power feeds partial update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerFeedsPartialUpdateParams struct {

	// Data.
	Data *models.WritablePowerFeed

	/* ID.

	   A UUID string identifying this power feed.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power feeds partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsPartialUpdateParams) WithDefaults() *DcimPowerFeedsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) WithContext(ctx context.Context) *DcimPowerFeedsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) WithData(data *models.WritablePowerFeed) *DcimPowerFeedsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) SetData(data *models.WritablePowerFeed) {
	o.Data = data
}

// WithID adds the id to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) WithID(id strfmt.UUID) *DcimPowerFeedsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power feeds partial update params
func (o *DcimPowerFeedsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
