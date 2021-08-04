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

// NewDcimConsolePortTemplatesBulkDeleteParams creates a new DcimConsolePortTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortTemplatesBulkDeleteParams() *DcimConsolePortTemplatesBulkDeleteParams {
	return &DcimConsolePortTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortTemplatesBulkDeleteParamsWithTimeout creates a new DcimConsolePortTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimConsolePortTemplatesBulkDeleteParams {
	return &DcimConsolePortTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortTemplatesBulkDeleteParamsWithContext creates a new DcimConsolePortTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimConsolePortTemplatesBulkDeleteParamsWithContext(ctx context.Context) *DcimConsolePortTemplatesBulkDeleteParams {
	return &DcimConsolePortTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimConsolePortTemplatesBulkDeleteParamsWithHTTPClient creates a new DcimConsolePortTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimConsolePortTemplatesBulkDeleteParams {
	return &DcimConsolePortTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim console port templates bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesBulkDeleteParams) WithDefaults() *DcimConsolePortTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console port templates bulk delete params
func (o *DcimConsolePortTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimConsolePortTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console port templates bulk delete params
func (o *DcimConsolePortTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console port templates bulk delete params
func (o *DcimConsolePortTemplatesBulkDeleteParams) WithContext(ctx context.Context) *DcimConsolePortTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console port templates bulk delete params
func (o *DcimConsolePortTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console port templates bulk delete params
func (o *DcimConsolePortTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimConsolePortTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console port templates bulk delete params
func (o *DcimConsolePortTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
