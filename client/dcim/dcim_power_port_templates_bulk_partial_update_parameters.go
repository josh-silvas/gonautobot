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

// NewDcimPowerPortTemplatesBulkPartialUpdateParams creates a new DcimPowerPortTemplatesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortTemplatesBulkPartialUpdateParams() *DcimPowerPortTemplatesBulkPartialUpdateParams {
	return &DcimPowerPortTemplatesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesBulkPartialUpdateParamsWithTimeout creates a new DcimPowerPortTemplatesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortTemplatesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesBulkPartialUpdateParams {
	return &DcimPowerPortTemplatesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesBulkPartialUpdateParamsWithContext creates a new DcimPowerPortTemplatesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerPortTemplatesBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesBulkPartialUpdateParams {
	return &DcimPowerPortTemplatesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesBulkPartialUpdateParamsWithHTTPClient creates a new DcimPowerPortTemplatesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortTemplatesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesBulkPartialUpdateParams {
	return &DcimPowerPortTemplatesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortTemplatesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power port templates bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortTemplatesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritablePowerPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) WithDefaults() *DcimPowerPortTemplatesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power port templates bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) WithData(data *models.WritablePowerPortTemplate) *DcimPowerPortTemplatesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power port templates bulk partial update params
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) SetData(data *models.WritablePowerPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
