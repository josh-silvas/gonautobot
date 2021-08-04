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

// NewDcimPowerOutletsListParams creates a new DcimPowerOutletsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerOutletsListParams() *DcimPowerOutletsListParams {
	return &DcimPowerOutletsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletsListParamsWithTimeout creates a new DcimPowerOutletsListParams object
// with the ability to set a timeout on a request.
func NewDcimPowerOutletsListParamsWithTimeout(timeout time.Duration) *DcimPowerOutletsListParams {
	return &DcimPowerOutletsListParams{
		timeout: timeout,
	}
}

// NewDcimPowerOutletsListParamsWithContext creates a new DcimPowerOutletsListParams object
// with the ability to set a context for a request.
func NewDcimPowerOutletsListParamsWithContext(ctx context.Context) *DcimPowerOutletsListParams {
	return &DcimPowerOutletsListParams{
		Context: ctx,
	}
}

// NewDcimPowerOutletsListParamsWithHTTPClient creates a new DcimPowerOutletsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerOutletsListParamsWithHTTPClient(client *http.Client) *DcimPowerOutletsListParams {
	return &DcimPowerOutletsListParams{
		HTTPClient: client,
	}
}

/* DcimPowerOutletsListParams contains all the parameters to send to the API endpoint
   for the dcim power outlets list operation.

   Typically these are written to a http.Request.
*/
type DcimPowerOutletsListParams struct {

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

	// FeedLeg.
	FeedLeg *string

	// FeedLegn.
	FeedLegn *string

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

// WithDefaults hydrates default values in the dcim power outlets list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsListParams) WithDefaults() *DcimPowerOutletsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power outlets list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerOutletsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithTimeout(timeout time.Duration) *DcimPowerOutletsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithContext(ctx context.Context) *DcimPowerOutletsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithHTTPClient(client *http.Client) *DcimPowerOutletsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCabled adds the cabled to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithCabled(cabled *string) *DcimPowerOutletsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnected adds the connected to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithConnected(connected *string) *DcimPowerOutletsListParams {
	o.SetConnected(connected)
	return o
}

// SetConnected adds the connected to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetConnected(connected *string) {
	o.Connected = connected
}

// WithDescription adds the description to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescription(description *string) *DcimPowerOutletsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionIc(descriptionIc *string) *DcimPowerOutletsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionIe(descriptionIe *string) *DcimPowerOutletsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionIew(descriptionIew *string) *DcimPowerOutletsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionIsw(descriptionIsw *string) *DcimPowerOutletsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionn(descriptionn *string) *DcimPowerOutletsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionNic(descriptionNic *string) *DcimPowerOutletsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionNie(descriptionNie *string) *DcimPowerOutletsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionNiew(descriptionNiew *string) *DcimPowerOutletsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDescriptionNisw(descriptionNisw *string) *DcimPowerOutletsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithDevice adds the device to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDevice(device *string) *DcimPowerOutletsListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDevicen(devicen *string) *DcimPowerOutletsListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDeviceID(deviceID *string) *DcimPowerOutletsListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithDeviceIDn(deviceIDn *string) *DcimPowerOutletsListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithFeedLeg adds the feedLeg to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithFeedLeg(feedLeg *string) *DcimPowerOutletsListParams {
	o.SetFeedLeg(feedLeg)
	return o
}

// SetFeedLeg adds the feedLeg to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetFeedLeg(feedLeg *string) {
	o.FeedLeg = feedLeg
}

// WithFeedLegn adds the feedLegn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithFeedLegn(feedLegn *string) *DcimPowerOutletsListParams {
	o.SetFeedLegn(feedLegn)
	return o
}

// SetFeedLegn adds the feedLegN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetFeedLegn(feedLegn *string) {
	o.FeedLegn = feedLegn
}

// WithID adds the id to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithID(id *string) *DcimPowerOutletsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDIc(iDIc *string) *DcimPowerOutletsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDIe(iDIe *string) *DcimPowerOutletsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDIew(iDIew *string) *DcimPowerOutletsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDIsw(iDIsw *string) *DcimPowerOutletsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDn(iDn *string) *DcimPowerOutletsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDNic(iDNic *string) *DcimPowerOutletsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDNie(iDNie *string) *DcimPowerOutletsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDNiew(iDNiew *string) *DcimPowerOutletsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithIDNisw(iDNisw *string) *DcimPowerOutletsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithLimit(limit *int64) *DcimPowerOutletsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithName(name *string) *DcimPowerOutletsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameIc(nameIc *string) *DcimPowerOutletsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameIe(nameIe *string) *DcimPowerOutletsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameIew(nameIew *string) *DcimPowerOutletsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameIsw(nameIsw *string) *DcimPowerOutletsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNamen(namen *string) *DcimPowerOutletsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameNic(nameNic *string) *DcimPowerOutletsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameNie(nameNie *string) *DcimPowerOutletsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameNiew(nameNiew *string) *DcimPowerOutletsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithNameNisw(nameNisw *string) *DcimPowerOutletsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithOffset(offset *int64) *DcimPowerOutletsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithQ(q *string) *DcimPowerOutletsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithRegion(region *string) *DcimPowerOutletsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithRegionn(regionn *string) *DcimPowerOutletsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithRegionID(regionID *string) *DcimPowerOutletsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithRegionIDn(regionIDn *string) *DcimPowerOutletsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithSite(site *string) *DcimPowerOutletsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithSiten(siten *string) *DcimPowerOutletsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithSiteID(siteID *string) *DcimPowerOutletsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithSiteIDn(siteIDn *string) *DcimPowerOutletsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithTag(tag *string) *DcimPowerOutletsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithTagn(tagn *string) *DcimPowerOutletsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithType adds the typeVar to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithType(typeVar *string) *DcimPowerOutletsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) WithTypen(typen *string) *DcimPowerOutletsListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim power outlets list params
func (o *DcimPowerOutletsListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FeedLeg != nil {

		// query param feed_leg
		var qrFeedLeg string

		if o.FeedLeg != nil {
			qrFeedLeg = *o.FeedLeg
		}
		qFeedLeg := qrFeedLeg
		if qFeedLeg != "" {

			if err := r.SetQueryParam("feed_leg", qFeedLeg); err != nil {
				return err
			}
		}
	}

	if o.FeedLegn != nil {

		// query param feed_leg__n
		var qrFeedLegn string

		if o.FeedLegn != nil {
			qrFeedLegn = *o.FeedLegn
		}
		qFeedLegn := qrFeedLegn
		if qFeedLegn != "" {

			if err := r.SetQueryParam("feed_leg__n", qFeedLegn); err != nil {
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
