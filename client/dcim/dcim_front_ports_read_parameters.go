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

// NewDcimFrontPortsReadParams creates a new DcimFrontPortsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortsReadParams() *DcimFrontPortsReadParams {
	return &DcimFrontPortsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortsReadParamsWithTimeout creates a new DcimFrontPortsReadParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortsReadParamsWithTimeout(timeout time.Duration) *DcimFrontPortsReadParams {
	return &DcimFrontPortsReadParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortsReadParamsWithContext creates a new DcimFrontPortsReadParams object
// with the ability to set a context for a request.
func NewDcimFrontPortsReadParamsWithContext(ctx context.Context) *DcimFrontPortsReadParams {
	return &DcimFrontPortsReadParams{
		Context: ctx,
	}
}

// NewDcimFrontPortsReadParamsWithHTTPClient creates a new DcimFrontPortsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortsReadParamsWithHTTPClient(client *http.Client) *DcimFrontPortsReadParams {
	return &DcimFrontPortsReadParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortsReadParams contains all the parameters to send to the API endpoint
   for the dcim front ports read operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortsReadParams struct {

	/* ID.

	   A UUID string identifying this front port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front ports read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsReadParams) WithDefaults() *DcimFrontPortsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front ports read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front ports read params
func (o *DcimFrontPortsReadParams) WithTimeout(timeout time.Duration) *DcimFrontPortsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front ports read params
func (o *DcimFrontPortsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front ports read params
func (o *DcimFrontPortsReadParams) WithContext(ctx context.Context) *DcimFrontPortsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front ports read params
func (o *DcimFrontPortsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front ports read params
func (o *DcimFrontPortsReadParams) WithHTTPClient(client *http.Client) *DcimFrontPortsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front ports read params
func (o *DcimFrontPortsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim front ports read params
func (o *DcimFrontPortsReadParams) WithID(id strfmt.UUID) *DcimFrontPortsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front ports read params
func (o *DcimFrontPortsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
