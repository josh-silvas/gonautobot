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

// NewExtrasWebhooksBulkPartialUpdateParams creates a new ExtrasWebhooksBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasWebhooksBulkPartialUpdateParams() *ExtrasWebhooksBulkPartialUpdateParams {
	return &ExtrasWebhooksBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasWebhooksBulkPartialUpdateParamsWithTimeout creates a new ExtrasWebhooksBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasWebhooksBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *ExtrasWebhooksBulkPartialUpdateParams {
	return &ExtrasWebhooksBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasWebhooksBulkPartialUpdateParamsWithContext creates a new ExtrasWebhooksBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewExtrasWebhooksBulkPartialUpdateParamsWithContext(ctx context.Context) *ExtrasWebhooksBulkPartialUpdateParams {
	return &ExtrasWebhooksBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewExtrasWebhooksBulkPartialUpdateParamsWithHTTPClient creates a new ExtrasWebhooksBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasWebhooksBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *ExtrasWebhooksBulkPartialUpdateParams {
	return &ExtrasWebhooksBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasWebhooksBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the extras webhooks bulk partial update operation.

   Typically these are written to a http.Request.
*/
type ExtrasWebhooksBulkPartialUpdateParams struct {

	// Data.
	Data *models.Webhook

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras webhooks bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksBulkPartialUpdateParams) WithDefaults() *ExtrasWebhooksBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras webhooks bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *ExtrasWebhooksBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) WithContext(ctx context.Context) *ExtrasWebhooksBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *ExtrasWebhooksBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) WithData(data *models.Webhook) *ExtrasWebhooksBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras webhooks bulk partial update params
func (o *ExtrasWebhooksBulkPartialUpdateParams) SetData(data *models.Webhook) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasWebhooksBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
