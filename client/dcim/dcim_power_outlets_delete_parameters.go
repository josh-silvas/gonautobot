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

// NewDcimPowerOutletsDeleteParams creates a new DcimPowerOutletsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletsDeleteParams() *DcimPowerOutletsDeleteParams {
	return &DcimPowerOutletsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsDeleteParamsWithTimeout creates a new DcimPowerOutletsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletsDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsDeleteParams {
	return &DcimPowerOutletsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletsDeleteParamsWithContext creates a new DcimPowerOutletsDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletsDeleteParamsWithContext(ctx context.Context) *DcimPowerOutletsDeleteParams {
	return &DcimPowerOutletsDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletsDeleteParamsWithHTTPClient creates a new DcimPowerOutletsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletsDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsDeleteParams {
	return &DcimPowerOutletsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power outlets delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletsDeleteParams struct {

	/* ID.

	   A UUID string identifying this power outlet.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlets delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsDeleteParams) WithDefaults() *DcimPowerOutletsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlets delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) WithContext(ctx context.Context) *DcimPowerOutletsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) WithID(id strfmt.UUID) *DcimPowerOutletsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets delete params
func (o *DcimPowerOutletsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
