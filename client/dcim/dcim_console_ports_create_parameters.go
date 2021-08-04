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

// NewDcimConsolePortsCreateParams creates a new DcimConsolePortsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsCreateParams() *DcimConsolePortsCreateParams {
	return &DcimConsolePortsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsCreateParamsWithTimeout creates a new DcimConsolePortsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsCreateParamsWithTimeout(timeout time.Duration) *DcimConsolePortsCreateParams {
	return &DcimConsolePortsCreateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsCreateParamsWithContext creates a new DcimConsolePortsCreateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsCreateParamsWithContext(ctx context.Context) *DcimConsolePortsCreateParams {
	return &DcimConsolePortsCreateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsCreateParamsWithHTTPClient creates a new DcimConsolePortsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsCreateParamsWithHTTPClient(client *http.Client) *DcimConsolePortsCreateParams {
	return &DcimConsolePortsCreateParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsCreateParams contains all the parameters to send to the API endpoint
   for the dcim console ports create operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsCreateParams struct {

	// Data.
	Data *models.WritableConsolePort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsCreateParams) WithDefaults() *DcimConsolePortsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithTimeout(timeout time.Duration) *DcimConsolePortsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithContext(ctx context.Context) *DcimConsolePortsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithHTTPClient(client *http.Client) *DcimConsolePortsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) WithData(data *models.WritableConsolePort) *DcimConsolePortsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console ports create params
func (o *DcimConsolePortsCreateParams) SetData(data *models.WritableConsolePort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
