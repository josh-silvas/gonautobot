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

// NewDcimPowerPortsTraceParams creates a new DcimPowerPortsTraceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortsTraceParams() *DcimPowerPortsTraceParams {
	return &DcimPowerPortsTraceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsTraceParamsWithTimeout creates a new DcimPowerPortsTraceParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortsTraceParamsWithTimeout(timeout time.Duration) *DcimPowerPortsTraceParams {
	return &DcimPowerPortsTraceParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortsTraceParamsWithContext creates a new DcimPowerPortsTraceParams object
// with the ability to set a context for a request.
func NewDcimPowerPortsTraceParamsWithContext(ctx context.Context) *DcimPowerPortsTraceParams {
	return &DcimPowerPortsTraceParams{
		Context: ctx,
	}
}

// NewDcimPowerPortsTraceParamsWithHTTPClient creates a new DcimPowerPortsTraceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortsTraceParamsWithHTTPClient(client *http.Client) *DcimPowerPortsTraceParams {
	return &DcimPowerPortsTraceParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortsTraceParams contains all the parameters to send to the API endpoint
   for the dcim power ports trace operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortsTraceParams struct {

	/* ID.

	   A UUID string identifying this power port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power ports trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsTraceParams) WithDefaults() *DcimPowerPortsTraceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power ports trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsTraceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) WithTimeout(timeout time.Duration) *DcimPowerPortsTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) WithContext(ctx context.Context) *DcimPowerPortsTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) WithHTTPClient(client *http.Client) *DcimPowerPortsTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) WithID(id strfmt.UUID) *DcimPowerPortsTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power ports trace params
func (o *DcimPowerPortsTraceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
