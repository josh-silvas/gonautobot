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

// NewDcimInterfaceTemplatesDeleteParams creates a new DcimInterfaceTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfaceTemplatesDeleteParams() *DcimInterfaceTemplatesDeleteParams {
	return &DcimInterfaceTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceTemplatesDeleteParamsWithTimeout creates a new DcimInterfaceTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimInterfaceTemplatesDeleteParamsWithTimeout(timeout time.Duration) *DcimInterfaceTemplatesDeleteParams {
	return &DcimInterfaceTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimInterfaceTemplatesDeleteParamsWithContext creates a new DcimInterfaceTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewDcimInterfaceTemplatesDeleteParamsWithContext(ctx context.Context) *DcimInterfaceTemplatesDeleteParams {
	return &DcimInterfaceTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewDcimInterfaceTemplatesDeleteParamsWithHTTPClient creates a new DcimInterfaceTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfaceTemplatesDeleteParamsWithHTTPClient(client *http.Client) *DcimInterfaceTemplatesDeleteParams {
	return &DcimInterfaceTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimInterfaceTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim interface templates delete operation.

   Typically these are written to a http.Request.
*/
type DcimInterfaceTemplatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this interface template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interface templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesDeleteParams) WithDefaults() *DcimInterfaceTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interface templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) WithTimeout(timeout time.Duration) *DcimInterfaceTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) WithContext(ctx context.Context) *DcimInterfaceTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) WithHTTPClient(client *http.Client) *DcimInterfaceTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) WithID(id strfmt.UUID) *DcimInterfaceTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interface templates delete params
func (o *DcimInterfaceTemplatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
