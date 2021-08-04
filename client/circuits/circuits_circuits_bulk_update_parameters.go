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

// NewCircuitsCircuitsBulkUpdateParams creates a new CircuitsCircuitsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitsBulkUpdateParams() *CircuitsCircuitsBulkUpdateParams {
	return &CircuitsCircuitsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitsBulkUpdateParamsWithTimeout creates a new CircuitsCircuitsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitsBulkUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitsBulkUpdateParams {
	return &CircuitsCircuitsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitsBulkUpdateParamsWithContext creates a new CircuitsCircuitsBulkUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitsBulkUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitsBulkUpdateParams {
	return &CircuitsCircuitsBulkUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitsBulkUpdateParamsWithHTTPClient creates a new CircuitsCircuitsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitsBulkUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitsBulkUpdateParams {
	return &CircuitsCircuitsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the circuits circuits bulk update operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitsBulkUpdateParams struct {

	// Data.
	Data *models.WritableCircuit

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuits bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitsBulkUpdateParams) WithDefaults() *CircuitsCircuitsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuits bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) WithData(data *models.WritableCircuit) *CircuitsCircuitsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuits bulk update params
func (o *CircuitsCircuitsBulkUpdateParams) SetData(data *models.WritableCircuit) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
