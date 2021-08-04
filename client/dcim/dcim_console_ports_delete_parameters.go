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

// NewDcimConsolePortsDeleteParams creates a new DcimConsolePortsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsolePortsDeleteParams() *DcimConsolePortsDeleteParams {
	return &DcimConsolePortsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsDeleteParamsWithTimeout creates a new DcimConsolePortsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimConsolePortsDeleteParamsWithTimeout(timeout time.Duration) *DcimConsolePortsDeleteParams {
	return &DcimConsolePortsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimConsolePortsDeleteParamsWithContext creates a new DcimConsolePortsDeleteParams object
// with the ability to set a context for a request.
func NewDcimConsolePortsDeleteParamsWithContext(ctx context.Context) *DcimConsolePortsDeleteParams {
	return &DcimConsolePortsDeleteParams{
		Context: ctx,
	}
}

// NewDcimConsolePortsDeleteParamsWithHTTPClient creates a new DcimConsolePortsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsolePortsDeleteParamsWithHTTPClient(client *http.Client) *DcimConsolePortsDeleteParams {
	return &DcimConsolePortsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimConsolePortsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim console ports delete operation.

   Typically these are written to a http.Request.
*/
type DcimConsolePortsDeleteParams struct {

	/* ID.

	   A UUID string identifying this console port.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim console ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsDeleteParams) WithDefaults() *DcimConsolePortsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console ports delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsolePortsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) WithTimeout(timeout time.Duration) *DcimConsolePortsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) WithContext(ctx context.Context) *DcimConsolePortsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) WithHTTPClient(client *http.Client) *DcimConsolePortsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) WithID(id strfmt.UUID) *DcimConsolePortsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports delete params
func (o *DcimConsolePortsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
