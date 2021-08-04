package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExtrasExportTemplatesDeleteParams creates a new ExtrasExportTemplatesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasExportTemplatesDeleteParams() *ExtrasExportTemplatesDeleteParams {
	return &ExtrasExportTemplatesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesDeleteParamsWithTimeout creates a new ExtrasExportTemplatesDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasExportTemplatesDeleteParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesDeleteParams {
	return &ExtrasExportTemplatesDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasExportTemplatesDeleteParamsWithContext creates a new ExtrasExportTemplatesDeleteParams object
// with the ability to set a context for a request.
func NewExtrasExportTemplatesDeleteParamsWithContext(ctx context.Context) *ExtrasExportTemplatesDeleteParams {
	return &ExtrasExportTemplatesDeleteParams{
		Context: ctx,
	}
}

// NewExtrasExportTemplatesDeleteParamsWithHTTPClient creates a new ExtrasExportTemplatesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasExportTemplatesDeleteParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesDeleteParams {
	return &ExtrasExportTemplatesDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasExportTemplatesDeleteParams contains all the parameters to send to the API endpoint
   for the extras export templates delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasExportTemplatesDeleteParams struct {

	/* ID.

	   A UUID string identifying this export template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras export templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesDeleteParams) WithDefaults() *ExtrasExportTemplatesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras export templates delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) WithContext(ctx context.Context) *ExtrasExportTemplatesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) WithID(id strfmt.UUID) *ExtrasExportTemplatesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates delete params
func (o *ExtrasExportTemplatesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
