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

// NewDcimConnectedDeviceListParams creates a new DcimConnectedDeviceListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConnectedDeviceListParams() *DcimConnectedDeviceListParams {
	return &DcimConnectedDeviceListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConnectedDeviceListParamsWithTimeout creates a new DcimConnectedDeviceListParams object
// with the ability to set a timeout on a request.
func NewDcimConnectedDeviceListParamsWithTimeout(timeout time.Duration) *DcimConnectedDeviceListParams {
	return &DcimConnectedDeviceListParams{
		timeout: timeout,
	}
}

// NewDcimConnectedDeviceListParamsWithContext creates a new DcimConnectedDeviceListParams object
// with the ability to set a context for a request.
func NewDcimConnectedDeviceListParamsWithContext(ctx context.Context) *DcimConnectedDeviceListParams {
	return &DcimConnectedDeviceListParams{
		Context: ctx,
	}
}

// NewDcimConnectedDeviceListParamsWithHTTPClient creates a new DcimConnectedDeviceListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConnectedDeviceListParamsWithHTTPClient(client *http.Client) *DcimConnectedDeviceListParams {
	return &DcimConnectedDeviceListParams{
		HTTPClient: client,
	}
}

/* DcimConnectedDeviceListParams contains all the parameters to send to the API endpoint
   for the dcim connected device list operation.

   Typically these are written to a http.Request.
*/
type DcimConnectedDeviceListParams struct {

	/* PeerDevice.

	   The name of the peer device
	*/
	PeerDevice string

	/* PeerInterface.

	   The name of the peer interface
	*/
	PeerInterface string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim connected device list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConnectedDeviceListParams) WithDefaults() *DcimConnectedDeviceListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim connected device list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConnectedDeviceListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) WithTimeout(timeout time.Duration) *DcimConnectedDeviceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) WithContext(ctx context.Context) *DcimConnectedDeviceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) WithHTTPClient(client *http.Client) *DcimConnectedDeviceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPeerDevice adds the peerDevice to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) WithPeerDevice(peerDevice string) *DcimConnectedDeviceListParams {
	o.SetPeerDevice(peerDevice)
	return o
}

// SetPeerDevice adds the peerDevice to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) SetPeerDevice(peerDevice string) {
	o.PeerDevice = peerDevice
}

// WithPeerInterface adds the peerInterface to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) WithPeerInterface(peerInterface string) *DcimConnectedDeviceListParams {
	o.SetPeerInterface(peerInterface)
	return o
}

// SetPeerInterface adds the peerInterface to the dcim connected device list params
func (o *DcimConnectedDeviceListParams) SetPeerInterface(peerInterface string) {
	o.PeerInterface = peerInterface
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConnectedDeviceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param peer_device
	qrPeerDevice := o.PeerDevice
	qPeerDevice := qrPeerDevice
	if qPeerDevice != "" {

		if err := r.SetQueryParam("peer_device", qPeerDevice); err != nil {
			return err
		}
	}

	// query param peer_interface
	qrPeerInterface := o.PeerInterface
	qPeerInterface := qrPeerInterface
	if qPeerInterface != "" {

		if err := r.SetQueryParam("peer_interface", qPeerInterface); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
