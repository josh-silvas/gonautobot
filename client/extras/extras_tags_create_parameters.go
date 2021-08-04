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

// NewExtrasTagsCreateParams creates a new ExtrasTagsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasTagsCreateParams() *ExtrasTagsCreateParams {
	return &ExtrasTagsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTagsCreateParamsWithTimeout creates a new ExtrasTagsCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasTagsCreateParamsWithTimeout(timeout time.Duration) *ExtrasTagsCreateParams {
	return &ExtrasTagsCreateParams{
		timeout: timeout,
	}
}

// NewExtrasTagsCreateParamsWithContext creates a new ExtrasTagsCreateParams object
// with the ability to set a context for a request.
func NewExtrasTagsCreateParamsWithContext(ctx context.Context) *ExtrasTagsCreateParams {
	return &ExtrasTagsCreateParams{
		Context: ctx,
	}
}

// NewExtrasTagsCreateParamsWithHTTPClient creates a new ExtrasTagsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasTagsCreateParamsWithHTTPClient(client *http.Client) *ExtrasTagsCreateParams {
	return &ExtrasTagsCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasTagsCreateParams contains all the parameters to send to the API endpoint
   for the extras tags create operation.

   Typically these are written to a http.Request.
*/
type ExtrasTagsCreateParams struct {

	// Data.
	Data *models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras tags create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsCreateParams) WithDefaults() *ExtrasTagsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras tags create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasTagsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras tags create params
func (o *ExtrasTagsCreateParams) WithTimeout(timeout time.Duration) *ExtrasTagsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras tags create params
func (o *ExtrasTagsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras tags create params
func (o *ExtrasTagsCreateParams) WithContext(ctx context.Context) *ExtrasTagsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras tags create params
func (o *ExtrasTagsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras tags create params
func (o *ExtrasTagsCreateParams) WithHTTPClient(client *http.Client) *ExtrasTagsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras tags create params
func (o *ExtrasTagsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras tags create params
func (o *ExtrasTagsCreateParams) WithData(data *models.Tag) *ExtrasTagsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras tags create params
func (o *ExtrasTagsCreateParams) SetData(data *models.Tag) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTagsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
