package dcim

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

// NewDcimFrontPortTemplatesCreateParams creates a new DcimFrontPortTemplatesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesCreateParams() *DcimFrontPortTemplatesCreateParams {
	return &DcimFrontPortTemplatesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesCreateParamsWithTimeout creates a new DcimFrontPortTemplatesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesCreateParams {
	return &DcimFrontPortTemplatesCreateParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesCreateParamsWithContext creates a new DcimFrontPortTemplatesCreateParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesCreateParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesCreateParams {
	return &DcimFrontPortTemplatesCreateParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesCreateParamsWithHTTPClient creates a new DcimFrontPortTemplatesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesCreateParams {
	return &DcimFrontPortTemplatesCreateParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortTemplatesCreateParams contains all the parameters to send to the API endpoint
   for the dcim front port templates create operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesCreateParams struct {

	// Data.
	Data *models.WritableFrontPortTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesCreateParams) WithDefaults() *DcimFrontPortTemplatesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) WithData(data *models.WritableFrontPortTemplate) *DcimFrontPortTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim front port templates create params
func (o *DcimFrontPortTemplatesCreateParams) SetData(data *models.WritableFrontPortTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
