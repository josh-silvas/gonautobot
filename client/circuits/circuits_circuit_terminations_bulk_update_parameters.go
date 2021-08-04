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

// NewCircuitsCircuitTerminationsBulkUpdateParams creates a new CircuitsCircuitTerminationsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTerminationsBulkUpdateParams() *CircuitsCircuitTerminationsBulkUpdateParams {
	return &CircuitsCircuitTerminationsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateParamsWithTimeout creates a new CircuitsCircuitTerminationsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTerminationsBulkUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsBulkUpdateParams {
	return &CircuitsCircuitTerminationsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateParamsWithContext creates a new CircuitsCircuitTerminationsBulkUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTerminationsBulkUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsBulkUpdateParams {
	return &CircuitsCircuitTerminationsBulkUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsBulkUpdateParamsWithHTTPClient creates a new CircuitsCircuitTerminationsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTerminationsBulkUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsBulkUpdateParams {
	return &CircuitsCircuitTerminationsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTerminationsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the circuits circuit terminations bulk update operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTerminationsBulkUpdateParams struct {

	// Data.
	Data *models.WritableCircuitTermination

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit terminations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithDefaults() *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit terminations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WithData(data *models.WritableCircuitTermination) *CircuitsCircuitTerminationsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit terminations bulk update params
func (o *CircuitsCircuitTerminationsBulkUpdateParams) SetData(data *models.WritableCircuitTermination) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
