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

// NewDcimRearPortsBulkDeleteParams creates a new DcimRearPortsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsBulkDeleteParams() *DcimRearPortsBulkDeleteParams {
	return &DcimRearPortsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsBulkDeleteParamsWithTimeout creates a new DcimRearPortsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimRearPortsBulkDeleteParams {
	return &DcimRearPortsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsBulkDeleteParamsWithContext creates a new DcimRearPortsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimRearPortsBulkDeleteParamsWithContext(ctx context.Context) *DcimRearPortsBulkDeleteParams {
	return &DcimRearPortsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimRearPortsBulkDeleteParamsWithHTTPClient creates a new DcimRearPortsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimRearPortsBulkDeleteParams {
	return &DcimRearPortsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRearPortsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rear ports bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsBulkDeleteParams) WithDefaults() *DcimRearPortsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports bulk delete params
func (o *DcimRearPortsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimRearPortsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports bulk delete params
func (o *DcimRearPortsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports bulk delete params
func (o *DcimRearPortsBulkDeleteParams) WithContext(ctx context.Context) *DcimRearPortsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports bulk delete params
func (o *DcimRearPortsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports bulk delete params
func (o *DcimRearPortsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimRearPortsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports bulk delete params
func (o *DcimRearPortsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
