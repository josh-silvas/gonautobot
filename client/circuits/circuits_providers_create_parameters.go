package circuits

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

// NewCircuitsProvidersCreateParams creates a new CircuitsProvidersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersCreateParams() *CircuitsProvidersCreateParams {
	return &CircuitsProvidersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersCreateParamsWithTimeout creates a new CircuitsProvidersCreateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersCreateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersCreateParams {
	return &CircuitsProvidersCreateParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersCreateParamsWithContext creates a new CircuitsProvidersCreateParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersCreateParamsWithContext(ctx context.Context) *CircuitsProvidersCreateParams {
	return &CircuitsProvidersCreateParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersCreateParamsWithHTTPClient creates a new CircuitsProvidersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersCreateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersCreateParams {
	return &CircuitsProvidersCreateParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersCreateParams contains all the parameters to send to the API endpoint
   for the circuits providers create operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersCreateParams struct {

	// Data.
	Data *models.Provider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersCreateParams) WithDefaults() *CircuitsProvidersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithContext(ctx context.Context) *CircuitsProvidersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers create params
func (o *CircuitsProvidersCreateParams) WithData(data *models.Provider) *CircuitsProvidersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers create params
func (o *CircuitsProvidersCreateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
