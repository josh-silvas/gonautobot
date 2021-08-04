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

// NewIpamRirsReadParams creates a new IpamRirsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRirsReadParams() *IpamRirsReadParams {
	return &IpamRirsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsReadParamsWithTimeout creates a new IpamRirsReadParams object
// with the ability to set a timeout on a request.
func NewIpamRirsReadParamsWithTimeout(timeout time.Duration) *IpamRirsReadParams {
	return &IpamRirsReadParams{
		timeout: timeout,
	}
}

// NewIpamRirsReadParamsWithContext creates a new IpamRirsReadParams object
// with the ability to set a context for a request.
func NewIpamRirsReadParamsWithContext(ctx context.Context) *IpamRirsReadParams {
	return &IpamRirsReadParams{
		Context: ctx,
	}
}

// NewIpamRirsReadParamsWithHTTPClient creates a new IpamRirsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRirsReadParamsWithHTTPClient(client *http.Client) *IpamRirsReadParams {
	return &IpamRirsReadParams{
		HTTPClient: client,
	}
}

/* IpamRirsReadParams contains all the parameters to send to the API endpoint
   for the ipam rirs read operation.

   Typically these are written to a http.Request.
*/
type IpamRirsReadParams struct {

	/* ID.

	   A UUID string identifying this RIR.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam rirs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsReadParams) WithDefaults() *IpamRirsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam rirs read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam rirs read params
func (o *IpamRirsReadParams) WithTimeout(timeout time.Duration) *IpamRirsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs read params
func (o *IpamRirsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs read params
func (o *IpamRirsReadParams) WithContext(ctx context.Context) *IpamRirsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs read params
func (o *IpamRirsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs read params
func (o *IpamRirsReadParams) WithHTTPClient(client *http.Client) *IpamRirsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs read params
func (o *IpamRirsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam rirs read params
func (o *IpamRirsReadParams) WithID(id strfmt.UUID) *IpamRirsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam rirs read params
func (o *IpamRirsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
