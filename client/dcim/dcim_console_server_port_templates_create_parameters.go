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

// NewDcimConsoleServerPortTemplatesCreateParams creates a new DcimConsoleServerPortTemplatesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortTemplatesCreateParams() *DcimConsoleServerPortTemplatesCreateParams {
	return &DcimConsoleServerPortTemplatesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortTemplatesCreateParamsWithTimeout creates a new DcimConsoleServerPortTemplatesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesCreateParams {
	return &DcimConsoleServerPortTemplatesCreateParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortTemplatesCreateParamsWithContext creates a new DcimConsoleServerPortTemplatesCreateParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortTemplatesCreateParamsWithContext(ctx context.Context) *DcimConsoleServerPortTemplatesCreateParams {
	return &DcimConsoleServerPortTemplatesCreateParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortTemplatesCreateParamsWithHTTPClient creates a new DcimConsoleServerPortTemplatesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesCreateParams {
	return &DcimConsoleServerPortTemplatesCreateParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortTemplatesCreateParams contains all the parameters to send to the API endpoint
   for the dcim console server port templates create operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortTemplatesCreateParams struct {

	// Data.
	Data *models.WritableConsoleServerPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server port templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesCreateParams) WithDefaults() *DcimConsoleServerPortTemplatesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server port templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortTemplatesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) WithContext(ctx context.Context) *DcimConsoleServerPortTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) WithData(data *models.WritableConsoleServerPortTemplate) *DcimConsoleServerPortTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server port templates create params
func (o *DcimConsoleServerPortTemplatesCreateParams) SetData(data *models.WritableConsoleServerPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
