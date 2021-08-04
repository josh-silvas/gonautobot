package extras

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

// NewExtrasConfigContextsUpdateParams creates a new ExtrasConfigContextsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsUpdateParams() *ExtrasConfigContextsUpdateParams {
	return &ExtrasConfigContextsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsUpdateParamsWithTimeout creates a new ExtrasConfigContextsUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsUpdateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsUpdateParams {
	return &ExtrasConfigContextsUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsUpdateParamsWithContext creates a new ExtrasConfigContextsUpdateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsUpdateParamsWithContext(ctx context.Context) *ExtrasConfigContextsUpdateParams {
	return &ExtrasConfigContextsUpdateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsUpdateParamsWithHTTPClient creates a new ExtrasConfigContextsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsUpdateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsUpdateParams {
	return &ExtrasConfigContextsUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextsUpdateParams contains all the parameters to send to the API endpoint
   for the extras config contexts update operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextsUpdateParams struct {

	// Data.
	Data *models.WritableConfigContext

	/* ID.

	   A UUID string identifying this config context.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsUpdateParams) WithDefaults() *ExtrasConfigContextsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) WithContext(ctx context.Context) *ExtrasConfigContextsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) WithData(data *models.WritableConfigContext) *ExtrasConfigContextsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) SetData(data *models.WritableConfigContext) {
	o.Data = data
}

// WithID adds the id to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) WithID(id strfmt.UUID) *ExtrasConfigContextsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config contexts update params
func (o *ExtrasConfigContextsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
