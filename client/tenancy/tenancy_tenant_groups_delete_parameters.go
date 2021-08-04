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

// NewTenancyTenantGroupsDeleteParams creates a new TenancyTenantGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantGroupsDeleteParams() *TenancyTenantGroupsDeleteParams {
	return &TenancyTenantGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsDeleteParamsWithTimeout creates a new TenancyTenantGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantGroupsDeleteParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsDeleteParams {
	return &TenancyTenantGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewTenancyTenantGroupsDeleteParamsWithContext creates a new TenancyTenantGroupsDeleteParams object
// with the ability to set a context for a request.
func NewTenancyTenantGroupsDeleteParamsWithContext(ctx context.Context) *TenancyTenantGroupsDeleteParams {
	return &TenancyTenantGroupsDeleteParams{
		Context: ctx,
	}
}

// NewTenancyTenantGroupsDeleteParamsWithHTTPClient creates a new TenancyTenantGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantGroupsDeleteParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsDeleteParams {
	return &TenancyTenantGroupsDeleteParams{
		HTTPClient: client,
	}
}

/* TenancyTenantGroupsDeleteParams contains all the parameters to send to the API endpoint
   for the tenancy tenant groups delete operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantGroupsDeleteParams struct {

	/* ID.

	   A UUID string identifying this tenant group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenant groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsDeleteParams) WithDefaults() *TenancyTenantGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenant groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) WithContext(ctx context.Context) *TenancyTenantGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) WithID(id strfmt.UUID) *TenancyTenantGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenant groups delete params
func (o *TenancyTenantGroupsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
