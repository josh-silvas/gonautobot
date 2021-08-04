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

// NewDcimCablesPartialUpdateParams creates a new DcimCablesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesPartialUpdateParams() *DcimCablesPartialUpdateParams {
	return &DcimCablesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesPartialUpdateParamsWithTimeout creates a new DcimCablesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimCablesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimCablesPartialUpdateParams {
	return &DcimCablesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimCablesPartialUpdateParamsWithContext creates a new DcimCablesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimCablesPartialUpdateParamsWithContext(ctx context.Context) *DcimCablesPartialUpdateParams {
	return &DcimCablesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimCablesPartialUpdateParamsWithHTTPClient creates a new DcimCablesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimCablesPartialUpdateParams {
	return &DcimCablesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimCablesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim cables partial update operation.

   Typically these are written to a http.Request.
*/
type DcimCablesPartialUpdateParams struct {

	// Data.
	Data *models.WritableCable

	/* ID.

	   A UUID string identifying this cable.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesPartialUpdateParams) WithDefaults() *DcimCablesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimCablesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) WithContext(ctx context.Context) *DcimCablesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimCablesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) WithData(data *models.WritableCable) *DcimCablesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) SetData(data *models.WritableCable) {
	o.Data = data
}

// WithID adds the id to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) WithID(id strfmt.UUID) *DcimCablesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cables partial update params
func (o *DcimCablesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
