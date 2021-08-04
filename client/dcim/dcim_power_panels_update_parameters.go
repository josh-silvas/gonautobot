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

// NewDcimPowerPanelsUpdateParams creates a new DcimPowerPanelsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPanelsUpdateParams() *DcimPowerPanelsUpdateParams {
	return &DcimPowerPanelsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsUpdateParamsWithTimeout creates a new DcimPowerPanelsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPanelsUpdateParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsUpdateParams {
	return &DcimPowerPanelsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPanelsUpdateParamsWithContext creates a new DcimPowerPanelsUpdateParams object
// with the ability to set a context for a request.
func NewDcimPowerPanelsUpdateParamsWithContext(ctx context.Context) *DcimPowerPanelsUpdateParams {
	return &DcimPowerPanelsUpdateParams{
		Context: ctx,
	}
}

// NewDcimPowerPanelsUpdateParamsWithHTTPClient creates a new DcimPowerPanelsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPanelsUpdateParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsUpdateParams {
	return &DcimPowerPanelsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPowerPanelsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim power panels update operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPanelsUpdateParams struct {

	// Data.
	Data *models.WritablePowerPanel

	/* ID.

	   A UUID string identifying this power panel.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power panels update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsUpdateParams) WithDefaults() *DcimPowerPanelsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power panels update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithContext(ctx context.Context) *DcimPowerPanelsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithData(data *models.WritablePowerPanel) *DcimPowerPanelsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetData(data *models.WritablePowerPanel) {
	o.Data = data
}

// WithID adds the id to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) WithID(id strfmt.UUID) *DcimPowerPanelsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power panels update params
func (o *DcimPowerPanelsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
