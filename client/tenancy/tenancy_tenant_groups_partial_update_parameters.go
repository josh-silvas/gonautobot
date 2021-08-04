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

// NewTenancyTenantGroupsPartialUpdateParams creates a new TenancyTenantGroupsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTenancyTenantGroupsPartialUpdateParams() *TenancyTenantGroupsPartialUpdateParams {
	return &TenancyTenantGroupsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTenancyTenantGroupsPartialUpdateParamsWithTimeout creates a new TenancyTenantGroupsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewTenancyTenantGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *TenancyTenantGroupsPartialUpdateParams {
	return &TenancyTenantGroupsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewTenancyTenantGroupsPartialUpdateParamsWithContext creates a new TenancyTenantGroupsPartialUpdateParams object
// with the ability to set a context for a request.
func NewTenancyTenantGroupsPartialUpdateParamsWithContext(ctx context.Context) *TenancyTenantGroupsPartialUpdateParams {
	return &TenancyTenantGroupsPartialUpdateParams{
		Context: ctx,
	}
}

// NewTenancyTenantGroupsPartialUpdateParamsWithHTTPClient creates a new TenancyTenantGroupsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewTenancyTenantGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *TenancyTenantGroupsPartialUpdateParams {
	return &TenancyTenantGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/* TenancyTenantGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
   for the tenancy tenant groups partial update operation.

   Typically these are written to a http.Request.
*/
type TenancyTenantGroupsPartialUpdateParams struct {

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

// WithDefaults hydrates default values in the tenancy tenant groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsPartialUpdateParams) WithDefaults() *TenancyTenantGroupsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tenancy tenant groups partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TenancyTenantGroupsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *TenancyTenantGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) WithContext(ctx context.Context) *TenancyTenantGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *TenancyTenantGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) WithData(data *models.WritableTenantGroup) *TenancyTenantGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) SetData(data *models.WritableTenantGroup) {
	o.Data = data
}

// WithID adds the id to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) WithID(id strfmt.UUID) *TenancyTenantGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the tenancy tenant groups partial update params
func (o *TenancyTenantGroupsPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TenancyTenantGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
