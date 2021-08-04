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

// NewIpamVlanGroupsDeleteParams creates a new IpamVlanGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsDeleteParams() *IpamVlanGroupsDeleteParams {
	return &IpamVlanGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsDeleteParamsWithTimeout creates a new IpamVlanGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsDeleteParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsDeleteParams {
	return &IpamVlanGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsDeleteParamsWithContext creates a new IpamVlanGroupsDeleteParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsDeleteParamsWithContext(ctx context.Context) *IpamVlanGroupsDeleteParams {
	return &IpamVlanGroupsDeleteParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsDeleteParamsWithHTTPClient creates a new IpamVlanGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsDeleteParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsDeleteParams {
	return &IpamVlanGroupsDeleteParams{
		HTTPClient: client,
	}
}

/* IpamVlanGroupsDeleteParams contains all the parameters to send to the API endpoint
   for the ipam vlan groups delete operation.

   Typically these are written to a http.Request.
*/
type IpamVlanGroupsDeleteParams struct {

	/* ID.

	   A UUID string identifying this VLAN group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsDeleteParams) WithDefaults() *IpamVlanGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) WithContext(ctx context.Context) *IpamVlanGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) WithID(id strfmt.UUID) *IpamVlanGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups delete params
func (o *IpamVlanGroupsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
