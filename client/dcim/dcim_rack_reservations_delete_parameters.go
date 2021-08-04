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

// NewDcimRackReservationsDeleteParams creates a new DcimRackReservationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackReservationsDeleteParams() *DcimRackReservationsDeleteParams {
	return &DcimRackReservationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsDeleteParamsWithTimeout creates a new DcimRackReservationsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRackReservationsDeleteParamsWithTimeout(timeout time.Duration) *DcimRackReservationsDeleteParams {
	return &DcimRackReservationsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRackReservationsDeleteParamsWithContext creates a new DcimRackReservationsDeleteParams object
// with the ability to set a context for a request.
func NewDcimRackReservationsDeleteParamsWithContext(ctx context.Context) *DcimRackReservationsDeleteParams {
	return &DcimRackReservationsDeleteParams{
		Context: ctx,
	}
}

// NewDcimRackReservationsDeleteParamsWithHTTPClient creates a new DcimRackReservationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackReservationsDeleteParamsWithHTTPClient(client *http.Client) *DcimRackReservationsDeleteParams {
	return &DcimRackReservationsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRackReservationsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim rack reservations delete operation.

   Typically these are written to a http.Request.
*/
type DcimRackReservationsDeleteParams struct {

	/* ID.

	   A UUID string identifying this rack reservation.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack reservations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsDeleteParams) WithDefaults() *DcimRackReservationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack reservations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) WithTimeout(timeout time.Duration) *DcimRackReservationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) WithContext(ctx context.Context) *DcimRackReservationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) WithHTTPClient(client *http.Client) *DcimRackReservationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) WithID(id strfmt.UUID) *DcimRackReservationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack reservations delete params
func (o *DcimRackReservationsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
