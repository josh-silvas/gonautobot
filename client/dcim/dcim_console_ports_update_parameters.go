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

// NewDcimConsolePortsUpdateParams creates a new DcimConsolePortsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsUpdateParams() *DcimConsolePortsUpdateParams {
	return &DcimConsolePortsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsUpdateParamsWithTimeout creates a new DcimConsolePortsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortsUpdateParams {
	return &DcimConsolePortsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsUpdateParamsWithContext creates a new DcimConsolePortsUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsUpdateParamsWithContext(ctx context.Context) *DcimConsolePortsUpdateParams {
	return &DcimConsolePortsUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsUpdateParamsWithHTTPClient creates a new DcimConsolePortsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortsUpdateParams {
	return &DcimConsolePortsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console ports update operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsUpdateParams struct {

	// Data.
	Data *models.WritableConsolePort

	/* ID.

	   A UUID string identifying this console port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsUpdateParams) WithDefaults() *DcimConsolePortsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithContext(ctx context.Context) *DcimConsolePortsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithData(data *models.WritableConsolePort) *DcimConsolePortsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetData(data *models.WritableConsolePort) {
	o.Data = data
}

// WithID adds the id to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) WithID(id strfmt.UUID) *DcimConsolePortsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports update params
func (o *DcimConsolePortsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
