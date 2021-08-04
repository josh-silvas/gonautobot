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

// NewIpamPrefixesDeleteParams creates a new IpamPrefixesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesDeleteParams() *IpamPrefixesDeleteParams {
	return &IpamPrefixesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesDeleteParamsWithTimeout creates a new IpamPrefixesDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesDeleteParamsWithTimeout(timeout time.Duration) *IpamPrefixesDeleteParams {
	return &IpamPrefixesDeleteParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesDeleteParamsWithContext creates a new IpamPrefixesDeleteParams object
// with the ability to set a context for a request.
func NewIpamPrefixesDeleteParamsWithContext(ctx context.Context) *IpamPrefixesDeleteParams {
	return &IpamPrefixesDeleteParams{
		Context: ctx,
	}
}

// NewIpamPrefixesDeleteParamsWithHTTPClient creates a new IpamPrefixesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesDeleteParamsWithHTTPClient(client *http.Client) *IpamPrefixesDeleteParams {
	return &IpamPrefixesDeleteParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesDeleteParams contains all the parameters to send to the API endpoint
   for the ipam prefixes delete operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesDeleteParams struct {

	/* ID.

	   A UUID string identifying this prefix.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesDeleteParams) WithDefaults() *IpamPrefixesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithTimeout(timeout time.Duration) *IpamPrefixesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithContext(ctx context.Context) *IpamPrefixesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithHTTPClient(client *http.Client) *IpamPrefixesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) WithID(id strfmt.UUID) *IpamPrefixesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes delete params
func (o *IpamPrefixesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
