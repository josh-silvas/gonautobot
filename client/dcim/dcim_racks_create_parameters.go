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

// NewDcimRacksCreateParams creates a new DcimRacksCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksCreateParams() *DcimRacksCreateParams {
	return &DcimRacksCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksCreateParamsWithTimeout creates a new DcimRacksCreateParams object
// with the ability to set a timeout on a request.
func NewDcimRacksCreateParamsWithTimeout(timeout time.Duration) *DcimRacksCreateParams {
	return &DcimRacksCreateParams{
		timeout: timeout,
	}
}

// NewDcimRacksCreateParamsWithContext creates a new DcimRacksCreateParams object
// with the ability to set a context for a request.
func NewDcimRacksCreateParamsWithContext(ctx context.Context) *DcimRacksCreateParams {
	return &DcimRacksCreateParams{
		Context: ctx,
	}
}

// NewDcimRacksCreateParamsWithHTTPClient creates a new DcimRacksCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksCreateParamsWithHTTPClient(client *http.Client) *DcimRacksCreateParams {
	return &DcimRacksCreateParams{
		HTTPClient: client,
	}
}

/* DcimRacksCreateParams contains all the parameters to send to the API endpoint
   for the dcim racks create operation.

   Typically these are written to a http.Request.
*/
type DcimRacksCreateParams struct {

	// Data.
	Data *models.WritableRack

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksCreateParams) WithDefaults() *DcimRacksCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks create params
func (o *DcimRacksCreateParams) WithTimeout(timeout time.Duration) *DcimRacksCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks create params
func (o *DcimRacksCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks create params
func (o *DcimRacksCreateParams) WithContext(ctx context.Context) *DcimRacksCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks create params
func (o *DcimRacksCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks create params
func (o *DcimRacksCreateParams) WithHTTPClient(client *http.Client) *DcimRacksCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks create params
func (o *DcimRacksCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim racks create params
func (o *DcimRacksCreateParams) WithData(data *models.WritableRack) *DcimRacksCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim racks create params
func (o *DcimRacksCreateParams) SetData(data *models.WritableRack) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
