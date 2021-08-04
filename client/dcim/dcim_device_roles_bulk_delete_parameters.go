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

// NewDcimDeviceRolesBulkDeleteParams creates a new DcimDeviceRolesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceRolesBulkDeleteParams() *DcimDeviceRolesBulkDeleteParams {
	return &DcimDeviceRolesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesBulkDeleteParamsWithTimeout creates a new DcimDeviceRolesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceRolesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesBulkDeleteParams {
	return &DcimDeviceRolesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimDeviceRolesBulkDeleteParamsWithContext creates a new DcimDeviceRolesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimDeviceRolesBulkDeleteParamsWithContext(ctx context.Context) *DcimDeviceRolesBulkDeleteParams {
	return &DcimDeviceRolesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimDeviceRolesBulkDeleteParamsWithHTTPClient creates a new DcimDeviceRolesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceRolesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesBulkDeleteParams {
	return &DcimDeviceRolesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimDeviceRolesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim device roles bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceRolesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device roles bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesBulkDeleteParams) WithDefaults() *DcimDeviceRolesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device roles bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceRolesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device roles bulk delete params
func (o *DcimDeviceRolesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles bulk delete params
func (o *DcimDeviceRolesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles bulk delete params
func (o *DcimDeviceRolesBulkDeleteParams) WithContext(ctx context.Context) *DcimDeviceRolesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles bulk delete params
func (o *DcimDeviceRolesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles bulk delete params
func (o *DcimDeviceRolesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles bulk delete params
func (o *DcimDeviceRolesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
