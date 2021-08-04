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

// NewIpamPrefixesAvailableIpsReadParams creates a new IpamPrefixesAvailableIpsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesAvailableIpsReadParams() *IpamPrefixesAvailableIpsReadParams {
	return &IpamPrefixesAvailableIpsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesAvailableIpsReadParamsWithTimeout creates a new IpamPrefixesAvailableIpsReadParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesAvailableIpsReadParamsWithTimeout(timeout time.Duration) *IpamPrefixesAvailableIpsReadParams {
	return &IpamPrefixesAvailableIpsReadParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesAvailableIpsReadParamsWithContext creates a new IpamPrefixesAvailableIpsReadParams object
// with the ability to set a context for a request.
func NewIpamPrefixesAvailableIpsReadParamsWithContext(ctx context.Context) *IpamPrefixesAvailableIpsReadParams {
	return &IpamPrefixesAvailableIpsReadParams{
		Context: ctx,
	}
}

// NewIpamPrefixesAvailableIpsReadParamsWithHTTPClient creates a new IpamPrefixesAvailableIpsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesAvailableIpsReadParamsWithHTTPClient(client *http.Client) *IpamPrefixesAvailableIpsReadParams {
	return &IpamPrefixesAvailableIpsReadParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesAvailableIpsReadParams contains all the parameters to send to the API endpoint
   for the ipam prefixes available ips read operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesAvailableIpsReadParams struct {

	/* ID.

	   A UUID string identifying this prefix.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes available ips read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailableIpsReadParams) WithDefaults() *IpamPrefixesAvailableIpsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes available ips read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailableIpsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) WithTimeout(timeout time.Duration) *IpamPrefixesAvailableIpsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) WithContext(ctx context.Context) *IpamPrefixesAvailableIpsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) WithHTTPClient(client *http.Client) *IpamPrefixesAvailableIpsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) WithID(id strfmt.UUID) *IpamPrefixesAvailableIpsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes available ips read params
func (o *IpamPrefixesAvailableIpsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesAvailableIpsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
