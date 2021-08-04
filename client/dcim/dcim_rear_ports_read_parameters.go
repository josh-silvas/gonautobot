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

// NewDcimRearPortsReadParams creates a new DcimRearPortsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsReadParams() *DcimRearPortsReadParams {
	return &DcimRearPortsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsReadParamsWithTimeout creates a new DcimRearPortsReadParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsReadParamsWithTimeout(timeout time.Duration) *DcimRearPortsReadParams {
	return &DcimRearPortsReadParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsReadParamsWithContext creates a new DcimRearPortsReadParams object
// with the ability to set a context for a request.
func NewDcimRearPortsReadParamsWithContext(ctx context.Context) *DcimRearPortsReadParams {
	return &DcimRearPortsReadParams{
		Context: ctx,
	}
}

// NewDcimRearPortsReadParamsWithHTTPClient creates a new DcimRearPortsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsReadParamsWithHTTPClient(client *http.Client) *DcimRearPortsReadParams {
	return &DcimRearPortsReadParams{
		HTTPClient: client,
	}
}

/* DcimRearPortsReadParams contains all the parameters to send to the API endpoint
   for the dcim rear ports read operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortsReadParams struct {

	/* ID.

	   A UUID string identifying this rear port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rear ports read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsReadParams) WithDefaults() *DcimRearPortsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports read params
func (o *DcimRearPortsReadParams) WithTimeout(timeout time.Duration) *DcimRearPortsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports read params
func (o *DcimRearPortsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports read params
func (o *DcimRearPortsReadParams) WithContext(ctx context.Context) *DcimRearPortsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports read params
func (o *DcimRearPortsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports read params
func (o *DcimRearPortsReadParams) WithHTTPClient(client *http.Client) *DcimRearPortsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports read params
func (o *DcimRearPortsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim rear ports read params
func (o *DcimRearPortsReadParams) WithID(id strfmt.UUID) *DcimRearPortsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear ports read params
func (o *DcimRearPortsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
