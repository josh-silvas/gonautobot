package dcim

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDcimDeviceTypesDeleteParams creates a new DcimDeviceTypesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceTypesDeleteParams() *DcimDeviceTypesDeleteParams {
	return &DcimDeviceTypesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesDeleteParamsWithTimeout creates a new DcimDeviceTypesDeleteParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceTypesDeleteParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesDeleteParams {
	return &DcimDeviceTypesDeleteParams{
		timeout: timeout,
	}
}

// NewDcimDeviceTypesDeleteParamsWithContext creates a new DcimDeviceTypesDeleteParams object
// with the ability to set a context for a request.
func NewDcimDeviceTypesDeleteParamsWithContext(ctx context.Context) *DcimDeviceTypesDeleteParams {
	return &DcimDeviceTypesDeleteParams{
		Context: ctx,
	}
}

// NewDcimDeviceTypesDeleteParamsWithHTTPClient creates a new DcimDeviceTypesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceTypesDeleteParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesDeleteParams {
	return &DcimDeviceTypesDeleteParams{
		HTTPClient: client,
	}
}

/* DcimDeviceTypesDeleteParams contains all the parameters to send to the API endpoint
   for the dcim device types delete operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceTypesDeleteParams struct {

	/* ID.

	   A UUID string identifying this device type.

	   Format: uuid
	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesDeleteParams) WithDefaults() *DcimDeviceTypesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device types delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) WithContext(ctx context.Context) *DcimDeviceTypesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) WithID(id strfmt.UUID) *DcimDeviceTypesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device types delete params
func (o *DcimDeviceTypesDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
