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

// NewExtrasExportTemplatesReadParams creates a new ExtrasExportTemplatesReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasExportTemplatesReadParams() *ExtrasExportTemplatesReadParams {
	return &ExtrasExportTemplatesReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasExportTemplatesReadParamsWithTimeout creates a new ExtrasExportTemplatesReadParams object
// with the ability to set a timeout on a request.
func NewExtrasExportTemplatesReadParamsWithTimeout(timeout time.Duration) *ExtrasExportTemplatesReadParams {
	return &ExtrasExportTemplatesReadParams{
		timeout: timeout,
	}
}

// NewExtrasExportTemplatesReadParamsWithContext creates a new ExtrasExportTemplatesReadParams object
// with the ability to set a context for a request.
func NewExtrasExportTemplatesReadParamsWithContext(ctx context.Context) *ExtrasExportTemplatesReadParams {
	return &ExtrasExportTemplatesReadParams{
		Context: ctx,
	}
}

// NewExtrasExportTemplatesReadParamsWithHTTPClient creates a new ExtrasExportTemplatesReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasExportTemplatesReadParamsWithHTTPClient(client *http.Client) *ExtrasExportTemplatesReadParams {
	return &ExtrasExportTemplatesReadParams{
		HTTPClient: client,
	}
}

/* ExtrasExportTemplatesReadParams contains all the parameters to send to the API endpoint
   for the extras export templates read operation.

   Typically these are written to a http.Request.
*/
type ExtrasExportTemplatesReadParams struct {

	/* ID.

	   A UUID string identifying this export template.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras export templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesReadParams) WithDefaults() *ExtrasExportTemplatesReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras export templates read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasExportTemplatesReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithTimeout(timeout time.Duration) *ExtrasExportTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithContext(ctx context.Context) *ExtrasExportTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithHTTPClient(client *http.Client) *ExtrasExportTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) WithID(id strfmt.UUID) *ExtrasExportTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras export templates read params
func (o *ExtrasExportTemplatesReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasExportTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
