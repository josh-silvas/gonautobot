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

// NewDcimPowerFeedsUpdateParams creates a new DcimPowerFeedsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsUpdateParams() *DcimPowerFeedsUpdateParams {
	return &DcimPowerFeedsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsUpdateParamsWithTimeout creates a new DcimPowerFeedsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsUpdateParams {
	return &DcimPowerFeedsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsUpdateParamsWithContext creates a new DcimPowerFeedsUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsUpdateParamsWithContext(ctx context.Context) *DcimPowerFeedsUpdateParams {
	return &DcimPowerFeedsUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsUpdateParamsWithHTTPClient creates a new DcimPowerFeedsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsUpdateParams {
	return &DcimPowerFeedsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerFeedsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power feeds update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerFeedsUpdateParams struct {

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

// WithDefaults hydrates default values in the dcim power feeds update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsUpdateParams) WithDefaults() *DcimPowerFeedsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) WithContext(ctx context.Context) *DcimPowerFeedsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) WithData(data *models.WritablePowerFeed) *DcimPowerFeedsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) SetData(data *models.WritablePowerFeed) {
	o.Data = data
}

// WithID adds the id to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) WithID(id strfmt.UUID) *DcimPowerFeedsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power feeds update params
func (o *DcimPowerFeedsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
