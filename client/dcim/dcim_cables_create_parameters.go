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

// NewDcimCablesCreateParams creates a new DcimCablesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesCreateParams() *DcimCablesCreateParams {
	return &DcimCablesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesCreateParamsWithTimeout creates a new DcimCablesCreateParams object
// with the ability to set a timeout on a request.
func NewDcimCablesCreateParamsWithTimeout(timeout time.Duration) *DcimCablesCreateParams {
	return &DcimCablesCreateParams{
		timeout: timeout,
	}
}

// NewDcimCablesCreateParamsWithContext creates a new DcimCablesCreateParams object
// with the ability to set a context for a request.
func NewDcimCablesCreateParamsWithContext(ctx context.Context) *DcimCablesCreateParams {
	return &DcimCablesCreateParams{
		Context: ctx,
	}
}

// NewDcimCablesCreateParamsWithHTTPClient creates a new DcimCablesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesCreateParamsWithHTTPClient(client *http.Client) *DcimCablesCreateParams {
	return &DcimCablesCreateParams{
		HTTPClient: client,
	}
}

/* DcimCablesCreateParams contains all the parameters to send to the API endpoint
   for the dcim cables create operation.

   Typically these are written to a http.Request.
*/
type DcimCablesCreateParams struct {

	// Data.
	Data *models.WritableCable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesCreateParams) WithDefaults() *DcimCablesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables create params
func (o *DcimCablesCreateParams) WithTimeout(timeout time.Duration) *DcimCablesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables create params
func (o *DcimCablesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables create params
func (o *DcimCablesCreateParams) WithContext(ctx context.Context) *DcimCablesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables create params
func (o *DcimCablesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables create params
func (o *DcimCablesCreateParams) WithHTTPClient(client *http.Client) *DcimCablesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables create params
func (o *DcimCablesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim cables create params
func (o *DcimCablesCreateParams) WithData(data *models.WritableCable) *DcimCablesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim cables create params
func (o *DcimCablesCreateParams) SetData(data *models.WritableCable) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
