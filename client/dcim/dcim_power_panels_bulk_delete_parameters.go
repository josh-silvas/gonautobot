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

// NewDcimPowerPanelsBulkDeleteParams creates a new DcimPowerPanelsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPanelsBulkDeleteParams() *DcimPowerPanelsBulkDeleteParams {
	return &DcimPowerPanelsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsBulkDeleteParamsWithTimeout creates a new DcimPowerPanelsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPanelsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsBulkDeleteParams {
	return &DcimPowerPanelsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerPanelsBulkDeleteParamsWithContext creates a new DcimPowerPanelsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerPanelsBulkDeleteParamsWithContext(ctx context.Context) *DcimPowerPanelsBulkDeleteParams {
	return &DcimPowerPanelsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerPanelsBulkDeleteParamsWithHTTPClient creates a new DcimPowerPanelsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPanelsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsBulkDeleteParams {
	return &DcimPowerPanelsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerPanelsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power panels bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPanelsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power panels bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsBulkDeleteParams) WithDefaults() *DcimPowerPanelsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power panels bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) WithContext(ctx context.Context) *DcimPowerPanelsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels bulk delete params
func (o *DcimPowerPanelsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
