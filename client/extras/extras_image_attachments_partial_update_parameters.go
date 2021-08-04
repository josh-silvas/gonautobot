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

// NewExtrasImageAttachmentsPartialUpdateParams creates a new ExtrasImageAttachmentsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsPartialUpdateParams() *ExtrasImageAttachmentsPartialUpdateParams {
	return &ExtrasImageAttachmentsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsPartialUpdateParamsWithTimeout creates a new ExtrasImageAttachmentsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsPartialUpdateParams {
	return &ExtrasImageAttachmentsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsPartialUpdateParamsWithContext creates a new ExtrasImageAttachmentsPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsPartialUpdateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsPartialUpdateParams {
	return &ExtrasImageAttachmentsPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsPartialUpdateParamsWithHTTPClient creates a new ExtrasImageAttachmentsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsPartialUpdateParams {
	return &ExtrasImageAttachmentsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasImageAttachmentsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras image attachments partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsPartialUpdateParams struct {

	// Data.
	Data *models.ImageAttachment

	/* ID.

	   A UUID string identifying this image attachment.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras image attachments partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithDefaults() *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WithID adds the id to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) WithID(id strfmt.UUID) *ExtrasImageAttachmentsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments partial update params
func (o *ExtrasImageAttachmentsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
