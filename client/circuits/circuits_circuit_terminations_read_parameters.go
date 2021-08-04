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

// NewCircuitsCircuitTerminationsReadParams creates a new CircuitsCircuitTerminationsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTerminationsReadParams() *CircuitsCircuitTerminationsReadParams {
	return &CircuitsCircuitTerminationsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTerminationsReadParamsWithTimeout creates a new CircuitsCircuitTerminationsReadParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTerminationsReadParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsReadParams {
	return &CircuitsCircuitTerminationsReadParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTerminationsReadParamsWithContext creates a new CircuitsCircuitTerminationsReadParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTerminationsReadParamsWithContext(ctx context.Context) *CircuitsCircuitTerminationsReadParams {
	return &CircuitsCircuitTerminationsReadParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTerminationsReadParamsWithHTTPClient creates a new CircuitsCircuitTerminationsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTerminationsReadParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsReadParams {
	return &CircuitsCircuitTerminationsReadParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTerminationsReadParams contains all the parameters to send to the API endpoint
   for the circuits circuit terminations read operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTerminationsReadParams struct {

	/* ID.

	   A UUID string identifying this circuit termination.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit terminations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsReadParams) WithDefaults() *CircuitsCircuitTerminationsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit terminations read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTerminationsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTerminationsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) WithContext(ctx context.Context) *CircuitsCircuitTerminationsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTerminationsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) WithID(id strfmt.UUID) *CircuitsCircuitTerminationsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit terminations read params
func (o *CircuitsCircuitTerminationsReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTerminationsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
