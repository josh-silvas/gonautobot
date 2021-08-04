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

// NewDcimRearPortTemplatesBulkDeleteParams creates a new DcimRearPortTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortTemplatesBulkDeleteParams() *DcimRearPortTemplatesBulkDeleteParams {
	return &DcimRearPortTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortTemplatesBulkDeleteParamsWithTimeout creates a new DcimRearPortTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimRearPortTemplatesBulkDeleteParams {
	return &DcimRearPortTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRearPortTemplatesBulkDeleteParamsWithContext creates a new DcimRearPortTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimRearPortTemplatesBulkDeleteParamsWithContext(ctx context.Context) *DcimRearPortTemplatesBulkDeleteParams {
	return &DcimRearPortTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimRearPortTemplatesBulkDeleteParamsWithHTTPClient creates a new DcimRearPortTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimRearPortTemplatesBulkDeleteParams {
	return &DcimRearPortTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRearPortTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rear port templates bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesBulkDeleteParams) WithDefaults() *DcimRearPortTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear port templates bulk delete params
func (o *DcimRearPortTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimRearPortTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear port templates bulk delete params
func (o *DcimRearPortTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear port templates bulk delete params
func (o *DcimRearPortTemplatesBulkDeleteParams) WithContext(ctx context.Context) *DcimRearPortTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear port templates bulk delete params
func (o *DcimRearPortTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear port templates bulk delete params
func (o *DcimRearPortTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimRearPortTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear port templates bulk delete params
func (o *DcimRearPortTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
