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

// NewDcimManufacturersPartialUpdateParams creates a new DcimManufacturersPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimManufacturersPartialUpdateParams() *DcimManufacturersPartialUpdateParams {
	return &DcimManufacturersPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersPartialUpdateParamsWithTimeout creates a new DcimManufacturersPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimManufacturersPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimManufacturersPartialUpdateParams {
	return &DcimManufacturersPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimManufacturersPartialUpdateParamsWithContext creates a new DcimManufacturersPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimManufacturersPartialUpdateParamsWithContext(ctx context.Context) *DcimManufacturersPartialUpdateParams {
	return &DcimManufacturersPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimManufacturersPartialUpdateParamsWithHTTPClient creates a new DcimManufacturersPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimManufacturersPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimManufacturersPartialUpdateParams {
	return &DcimManufacturersPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimManufacturersPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim manufacturers partial update operation.

   Typically these are written to a http.Request.
*/
type DcimManufacturersPartialUpdateParams struct {

	// Data.
	Data *models.Manufacturer

	/* ID.

	   A UUID string identifying this manufacturer.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim manufacturers partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersPartialUpdateParams) WithDefaults() *DcimManufacturersPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim manufacturers partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimManufacturersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithContext(ctx context.Context) *DcimManufacturersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimManufacturersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithData(data *models.Manufacturer) *DcimManufacturersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetData(data *models.Manufacturer) {
	o.Data = data
}

// WithID adds the id to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) WithID(id strfmt.UUID) *DcimManufacturersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers partial update params
func (o *DcimManufacturersPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
