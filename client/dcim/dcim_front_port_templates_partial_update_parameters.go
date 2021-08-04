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

// NewDcimFrontPortTemplatesPartialUpdateParams creates a new DcimFrontPortTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesPartialUpdateParams() *DcimFrontPortTemplatesPartialUpdateParams {
	return &DcimFrontPortTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesPartialUpdateParamsWithTimeout creates a new DcimFrontPortTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesPartialUpdateParams {
	return &DcimFrontPortTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesPartialUpdateParamsWithContext creates a new DcimFrontPortTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesPartialUpdateParams {
	return &DcimFrontPortTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimFrontPortTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesPartialUpdateParams {
	return &DcimFrontPortTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim front port templates partial update operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesPartialUpdateParams struct {

	// Data.
	Data *models.WritableFrontPortTemplate

	/* ID.

	   A UUID string identifying this front port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithDefaults() *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithData(data *models.WritableFrontPortTemplate) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetData(data *models.WritableFrontPortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) WithID(id strfmt.UUID) *DcimFrontPortTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front port templates partial update params
func (o *DcimFrontPortTemplatesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
