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

// NewExtrasTagsBulkUpdateParams creates a new ExtrasTagsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasTagsBulkUpdateParams() *ExtrasTagsBulkUpdateParams {
	return &ExtrasTagsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsBulkUpdateParamsWithTimeout creates a new ExtrasTagsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewExtrasTagsBulkUpdateParamsWithTimeout(timeout time.Duration) *ExtrasTagsBulkUpdateParams {
	return &ExtrasTagsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewExtrasTagsBulkUpdateParamsWithContext creates a new ExtrasTagsBulkUpdateParams object
// with the ability to set a context for a request.
func NewExtrasTagsBulkUpdateParamsWithContext(ctx context.Context) *ExtrasTagsBulkUpdateParams {
	return &ExtrasTagsBulkUpdateParams{
		Context: ctx,
	}
}

// NewExtrasTagsBulkUpdateParamsWithHTTPClient creates a new ExtrasTagsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasTagsBulkUpdateParamsWithHTTPClient(client *http.Client) *ExtrasTagsBulkUpdateParams {
	return &ExtrasTagsBulkUpdateParams{
		HTTPClient: client,
	}
}

/* ExtrasTagsBulkUpdateParams contains all the parameters to send to the API endpoint
   for the extras tags bulk update operation.

   Typically these are written to a http.Request.
*/
type ExtrasTagsBulkUpdateParams struct {

	// Data.
	Data *models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras tags bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsBulkUpdateParams) WithDefaults() *ExtrasTagsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras tags bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithTimeout(timeout time.Duration) *ExtrasTagsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithContext(ctx context.Context) *ExtrasTagsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithHTTPClient(client *http.Client) *ExtrasTagsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) WithData(data *models.Tag) *ExtrasTagsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras tags bulk update params
func (o *ExtrasTagsBulkUpdateParams) SetData(data *models.Tag) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
