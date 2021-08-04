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

// NewDcimPlatformsBulkDeleteParams creates a new DcimPlatformsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPlatformsBulkDeleteParams() *DcimPlatformsBulkDeleteParams {
	return &DcimPlatformsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsBulkDeleteParamsWithTimeout creates a new DcimPlatformsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPlatformsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPlatformsBulkDeleteParams {
	return &DcimPlatformsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPlatformsBulkDeleteParamsWithContext creates a new DcimPlatformsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimPlatformsBulkDeleteParamsWithContext(ctx context.Context) *DcimPlatformsBulkDeleteParams {
	return &DcimPlatformsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimPlatformsBulkDeleteParamsWithHTTPClient creates a new DcimPlatformsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPlatformsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPlatformsBulkDeleteParams {
	return &DcimPlatformsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPlatformsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim platforms bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimPlatformsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim platforms bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsBulkDeleteParams) WithDefaults() *DcimPlatformsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim platforms bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim platforms bulk delete params
func (o *DcimPlatformsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPlatformsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms bulk delete params
func (o *DcimPlatformsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms bulk delete params
func (o *DcimPlatformsBulkDeleteParams) WithContext(ctx context.Context) *DcimPlatformsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms bulk delete params
func (o *DcimPlatformsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms bulk delete params
func (o *DcimPlatformsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPlatformsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms bulk delete params
func (o *DcimPlatformsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
