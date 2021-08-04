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

// NewDcimRegionsBulkDeleteParams creates a new DcimRegionsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRegionsBulkDeleteParams() *DcimRegionsBulkDeleteParams {
	return &DcimRegionsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRegionsBulkDeleteParamsWithTimeout creates a new DcimRegionsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRegionsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimRegionsBulkDeleteParams {
	return &DcimRegionsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRegionsBulkDeleteParamsWithContext creates a new DcimRegionsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimRegionsBulkDeleteParamsWithContext(ctx context.Context) *DcimRegionsBulkDeleteParams {
	return &DcimRegionsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimRegionsBulkDeleteParamsWithHTTPClient creates a new DcimRegionsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRegionsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimRegionsBulkDeleteParams {
	return &DcimRegionsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRegionsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim regions bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimRegionsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim regions bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsBulkDeleteParams) WithDefaults() *DcimRegionsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim regions bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim regions bulk delete params
func (o *DcimRegionsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimRegionsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim regions bulk delete params
func (o *DcimRegionsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim regions bulk delete params
func (o *DcimRegionsBulkDeleteParams) WithContext(ctx context.Context) *DcimRegionsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim regions bulk delete params
func (o *DcimRegionsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim regions bulk delete params
func (o *DcimRegionsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimRegionsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim regions bulk delete params
func (o *DcimRegionsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRegionsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
