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

// NewCircuitsProvidersBulkUpdateParams creates a new CircuitsProvidersBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersBulkUpdateParams() *CircuitsProvidersBulkUpdateParams {
	return &CircuitsProvidersBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersBulkUpdateParamsWithTimeout creates a new CircuitsProvidersBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersBulkUpdateParamsWithTimeout(timeout time.Duration) *CircuitsProvidersBulkUpdateParams {
	return &CircuitsProvidersBulkUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersBulkUpdateParamsWithContext creates a new CircuitsProvidersBulkUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersBulkUpdateParamsWithContext(ctx context.Context) *CircuitsProvidersBulkUpdateParams {
	return &CircuitsProvidersBulkUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersBulkUpdateParamsWithHTTPClient creates a new CircuitsProvidersBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersBulkUpdateParamsWithHTTPClient(client *http.Client) *CircuitsProvidersBulkUpdateParams {
	return &CircuitsProvidersBulkUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersBulkUpdateParams contains all the parameters to send to the API endpoint
   for the circuits providers bulk update operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersBulkUpdateParams struct {

	// Data.
	Data *models.Provider

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersBulkUpdateParams) WithDefaults() *CircuitsProvidersBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithTimeout(timeout time.Duration) *CircuitsProvidersBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithContext(ctx context.Context) *CircuitsProvidersBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithHTTPClient(client *http.Client) *CircuitsProvidersBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) WithData(data *models.Provider) *CircuitsProvidersBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits providers bulk update params
func (o *CircuitsProvidersBulkUpdateParams) SetData(data *models.Provider) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
