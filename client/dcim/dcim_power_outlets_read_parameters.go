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

// NewDcimPowerOutletsReadParams creates a new DcimPowerOutletsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletsReadParams() *DcimPowerOutletsReadParams {
	return &DcimPowerOutletsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsReadParamsWithTimeout creates a new DcimPowerOutletsReadParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletsReadParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsReadParams {
	return &DcimPowerOutletsReadParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletsReadParamsWithContext creates a new DcimPowerOutletsReadParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletsReadParamsWithContext(ctx context.Context) *DcimPowerOutletsReadParams {
	return &DcimPowerOutletsReadParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletsReadParamsWithHTTPClient creates a new DcimPowerOutletsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletsReadParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsReadParams {
	return &DcimPowerOutletsReadParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletsReadParams contains all the parameters to send to the API endpoint
   for the dcim power outlets read operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletsReadParams struct {

	/* ID.

	   A UUID string identifying this power outlet.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power outlets read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsReadParams) WithDefaults() *DcimPowerOutletsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlets read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) WithContext(ctx context.Context) *DcimPowerOutletsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) WithID(id strfmt.UUID) *DcimPowerOutletsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets read params
func (o *DcimPowerOutletsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
