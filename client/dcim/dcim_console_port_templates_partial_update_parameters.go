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

// NewDcimConsolePortTemplatesPartialUpdateParams creates a new DcimConsolePortTemplatesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortTemplatesPartialUpdateParams() *DcimConsolePortTemplatesPartialUpdateParams {
	return &DcimConsolePortTemplatesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesPartialUpdateParamsWithTimeout creates a new DcimConsolePortTemplatesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesPartialUpdateParams {
	return &DcimConsolePortTemplatesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesPartialUpdateParamsWithContext creates a new DcimConsolePortTemplatesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimConsolePortTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesPartialUpdateParams {
	return &DcimConsolePortTemplatesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimConsolePortTemplatesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesPartialUpdateParams {
	return &DcimConsolePortTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim console port templates partial update operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortTemplatesPartialUpdateParams struct {

	// Data.
	Data *models.WritableConsolePortTemplate

	/* ID.

	   A UUID string identifying this console port template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithDefaults() *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console port templates partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithData(data *models.WritableConsolePortTemplate) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetData(data *models.WritableConsolePortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) WithID(id strfmt.UUID) *DcimConsolePortTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console port templates partial update params
func (o *DcimConsolePortTemplatesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
