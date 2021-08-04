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

// NewDcimCablesBulkDeleteParams creates a new DcimCablesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesBulkDeleteParams() *DcimCablesBulkDeleteParams {
	return &DcimCablesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesBulkDeleteParamsWithTimeout creates a new DcimCablesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimCablesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimCablesBulkDeleteParams {
	return &DcimCablesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimCablesBulkDeleteParamsWithContext creates a new DcimCablesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimCablesBulkDeleteParamsWithContext(ctx context.Context) *DcimCablesBulkDeleteParams {
	return &DcimCablesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimCablesBulkDeleteParamsWithHTTPClient creates a new DcimCablesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimCablesBulkDeleteParams {
	return &DcimCablesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimCablesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim cables bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimCablesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesBulkDeleteParams) WithDefaults() *DcimCablesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables bulk delete params
func (o *DcimCablesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimCablesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables bulk delete params
func (o *DcimCablesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables bulk delete params
func (o *DcimCablesBulkDeleteParams) WithContext(ctx context.Context) *DcimCablesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables bulk delete params
func (o *DcimCablesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables bulk delete params
func (o *DcimCablesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimCablesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables bulk delete params
func (o *DcimCablesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
