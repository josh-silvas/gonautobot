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

// NewDcimDevicesReadParams creates a new DcimDevicesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesReadParams() *DcimDevicesReadParams {
	return &DcimDevicesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesReadParamsWithTimeout creates a new DcimDevicesReadParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesReadParamsWithTimeout(timeout time.Duration) *DcimDevicesReadParams {
	return &DcimDevicesReadParams{
		timeout: timeout,
	}
}

// NewDcimDevicesReadParamsWithContext creates a new DcimDevicesReadParams object
// with the ability to set a context for a request.
func NewDcimDevicesReadParamsWithContext(ctx context.Context) *DcimDevicesReadParams {
	return &DcimDevicesReadParams{
		Context: ctx,
	}
}

// NewDcimDevicesReadParamsWithHTTPClient creates a new DcimDevicesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesReadParamsWithHTTPClient(client *http.Client) *DcimDevicesReadParams {
	return &DcimDevicesReadParams{
		HTTPClient: client,
	}
}

/* DcimDevicesReadParams contains all the parameters to send to the API endpoint
   for the dcim devices read operation.

   Typically these are written to a http.Request.
*/
type DcimDevicesReadParams struct {

	/* ID.

	   A UUID string identifying this device.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesReadParams) WithDefaults() *DcimDevicesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices read params
func (o *DcimDevicesReadParams) WithTimeout(timeout time.Duration) *DcimDevicesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices read params
func (o *DcimDevicesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices read params
func (o *DcimDevicesReadParams) WithContext(ctx context.Context) *DcimDevicesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices read params
func (o *DcimDevicesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices read params
func (o *DcimDevicesReadParams) WithHTTPClient(client *http.Client) *DcimDevicesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices read params
func (o *DcimDevicesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim devices read params
func (o *DcimDevicesReadParams) WithID(id strfmt.UUID) *DcimDevicesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices read params
func (o *DcimDevicesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
