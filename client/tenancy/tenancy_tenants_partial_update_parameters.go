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

// NewTenancyTenantsPartialUpdateParams creates a new TenancyTenantsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsPartialUpdateParams() *TenancyTenantsPartialUpdateParams {
	return &TenancyTenantsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsPartialUpdateParamsWithTimeout creates a new TenancyTenantsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsPartialUpdateParamsWithTimeout(timeout time.Duration) *TenancyTenantsPartialUpdateParams {
	return &TenancyTenantsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsPartialUpdateParamsWithContext creates a new TenancyTenantsPartialUpdateParams object
// with the ability to set a context for a request.
func NewTenancyTenantsPartialUpdateParamsWithContext(ctx context.Context) *TenancyTenantsPartialUpdateParams {
	return &TenancyTenantsPartialUpdateParams{
		Context: ctx,
	}
}

// NewTenancyTenantsPartialUpdateParamsWithHTTPClient creates a new TenancyTenantsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsPartialUpdateParamsWithHTTPClient(client *http.Client) *TenancyTenantsPartialUpdateParams {
	return &TenancyTenantsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the tenancy tenants partial update operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsPartialUpdateParams struct {

	// Data.
	Data *models.WritableTenant

	/* ID.

	   A UUID string identifying this tenant.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsPartialUpdateParams) WithDefaults() *TenancyTenantsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) WithTimeout(timeout time.Duration) *TenancyTenantsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) WithContext(ctx context.Context) *TenancyTenantsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) WithHTTPClient(client *http.Client) *TenancyTenantsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) WithData(data *models.WritableTenant) *TenancyTenantsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) SetData(data *models.WritableTenant) {
	o.Data = data
}

// WithID adds the id to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) WithID(id strfmt.UUID) *TenancyTenantsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenants partial update params
func (o *TenancyTenantsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
