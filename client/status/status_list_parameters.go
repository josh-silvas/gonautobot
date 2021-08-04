package status

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewStatusListParams creates a new StatusListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStatusListParams() *StatusListParams {
	return &StatusListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStatusListParamsWithTimeout creates a new StatusListParams object
// with the ability to set a timeout on a request.
func NewStatusListParamsWithTimeout(timeout time.Duration) *StatusListParams {
	return &StatusListParams{
		timeout: timeout,
	}
}

// NewStatusListParamsWithContext creates a new StatusListParams object
// with the ability to set a context for a request.
func NewStatusListParamsWithContext(ctx context.Context) *StatusListParams {
	return &StatusListParams{
		Context: ctx,
	}
}

// NewStatusListParamsWithHTTPClient creates a new StatusListParams object
// with the ability to set a custom HTTPClient for a request.
func NewStatusListParamsWithHTTPClient(client *http.Client) *StatusListParams {
	return &StatusListParams{
		HTTPClient: client,
	}
}

/* StatusListParams contains all the parameters to send to the API endpoint
   for the status list operation.

   Typically these are written to a http.Request.
*/
type StatusListParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the status list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusListParams) WithDefaults() *StatusListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the status list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StatusListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the status list params
func (o *StatusListParams) WithTimeout(timeout time.Duration) *StatusListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status list params
func (o *StatusListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status list params
func (o *StatusListParams) WithContext(ctx context.Context) *StatusListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status list params
func (o *StatusListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status list params
func (o *StatusListParams) WithHTTPClient(client *http.Client) *StatusListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status list params
func (o *StatusListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StatusListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
