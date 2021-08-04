package tenancy

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewTenancyTenantsReadParams creates a new TenancyTenantsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantsReadParams() *TenancyTenantsReadParams {
	return &TenancyTenantsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantsReadParamsWithTimeout creates a new TenancyTenantsReadParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantsReadParamsWithTimeout(timeout time.Duration) *TenancyTenantsReadParams {
	return &TenancyTenantsReadParams{
		timeout: timeout,
	}
}

// NewTenancyTenantsReadParamsWithContext creates a new TenancyTenantsReadParams object
// with the ability to set a context for a request.
func NewTenancyTenantsReadParamsWithContext(ctx context.Context) *TenancyTenantsReadParams {
	return &TenancyTenantsReadParams{
		Context: ctx,
	}
}

// NewTenancyTenantsReadParamsWithHTTPClient creates a new TenancyTenantsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantsReadParamsWithHTTPClient(client *http.Client) *TenancyTenantsReadParams {
	return &TenancyTenantsReadParams{
		HTTPClient: client,
	}
}

/* TenancyTenantsReadParams contains all the parameters to send to the API endpoint
   for the tenancy tenants read operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantsReadParams struct {

	/* ID.

	   A UUID string identifying this tenant.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenants read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsReadParams) WithDefaults() *TenancyTenantsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenants read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenants read params
func (o *TenancyTenantsReadParams) WithTimeout(timeout time.Duration) *TenancyTenantsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenants read params
func (o *TenancyTenantsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenants read params
func (o *TenancyTenantsReadParams) WithContext(ctx context.Context) *TenancyTenantsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenants read params
func (o *TenancyTenantsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenants read params
func (o *TenancyTenantsReadParams) WithHTTPClient(client *http.Client) *TenancyTenantsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenants read params
func (o *TenancyTenantsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenants read params
func (o *TenancyTenantsReadParams) WithID(id strfmt.UUID) *TenancyTenantsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenants read params
func (o *TenancyTenantsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
