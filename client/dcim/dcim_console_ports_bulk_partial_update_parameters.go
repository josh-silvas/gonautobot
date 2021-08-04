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

// NewDcimConsolePortsBulkPartialUpdateParams creates a new DcimConsolePortsBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsBulkPartialUpdateParams() *DcimConsolePortsBulkPartialUpdateParams {
	return &DcimConsolePortsBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsBulkPartialUpdateParamsWithTimeout creates a new DcimConsolePortsBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortsBulkPartialUpdateParams {
	return &DcimConsolePortsBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsBulkPartialUpdateParamsWithContext creates a new DcimConsolePortsBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimConsolePortsBulkPartialUpdateParams {
	return &DcimConsolePortsBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsBulkPartialUpdateParamsWithHTTPClient creates a new DcimConsolePortsBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortsBulkPartialUpdateParams {
	return &DcimConsolePortsBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console ports bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableConsolePort

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsBulkPartialUpdateParams) WithDefaults() *DcimConsolePortsBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortsBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimConsolePortsBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortsBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) WithData(data *models.WritableConsolePort) *DcimConsolePortsBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console ports bulk partial update params
func (o *DcimConsolePortsBulkPartialUpdateParams) SetData(data *models.WritableConsolePort) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
