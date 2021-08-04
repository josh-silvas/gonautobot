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

// NewDcimConsolePortsReadParams creates a new DcimConsolePortsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsReadParams() *DcimConsolePortsReadParams {
	return &DcimConsolePortsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsReadParamsWithTimeout creates a new DcimConsolePortsReadParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsReadParamsWithTimeout(timeout time.Duration) *DcimConsolePortsReadParams {
	return &DcimConsolePortsReadParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsReadParamsWithContext creates a new DcimConsolePortsReadParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsReadParamsWithContext(ctx context.Context) *DcimConsolePortsReadParams {
	return &DcimConsolePortsReadParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsReadParamsWithHTTPClient creates a new DcimConsolePortsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsReadParamsWithHTTPClient(client *http.Client) *DcimConsolePortsReadParams {
	return &DcimConsolePortsReadParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsReadParams contains all the parameters to send to the API endpoint
   for the dcim console ports read operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsReadParams struct {

	/* ID.

	   A UUID string identifying this console port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsReadParams) WithDefaults() *DcimConsolePortsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithTimeout(timeout time.Duration) *DcimConsolePortsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithContext(ctx context.Context) *DcimConsolePortsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithHTTPClient(client *http.Client) *DcimConsolePortsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithID(id strfmt.UUID) *DcimConsolePortsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
