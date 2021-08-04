package ipam

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/josh5276/gonautobot/gonautobot/models"
)

// NewIpamVlanGroupsCreateParams creates a new IpamVlanGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlanGroupsCreateParams() *IpamVlanGroupsCreateParams {
	return &IpamVlanGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlanGroupsCreateParamsWithTimeout creates a new IpamVlanGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewIpamVlanGroupsCreateParamsWithTimeout(timeout time.Duration) *IpamVlanGroupsCreateParams {
	return &IpamVlanGroupsCreateParams{
		timeout: timeout,
	}
}

// NewIpamVlanGroupsCreateParamsWithContext creates a new IpamVlanGroupsCreateParams object
// with the ability to set a context for a request.
func NewIpamVlanGroupsCreateParamsWithContext(ctx context.Context) *IpamVlanGroupsCreateParams {
	return &IpamVlanGroupsCreateParams{
		Context: ctx,
	}
}

// NewIpamVlanGroupsCreateParamsWithHTTPClient creates a new IpamVlanGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlanGroupsCreateParamsWithHTTPClient(client *http.Client) *IpamVlanGroupsCreateParams {
	return &IpamVlanGroupsCreateParams{
		HTTPClient: client,
	}
}

/* IpamVlanGroupsCreateParams contains all the parameters to send to the API endpoint
   for the ipam vlan groups create operation.

   Typically these are written to a http.Request.
*/
type IpamVlanGroupsCreateParams struct {

	// Data.
	Data *models.WritableVLANGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlan groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsCreateParams) WithDefaults() *IpamVlanGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlan groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlanGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) WithTimeout(timeout time.Duration) *IpamVlanGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) WithContext(ctx context.Context) *IpamVlanGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) WithHTTPClient(client *http.Client) *IpamVlanGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) WithData(data *models.WritableVLANGroup) *IpamVlanGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlan groups create params
func (o *IpamVlanGroupsCreateParams) SetData(data *models.WritableVLANGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlanGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
