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

// NewDcimSitesCreateParams creates a new DcimSitesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesCreateParams() *DcimSitesCreateParams {
	return &DcimSitesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesCreateParamsWithTimeout creates a new DcimSitesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimSitesCreateParamsWithTimeout(timeout time.Duration) *DcimSitesCreateParams {
	return &DcimSitesCreateParams{
		timeout: timeout,
	}
}

// NewDcimSitesCreateParamsWithContext creates a new DcimSitesCreateParams object
// with the ability to set a context for a request.
func NewDcimSitesCreateParamsWithContext(ctx context.Context) *DcimSitesCreateParams {
	return &DcimSitesCreateParams{
		Context: ctx,
	}
}

// NewDcimSitesCreateParamsWithHTTPClient creates a new DcimSitesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesCreateParamsWithHTTPClient(client *http.Client) *DcimSitesCreateParams {
	return &DcimSitesCreateParams{
		HTTPClient: client,
	}
}

/* DcimSitesCreateParams contains all the parameters to send to the API endpoint
   for the dcim sites create operation.

   Typically these are written to a http.Request.
*/
type DcimSitesCreateParams struct {

	// Data.
	Data *models.WritableSite

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim sites create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesCreateParams) WithDefaults() *DcimSitesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites create params
func (o *DcimSitesCreateParams) WithTimeout(timeout time.Duration) *DcimSitesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites create params
func (o *DcimSitesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites create params
func (o *DcimSitesCreateParams) WithContext(ctx context.Context) *DcimSitesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites create params
func (o *DcimSitesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites create params
func (o *DcimSitesCreateParams) WithHTTPClient(client *http.Client) *DcimSitesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites create params
func (o *DcimSitesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim sites create params
func (o *DcimSitesCreateParams) WithData(data *models.WritableSite) *DcimSitesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim sites create params
func (o *DcimSitesCreateParams) SetData(data *models.WritableSite) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
