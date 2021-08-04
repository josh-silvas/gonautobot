package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDcimConsolePortsBulkDeleteParams creates a new DcimConsolePortsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsBulkDeleteParams() *DcimConsolePortsBulkDeleteParams {
	return &DcimConsolePortsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsBulkDeleteParamsWithTimeout creates a new DcimConsolePortsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimConsolePortsBulkDeleteParams {
	return &DcimConsolePortsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsBulkDeleteParamsWithContext creates a new DcimConsolePortsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsBulkDeleteParamsWithContext(ctx context.Context) *DcimConsolePortsBulkDeleteParams {
	return &DcimConsolePortsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsBulkDeleteParamsWithHTTPClient creates a new DcimConsolePortsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimConsolePortsBulkDeleteParams {
	return &DcimConsolePortsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim console ports bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsBulkDeleteParams) WithDefaults() *DcimConsolePortsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports bulk delete params
func (o *DcimConsolePortsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimConsolePortsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports bulk delete params
func (o *DcimConsolePortsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports bulk delete params
func (o *DcimConsolePortsBulkDeleteParams) WithContext(ctx context.Context) *DcimConsolePortsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports bulk delete params
func (o *DcimConsolePortsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports bulk delete params
func (o *DcimConsolePortsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimConsolePortsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports bulk delete params
func (o *DcimConsolePortsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
