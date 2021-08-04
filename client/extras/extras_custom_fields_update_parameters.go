package extras

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

// NewExtrasCustomFieldsUpdateParams creates a new ExtrasCustomFieldsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldsUpdateParams() *ExtrasCustomFieldsUpdateParams {
	return &ExtrasCustomFieldsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldsUpdateParamsWithTimeout creates a new ExtrasCustomFieldsUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldsUpdateParams {
	return &ExtrasCustomFieldsUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldsUpdateParamsWithContext creates a new ExtrasCustomFieldsUpdateParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldsUpdateParamsWithContext(ctx context.Context) *ExtrasCustomFieldsUpdateParams {
	return &ExtrasCustomFieldsUpdateParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldsUpdateParamsWithHTTPClient creates a new ExtrasCustomFieldsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldsUpdateParams {
	return &ExtrasCustomFieldsUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldsUpdateParams contains all the parameters to send to the API endpoint
   for the extras custom fields update operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldsUpdateParams struct {

	// Data.
	Data *models.WritableCustomField

	/* ID.

	   A UUID string identifying this custom field.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom fields update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsUpdateParams) WithDefaults() *ExtrasCustomFieldsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom fields update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) WithContext(ctx context.Context) *ExtrasCustomFieldsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) WithData(data *models.WritableCustomField) *ExtrasCustomFieldsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) SetData(data *models.WritableCustomField) {
	o.Data = data
}

// WithID adds the id to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) WithID(id strfmt.UUID) *ExtrasCustomFieldsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom fields update params
func (o *ExtrasCustomFieldsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
