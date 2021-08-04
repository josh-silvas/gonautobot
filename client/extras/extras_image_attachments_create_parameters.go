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

// NewExtrasImageAttachmentsCreateParams creates a new ExtrasImageAttachmentsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasImageAttachmentsCreateParams() *ExtrasImageAttachmentsCreateParams {
	return &ExtrasImageAttachmentsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasImageAttachmentsCreateParamsWithTimeout creates a new ExtrasImageAttachmentsCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasImageAttachmentsCreateParamsWithTimeout(timeout time.Duration) *ExtrasImageAttachmentsCreateParams {
	return &ExtrasImageAttachmentsCreateParams{
		timeout: timeout,
	}
}

// NewExtrasImageAttachmentsCreateParamsWithContext creates a new ExtrasImageAttachmentsCreateParams object
// with the ability to set a context for a request.
func NewExtrasImageAttachmentsCreateParamsWithContext(ctx context.Context) *ExtrasImageAttachmentsCreateParams {
	return &ExtrasImageAttachmentsCreateParams{
		Context: ctx,
	}
}

// NewExtrasImageAttachmentsCreateParamsWithHTTPClient creates a new ExtrasImageAttachmentsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasImageAttachmentsCreateParamsWithHTTPClient(client *http.Client) *ExtrasImageAttachmentsCreateParams {
	return &ExtrasImageAttachmentsCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasImageAttachmentsCreateParams contains all the parameters to send to the API endpoint
   for the extras image attachments create operation.

   Typically these are written to a http.Request.
*/
type ExtrasImageAttachmentsCreateParams struct {

	// Data.
	Data *models.ImageAttachment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras image attachments create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsCreateParams) WithDefaults() *ExtrasImageAttachmentsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras image attachments create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasImageAttachmentsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithTimeout(timeout time.Duration) *ExtrasImageAttachmentsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithContext(ctx context.Context) *ExtrasImageAttachmentsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithHTTPClient(client *http.Client) *ExtrasImageAttachmentsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) WithData(data *models.ImageAttachment) *ExtrasImageAttachmentsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras image attachments create params
func (o *ExtrasImageAttachmentsCreateParams) SetData(data *models.ImageAttachment) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasImageAttachmentsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
