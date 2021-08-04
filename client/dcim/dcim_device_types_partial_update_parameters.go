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

// NewDcimDeviceTypesPartialUpdateParams creates a new DcimDeviceTypesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceTypesPartialUpdateParams() *DcimDeviceTypesPartialUpdateParams {
	return &DcimDeviceTypesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesPartialUpdateParamsWithTimeout creates a new DcimDeviceTypesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceTypesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesPartialUpdateParams {
	return &DcimDeviceTypesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimDeviceTypesPartialUpdateParamsWithContext creates a new DcimDeviceTypesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimDeviceTypesPartialUpdateParamsWithContext(ctx context.Context) *DcimDeviceTypesPartialUpdateParams {
	return &DcimDeviceTypesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimDeviceTypesPartialUpdateParamsWithHTTPClient creates a new DcimDeviceTypesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceTypesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesPartialUpdateParams {
	return &DcimDeviceTypesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimDeviceTypesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim device types partial update operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceTypesPartialUpdateParams struct {

	// Data.
	Data *models.WritableDeviceType

	/* ID.

	   A UUID string identifying this device type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device types partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesPartialUpdateParams) WithDefaults() *DcimDeviceTypesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device types partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) WithContext(ctx context.Context) *DcimDeviceTypesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) WithData(data *models.WritableDeviceType) *DcimDeviceTypesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) SetData(data *models.WritableDeviceType) {
	o.Data = data
}

// WithID adds the id to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) WithID(id strfmt.UUID) *DcimDeviceTypesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device types partial update params
func (o *DcimDeviceTypesPartialUpdateParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
