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

// NewIpamRouteTargetsDeleteParams creates a new IpamRouteTargetsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamRouteTargetsDeleteParams() *IpamRouteTargetsDeleteParams {
	return &IpamRouteTargetsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamRouteTargetsDeleteParamsWithTimeout creates a new IpamRouteTargetsDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamRouteTargetsDeleteParamsWithTimeout(timeout time.Duration) *IpamRouteTargetsDeleteParams {
	return &IpamRouteTargetsDeleteParams{
		timeout: timeout,
	}
}

// NewIpamRouteTargetsDeleteParamsWithContext creates a new IpamRouteTargetsDeleteParams object
// with the ability to set a context for a request.
func NewIpamRouteTargetsDeleteParamsWithContext(ctx context.Context) *IpamRouteTargetsDeleteParams {
	return &IpamRouteTargetsDeleteParams{
		Context: ctx,
	}
}

// NewIpamRouteTargetsDeleteParamsWithHTTPClient creates a new IpamRouteTargetsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamRouteTargetsDeleteParamsWithHTTPClient(client *http.Client) *IpamRouteTargetsDeleteParams {
	return &IpamRouteTargetsDeleteParams{
		HTTPClient: client,
	}
}

/* IpamRouteTargetsDeleteParams contains all the parameters to send to the API endpoint
   for the ipam route targets delete operation.

   Typically these are written to a http.Request.
*/
type IpamRouteTargetsDeleteParams struct {

	/* ID.

	   A UUID string identifying this route target.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam route targets delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsDeleteParams) WithDefaults() *IpamRouteTargetsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam route targets delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamRouteTargetsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) WithTimeout(timeout time.Duration) *IpamRouteTargetsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) WithContext(ctx context.Context) *IpamRouteTargetsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) WithHTTPClient(client *http.Client) *IpamRouteTargetsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) WithID(id strfmt.UUID) *IpamRouteTargetsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam route targets delete params
func (o *IpamRouteTargetsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamRouteTargetsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
