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

// NewDcimDevicesUpdateParams creates a new DcimDevicesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesUpdateParams() *DcimDevicesUpdateParams {
	return &DcimDevicesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesUpdateParamsWithTimeout creates a new DcimDevicesUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesUpdateParamsWithTimeout(timeout time.Duration) *DcimDevicesUpdateParams {
	return &DcimDevicesUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDevicesUpdateParamsWithContext creates a new DcimDevicesUpdateParams object
// with the ability to set a context for a request.
func NewDcimDevicesUpdateParamsWithContext(ctx context.Context) *DcimDevicesUpdateParams {
	return &DcimDevicesUpdateParams{
		Context: ctx,
	}
}

// NewDcimDevicesUpdateParamsWithHTTPClient creates a new DcimDevicesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesUpdateParamsWithHTTPClient(client *http.Client) *DcimDevicesUpdateParams {
	return &DcimDevicesUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDevicesUpdateParams contains all the parameters to send to the API endpoint
   for the dcim devices update operation.

   Typically these are written to a http.Request.
*/
type DcimDevicesUpdateParams struct {

	// Data.
	Data *models.WritableDeviceWithConfigContext

	/* ID.

	   A UUID string identifying this device.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesUpdateParams) WithDefaults() *DcimDevicesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices update params
func (o *DcimDevicesUpdateParams) WithTimeout(timeout time.Duration) *DcimDevicesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices update params
func (o *DcimDevicesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices update params
func (o *DcimDevicesUpdateParams) WithContext(ctx context.Context) *DcimDevicesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices update params
func (o *DcimDevicesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices update params
func (o *DcimDevicesUpdateParams) WithHTTPClient(client *http.Client) *DcimDevicesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices update params
func (o *DcimDevicesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim devices update params
func (o *DcimDevicesUpdateParams) WithData(data *models.WritableDeviceWithConfigContext) *DcimDevicesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim devices update params
func (o *DcimDevicesUpdateParams) SetData(data *models.WritableDeviceWithConfigContext) {
	o.Data = data
}

// WithID adds the id to the dcim devices update params
func (o *DcimDevicesUpdateParams) WithID(id strfmt.UUID) *DcimDevicesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices update params
func (o *DcimDevicesUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
