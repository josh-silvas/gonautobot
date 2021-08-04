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

// NewDcimPowerPortTemplatesBulkDeleteParams creates a new DcimPowerPortTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortTemplatesBulkDeleteParams() *DcimPowerPortTemplatesBulkDeleteParams {
	return &DcimPowerPortTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortTemplatesBulkDeleteParamsWithTimeout creates a new DcimPowerPortTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPortTemplatesBulkDeleteParams {
	return &DcimPowerPortTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortTemplatesBulkDeleteParamsWithContext creates a new DcimPowerPortTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerPortTemplatesBulkDeleteParamsWithContext(ctx context.Context) *DcimPowerPortTemplatesBulkDeleteParams {
	return &DcimPowerPortTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerPortTemplatesBulkDeleteParamsWithHTTPClient creates a new DcimPowerPortTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPortTemplatesBulkDeleteParams {
	return &DcimPowerPortTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power port templates bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesBulkDeleteParams) WithDefaults() *DcimPowerPortTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power port templates bulk delete params
func (o *DcimPowerPortTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPortTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power port templates bulk delete params
func (o *DcimPowerPortTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power port templates bulk delete params
func (o *DcimPowerPortTemplatesBulkDeleteParams) WithContext(ctx context.Context) *DcimPowerPortTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power port templates bulk delete params
func (o *DcimPowerPortTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power port templates bulk delete params
func (o *DcimPowerPortTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPortTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power port templates bulk delete params
func (o *DcimPowerPortTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
