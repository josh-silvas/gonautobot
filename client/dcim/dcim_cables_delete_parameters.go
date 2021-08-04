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

// NewDcimCablesDeleteParams creates a new DcimCablesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesDeleteParams() *DcimCablesDeleteParams {
	return &DcimCablesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesDeleteParamsWithTimeout creates a new DcimCablesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimCablesDeleteParamsWithTimeout(timeout time.Duration) *DcimCablesDeleteParams {
	return &DcimCablesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimCablesDeleteParamsWithContext creates a new DcimCablesDeleteParams object
// with the ability to set a context for a request.
func NewDcimCablesDeleteParamsWithContext(ctx context.Context) *DcimCablesDeleteParams {
	return &DcimCablesDeleteParams{
		Context: ctx,
	}
}

// NewDcimCablesDeleteParamsWithHTTPClient creates a new DcimCablesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesDeleteParamsWithHTTPClient(client *http.Client) *DcimCablesDeleteParams {
	return &DcimCablesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimCablesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim cables delete operation.

   Typically these are written to a http.Request.
*/
type DcimCablesDeleteParams struct {

	/* ID.

	   A UUID string identifying this cable.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesDeleteParams) WithDefaults() *DcimCablesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithTimeout(timeout time.Duration) *DcimCablesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithContext(ctx context.Context) *DcimCablesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithHTTPClient(client *http.Client) *DcimCablesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim cables delete params
func (o *DcimCablesDeleteParams) WithID(id strfmt.UUID) *DcimCablesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cables delete params
func (o *DcimCablesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
