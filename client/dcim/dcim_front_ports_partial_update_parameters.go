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

// NewDcimFrontPortsPartialUpdateParams creates a new DcimFrontPortsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsPartialUpdateParams() *DcimFrontPortsPartialUpdateParams {
	return &DcimFrontPortsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsPartialUpdateParamsWithTimeout creates a new DcimFrontPortsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortsPartialUpdateParams {
	return &DcimFrontPortsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsPartialUpdateParamsWithContext creates a new DcimFrontPortsPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsPartialUpdateParamsWithContext(ctx context.Context) *DcimFrontPortsPartialUpdateParams {
	return &DcimFrontPortsPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsPartialUpdateParamsWithHTTPClient creates a new DcimFrontPortsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortsPartialUpdateParams {
	return &DcimFrontPortsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim front ports partial update operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortsPartialUpdateParams struct {

	// Data.
	Data *models.WritableFrontPort

	/* ID.

	   A UUID string identifying this front port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsPartialUpdateParams) WithDefaults() *DcimFrontPortsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithContext(ctx context.Context) *DcimFrontPortsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithData(data *models.WritableFrontPort) *DcimFrontPortsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetData(data *models.WritableFrontPort) {
	o.Data = data
}

// WithID adds the id to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) WithID(id strfmt.UUID) *DcimFrontPortsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports partial update params
func (o *DcimFrontPortsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
