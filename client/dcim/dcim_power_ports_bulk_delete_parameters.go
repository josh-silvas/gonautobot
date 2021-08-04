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

// NewDcimPowerPortsBulkDeleteParams creates a new DcimPowerPortsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortsBulkDeleteParams() *DcimPowerPortsBulkDeleteParams {
	return &DcimPowerPortsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsBulkDeleteParamsWithTimeout creates a new DcimPowerPortsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPortsBulkDeleteParams {
	return &DcimPowerPortsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortsBulkDeleteParamsWithContext creates a new DcimPowerPortsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerPortsBulkDeleteParamsWithContext(ctx context.Context) *DcimPowerPortsBulkDeleteParams {
	return &DcimPowerPortsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerPortsBulkDeleteParamsWithHTTPClient creates a new DcimPowerPortsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPortsBulkDeleteParams {
	return &DcimPowerPortsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power ports bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsBulkDeleteParams) WithDefaults() *DcimPowerPortsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power ports bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power ports bulk delete params
func (o *DcimPowerPortsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPortsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports bulk delete params
func (o *DcimPowerPortsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports bulk delete params
func (o *DcimPowerPortsBulkDeleteParams) WithContext(ctx context.Context) *DcimPowerPortsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports bulk delete params
func (o *DcimPowerPortsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports bulk delete params
func (o *DcimPowerPortsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPortsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports bulk delete params
func (o *DcimPowerPortsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
