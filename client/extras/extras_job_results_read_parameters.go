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

// NewExtrasJobResultsReadParams creates a new ExtrasJobResultsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasJobResultsReadParams() *ExtrasJobResultsReadParams {
	return &ExtrasJobResultsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasJobResultsReadParamsWithTimeout creates a new ExtrasJobResultsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasJobResultsReadParamsWithTimeout(timeout time.Duration) *ExtrasJobResultsReadParams {
	return &ExtrasJobResultsReadParams{
		timeout: timeout,
	}
}

// NewExtrasJobResultsReadParamsWithContext creates a new ExtrasJobResultsReadParams object
// with the ability to set a context for a request.
func NewExtrasJobResultsReadParamsWithContext(ctx context.Context) *ExtrasJobResultsReadParams {
	return &ExtrasJobResultsReadParams{
		Context: ctx,
	}
}

// NewExtrasJobResultsReadParamsWithHTTPClient creates a new ExtrasJobResultsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasJobResultsReadParamsWithHTTPClient(client *http.Client) *ExtrasJobResultsReadParams {
	return &ExtrasJobResultsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasJobResultsReadParams contains all the parameters to send to the API endpoint
   for the extras job results read operation.

   Typically these are written to a http.Request.
*/
type ExtrasJobResultsReadParams struct {

	/* ID.

	   A UUID string identifying this job result.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras job results read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsReadParams) WithDefaults() *ExtrasJobResultsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras job results read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasJobResultsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras job results read params
func (o *ExtrasJobResultsReadParams) WithTimeout(timeout time.Duration) *ExtrasJobResultsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras job results read params
func (o *ExtrasJobResultsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras job results read params
func (o *ExtrasJobResultsReadParams) WithContext(ctx context.Context) *ExtrasJobResultsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras job results read params
func (o *ExtrasJobResultsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras job results read params
func (o *ExtrasJobResultsReadParams) WithHTTPClient(client *http.Client) *ExtrasJobResultsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras job results read params
func (o *ExtrasJobResultsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras job results read params
func (o *ExtrasJobResultsReadParams) WithID(id strfmt.UUID) *ExtrasJobResultsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras job results read params
func (o *ExtrasJobResultsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasJobResultsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
