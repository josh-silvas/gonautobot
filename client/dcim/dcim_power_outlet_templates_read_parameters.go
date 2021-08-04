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

// NewDcimPowerOutletTemplatesReadParams creates a new DcimPowerOutletTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletTemplatesReadParams() *DcimPowerOutletTemplatesReadParams {
	return &DcimPowerOutletTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesReadParamsWithTimeout creates a new DcimPowerOutletTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesReadParams {
	return &DcimPowerOutletTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesReadParamsWithContext creates a new DcimPowerOutletTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletTemplatesReadParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesReadParams {
	return &DcimPowerOutletTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesReadParamsWithHTTPClient creates a new DcimPowerOutletTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesReadParams {
	return &DcimPowerOutletTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim power outlet templates read operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this power outlet template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlet templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesReadParams) WithDefaults() *DcimPowerOutletTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlet templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithID(id strfmt.UUID) *DcimPowerOutletTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
