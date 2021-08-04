package extras

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

// NewExtrasConfigContextsCreateParams creates a new ExtrasConfigContextsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsCreateParams() *ExtrasConfigContextsCreateParams {
	return &ExtrasConfigContextsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsCreateParamsWithTimeout creates a new ExtrasConfigContextsCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsCreateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsCreateParams {
	return &ExtrasConfigContextsCreateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsCreateParamsWithContext creates a new ExtrasConfigContextsCreateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsCreateParamsWithContext(ctx context.Context) *ExtrasConfigContextsCreateParams {
	return &ExtrasConfigContextsCreateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsCreateParamsWithHTTPClient creates a new ExtrasConfigContextsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsCreateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsCreateParams {
	return &ExtrasConfigContextsCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextsCreateParams contains all the parameters to send to the API endpoint
   for the extras config contexts create operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextsCreateParams struct {

	// Data.
	Data *models.WritableConfigContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsCreateParams) WithDefaults() *ExtrasConfigContextsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) WithContext(ctx context.Context) *ExtrasConfigContextsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) WithData(data *models.WritableConfigContext) *ExtrasConfigContextsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config contexts create params
func (o *ExtrasConfigContextsCreateParams) SetData(data *models.WritableConfigContext) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
