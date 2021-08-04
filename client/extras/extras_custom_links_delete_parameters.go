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

// NewExtrasCustomLinksDeleteParams creates a new ExtrasCustomLinksDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomLinksDeleteParams() *ExtrasCustomLinksDeleteParams {
	return &ExtrasCustomLinksDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomLinksDeleteParamsWithTimeout creates a new ExtrasCustomLinksDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomLinksDeleteParamsWithTimeout(timeout time.Duration) *ExtrasCustomLinksDeleteParams {
	return &ExtrasCustomLinksDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasCustomLinksDeleteParamsWithContext creates a new ExtrasCustomLinksDeleteParams object
// with the ability to set a context for a request.
func NewExtrasCustomLinksDeleteParamsWithContext(ctx context.Context) *ExtrasCustomLinksDeleteParams {
	return &ExtrasCustomLinksDeleteParams{
		Context: ctx,
	}
}

// NewExtrasCustomLinksDeleteParamsWithHTTPClient creates a new ExtrasCustomLinksDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomLinksDeleteParamsWithHTTPClient(client *http.Client) *ExtrasCustomLinksDeleteParams {
	return &ExtrasCustomLinksDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomLinksDeleteParams contains all the parameters to send to the API endpoint
   for the extras custom links delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomLinksDeleteParams struct {

	/* ID.

	   A UUID string identifying this custom link.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom links delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksDeleteParams) WithDefaults() *ExtrasCustomLinksDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom links delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) WithTimeout(timeout time.Duration) *ExtrasCustomLinksDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) WithContext(ctx context.Context) *ExtrasCustomLinksDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) WithHTTPClient(client *http.Client) *ExtrasCustomLinksDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) WithID(id strfmt.UUID) *ExtrasCustomLinksDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom links delete params
func (o *ExtrasCustomLinksDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomLinksDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
