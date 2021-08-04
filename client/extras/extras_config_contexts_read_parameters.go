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

// NewExtrasConfigContextsReadParams creates a new ExtrasConfigContextsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsReadParams() *ExtrasConfigContextsReadParams {
	return &ExtrasConfigContextsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsReadParamsWithTimeout creates a new ExtrasConfigContextsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsReadParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsReadParams {
	return &ExtrasConfigContextsReadParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsReadParamsWithContext creates a new ExtrasConfigContextsReadParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsReadParamsWithContext(ctx context.Context) *ExtrasConfigContextsReadParams {
	return &ExtrasConfigContextsReadParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsReadParamsWithHTTPClient creates a new ExtrasConfigContextsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsReadParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsReadParams {
	return &ExtrasConfigContextsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextsReadParams contains all the parameters to send to the API endpoint
   for the extras config contexts read operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextsReadParams struct {

	/* ID.

	   A UUID string identifying this config context.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsReadParams) WithDefaults() *ExtrasConfigContextsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) WithContext(ctx context.Context) *ExtrasConfigContextsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) WithID(id strfmt.UUID) *ExtrasConfigContextsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config contexts read params
func (o *ExtrasConfigContextsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
