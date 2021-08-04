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

// NewDcimPowerPanelsReadParams creates a new DcimPowerPanelsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPanelsReadParams() *DcimPowerPanelsReadParams {
	return &DcimPowerPanelsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPanelsReadParamsWithTimeout creates a new DcimPowerPanelsReadParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPanelsReadParamsWithTimeout(timeout time.Duration) *DcimPowerPanelsReadParams {
	return &DcimPowerPanelsReadParams{
		timeout: timeout,
	}
}

// NewDcimPowerPanelsReadParamsWithContext creates a new DcimPowerPanelsReadParams object
// with the ability to set a context for a request.
func NewDcimPowerPanelsReadParamsWithContext(ctx context.Context) *DcimPowerPanelsReadParams {
	return &DcimPowerPanelsReadParams{
		Context: ctx,
	}
}

// NewDcimPowerPanelsReadParamsWithHTTPClient creates a new DcimPowerPanelsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPanelsReadParamsWithHTTPClient(client *http.Client) *DcimPowerPanelsReadParams {
	return &DcimPowerPanelsReadParams{
		HTTPClient: client,
	}
}

/* DcimPowerPanelsReadParams contains all the parameters to send to the API endpoint
   for the dcim power panels read operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPanelsReadParams struct {

	/* ID.

	   A UUID string identifying this power panel.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power panels read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsReadParams) WithDefaults() *DcimPowerPanelsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power panels read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPanelsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) WithTimeout(timeout time.Duration) *DcimPowerPanelsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) WithContext(ctx context.Context) *DcimPowerPanelsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) WithHTTPClient(client *http.Client) *DcimPowerPanelsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) WithID(id strfmt.UUID) *DcimPowerPanelsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power panels read params
func (o *DcimPowerPanelsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPanelsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
