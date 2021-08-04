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

// NewIpamServicesUpdateParams creates a new IpamServicesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesUpdateParams() *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesUpdateParamsWithTimeout creates a new IpamServicesUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamServicesUpdateParamsWithTimeout(timeout time.Duration) *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		timeout: timeout,
	}
}

// NewIpamServicesUpdateParamsWithContext creates a new IpamServicesUpdateParams object
// with the ability to set a context for a request.
func NewIpamServicesUpdateParamsWithContext(ctx context.Context) *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		Context: ctx,
	}
}

// NewIpamServicesUpdateParamsWithHTTPClient creates a new IpamServicesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesUpdateParamsWithHTTPClient(client *http.Client) *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		HTTPClient: client,
	}
}

/* IpamServicesUpdateParams contains all the parameters to send to the API endpoint
   for the ipam services update operation.

   Typically these are written to a http.Request.
*/
type IpamServicesUpdateParams struct {

	// Data.
	Data *models.WritableService

	/* ID.

	   A UUID string identifying this service.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesUpdateParams) WithDefaults() *IpamServicesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services update params
func (o *IpamServicesUpdateParams) WithTimeout(timeout time.Duration) *IpamServicesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services update params
func (o *IpamServicesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services update params
func (o *IpamServicesUpdateParams) WithContext(ctx context.Context) *IpamServicesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services update params
func (o *IpamServicesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services update params
func (o *IpamServicesUpdateParams) WithHTTPClient(client *http.Client) *IpamServicesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services update params
func (o *IpamServicesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services update params
func (o *IpamServicesUpdateParams) WithData(data *models.WritableService) *IpamServicesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services update params
func (o *IpamServicesUpdateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WithID adds the id to the ipam services update params
func (o *IpamServicesUpdateParams) WithID(id strfmt.UUID) *IpamServicesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services update params
func (o *IpamServicesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
