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

// NewCircuitsCircuitTypesPartialUpdateParams creates a new CircuitsCircuitTypesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTypesPartialUpdateParams() *CircuitsCircuitTypesPartialUpdateParams {
	return &CircuitsCircuitTypesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesPartialUpdateParamsWithTimeout creates a new CircuitsCircuitTypesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTypesPartialUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesPartialUpdateParams {
	return &CircuitsCircuitTypesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesPartialUpdateParamsWithContext creates a new CircuitsCircuitTypesPartialUpdateParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTypesPartialUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTypesPartialUpdateParams {
	return &CircuitsCircuitTypesPartialUpdateParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTypesPartialUpdateParamsWithHTTPClient creates a new CircuitsCircuitTypesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTypesPartialUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesPartialUpdateParams {
	return &CircuitsCircuitTypesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTypesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the circuits circuit types partial update operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTypesPartialUpdateParams struct {

	// Data.
	Data *models.CircuitType

	/* ID.

	   A UUID string identifying this circuit type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit types partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesPartialUpdateParams) WithDefaults() *CircuitsCircuitTypesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit types partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithData(data *models.CircuitType) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetData(data *models.CircuitType) {
	o.Data = data
}

// WithID adds the id to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) WithID(id strfmt.UUID) *CircuitsCircuitTypesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types partial update params
func (o *CircuitsCircuitTypesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
