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

// NewExtrasCustomLinksBulkDeleteParams creates a new ExtrasCustomLinksBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomLinksBulkDeleteParams() *ExtrasCustomLinksBulkDeleteParams {
	return &ExtrasCustomLinksBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomLinksBulkDeleteParamsWithTimeout creates a new ExtrasCustomLinksBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomLinksBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasCustomLinksBulkDeleteParams {
	return &ExtrasCustomLinksBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasCustomLinksBulkDeleteParamsWithContext creates a new ExtrasCustomLinksBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasCustomLinksBulkDeleteParamsWithContext(ctx context.Context) *ExtrasCustomLinksBulkDeleteParams {
	return &ExtrasCustomLinksBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasCustomLinksBulkDeleteParamsWithHTTPClient creates a new ExtrasCustomLinksBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomLinksBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasCustomLinksBulkDeleteParams {
	return &ExtrasCustomLinksBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomLinksBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras custom links bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomLinksBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom links bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksBulkDeleteParams) WithDefaults() *ExtrasCustomLinksBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom links bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom links bulk delete params
func (o *ExtrasCustomLinksBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasCustomLinksBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom links bulk delete params
func (o *ExtrasCustomLinksBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom links bulk delete params
func (o *ExtrasCustomLinksBulkDeleteParams) WithContext(ctx context.Context) *ExtrasCustomLinksBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom links bulk delete params
func (o *ExtrasCustomLinksBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom links bulk delete params
func (o *ExtrasCustomLinksBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasCustomLinksBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom links bulk delete params
func (o *ExtrasCustomLinksBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomLinksBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
