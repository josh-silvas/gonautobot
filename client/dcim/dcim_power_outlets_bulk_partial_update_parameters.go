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

// NewDcimPowerOutletsBulkPartialUpdateParams creates a new DcimPowerOutletsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletsBulkPartialUpdateParams() *DcimPowerOutletsBulkPartialUpdateParams {
	return &DcimPowerOutletsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsBulkPartialUpdateParamsWithTimeout creates a new DcimPowerOutletsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsBulkPartialUpdateParams {
	return &DcimPowerOutletsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletsBulkPartialUpdateParamsWithContext creates a new DcimPowerOutletsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimPowerOutletsBulkPartialUpdateParams {
	return &DcimPowerOutletsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletsBulkPartialUpdateParamsWithHTTPClient creates a new DcimPowerOutletsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsBulkPartialUpdateParams {
	return &DcimPowerOutletsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power outlets bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritablePowerOutlet

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlets bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsBulkPartialUpdateParams) WithDefaults() *DcimPowerOutletsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlets bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimPowerOutletsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) WithData(data *models.WritablePowerOutlet) *DcimPowerOutletsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlets bulk partial update params
func (o *DcimPowerOutletsBulkPartialUpdateParams) SetData(data *models.WritablePowerOutlet) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
