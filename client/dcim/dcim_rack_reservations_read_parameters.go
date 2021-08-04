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

// NewDcimRackReservationsReadParams creates a new DcimRackReservationsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackReservationsReadParams() *DcimRackReservationsReadParams {
	return &DcimRackReservationsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsReadParamsWithTimeout creates a new DcimRackReservationsReadParams object
// with the ability to set a timeout on a request.
func NewDcimRackReservationsReadParamsWithTimeout(timeout time.Duration) *DcimRackReservationsReadParams {
	return &DcimRackReservationsReadParams{
		timeout: timeout,
	}
}

// NewDcimRackReservationsReadParamsWithContext creates a new DcimRackReservationsReadParams object
// with the ability to set a context for a request.
func NewDcimRackReservationsReadParamsWithContext(ctx context.Context) *DcimRackReservationsReadParams {
	return &DcimRackReservationsReadParams{
		Context: ctx,
	}
}

// NewDcimRackReservationsReadParamsWithHTTPClient creates a new DcimRackReservationsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackReservationsReadParamsWithHTTPClient(client *http.Client) *DcimRackReservationsReadParams {
	return &DcimRackReservationsReadParams{
		HTTPClient: client,
	}
}

/* DcimRackReservationsReadParams contains all the parameters to send to the API endpoint
   for the dcim rack reservations read operation.

   Typically these are written to a http.Request.
*/
type DcimRackReservationsReadParams struct {

	/* ID.

	   A UUID string identifying this rack reservation.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack reservations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsReadParams) WithDefaults() *DcimRackReservationsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack reservations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithTimeout(timeout time.Duration) *DcimRackReservationsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithContext(ctx context.Context) *DcimRackReservationsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithHTTPClient(client *http.Client) *DcimRackReservationsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) WithID(id strfmt.UUID) *DcimRackReservationsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack reservations read params
func (o *DcimRackReservationsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
