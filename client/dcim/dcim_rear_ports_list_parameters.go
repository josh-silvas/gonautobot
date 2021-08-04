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

// NewDcimRearPortsListParams creates a new DcimRearPortsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRearPortsListParams() *DcimRearPortsListParams {
	return &DcimRearPortsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRearPortsListParamsWithTimeout creates a new DcimRearPortsListParams object
// with the ability to set a timeout on a request.
func NewDcimRearPortsListParamsWithTimeout(timeout time.Duration) *DcimRearPortsListParams {
	return &DcimRearPortsListParams{
		timeout: timeout,
	}
}

// NewDcimRearPortsListParamsWithContext creates a new DcimRearPortsListParams object
// with the ability to set a context for a request.
func NewDcimRearPortsListParamsWithContext(ctx context.Context) *DcimRearPortsListParams {
	return &DcimRearPortsListParams{
		Context: ctx,
	}
}

// NewDcimRearPortsListParamsWithHTTPClient creates a new DcimRearPortsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRearPortsListParamsWithHTTPClient(client *http.Client) *DcimRearPortsListParams {
	return &DcimRearPortsListParams{
		HTTPClient: client,
	}
}

/* DcimRearPortsListParams contains all the parameters to send to the API endpoint
   for the dcim rear ports list operation.

   Typically these are written to a http.Request.
*/
type DcimRearPortsListParams struct {

	// Cabled.
	Cabled *string

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

	// Positions.
	Positions *string

	// PositionsGt.
	PositionsGt *string

	// PositionsGte.
	PositionsGte *string

	// PositionsLt.
	PositionsLt *string

	// PositionsLte.
	PositionsLte *string

	// Positionsn.
	Positionsn *string

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

// WithDefaults hydrates default values in the dcim rear ports list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsListParams) WithDefaults() *DcimRearPortsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rear ports list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRearPortsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithTimeout(timeout time.Duration) *DcimRearPortsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithContext(ctx context.Context) *DcimRearPortsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithHTTPClient(client *http.Client) *DcimRearPortsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithCabled(cabled *string) *DcimRearPortsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithDescription adds the description to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescription(description *string) *DcimRearPortsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionIc(descriptionIc *string) *DcimRearPortsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionIe(descriptionIe *string) *DcimRearPortsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionIew(descriptionIew *string) *DcimRearPortsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionIsw(descriptionIsw *string) *DcimRearPortsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionn(descriptionn *string) *DcimRearPortsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionNic(descriptionNic *string) *DcimRearPortsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionNie(descriptionNie *string) *DcimRearPortsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionNiew(descriptionNiew *string) *DcimRearPortsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDescriptionNisw(descriptionNisw *string) *DcimRearPortsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithDevice adds the device to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDevice(device *string) *DcimRearPortsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDevicen(devicen *string) *DcimRearPortsListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDeviceID(deviceID *string) *DcimRearPortsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithDeviceIDn(deviceIDn *string) *DcimRearPortsListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithID adds the id to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithID(id *string) *DcimRearPortsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDIc(iDIc *string) *DcimRearPortsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDIe(iDIe *string) *DcimRearPortsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDIew(iDIew *string) *DcimRearPortsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDIsw(iDIsw *string) *DcimRearPortsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDn(iDn *string) *DcimRearPortsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDNic(iDNic *string) *DcimRearPortsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDNie(iDNie *string) *DcimRearPortsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDNiew(iDNiew *string) *DcimRearPortsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithIDNisw(iDNisw *string) *DcimRearPortsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithLimit(limit *int64) *DcimRearPortsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithName(name *string) *DcimRearPortsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameIc(nameIc *string) *DcimRearPortsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameIe(nameIe *string) *DcimRearPortsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameIew(nameIew *string) *DcimRearPortsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameIsw(nameIsw *string) *DcimRearPortsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNamen(namen *string) *DcimRearPortsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameNic(nameNic *string) *DcimRearPortsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameNie(nameNie *string) *DcimRearPortsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameNiew(nameNiew *string) *DcimRearPortsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithNameNisw(nameNisw *string) *DcimRearPortsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithOffset(offset *int64) *DcimRearPortsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPositions adds the positions to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithPositions(positions *string) *DcimRearPortsListParams {
	o.SetPositions(positions)
	return o
}

// SetPositions adds the positions to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetPositions(positions *string) {
	o.Positions = positions
}

// WithPositionsGt adds the positionsGt to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithPositionsGt(positionsGt *string) *DcimRearPortsListParams {
	o.SetPositionsGt(positionsGt)
	return o
}

// SetPositionsGt adds the positionsGt to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetPositionsGt(positionsGt *string) {
	o.PositionsGt = positionsGt
}

// WithPositionsGte adds the positionsGte to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithPositionsGte(positionsGte *string) *DcimRearPortsListParams {
	o.SetPositionsGte(positionsGte)
	return o
}

// SetPositionsGte adds the positionsGte to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetPositionsGte(positionsGte *string) {
	o.PositionsGte = positionsGte
}

// WithPositionsLt adds the positionsLt to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithPositionsLt(positionsLt *string) *DcimRearPortsListParams {
	o.SetPositionsLt(positionsLt)
	return o
}

// SetPositionsLt adds the positionsLt to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetPositionsLt(positionsLt *string) {
	o.PositionsLt = positionsLt
}

// WithPositionsLte adds the positionsLte to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithPositionsLte(positionsLte *string) *DcimRearPortsListParams {
	o.SetPositionsLte(positionsLte)
	return o
}

// SetPositionsLte adds the positionsLte to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetPositionsLte(positionsLte *string) {
	o.PositionsLte = positionsLte
}

// WithPositionsn adds the positionsn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithPositionsn(positionsn *string) *DcimRearPortsListParams {
	o.SetPositionsn(positionsn)
	return o
}

// SetPositionsn adds the positionsN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetPositionsn(positionsn *string) {
	o.Positionsn = positionsn
}

// WithQ adds the q to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithQ(q *string) *DcimRearPortsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithRegion(region *string) *DcimRearPortsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithRegionn(regionn *string) *DcimRearPortsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithRegionID(regionID *string) *DcimRearPortsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithRegionIDn(regionIDn *string) *DcimRearPortsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithSite(site *string) *DcimRearPortsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithSiten(siten *string) *DcimRearPortsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithSiteID(siteID *string) *DcimRearPortsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithSiteIDn(siteIDn *string) *DcimRearPortsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithTag(tag *string) *DcimRearPortsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithTagn(tagn *string) *DcimRearPortsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithType adds the typeVar to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithType(typeVar *string) *DcimRearPortsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim rear ports list params
func (o *DcimRearPortsListParams) WithTypen(typen *string) *DcimRearPortsListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim rear ports list params
func (o *DcimRearPortsListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRearPortsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Positions != nil {

		// query param positions
		var qrPositions string

		if o.Positions != nil {
			qrPositions = *o.Positions
		}
		qPositions := qrPositions
		if qPositions != "" {

			if err := r.SetQueryParam("positions", qPositions); err != nil {
				return err
			}
		}
	}

	if o.PositionsGt != nil {

		// query param positions__gt
		var qrPositionsGt string

		if o.PositionsGt != nil {
			qrPositionsGt = *o.PositionsGt
		}
		qPositionsGt := qrPositionsGt
		if qPositionsGt != "" {

			if err := r.SetQueryParam("positions__gt", qPositionsGt); err != nil {
				return err
			}
		}
	}

	if o.PositionsGte != nil {

		// query param positions__gte
		var qrPositionsGte string

		if o.PositionsGte != nil {
			qrPositionsGte = *o.PositionsGte
		}
		qPositionsGte := qrPositionsGte
		if qPositionsGte != "" {

			if err := r.SetQueryParam("positions__gte", qPositionsGte); err != nil {
				return err
			}
		}
	}

	if o.PositionsLt != nil {

		// query param positions__lt
		var qrPositionsLt string

		if o.PositionsLt != nil {
			qrPositionsLt = *o.PositionsLt
		}
		qPositionsLt := qrPositionsLt
		if qPositionsLt != "" {

			if err := r.SetQueryParam("positions__lt", qPositionsLt); err != nil {
				return err
			}
		}
	}

	if o.PositionsLte != nil {

		// query param positions__lte
		var qrPositionsLte string

		if o.PositionsLte != nil {
			qrPositionsLte = *o.PositionsLte
		}
		qPositionsLte := qrPositionsLte
		if qPositionsLte != "" {

			if err := r.SetQueryParam("positions__lte", qPositionsLte); err != nil {
				return err
			}
		}
	}

	if o.Positionsn != nil {

		// query param positions__n
		var qrPositionsn string

		if o.Positionsn != nil {
			qrPositionsn = *o.Positionsn
		}
		qPositionsn := qrPositionsn
		if qPositionsn != "" {

			if err := r.SetQueryParam("positions__n", qPositionsn); err != nil {
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
