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

// NewDcimPowerPortsListParams creates a new DcimPowerPortsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerPortsListParams() *DcimPowerPortsListParams {
	return &DcimPowerPortsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerPortsListParamsWithTimeout creates a new DcimPowerPortsListParams object
// with the ability to set a timeout on a request.
func NewDcimPowerPortsListParamsWithTimeout(timeout time.Duration) *DcimPowerPortsListParams {
	return &DcimPowerPortsListParams{
		timeout: timeout,
	}
}

// NewDcimPowerPortsListParamsWithContext creates a new DcimPowerPortsListParams object
// with the ability to set a context for a request.
func NewDcimPowerPortsListParamsWithContext(ctx context.Context) *DcimPowerPortsListParams {
	return &DcimPowerPortsListParams{
		Context: ctx,
	}
}

// NewDcimPowerPortsListParamsWithHTTPClient creates a new DcimPowerPortsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerPortsListParamsWithHTTPClient(client *http.Client) *DcimPowerPortsListParams {
	return &DcimPowerPortsListParams{
		HTTPClient: client,
	}
}

/* DcimPowerPortsListParams contains all the parameters to send to the API endpoint
   for the dcim power ports list operation.

   Typically these are written to a http.Request.
*/
type DcimPowerPortsListParams struct {

	// AllocatedDraw.
	AllocatedDraw *string

	// AllocatedDrawGt.
	AllocatedDrawGt *string

	// AllocatedDrawGte.
	AllocatedDrawGte *string

	// AllocatedDrawLt.
	AllocatedDrawLt *string

	// AllocatedDrawLte.
	AllocatedDrawLte *string

	// AllocatedDrawn.
	AllocatedDrawn *string

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

	// Devicen.
	Devicen *string

	// DeviceID.
	DeviceID *string

	// DeviceIDn.
	DeviceIDn *string

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

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// MaximumDraw.
	MaximumDraw *string

	// MaximumDrawGt.
	MaximumDrawGt *string

	// MaximumDrawGte.
	MaximumDrawGte *string

	// MaximumDrawLt.
	MaximumDrawLt *string

	// MaximumDrawLte.
	MaximumDrawLte *string

	// MaximumDrawn.
	MaximumDrawn *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power ports list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsListParams) WithDefaults() *DcimPowerPortsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power ports list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerPortsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithTimeout(timeout time.Duration) *DcimPowerPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithContext(ctx context.Context) *DcimPowerPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithHTTPClient(client *http.Client) *DcimPowerPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllocatedDraw adds the allocatedDraw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithAllocatedDraw(allocatedDraw *string) *DcimPowerPortsListParams {
	o.SetAllocatedDraw(allocatedDraw)
	return o
}

// SetAllocatedDraw adds the allocatedDraw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetAllocatedDraw(allocatedDraw *string) {
	o.AllocatedDraw = allocatedDraw
}

// WithAllocatedDrawGt adds the allocatedDrawGt to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithAllocatedDrawGt(allocatedDrawGt *string) *DcimPowerPortsListParams {
	o.SetAllocatedDrawGt(allocatedDrawGt)
	return o
}

// SetAllocatedDrawGt adds the allocatedDrawGt to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetAllocatedDrawGt(allocatedDrawGt *string) {
	o.AllocatedDrawGt = allocatedDrawGt
}

// WithAllocatedDrawGte adds the allocatedDrawGte to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithAllocatedDrawGte(allocatedDrawGte *string) *DcimPowerPortsListParams {
	o.SetAllocatedDrawGte(allocatedDrawGte)
	return o
}

// SetAllocatedDrawGte adds the allocatedDrawGte to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetAllocatedDrawGte(allocatedDrawGte *string) {
	o.AllocatedDrawGte = allocatedDrawGte
}

// WithAllocatedDrawLt adds the allocatedDrawLt to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithAllocatedDrawLt(allocatedDrawLt *string) *DcimPowerPortsListParams {
	o.SetAllocatedDrawLt(allocatedDrawLt)
	return o
}

// SetAllocatedDrawLt adds the allocatedDrawLt to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetAllocatedDrawLt(allocatedDrawLt *string) {
	o.AllocatedDrawLt = allocatedDrawLt
}

// WithAllocatedDrawLte adds the allocatedDrawLte to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithAllocatedDrawLte(allocatedDrawLte *string) *DcimPowerPortsListParams {
	o.SetAllocatedDrawLte(allocatedDrawLte)
	return o
}

// SetAllocatedDrawLte adds the allocatedDrawLte to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetAllocatedDrawLte(allocatedDrawLte *string) {
	o.AllocatedDrawLte = allocatedDrawLte
}

// WithAllocatedDrawn adds the allocatedDrawn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithAllocatedDrawn(allocatedDrawn *string) *DcimPowerPortsListParams {
	o.SetAllocatedDrawn(allocatedDrawn)
	return o
}

// SetAllocatedDrawn adds the allocatedDrawN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetAllocatedDrawn(allocatedDrawn *string) {
	o.AllocatedDrawn = allocatedDrawn
}

// WithCabled adds the cabled to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithCabled(cabled *string) *DcimPowerPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnected adds the connected to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithConnected(connected *string) *DcimPowerPortsListParams {
	o.SetConnected(connected)
	return o
}

// SetConnected adds the connected to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetConnected(connected *string) {
	o.Connected = connected
}

// WithDescription adds the description to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescription(description *string) *DcimPowerPortsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionIc(descriptionIc *string) *DcimPowerPortsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionIe(descriptionIe *string) *DcimPowerPortsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionIew(descriptionIew *string) *DcimPowerPortsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionIsw(descriptionIsw *string) *DcimPowerPortsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionn(descriptionn *string) *DcimPowerPortsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionNic(descriptionNic *string) *DcimPowerPortsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionNie(descriptionNie *string) *DcimPowerPortsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionNiew(descriptionNiew *string) *DcimPowerPortsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDescriptionNisw(descriptionNisw *string) *DcimPowerPortsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithDevice adds the device to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDevice(device *string) *DcimPowerPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDevicen(devicen *string) *DcimPowerPortsListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDeviceID(deviceID *string) *DcimPowerPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithDeviceIDn(deviceIDn *string) *DcimPowerPortsListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithID adds the id to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithID(id *string) *DcimPowerPortsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDIc(iDIc *string) *DcimPowerPortsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDIe(iDIe *string) *DcimPowerPortsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDIew(iDIew *string) *DcimPowerPortsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDIsw(iDIsw *string) *DcimPowerPortsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDn(iDn *string) *DcimPowerPortsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDNic(iDNic *string) *DcimPowerPortsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDNie(iDNie *string) *DcimPowerPortsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDNiew(iDNiew *string) *DcimPowerPortsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithIDNisw(iDNisw *string) *DcimPowerPortsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithLimit(limit *int64) *DcimPowerPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaximumDraw adds the maximumDraw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithMaximumDraw(maximumDraw *string) *DcimPowerPortsListParams {
	o.SetMaximumDraw(maximumDraw)
	return o
}

// SetMaximumDraw adds the maximumDraw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetMaximumDraw(maximumDraw *string) {
	o.MaximumDraw = maximumDraw
}

// WithMaximumDrawGt adds the maximumDrawGt to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithMaximumDrawGt(maximumDrawGt *string) *DcimPowerPortsListParams {
	o.SetMaximumDrawGt(maximumDrawGt)
	return o
}

// SetMaximumDrawGt adds the maximumDrawGt to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetMaximumDrawGt(maximumDrawGt *string) {
	o.MaximumDrawGt = maximumDrawGt
}

// WithMaximumDrawGte adds the maximumDrawGte to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithMaximumDrawGte(maximumDrawGte *string) *DcimPowerPortsListParams {
	o.SetMaximumDrawGte(maximumDrawGte)
	return o
}

// SetMaximumDrawGte adds the maximumDrawGte to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetMaximumDrawGte(maximumDrawGte *string) {
	o.MaximumDrawGte = maximumDrawGte
}

// WithMaximumDrawLt adds the maximumDrawLt to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithMaximumDrawLt(maximumDrawLt *string) *DcimPowerPortsListParams {
	o.SetMaximumDrawLt(maximumDrawLt)
	return o
}

// SetMaximumDrawLt adds the maximumDrawLt to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetMaximumDrawLt(maximumDrawLt *string) {
	o.MaximumDrawLt = maximumDrawLt
}

// WithMaximumDrawLte adds the maximumDrawLte to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithMaximumDrawLte(maximumDrawLte *string) *DcimPowerPortsListParams {
	o.SetMaximumDrawLte(maximumDrawLte)
	return o
}

// SetMaximumDrawLte adds the maximumDrawLte to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetMaximumDrawLte(maximumDrawLte *string) {
	o.MaximumDrawLte = maximumDrawLte
}

// WithMaximumDrawn adds the maximumDrawn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithMaximumDrawn(maximumDrawn *string) *DcimPowerPortsListParams {
	o.SetMaximumDrawn(maximumDrawn)
	return o
}

// SetMaximumDrawn adds the maximumDrawN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetMaximumDrawn(maximumDrawn *string) {
	o.MaximumDrawn = maximumDrawn
}

// WithName adds the name to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithName(name *string) *DcimPowerPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameIc(nameIc *string) *DcimPowerPortsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameIe(nameIe *string) *DcimPowerPortsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameIew(nameIew *string) *DcimPowerPortsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameIsw(nameIsw *string) *DcimPowerPortsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNamen(namen *string) *DcimPowerPortsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameNic(nameNic *string) *DcimPowerPortsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameNie(nameNie *string) *DcimPowerPortsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameNiew(nameNiew *string) *DcimPowerPortsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithNameNisw(nameNisw *string) *DcimPowerPortsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithOffset(offset *int64) *DcimPowerPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithQ(q *string) *DcimPowerPortsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithRegion(region *string) *DcimPowerPortsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithRegionn(regionn *string) *DcimPowerPortsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithRegionID(regionID *string) *DcimPowerPortsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithRegionIDn(regionIDn *string) *DcimPowerPortsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithSite(site *string) *DcimPowerPortsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithSiten(siten *string) *DcimPowerPortsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithSiteID(siteID *string) *DcimPowerPortsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithSiteIDn(siteIDn *string) *DcimPowerPortsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithTag(tag *string) *DcimPowerPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithTagn(tagn *string) *DcimPowerPortsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithType adds the typeVar to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithType(typeVar *string) *DcimPowerPortsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim power ports list params
func (o *DcimPowerPortsListParams) WithTypen(typen *string) *DcimPowerPortsListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim power ports list params
func (o *DcimPowerPortsListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AllocatedDraw != nil {

		// query param allocated_draw
		var qrAllocatedDraw string

		if o.AllocatedDraw != nil {
			qrAllocatedDraw = *o.AllocatedDraw
		}
		qAllocatedDraw := qrAllocatedDraw
		if qAllocatedDraw != "" {

			if err := r.SetQueryParam("allocated_draw", qAllocatedDraw); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawGt != nil {

		// query param allocated_draw__gt
		var qrAllocatedDrawGt string

		if o.AllocatedDrawGt != nil {
			qrAllocatedDrawGt = *o.AllocatedDrawGt
		}
		qAllocatedDrawGt := qrAllocatedDrawGt
		if qAllocatedDrawGt != "" {

			if err := r.SetQueryParam("allocated_draw__gt", qAllocatedDrawGt); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawGte != nil {

		// query param allocated_draw__gte
		var qrAllocatedDrawGte string

		if o.AllocatedDrawGte != nil {
			qrAllocatedDrawGte = *o.AllocatedDrawGte
		}
		qAllocatedDrawGte := qrAllocatedDrawGte
		if qAllocatedDrawGte != "" {

			if err := r.SetQueryParam("allocated_draw__gte", qAllocatedDrawGte); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawLt != nil {

		// query param allocated_draw__lt
		var qrAllocatedDrawLt string

		if o.AllocatedDrawLt != nil {
			qrAllocatedDrawLt = *o.AllocatedDrawLt
		}
		qAllocatedDrawLt := qrAllocatedDrawLt
		if qAllocatedDrawLt != "" {

			if err := r.SetQueryParam("allocated_draw__lt", qAllocatedDrawLt); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawLte != nil {

		// query param allocated_draw__lte
		var qrAllocatedDrawLte string

		if o.AllocatedDrawLte != nil {
			qrAllocatedDrawLte = *o.AllocatedDrawLte
		}
		qAllocatedDrawLte := qrAllocatedDrawLte
		if qAllocatedDrawLte != "" {

			if err := r.SetQueryParam("allocated_draw__lte", qAllocatedDrawLte); err != nil {
				return err
			}
		}
	}

	if o.AllocatedDrawn != nil {

		// query param allocated_draw__n
		var qrAllocatedDrawn string

		if o.AllocatedDrawn != nil {
			qrAllocatedDrawn = *o.AllocatedDrawn
		}
		qAllocatedDrawn := qrAllocatedDrawn
		if qAllocatedDrawn != "" {

			if err := r.SetQueryParam("allocated_draw__n", qAllocatedDrawn); err != nil {
				return err
			}
		}
	}

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

	if o.Devicen != nil {

		// query param device__n
		var qrDevicen string

		if o.Devicen != nil {
			qrDevicen = *o.Devicen
		}
		qDevicen := qrDevicen
		if qDevicen != "" {

			if err := r.SetQueryParam("device__n", qDevicen); err != nil {
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

	if o.DeviceIDn != nil {

		// query param device_id__n
		var qrDeviceIDn string

		if o.DeviceIDn != nil {
			qrDeviceIDn = *o.DeviceIDn
		}
		qDeviceIDn := qrDeviceIDn
		if qDeviceIDn != "" {

			if err := r.SetQueryParam("device_id__n", qDeviceIDn); err != nil {
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

	if o.MaximumDraw != nil {

		// query param maximum_draw
		var qrMaximumDraw string

		if o.MaximumDraw != nil {
			qrMaximumDraw = *o.MaximumDraw
		}
		qMaximumDraw := qrMaximumDraw
		if qMaximumDraw != "" {

			if err := r.SetQueryParam("maximum_draw", qMaximumDraw); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawGt != nil {

		// query param maximum_draw__gt
		var qrMaximumDrawGt string

		if o.MaximumDrawGt != nil {
			qrMaximumDrawGt = *o.MaximumDrawGt
		}
		qMaximumDrawGt := qrMaximumDrawGt
		if qMaximumDrawGt != "" {

			if err := r.SetQueryParam("maximum_draw__gt", qMaximumDrawGt); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawGte != nil {

		// query param maximum_draw__gte
		var qrMaximumDrawGte string

		if o.MaximumDrawGte != nil {
			qrMaximumDrawGte = *o.MaximumDrawGte
		}
		qMaximumDrawGte := qrMaximumDrawGte
		if qMaximumDrawGte != "" {

			if err := r.SetQueryParam("maximum_draw__gte", qMaximumDrawGte); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawLt != nil {

		// query param maximum_draw__lt
		var qrMaximumDrawLt string

		if o.MaximumDrawLt != nil {
			qrMaximumDrawLt = *o.MaximumDrawLt
		}
		qMaximumDrawLt := qrMaximumDrawLt
		if qMaximumDrawLt != "" {

			if err := r.SetQueryParam("maximum_draw__lt", qMaximumDrawLt); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawLte != nil {

		// query param maximum_draw__lte
		var qrMaximumDrawLte string

		if o.MaximumDrawLte != nil {
			qrMaximumDrawLte = *o.MaximumDrawLte
		}
		qMaximumDrawLte := qrMaximumDrawLte
		if qMaximumDrawLte != "" {

			if err := r.SetQueryParam("maximum_draw__lte", qMaximumDrawLte); err != nil {
				return err
			}
		}
	}

	if o.MaximumDrawn != nil {

		// query param maximum_draw__n
		var qrMaximumDrawn string

		if o.MaximumDrawn != nil {
			qrMaximumDrawn = *o.MaximumDrawn
		}
		qMaximumDrawn := qrMaximumDrawn
		if qMaximumDrawn != "" {

			if err := r.SetQueryParam("maximum_draw__n", qMaximumDrawn); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
