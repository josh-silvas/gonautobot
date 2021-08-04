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

// NewDcimDeviceBaysListParams creates a new DcimDeviceBaysListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceBaysListParams() *DcimDeviceBaysListParams {
	return &DcimDeviceBaysListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceBaysListParamsWithTimeout creates a new DcimDeviceBaysListParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceBaysListParamsWithTimeout(timeout time.Duration) *DcimDeviceBaysListParams {
	return &DcimDeviceBaysListParams{
		timeout: timeout,
	}
}

// NewDcimDeviceBaysListParamsWithContext creates a new DcimDeviceBaysListParams object
// with the ability to set a context for a request.
func NewDcimDeviceBaysListParamsWithContext(ctx context.Context) *DcimDeviceBaysListParams {
	return &DcimDeviceBaysListParams{
		Context: ctx,
	}
}

// NewDcimDeviceBaysListParamsWithHTTPClient creates a new DcimDeviceBaysListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceBaysListParamsWithHTTPClient(client *http.Client) *DcimDeviceBaysListParams {
	return &DcimDeviceBaysListParams{
		HTTPClient: client,
	}
}

/* DcimDeviceBaysListParams contains all the parameters to send to the API endpoint
   for the dcim device bays list operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceBaysListParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device bays list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysListParams) WithDefaults() *DcimDeviceBaysListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device bays list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceBaysListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithTimeout(timeout time.Duration) *DcimDeviceBaysListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithContext(ctx context.Context) *DcimDeviceBaysListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithHTTPClient(client *http.Client) *DcimDeviceBaysListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescription(description *string) *DcimDeviceBaysListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionIc(descriptionIc *string) *DcimDeviceBaysListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionIe(descriptionIe *string) *DcimDeviceBaysListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionIew(descriptionIew *string) *DcimDeviceBaysListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionIsw(descriptionIsw *string) *DcimDeviceBaysListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionn(descriptionn *string) *DcimDeviceBaysListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionNic(descriptionNic *string) *DcimDeviceBaysListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionNie(descriptionNie *string) *DcimDeviceBaysListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionNiew(descriptionNiew *string) *DcimDeviceBaysListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDescriptionNisw(descriptionNisw *string) *DcimDeviceBaysListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithDevice adds the device to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDevice(device *string) *DcimDeviceBaysListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDevicen(devicen *string) *DcimDeviceBaysListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDeviceID(deviceID *string) *DcimDeviceBaysListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithDeviceIDn(deviceIDn *string) *DcimDeviceBaysListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithID adds the id to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithID(id *string) *DcimDeviceBaysListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDIc(iDIc *string) *DcimDeviceBaysListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDIe(iDIe *string) *DcimDeviceBaysListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDIew(iDIew *string) *DcimDeviceBaysListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDIsw(iDIsw *string) *DcimDeviceBaysListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDn(iDn *string) *DcimDeviceBaysListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDNic(iDNic *string) *DcimDeviceBaysListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDNie(iDNie *string) *DcimDeviceBaysListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDNiew(iDNiew *string) *DcimDeviceBaysListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithIDNisw(iDNisw *string) *DcimDeviceBaysListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithLimit(limit *int64) *DcimDeviceBaysListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithName(name *string) *DcimDeviceBaysListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameIc(nameIc *string) *DcimDeviceBaysListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameIe(nameIe *string) *DcimDeviceBaysListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameIew(nameIew *string) *DcimDeviceBaysListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameIsw(nameIsw *string) *DcimDeviceBaysListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNamen(namen *string) *DcimDeviceBaysListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameNic(nameNic *string) *DcimDeviceBaysListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameNie(nameNie *string) *DcimDeviceBaysListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameNiew(nameNiew *string) *DcimDeviceBaysListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithNameNisw(nameNisw *string) *DcimDeviceBaysListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithOffset(offset *int64) *DcimDeviceBaysListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithQ(q *string) *DcimDeviceBaysListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithRegion(region *string) *DcimDeviceBaysListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithRegionn(regionn *string) *DcimDeviceBaysListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithRegionID(regionID *string) *DcimDeviceBaysListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithRegionIDn(regionIDn *string) *DcimDeviceBaysListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithSite(site *string) *DcimDeviceBaysListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithSiten(siten *string) *DcimDeviceBaysListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithSiteID(siteID *string) *DcimDeviceBaysListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithSiteIDn(siteIDn *string) *DcimDeviceBaysListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithTag(tag *string) *DcimDeviceBaysListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim device bays list params
func (o *DcimDeviceBaysListParams) WithTagn(tagn *string) *DcimDeviceBaysListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim device bays list params
func (o *DcimDeviceBaysListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceBaysListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
