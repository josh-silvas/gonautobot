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

// NewIpamRirsDeleteParams creates a new IpamRirsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRirsDeleteParams() *IpamRirsDeleteParams {
	return &IpamRirsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRirsDeleteParamsWithTimeout creates a new IpamRirsDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamRirsDeleteParamsWithTimeout(timeout time.Duration) *IpamRirsDeleteParams {
	return &IpamRirsDeleteParams{
		timeout: timeout,
	}
}

// NewIpamRirsDeleteParamsWithContext creates a new IpamRirsDeleteParams object
// with the ability to set a context for a request.
func NewIpamRirsDeleteParamsWithContext(ctx context.Context) *IpamRirsDeleteParams {
	return &IpamRirsDeleteParams{
		Context: ctx,
	}
}

// NewIpamRirsDeleteParamsWithHTTPClient creates a new IpamRirsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRirsDeleteParamsWithHTTPClient(client *http.Client) *IpamRirsDeleteParams {
	return &IpamRirsDeleteParams{
		HTTPClient: client,
	}
}

/* IpamRirsDeleteParams contains all the parameters to send to the API endpoint
   for the ipam rirs delete operation.

   Typically these are written to a http.Request.
*/
type IpamRirsDeleteParams struct {

	/* ID.

	   A UUID string identifying this RIR.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam rirs delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsDeleteParams) WithDefaults() *IpamRirsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam rirs delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRirsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam rirs delete params
func (o *IpamRirsDeleteParams) WithTimeout(timeout time.Duration) *IpamRirsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam rirs delete params
func (o *IpamRirsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam rirs delete params
func (o *IpamRirsDeleteParams) WithContext(ctx context.Context) *IpamRirsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam rirs delete params
func (o *IpamRirsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam rirs delete params
func (o *IpamRirsDeleteParams) WithHTTPClient(client *http.Client) *IpamRirsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam rirs delete params
func (o *IpamRirsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam rirs delete params
func (o *IpamRirsDeleteParams) WithID(id strfmt.UUID) *IpamRirsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam rirs delete params
func (o *IpamRirsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRirsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
