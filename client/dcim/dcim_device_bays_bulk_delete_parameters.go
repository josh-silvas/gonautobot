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

// NewDcimDeviceBaysBulkDeleteParams creates a new DcimDeviceBaysBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBaysBulkDeleteParams() *DcimDeviceBaysBulkDeleteParams {
	return &DcimDeviceBaysBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBaysBulkDeleteParamsWithTimeout creates a new DcimDeviceBaysBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBaysBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceBaysBulkDeleteParams {
	return &DcimDeviceBaysBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBaysBulkDeleteParamsWithContext creates a new DcimDeviceBaysBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimDeviceBaysBulkDeleteParamsWithContext(ctx context.Context) *DcimDeviceBaysBulkDeleteParams {
	return &DcimDeviceBaysBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimDeviceBaysBulkDeleteParamsWithHTTPClient creates a new DcimDeviceBaysBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBaysBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceBaysBulkDeleteParams {
	return &DcimDeviceBaysBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBaysBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim device bays bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBaysBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bays bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysBulkDeleteParams) WithDefaults() *DcimDeviceBaysBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bays bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bays bulk delete params
func (o *DcimDeviceBaysBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceBaysBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bays bulk delete params
func (o *DcimDeviceBaysBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bays bulk delete params
func (o *DcimDeviceBaysBulkDeleteParams) WithContext(ctx context.Context) *DcimDeviceBaysBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bays bulk delete params
func (o *DcimDeviceBaysBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bays bulk delete params
func (o *DcimDeviceBaysBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceBaysBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bays bulk delete params
func (o *DcimDeviceBaysBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBaysBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
