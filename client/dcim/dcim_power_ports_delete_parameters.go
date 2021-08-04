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

// NewDcimPowerPortsDeleteParams creates a new DcimPowerPortsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortsDeleteParams() *DcimPowerPortsDeleteParams {
	return &DcimPowerPortsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsDeleteParamsWithTimeout creates a new DcimPowerPortsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortsDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPortsDeleteParams {
	return &DcimPowerPortsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortsDeleteParamsWithContext creates a new DcimPowerPortsDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerPortsDeleteParamsWithContext(ctx context.Context) *DcimPowerPortsDeleteParams {
	return &DcimPowerPortsDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerPortsDeleteParamsWithHTTPClient creates a new DcimPowerPortsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortsDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPortsDeleteParams {
	return &DcimPowerPortsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power ports delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortsDeleteParams struct {

	/* ID.

	   A UUID string identifying this power port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsDeleteParams) WithDefaults() *DcimPowerPortsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPortsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithContext(ctx context.Context) *DcimPowerPortsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPortsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) WithID(id strfmt.UUID) *DcimPowerPortsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power ports delete params
func (o *DcimPowerPortsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
