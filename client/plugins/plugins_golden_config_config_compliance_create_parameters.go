package plugins

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

// NewPluginsGoldenConfigConfigComplianceCreateParams creates a new PluginsGoldenConfigConfigComplianceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginsGoldenConfigConfigComplianceCreateParams() *PluginsGoldenConfigConfigComplianceCreateParams {
	return &PluginsGoldenConfigConfigComplianceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceCreateParamsWithTimeout creates a new PluginsGoldenConfigConfigComplianceCreateParams object
// with the ability to set a timeout on a request.
func NewPluginsGoldenConfigConfigComplianceCreateParamsWithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceCreateParams {
	return &PluginsGoldenConfigConfigComplianceCreateParams{
		timeout: timeout,
	}
}

// NewPluginsGoldenConfigConfigComplianceCreateParamsWithContext creates a new PluginsGoldenConfigConfigComplianceCreateParams object
// with the ability to set a context for a request.
func NewPluginsGoldenConfigConfigComplianceCreateParamsWithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceCreateParams {
	return &PluginsGoldenConfigConfigComplianceCreateParams{
		Context: ctx,
	}
}

// NewPluginsGoldenConfigConfigComplianceCreateParamsWithHTTPClient creates a new PluginsGoldenConfigConfigComplianceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginsGoldenConfigConfigComplianceCreateParamsWithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceCreateParams {
	return &PluginsGoldenConfigConfigComplianceCreateParams{
		HTTPClient: client,
	}
}

/* PluginsGoldenConfigConfigComplianceCreateParams contains all the parameters to send to the API endpoint
   for the plugins golden config config compliance create operation.

   Typically these are written to a http.Request.
*/
type PluginsGoldenConfigConfigComplianceCreateParams struct {

	// Data.
	Data *models.ConfigCompliance

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugins golden config config compliance create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceCreateParams) WithDefaults() *PluginsGoldenConfigConfigComplianceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugins golden config config compliance create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginsGoldenConfigConfigComplianceCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) WithTimeout(timeout time.Duration) *PluginsGoldenConfigConfigComplianceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) WithContext(ctx context.Context) *PluginsGoldenConfigConfigComplianceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) WithHTTPClient(client *http.Client) *PluginsGoldenConfigConfigComplianceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) WithData(data *models.ConfigCompliance) *PluginsGoldenConfigConfigComplianceCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the plugins golden config config compliance create params
func (o *PluginsGoldenConfigConfigComplianceCreateParams) SetData(data *models.ConfigCompliance) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *PluginsGoldenConfigConfigComplianceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
