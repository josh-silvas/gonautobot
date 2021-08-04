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

// NewExtrasTagsDeleteParams creates a new ExtrasTagsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasTagsDeleteParams() *ExtrasTagsDeleteParams {
	return &ExtrasTagsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsDeleteParamsWithTimeout creates a new ExtrasTagsDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasTagsDeleteParamsWithTimeout(timeout time.Duration) *ExtrasTagsDeleteParams {
	return &ExtrasTagsDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasTagsDeleteParamsWithContext creates a new ExtrasTagsDeleteParams object
// with the ability to set a context for a request.
func NewExtrasTagsDeleteParamsWithContext(ctx context.Context) *ExtrasTagsDeleteParams {
	return &ExtrasTagsDeleteParams{
		Context: ctx,
	}
}

// NewExtrasTagsDeleteParamsWithHTTPClient creates a new ExtrasTagsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasTagsDeleteParamsWithHTTPClient(client *http.Client) *ExtrasTagsDeleteParams {
	return &ExtrasTagsDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasTagsDeleteParams contains all the parameters to send to the API endpoint
   for the extras tags delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasTagsDeleteParams struct {

	/* ID.

	   A UUID string identifying this tag.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras tags delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsDeleteParams) WithDefaults() *ExtrasTagsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras tags delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras tags delete params
func (o *ExtrasTagsDeleteParams) WithTimeout(timeout time.Duration) *ExtrasTagsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags delete params
func (o *ExtrasTagsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags delete params
func (o *ExtrasTagsDeleteParams) WithContext(ctx context.Context) *ExtrasTagsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags delete params
func (o *ExtrasTagsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags delete params
func (o *ExtrasTagsDeleteParams) WithHTTPClient(client *http.Client) *ExtrasTagsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags delete params
func (o *ExtrasTagsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras tags delete params
func (o *ExtrasTagsDeleteParams) WithID(id strfmt.UUID) *ExtrasTagsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras tags delete params
func (o *ExtrasTagsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
