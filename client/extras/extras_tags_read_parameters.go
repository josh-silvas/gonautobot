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

// NewExtrasTagsReadParams creates a new ExtrasTagsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasTagsReadParams() *ExtrasTagsReadParams {
	return &ExtrasTagsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsReadParamsWithTimeout creates a new ExtrasTagsReadParams object
// with the ability to set a timeout on a request.
func NewExtrasTagsReadParamsWithTimeout(timeout time.Duration) *ExtrasTagsReadParams {
	return &ExtrasTagsReadParams{
		timeout: timeout,
	}
}

// NewExtrasTagsReadParamsWithContext creates a new ExtrasTagsReadParams object
// with the ability to set a context for a request.
func NewExtrasTagsReadParamsWithContext(ctx context.Context) *ExtrasTagsReadParams {
	return &ExtrasTagsReadParams{
		Context: ctx,
	}
}

// NewExtrasTagsReadParamsWithHTTPClient creates a new ExtrasTagsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasTagsReadParamsWithHTTPClient(client *http.Client) *ExtrasTagsReadParams {
	return &ExtrasTagsReadParams{
		HTTPClient: client,
	}
}

/* ExtrasTagsReadParams contains all the parameters to send to the API endpoint
   for the extras tags read operation.

   Typically these are written to a http.Request.
*/
type ExtrasTagsReadParams struct {

	/* ID.

	   A UUID string identifying this tag.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras tags read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsReadParams) WithDefaults() *ExtrasTagsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras tags read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras tags read params
func (o *ExtrasTagsReadParams) WithTimeout(timeout time.Duration) *ExtrasTagsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags read params
func (o *ExtrasTagsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags read params
func (o *ExtrasTagsReadParams) WithContext(ctx context.Context) *ExtrasTagsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags read params
func (o *ExtrasTagsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags read params
func (o *ExtrasTagsReadParams) WithHTTPClient(client *http.Client) *ExtrasTagsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags read params
func (o *ExtrasTagsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras tags read params
func (o *ExtrasTagsReadParams) WithID(id strfmt.UUID) *ExtrasTagsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras tags read params
func (o *ExtrasTagsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
