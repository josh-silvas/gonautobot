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

// NewDcimConsoleServerPortsCreateParams creates a new DcimConsoleServerPortsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortsCreateParams() *DcimConsoleServerPortsCreateParams {
	return &DcimConsoleServerPortsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsCreateParamsWithTimeout creates a new DcimConsoleServerPortsCreateParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortsCreateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsCreateParams {
	return &DcimConsoleServerPortsCreateParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsCreateParamsWithContext creates a new DcimConsoleServerPortsCreateParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortsCreateParamsWithContext(ctx context.Context) *DcimConsoleServerPortsCreateParams {
	return &DcimConsoleServerPortsCreateParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortsCreateParamsWithHTTPClient creates a new DcimConsoleServerPortsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortsCreateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsCreateParams {
	return &DcimConsoleServerPortsCreateParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortsCreateParams contains all the parameters to send to the API endpoint
   for the dcim console server ports create operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortsCreateParams struct {

	// Data.
	Data *models.WritableConsoleServerPort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsCreateParams) WithDefaults() *DcimConsoleServerPortsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server ports create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) WithContext(ctx context.Context) *DcimConsoleServerPortsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) WithData(data *models.WritableConsoleServerPort) *DcimConsoleServerPortsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server ports create params
func (o *DcimConsoleServerPortsCreateParams) SetData(data *models.WritableConsoleServerPort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
