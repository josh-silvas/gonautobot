package circuits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCircuitsProvidersBulkDeleteParams creates a new CircuitsProvidersBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersBulkDeleteParams() *CircuitsProvidersBulkDeleteParams {
	return &CircuitsProvidersBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersBulkDeleteParamsWithTimeout creates a new CircuitsProvidersBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersBulkDeleteParamsWithTimeout(timeout time.Duration) *CircuitsProvidersBulkDeleteParams {
	return &CircuitsProvidersBulkDeleteParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersBulkDeleteParamsWithContext creates a new CircuitsProvidersBulkDeleteParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersBulkDeleteParamsWithContext(ctx context.Context) *CircuitsProvidersBulkDeleteParams {
	return &CircuitsProvidersBulkDeleteParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersBulkDeleteParamsWithHTTPClient creates a new CircuitsProvidersBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersBulkDeleteParamsWithHTTPClient(client *http.Client) *CircuitsProvidersBulkDeleteParams {
	return &CircuitsProvidersBulkDeleteParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersBulkDeleteParams contains all the parameters to send to the API endpoint
   for the circuits providers bulk delete operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersBulkDeleteParams) WithDefaults() *CircuitsProvidersBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers bulk delete params
func (o *CircuitsProvidersBulkDeleteParams) WithTimeout(timeout time.Duration) *CircuitsProvidersBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers bulk delete params
func (o *CircuitsProvidersBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers bulk delete params
func (o *CircuitsProvidersBulkDeleteParams) WithContext(ctx context.Context) *CircuitsProvidersBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers bulk delete params
func (o *CircuitsProvidersBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers bulk delete params
func (o *CircuitsProvidersBulkDeleteParams) WithHTTPClient(client *http.Client) *CircuitsProvidersBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers bulk delete params
func (o *CircuitsProvidersBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
