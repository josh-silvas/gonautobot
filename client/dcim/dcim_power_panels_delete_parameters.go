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

// NewDcimPowerPanelsDeleteParams creates a new DcimPowerPanelsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPanelsDeleteParams() *DcimPowerPanelsDeleteParams {
	return &DcimPowerPanelsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsDeleteParamsWithTimeout creates a new DcimPowerPanelsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPanelsDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsDeleteParams {
	return &DcimPowerPanelsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerPanelsDeleteParamsWithContext creates a new DcimPowerPanelsDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerPanelsDeleteParamsWithContext(ctx context.Context) *DcimPowerPanelsDeleteParams {
	return &DcimPowerPanelsDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerPanelsDeleteParamsWithHTTPClient creates a new DcimPowerPanelsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPanelsDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsDeleteParams {
	return &DcimPowerPanelsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerPanelsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power panels delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPanelsDeleteParams struct {

	/* ID.

	   A UUID string identifying this power panel.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power panels delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsDeleteParams) WithDefaults() *DcimPowerPanelsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power panels delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) WithContext(ctx context.Context) *DcimPowerPanelsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) WithID(id strfmt.UUID) *DcimPowerPanelsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power panels delete params
func (o *DcimPowerPanelsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
