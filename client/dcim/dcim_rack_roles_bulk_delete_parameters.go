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

// NewDcimRackRolesBulkDeleteParams creates a new DcimRackRolesBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackRolesBulkDeleteParams() *DcimRackRolesBulkDeleteParams {
	return &DcimRackRolesBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackRolesBulkDeleteParamsWithTimeout creates a new DcimRackRolesBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRackRolesBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimRackRolesBulkDeleteParams {
	return &DcimRackRolesBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRackRolesBulkDeleteParamsWithContext creates a new DcimRackRolesBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimRackRolesBulkDeleteParamsWithContext(ctx context.Context) *DcimRackRolesBulkDeleteParams {
	return &DcimRackRolesBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimRackRolesBulkDeleteParamsWithHTTPClient creates a new DcimRackRolesBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackRolesBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimRackRolesBulkDeleteParams {
	return &DcimRackRolesBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRackRolesBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rack roles bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimRackRolesBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack roles bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesBulkDeleteParams) WithDefaults() *DcimRackRolesBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack roles bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackRolesBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack roles bulk delete params
func (o *DcimRackRolesBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimRackRolesBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack roles bulk delete params
func (o *DcimRackRolesBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack roles bulk delete params
func (o *DcimRackRolesBulkDeleteParams) WithContext(ctx context.Context) *DcimRackRolesBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack roles bulk delete params
func (o *DcimRackRolesBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack roles bulk delete params
func (o *DcimRackRolesBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimRackRolesBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack roles bulk delete params
func (o *DcimRackRolesBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackRolesBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
