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

// NewCircuitsCircuitTerminationsUpdateParams creates a new CircuitsCircuitTerminationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTerminationsUpdateParams() *CircuitsCircuitTerminationsUpdateParams {
	return &CircuitsCircuitTerminationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsUpdateParamsWithTimeout creates a new CircuitsCircuitTerminationsUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTerminationsUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsUpdateParams {
	return &CircuitsCircuitTerminationsUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsUpdateParamsWithContext creates a new CircuitsCircuitTerminationsUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTerminationsUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsUpdateParams {
	return &CircuitsCircuitTerminationsUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsUpdateParamsWithHTTPClient creates a new CircuitsCircuitTerminationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTerminationsUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsUpdateParams {
	return &CircuitsCircuitTerminationsUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTerminationsUpdateParams contains all the parameters to send to the API endpoint
   for the circuits circuit terminations update operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTerminationsUpdateParams struct {

	// Data.
	Data *models.WritableCircuitTermination

	/* ID.

	   A UUID string identifying this circuit termination.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit terminations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsUpdateParams) WithDefaults() *CircuitsCircuitTerminationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit terminations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) WithData(data *models.WritableCircuitTermination) *CircuitsCircuitTerminationsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) SetData(data *models.WritableCircuitTermination) {
	o.Data = data
}

// WithID adds the id to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) WithID(id strfmt.UUID) *CircuitsCircuitTerminationsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit terminations update params
func (o *CircuitsCircuitTerminationsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
