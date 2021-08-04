package circuits

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewCircuitsCircuitTerminationsDeleteParams creates a new CircuitsCircuitTerminationsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTerminationsDeleteParams() *CircuitsCircuitTerminationsDeleteParams {
	return &CircuitsCircuitTerminationsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsDeleteParamsWithTimeout creates a new CircuitsCircuitTerminationsDeleteParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTerminationsDeleteParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsDeleteParams {
	return &CircuitsCircuitTerminationsDeleteParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsDeleteParamsWithContext creates a new CircuitsCircuitTerminationsDeleteParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTerminationsDeleteParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsDeleteParams {
	return &CircuitsCircuitTerminationsDeleteParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsDeleteParamsWithHTTPClient creates a new CircuitsCircuitTerminationsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTerminationsDeleteParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsDeleteParams {
	return &CircuitsCircuitTerminationsDeleteParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTerminationsDeleteParams contains all the parameters to send to the API endpoint
   for the circuits circuit terminations delete operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTerminationsDeleteParams struct {

	/* ID.

	   A UUID string identifying this circuit termination.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit terminations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsDeleteParams) WithDefaults() *CircuitsCircuitTerminationsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit terminations delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) WithID(id strfmt.UUID) *CircuitsCircuitTerminationsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit terminations delete params
func (o *CircuitsCircuitTerminationsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
