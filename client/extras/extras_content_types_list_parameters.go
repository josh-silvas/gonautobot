package extras

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewExtrasContentTypesListParams creates a new ExtrasContentTypesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasContentTypesListParams() *ExtrasContentTypesListParams {
	return &ExtrasContentTypesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasContentTypesListParamsWithTimeout creates a new ExtrasContentTypesListParams object
// with the ability to set a timeout on a request.
func NewExtrasContentTypesListParamsWithTimeout(timeout time.Duration) *ExtrasContentTypesListParams {
	return &ExtrasContentTypesListParams{
		timeout: timeout,
	}
}

// NewExtrasContentTypesListParamsWithContext creates a new ExtrasContentTypesListParams object
// with the ability to set a context for a request.
func NewExtrasContentTypesListParamsWithContext(ctx context.Context) *ExtrasContentTypesListParams {
	return &ExtrasContentTypesListParams{
		Context: ctx,
	}
}

// NewExtrasContentTypesListParamsWithHTTPClient creates a new ExtrasContentTypesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasContentTypesListParamsWithHTTPClient(client *http.Client) *ExtrasContentTypesListParams {
	return &ExtrasContentTypesListParams{
		HTTPClient: client,
	}
}

/* ExtrasContentTypesListParams contains all the parameters to send to the API endpoint
   for the extras content types list operation.

   Typically these are written to a http.Request.
*/
type ExtrasContentTypesListParams struct {

	// AppLabel.
	AppLabel *string

	// ID.
	ID *float64

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Model.
	Model *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras content types list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasContentTypesListParams) WithDefaults() *ExtrasContentTypesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras content types list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasContentTypesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras content types list params
func (o *ExtrasContentTypesListParams) WithTimeout(timeout time.Duration) *ExtrasContentTypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras content types list params
func (o *ExtrasContentTypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras content types list params
func (o *ExtrasContentTypesListParams) WithContext(ctx context.Context) *ExtrasContentTypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras content types list params
func (o *ExtrasContentTypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras content types list params
func (o *ExtrasContentTypesListParams) WithHTTPClient(client *http.Client) *ExtrasContentTypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras content types list params
func (o *ExtrasContentTypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppLabel adds the appLabel to the extras content types list params
func (o *ExtrasContentTypesListParams) WithAppLabel(appLabel *string) *ExtrasContentTypesListParams {
	o.SetAppLabel(appLabel)
	return o
}

// SetAppLabel adds the appLabel to the extras content types list params
func (o *ExtrasContentTypesListParams) SetAppLabel(appLabel *string) {
	o.AppLabel = appLabel
}

// WithID adds the id to the extras content types list params
func (o *ExtrasContentTypesListParams) WithID(id *float64) *ExtrasContentTypesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras content types list params
func (o *ExtrasContentTypesListParams) SetID(id *float64) {
	o.ID = id
}

// WithLimit adds the limit to the extras content types list params
func (o *ExtrasContentTypesListParams) WithLimit(limit *int64) *ExtrasContentTypesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras content types list params
func (o *ExtrasContentTypesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithModel adds the model to the extras content types list params
func (o *ExtrasContentTypesListParams) WithModel(model *string) *ExtrasContentTypesListParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the extras content types list params
func (o *ExtrasContentTypesListParams) SetModel(model *string) {
	o.Model = model
}

// WithOffset adds the offset to the extras content types list params
func (o *ExtrasContentTypesListParams) WithOffset(offset *int64) *ExtrasContentTypesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras content types list params
func (o *ExtrasContentTypesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasContentTypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AppLabel != nil {

		// query param app_label
		var qrAppLabel string

		if o.AppLabel != nil {
			qrAppLabel = *o.AppLabel
		}
		qAppLabel := qrAppLabel
		if qAppLabel != "" {

			if err := r.SetQueryParam("app_label", qAppLabel); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID float64

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := swag.FormatFloat64(qrID)
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Model != nil {

		// query param model
		var qrModel string

		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {

			if err := r.SetQueryParam("model", qModel); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
