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

// NewDcimInventoryItemsListParams creates a new DcimInventoryItemsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimInventoryItemsListParams() *DcimInventoryItemsListParams {
	return &DcimInventoryItemsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInventoryItemsListParamsWithTimeout creates a new DcimInventoryItemsListParams object
// with the ability to set a timeout on a request.
func NewDcimInventoryItemsListParamsWithTimeout(timeout time.Duration) *DcimInventoryItemsListParams {
	return &DcimInventoryItemsListParams{
		timeout: timeout,
	}
}

// NewDcimInventoryItemsListParamsWithContext creates a new DcimInventoryItemsListParams object
// with the ability to set a context for a request.
func NewDcimInventoryItemsListParamsWithContext(ctx context.Context) *DcimInventoryItemsListParams {
	return &DcimInventoryItemsListParams{
		Context: ctx,
	}
}

// NewDcimInventoryItemsListParamsWithHTTPClient creates a new DcimInventoryItemsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimInventoryItemsListParamsWithHTTPClient(client *http.Client) *DcimInventoryItemsListParams {
	return &DcimInventoryItemsListParams{
		HTTPClient: client,
	}
}

/* DcimInventoryItemsListParams contains all the parameters to send to the API endpoint
   for the dcim inventory items list operation.

   Typically these are written to a http.Request.
*/
type DcimInventoryItemsListParams struct {

	// AssetTag.
	AssetTag *string

	// AssetTagIc.
	AssetTagIc *string

	// AssetTagIe.
	AssetTagIe *string

	// AssetTagIew.
	AssetTagIew *string

	// AssetTagIsw.
	AssetTagIsw *string

	// AssetTagn.
	AssetTagn *string

	// AssetTagNic.
	AssetTagNic *string

	// AssetTagNie.
	AssetTagNie *string

	// AssetTagNiew.
	AssetTagNiew *string

	// AssetTagNisw.
	AssetTagNisw *string

	// Device.
	Device *string

	// Devicen.
	Devicen *string

	// DeviceID.
	DeviceID *string

	// DeviceIDn.
	DeviceIDn *string

	// Discovered.
	Discovered *string

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

	// Manufacturer.
	Manufacturer *string

	// Manufacturern.
	Manufacturern *string

	// ManufacturerID.
	ManufacturerID *string

	// ManufacturerIDn.
	ManufacturerIDn *string

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

	// ParentID.
	ParentID *string

	// ParentIDn.
	ParentIDn *string

	// PartID.
	PartID *string

	// PartIDIc.
	PartIDIc *string

	// PartIDIe.
	PartIDIe *string

	// PartIDIew.
	PartIDIew *string

	// PartIDIsw.
	PartIDIsw *string

	// PartIDn.
	PartIDn *string

	// PartIDNic.
	PartIDNic *string

	// PartIDNie.
	PartIDNie *string

	// PartIDNiew.
	PartIDNiew *string

	// PartIDNisw.
	PartIDNisw *string

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

	// Serial.
	Serial *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim inventory items list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsListParams) WithDefaults() *DcimInventoryItemsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim inventory items list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimInventoryItemsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithTimeout(timeout time.Duration) *DcimInventoryItemsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithContext(ctx context.Context) *DcimInventoryItemsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithHTTPClient(client *http.Client) *DcimInventoryItemsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetTag adds the assetTag to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTag(assetTag *string) *DcimInventoryItemsListParams {
	o.SetAssetTag(assetTag)
	return o
}

// SetAssetTag adds the assetTag to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTag(assetTag *string) {
	o.AssetTag = assetTag
}

// WithAssetTagIc adds the assetTagIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagIc(assetTagIc *string) *DcimInventoryItemsListParams {
	o.SetAssetTagIc(assetTagIc)
	return o
}

// SetAssetTagIc adds the assetTagIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagIc(assetTagIc *string) {
	o.AssetTagIc = assetTagIc
}

// WithAssetTagIe adds the assetTagIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagIe(assetTagIe *string) *DcimInventoryItemsListParams {
	o.SetAssetTagIe(assetTagIe)
	return o
}

// SetAssetTagIe adds the assetTagIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagIe(assetTagIe *string) {
	o.AssetTagIe = assetTagIe
}

// WithAssetTagIew adds the assetTagIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagIew(assetTagIew *string) *DcimInventoryItemsListParams {
	o.SetAssetTagIew(assetTagIew)
	return o
}

// SetAssetTagIew adds the assetTagIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagIew(assetTagIew *string) {
	o.AssetTagIew = assetTagIew
}

// WithAssetTagIsw adds the assetTagIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagIsw(assetTagIsw *string) *DcimInventoryItemsListParams {
	o.SetAssetTagIsw(assetTagIsw)
	return o
}

// SetAssetTagIsw adds the assetTagIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagIsw(assetTagIsw *string) {
	o.AssetTagIsw = assetTagIsw
}

// WithAssetTagn adds the assetTagn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagn(assetTagn *string) *DcimInventoryItemsListParams {
	o.SetAssetTagn(assetTagn)
	return o
}

// SetAssetTagn adds the assetTagN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagn(assetTagn *string) {
	o.AssetTagn = assetTagn
}

// WithAssetTagNic adds the assetTagNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagNic(assetTagNic *string) *DcimInventoryItemsListParams {
	o.SetAssetTagNic(assetTagNic)
	return o
}

// SetAssetTagNic adds the assetTagNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagNic(assetTagNic *string) {
	o.AssetTagNic = assetTagNic
}

// WithAssetTagNie adds the assetTagNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagNie(assetTagNie *string) *DcimInventoryItemsListParams {
	o.SetAssetTagNie(assetTagNie)
	return o
}

// SetAssetTagNie adds the assetTagNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagNie(assetTagNie *string) {
	o.AssetTagNie = assetTagNie
}

// WithAssetTagNiew adds the assetTagNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagNiew(assetTagNiew *string) *DcimInventoryItemsListParams {
	o.SetAssetTagNiew(assetTagNiew)
	return o
}

// SetAssetTagNiew adds the assetTagNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagNiew(assetTagNiew *string) {
	o.AssetTagNiew = assetTagNiew
}

// WithAssetTagNisw adds the assetTagNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithAssetTagNisw(assetTagNisw *string) *DcimInventoryItemsListParams {
	o.SetAssetTagNisw(assetTagNisw)
	return o
}

// SetAssetTagNisw adds the assetTagNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetAssetTagNisw(assetTagNisw *string) {
	o.AssetTagNisw = assetTagNisw
}

// WithDevice adds the device to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDevice(device *string) *DcimInventoryItemsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDevicen(devicen *string) *DcimInventoryItemsListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDeviceID(deviceID *string) *DcimInventoryItemsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDeviceIDn(deviceIDn *string) *DcimInventoryItemsListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithDiscovered adds the discovered to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithDiscovered(discovered *string) *DcimInventoryItemsListParams {
	o.SetDiscovered(discovered)
	return o
}

// SetDiscovered adds the discovered to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetDiscovered(discovered *string) {
	o.Discovered = discovered
}

// WithID adds the id to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithID(id *string) *DcimInventoryItemsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDIc(iDIc *string) *DcimInventoryItemsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDIe(iDIe *string) *DcimInventoryItemsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDIew(iDIew *string) *DcimInventoryItemsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDIsw(iDIsw *string) *DcimInventoryItemsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDn(iDn *string) *DcimInventoryItemsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDNic(iDNic *string) *DcimInventoryItemsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDNie(iDNie *string) *DcimInventoryItemsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDNiew(iDNiew *string) *DcimInventoryItemsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithIDNisw(iDNisw *string) *DcimInventoryItemsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithLimit(limit *int64) *DcimInventoryItemsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithManufacturer(manufacturer *string) *DcimInventoryItemsListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturern adds the manufacturern to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithManufacturern(manufacturern *string) *DcimInventoryItemsListParams {
	o.SetManufacturern(manufacturern)
	return o
}

// SetManufacturern adds the manufacturerN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetManufacturern(manufacturern *string) {
	o.Manufacturern = manufacturern
}

// WithManufacturerID adds the manufacturerID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithManufacturerID(manufacturerID *string) *DcimInventoryItemsListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithManufacturerIDn adds the manufacturerIDn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithManufacturerIDn(manufacturerIDn *string) *DcimInventoryItemsListParams {
	o.SetManufacturerIDn(manufacturerIDn)
	return o
}

// SetManufacturerIDn adds the manufacturerIdN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetManufacturerIDn(manufacturerIDn *string) {
	o.ManufacturerIDn = manufacturerIDn
}

// WithName adds the name to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithName(name *string) *DcimInventoryItemsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameIc(nameIc *string) *DcimInventoryItemsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameIe(nameIe *string) *DcimInventoryItemsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameIew(nameIew *string) *DcimInventoryItemsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameIsw(nameIsw *string) *DcimInventoryItemsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNamen(namen *string) *DcimInventoryItemsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameNic(nameNic *string) *DcimInventoryItemsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameNie(nameNie *string) *DcimInventoryItemsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameNiew(nameNiew *string) *DcimInventoryItemsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithNameNisw(nameNisw *string) *DcimInventoryItemsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithOffset(offset *int64) *DcimInventoryItemsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParentID adds the parentID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithParentID(parentID *string) *DcimInventoryItemsListParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetParentID(parentID *string) {
	o.ParentID = parentID
}

// WithParentIDn adds the parentIDn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithParentIDn(parentIDn *string) *DcimInventoryItemsListParams {
	o.SetParentIDn(parentIDn)
	return o
}

// SetParentIDn adds the parentIdN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetParentIDn(parentIDn *string) {
	o.ParentIDn = parentIDn
}

// WithPartID adds the partID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartID(partID *string) *DcimInventoryItemsListParams {
	o.SetPartID(partID)
	return o
}

// SetPartID adds the partId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartID(partID *string) {
	o.PartID = partID
}

// WithPartIDIc adds the partIDIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDIc(partIDIc *string) *DcimInventoryItemsListParams {
	o.SetPartIDIc(partIDIc)
	return o
}

// SetPartIDIc adds the partIdIc to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDIc(partIDIc *string) {
	o.PartIDIc = partIDIc
}

// WithPartIDIe adds the partIDIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDIe(partIDIe *string) *DcimInventoryItemsListParams {
	o.SetPartIDIe(partIDIe)
	return o
}

// SetPartIDIe adds the partIdIe to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDIe(partIDIe *string) {
	o.PartIDIe = partIDIe
}

// WithPartIDIew adds the partIDIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDIew(partIDIew *string) *DcimInventoryItemsListParams {
	o.SetPartIDIew(partIDIew)
	return o
}

// SetPartIDIew adds the partIdIew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDIew(partIDIew *string) {
	o.PartIDIew = partIDIew
}

// WithPartIDIsw adds the partIDIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDIsw(partIDIsw *string) *DcimInventoryItemsListParams {
	o.SetPartIDIsw(partIDIsw)
	return o
}

// SetPartIDIsw adds the partIdIsw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDIsw(partIDIsw *string) {
	o.PartIDIsw = partIDIsw
}

// WithPartIDn adds the partIDn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDn(partIDn *string) *DcimInventoryItemsListParams {
	o.SetPartIDn(partIDn)
	return o
}

// SetPartIDn adds the partIdN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDn(partIDn *string) {
	o.PartIDn = partIDn
}

// WithPartIDNic adds the partIDNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDNic(partIDNic *string) *DcimInventoryItemsListParams {
	o.SetPartIDNic(partIDNic)
	return o
}

// SetPartIDNic adds the partIdNic to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDNic(partIDNic *string) {
	o.PartIDNic = partIDNic
}

// WithPartIDNie adds the partIDNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDNie(partIDNie *string) *DcimInventoryItemsListParams {
	o.SetPartIDNie(partIDNie)
	return o
}

// SetPartIDNie adds the partIdNie to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDNie(partIDNie *string) {
	o.PartIDNie = partIDNie
}

// WithPartIDNiew adds the partIDNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDNiew(partIDNiew *string) *DcimInventoryItemsListParams {
	o.SetPartIDNiew(partIDNiew)
	return o
}

// SetPartIDNiew adds the partIdNiew to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDNiew(partIDNiew *string) {
	o.PartIDNiew = partIDNiew
}

// WithPartIDNisw adds the partIDNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithPartIDNisw(partIDNisw *string) *DcimInventoryItemsListParams {
	o.SetPartIDNisw(partIDNisw)
	return o
}

// SetPartIDNisw adds the partIdNisw to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetPartIDNisw(partIDNisw *string) {
	o.PartIDNisw = partIDNisw
}

// WithQ adds the q to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithQ(q *string) *DcimInventoryItemsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithRegion(region *string) *DcimInventoryItemsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithRegionn(regionn *string) *DcimInventoryItemsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithRegionID(regionID *string) *DcimInventoryItemsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithRegionIDn(regionIDn *string) *DcimInventoryItemsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSerial adds the serial to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithSerial(serial *string) *DcimInventoryItemsListParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithSite adds the site to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithSite(site *string) *DcimInventoryItemsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithSiten(siten *string) *DcimInventoryItemsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithSiteID(siteID *string) *DcimInventoryItemsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithSiteIDn(siteIDn *string) *DcimInventoryItemsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithTag(tag *string) *DcimInventoryItemsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) WithTagn(tagn *string) *DcimInventoryItemsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim inventory items list params
func (o *DcimInventoryItemsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInventoryItemsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AssetTag != nil {

		// query param asset_tag
		var qrAssetTag string

		if o.AssetTag != nil {
			qrAssetTag = *o.AssetTag
		}
		qAssetTag := qrAssetTag
		if qAssetTag != "" {

			if err := r.SetQueryParam("asset_tag", qAssetTag); err != nil {
				return err
			}
		}
	}

	if o.AssetTagIc != nil {

		// query param asset_tag__ic
		var qrAssetTagIc string

		if o.AssetTagIc != nil {
			qrAssetTagIc = *o.AssetTagIc
		}
		qAssetTagIc := qrAssetTagIc
		if qAssetTagIc != "" {

			if err := r.SetQueryParam("asset_tag__ic", qAssetTagIc); err != nil {
				return err
			}
		}
	}

	if o.AssetTagIe != nil {

		// query param asset_tag__ie
		var qrAssetTagIe string

		if o.AssetTagIe != nil {
			qrAssetTagIe = *o.AssetTagIe
		}
		qAssetTagIe := qrAssetTagIe
		if qAssetTagIe != "" {

			if err := r.SetQueryParam("asset_tag__ie", qAssetTagIe); err != nil {
				return err
			}
		}
	}

	if o.AssetTagIew != nil {

		// query param asset_tag__iew
		var qrAssetTagIew string

		if o.AssetTagIew != nil {
			qrAssetTagIew = *o.AssetTagIew
		}
		qAssetTagIew := qrAssetTagIew
		if qAssetTagIew != "" {

			if err := r.SetQueryParam("asset_tag__iew", qAssetTagIew); err != nil {
				return err
			}
		}
	}

	if o.AssetTagIsw != nil {

		// query param asset_tag__isw
		var qrAssetTagIsw string

		if o.AssetTagIsw != nil {
			qrAssetTagIsw = *o.AssetTagIsw
		}
		qAssetTagIsw := qrAssetTagIsw
		if qAssetTagIsw != "" {

			if err := r.SetQueryParam("asset_tag__isw", qAssetTagIsw); err != nil {
				return err
			}
		}
	}

	if o.AssetTagn != nil {

		// query param asset_tag__n
		var qrAssetTagn string

		if o.AssetTagn != nil {
			qrAssetTagn = *o.AssetTagn
		}
		qAssetTagn := qrAssetTagn
		if qAssetTagn != "" {

			if err := r.SetQueryParam("asset_tag__n", qAssetTagn); err != nil {
				return err
			}
		}
	}

	if o.AssetTagNic != nil {

		// query param asset_tag__nic
		var qrAssetTagNic string

		if o.AssetTagNic != nil {
			qrAssetTagNic = *o.AssetTagNic
		}
		qAssetTagNic := qrAssetTagNic
		if qAssetTagNic != "" {

			if err := r.SetQueryParam("asset_tag__nic", qAssetTagNic); err != nil {
				return err
			}
		}
	}

	if o.AssetTagNie != nil {

		// query param asset_tag__nie
		var qrAssetTagNie string

		if o.AssetTagNie != nil {
			qrAssetTagNie = *o.AssetTagNie
		}
		qAssetTagNie := qrAssetTagNie
		if qAssetTagNie != "" {

			if err := r.SetQueryParam("asset_tag__nie", qAssetTagNie); err != nil {
				return err
			}
		}
	}

	if o.AssetTagNiew != nil {

		// query param asset_tag__niew
		var qrAssetTagNiew string

		if o.AssetTagNiew != nil {
			qrAssetTagNiew = *o.AssetTagNiew
		}
		qAssetTagNiew := qrAssetTagNiew
		if qAssetTagNiew != "" {

			if err := r.SetQueryParam("asset_tag__niew", qAssetTagNiew); err != nil {
				return err
			}
		}
	}

	if o.AssetTagNisw != nil {

		// query param asset_tag__nisw
		var qrAssetTagNisw string

		if o.AssetTagNisw != nil {
			qrAssetTagNisw = *o.AssetTagNisw
		}
		qAssetTagNisw := qrAssetTagNisw
		if qAssetTagNisw != "" {

			if err := r.SetQueryParam("asset_tag__nisw", qAssetTagNisw); err != nil {
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

	if o.Discovered != nil {

		// query param discovered
		var qrDiscovered string

		if o.Discovered != nil {
			qrDiscovered = *o.Discovered
		}
		qDiscovered := qrDiscovered
		if qDiscovered != "" {

			if err := r.SetQueryParam("discovered", qDiscovered); err != nil {
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

	if o.Manufacturer != nil {

		// query param manufacturer
		var qrManufacturer string

		if o.Manufacturer != nil {
			qrManufacturer = *o.Manufacturer
		}
		qManufacturer := qrManufacturer
		if qManufacturer != "" {

			if err := r.SetQueryParam("manufacturer", qManufacturer); err != nil {
				return err
			}
		}
	}

	if o.Manufacturern != nil {

		// query param manufacturer__n
		var qrManufacturern string

		if o.Manufacturern != nil {
			qrManufacturern = *o.Manufacturern
		}
		qManufacturern := qrManufacturern
		if qManufacturern != "" {

			if err := r.SetQueryParam("manufacturer__n", qManufacturern); err != nil {
				return err
			}
		}
	}

	if o.ManufacturerID != nil {

		// query param manufacturer_id
		var qrManufacturerID string

		if o.ManufacturerID != nil {
			qrManufacturerID = *o.ManufacturerID
		}
		qManufacturerID := qrManufacturerID
		if qManufacturerID != "" {

			if err := r.SetQueryParam("manufacturer_id", qManufacturerID); err != nil {
				return err
			}
		}
	}

	if o.ManufacturerIDn != nil {

		// query param manufacturer_id__n
		var qrManufacturerIDn string

		if o.ManufacturerIDn != nil {
			qrManufacturerIDn = *o.ManufacturerIDn
		}
		qManufacturerIDn := qrManufacturerIDn
		if qManufacturerIDn != "" {

			if err := r.SetQueryParam("manufacturer_id__n", qManufacturerIDn); err != nil {
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

	if o.ParentID != nil {

		// query param parent_id
		var qrParentID string

		if o.ParentID != nil {
			qrParentID = *o.ParentID
		}
		qParentID := qrParentID
		if qParentID != "" {

			if err := r.SetQueryParam("parent_id", qParentID); err != nil {
				return err
			}
		}
	}

	if o.ParentIDn != nil {

		// query param parent_id__n
		var qrParentIDn string

		if o.ParentIDn != nil {
			qrParentIDn = *o.ParentIDn
		}
		qParentIDn := qrParentIDn
		if qParentIDn != "" {

			if err := r.SetQueryParam("parent_id__n", qParentIDn); err != nil {
				return err
			}
		}
	}

	if o.PartID != nil {

		// query param part_id
		var qrPartID string

		if o.PartID != nil {
			qrPartID = *o.PartID
		}
		qPartID := qrPartID
		if qPartID != "" {

			if err := r.SetQueryParam("part_id", qPartID); err != nil {
				return err
			}
		}
	}

	if o.PartIDIc != nil {

		// query param part_id__ic
		var qrPartIDIc string

		if o.PartIDIc != nil {
			qrPartIDIc = *o.PartIDIc
		}
		qPartIDIc := qrPartIDIc
		if qPartIDIc != "" {

			if err := r.SetQueryParam("part_id__ic", qPartIDIc); err != nil {
				return err
			}
		}
	}

	if o.PartIDIe != nil {

		// query param part_id__ie
		var qrPartIDIe string

		if o.PartIDIe != nil {
			qrPartIDIe = *o.PartIDIe
		}
		qPartIDIe := qrPartIDIe
		if qPartIDIe != "" {

			if err := r.SetQueryParam("part_id__ie", qPartIDIe); err != nil {
				return err
			}
		}
	}

	if o.PartIDIew != nil {

		// query param part_id__iew
		var qrPartIDIew string

		if o.PartIDIew != nil {
			qrPartIDIew = *o.PartIDIew
		}
		qPartIDIew := qrPartIDIew
		if qPartIDIew != "" {

			if err := r.SetQueryParam("part_id__iew", qPartIDIew); err != nil {
				return err
			}
		}
	}

	if o.PartIDIsw != nil {

		// query param part_id__isw
		var qrPartIDIsw string

		if o.PartIDIsw != nil {
			qrPartIDIsw = *o.PartIDIsw
		}
		qPartIDIsw := qrPartIDIsw
		if qPartIDIsw != "" {

			if err := r.SetQueryParam("part_id__isw", qPartIDIsw); err != nil {
				return err
			}
		}
	}

	if o.PartIDn != nil {

		// query param part_id__n
		var qrPartIDn string

		if o.PartIDn != nil {
			qrPartIDn = *o.PartIDn
		}
		qPartIDn := qrPartIDn
		if qPartIDn != "" {

			if err := r.SetQueryParam("part_id__n", qPartIDn); err != nil {
				return err
			}
		}
	}

	if o.PartIDNic != nil {

		// query param part_id__nic
		var qrPartIDNic string

		if o.PartIDNic != nil {
			qrPartIDNic = *o.PartIDNic
		}
		qPartIDNic := qrPartIDNic
		if qPartIDNic != "" {

			if err := r.SetQueryParam("part_id__nic", qPartIDNic); err != nil {
				return err
			}
		}
	}

	if o.PartIDNie != nil {

		// query param part_id__nie
		var qrPartIDNie string

		if o.PartIDNie != nil {
			qrPartIDNie = *o.PartIDNie
		}
		qPartIDNie := qrPartIDNie
		if qPartIDNie != "" {

			if err := r.SetQueryParam("part_id__nie", qPartIDNie); err != nil {
				return err
			}
		}
	}

	if o.PartIDNiew != nil {

		// query param part_id__niew
		var qrPartIDNiew string

		if o.PartIDNiew != nil {
			qrPartIDNiew = *o.PartIDNiew
		}
		qPartIDNiew := qrPartIDNiew
		if qPartIDNiew != "" {

			if err := r.SetQueryParam("part_id__niew", qPartIDNiew); err != nil {
				return err
			}
		}
	}

	if o.PartIDNisw != nil {

		// query param part_id__nisw
		var qrPartIDNisw string

		if o.PartIDNisw != nil {
			qrPartIDNisw = *o.PartIDNisw
		}
		qPartIDNisw := qrPartIDNisw
		if qPartIDNisw != "" {

			if err := r.SetQueryParam("part_id__nisw", qPartIDNisw); err != nil {
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

	if o.Serial != nil {

		// query param serial
		var qrSerial string

		if o.Serial != nil {
			qrSerial = *o.Serial
		}
		qSerial := qrSerial
		if qSerial != "" {

			if err := r.SetQueryParam("serial", qSerial); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
