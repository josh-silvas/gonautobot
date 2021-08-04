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

// NewDcimPowerOutletTemplatesDeleteParams creates a new DcimPowerOutletTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletTemplatesDeleteParams() *DcimPowerOutletTemplatesDeleteParams {
	return &DcimPowerOutletTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesDeleteParamsWithTimeout creates a new DcimPowerOutletTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesDeleteParams {
	return &DcimPowerOutletTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesDeleteParamsWithContext creates a new DcimPowerOutletTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletTemplatesDeleteParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesDeleteParams {
	return &DcimPowerOutletTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesDeleteParamsWithHTTPClient creates a new DcimPowerOutletTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesDeleteParams {
	return &DcimPowerOutletTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim power outlet templates delete operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletTemplatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this power outlet template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlet templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesDeleteParams) WithDefaults() *DcimPowerOutletTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlet templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) WithID(id strfmt.UUID) *DcimPowerOutletTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlet templates delete params
func (o *DcimPowerOutletTemplatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
