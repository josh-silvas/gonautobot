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

// NewDcimDeviceBayTemplatesPartialUpdateParams creates a new DcimDeviceBayTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBayTemplatesPartialUpdateParams() *DcimDeviceBayTemplatesPartialUpdateParams {
	return &DcimDeviceBayTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateParamsWithTimeout creates a new DcimDeviceBayTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBayTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesPartialUpdateParams {
	return &DcimDeviceBayTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateParamsWithContext creates a new DcimDeviceBayTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceBayTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimDeviceBayTemplatesPartialUpdateParams {
	return &DcimDeviceBayTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceBayTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimDeviceBayTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBayTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesPartialUpdateParams {
	return &DcimDeviceBayTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBayTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device bay templates partial update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBayTemplatesPartialUpdateParams struct {

	// Data.
	Data *models.WritableDeviceBayTemplate

	/* ID.

	   A UUID string identifying this device bay template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bay templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithDefaults() *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bay templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithData(data *models.WritableDeviceBayTemplate) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetData(data *models.WritableDeviceBayTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WithID(id strfmt.UUID) *DcimDeviceBayTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bay templates partial update params
func (o *DcimDeviceBayTemplatesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBayTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
