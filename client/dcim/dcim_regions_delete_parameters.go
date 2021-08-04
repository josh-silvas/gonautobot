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

// NewDcimRegionsDeleteParams creates a new DcimRegionsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRegionsDeleteParams() *DcimRegionsDeleteParams {
	return &DcimRegionsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRegionsDeleteParamsWithTimeout creates a new DcimRegionsDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimRegionsDeleteParamsWithTimeout(timeout time.Duration) *DcimRegionsDeleteParams {
	return &DcimRegionsDeleteParams{
		timeout: timeout,
	}
}

// NewDcimRegionsDeleteParamsWithContext creates a new DcimRegionsDeleteParams object
// with the ability to set a context for a request.
func NewDcimRegionsDeleteParamsWithContext(ctx context.Context) *DcimRegionsDeleteParams {
	return &DcimRegionsDeleteParams{
		Context: ctx,
	}
}

// NewDcimRegionsDeleteParamsWithHTTPClient creates a new DcimRegionsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRegionsDeleteParamsWithHTTPClient(client *http.Client) *DcimRegionsDeleteParams {
	return &DcimRegionsDeleteParams{
		HTTPClient: client,
	}
}

/* DcimRegionsDeleteParams contains all the parameters to send to the API endpoint
   for the dcim regions delete operation.

   Typically these are written to a http.Request.
*/
type DcimRegionsDeleteParams struct {

	/* ID.

	   A UUID string identifying this region.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim regions delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsDeleteParams) WithDefaults() *DcimRegionsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim regions delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim regions delete params
func (o *DcimRegionsDeleteParams) WithTimeout(timeout time.Duration) *DcimRegionsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim regions delete params
func (o *DcimRegionsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim regions delete params
func (o *DcimRegionsDeleteParams) WithContext(ctx context.Context) *DcimRegionsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim regions delete params
func (o *DcimRegionsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim regions delete params
func (o *DcimRegionsDeleteParams) WithHTTPClient(client *http.Client) *DcimRegionsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim regions delete params
func (o *DcimRegionsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim regions delete params
func (o *DcimRegionsDeleteParams) WithID(id strfmt.UUID) *DcimRegionsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim regions delete params
func (o *DcimRegionsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRegionsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
