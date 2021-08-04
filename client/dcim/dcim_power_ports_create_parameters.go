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

// NewDcimPowerPortsCreateParams creates a new DcimPowerPortsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortsCreateParams() *DcimPowerPortsCreateParams {
	return &DcimPowerPortsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsCreateParamsWithTimeout creates a new DcimPowerPortsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortsCreateParamsWithTimeout(timeout time.Duration) *DcimPowerPortsCreateParams {
	return &DcimPowerPortsCreateParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortsCreateParamsWithContext creates a new DcimPowerPortsCreateParams object
// with the ability to set a context for a request.
func NewDcimPowerPortsCreateParamsWithContext(ctx context.Context) *DcimPowerPortsCreateParams {
	return &DcimPowerPortsCreateParams{
		Context: ctx,
	}
}

// NewDcimPowerPortsCreateParamsWithHTTPClient creates a new DcimPowerPortsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortsCreateParamsWithHTTPClient(client *http.Client) *DcimPowerPortsCreateParams {
	return &DcimPowerPortsCreateParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortsCreateParams contains all the parameters to send to the API endpoint
   for the dcim power ports create operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortsCreateParams struct {

	// Data.
	Data *models.WritablePowerPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsCreateParams) WithDefaults() *DcimPowerPortsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) WithTimeout(timeout time.Duration) *DcimPowerPortsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) WithContext(ctx context.Context) *DcimPowerPortsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) WithHTTPClient(client *http.Client) *DcimPowerPortsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) WithData(data *models.WritablePowerPort) *DcimPowerPortsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power ports create params
func (o *DcimPowerPortsCreateParams) SetData(data *models.WritablePowerPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
