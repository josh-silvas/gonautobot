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

// NewExtrasConfigContextsDeleteParams creates a new ExtrasConfigContextsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsDeleteParams() *ExtrasConfigContextsDeleteParams {
	return &ExtrasConfigContextsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsDeleteParamsWithTimeout creates a new ExtrasConfigContextsDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsDeleteParams {
	return &ExtrasConfigContextsDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsDeleteParamsWithContext creates a new ExtrasConfigContextsDeleteParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsDeleteParamsWithContext(ctx context.Context) *ExtrasConfigContextsDeleteParams {
	return &ExtrasConfigContextsDeleteParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsDeleteParamsWithHTTPClient creates a new ExtrasConfigContextsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsDeleteParams {
	return &ExtrasConfigContextsDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextsDeleteParams contains all the parameters to send to the API endpoint
   for the extras config contexts delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextsDeleteParams struct {

	/* ID.

	   A UUID string identifying this config context.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsDeleteParams) WithDefaults() *ExtrasConfigContextsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) WithContext(ctx context.Context) *ExtrasConfigContextsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) WithID(id strfmt.UUID) *ExtrasConfigContextsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config contexts delete params
func (o *ExtrasConfigContextsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
