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

// NewExtrasImageAttachmentsUpdateParams creates a new ExtrasImageAttachmentsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsUpdateParams() *ExtrasImageAttachmentsUpdateParams {
	return &ExtrasImageAttachmentsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsUpdateParamsWithTimeout creates a new ExtrasImageAttachmentsUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsUpdateParams {
	return &ExtrasImageAttachmentsUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsUpdateParamsWithContext creates a new ExtrasImageAttachmentsUpdateParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsUpdateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsUpdateParams {
	return &ExtrasImageAttachmentsUpdateParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsUpdateParamsWithHTTPClient creates a new ExtrasImageAttachmentsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsUpdateParams {
	return &ExtrasImageAttachmentsUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasImageAttachmentsUpdateParams contains all the parameters to send to the API endpoint
   for the extras image attachments update operation.

   Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsUpdateParams struct {

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

// WithDefaults hydrates default values in the extras image attachments update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsUpdateParams) WithDefaults() *ExtrasImageAttachmentsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WithID adds the id to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) WithID(id strfmt.UUID) *ExtrasImageAttachmentsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras image attachments update params
func (o *ExtrasImageAttachmentsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
