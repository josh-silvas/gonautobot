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

// NewDcimInterfacesTraceParams creates a new DcimInterfacesTraceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfacesTraceParams() *DcimInterfacesTraceParams {
	return &DcimInterfacesTraceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesTraceParamsWithTimeout creates a new DcimInterfacesTraceParams object
// with the ability to set a timeout on a request.
func NewDcimInterfacesTraceParamsWithTimeout(timeout time.Duration) *DcimInterfacesTraceParams {
	return &DcimInterfacesTraceParams{
		timeout: timeout,
	}
}

// NewDcimInterfacesTraceParamsWithContext creates a new DcimInterfacesTraceParams object
// with the ability to set a context for a request.
func NewDcimInterfacesTraceParamsWithContext(ctx context.Context) *DcimInterfacesTraceParams {
	return &DcimInterfacesTraceParams{
		Context: ctx,
	}
}

// NewDcimInterfacesTraceParamsWithHTTPClient creates a new DcimInterfacesTraceParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfacesTraceParamsWithHTTPClient(client *http.Client) *DcimInterfacesTraceParams {
	return &DcimInterfacesTraceParams{
		HTTPClient: client,
	}
}

/* DcimInterfacesTraceParams contains all the parameters to send to the API endpoint
   for the dcim interfaces trace operation.

   Typically these are written to a http.Request.
*/
type DcimInterfacesTraceParams struct {

	/* ID.

	   A UUID string identifying this interface.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interfaces trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesTraceParams) WithDefaults() *DcimInterfacesTraceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interfaces trace params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesTraceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) WithTimeout(timeout time.Duration) *DcimInterfacesTraceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) WithContext(ctx context.Context) *DcimInterfacesTraceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) WithHTTPClient(client *http.Client) *DcimInterfacesTraceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) WithID(id strfmt.UUID) *DcimInterfacesTraceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interfaces trace params
func (o *DcimInterfacesTraceParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesTraceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
