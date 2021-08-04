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

// NewDcimRearPortTemplatesPartialUpdateParams creates a new DcimRearPortTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesPartialUpdateParams() *DcimRearPortTemplatesPartialUpdateParams {
	return &DcimRearPortTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesPartialUpdateParamsWithTimeout creates a new DcimRearPortTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesPartialUpdateParams {
	return &DcimRearPortTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesPartialUpdateParamsWithContext creates a new DcimRearPortTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimRearPortTemplatesPartialUpdateParams {
	return &DcimRearPortTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimRearPortTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesPartialUpdateParams {
	return &DcimRearPortTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRearPortTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rear port templates partial update operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesPartialUpdateParams struct {

	// Data.
	Data *models.WritableRearPortTemplate

	/* ID.

	   A UUID string identifying this rear port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesPartialUpdateParams) WithDefaults() *DcimRearPortTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimRearPortTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) WithData(data *models.WritableRearPortTemplate) *DcimRearPortTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) SetData(data *models.WritableRearPortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) WithID(id strfmt.UUID) *DcimRearPortTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear port templates partial update params
func (o *DcimRearPortTemplatesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
