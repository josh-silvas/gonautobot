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

// NewDcimInterfaceTemplatesBulkDeleteParams creates a new DcimInterfaceTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfaceTemplatesBulkDeleteParams() *DcimInterfaceTemplatesBulkDeleteParams {
	return &DcimInterfaceTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesBulkDeleteParamsWithTimeout creates a new DcimInterfaceTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimInterfaceTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesBulkDeleteParams {
	return &DcimInterfaceTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesBulkDeleteParamsWithContext creates a new DcimInterfaceTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimInterfaceTemplatesBulkDeleteParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesBulkDeleteParams {
	return &DcimInterfaceTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesBulkDeleteParamsWithHTTPClient creates a new DcimInterfaceTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfaceTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesBulkDeleteParams {
	return &DcimInterfaceTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimInterfaceTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim interface templates bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimInterfaceTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interface templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesBulkDeleteParams) WithDefaults() *DcimInterfaceTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interface templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interface templates bulk delete params
func (o *DcimInterfaceTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates bulk delete params
func (o *DcimInterfaceTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates bulk delete params
func (o *DcimInterfaceTemplatesBulkDeleteParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates bulk delete params
func (o *DcimInterfaceTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates bulk delete params
func (o *DcimInterfaceTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates bulk delete params
func (o *DcimInterfaceTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
