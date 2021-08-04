package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExtrasWebhooksDeleteParams creates a new ExtrasWebhooksDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasWebhooksDeleteParams() *ExtrasWebhooksDeleteParams {
	return &ExtrasWebhooksDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasWebhooksDeleteParamsWithTimeout creates a new ExtrasWebhooksDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasWebhooksDeleteParamsWithTimeout(timeout time.Duration) *ExtrasWebhooksDeleteParams {
	return &ExtrasWebhooksDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasWebhooksDeleteParamsWithContext creates a new ExtrasWebhooksDeleteParams object
// with the ability to set a context for a request.
func NewExtrasWebhooksDeleteParamsWithContext(ctx context.Context) *ExtrasWebhooksDeleteParams {
	return &ExtrasWebhooksDeleteParams{
		Context: ctx,
	}
}

// NewExtrasWebhooksDeleteParamsWithHTTPClient creates a new ExtrasWebhooksDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasWebhooksDeleteParamsWithHTTPClient(client *http.Client) *ExtrasWebhooksDeleteParams {
	return &ExtrasWebhooksDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasWebhooksDeleteParams contains all the parameters to send to the API endpoint
   for the extras webhooks delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasWebhooksDeleteParams struct {

	/* ID.

	   A UUID string identifying this webhook.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras webhooks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksDeleteParams) WithDefaults() *ExtrasWebhooksDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras webhooks delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) WithTimeout(timeout time.Duration) *ExtrasWebhooksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) WithContext(ctx context.Context) *ExtrasWebhooksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) WithHTTPClient(client *http.Client) *ExtrasWebhooksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) WithID(id strfmt.UUID) *ExtrasWebhooksDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras webhooks delete params
func (o *ExtrasWebhooksDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasWebhooksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
