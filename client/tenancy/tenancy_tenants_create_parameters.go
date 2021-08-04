package tenancy

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

// NewTenancyTenantsCreateParams creates a new TenancyTenantsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsCreateParams() *TenancyTenantsCreateParams {
	return &TenancyTenantsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsCreateParamsWithTimeout creates a new TenancyTenantsCreateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsCreateParamsWithTimeout(timeout time.Duration) *TenancyTenantsCreateParams {
	return &TenancyTenantsCreateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsCreateParamsWithContext creates a new TenancyTenantsCreateParams object
// with the ability to set a context for a request.
func NewTenancyTenantsCreateParamsWithContext(ctx context.Context) *TenancyTenantsCreateParams {
	return &TenancyTenantsCreateParams{
		Context: ctx,
	}
}

// NewTenancyTenantsCreateParamsWithHTTPClient creates a new TenancyTenantsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsCreateParamsWithHTTPClient(client *http.Client) *TenancyTenantsCreateParams {
	return &TenancyTenantsCreateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsCreateParams contains all the parameters to send to the API endpoint
   for the tenancy tenants create operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsCreateParams struct {

	// Data.
	Data *models.WritableTenant

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsCreateParams) WithDefaults() *TenancyTenantsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithTimeout(timeout time.Duration) *TenancyTenantsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithContext(ctx context.Context) *TenancyTenantsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithHTTPClient(client *http.Client) *TenancyTenantsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) WithData(data *models.WritableTenant) *TenancyTenantsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenants create params
func (o *TenancyTenantsCreateParams) SetData(data *models.WritableTenant) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
