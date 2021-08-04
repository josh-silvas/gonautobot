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

// NewTenancyTenantGroupsReadParams creates a new TenancyTenantGroupsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantGroupsReadParams() *TenancyTenantGroupsReadParams {
	return &TenancyTenantGroupsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsReadParamsWithTimeout creates a new TenancyTenantGroupsReadParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantGroupsReadParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsReadParams {
	return &TenancyTenantGroupsReadParams{
		timeout: timeout,
	}
}

// NewTenancyTenantGroupsReadParamsWithContext creates a new TenancyTenantGroupsReadParams object
// with the ability to set a context for a request.
func NewTenancyTenantGroupsReadParamsWithContext(ctx context.Context) *TenancyTenantGroupsReadParams {
	return &TenancyTenantGroupsReadParams{
		Context: ctx,
	}
}

// NewTenancyTenantGroupsReadParamsWithHTTPClient creates a new TenancyTenantGroupsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantGroupsReadParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsReadParams {
	return &TenancyTenantGroupsReadParams{
		HTTPClient: client,
	}
}

/* TenancyTenantGroupsReadParams contains all the parameters to send to the API endpoint
   for the tenancy tenant groups read operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantGroupsReadParams struct {

	/* ID.

	   A UUID string identifying this tenant group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenant groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsReadParams) WithDefaults() *TenancyTenantGroupsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenant groups read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithContext(ctx context.Context) *TenancyTenantGroupsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) WithID(id strfmt.UUID) *TenancyTenantGroupsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenant groups read params
func (o *TenancyTenantGroupsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
