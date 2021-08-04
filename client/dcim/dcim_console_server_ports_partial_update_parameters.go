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

// NewDcimConsoleServerPortsPartialUpdateParams creates a new DcimConsoleServerPortsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleServerPortsPartialUpdateParams() *DcimConsoleServerPortsPartialUpdateParams {
	return &DcimConsoleServerPortsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortsPartialUpdateParamsWithTimeout creates a new DcimConsoleServerPortsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleServerPortsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortsPartialUpdateParams {
	return &DcimConsoleServerPortsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsoleServerPortsPartialUpdateParamsWithContext creates a new DcimConsoleServerPortsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsoleServerPortsPartialUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortsPartialUpdateParams {
	return &DcimConsoleServerPortsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsoleServerPortsPartialUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleServerPortsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortsPartialUpdateParams {
	return &DcimConsoleServerPortsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsoleServerPortsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console server ports partial update operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleServerPortsPartialUpdateParams struct {

	// Data.
	Data *models.WritableConsoleServerPort

	/* ID.

	   A UUID string identifying this console server port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console server ports partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsPartialUpdateParams) WithDefaults() *DcimConsoleServerPortsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console server ports partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleServerPortsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) WithData(data *models.WritableConsoleServerPort) *DcimConsoleServerPortsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) SetData(data *models.WritableConsoleServerPort) {
	o.Data = data
}

// WithID adds the id to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) WithID(id strfmt.UUID) *DcimConsoleServerPortsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server ports partial update params
func (o *DcimConsoleServerPortsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
