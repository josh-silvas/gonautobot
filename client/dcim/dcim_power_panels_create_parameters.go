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

// NewDcimPowerPanelsCreateParams creates a new DcimPowerPanelsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPanelsCreateParams() *DcimPowerPanelsCreateParams {
	return &DcimPowerPanelsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsCreateParamsWithTimeout creates a new DcimPowerPanelsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPanelsCreateParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsCreateParams {
	return &DcimPowerPanelsCreateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPanelsCreateParamsWithContext creates a new DcimPowerPanelsCreateParams object
// with the ability to set a context for a request.
func NewDcimPowerPanelsCreateParamsWithContext(ctx context.Context) *DcimPowerPanelsCreateParams {
	return &DcimPowerPanelsCreateParams{
		Context: ctx,
	}
}

// NewDcimPowerPanelsCreateParamsWithHTTPClient creates a new DcimPowerPanelsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPanelsCreateParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsCreateParams {
	return &DcimPowerPanelsCreateParams{
		HTTPClient: client,
	}
}

/* DcimPowerPanelsCreateParams contains all the parameters to send to the API endpoint
   for the dcim power panels create operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPanelsCreateParams struct {

	// Data.
	Data *models.WritablePowerPanel

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power panels create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsCreateParams) WithDefaults() *DcimPowerPanelsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power panels create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) WithContext(ctx context.Context) *DcimPowerPanelsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) WithData(data *models.WritablePowerPanel) *DcimPowerPanelsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power panels create params
func (o *DcimPowerPanelsCreateParams) SetData(data *models.WritablePowerPanel) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
