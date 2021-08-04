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

// NewTenancyTenantsBulkUpdateParams creates a new TenancyTenantsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsBulkUpdateParams() *TenancyTenantsBulkUpdateParams {
	return &TenancyTenantsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsBulkUpdateParamsWithTimeout creates a new TenancyTenantsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsBulkUpdateParamsWithTimeout(timeout time.Duration) *TenancyTenantsBulkUpdateParams {
	return &TenancyTenantsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsBulkUpdateParamsWithContext creates a new TenancyTenantsBulkUpdateParams object
// with the ability to set a context for a request.
func NewTenancyTenantsBulkUpdateParamsWithContext(ctx context.Context) *TenancyTenantsBulkUpdateParams {
	return &TenancyTenantsBulkUpdateParams{
		Context: ctx,
	}
}

// NewTenancyTenantsBulkUpdateParamsWithHTTPClient creates a new TenancyTenantsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsBulkUpdateParamsWithHTTPClient(client *http.Client) *TenancyTenantsBulkUpdateParams {
	return &TenancyTenantsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the tenancy tenants bulk update operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsBulkUpdateParams struct {

	// Data.
	Data *models.WritableTenant

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsBulkUpdateParams) WithDefaults() *TenancyTenantsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) WithTimeout(timeout time.Duration) *TenancyTenantsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) WithContext(ctx context.Context) *TenancyTenantsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) WithHTTPClient(client *http.Client) *TenancyTenantsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) WithData(data *models.WritableTenant) *TenancyTenantsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenants bulk update params
func (o *TenancyTenantsBulkUpdateParams) SetData(data *models.WritableTenant) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
