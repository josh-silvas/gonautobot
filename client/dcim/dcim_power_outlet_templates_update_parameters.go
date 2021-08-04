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

// NewDcimPowerOutletTemplatesUpdateParams creates a new DcimPowerOutletTemplatesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletTemplatesUpdateParams() *DcimPowerOutletTemplatesUpdateParams {
	return &DcimPowerOutletTemplatesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesUpdateParamsWithTimeout creates a new DcimPowerOutletTemplatesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletTemplatesUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesUpdateParams {
	return &DcimPowerOutletTemplatesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesUpdateParamsWithContext creates a new DcimPowerOutletTemplatesUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletTemplatesUpdateParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesUpdateParams {
	return &DcimPowerOutletTemplatesUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesUpdateParamsWithHTTPClient creates a new DcimPowerOutletTemplatesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletTemplatesUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesUpdateParams {
	return &DcimPowerOutletTemplatesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletTemplatesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power outlet templates update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletTemplatesUpdateParams struct {

	// Data.
	Data *models.WritablePowerOutletTemplate

	/* ID.

	   A UUID string identifying this power outlet template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlet templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesUpdateParams) WithDefaults() *DcimPowerOutletTemplatesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlet templates update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) WithData(data *models.WritablePowerOutletTemplate) *DcimPowerOutletTemplatesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) SetData(data *models.WritablePowerOutletTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) WithID(id strfmt.UUID) *DcimPowerOutletTemplatesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlet templates update params
func (o *DcimPowerOutletTemplatesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
