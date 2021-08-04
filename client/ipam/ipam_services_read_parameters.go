package ipam

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIpamServicesReadParams creates a new IpamServicesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesReadParams() *IpamServicesReadParams {
	return &IpamServicesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesReadParamsWithTimeout creates a new IpamServicesReadParams object
// with the ability to set a timeout on a request.
func NewIpamServicesReadParamsWithTimeout(timeout time.Duration) *IpamServicesReadParams {
	return &IpamServicesReadParams{
		timeout: timeout,
	}
}

// NewIpamServicesReadParamsWithContext creates a new IpamServicesReadParams object
// with the ability to set a context for a request.
func NewIpamServicesReadParamsWithContext(ctx context.Context) *IpamServicesReadParams {
	return &IpamServicesReadParams{
		Context: ctx,
	}
}

// NewIpamServicesReadParamsWithHTTPClient creates a new IpamServicesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesReadParamsWithHTTPClient(client *http.Client) *IpamServicesReadParams {
	return &IpamServicesReadParams{
		HTTPClient: client,
	}
}

/* IpamServicesReadParams contains all the parameters to send to the API endpoint
   for the ipam services read operation.

   Typically these are written to a http.Request.
*/
type IpamServicesReadParams struct {

	/* ID.

	   A UUID string identifying this service.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesReadParams) WithDefaults() *IpamServicesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services read params
func (o *IpamServicesReadParams) WithTimeout(timeout time.Duration) *IpamServicesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services read params
func (o *IpamServicesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services read params
func (o *IpamServicesReadParams) WithContext(ctx context.Context) *IpamServicesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services read params
func (o *IpamServicesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services read params
func (o *IpamServicesReadParams) WithHTTPClient(client *http.Client) *IpamServicesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services read params
func (o *IpamServicesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam services read params
func (o *IpamServicesReadParams) WithID(id strfmt.UUID) *IpamServicesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services read params
func (o *IpamServicesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
