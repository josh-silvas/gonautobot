package plugins

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPluginsGoldenConfigConfigComplianceReadParams creates a new PluginsGoldenConfigConfigComplianceReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigComplianceReadParams() *PluginsGoldenConfigConfigComplianceReadParams {
	return &PluginsGoldenConfigConfigComplianceReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceReadParamsWithTimeout creates a new PluginsGoldenConfigConfigComplianceReadParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigComplianceReadParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceReadParams {
	return &PluginsGoldenConfigConfigComplianceReadParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceReadParamsWithContext creates a new PluginsGoldenConfigConfigComplianceReadParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigComplianceReadParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceReadParams {
	return &PluginsGoldenConfigConfigComplianceReadParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigComplianceReadParamsWithHTTPClient creates a new PluginsGoldenConfigConfigComplianceReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigComplianceReadParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceReadParams {
	return &PluginsGoldenConfigConfigComplianceReadParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigComplianceReadParams contains all the parameters to send to the API endpoint
   for the plugins golden config config compliance read operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigComplianceReadParams struct {

	/* ID.

	   A UUID string identifying this config compliance.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config compliance read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceReadParams) WithDefaults() *PluginsGoldenConfigConfigComplianceReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config compliance read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) WithID(id strfmt.UUID) *PluginsGoldenConfigConfigComplianceReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the plugins golden config config compliance read params
func (o *PluginsGoldenConfigConfigComplianceReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigComplianceReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
