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

// NewDcimDeviceBaysPartialUpdateParams creates a new DcimDeviceBaysPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBaysPartialUpdateParams() *DcimDeviceBaysPartialUpdateParams {
	return &DcimDeviceBaysPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBaysPartialUpdateParamsWithTimeout creates a new DcimDeviceBaysPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBaysPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBaysPartialUpdateParams {
	return &DcimDeviceBaysPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBaysPartialUpdateParamsWithContext creates a new DcimDeviceBaysPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceBaysPartialUpdateParamsWithContext(ctx context.Context) *DcimDeviceBaysPartialUpdateParams {
	return &DcimDeviceBaysPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceBaysPartialUpdateParamsWithHTTPClient creates a new DcimDeviceBaysPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBaysPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBaysPartialUpdateParams {
	return &DcimDeviceBaysPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBaysPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device bays partial update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBaysPartialUpdateParams struct {

	// Data.
	Data *models.WritableDeviceBay

	/* ID.

	   A UUID string identifying this device bay.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bays partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysPartialUpdateParams) WithDefaults() *DcimDeviceBaysPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bays partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBaysPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) WithContext(ctx context.Context) *DcimDeviceBaysPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBaysPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) WithData(data *models.WritableDeviceBay) *DcimDeviceBaysPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) SetData(data *models.WritableDeviceBay) {
	o.Data = data
}

// WithID adds the id to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) WithID(id strfmt.UUID) *DcimDeviceBaysPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bays partial update params
func (o *DcimDeviceBaysPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBaysPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
