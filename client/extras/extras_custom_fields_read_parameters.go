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

// NewExtrasCustomFieldsReadParams creates a new ExtrasCustomFieldsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomFieldsReadParams() *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomFieldsReadParamsWithTimeout creates a new ExtrasCustomFieldsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomFieldsReadParamsWithTimeout(timeout time.Duration) *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		timeout: timeout,
	}
}

// NewExtrasCustomFieldsReadParamsWithContext creates a new ExtrasCustomFieldsReadParams object
// with the ability to set a context for a request.
func NewExtrasCustomFieldsReadParamsWithContext(ctx context.Context) *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		Context: ctx,
	}
}

// NewExtrasCustomFieldsReadParamsWithHTTPClient creates a new ExtrasCustomFieldsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomFieldsReadParamsWithHTTPClient(client *http.Client) *ExtrasCustomFieldsReadParams {
	return &ExtrasCustomFieldsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomFieldsReadParams contains all the parameters to send to the API endpoint
   for the extras custom fields read operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomFieldsReadParams struct {

	/* ID.

	   A UUID string identifying this custom field.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom fields read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsReadParams) WithDefaults() *ExtrasCustomFieldsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom fields read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomFieldsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithTimeout(timeout time.Duration) *ExtrasCustomFieldsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithContext(ctx context.Context) *ExtrasCustomFieldsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithHTTPClient(client *http.Client) *ExtrasCustomFieldsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) WithID(id strfmt.UUID) *ExtrasCustomFieldsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom fields read params
func (o *ExtrasCustomFieldsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomFieldsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
