package dcim

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

// NewDcimInterfaceConnectionsListParams creates a new DcimInterfaceConnectionsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfaceConnectionsListParams() *DcimInterfaceConnectionsListParams {
	return &DcimInterfaceConnectionsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfaceConnectionsListParamsWithTimeout creates a new DcimInterfaceConnectionsListParams object
// with the ability to set a timeout on a request.
func NewDcimInterfaceConnectionsListParamsWithTimeout(timeout time.Duration) *DcimInterfaceConnectionsListParams {
	return &DcimInterfaceConnectionsListParams{
		timeout: timeout,
	}
}

// NewDcimInterfaceConnectionsListParamsWithContext creates a new DcimInterfaceConnectionsListParams object
// with the ability to set a context for a request.
func NewDcimInterfaceConnectionsListParamsWithContext(ctx context.Context) *DcimInterfaceConnectionsListParams {
	return &DcimInterfaceConnectionsListParams{
		Context: ctx,
	}
}

// NewDcimInterfaceConnectionsListParamsWithHTTPClient creates a new DcimInterfaceConnectionsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfaceConnectionsListParamsWithHTTPClient(client *http.Client) *DcimInterfaceConnectionsListParams {
	return &DcimInterfaceConnectionsListParams{
		HTTPClient: client,
	}
}

/* DcimInterfaceConnectionsListParams contains all the parameters to send to the API endpoint
   for the dcim interface connections list operation.

   Typically these are written to a http.Request.
*/
type DcimInterfaceConnectionsListParams struct {

	// Device.
	Device *string

	// DeviceID.
	DeviceID *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Site.
	Site *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interface connections list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceConnectionsListParams) WithDefaults() *DcimInterfaceConnectionsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interface connections list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfaceConnectionsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithTimeout(timeout time.Duration) *DcimInterfaceConnectionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithContext(ctx context.Context) *DcimInterfaceConnectionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithHTTPClient(client *http.Client) *DcimInterfaceConnectionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithDevice(device *string) *DcimInterfaceConnectionsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithDeviceID(deviceID *string) *DcimInterfaceConnectionsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithLimit(limit *int64) *DcimInterfaceConnectionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithOffset(offset *int64) *DcimInterfaceConnectionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSite adds the site to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) WithSite(site *string) *DcimInterfaceConnectionsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim interface connections list params
func (o *DcimInterfaceConnectionsListParams) SetSite(site *string) {
	o.Site = site
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfaceConnectionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Device != nil {

		// query param device
		var qrDevice string

		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {

			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}
	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string

		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {

			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
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

	if o.Site != nil {

		// query param site
		var qrSite string

		if o.Site != nil {
			qrSite = *o.Site
		}
		qSite := qrSite
		if qSite != "" {

			if err := r.SetQueryParam("site", qSite); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
