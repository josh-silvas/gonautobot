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

// NewDcimConsoleConnectionsListParams creates a new DcimConsoleConnectionsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimConsoleConnectionsListParams() *DcimConsoleConnectionsListParams {
	return &DcimConsoleConnectionsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleConnectionsListParamsWithTimeout creates a new DcimConsoleConnectionsListParams object
// with the ability to set a timeout on a request.
func NewDcimConsoleConnectionsListParamsWithTimeout(timeout time.Duration) *DcimConsoleConnectionsListParams {
	return &DcimConsoleConnectionsListParams{
		timeout: timeout,
	}
}

// NewDcimConsoleConnectionsListParamsWithContext creates a new DcimConsoleConnectionsListParams object
// with the ability to set a context for a request.
func NewDcimConsoleConnectionsListParamsWithContext(ctx context.Context) *DcimConsoleConnectionsListParams {
	return &DcimConsoleConnectionsListParams{
		Context: ctx,
	}
}

// NewDcimConsoleConnectionsListParamsWithHTTPClient creates a new DcimConsoleConnectionsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimConsoleConnectionsListParamsWithHTTPClient(client *http.Client) *DcimConsoleConnectionsListParams {
	return &DcimConsoleConnectionsListParams{
		HTTPClient: client,
	}
}

/* DcimConsoleConnectionsListParams contains all the parameters to send to the API endpoint
   for the dcim console connections list operation.

   Typically these are written to a http.Request.
*/
type DcimConsoleConnectionsListParams struct {

	// Device.
	Device *string

	// DeviceID.
	DeviceID *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

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

// WithDefaults hydrates default values in the dcim console connections list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleConnectionsListParams) WithDefaults() *DcimConsoleConnectionsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim console connections list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimConsoleConnectionsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithTimeout(timeout time.Duration) *DcimConsoleConnectionsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithContext(ctx context.Context) *DcimConsoleConnectionsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithHTTPClient(client *http.Client) *DcimConsoleConnectionsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDevice adds the device to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithDevice(device *string) *DcimConsoleConnectionsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithDeviceID(deviceID *string) *DcimConsoleConnectionsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithLimit adds the limit to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithLimit(limit *int64) *DcimConsoleConnectionsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithName(name *string) *DcimConsoleConnectionsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameIc(nameIc *string) *DcimConsoleConnectionsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameIe(nameIe *string) *DcimConsoleConnectionsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameIew(nameIew *string) *DcimConsoleConnectionsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameIsw(nameIsw *string) *DcimConsoleConnectionsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNamen(namen *string) *DcimConsoleConnectionsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameNic(nameNic *string) *DcimConsoleConnectionsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameNie(nameNie *string) *DcimConsoleConnectionsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameNiew(nameNiew *string) *DcimConsoleConnectionsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithNameNisw(nameNisw *string) *DcimConsoleConnectionsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithOffset(offset *int64) *DcimConsoleConnectionsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSite adds the site to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) WithSite(site *string) *DcimConsoleConnectionsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim console connections list params
func (o *DcimConsoleConnectionsListParams) SetSite(site *string) {
	o.Site = site
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleConnectionsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
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
