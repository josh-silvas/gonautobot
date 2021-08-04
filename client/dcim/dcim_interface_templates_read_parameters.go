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

// NewDcimInterfaceTemplatesReadParams creates a new DcimInterfaceTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfaceTemplatesReadParams() *DcimInterfaceTemplatesReadParams {
	return &DcimInterfaceTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesReadParamsWithTimeout creates a new DcimInterfaceTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewDcimInterfaceTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesReadParams {
	return &DcimInterfaceTemplatesReadParams{
		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesReadParamsWithContext creates a new DcimInterfaceTemplatesReadParams object
// with the ability to set a context for a request.
func NewDcimInterfaceTemplatesReadParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesReadParams {
	return &DcimInterfaceTemplatesReadParams{
		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesReadParamsWithHTTPClient creates a new DcimInterfaceTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfaceTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesReadParams {
	return &DcimInterfaceTemplatesReadParams{
		HTTPClient: client,
	}
}

/* DcimInterfaceTemplatesReadParams contains all the parameters to send to the API endpoint
   for the dcim interface templates read operation.

   Typically these are written to a http.Request.
*/
type DcimInterfaceTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this interface template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interface templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesReadParams) WithDefaults() *DcimInterfaceTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interface templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) WithID(id strfmt.UUID) *DcimInterfaceTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interface templates read params
func (o *DcimInterfaceTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
