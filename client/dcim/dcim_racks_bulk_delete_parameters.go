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

// NewDcimRacksBulkDeleteParams creates a new DcimRacksBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksBulkDeleteParams() *DcimRacksBulkDeleteParams {
	return &DcimRacksBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksBulkDeleteParamsWithTimeout creates a new DcimRacksBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRacksBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimRacksBulkDeleteParams {
	return &DcimRacksBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRacksBulkDeleteParamsWithContext creates a new DcimRacksBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimRacksBulkDeleteParamsWithContext(ctx context.Context) *DcimRacksBulkDeleteParams {
	return &DcimRacksBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimRacksBulkDeleteParamsWithHTTPClient creates a new DcimRacksBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimRacksBulkDeleteParams {
	return &DcimRacksBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRacksBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim racks bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimRacksBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksBulkDeleteParams) WithDefaults() *DcimRacksBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks bulk delete params
func (o *DcimRacksBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimRacksBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks bulk delete params
func (o *DcimRacksBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks bulk delete params
func (o *DcimRacksBulkDeleteParams) WithContext(ctx context.Context) *DcimRacksBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks bulk delete params
func (o *DcimRacksBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks bulk delete params
func (o *DcimRacksBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimRacksBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks bulk delete params
func (o *DcimRacksBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
