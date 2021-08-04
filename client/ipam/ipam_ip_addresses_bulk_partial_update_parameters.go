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

// NewIpamIPAddressesBulkPartialUpdateParams creates a new IpamIPAddressesBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPAddressesBulkPartialUpdateParams() *IpamIPAddressesBulkPartialUpdateParams {
	return &IpamIPAddressesBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesBulkPartialUpdateParamsWithTimeout creates a new IpamIPAddressesBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamIPAddressesBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamIPAddressesBulkPartialUpdateParams {
	return &IpamIPAddressesBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamIPAddressesBulkPartialUpdateParamsWithContext creates a new IpamIPAddressesBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamIPAddressesBulkPartialUpdateParamsWithContext(ctx context.Context) *IpamIPAddressesBulkPartialUpdateParams {
	return &IpamIPAddressesBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamIPAddressesBulkPartialUpdateParamsWithHTTPClient creates a new IpamIPAddressesBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPAddressesBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamIPAddressesBulkPartialUpdateParams {
	return &IpamIPAddressesBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* IpamIPAddressesBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the ipam ip addresses bulk partial update operation.

   Typically these are written to a http.Request.
*/
type IpamIPAddressesBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableIPAddress

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip addresses bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesBulkPartialUpdateParams) WithDefaults() *IpamIPAddressesBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip addresses bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamIPAddressesBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) WithContext(ctx context.Context) *IpamIPAddressesBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamIPAddressesBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) WithData(data *models.WritableIPAddress) *IpamIPAddressesBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam ip addresses bulk partial update params
func (o *IpamIPAddressesBulkPartialUpdateParams) SetData(data *models.WritableIPAddress) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
