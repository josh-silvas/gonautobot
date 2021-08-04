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

// NewDcimRackReservationsBulkDeleteParams creates a new DcimRackReservationsBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackReservationsBulkDeleteParams() *DcimRackReservationsBulkDeleteParams {
	return &DcimRackReservationsBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsBulkDeleteParamsWithTimeout creates a new DcimRackReservationsBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRackReservationsBulkDeleteParamsWithTimeout(timeout time.Duration) *DcimRackReservationsBulkDeleteParams {
	return &DcimRackReservationsBulkDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRackReservationsBulkDeleteParamsWithContext creates a new DcimRackReservationsBulkDeleteParams object
// with the ability to set a context for a request.
func NewDcimRackReservationsBulkDeleteParamsWithContext(ctx context.Context) *DcimRackReservationsBulkDeleteParams {
	return &DcimRackReservationsBulkDeleteParams{
		Context: ctx,
	}
}

// NewDcimRackReservationsBulkDeleteParamsWithHTTPClient creates a new DcimRackReservationsBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackReservationsBulkDeleteParamsWithHTTPClient(client *http.Client) *DcimRackReservationsBulkDeleteParams {
	return &DcimRackReservationsBulkDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRackReservationsBulkDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rack reservations bulk delete operation.

   Typically these are written to a http.Request.
*/
type DcimRackReservationsBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack reservations bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsBulkDeleteParams) WithDefaults() *DcimRackReservationsBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack reservations bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack reservations bulk delete params
func (o *DcimRackReservationsBulkDeleteParams) WithTimeout(timeout time.Duration) *DcimRackReservationsBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations bulk delete params
func (o *DcimRackReservationsBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations bulk delete params
func (o *DcimRackReservationsBulkDeleteParams) WithContext(ctx context.Context) *DcimRackReservationsBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations bulk delete params
func (o *DcimRackReservationsBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations bulk delete params
func (o *DcimRackReservationsBulkDeleteParams) WithHTTPClient(client *http.Client) *DcimRackReservationsBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations bulk delete params
func (o *DcimRackReservationsBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
