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

// NewIpamPrefixesAvailableIpsCreateParams creates a new IpamPrefixesAvailableIpsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesAvailableIpsCreateParams() *IpamPrefixesAvailableIpsCreateParams {
	return &IpamPrefixesAvailableIpsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesAvailableIpsCreateParamsWithTimeout creates a new IpamPrefixesAvailableIpsCreateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesAvailableIpsCreateParamsWithTimeout(timeout time.Duration) *IpamPrefixesAvailableIpsCreateParams {
	return &IpamPrefixesAvailableIpsCreateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesAvailableIpsCreateParamsWithContext creates a new IpamPrefixesAvailableIpsCreateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesAvailableIpsCreateParamsWithContext(ctx context.Context) *IpamPrefixesAvailableIpsCreateParams {
	return &IpamPrefixesAvailableIpsCreateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesAvailableIpsCreateParamsWithHTTPClient creates a new IpamPrefixesAvailableIpsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesAvailableIpsCreateParamsWithHTTPClient(client *http.Client) *IpamPrefixesAvailableIpsCreateParams {
	return &IpamPrefixesAvailableIpsCreateParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesAvailableIpsCreateParams contains all the parameters to send to the API endpoint
   for the ipam prefixes available ips create operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesAvailableIpsCreateParams struct {

	// Data.
	Data []*models.AvailableIP

	/* ID.

	   A UUID string identifying this prefix.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes available ips create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailableIpsCreateParams) WithDefaults() *IpamPrefixesAvailableIpsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes available ips create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesAvailableIpsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithTimeout(timeout time.Duration) *IpamPrefixesAvailableIpsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithContext(ctx context.Context) *IpamPrefixesAvailableIpsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithHTTPClient(client *http.Client) *IpamPrefixesAvailableIpsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithData(data []*models.AvailableIP) *IpamPrefixesAvailableIpsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetData(data []*models.AvailableIP) {
	o.Data = data
}

// WithID adds the id to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) WithID(id strfmt.UUID) *IpamPrefixesAvailableIpsCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes available ips create params
func (o *IpamPrefixesAvailableIpsCreateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesAvailableIpsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
