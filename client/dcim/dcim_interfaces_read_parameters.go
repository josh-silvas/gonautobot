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

// NewDcimInterfacesReadParams creates a new DcimInterfacesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfacesReadParams() *DcimInterfacesReadParams {
	return &DcimInterfacesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesReadParamsWithTimeout creates a new DcimInterfacesReadParams object
// with the ability to set a timeout on a request.
func NewDcimInterfacesReadParamsWithTimeout(timeout time.Duration) *DcimInterfacesReadParams {
	return &DcimInterfacesReadParams{
		timeout: timeout,
	}
}

// NewDcimInterfacesReadParamsWithContext creates a new DcimInterfacesReadParams object
// with the ability to set a context for a request.
func NewDcimInterfacesReadParamsWithContext(ctx context.Context) *DcimInterfacesReadParams {
	return &DcimInterfacesReadParams{
		Context: ctx,
	}
}

// NewDcimInterfacesReadParamsWithHTTPClient creates a new DcimInterfacesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfacesReadParamsWithHTTPClient(client *http.Client) *DcimInterfacesReadParams {
	return &DcimInterfacesReadParams{
		HTTPClient: client,
	}
}

/* DcimInterfacesReadParams contains all the parameters to send to the API endpoint
   for the dcim interfaces read operation.

   Typically these are written to a http.Request.
*/
type DcimInterfacesReadParams struct {

	/* ID.

	   A UUID string identifying this interface.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interfaces read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesReadParams) WithDefaults() *DcimInterfacesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interfaces read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interfaces read params
func (o *DcimInterfacesReadParams) WithTimeout(timeout time.Duration) *DcimInterfacesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces read params
func (o *DcimInterfacesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces read params
func (o *DcimInterfacesReadParams) WithContext(ctx context.Context) *DcimInterfacesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces read params
func (o *DcimInterfacesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces read params
func (o *DcimInterfacesReadParams) WithHTTPClient(client *http.Client) *DcimInterfacesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces read params
func (o *DcimInterfacesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim interfaces read params
func (o *DcimInterfacesReadParams) WithID(id strfmt.UUID) *DcimInterfacesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interfaces read params
func (o *DcimInterfacesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
