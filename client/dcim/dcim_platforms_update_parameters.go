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

// NewDcimPlatformsUpdateParams creates a new DcimPlatformsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPlatformsUpdateParams() *DcimPlatformsUpdateParams {
	return &DcimPlatformsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsUpdateParamsWithTimeout creates a new DcimPlatformsUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimPlatformsUpdateParamsWithTimeout(timeout time.Duration) *DcimPlatformsUpdateParams {
	return &DcimPlatformsUpdateParams{
		timeout: timeout,
	}
}

// NewDcimPlatformsUpdateParamsWithContext creates a new DcimPlatformsUpdateParams object
// with the ability to set a context for a request.
func NewDcimPlatformsUpdateParamsWithContext(ctx context.Context) *DcimPlatformsUpdateParams {
	return &DcimPlatformsUpdateParams{
		Context: ctx,
	}
}

// NewDcimPlatformsUpdateParamsWithHTTPClient creates a new DcimPlatformsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPlatformsUpdateParamsWithHTTPClient(client *http.Client) *DcimPlatformsUpdateParams {
	return &DcimPlatformsUpdateParams{
		HTTPClient: client,
	}
}

/* DcimPlatformsUpdateParams contains all the parameters to send to the API endpoint
   for the dcim platforms update operation.

   Typically these are written to a http.Request.
*/
type DcimPlatformsUpdateParams struct {

	// Data.
	Data *models.WritablePlatform

	/* ID.

	   A UUID string identifying this platform.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim platforms update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsUpdateParams) WithDefaults() *DcimPlatformsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim platforms update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) WithTimeout(timeout time.Duration) *DcimPlatformsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) WithContext(ctx context.Context) *DcimPlatformsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) WithHTTPClient(client *http.Client) *DcimPlatformsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) WithData(data *models.WritablePlatform) *DcimPlatformsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) SetData(data *models.WritablePlatform) {
	o.Data = data
}

// WithID adds the id to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) WithID(id strfmt.UUID) *DcimPlatformsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim platforms update params
func (o *DcimPlatformsUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
