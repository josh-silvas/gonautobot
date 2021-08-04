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

// NewExtrasGitRepositoriesDeleteParams creates a new ExtrasGitRepositoriesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasGitRepositoriesDeleteParams() *ExtrasGitRepositoriesDeleteParams {
	return &ExtrasGitRepositoriesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasGitRepositoriesDeleteParamsWithTimeout creates a new ExtrasGitRepositoriesDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasGitRepositoriesDeleteParamsWithTimeout(timeout time.Duration) *ExtrasGitRepositoriesDeleteParams {
	return &ExtrasGitRepositoriesDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasGitRepositoriesDeleteParamsWithContext creates a new ExtrasGitRepositoriesDeleteParams object
// with the ability to set a context for a request.
func NewExtrasGitRepositoriesDeleteParamsWithContext(ctx context.Context) *ExtrasGitRepositoriesDeleteParams {
	return &ExtrasGitRepositoriesDeleteParams{
		Context: ctx,
	}
}

// NewExtrasGitRepositoriesDeleteParamsWithHTTPClient creates a new ExtrasGitRepositoriesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasGitRepositoriesDeleteParamsWithHTTPClient(client *http.Client) *ExtrasGitRepositoriesDeleteParams {
	return &ExtrasGitRepositoriesDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasGitRepositoriesDeleteParams contains all the parameters to send to the API endpoint
   for the extras git repositories delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasGitRepositoriesDeleteParams struct {

	/* ID.

	   A UUID string identifying this Git repository.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras git repositories delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesDeleteParams) WithDefaults() *ExtrasGitRepositoriesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras git repositories delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasGitRepositoriesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) WithTimeout(timeout time.Duration) *ExtrasGitRepositoriesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) WithContext(ctx context.Context) *ExtrasGitRepositoriesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) WithHTTPClient(client *http.Client) *ExtrasGitRepositoriesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) WithID(id strfmt.UUID) *ExtrasGitRepositoriesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras git repositories delete params
func (o *ExtrasGitRepositoriesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasGitRepositoriesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
