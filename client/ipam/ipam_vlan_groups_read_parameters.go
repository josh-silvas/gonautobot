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

// NewIpamVlanGroupsReadParams creates a new IpamVlanGroupsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsReadParams() *IpamVlanGroupsReadParams {
	return &IpamVlanGroupsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsReadParamsWithTimeout creates a new IpamVlanGroupsReadParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsReadParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsReadParams {
	return &IpamVlanGroupsReadParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsReadParamsWithContext creates a new IpamVlanGroupsReadParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsReadParamsWithContext(ctx context.Context) *IpamVlanGroupsReadParams {
	return &IpamVlanGroupsReadParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsReadParamsWithHTTPClient creates a new IpamVlanGroupsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsReadParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsReadParams {
	return &IpamVlanGroupsReadParams{
		HTTPClient: client,
	}
}

/* IpamVlanGroupsReadParams contains all the parameters to send to the API endpoint
   for the ipam vlan groups read operation.

   Typically these are written to a http.Request.
*/
type IpamVlanGroupsReadParams struct {

	/* ID.

	   A UUID string identifying this VLAN group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsReadParams) WithDefaults() *IpamVlanGroupsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) WithContext(ctx context.Context) *IpamVlanGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) WithID(id strfmt.UUID) *IpamVlanGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups read params
func (o *IpamVlanGroupsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
