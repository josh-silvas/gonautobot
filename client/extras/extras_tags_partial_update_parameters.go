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

// NewExtrasTagsPartialUpdateParams creates a new ExtrasTagsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasTagsPartialUpdateParams() *ExtrasTagsPartialUpdateParams {
	return &ExtrasTagsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsPartialUpdateParamsWithTimeout creates a new ExtrasTagsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasTagsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTagsPartialUpdateParams {
	return &ExtrasTagsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasTagsPartialUpdateParamsWithContext creates a new ExtrasTagsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasTagsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasTagsPartialUpdateParams {
	return &ExtrasTagsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasTagsPartialUpdateParamsWithHTTPClient creates a new ExtrasTagsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasTagsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTagsPartialUpdateParams {
	return &ExtrasTagsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasTagsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras tags partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasTagsPartialUpdateParams struct {

	// Data.
	Data *models.Tag

	/* ID.

	   A UUID string identifying this tag.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras tags partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsPartialUpdateParams) WithDefaults() *ExtrasTagsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras tags partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTagsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasTagsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTagsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithData(data *models.Tag) *ExtrasTagsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetData(data *models.Tag) {
	o.Data = data
}

// WithID adds the id to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasTagsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras tags partial update params
func (o *ExtrasTagsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
