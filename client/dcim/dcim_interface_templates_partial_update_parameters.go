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

// NewDcimInterfaceTemplatesPartialUpdateParams creates a new DcimInterfaceTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfaceTemplatesPartialUpdateParams() *DcimInterfaceTemplatesPartialUpdateParams {
	return &DcimInterfaceTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesPartialUpdateParamsWithTimeout creates a new DcimInterfaceTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimInterfaceTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesPartialUpdateParams {
	return &DcimInterfaceTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesPartialUpdateParamsWithContext creates a new DcimInterfaceTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimInterfaceTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesPartialUpdateParams {
	return &DcimInterfaceTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimInterfaceTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfaceTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesPartialUpdateParams {
	return &DcimInterfaceTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimInterfaceTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim interface templates partial update operation.

   Typically these are written to a http.Request.
*/
type DcimInterfaceTemplatesPartialUpdateParams struct {

	// Data.
	Data *models.WritableInterfaceTemplate

	/* ID.

	   A UUID string identifying this interface template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interface templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithDefaults() *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interface templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithData(data *models.WritableInterfaceTemplate) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetData(data *models.WritableInterfaceTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) WithID(id strfmt.UUID) *DcimInterfaceTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interface templates partial update params
func (o *DcimInterfaceTemplatesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
