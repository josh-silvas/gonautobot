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

// NewDcimInterfacesListParams creates a new DcimInterfacesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInterfacesListParams() *DcimInterfacesListParams {
	return &DcimInterfacesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesListParamsWithTimeout creates a new DcimInterfacesListParams object
// with the ability to set a timeout on a request.
func NewDcimInterfacesListParamsWithTimeout(timeout time.Duration) *DcimInterfacesListParams {
	return &DcimInterfacesListParams{
		timeout: timeout,
	}
}

// NewDcimInterfacesListParamsWithContext creates a new DcimInterfacesListParams object
// with the ability to set a context for a request.
func NewDcimInterfacesListParamsWithContext(ctx context.Context) *DcimInterfacesListParams {
	return &DcimInterfacesListParams{
		Context: ctx,
	}
}

// NewDcimInterfacesListParamsWithHTTPClient creates a new DcimInterfacesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInterfacesListParamsWithHTTPClient(client *http.Client) *DcimInterfacesListParams {
	return &DcimInterfacesListParams{
		HTTPClient: client,
	}
}

/* DcimInterfacesListParams contains all the parameters to send to the API endpoint
   for the dcim interfaces list operation.

   Typically these are written to a http.Request.
*/
type DcimInterfacesListParams struct {

	// Cabled.
	Cabled *string

	// Connected.
	Connected *string

	// Description.
	Description *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

	// Device.
	Device *string

	// DeviceID.
	DeviceID *string

	// Enabled.
	Enabled *string

	// ID.
	ID *string

	// IDIc.
	IDIc *string

	// IDIe.
	IDIe *string

	// IDIew.
	IDIew *string

	// IDIsw.
	IDIsw *string

	// IDn.
	IDn *string

	// IDNic.
	IDNic *string

	// IDNie.
	IDNie *string

	// IDNiew.
	IDNiew *string

	// IDNisw.
	IDNisw *string

	// Kind.
	Kind *string

	// LagID.
	LagID *string

	// LagIDn.
	LagIDn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// MacAddress.
	MacAddress *string

	// MacAddressIc.
	MacAddressIc *string

	// MacAddressIe.
	MacAddressIe *string

	// MacAddressIew.
	MacAddressIew *string

	// MacAddressIsw.
	MacAddressIsw *string

	// MacAddressn.
	MacAddressn *string

	// MacAddressNic.
	MacAddressNic *string

	// MacAddressNie.
	MacAddressNie *string

	// MacAddressNiew.
	MacAddressNiew *string

	// MacAddressNisw.
	MacAddressNisw *string

	// MgmtOnly.
	MgmtOnly *string

	// Mode.
	Mode *string

	// Moden.
	Moden *string

	// Mtu.
	Mtu *string

	// MtuGt.
	MtuGt *string

	// MtuGte.
	MtuGte *string

	// MtuLt.
	MtuLt *string

	// MtuLte.
	MtuLte *string

	// Mtun.
	Mtun *string

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

	// Q.
	Q *string

	// Region.
	Region *string

	// Regionn.
	Regionn *string

	// RegionID.
	RegionID *string

	// RegionIDn.
	RegionIDn *string

	// Site.
	Site *string

	// Siten.
	Siten *string

	// SiteID.
	SiteID *string

	// SiteIDn.
	SiteIDn *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Type.
	Type *string

	// Typen.
	Typen *string

	// Vlan.
	Vlan *float64

	// VlanID.
	VlanID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim interfaces list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesListParams) WithDefaults() *DcimInterfacesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim interfaces list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInterfacesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithTimeout(timeout time.Duration) *DcimInterfacesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithContext(ctx context.Context) *DcimInterfacesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithHTTPClient(client *http.Client) *DcimInterfacesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithCabled(cabled *string) *DcimInterfacesListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnected adds the connected to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithConnected(connected *string) *DcimInterfacesListParams {
	o.SetConnected(connected)
	return o
}

// SetConnected adds the connected to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetConnected(connected *string) {
	o.Connected = connected
}

// WithDescription adds the description to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescription(description *string) *DcimInterfacesListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionIc(descriptionIc *string) *DcimInterfacesListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionIe(descriptionIe *string) *DcimInterfacesListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionIew(descriptionIew *string) *DcimInterfacesListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionIsw(descriptionIsw *string) *DcimInterfacesListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionn(descriptionn *string) *DcimInterfacesListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionNic(descriptionNic *string) *DcimInterfacesListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionNie(descriptionNie *string) *DcimInterfacesListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionNiew(descriptionNiew *string) *DcimInterfacesListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDescriptionNisw(descriptionNisw *string) *DcimInterfacesListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithDevice adds the device to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDevice(device *string) *DcimInterfacesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithDeviceID(deviceID *string) *DcimInterfacesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithEnabled adds the enabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithEnabled(enabled *string) *DcimInterfacesListParams {
	o.SetEnabled(enabled)
	return o
}

// SetEnabled adds the enabled to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetEnabled(enabled *string) {
	o.Enabled = enabled
}

// WithID adds the id to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithID(id *string) *DcimInterfacesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDIc(iDIc *string) *DcimInterfacesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDIe(iDIe *string) *DcimInterfacesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDIew(iDIew *string) *DcimInterfacesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDIsw(iDIsw *string) *DcimInterfacesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDn(iDn *string) *DcimInterfacesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDNic(iDNic *string) *DcimInterfacesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDNie(iDNie *string) *DcimInterfacesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDNiew(iDNiew *string) *DcimInterfacesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithIDNisw(iDNisw *string) *DcimInterfacesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithKind adds the kind to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithKind(kind *string) *DcimInterfacesListParams {
	o.SetKind(kind)
	return o
}

// SetKind adds the kind to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetKind(kind *string) {
	o.Kind = kind
}

// WithLagID adds the lagID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithLagID(lagID *string) *DcimInterfacesListParams {
	o.SetLagID(lagID)
	return o
}

// SetLagID adds the lagId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetLagID(lagID *string) {
	o.LagID = lagID
}

// WithLagIDn adds the lagIDn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithLagIDn(lagIDn *string) *DcimInterfacesListParams {
	o.SetLagIDn(lagIDn)
	return o
}

// SetLagIDn adds the lagIdN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetLagIDn(lagIDn *string) {
	o.LagIDn = lagIDn
}

// WithLimit adds the limit to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithLimit(limit *int64) *DcimInterfacesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMacAddress adds the macAddress to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddress(macAddress *string) *DcimInterfacesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMacAddressIc adds the macAddressIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressIc(macAddressIc *string) *DcimInterfacesListParams {
	o.SetMacAddressIc(macAddressIc)
	return o
}

// SetMacAddressIc adds the macAddressIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressIc(macAddressIc *string) {
	o.MacAddressIc = macAddressIc
}

// WithMacAddressIe adds the macAddressIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressIe(macAddressIe *string) *DcimInterfacesListParams {
	o.SetMacAddressIe(macAddressIe)
	return o
}

// SetMacAddressIe adds the macAddressIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressIe(macAddressIe *string) {
	o.MacAddressIe = macAddressIe
}

// WithMacAddressIew adds the macAddressIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressIew(macAddressIew *string) *DcimInterfacesListParams {
	o.SetMacAddressIew(macAddressIew)
	return o
}

// SetMacAddressIew adds the macAddressIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressIew(macAddressIew *string) {
	o.MacAddressIew = macAddressIew
}

// WithMacAddressIsw adds the macAddressIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressIsw(macAddressIsw *string) *DcimInterfacesListParams {
	o.SetMacAddressIsw(macAddressIsw)
	return o
}

// SetMacAddressIsw adds the macAddressIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressIsw(macAddressIsw *string) {
	o.MacAddressIsw = macAddressIsw
}

// WithMacAddressn adds the macAddressn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressn(macAddressn *string) *DcimInterfacesListParams {
	o.SetMacAddressn(macAddressn)
	return o
}

// SetMacAddressn adds the macAddressN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressn(macAddressn *string) {
	o.MacAddressn = macAddressn
}

// WithMacAddressNic adds the macAddressNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressNic(macAddressNic *string) *DcimInterfacesListParams {
	o.SetMacAddressNic(macAddressNic)
	return o
}

// SetMacAddressNic adds the macAddressNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressNic(macAddressNic *string) {
	o.MacAddressNic = macAddressNic
}

// WithMacAddressNie adds the macAddressNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressNie(macAddressNie *string) *DcimInterfacesListParams {
	o.SetMacAddressNie(macAddressNie)
	return o
}

// SetMacAddressNie adds the macAddressNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressNie(macAddressNie *string) {
	o.MacAddressNie = macAddressNie
}

// WithMacAddressNiew adds the macAddressNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressNiew(macAddressNiew *string) *DcimInterfacesListParams {
	o.SetMacAddressNiew(macAddressNiew)
	return o
}

// SetMacAddressNiew adds the macAddressNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressNiew(macAddressNiew *string) {
	o.MacAddressNiew = macAddressNiew
}

// WithMacAddressNisw adds the macAddressNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMacAddressNisw(macAddressNisw *string) *DcimInterfacesListParams {
	o.SetMacAddressNisw(macAddressNisw)
	return o
}

// SetMacAddressNisw adds the macAddressNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMacAddressNisw(macAddressNisw *string) {
	o.MacAddressNisw = macAddressNisw
}

// WithMgmtOnly adds the mgmtOnly to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMgmtOnly(mgmtOnly *string) *DcimInterfacesListParams {
	o.SetMgmtOnly(mgmtOnly)
	return o
}

// SetMgmtOnly adds the mgmtOnly to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMgmtOnly(mgmtOnly *string) {
	o.MgmtOnly = mgmtOnly
}

// WithMode adds the mode to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMode(mode *string) *DcimInterfacesListParams {
	o.SetMode(mode)
	return o
}

// SetMode adds the mode to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMode(mode *string) {
	o.Mode = mode
}

// WithModen adds the moden to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithModen(moden *string) *DcimInterfacesListParams {
	o.SetModen(moden)
	return o
}

// SetModen adds the modeN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetModen(moden *string) {
	o.Moden = moden
}

// WithMtu adds the mtu to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMtu(mtu *string) *DcimInterfacesListParams {
	o.SetMtu(mtu)
	return o
}

// SetMtu adds the mtu to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMtu(mtu *string) {
	o.Mtu = mtu
}

// WithMtuGt adds the mtuGt to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMtuGt(mtuGt *string) *DcimInterfacesListParams {
	o.SetMtuGt(mtuGt)
	return o
}

// SetMtuGt adds the mtuGt to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMtuGt(mtuGt *string) {
	o.MtuGt = mtuGt
}

// WithMtuGte adds the mtuGte to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMtuGte(mtuGte *string) *DcimInterfacesListParams {
	o.SetMtuGte(mtuGte)
	return o
}

// SetMtuGte adds the mtuGte to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMtuGte(mtuGte *string) {
	o.MtuGte = mtuGte
}

// WithMtuLt adds the mtuLt to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMtuLt(mtuLt *string) *DcimInterfacesListParams {
	o.SetMtuLt(mtuLt)
	return o
}

// SetMtuLt adds the mtuLt to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMtuLt(mtuLt *string) {
	o.MtuLt = mtuLt
}

// WithMtuLte adds the mtuLte to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMtuLte(mtuLte *string) *DcimInterfacesListParams {
	o.SetMtuLte(mtuLte)
	return o
}

// SetMtuLte adds the mtuLte to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMtuLte(mtuLte *string) {
	o.MtuLte = mtuLte
}

// WithMtun adds the mtun to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithMtun(mtun *string) *DcimInterfacesListParams {
	o.SetMtun(mtun)
	return o
}

// SetMtun adds the mtuN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetMtun(mtun *string) {
	o.Mtun = mtun
}

// WithName adds the name to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithName(name *string) *DcimInterfacesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameIc(nameIc *string) *DcimInterfacesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameIe(nameIe *string) *DcimInterfacesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameIew(nameIew *string) *DcimInterfacesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameIsw(nameIsw *string) *DcimInterfacesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNamen(namen *string) *DcimInterfacesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameNic(nameNic *string) *DcimInterfacesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameNie(nameNie *string) *DcimInterfacesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameNiew(nameNiew *string) *DcimInterfacesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithNameNisw(nameNisw *string) *DcimInterfacesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithOffset(offset *int64) *DcimInterfacesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithQ(q *string) *DcimInterfacesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithRegion(region *string) *DcimInterfacesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithRegionn(regionn *string) *DcimInterfacesListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithRegionID(regionID *string) *DcimInterfacesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithRegionIDn(regionIDn *string) *DcimInterfacesListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithSite(site *string) *DcimInterfacesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithSiten(siten *string) *DcimInterfacesListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithSiteID(siteID *string) *DcimInterfacesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithSiteIDn(siteIDn *string) *DcimInterfacesListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithTag(tag *string) *DcimInterfacesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithTagn(tagn *string) *DcimInterfacesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithType adds the typeVar to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithType(typeVar *string) *DcimInterfacesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithTypen(typen *string) *DcimInterfacesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WithVlan adds the vlan to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithVlan(vlan *float64) *DcimInterfacesListParams {
	o.SetVlan(vlan)
	return o
}

// SetVlan adds the vlan to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetVlan(vlan *float64) {
	o.Vlan = vlan
}

// WithVlanID adds the vlanID to the dcim interfaces list params
func (o *DcimInterfacesListParams) WithVlanID(vlanID *string) *DcimInterfacesListParams {
	o.SetVlanID(vlanID)
	return o
}

// SetVlanID adds the vlanId to the dcim interfaces list params
func (o *DcimInterfacesListParams) SetVlanID(vlanID *string) {
	o.VlanID = vlanID
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string

		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {

			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}
	}

	if o.Connected != nil {

		// query param connected
		var qrConnected string

		if o.Connected != nil {
			qrConnected = *o.Connected
		}
		qConnected := qrConnected
		if qConnected != "" {

			if err := r.SetQueryParam("connected", qConnected); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
				return err
			}
		}
	}

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

	if o.Enabled != nil {

		// query param enabled
		var qrEnabled string

		if o.Enabled != nil {
			qrEnabled = *o.Enabled
		}
		qEnabled := qrEnabled
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDIc != nil {

		// query param id__ic
		var qrIDIc string

		if o.IDIc != nil {
			qrIDIc = *o.IDIc
		}
		qIDIc := qrIDIc
		if qIDIc != "" {

			if err := r.SetQueryParam("id__ic", qIDIc); err != nil {
				return err
			}
		}
	}

	if o.IDIe != nil {

		// query param id__ie
		var qrIDIe string

		if o.IDIe != nil {
			qrIDIe = *o.IDIe
		}
		qIDIe := qrIDIe
		if qIDIe != "" {

			if err := r.SetQueryParam("id__ie", qIDIe); err != nil {
				return err
			}
		}
	}

	if o.IDIew != nil {

		// query param id__iew
		var qrIDIew string

		if o.IDIew != nil {
			qrIDIew = *o.IDIew
		}
		qIDIew := qrIDIew
		if qIDIew != "" {

			if err := r.SetQueryParam("id__iew", qIDIew); err != nil {
				return err
			}
		}
	}

	if o.IDIsw != nil {

		// query param id__isw
		var qrIDIsw string

		if o.IDIsw != nil {
			qrIDIsw = *o.IDIsw
		}
		qIDIsw := qrIDIsw
		if qIDIsw != "" {

			if err := r.SetQueryParam("id__isw", qIDIsw); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.IDNic != nil {

		// query param id__nic
		var qrIDNic string

		if o.IDNic != nil {
			qrIDNic = *o.IDNic
		}
		qIDNic := qrIDNic
		if qIDNic != "" {

			if err := r.SetQueryParam("id__nic", qIDNic); err != nil {
				return err
			}
		}
	}

	if o.IDNie != nil {

		// query param id__nie
		var qrIDNie string

		if o.IDNie != nil {
			qrIDNie = *o.IDNie
		}
		qIDNie := qrIDNie
		if qIDNie != "" {

			if err := r.SetQueryParam("id__nie", qIDNie); err != nil {
				return err
			}
		}
	}

	if o.IDNiew != nil {

		// query param id__niew
		var qrIDNiew string

		if o.IDNiew != nil {
			qrIDNiew = *o.IDNiew
		}
		qIDNiew := qrIDNiew
		if qIDNiew != "" {

			if err := r.SetQueryParam("id__niew", qIDNiew); err != nil {
				return err
			}
		}
	}

	if o.IDNisw != nil {

		// query param id__nisw
		var qrIDNisw string

		if o.IDNisw != nil {
			qrIDNisw = *o.IDNisw
		}
		qIDNisw := qrIDNisw
		if qIDNisw != "" {

			if err := r.SetQueryParam("id__nisw", qIDNisw); err != nil {
				return err
			}
		}
	}

	if o.Kind != nil {

		// query param kind
		var qrKind string

		if o.Kind != nil {
			qrKind = *o.Kind
		}
		qKind := qrKind
		if qKind != "" {

			if err := r.SetQueryParam("kind", qKind); err != nil {
				return err
			}
		}
	}

	if o.LagID != nil {

		// query param lag_id
		var qrLagID string

		if o.LagID != nil {
			qrLagID = *o.LagID
		}
		qLagID := qrLagID
		if qLagID != "" {

			if err := r.SetQueryParam("lag_id", qLagID); err != nil {
				return err
			}
		}
	}

	if o.LagIDn != nil {

		// query param lag_id__n
		var qrLagIDn string

		if o.LagIDn != nil {
			qrLagIDn = *o.LagIDn
		}
		qLagIDn := qrLagIDn
		if qLagIDn != "" {

			if err := r.SetQueryParam("lag_id__n", qLagIDn); err != nil {
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

	if o.MacAddress != nil {

		// query param mac_address
		var qrMacAddress string

		if o.MacAddress != nil {
			qrMacAddress = *o.MacAddress
		}
		qMacAddress := qrMacAddress
		if qMacAddress != "" {

			if err := r.SetQueryParam("mac_address", qMacAddress); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIc != nil {

		// query param mac_address__ic
		var qrMacAddressIc string

		if o.MacAddressIc != nil {
			qrMacAddressIc = *o.MacAddressIc
		}
		qMacAddressIc := qrMacAddressIc
		if qMacAddressIc != "" {

			if err := r.SetQueryParam("mac_address__ic", qMacAddressIc); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIe != nil {

		// query param mac_address__ie
		var qrMacAddressIe string

		if o.MacAddressIe != nil {
			qrMacAddressIe = *o.MacAddressIe
		}
		qMacAddressIe := qrMacAddressIe
		if qMacAddressIe != "" {

			if err := r.SetQueryParam("mac_address__ie", qMacAddressIe); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIew != nil {

		// query param mac_address__iew
		var qrMacAddressIew string

		if o.MacAddressIew != nil {
			qrMacAddressIew = *o.MacAddressIew
		}
		qMacAddressIew := qrMacAddressIew
		if qMacAddressIew != "" {

			if err := r.SetQueryParam("mac_address__iew", qMacAddressIew); err != nil {
				return err
			}
		}
	}

	if o.MacAddressIsw != nil {

		// query param mac_address__isw
		var qrMacAddressIsw string

		if o.MacAddressIsw != nil {
			qrMacAddressIsw = *o.MacAddressIsw
		}
		qMacAddressIsw := qrMacAddressIsw
		if qMacAddressIsw != "" {

			if err := r.SetQueryParam("mac_address__isw", qMacAddressIsw); err != nil {
				return err
			}
		}
	}

	if o.MacAddressn != nil {

		// query param mac_address__n
		var qrMacAddressn string

		if o.MacAddressn != nil {
			qrMacAddressn = *o.MacAddressn
		}
		qMacAddressn := qrMacAddressn
		if qMacAddressn != "" {

			if err := r.SetQueryParam("mac_address__n", qMacAddressn); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNic != nil {

		// query param mac_address__nic
		var qrMacAddressNic string

		if o.MacAddressNic != nil {
			qrMacAddressNic = *o.MacAddressNic
		}
		qMacAddressNic := qrMacAddressNic
		if qMacAddressNic != "" {

			if err := r.SetQueryParam("mac_address__nic", qMacAddressNic); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNie != nil {

		// query param mac_address__nie
		var qrMacAddressNie string

		if o.MacAddressNie != nil {
			qrMacAddressNie = *o.MacAddressNie
		}
		qMacAddressNie := qrMacAddressNie
		if qMacAddressNie != "" {

			if err := r.SetQueryParam("mac_address__nie", qMacAddressNie); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNiew != nil {

		// query param mac_address__niew
		var qrMacAddressNiew string

		if o.MacAddressNiew != nil {
			qrMacAddressNiew = *o.MacAddressNiew
		}
		qMacAddressNiew := qrMacAddressNiew
		if qMacAddressNiew != "" {

			if err := r.SetQueryParam("mac_address__niew", qMacAddressNiew); err != nil {
				return err
			}
		}
	}

	if o.MacAddressNisw != nil {

		// query param mac_address__nisw
		var qrMacAddressNisw string

		if o.MacAddressNisw != nil {
			qrMacAddressNisw = *o.MacAddressNisw
		}
		qMacAddressNisw := qrMacAddressNisw
		if qMacAddressNisw != "" {

			if err := r.SetQueryParam("mac_address__nisw", qMacAddressNisw); err != nil {
				return err
			}
		}
	}

	if o.MgmtOnly != nil {

		// query param mgmt_only
		var qrMgmtOnly string

		if o.MgmtOnly != nil {
			qrMgmtOnly = *o.MgmtOnly
		}
		qMgmtOnly := qrMgmtOnly
		if qMgmtOnly != "" {

			if err := r.SetQueryParam("mgmt_only", qMgmtOnly); err != nil {
				return err
			}
		}
	}

	if o.Mode != nil {

		// query param mode
		var qrMode string

		if o.Mode != nil {
			qrMode = *o.Mode
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}
	}

	if o.Moden != nil {

		// query param mode__n
		var qrModen string

		if o.Moden != nil {
			qrModen = *o.Moden
		}
		qModen := qrModen
		if qModen != "" {

			if err := r.SetQueryParam("mode__n", qModen); err != nil {
				return err
			}
		}
	}

	if o.Mtu != nil {

		// query param mtu
		var qrMtu string

		if o.Mtu != nil {
			qrMtu = *o.Mtu
		}
		qMtu := qrMtu
		if qMtu != "" {

			if err := r.SetQueryParam("mtu", qMtu); err != nil {
				return err
			}
		}
	}

	if o.MtuGt != nil {

		// query param mtu__gt
		var qrMtuGt string

		if o.MtuGt != nil {
			qrMtuGt = *o.MtuGt
		}
		qMtuGt := qrMtuGt
		if qMtuGt != "" {

			if err := r.SetQueryParam("mtu__gt", qMtuGt); err != nil {
				return err
			}
		}
	}

	if o.MtuGte != nil {

		// query param mtu__gte
		var qrMtuGte string

		if o.MtuGte != nil {
			qrMtuGte = *o.MtuGte
		}
		qMtuGte := qrMtuGte
		if qMtuGte != "" {

			if err := r.SetQueryParam("mtu__gte", qMtuGte); err != nil {
				return err
			}
		}
	}

	if o.MtuLt != nil {

		// query param mtu__lt
		var qrMtuLt string

		if o.MtuLt != nil {
			qrMtuLt = *o.MtuLt
		}
		qMtuLt := qrMtuLt
		if qMtuLt != "" {

			if err := r.SetQueryParam("mtu__lt", qMtuLt); err != nil {
				return err
			}
		}
	}

	if o.MtuLte != nil {

		// query param mtu__lte
		var qrMtuLte string

		if o.MtuLte != nil {
			qrMtuLte = *o.MtuLte
		}
		qMtuLte := qrMtuLte
		if qMtuLte != "" {

			if err := r.SetQueryParam("mtu__lte", qMtuLte); err != nil {
				return err
			}
		}
	}

	if o.Mtun != nil {

		// query param mtu__n
		var qrMtun string

		if o.Mtun != nil {
			qrMtun = *o.Mtun
		}
		qMtun := qrMtun
		if qMtun != "" {

			if err := r.SetQueryParam("mtu__n", qMtun); err != nil {
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

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Region != nil {

		// query param region
		var qrRegion string

		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {

			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}
	}

	if o.Regionn != nil {

		// query param region__n
		var qrRegionn string

		if o.Regionn != nil {
			qrRegionn = *o.Regionn
		}
		qRegionn := qrRegionn
		if qRegionn != "" {

			if err := r.SetQueryParam("region__n", qRegionn); err != nil {
				return err
			}
		}
	}

	if o.RegionID != nil {

		// query param region_id
		var qrRegionID string

		if o.RegionID != nil {
			qrRegionID = *o.RegionID
		}
		qRegionID := qrRegionID
		if qRegionID != "" {

			if err := r.SetQueryParam("region_id", qRegionID); err != nil {
				return err
			}
		}
	}

	if o.RegionIDn != nil {

		// query param region_id__n
		var qrRegionIDn string

		if o.RegionIDn != nil {
			qrRegionIDn = *o.RegionIDn
		}
		qRegionIDn := qrRegionIDn
		if qRegionIDn != "" {

			if err := r.SetQueryParam("region_id__n", qRegionIDn); err != nil {
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

	if o.Siten != nil {

		// query param site__n
		var qrSiten string

		if o.Siten != nil {
			qrSiten = *o.Siten
		}
		qSiten := qrSiten
		if qSiten != "" {

			if err := r.SetQueryParam("site__n", qSiten); err != nil {
				return err
			}
		}
	}

	if o.SiteID != nil {

		// query param site_id
		var qrSiteID string

		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {

			if err := r.SetQueryParam("site_id", qSiteID); err != nil {
				return err
			}
		}
	}

	if o.SiteIDn != nil {

		// query param site_id__n
		var qrSiteIDn string

		if o.SiteIDn != nil {
			qrSiteIDn = *o.SiteIDn
		}
		qSiteIDn := qrSiteIDn
		if qSiteIDn != "" {

			if err := r.SetQueryParam("site_id__n", qSiteIDn); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string

		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {

			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}
	}

	if o.Vlan != nil {

		// query param vlan
		var qrVlan float64

		if o.Vlan != nil {
			qrVlan = *o.Vlan
		}
		qVlan := swag.FormatFloat64(qrVlan)
		if qVlan != "" {

			if err := r.SetQueryParam("vlan", qVlan); err != nil {
				return err
			}
		}
	}

	if o.VlanID != nil {

		// query param vlan_id
		var qrVlanID string

		if o.VlanID != nil {
			qrVlanID = *o.VlanID
		}
		qVlanID := qrVlanID
		if qVlanID != "" {

			if err := r.SetQueryParam("vlan_id", qVlanID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
