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

// NewTenancyTenantGroupsUpdateParams creates a new TenancyTenantGroupsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantGroupsUpdateParams() *TenancyTenantGroupsUpdateParams {
	return &TenancyTenantGroupsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsUpdateParamsWithTimeout creates a new TenancyTenantGroupsUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantGroupsUpdateParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsUpdateParams {
	return &TenancyTenantGroupsUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantGroupsUpdateParamsWithContext creates a new TenancyTenantGroupsUpdateParams object
// with the ability to set a context for a request.
func NewTenancyTenantGroupsUpdateParamsWithContext(ctx context.Context) *TenancyTenantGroupsUpdateParams {
	return &TenancyTenantGroupsUpdateParams{
		Context: ctx,
	}
}

// NewTenancyTenantGroupsUpdateParamsWithHTTPClient creates a new TenancyTenantGroupsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantGroupsUpdateParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsUpdateParams {
	return &TenancyTenantGroupsUpdateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantGroupsUpdateParams contains all the parameters to send to the API endpoint
   for the tenancy tenant groups update operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantGroupsUpdateParams struct {

	// Data.
	Data *models.WritableTenantGroup

	/* ID.

	   A UUID string identifying this tenant group.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tenancy tenant groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsUpdateParams) WithDefaults() *TenancyTenantGroupsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenant groups update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) WithContext(ctx context.Context) *TenancyTenantGroupsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) WithData(data *models.WritableTenantGroup) *TenancyTenantGroupsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) SetData(data *models.WritableTenantGroup) {
	o.Data = data
}

// WithID adds the id to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) WithID(id strfmt.UUID) *TenancyTenantGroupsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenant groups update params
func (o *TenancyTenantGroupsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
