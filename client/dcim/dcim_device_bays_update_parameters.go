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

// NewDcimDeviceBaysUpdateParams creates a new DcimDeviceBaysUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBaysUpdateParams() *DcimDeviceBaysUpdateParams {
	return &DcimDeviceBaysUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBaysUpdateParamsWithTimeout creates a new DcimDeviceBaysUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBaysUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBaysUpdateParams {
	return &DcimDeviceBaysUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBaysUpdateParamsWithContext creates a new DcimDeviceBaysUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceBaysUpdateParamsWithContext(ctx context.Context) *DcimDeviceBaysUpdateParams {
	return &DcimDeviceBaysUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceBaysUpdateParamsWithHTTPClient creates a new DcimDeviceBaysUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBaysUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBaysUpdateParams {
	return &DcimDeviceBaysUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBaysUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device bays update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBaysUpdateParams struct {

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

// WithDefaults hydrates default values in the dcim device bays update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysUpdateParams) WithDefaults() *DcimDeviceBaysUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bays update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBaysUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithContext(ctx context.Context) *DcimDeviceBaysUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBaysUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithData(data *models.WritableDeviceBay) *DcimDeviceBaysUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetData(data *models.WritableDeviceBay) {
	o.Data = data
}

// WithID adds the id to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) WithID(id strfmt.UUID) *DcimDeviceBaysUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bays update params
func (o *DcimDeviceBaysUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBaysUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
