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

// NewDcimPowerOutletTemplatesBulkDeleteParams creates a new DcimPowerOutletTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletTemplatesBulkDeleteParams() *DcimPowerOutletTemplatesBulkDeleteParams {
	return &DcimPowerOutletTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesBulkDeleteParamsWithTimeout creates a new DcimPowerOutletTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesBulkDeleteParams {
	return &DcimPowerOutletTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesBulkDeleteParamsWithContext creates a new DcimPowerOutletTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletTemplatesBulkDeleteParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesBulkDeleteParams {
	return &DcimPowerOutletTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesBulkDeleteParamsWithHTTPClient creates a new DcimPowerOutletTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesBulkDeleteParams {
	return &DcimPowerOutletTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power outlet templates bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlet templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesBulkDeleteParams) WithDefaults() *DcimPowerOutletTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlet templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlet templates bulk delete params
func (o *DcimPowerOutletTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates bulk delete params
func (o *DcimPowerOutletTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates bulk delete params
func (o *DcimPowerOutletTemplatesBulkDeleteParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates bulk delete params
func (o *DcimPowerOutletTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates bulk delete params
func (o *DcimPowerOutletTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates bulk delete params
func (o *DcimPowerOutletTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
