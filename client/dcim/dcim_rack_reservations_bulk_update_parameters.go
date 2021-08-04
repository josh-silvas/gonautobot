package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewDcimRackReservationsBulkUpdateParams creates a new DcimRackReservationsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackReservationsBulkUpdateParams() *DcimRackReservationsBulkUpdateParams {
	return &DcimRackReservationsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsBulkUpdateParamsWithTimeout creates a new DcimRackReservationsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimRackReservationsBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimRackReservationsBulkUpdateParams {
	return &DcimRackReservationsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimRackReservationsBulkUpdateParamsWithContext creates a new DcimRackReservationsBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimRackReservationsBulkUpdateParamsWithContext(ctx context.Context) *DcimRackReservationsBulkUpdateParams {
	return &DcimRackReservationsBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimRackReservationsBulkUpdateParamsWithHTTPClient creates a new DcimRackReservationsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackReservationsBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimRackReservationsBulkUpdateParams {
	return &DcimRackReservationsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* DcimRackReservationsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the dcim rack reservations bulk update operation.

   Typically these are written to a http.Request.
*/
type DcimRackReservationsBulkUpdateParams struct {

	// Data.
	Data *models.WritableRackReservation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack reservations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsBulkUpdateParams) WithDefaults() *DcimRackReservationsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack reservations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimRackReservationsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) WithContext(ctx context.Context) *DcimRackReservationsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimRackReservationsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) WithData(data *models.WritableRackReservation) *DcimRackReservationsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim rack reservations bulk update params
func (o *DcimRackReservationsBulkUpdateParams) SetData(data *models.WritableRackReservation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
