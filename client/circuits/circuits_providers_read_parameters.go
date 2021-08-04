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

// NewCircuitsProvidersReadParams creates a new CircuitsProvidersReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersReadParams() *CircuitsProvidersReadParams {
	return &CircuitsProvidersReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersReadParamsWithTimeout creates a new CircuitsProvidersReadParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersReadParamsWithTimeout(timeout time.Duration) *CircuitsProvidersReadParams {
	return &CircuitsProvidersReadParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersReadParamsWithContext creates a new CircuitsProvidersReadParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersReadParamsWithContext(ctx context.Context) *CircuitsProvidersReadParams {
	return &CircuitsProvidersReadParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersReadParamsWithHTTPClient creates a new CircuitsProvidersReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersReadParamsWithHTTPClient(client *http.Client) *CircuitsProvidersReadParams {
	return &CircuitsProvidersReadParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersReadParams contains all the parameters to send to the API endpoint
   for the circuits providers read operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersReadParams struct {

	/* ID.

	   A UUID string identifying this provider.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersReadParams) WithDefaults() *CircuitsProvidersReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers read params
func (o *CircuitsProvidersReadParams) WithTimeout(timeout time.Duration) *CircuitsProvidersReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers read params
func (o *CircuitsProvidersReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers read params
func (o *CircuitsProvidersReadParams) WithContext(ctx context.Context) *CircuitsProvidersReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers read params
func (o *CircuitsProvidersReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers read params
func (o *CircuitsProvidersReadParams) WithHTTPClient(client *http.Client) *CircuitsProvidersReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers read params
func (o *CircuitsProvidersReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the circuits providers read params
func (o *CircuitsProvidersReadParams) WithID(id strfmt.UUID) *CircuitsProvidersReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits providers read params
func (o *CircuitsProvidersReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
