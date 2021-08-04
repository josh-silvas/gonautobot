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

// NewDcimRacksPartialUpdateParams creates a new DcimRacksPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksPartialUpdateParams() *DcimRacksPartialUpdateParams {
	return &DcimRacksPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksPartialUpdateParamsWithTimeout creates a new DcimRacksPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRacksPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRacksPartialUpdateParams {
	return &DcimRacksPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRacksPartialUpdateParamsWithContext creates a new DcimRacksPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRacksPartialUpdateParamsWithContext(ctx context.Context) *DcimRacksPartialUpdateParams {
	return &DcimRacksPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRacksPartialUpdateParamsWithHTTPClient creates a new DcimRacksPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRacksPartialUpdateParams {
	return &DcimRacksPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRacksPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim racks partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRacksPartialUpdateParams struct {

	// Data.
	Data *models.WritableRack

	/* ID.

	   A UUID string identifying this rack.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksPartialUpdateParams) WithDefaults() *DcimRacksPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRacksPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithContext(ctx context.Context) *DcimRacksPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRacksPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithData(data *models.WritableRack) *DcimRacksPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WithID adds the id to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) WithID(id strfmt.UUID) *DcimRacksPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks partial update params
func (o *DcimRacksPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
