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

// NewIpamIPAddressesBulkUpdateParams creates a new IpamIPAddressesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPAddressesBulkUpdateParams() *IpamIPAddressesBulkUpdateParams {
	return &IpamIPAddressesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesBulkUpdateParamsWithTimeout creates a new IpamIPAddressesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamIPAddressesBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamIPAddressesBulkUpdateParams {
	return &IpamIPAddressesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamIPAddressesBulkUpdateParamsWithContext creates a new IpamIPAddressesBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamIPAddressesBulkUpdateParamsWithContext(ctx context.Context) *IpamIPAddressesBulkUpdateParams {
	return &IpamIPAddressesBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamIPAddressesBulkUpdateParamsWithHTTPClient creates a new IpamIPAddressesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPAddressesBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamIPAddressesBulkUpdateParams {
	return &IpamIPAddressesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* IpamIPAddressesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the ipam ip addresses bulk update operation.

   Typically these are written to a http.Request.
*/
type IpamIPAddressesBulkUpdateParams struct {

	// Data.
	Data *models.WritableIPAddress

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip addresses bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesBulkUpdateParams) WithDefaults() *IpamIPAddressesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip addresses bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamIPAddressesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) WithContext(ctx context.Context) *IpamIPAddressesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamIPAddressesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) WithData(data *models.WritableIPAddress) *IpamIPAddressesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam ip addresses bulk update params
func (o *IpamIPAddressesBulkUpdateParams) SetData(data *models.WritableIPAddress) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
