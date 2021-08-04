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

// NewDcimManufacturersDeleteParams creates a new DcimManufacturersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimManufacturersDeleteParams() *DcimManufacturersDeleteParams {
	return &DcimManufacturersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersDeleteParamsWithTimeout creates a new DcimManufacturersDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimManufacturersDeleteParamsWithTimeout(timeout time.Duration) *DcimManufacturersDeleteParams {
	return &DcimManufacturersDeleteParams{
		timeout: timeout,
	}
}

// NewDcimManufacturersDeleteParamsWithContext creates a new DcimManufacturersDeleteParams object
// with the ability to set a context for a request.
func NewDcimManufacturersDeleteParamsWithContext(ctx context.Context) *DcimManufacturersDeleteParams {
	return &DcimManufacturersDeleteParams{
		Context: ctx,
	}
}

// NewDcimManufacturersDeleteParamsWithHTTPClient creates a new DcimManufacturersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimManufacturersDeleteParamsWithHTTPClient(client *http.Client) *DcimManufacturersDeleteParams {
	return &DcimManufacturersDeleteParams{
		HTTPClient: client,
	}
}

/* DcimManufacturersDeleteParams contains all the parameters to send to the API endpoint
   for the dcim manufacturers delete operation.

   Typically these are written to a http.Request.
*/
type DcimManufacturersDeleteParams struct {

	/* ID.

	   A UUID string identifying this manufacturer.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim manufacturers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersDeleteParams) WithDefaults() *DcimManufacturersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim manufacturers delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) WithTimeout(timeout time.Duration) *DcimManufacturersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) WithContext(ctx context.Context) *DcimManufacturersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) WithHTTPClient(client *http.Client) *DcimManufacturersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) WithID(id strfmt.UUID) *DcimManufacturersDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers delete params
func (o *DcimManufacturersDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
