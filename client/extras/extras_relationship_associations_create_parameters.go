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

// NewExtrasRelationshipAssociationsCreateParams creates a new ExtrasRelationshipAssociationsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasRelationshipAssociationsCreateParams() *ExtrasRelationshipAssociationsCreateParams {
	return &ExtrasRelationshipAssociationsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRelationshipAssociationsCreateParamsWithTimeout creates a new ExtrasRelationshipAssociationsCreateParams object
// with the ability to set a timeout on a request.
func NewExtrasRelationshipAssociationsCreateParamsWithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsCreateParams {
	return &ExtrasRelationshipAssociationsCreateParams{
		timeout: timeout,
	}
}

// NewExtrasRelationshipAssociationsCreateParamsWithContext creates a new ExtrasRelationshipAssociationsCreateParams object
// with the ability to set a context for a request.
func NewExtrasRelationshipAssociationsCreateParamsWithContext(ctx context.Context) *ExtrasRelationshipAssociationsCreateParams {
	return &ExtrasRelationshipAssociationsCreateParams{
		Context: ctx,
	}
}

// NewExtrasRelationshipAssociationsCreateParamsWithHTTPClient creates a new ExtrasRelationshipAssociationsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasRelationshipAssociationsCreateParamsWithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsCreateParams {
	return &ExtrasRelationshipAssociationsCreateParams{
		HTTPClient: client,
	}
}

/* ExtrasRelationshipAssociationsCreateParams contains all the parameters to send to the API endpoint
   for the extras relationship associations create operation.

   Typically these are written to a http.Request.
*/
type ExtrasRelationshipAssociationsCreateParams struct {

	// Data.
	Data *models.WritableRelationshipAssociation

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras relationship associations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsCreateParams) WithDefaults() *ExtrasRelationshipAssociationsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras relationship associations create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasRelationshipAssociationsCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) WithTimeout(timeout time.Duration) *ExtrasRelationshipAssociationsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) WithContext(ctx context.Context) *ExtrasRelationshipAssociationsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) WithHTTPClient(client *http.Client) *ExtrasRelationshipAssociationsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) WithData(data *models.WritableRelationshipAssociation) *ExtrasRelationshipAssociationsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras relationship associations create params
func (o *ExtrasRelationshipAssociationsCreateParams) SetData(data *models.WritableRelationshipAssociation) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRelationshipAssociationsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
