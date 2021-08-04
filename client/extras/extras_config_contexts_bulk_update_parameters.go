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

// NewExtrasConfigContextsBulkUpdateParams creates a new ExtrasConfigContextsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsBulkUpdateParams() *ExtrasConfigContextsBulkUpdateParams {
	return &ExtrasConfigContextsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsBulkUpdateParamsWithTimeout creates a new ExtrasConfigContextsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsBulkUpdateParams {
	return &ExtrasConfigContextsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsBulkUpdateParamsWithContext creates a new ExtrasConfigContextsBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasConfigContextsBulkUpdateParams {
	return &ExtrasConfigContextsBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsBulkUpdateParamsWithHTTPClient creates a new ExtrasConfigContextsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsBulkUpdateParams {
	return &ExtrasConfigContextsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras config contexts bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextsBulkUpdateParams struct {

	// Data.
	Data *models.WritableConfigContext

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsBulkUpdateParams) WithDefaults() *ExtrasConfigContextsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasConfigContextsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) WithData(data *models.WritableConfigContext) *ExtrasConfigContextsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras config contexts bulk update params
func (o *ExtrasConfigContextsBulkUpdateParams) SetData(data *models.WritableConfigContext) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
