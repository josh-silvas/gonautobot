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

// NewTenancyTenantGroupsCreateParams creates a new TenancyTenantGroupsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantGroupsCreateParams() *TenancyTenantGroupsCreateParams {
	return &TenancyTenantGroupsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsCreateParamsWithTimeout creates a new TenancyTenantGroupsCreateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantGroupsCreateParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsCreateParams {
	return &TenancyTenantGroupsCreateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantGroupsCreateParamsWithContext creates a new TenancyTenantGroupsCreateParams object
// with the ability to set a context for a request.
func NewTenancyTenantGroupsCreateParamsWithContext(ctx context.Context) *TenancyTenantGroupsCreateParams {
	return &TenancyTenantGroupsCreateParams{
		Context: ctx,
	}
}

// NewTenancyTenantGroupsCreateParamsWithHTTPClient creates a new TenancyTenantGroupsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantGroupsCreateParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsCreateParams {
	return &TenancyTenantGroupsCreateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantGroupsCreateParams contains all the parameters to send to the API endpoint
   for the tenancy tenant groups create operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantGroupsCreateParams struct {

	// Data.
	Data *models.WritableTenantGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenant groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsCreateParams) WithDefaults() *TenancyTenantGroupsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenant groups create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithContext(ctx context.Context) *TenancyTenantGroupsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) WithData(data *models.WritableTenantGroup) *TenancyTenantGroupsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenant groups create params
func (o *TenancyTenantGroupsCreateParams) SetData(data *models.WritableTenantGroup) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
