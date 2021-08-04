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

// NewDcimPlatformsReadParams creates a new DcimPlatformsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPlatformsReadParams() *DcimPlatformsReadParams {
	return &DcimPlatformsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsReadParamsWithTimeout creates a new DcimPlatformsReadParams object
// with the ability to set a timeout on a request.
func NewDcimPlatformsReadParamsWithTimeout(timeout time.Duration) *DcimPlatformsReadParams {
	return &DcimPlatformsReadParams{
		timeout: timeout,
	}
}

// NewDcimPlatformsReadParamsWithContext creates a new DcimPlatformsReadParams object
// with the ability to set a context for a request.
func NewDcimPlatformsReadParamsWithContext(ctx context.Context) *DcimPlatformsReadParams {
	return &DcimPlatformsReadParams{
		Context: ctx,
	}
}

// NewDcimPlatformsReadParamsWithHTTPClient creates a new DcimPlatformsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPlatformsReadParamsWithHTTPClient(client *http.Client) *DcimPlatformsReadParams {
	return &DcimPlatformsReadParams{
		HTTPClient: client,
	}
}

/* DcimPlatformsReadParams contains all the parameters to send to the API endpoint
   for the dcim platforms read operation.

   Typically these are written to a http.Request.
*/
type DcimPlatformsReadParams struct {

	/* ID.

	   A UUID string identifying this platform.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim platforms read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsReadParams) WithDefaults() *DcimPlatformsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim platforms read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithTimeout(timeout time.Duration) *DcimPlatformsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithContext(ctx context.Context) *DcimPlatformsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithHTTPClient(client *http.Client) *DcimPlatformsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim platforms read params
func (o *DcimPlatformsReadParams) WithID(id strfmt.UUID) *DcimPlatformsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim platforms read params
func (o *DcimPlatformsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
