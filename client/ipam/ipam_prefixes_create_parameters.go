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

// NewIpamPrefixesCreateParams creates a new IpamPrefixesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesCreateParams() *IpamPrefixesCreateParams {
	return &IpamPrefixesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesCreateParamsWithTimeout creates a new IpamPrefixesCreateParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesCreateParamsWithTimeout(timeout time.Duration) *IpamPrefixesCreateParams {
	return &IpamPrefixesCreateParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesCreateParamsWithContext creates a new IpamPrefixesCreateParams object
// with the ability to set a context for a request.
func NewIpamPrefixesCreateParamsWithContext(ctx context.Context) *IpamPrefixesCreateParams {
	return &IpamPrefixesCreateParams{
		Context: ctx,
	}
}

// NewIpamPrefixesCreateParamsWithHTTPClient creates a new IpamPrefixesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesCreateParamsWithHTTPClient(client *http.Client) *IpamPrefixesCreateParams {
	return &IpamPrefixesCreateParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesCreateParams contains all the parameters to send to the API endpoint
   for the ipam prefixes create operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesCreateParams struct {

	// Data.
	Data *models.WritablePrefix

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesCreateParams) WithDefaults() *IpamPrefixesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) WithTimeout(timeout time.Duration) *IpamPrefixesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) WithContext(ctx context.Context) *IpamPrefixesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) WithHTTPClient(client *http.Client) *IpamPrefixesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) WithData(data *models.WritablePrefix) *IpamPrefixesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam prefixes create params
func (o *IpamPrefixesCreateParams) SetData(data *models.WritablePrefix) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
