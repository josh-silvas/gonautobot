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

// NewExtrasGitRepositoriesBulkDeleteParams creates a new ExtrasGitRepositoriesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesBulkDeleteParams() *ExtrasGitRepositoriesBulkDeleteParams {
	return &ExtrasGitRepositoriesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesBulkDeleteParamsWithTimeout creates a new ExtrasGitRepositoriesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesBulkDeleteParams {
	return &ExtrasGitRepositoriesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesBulkDeleteParamsWithContext creates a new ExtrasGitRepositoriesBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesBulkDeleteParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesBulkDeleteParams {
	return &ExtrasGitRepositoriesBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesBulkDeleteParamsWithHTTPClient creates a new ExtrasGitRepositoriesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesBulkDeleteParams {
	return &ExtrasGitRepositoriesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras git repositories bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesBulkDeleteParams) WithDefaults() *ExtrasGitRepositoriesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories bulk delete params
func (o *ExtrasGitRepositoriesBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories bulk delete params
func (o *ExtrasGitRepositoriesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories bulk delete params
func (o *ExtrasGitRepositoriesBulkDeleteParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories bulk delete params
func (o *ExtrasGitRepositoriesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories bulk delete params
func (o *ExtrasGitRepositoriesBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories bulk delete params
func (o *ExtrasGitRepositoriesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
