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

// NewDcimCablesListParams creates a new DcimCablesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCablesListParams() *DcimCablesListParams {
	return &DcimCablesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCablesListParamsWithTimeout creates a new DcimCablesListParams object
// with the ability to set a timeout on a request.
func NewDcimCablesListParamsWithTimeout(timeout time.Duration) *DcimCablesListParams {
	return &DcimCablesListParams{
		timeout: timeout,
	}
}

// NewDcimCablesListParamsWithContext creates a new DcimCablesListParams object
// with the ability to set a context for a request.
func NewDcimCablesListParamsWithContext(ctx context.Context) *DcimCablesListParams {
	return &DcimCablesListParams{
		Context: ctx,
	}
}

// NewDcimCablesListParamsWithHTTPClient creates a new DcimCablesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCablesListParamsWithHTTPClient(client *http.Client) *DcimCablesListParams {
	return &DcimCablesListParams{
		HTTPClient: client,
	}
}

/* DcimCablesListParams contains all the parameters to send to the API endpoint
   for the dcim cables list operation.

   Typically these are written to a http.Request.
*/
type DcimCablesListParams struct {

	// Color.
	Color *string

	// Colorn.
	Colorn *string

	// Device.
	Device *string

	// DeviceID.
	DeviceID *string

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

	// Label.
	Label *string

	// LabelIc.
	LabelIc *string

	// LabelIe.
	LabelIe *string

	// LabelIew.
	LabelIew *string

	// LabelIsw.
	LabelIsw *string

	// Labeln.
	Labeln *string

	// LabelNic.
	LabelNic *string

	// LabelNie.
	LabelNie *string

	// LabelNiew.
	LabelNiew *string

	// LabelNisw.
	LabelNisw *string

	// Length.
	Length *string

	// LengthGt.
	LengthGt *string

	// LengthGte.
	LengthGte *string

	// LengthLt.
	LengthLt *string

	// LengthLte.
	LengthLte *string

	// Lengthn.
	Lengthn *string

	// LengthUnit.
	LengthUnit *string

	// LengthUnitn.
	LengthUnitn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Rack.
	Rack *string

	// RackID.
	RackID *string

	// Site.
	Site *string

	// SiteID.
	SiteID *string

	// Status.
	Status *string

	// Statusn.
	Statusn *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Tenant.
	Tenant *string

	// TenantID.
	TenantID *string

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cables list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesListParams) WithDefaults() *DcimCablesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cables list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCablesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cables list params
func (o *DcimCablesListParams) WithTimeout(timeout time.Duration) *DcimCablesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cables list params
func (o *DcimCablesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cables list params
func (o *DcimCablesListParams) WithContext(ctx context.Context) *DcimCablesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cables list params
func (o *DcimCablesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cables list params
func (o *DcimCablesListParams) WithHTTPClient(client *http.Client) *DcimCablesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cables list params
func (o *DcimCablesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColor adds the color to the dcim cables list params
func (o *DcimCablesListParams) WithColor(color *string) *DcimCablesListParams {
	o.SetColor(color)
	return o
}

// SetColor adds the color to the dcim cables list params
func (o *DcimCablesListParams) SetColor(color *string) {
	o.Color = color
}

// WithColorn adds the colorn to the dcim cables list params
func (o *DcimCablesListParams) WithColorn(colorn *string) *DcimCablesListParams {
	o.SetColorn(colorn)
	return o
}

// SetColorn adds the colorN to the dcim cables list params
func (o *DcimCablesListParams) SetColorn(colorn *string) {
	o.Colorn = colorn
}

// WithDevice adds the device to the dcim cables list params
func (o *DcimCablesListParams) WithDevice(device *string) *DcimCablesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the dcim cables list params
func (o *DcimCablesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the dcim cables list params
func (o *DcimCablesListParams) WithDeviceID(deviceID *string) *DcimCablesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the dcim cables list params
func (o *DcimCablesListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithID adds the id to the dcim cables list params
func (o *DcimCablesListParams) WithID(id *string) *DcimCablesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cables list params
func (o *DcimCablesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim cables list params
func (o *DcimCablesListParams) WithIDIc(iDIc *string) *DcimCablesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim cables list params
func (o *DcimCablesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim cables list params
func (o *DcimCablesListParams) WithIDIe(iDIe *string) *DcimCablesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim cables list params
func (o *DcimCablesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim cables list params
func (o *DcimCablesListParams) WithIDIew(iDIew *string) *DcimCablesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim cables list params
func (o *DcimCablesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim cables list params
func (o *DcimCablesListParams) WithIDIsw(iDIsw *string) *DcimCablesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim cables list params
func (o *DcimCablesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim cables list params
func (o *DcimCablesListParams) WithIDn(iDn *string) *DcimCablesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim cables list params
func (o *DcimCablesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim cables list params
func (o *DcimCablesListParams) WithIDNic(iDNic *string) *DcimCablesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim cables list params
func (o *DcimCablesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim cables list params
func (o *DcimCablesListParams) WithIDNie(iDNie *string) *DcimCablesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim cables list params
func (o *DcimCablesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim cables list params
func (o *DcimCablesListParams) WithIDNiew(iDNiew *string) *DcimCablesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim cables list params
func (o *DcimCablesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim cables list params
func (o *DcimCablesListParams) WithIDNisw(iDNisw *string) *DcimCablesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim cables list params
func (o *DcimCablesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLabel adds the label to the dcim cables list params
func (o *DcimCablesListParams) WithLabel(label *string) *DcimCablesListParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the dcim cables list params
func (o *DcimCablesListParams) SetLabel(label *string) {
	o.Label = label
}

// WithLabelIc adds the labelIc to the dcim cables list params
func (o *DcimCablesListParams) WithLabelIc(labelIc *string) *DcimCablesListParams {
	o.SetLabelIc(labelIc)
	return o
}

// SetLabelIc adds the labelIc to the dcim cables list params
func (o *DcimCablesListParams) SetLabelIc(labelIc *string) {
	o.LabelIc = labelIc
}

// WithLabelIe adds the labelIe to the dcim cables list params
func (o *DcimCablesListParams) WithLabelIe(labelIe *string) *DcimCablesListParams {
	o.SetLabelIe(labelIe)
	return o
}

// SetLabelIe adds the labelIe to the dcim cables list params
func (o *DcimCablesListParams) SetLabelIe(labelIe *string) {
	o.LabelIe = labelIe
}

// WithLabelIew adds the labelIew to the dcim cables list params
func (o *DcimCablesListParams) WithLabelIew(labelIew *string) *DcimCablesListParams {
	o.SetLabelIew(labelIew)
	return o
}

// SetLabelIew adds the labelIew to the dcim cables list params
func (o *DcimCablesListParams) SetLabelIew(labelIew *string) {
	o.LabelIew = labelIew
}

// WithLabelIsw adds the labelIsw to the dcim cables list params
func (o *DcimCablesListParams) WithLabelIsw(labelIsw *string) *DcimCablesListParams {
	o.SetLabelIsw(labelIsw)
	return o
}

// SetLabelIsw adds the labelIsw to the dcim cables list params
func (o *DcimCablesListParams) SetLabelIsw(labelIsw *string) {
	o.LabelIsw = labelIsw
}

// WithLabeln adds the labeln to the dcim cables list params
func (o *DcimCablesListParams) WithLabeln(labeln *string) *DcimCablesListParams {
	o.SetLabeln(labeln)
	return o
}

// SetLabeln adds the labelN to the dcim cables list params
func (o *DcimCablesListParams) SetLabeln(labeln *string) {
	o.Labeln = labeln
}

// WithLabelNic adds the labelNic to the dcim cables list params
func (o *DcimCablesListParams) WithLabelNic(labelNic *string) *DcimCablesListParams {
	o.SetLabelNic(labelNic)
	return o
}

// SetLabelNic adds the labelNic to the dcim cables list params
func (o *DcimCablesListParams) SetLabelNic(labelNic *string) {
	o.LabelNic = labelNic
}

// WithLabelNie adds the labelNie to the dcim cables list params
func (o *DcimCablesListParams) WithLabelNie(labelNie *string) *DcimCablesListParams {
	o.SetLabelNie(labelNie)
	return o
}

// SetLabelNie adds the labelNie to the dcim cables list params
func (o *DcimCablesListParams) SetLabelNie(labelNie *string) {
	o.LabelNie = labelNie
}

// WithLabelNiew adds the labelNiew to the dcim cables list params
func (o *DcimCablesListParams) WithLabelNiew(labelNiew *string) *DcimCablesListParams {
	o.SetLabelNiew(labelNiew)
	return o
}

// SetLabelNiew adds the labelNiew to the dcim cables list params
func (o *DcimCablesListParams) SetLabelNiew(labelNiew *string) {
	o.LabelNiew = labelNiew
}

// WithLabelNisw adds the labelNisw to the dcim cables list params
func (o *DcimCablesListParams) WithLabelNisw(labelNisw *string) *DcimCablesListParams {
	o.SetLabelNisw(labelNisw)
	return o
}

// SetLabelNisw adds the labelNisw to the dcim cables list params
func (o *DcimCablesListParams) SetLabelNisw(labelNisw *string) {
	o.LabelNisw = labelNisw
}

// WithLength adds the length to the dcim cables list params
func (o *DcimCablesListParams) WithLength(length *string) *DcimCablesListParams {
	o.SetLength(length)
	return o
}

// SetLength adds the length to the dcim cables list params
func (o *DcimCablesListParams) SetLength(length *string) {
	o.Length = length
}

// WithLengthGt adds the lengthGt to the dcim cables list params
func (o *DcimCablesListParams) WithLengthGt(lengthGt *string) *DcimCablesListParams {
	o.SetLengthGt(lengthGt)
	return o
}

// SetLengthGt adds the lengthGt to the dcim cables list params
func (o *DcimCablesListParams) SetLengthGt(lengthGt *string) {
	o.LengthGt = lengthGt
}

// WithLengthGte adds the lengthGte to the dcim cables list params
func (o *DcimCablesListParams) WithLengthGte(lengthGte *string) *DcimCablesListParams {
	o.SetLengthGte(lengthGte)
	return o
}

// SetLengthGte adds the lengthGte to the dcim cables list params
func (o *DcimCablesListParams) SetLengthGte(lengthGte *string) {
	o.LengthGte = lengthGte
}

// WithLengthLt adds the lengthLt to the dcim cables list params
func (o *DcimCablesListParams) WithLengthLt(lengthLt *string) *DcimCablesListParams {
	o.SetLengthLt(lengthLt)
	return o
}

// SetLengthLt adds the lengthLt to the dcim cables list params
func (o *DcimCablesListParams) SetLengthLt(lengthLt *string) {
	o.LengthLt = lengthLt
}

// WithLengthLte adds the lengthLte to the dcim cables list params
func (o *DcimCablesListParams) WithLengthLte(lengthLte *string) *DcimCablesListParams {
	o.SetLengthLte(lengthLte)
	return o
}

// SetLengthLte adds the lengthLte to the dcim cables list params
func (o *DcimCablesListParams) SetLengthLte(lengthLte *string) {
	o.LengthLte = lengthLte
}

// WithLengthn adds the lengthn to the dcim cables list params
func (o *DcimCablesListParams) WithLengthn(lengthn *string) *DcimCablesListParams {
	o.SetLengthn(lengthn)
	return o
}

// SetLengthn adds the lengthN to the dcim cables list params
func (o *DcimCablesListParams) SetLengthn(lengthn *string) {
	o.Lengthn = lengthn
}

// WithLengthUnit adds the lengthUnit to the dcim cables list params
func (o *DcimCablesListParams) WithLengthUnit(lengthUnit *string) *DcimCablesListParams {
	o.SetLengthUnit(lengthUnit)
	return o
}

// SetLengthUnit adds the lengthUnit to the dcim cables list params
func (o *DcimCablesListParams) SetLengthUnit(lengthUnit *string) {
	o.LengthUnit = lengthUnit
}

// WithLengthUnitn adds the lengthUnitn to the dcim cables list params
func (o *DcimCablesListParams) WithLengthUnitn(lengthUnitn *string) *DcimCablesListParams {
	o.SetLengthUnitn(lengthUnitn)
	return o
}

// SetLengthUnitn adds the lengthUnitN to the dcim cables list params
func (o *DcimCablesListParams) SetLengthUnitn(lengthUnitn *string) {
	o.LengthUnitn = lengthUnitn
}

// WithLimit adds the limit to the dcim cables list params
func (o *DcimCablesListParams) WithLimit(limit *int64) *DcimCablesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim cables list params
func (o *DcimCablesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the dcim cables list params
func (o *DcimCablesListParams) WithOffset(offset *int64) *DcimCablesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim cables list params
func (o *DcimCablesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim cables list params
func (o *DcimCablesListParams) WithQ(q *string) *DcimCablesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim cables list params
func (o *DcimCablesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRack adds the rack to the dcim cables list params
func (o *DcimCablesListParams) WithRack(rack *string) *DcimCablesListParams {
	o.SetRack(rack)
	return o
}

// SetRack adds the rack to the dcim cables list params
func (o *DcimCablesListParams) SetRack(rack *string) {
	o.Rack = rack
}

// WithRackID adds the rackID to the dcim cables list params
func (o *DcimCablesListParams) WithRackID(rackID *string) *DcimCablesListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim cables list params
func (o *DcimCablesListParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithSite adds the site to the dcim cables list params
func (o *DcimCablesListParams) WithSite(site *string) *DcimCablesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim cables list params
func (o *DcimCablesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiteID adds the siteID to the dcim cables list params
func (o *DcimCablesListParams) WithSiteID(siteID *string) *DcimCablesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim cables list params
func (o *DcimCablesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithStatus adds the status to the dcim cables list params
func (o *DcimCablesListParams) WithStatus(status *string) *DcimCablesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim cables list params
func (o *DcimCablesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the dcim cables list params
func (o *DcimCablesListParams) WithStatusn(statusn *string) *DcimCablesListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the dcim cables list params
func (o *DcimCablesListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the dcim cables list params
func (o *DcimCablesListParams) WithTag(tag *string) *DcimCablesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim cables list params
func (o *DcimCablesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim cables list params
func (o *DcimCablesListParams) WithTagn(tagn *string) *DcimCablesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim cables list params
func (o *DcimCablesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the dcim cables list params
func (o *DcimCablesListParams) WithTenant(tenant *string) *DcimCablesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim cables list params
func (o *DcimCablesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantID adds the tenantID to the dcim cables list params
func (o *DcimCablesListParams) WithTenantID(tenantID *string) *DcimCablesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim cables list params
func (o *DcimCablesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithType adds the typeVar to the dcim cables list params
func (o *DcimCablesListParams) WithType(typeVar *string) *DcimCablesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim cables list params
func (o *DcimCablesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim cables list params
func (o *DcimCablesListParams) WithTypen(typen *string) *DcimCablesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim cables list params
func (o *DcimCablesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCablesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Color != nil {

		// query param color
		var qrColor string

		if o.Color != nil {
			qrColor = *o.Color
		}
		qColor := qrColor
		if qColor != "" {

			if err := r.SetQueryParam("color", qColor); err != nil {
				return err
			}
		}
	}

	if o.Colorn != nil {

		// query param color__n
		var qrColorn string

		if o.Colorn != nil {
			qrColorn = *o.Colorn
		}
		qColorn := qrColorn
		if qColorn != "" {

			if err := r.SetQueryParam("color__n", qColorn); err != nil {
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

	if o.Label != nil {

		// query param label
		var qrLabel string

		if o.Label != nil {
			qrLabel = *o.Label
		}
		qLabel := qrLabel
		if qLabel != "" {

			if err := r.SetQueryParam("label", qLabel); err != nil {
				return err
			}
		}
	}

	if o.LabelIc != nil {

		// query param label__ic
		var qrLabelIc string

		if o.LabelIc != nil {
			qrLabelIc = *o.LabelIc
		}
		qLabelIc := qrLabelIc
		if qLabelIc != "" {

			if err := r.SetQueryParam("label__ic", qLabelIc); err != nil {
				return err
			}
		}
	}

	if o.LabelIe != nil {

		// query param label__ie
		var qrLabelIe string

		if o.LabelIe != nil {
			qrLabelIe = *o.LabelIe
		}
		qLabelIe := qrLabelIe
		if qLabelIe != "" {

			if err := r.SetQueryParam("label__ie", qLabelIe); err != nil {
				return err
			}
		}
	}

	if o.LabelIew != nil {

		// query param label__iew
		var qrLabelIew string

		if o.LabelIew != nil {
			qrLabelIew = *o.LabelIew
		}
		qLabelIew := qrLabelIew
		if qLabelIew != "" {

			if err := r.SetQueryParam("label__iew", qLabelIew); err != nil {
				return err
			}
		}
	}

	if o.LabelIsw != nil {

		// query param label__isw
		var qrLabelIsw string

		if o.LabelIsw != nil {
			qrLabelIsw = *o.LabelIsw
		}
		qLabelIsw := qrLabelIsw
		if qLabelIsw != "" {

			if err := r.SetQueryParam("label__isw", qLabelIsw); err != nil {
				return err
			}
		}
	}

	if o.Labeln != nil {

		// query param label__n
		var qrLabeln string

		if o.Labeln != nil {
			qrLabeln = *o.Labeln
		}
		qLabeln := qrLabeln
		if qLabeln != "" {

			if err := r.SetQueryParam("label__n", qLabeln); err != nil {
				return err
			}
		}
	}

	if o.LabelNic != nil {

		// query param label__nic
		var qrLabelNic string

		if o.LabelNic != nil {
			qrLabelNic = *o.LabelNic
		}
		qLabelNic := qrLabelNic
		if qLabelNic != "" {

			if err := r.SetQueryParam("label__nic", qLabelNic); err != nil {
				return err
			}
		}
	}

	if o.LabelNie != nil {

		// query param label__nie
		var qrLabelNie string

		if o.LabelNie != nil {
			qrLabelNie = *o.LabelNie
		}
		qLabelNie := qrLabelNie
		if qLabelNie != "" {

			if err := r.SetQueryParam("label__nie", qLabelNie); err != nil {
				return err
			}
		}
	}

	if o.LabelNiew != nil {

		// query param label__niew
		var qrLabelNiew string

		if o.LabelNiew != nil {
			qrLabelNiew = *o.LabelNiew
		}
		qLabelNiew := qrLabelNiew
		if qLabelNiew != "" {

			if err := r.SetQueryParam("label__niew", qLabelNiew); err != nil {
				return err
			}
		}
	}

	if o.LabelNisw != nil {

		// query param label__nisw
		var qrLabelNisw string

		if o.LabelNisw != nil {
			qrLabelNisw = *o.LabelNisw
		}
		qLabelNisw := qrLabelNisw
		if qLabelNisw != "" {

			if err := r.SetQueryParam("label__nisw", qLabelNisw); err != nil {
				return err
			}
		}
	}

	if o.Length != nil {

		// query param length
		var qrLength string

		if o.Length != nil {
			qrLength = *o.Length
		}
		qLength := qrLength
		if qLength != "" {

			if err := r.SetQueryParam("length", qLength); err != nil {
				return err
			}
		}
	}

	if o.LengthGt != nil {

		// query param length__gt
		var qrLengthGt string

		if o.LengthGt != nil {
			qrLengthGt = *o.LengthGt
		}
		qLengthGt := qrLengthGt
		if qLengthGt != "" {

			if err := r.SetQueryParam("length__gt", qLengthGt); err != nil {
				return err
			}
		}
	}

	if o.LengthGte != nil {

		// query param length__gte
		var qrLengthGte string

		if o.LengthGte != nil {
			qrLengthGte = *o.LengthGte
		}
		qLengthGte := qrLengthGte
		if qLengthGte != "" {

			if err := r.SetQueryParam("length__gte", qLengthGte); err != nil {
				return err
			}
		}
	}

	if o.LengthLt != nil {

		// query param length__lt
		var qrLengthLt string

		if o.LengthLt != nil {
			qrLengthLt = *o.LengthLt
		}
		qLengthLt := qrLengthLt
		if qLengthLt != "" {

			if err := r.SetQueryParam("length__lt", qLengthLt); err != nil {
				return err
			}
		}
	}

	if o.LengthLte != nil {

		// query param length__lte
		var qrLengthLte string

		if o.LengthLte != nil {
			qrLengthLte = *o.LengthLte
		}
		qLengthLte := qrLengthLte
		if qLengthLte != "" {

			if err := r.SetQueryParam("length__lte", qLengthLte); err != nil {
				return err
			}
		}
	}

	if o.Lengthn != nil {

		// query param length__n
		var qrLengthn string

		if o.Lengthn != nil {
			qrLengthn = *o.Lengthn
		}
		qLengthn := qrLengthn
		if qLengthn != "" {

			if err := r.SetQueryParam("length__n", qLengthn); err != nil {
				return err
			}
		}
	}

	if o.LengthUnit != nil {

		// query param length_unit
		var qrLengthUnit string

		if o.LengthUnit != nil {
			qrLengthUnit = *o.LengthUnit
		}
		qLengthUnit := qrLengthUnit
		if qLengthUnit != "" {

			if err := r.SetQueryParam("length_unit", qLengthUnit); err != nil {
				return err
			}
		}
	}

	if o.LengthUnitn != nil {

		// query param length_unit__n
		var qrLengthUnitn string

		if o.LengthUnitn != nil {
			qrLengthUnitn = *o.LengthUnitn
		}
		qLengthUnitn := qrLengthUnitn
		if qLengthUnitn != "" {

			if err := r.SetQueryParam("length_unit__n", qLengthUnitn); err != nil {
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

	if o.Rack != nil {

		// query param rack
		var qrRack string

		if o.Rack != nil {
			qrRack = *o.Rack
		}
		qRack := qrRack
		if qRack != "" {

			if err := r.SetQueryParam("rack", qRack); err != nil {
				return err
			}
		}
	}

	if o.RackID != nil {

		// query param rack_id
		var qrRackID string

		if o.RackID != nil {
			qrRackID = *o.RackID
		}
		qRackID := qrRackID
		if qRackID != "" {

			if err := r.SetQueryParam("rack_id", qRackID); err != nil {
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

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if o.Statusn != nil {

		// query param status__n
		var qrStatusn string

		if o.Statusn != nil {
			qrStatusn = *o.Statusn
		}
		qStatusn := qrStatusn
		if qStatusn != "" {

			if err := r.SetQueryParam("status__n", qStatusn); err != nil {
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

	if o.Tenant != nil {

		// query param tenant
		var qrTenant string

		if o.Tenant != nil {
			qrTenant = *o.Tenant
		}
		qTenant := qrTenant
		if qTenant != "" {

			if err := r.SetQueryParam("tenant", qTenant); err != nil {
				return err
			}
		}
	}

	if o.TenantID != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantID != nil {
			qrTenantID = *o.TenantID
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
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
