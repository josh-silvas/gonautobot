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

// NewDcimRegionsReadParams creates a new DcimRegionsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRegionsReadParams() *DcimRegionsReadParams {
	return &DcimRegionsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRegionsReadParamsWithTimeout creates a new DcimRegionsReadParams object
// with the ability to set a timeout on a request.
func NewDcimRegionsReadParamsWithTimeout(timeout time.Duration) *DcimRegionsReadParams {
	return &DcimRegionsReadParams{
		timeout: timeout,
	}
}

// NewDcimRegionsReadParamsWithContext creates a new DcimRegionsReadParams object
// with the ability to set a context for a request.
func NewDcimRegionsReadParamsWithContext(ctx context.Context) *DcimRegionsReadParams {
	return &DcimRegionsReadParams{
		Context: ctx,
	}
}

// NewDcimRegionsReadParamsWithHTTPClient creates a new DcimRegionsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRegionsReadParamsWithHTTPClient(client *http.Client) *DcimRegionsReadParams {
	return &DcimRegionsReadParams{
		HTTPClient: client,
	}
}

/* DcimRegionsReadParams contains all the parameters to send to the API endpoint
   for the dcim regions read operation.

   Typically these are written to a http.Request.
*/
type DcimRegionsReadParams struct {

	/* ID.

	   A UUID string identifying this region.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim regions read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsReadParams) WithDefaults() *DcimRegionsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim regions read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRegionsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim regions read params
func (o *DcimRegionsReadParams) WithTimeout(timeout time.Duration) *DcimRegionsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim regions read params
func (o *DcimRegionsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim regions read params
func (o *DcimRegionsReadParams) WithContext(ctx context.Context) *DcimRegionsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim regions read params
func (o *DcimRegionsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim regions read params
func (o *DcimRegionsReadParams) WithHTTPClient(client *http.Client) *DcimRegionsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim regions read params
func (o *DcimRegionsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim regions read params
func (o *DcimRegionsReadParams) WithID(id strfmt.UUID) *DcimRegionsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim regions read params
func (o *DcimRegionsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRegionsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
