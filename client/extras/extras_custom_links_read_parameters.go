package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExtrasCustomLinksReadParams creates a new ExtrasCustomLinksReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasCustomLinksReadParams() *ExtrasCustomLinksReadParams {
	return &ExtrasCustomLinksReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasCustomLinksReadParamsWithTimeout creates a new ExtrasCustomLinksReadParams object
// with the ability to set a timeout on a request.
func NewExtrasCustomLinksReadParamsWithTimeout(timeout time.Duration) *ExtrasCustomLinksReadParams {
	return &ExtrasCustomLinksReadParams{
		timeout: timeout,
	}
}

// NewExtrasCustomLinksReadParamsWithContext creates a new ExtrasCustomLinksReadParams object
// with the ability to set a context for a request.
func NewExtrasCustomLinksReadParamsWithContext(ctx context.Context) *ExtrasCustomLinksReadParams {
	return &ExtrasCustomLinksReadParams{
		Context: ctx,
	}
}

// NewExtrasCustomLinksReadParamsWithHTTPClient creates a new ExtrasCustomLinksReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasCustomLinksReadParamsWithHTTPClient(client *http.Client) *ExtrasCustomLinksReadParams {
	return &ExtrasCustomLinksReadParams{
		HTTPClient: client,
	}
}

/* ExtrasCustomLinksReadParams contains all the parameters to send to the API endpoint
   for the extras custom links read operation.

   Typically these are written to a http.Request.
*/
type ExtrasCustomLinksReadParams struct {

	/* ID.

	   A UUID string identifying this custom link.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras custom links read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksReadParams) WithDefaults() *ExtrasCustomLinksReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras custom links read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasCustomLinksReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) WithTimeout(timeout time.Duration) *ExtrasCustomLinksReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) WithContext(ctx context.Context) *ExtrasCustomLinksReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) WithHTTPClient(client *http.Client) *ExtrasCustomLinksReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) WithID(id strfmt.UUID) *ExtrasCustomLinksReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras custom links read params
func (o *ExtrasCustomLinksReadParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasCustomLinksReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
