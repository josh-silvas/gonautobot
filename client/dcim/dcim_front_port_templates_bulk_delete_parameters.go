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

// NewDcimFrontPortTemplatesBulkDeleteParams creates a new DcimFrontPortTemplatesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesBulkDeleteParams() *DcimFrontPortTemplatesBulkDeleteParams {
	return &DcimFrontPortTemplatesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesBulkDeleteParamsWithTimeout creates a new DcimFrontPortTemplatesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesBulkDeleteParams {
	return &DcimFrontPortTemplatesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesBulkDeleteParamsWithContext creates a new DcimFrontPortTemplatesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesBulkDeleteParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesBulkDeleteParams {
	return &DcimFrontPortTemplatesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesBulkDeleteParamsWithHTTPClient creates a new DcimFrontPortTemplatesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesBulkDeleteParams {
	return &DcimFrontPortTemplatesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortTemplatesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim front port templates bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesBulkDeleteParams) WithDefaults() *DcimFrontPortTemplatesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates bulk delete params
func (o *DcimFrontPortTemplatesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates bulk delete params
func (o *DcimFrontPortTemplatesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates bulk delete params
func (o *DcimFrontPortTemplatesBulkDeleteParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates bulk delete params
func (o *DcimFrontPortTemplatesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates bulk delete params
func (o *DcimFrontPortTemplatesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates bulk delete params
func (o *DcimFrontPortTemplatesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
