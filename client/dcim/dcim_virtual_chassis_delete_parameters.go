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

// NewDcimVirtualChassisDeleteParams creates a new DcimVirtualChassisDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimVirtualChassisDeleteParams() *DcimVirtualChassisDeleteParams {
	return &DcimVirtualChassisDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisDeleteParamsWithTimeout creates a new DcimVirtualChassisDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimVirtualChassisDeleteParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisDeleteParams {
	return &DcimVirtualChassisDeleteParams{
		timeout: timeout,
	}
}

// NewDcimVirtualChassisDeleteParamsWithContext creates a new DcimVirtualChassisDeleteParams object
// with the ability to set a context for a request.
func NewDcimVirtualChassisDeleteParamsWithContext(ctx context.Context) *DcimVirtualChassisDeleteParams {
	return &DcimVirtualChassisDeleteParams{
		Context: ctx,
	}
}

// NewDcimVirtualChassisDeleteParamsWithHTTPClient creates a new DcimVirtualChassisDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimVirtualChassisDeleteParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisDeleteParams {
	return &DcimVirtualChassisDeleteParams{
		HTTPClient: client,
	}
}

/* DcimVirtualChassisDeleteParams contains all the parameters to send to the API endpoint
   for the dcim virtual chassis delete operation.

   Typically these are written to a http.Request.
*/
type DcimVirtualChassisDeleteParams struct {

	/* ID.

	   A UUID string identifying this virtual chassis.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim virtual chassis delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisDeleteParams) WithDefaults() *DcimVirtualChassisDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim virtual chassis delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimVirtualChassisDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithContext(ctx context.Context) *DcimVirtualChassisDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithID(id strfmt.UUID) *DcimVirtualChassisDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
