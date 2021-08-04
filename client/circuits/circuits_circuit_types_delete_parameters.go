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

// NewCircuitsCircuitTypesDeleteParams creates a new CircuitsCircuitTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitTypesDeleteParams() *CircuitsCircuitTypesDeleteParams {
	return &CircuitsCircuitTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesDeleteParamsWithTimeout creates a new CircuitsCircuitTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitTypesDeleteParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesDeleteParams {
	return &CircuitsCircuitTypesDeleteParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesDeleteParamsWithContext creates a new CircuitsCircuitTypesDeleteParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitTypesDeleteParamsWithContext(ctx context.Context) *CircuitsCircuitTypesDeleteParams {
	return &CircuitsCircuitTypesDeleteParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitTypesDeleteParamsWithHTTPClient creates a new CircuitsCircuitTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitTypesDeleteParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesDeleteParams {
	return &CircuitsCircuitTypesDeleteParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitTypesDeleteParams contains all the parameters to send to the API endpoint
   for the circuits circuit types delete operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitTypesDeleteParams struct {

	/* ID.

	   A UUID string identifying this circuit type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuit types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesDeleteParams) WithDefaults() *CircuitsCircuitTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuit types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithContext(ctx context.Context) *CircuitsCircuitTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) WithID(id strfmt.UUID) *CircuitsCircuitTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types delete params
func (o *CircuitsCircuitTypesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
