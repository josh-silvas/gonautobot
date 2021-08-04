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

// NewIpamVrfsDeleteParams creates a new IpamVrfsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsDeleteParams() *IpamVrfsDeleteParams {
	return &IpamVrfsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsDeleteParamsWithTimeout creates a new IpamVrfsDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsDeleteParamsWithTimeout(timeout time.Duration) *IpamVrfsDeleteParams {
	return &IpamVrfsDeleteParams{
		timeout: timeout,
	}
}

// NewIpamVrfsDeleteParamsWithContext creates a new IpamVrfsDeleteParams object
// with the ability to set a context for a request.
func NewIpamVrfsDeleteParamsWithContext(ctx context.Context) *IpamVrfsDeleteParams {
	return &IpamVrfsDeleteParams{
		Context: ctx,
	}
}

// NewIpamVrfsDeleteParamsWithHTTPClient creates a new IpamVrfsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsDeleteParamsWithHTTPClient(client *http.Client) *IpamVrfsDeleteParams {
	return &IpamVrfsDeleteParams{
		HTTPClient: client,
	}
}

/* IpamVrfsDeleteParams contains all the parameters to send to the API endpoint
   for the ipam vrfs delete operation.

   Typically these are written to a http.Request.
*/
type IpamVrfsDeleteParams struct {

	/* ID.

	   A UUID string identifying this VRF.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vrfs delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsDeleteParams) WithDefaults() *IpamVrfsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) WithTimeout(timeout time.Duration) *IpamVrfsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) WithContext(ctx context.Context) *IpamVrfsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) WithHTTPClient(client *http.Client) *IpamVrfsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) WithID(id strfmt.UUID) *IpamVrfsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vrfs delete params
func (o *IpamVrfsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
