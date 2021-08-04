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

// NewDcimManufacturersReadParams creates a new DcimManufacturersReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimManufacturersReadParams() *DcimManufacturersReadParams {
	return &DcimManufacturersReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersReadParamsWithTimeout creates a new DcimManufacturersReadParams object
// with the ability to set a timeout on a request.
func NewDcimManufacturersReadParamsWithTimeout(timeout time.Duration) *DcimManufacturersReadParams {
	return &DcimManufacturersReadParams{
		timeout: timeout,
	}
}

// NewDcimManufacturersReadParamsWithContext creates a new DcimManufacturersReadParams object
// with the ability to set a context for a request.
func NewDcimManufacturersReadParamsWithContext(ctx context.Context) *DcimManufacturersReadParams {
	return &DcimManufacturersReadParams{
		Context: ctx,
	}
}

// NewDcimManufacturersReadParamsWithHTTPClient creates a new DcimManufacturersReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimManufacturersReadParamsWithHTTPClient(client *http.Client) *DcimManufacturersReadParams {
	return &DcimManufacturersReadParams{
		HTTPClient: client,
	}
}

/* DcimManufacturersReadParams contains all the parameters to send to the API endpoint
   for the dcim manufacturers read operation.

   Typically these are written to a http.Request.
*/
type DcimManufacturersReadParams struct {

	/* ID.

	   A UUID string identifying this manufacturer.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim manufacturers read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersReadParams) WithDefaults() *DcimManufacturersReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim manufacturers read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithTimeout(timeout time.Duration) *DcimManufacturersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithContext(ctx context.Context) *DcimManufacturersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithHTTPClient(client *http.Client) *DcimManufacturersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) WithID(id strfmt.UUID) *DcimManufacturersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim manufacturers read params
func (o *DcimManufacturersReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
