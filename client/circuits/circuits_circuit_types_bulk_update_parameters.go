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

// NewCircuitsCircuitTypesBulkUpdateParams creates a new CircuitsCircuitTypesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTypesBulkUpdateParams() *CircuitsCircuitTypesBulkUpdateParams {
	return &CircuitsCircuitTypesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesBulkUpdateParamsWithTimeout creates a new CircuitsCircuitTypesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTypesBulkUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesBulkUpdateParams {
	return &CircuitsCircuitTypesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesBulkUpdateParamsWithContext creates a new CircuitsCircuitTypesBulkUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTypesBulkUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTypesBulkUpdateParams {
	return &CircuitsCircuitTypesBulkUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTypesBulkUpdateParamsWithHTTPClient creates a new CircuitsCircuitTypesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTypesBulkUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesBulkUpdateParams {
	return &CircuitsCircuitTypesBulkUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTypesBulkUpdateParams contains all the parameters to send to the API endpoint
   for the circuits circuit types bulk update operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTypesBulkUpdateParams struct {

	// Data.
	Data *models.CircuitType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit types bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesBulkUpdateParams) WithDefaults() *CircuitsCircuitTypesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit types bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTypesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) WithData(data *models.CircuitType) *CircuitsCircuitTypesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit types bulk update params
func (o *CircuitsCircuitTypesBulkUpdateParams) SetData(data *models.CircuitType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
