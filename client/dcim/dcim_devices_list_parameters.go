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

// NewDcimDevicesListParams creates a new DcimDevicesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDevicesListParams() *DcimDevicesListParams {
	return &DcimDevicesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDevicesListParamsWithTimeout creates a new DcimDevicesListParams object
// with the ability to set a timeout on a request.
func NewDcimDevicesListParamsWithTimeout(timeout time.Duration) *DcimDevicesListParams {
	return &DcimDevicesListParams{
		timeout: timeout,
	}
}

// NewDcimDevicesListParamsWithContext creates a new DcimDevicesListParams object
// with the ability to set a context for a request.
func NewDcimDevicesListParamsWithContext(ctx context.Context) *DcimDevicesListParams {
	return &DcimDevicesListParams{
		Context: ctx,
	}
}

// NewDcimDevicesListParamsWithHTTPClient creates a new DcimDevicesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDevicesListParamsWithHTTPClient(client *http.Client) *DcimDevicesListParams {
	return &DcimDevicesListParams{
		HTTPClient: client,
	}
}

/* DcimDevicesListParams contains all the parameters to send to the API endpoint
   for the dcim devices list operation.

   Typically these are written to a http.Request.
*/
type DcimDevicesListParams struct {

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

	// ClusterID.
	ClusterID *string

	// ClusterIDn.
	ClusterIDn *string

	// ConsolePorts.
	ConsolePorts *string

	// ConsoleServerPorts.
	ConsoleServerPorts *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DeviceBays.
	DeviceBays *string

	// DeviceTypeID.
	DeviceTypeID *string

	// DeviceTypeIDn.
	DeviceTypeIDn *string

	// Face.
	Face *string

	// Facen.
	Facen *string

	// HasPrimaryIP.
	HasPrimaryIP *string

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

	// Interfaces.
	Interfaces *string

	// IsFullDepth.
	IsFullDepth *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// LocalContextData.
	LocalContextData *string

	// LocalContextSchema.
	LocalContextSchema *string

	// LocalContextScheman.
	LocalContextScheman *string

	// LocalContextSchemaID.
	LocalContextSchemaID *string

	// LocalContextSchemaIDn.
	LocalContextSchemaIDn *string

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

	// Manufacturer.
	Manufacturer *string

	// Manufacturern.
	Manufacturern *string

	// ManufacturerID.
	ManufacturerID *string

	// ManufacturerIDn.
	ManufacturerIDn *string

	// Model.
	Model *string

	// Modeln.
	Modeln *string

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

	// PassThroughPorts.
	PassThroughPorts *string

	// Platform.
	Platform *string

	// Platformn.
	Platformn *string

	// PlatformID.
	PlatformID *string

	// PlatformIDn.
	PlatformIDn *string

	// Position.
	Position *string

	// PositionGt.
	PositionGt *string

	// PositionGte.
	PositionGte *string

	// PositionLt.
	PositionLt *string

	// PositionLte.
	PositionLte *string

	// Positionn.
	Positionn *string

	// PowerOutlets.
	PowerOutlets *string

	// PowerPorts.
	PowerPorts *string

	// Q.
	Q *string

	// RackGroupID.
	RackGroupID *string

	// RackGroupIDn.
	RackGroupIDn *string

	// RackID.
	RackID *string

	// RackIDn.
	RackIDn *string

	// Region.
	Region *string

	// Regionn.
	Regionn *string

	// RegionID.
	RegionID *string

	// RegionIDn.
	RegionIDn *string

	// Role.
	Role *string

	// Rolen.
	Rolen *string

	// RoleID.
	RoleID *string

	// RoleIDn.
	RoleIDn *string

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

	// Tenantn.
	Tenantn *string

	// TenantGroup.
	TenantGroup *string

	// TenantGroupn.
	TenantGroupn *string

	// TenantGroupID.
	TenantGroupID *string

	// TenantGroupIDn.
	TenantGroupIDn *string

	// TenantID.
	TenantID *string

	// TenantIDn.
	TenantIDn *string

	// VcPosition.
	VcPosition *string

	// VcPositionGt.
	VcPositionGt *string

	// VcPositionGte.
	VcPositionGte *string

	// VcPositionLt.
	VcPositionLt *string

	// VcPositionLte.
	VcPositionLte *string

	// VcPositionn.
	VcPositionn *string

	// VcPriority.
	VcPriority *string

	// VcPriorityGt.
	VcPriorityGt *string

	// VcPriorityGte.
	VcPriorityGte *string

	// VcPriorityLt.
	VcPriorityLt *string

	// VcPriorityLte.
	VcPriorityLte *string

	// VcPriorityn.
	VcPriorityn *string

	// VirtualChassisID.
	VirtualChassisID *string

	// VirtualChassisIDn.
	VirtualChassisIDn *string

	// VirtualChassisMember.
	VirtualChassisMember *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim devices list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesListParams) WithDefaults() *DcimDevicesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim devices list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDevicesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim devices list params
func (o *DcimDevicesListParams) WithTimeout(timeout time.Duration) *DcimDevicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim devices list params
func (o *DcimDevicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim devices list params
func (o *DcimDevicesListParams) WithContext(ctx context.Context) *DcimDevicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim devices list params
func (o *DcimDevicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim devices list params
func (o *DcimDevicesListParams) WithHTTPClient(client *http.Client) *DcimDevicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim devices list params
func (o *DcimDevicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetTag adds the assetTag to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTag(assetTag *string) *DcimDevicesListParams {
	o.SetAssetTag(assetTag)
	return o
}

// SetAssetTag adds the assetTag to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTag(assetTag *string) {
	o.AssetTag = assetTag
}

// WithAssetTagIc adds the assetTagIc to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagIc(assetTagIc *string) *DcimDevicesListParams {
	o.SetAssetTagIc(assetTagIc)
	return o
}

// SetAssetTagIc adds the assetTagIc to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagIc(assetTagIc *string) {
	o.AssetTagIc = assetTagIc
}

// WithAssetTagIe adds the assetTagIe to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagIe(assetTagIe *string) *DcimDevicesListParams {
	o.SetAssetTagIe(assetTagIe)
	return o
}

// SetAssetTagIe adds the assetTagIe to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagIe(assetTagIe *string) {
	o.AssetTagIe = assetTagIe
}

// WithAssetTagIew adds the assetTagIew to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagIew(assetTagIew *string) *DcimDevicesListParams {
	o.SetAssetTagIew(assetTagIew)
	return o
}

// SetAssetTagIew adds the assetTagIew to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagIew(assetTagIew *string) {
	o.AssetTagIew = assetTagIew
}

// WithAssetTagIsw adds the assetTagIsw to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagIsw(assetTagIsw *string) *DcimDevicesListParams {
	o.SetAssetTagIsw(assetTagIsw)
	return o
}

// SetAssetTagIsw adds the assetTagIsw to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagIsw(assetTagIsw *string) {
	o.AssetTagIsw = assetTagIsw
}

// WithAssetTagn adds the assetTagn to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagn(assetTagn *string) *DcimDevicesListParams {
	o.SetAssetTagn(assetTagn)
	return o
}

// SetAssetTagn adds the assetTagN to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagn(assetTagn *string) {
	o.AssetTagn = assetTagn
}

// WithAssetTagNic adds the assetTagNic to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagNic(assetTagNic *string) *DcimDevicesListParams {
	o.SetAssetTagNic(assetTagNic)
	return o
}

// SetAssetTagNic adds the assetTagNic to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagNic(assetTagNic *string) {
	o.AssetTagNic = assetTagNic
}

// WithAssetTagNie adds the assetTagNie to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagNie(assetTagNie *string) *DcimDevicesListParams {
	o.SetAssetTagNie(assetTagNie)
	return o
}

// SetAssetTagNie adds the assetTagNie to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagNie(assetTagNie *string) {
	o.AssetTagNie = assetTagNie
}

// WithAssetTagNiew adds the assetTagNiew to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagNiew(assetTagNiew *string) *DcimDevicesListParams {
	o.SetAssetTagNiew(assetTagNiew)
	return o
}

// SetAssetTagNiew adds the assetTagNiew to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagNiew(assetTagNiew *string) {
	o.AssetTagNiew = assetTagNiew
}

// WithAssetTagNisw adds the assetTagNisw to the dcim devices list params
func (o *DcimDevicesListParams) WithAssetTagNisw(assetTagNisw *string) *DcimDevicesListParams {
	o.SetAssetTagNisw(assetTagNisw)
	return o
}

// SetAssetTagNisw adds the assetTagNisw to the dcim devices list params
func (o *DcimDevicesListParams) SetAssetTagNisw(assetTagNisw *string) {
	o.AssetTagNisw = assetTagNisw
}

// WithClusterID adds the clusterID to the dcim devices list params
func (o *DcimDevicesListParams) WithClusterID(clusterID *string) *DcimDevicesListParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the dcim devices list params
func (o *DcimDevicesListParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithClusterIDn adds the clusterIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithClusterIDn(clusterIDn *string) *DcimDevicesListParams {
	o.SetClusterIDn(clusterIDn)
	return o
}

// SetClusterIDn adds the clusterIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetClusterIDn(clusterIDn *string) {
	o.ClusterIDn = clusterIDn
}

// WithConsolePorts adds the consolePorts to the dcim devices list params
func (o *DcimDevicesListParams) WithConsolePorts(consolePorts *string) *DcimDevicesListParams {
	o.SetConsolePorts(consolePorts)
	return o
}

// SetConsolePorts adds the consolePorts to the dcim devices list params
func (o *DcimDevicesListParams) SetConsolePorts(consolePorts *string) {
	o.ConsolePorts = consolePorts
}

// WithConsoleServerPorts adds the consoleServerPorts to the dcim devices list params
func (o *DcimDevicesListParams) WithConsoleServerPorts(consoleServerPorts *string) *DcimDevicesListParams {
	o.SetConsoleServerPorts(consoleServerPorts)
	return o
}

// SetConsoleServerPorts adds the consoleServerPorts to the dcim devices list params
func (o *DcimDevicesListParams) SetConsoleServerPorts(consoleServerPorts *string) {
	o.ConsoleServerPorts = consoleServerPorts
}

// WithCreated adds the created to the dcim devices list params
func (o *DcimDevicesListParams) WithCreated(created *string) *DcimDevicesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim devices list params
func (o *DcimDevicesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim devices list params
func (o *DcimDevicesListParams) WithCreatedGte(createdGte *string) *DcimDevicesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim devices list params
func (o *DcimDevicesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim devices list params
func (o *DcimDevicesListParams) WithCreatedLte(createdLte *string) *DcimDevicesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim devices list params
func (o *DcimDevicesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDeviceBays adds the deviceBays to the dcim devices list params
func (o *DcimDevicesListParams) WithDeviceBays(deviceBays *string) *DcimDevicesListParams {
	o.SetDeviceBays(deviceBays)
	return o
}

// SetDeviceBays adds the deviceBays to the dcim devices list params
func (o *DcimDevicesListParams) SetDeviceBays(deviceBays *string) {
	o.DeviceBays = deviceBays
}

// WithDeviceTypeID adds the deviceTypeID to the dcim devices list params
func (o *DcimDevicesListParams) WithDeviceTypeID(deviceTypeID *string) *DcimDevicesListParams {
	o.SetDeviceTypeID(deviceTypeID)
	return o
}

// SetDeviceTypeID adds the deviceTypeId to the dcim devices list params
func (o *DcimDevicesListParams) SetDeviceTypeID(deviceTypeID *string) {
	o.DeviceTypeID = deviceTypeID
}

// WithDeviceTypeIDn adds the deviceTypeIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithDeviceTypeIDn(deviceTypeIDn *string) *DcimDevicesListParams {
	o.SetDeviceTypeIDn(deviceTypeIDn)
	return o
}

// SetDeviceTypeIDn adds the deviceTypeIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetDeviceTypeIDn(deviceTypeIDn *string) {
	o.DeviceTypeIDn = deviceTypeIDn
}

// WithFace adds the face to the dcim devices list params
func (o *DcimDevicesListParams) WithFace(face *string) *DcimDevicesListParams {
	o.SetFace(face)
	return o
}

// SetFace adds the face to the dcim devices list params
func (o *DcimDevicesListParams) SetFace(face *string) {
	o.Face = face
}

// WithFacen adds the facen to the dcim devices list params
func (o *DcimDevicesListParams) WithFacen(facen *string) *DcimDevicesListParams {
	o.SetFacen(facen)
	return o
}

// SetFacen adds the faceN to the dcim devices list params
func (o *DcimDevicesListParams) SetFacen(facen *string) {
	o.Facen = facen
}

// WithHasPrimaryIP adds the hasPrimaryIP to the dcim devices list params
func (o *DcimDevicesListParams) WithHasPrimaryIP(hasPrimaryIP *string) *DcimDevicesListParams {
	o.SetHasPrimaryIP(hasPrimaryIP)
	return o
}

// SetHasPrimaryIP adds the hasPrimaryIp to the dcim devices list params
func (o *DcimDevicesListParams) SetHasPrimaryIP(hasPrimaryIP *string) {
	o.HasPrimaryIP = hasPrimaryIP
}

// WithID adds the id to the dcim devices list params
func (o *DcimDevicesListParams) WithID(id *string) *DcimDevicesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim devices list params
func (o *DcimDevicesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim devices list params
func (o *DcimDevicesListParams) WithIDIc(iDIc *string) *DcimDevicesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim devices list params
func (o *DcimDevicesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim devices list params
func (o *DcimDevicesListParams) WithIDIe(iDIe *string) *DcimDevicesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim devices list params
func (o *DcimDevicesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim devices list params
func (o *DcimDevicesListParams) WithIDIew(iDIew *string) *DcimDevicesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim devices list params
func (o *DcimDevicesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim devices list params
func (o *DcimDevicesListParams) WithIDIsw(iDIsw *string) *DcimDevicesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim devices list params
func (o *DcimDevicesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim devices list params
func (o *DcimDevicesListParams) WithIDn(iDn *string) *DcimDevicesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim devices list params
func (o *DcimDevicesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim devices list params
func (o *DcimDevicesListParams) WithIDNic(iDNic *string) *DcimDevicesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim devices list params
func (o *DcimDevicesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim devices list params
func (o *DcimDevicesListParams) WithIDNie(iDNie *string) *DcimDevicesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim devices list params
func (o *DcimDevicesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim devices list params
func (o *DcimDevicesListParams) WithIDNiew(iDNiew *string) *DcimDevicesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim devices list params
func (o *DcimDevicesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim devices list params
func (o *DcimDevicesListParams) WithIDNisw(iDNisw *string) *DcimDevicesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim devices list params
func (o *DcimDevicesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithInterfaces adds the interfaces to the dcim devices list params
func (o *DcimDevicesListParams) WithInterfaces(interfaces *string) *DcimDevicesListParams {
	o.SetInterfaces(interfaces)
	return o
}

// SetInterfaces adds the interfaces to the dcim devices list params
func (o *DcimDevicesListParams) SetInterfaces(interfaces *string) {
	o.Interfaces = interfaces
}

// WithIsFullDepth adds the isFullDepth to the dcim devices list params
func (o *DcimDevicesListParams) WithIsFullDepth(isFullDepth *string) *DcimDevicesListParams {
	o.SetIsFullDepth(isFullDepth)
	return o
}

// SetIsFullDepth adds the isFullDepth to the dcim devices list params
func (o *DcimDevicesListParams) SetIsFullDepth(isFullDepth *string) {
	o.IsFullDepth = isFullDepth
}

// WithLastUpdated adds the lastUpdated to the dcim devices list params
func (o *DcimDevicesListParams) WithLastUpdated(lastUpdated *string) *DcimDevicesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim devices list params
func (o *DcimDevicesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim devices list params
func (o *DcimDevicesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimDevicesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim devices list params
func (o *DcimDevicesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim devices list params
func (o *DcimDevicesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimDevicesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim devices list params
func (o *DcimDevicesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim devices list params
func (o *DcimDevicesListParams) WithLimit(limit *int64) *DcimDevicesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim devices list params
func (o *DcimDevicesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLocalContextData adds the localContextData to the dcim devices list params
func (o *DcimDevicesListParams) WithLocalContextData(localContextData *string) *DcimDevicesListParams {
	o.SetLocalContextData(localContextData)
	return o
}

// SetLocalContextData adds the localContextData to the dcim devices list params
func (o *DcimDevicesListParams) SetLocalContextData(localContextData *string) {
	o.LocalContextData = localContextData
}

// WithLocalContextSchema adds the localContextSchema to the dcim devices list params
func (o *DcimDevicesListParams) WithLocalContextSchema(localContextSchema *string) *DcimDevicesListParams {
	o.SetLocalContextSchema(localContextSchema)
	return o
}

// SetLocalContextSchema adds the localContextSchema to the dcim devices list params
func (o *DcimDevicesListParams) SetLocalContextSchema(localContextSchema *string) {
	o.LocalContextSchema = localContextSchema
}

// WithLocalContextScheman adds the localContextScheman to the dcim devices list params
func (o *DcimDevicesListParams) WithLocalContextScheman(localContextScheman *string) *DcimDevicesListParams {
	o.SetLocalContextScheman(localContextScheman)
	return o
}

// SetLocalContextScheman adds the localContextSchemaN to the dcim devices list params
func (o *DcimDevicesListParams) SetLocalContextScheman(localContextScheman *string) {
	o.LocalContextScheman = localContextScheman
}

// WithLocalContextSchemaID adds the localContextSchemaID to the dcim devices list params
func (o *DcimDevicesListParams) WithLocalContextSchemaID(localContextSchemaID *string) *DcimDevicesListParams {
	o.SetLocalContextSchemaID(localContextSchemaID)
	return o
}

// SetLocalContextSchemaID adds the localContextSchemaId to the dcim devices list params
func (o *DcimDevicesListParams) SetLocalContextSchemaID(localContextSchemaID *string) {
	o.LocalContextSchemaID = localContextSchemaID
}

// WithLocalContextSchemaIDn adds the localContextSchemaIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithLocalContextSchemaIDn(localContextSchemaIDn *string) *DcimDevicesListParams {
	o.SetLocalContextSchemaIDn(localContextSchemaIDn)
	return o
}

// SetLocalContextSchemaIDn adds the localContextSchemaIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetLocalContextSchemaIDn(localContextSchemaIDn *string) {
	o.LocalContextSchemaIDn = localContextSchemaIDn
}

// WithMacAddress adds the macAddress to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddress(macAddress *string) *DcimDevicesListParams {
	o.SetMacAddress(macAddress)
	return o
}

// SetMacAddress adds the macAddress to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddress(macAddress *string) {
	o.MacAddress = macAddress
}

// WithMacAddressIc adds the macAddressIc to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressIc(macAddressIc *string) *DcimDevicesListParams {
	o.SetMacAddressIc(macAddressIc)
	return o
}

// SetMacAddressIc adds the macAddressIc to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressIc(macAddressIc *string) {
	o.MacAddressIc = macAddressIc
}

// WithMacAddressIe adds the macAddressIe to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressIe(macAddressIe *string) *DcimDevicesListParams {
	o.SetMacAddressIe(macAddressIe)
	return o
}

// SetMacAddressIe adds the macAddressIe to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressIe(macAddressIe *string) {
	o.MacAddressIe = macAddressIe
}

// WithMacAddressIew adds the macAddressIew to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressIew(macAddressIew *string) *DcimDevicesListParams {
	o.SetMacAddressIew(macAddressIew)
	return o
}

// SetMacAddressIew adds the macAddressIew to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressIew(macAddressIew *string) {
	o.MacAddressIew = macAddressIew
}

// WithMacAddressIsw adds the macAddressIsw to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressIsw(macAddressIsw *string) *DcimDevicesListParams {
	o.SetMacAddressIsw(macAddressIsw)
	return o
}

// SetMacAddressIsw adds the macAddressIsw to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressIsw(macAddressIsw *string) {
	o.MacAddressIsw = macAddressIsw
}

// WithMacAddressn adds the macAddressn to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressn(macAddressn *string) *DcimDevicesListParams {
	o.SetMacAddressn(macAddressn)
	return o
}

// SetMacAddressn adds the macAddressN to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressn(macAddressn *string) {
	o.MacAddressn = macAddressn
}

// WithMacAddressNic adds the macAddressNic to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressNic(macAddressNic *string) *DcimDevicesListParams {
	o.SetMacAddressNic(macAddressNic)
	return o
}

// SetMacAddressNic adds the macAddressNic to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressNic(macAddressNic *string) {
	o.MacAddressNic = macAddressNic
}

// WithMacAddressNie adds the macAddressNie to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressNie(macAddressNie *string) *DcimDevicesListParams {
	o.SetMacAddressNie(macAddressNie)
	return o
}

// SetMacAddressNie adds the macAddressNie to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressNie(macAddressNie *string) {
	o.MacAddressNie = macAddressNie
}

// WithMacAddressNiew adds the macAddressNiew to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressNiew(macAddressNiew *string) *DcimDevicesListParams {
	o.SetMacAddressNiew(macAddressNiew)
	return o
}

// SetMacAddressNiew adds the macAddressNiew to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressNiew(macAddressNiew *string) {
	o.MacAddressNiew = macAddressNiew
}

// WithMacAddressNisw adds the macAddressNisw to the dcim devices list params
func (o *DcimDevicesListParams) WithMacAddressNisw(macAddressNisw *string) *DcimDevicesListParams {
	o.SetMacAddressNisw(macAddressNisw)
	return o
}

// SetMacAddressNisw adds the macAddressNisw to the dcim devices list params
func (o *DcimDevicesListParams) SetMacAddressNisw(macAddressNisw *string) {
	o.MacAddressNisw = macAddressNisw
}

// WithManufacturer adds the manufacturer to the dcim devices list params
func (o *DcimDevicesListParams) WithManufacturer(manufacturer *string) *DcimDevicesListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim devices list params
func (o *DcimDevicesListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturern adds the manufacturern to the dcim devices list params
func (o *DcimDevicesListParams) WithManufacturern(manufacturern *string) *DcimDevicesListParams {
	o.SetManufacturern(manufacturern)
	return o
}

// SetManufacturern adds the manufacturerN to the dcim devices list params
func (o *DcimDevicesListParams) SetManufacturern(manufacturern *string) {
	o.Manufacturern = manufacturern
}

// WithManufacturerID adds the manufacturerID to the dcim devices list params
func (o *DcimDevicesListParams) WithManufacturerID(manufacturerID *string) *DcimDevicesListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim devices list params
func (o *DcimDevicesListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithManufacturerIDn adds the manufacturerIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithManufacturerIDn(manufacturerIDn *string) *DcimDevicesListParams {
	o.SetManufacturerIDn(manufacturerIDn)
	return o
}

// SetManufacturerIDn adds the manufacturerIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetManufacturerIDn(manufacturerIDn *string) {
	o.ManufacturerIDn = manufacturerIDn
}

// WithModel adds the model to the dcim devices list params
func (o *DcimDevicesListParams) WithModel(model *string) *DcimDevicesListParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the dcim devices list params
func (o *DcimDevicesListParams) SetModel(model *string) {
	o.Model = model
}

// WithModeln adds the modeln to the dcim devices list params
func (o *DcimDevicesListParams) WithModeln(modeln *string) *DcimDevicesListParams {
	o.SetModeln(modeln)
	return o
}

// SetModeln adds the modelN to the dcim devices list params
func (o *DcimDevicesListParams) SetModeln(modeln *string) {
	o.Modeln = modeln
}

// WithName adds the name to the dcim devices list params
func (o *DcimDevicesListParams) WithName(name *string) *DcimDevicesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim devices list params
func (o *DcimDevicesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim devices list params
func (o *DcimDevicesListParams) WithNameIc(nameIc *string) *DcimDevicesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim devices list params
func (o *DcimDevicesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim devices list params
func (o *DcimDevicesListParams) WithNameIe(nameIe *string) *DcimDevicesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim devices list params
func (o *DcimDevicesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim devices list params
func (o *DcimDevicesListParams) WithNameIew(nameIew *string) *DcimDevicesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim devices list params
func (o *DcimDevicesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim devices list params
func (o *DcimDevicesListParams) WithNameIsw(nameIsw *string) *DcimDevicesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim devices list params
func (o *DcimDevicesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim devices list params
func (o *DcimDevicesListParams) WithNamen(namen *string) *DcimDevicesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim devices list params
func (o *DcimDevicesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim devices list params
func (o *DcimDevicesListParams) WithNameNic(nameNic *string) *DcimDevicesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim devices list params
func (o *DcimDevicesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim devices list params
func (o *DcimDevicesListParams) WithNameNie(nameNie *string) *DcimDevicesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim devices list params
func (o *DcimDevicesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim devices list params
func (o *DcimDevicesListParams) WithNameNiew(nameNiew *string) *DcimDevicesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim devices list params
func (o *DcimDevicesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim devices list params
func (o *DcimDevicesListParams) WithNameNisw(nameNisw *string) *DcimDevicesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim devices list params
func (o *DcimDevicesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim devices list params
func (o *DcimDevicesListParams) WithOffset(offset *int64) *DcimDevicesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim devices list params
func (o *DcimDevicesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPassThroughPorts adds the passThroughPorts to the dcim devices list params
func (o *DcimDevicesListParams) WithPassThroughPorts(passThroughPorts *string) *DcimDevicesListParams {
	o.SetPassThroughPorts(passThroughPorts)
	return o
}

// SetPassThroughPorts adds the passThroughPorts to the dcim devices list params
func (o *DcimDevicesListParams) SetPassThroughPorts(passThroughPorts *string) {
	o.PassThroughPorts = passThroughPorts
}

// WithPlatform adds the platform to the dcim devices list params
func (o *DcimDevicesListParams) WithPlatform(platform *string) *DcimDevicesListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the dcim devices list params
func (o *DcimDevicesListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformn adds the platformn to the dcim devices list params
func (o *DcimDevicesListParams) WithPlatformn(platformn *string) *DcimDevicesListParams {
	o.SetPlatformn(platformn)
	return o
}

// SetPlatformn adds the platformN to the dcim devices list params
func (o *DcimDevicesListParams) SetPlatformn(platformn *string) {
	o.Platformn = platformn
}

// WithPlatformID adds the platformID to the dcim devices list params
func (o *DcimDevicesListParams) WithPlatformID(platformID *string) *DcimDevicesListParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the dcim devices list params
func (o *DcimDevicesListParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WithPlatformIDn adds the platformIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithPlatformIDn(platformIDn *string) *DcimDevicesListParams {
	o.SetPlatformIDn(platformIDn)
	return o
}

// SetPlatformIDn adds the platformIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetPlatformIDn(platformIDn *string) {
	o.PlatformIDn = platformIDn
}

// WithPosition adds the position to the dcim devices list params
func (o *DcimDevicesListParams) WithPosition(position *string) *DcimDevicesListParams {
	o.SetPosition(position)
	return o
}

// SetPosition adds the position to the dcim devices list params
func (o *DcimDevicesListParams) SetPosition(position *string) {
	o.Position = position
}

// WithPositionGt adds the positionGt to the dcim devices list params
func (o *DcimDevicesListParams) WithPositionGt(positionGt *string) *DcimDevicesListParams {
	o.SetPositionGt(positionGt)
	return o
}

// SetPositionGt adds the positionGt to the dcim devices list params
func (o *DcimDevicesListParams) SetPositionGt(positionGt *string) {
	o.PositionGt = positionGt
}

// WithPositionGte adds the positionGte to the dcim devices list params
func (o *DcimDevicesListParams) WithPositionGte(positionGte *string) *DcimDevicesListParams {
	o.SetPositionGte(positionGte)
	return o
}

// SetPositionGte adds the positionGte to the dcim devices list params
func (o *DcimDevicesListParams) SetPositionGte(positionGte *string) {
	o.PositionGte = positionGte
}

// WithPositionLt adds the positionLt to the dcim devices list params
func (o *DcimDevicesListParams) WithPositionLt(positionLt *string) *DcimDevicesListParams {
	o.SetPositionLt(positionLt)
	return o
}

// SetPositionLt adds the positionLt to the dcim devices list params
func (o *DcimDevicesListParams) SetPositionLt(positionLt *string) {
	o.PositionLt = positionLt
}

// WithPositionLte adds the positionLte to the dcim devices list params
func (o *DcimDevicesListParams) WithPositionLte(positionLte *string) *DcimDevicesListParams {
	o.SetPositionLte(positionLte)
	return o
}

// SetPositionLte adds the positionLte to the dcim devices list params
func (o *DcimDevicesListParams) SetPositionLte(positionLte *string) {
	o.PositionLte = positionLte
}

// WithPositionn adds the positionn to the dcim devices list params
func (o *DcimDevicesListParams) WithPositionn(positionn *string) *DcimDevicesListParams {
	o.SetPositionn(positionn)
	return o
}

// SetPositionn adds the positionN to the dcim devices list params
func (o *DcimDevicesListParams) SetPositionn(positionn *string) {
	o.Positionn = positionn
}

// WithPowerOutlets adds the powerOutlets to the dcim devices list params
func (o *DcimDevicesListParams) WithPowerOutlets(powerOutlets *string) *DcimDevicesListParams {
	o.SetPowerOutlets(powerOutlets)
	return o
}

// SetPowerOutlets adds the powerOutlets to the dcim devices list params
func (o *DcimDevicesListParams) SetPowerOutlets(powerOutlets *string) {
	o.PowerOutlets = powerOutlets
}

// WithPowerPorts adds the powerPorts to the dcim devices list params
func (o *DcimDevicesListParams) WithPowerPorts(powerPorts *string) *DcimDevicesListParams {
	o.SetPowerPorts(powerPorts)
	return o
}

// SetPowerPorts adds the powerPorts to the dcim devices list params
func (o *DcimDevicesListParams) SetPowerPorts(powerPorts *string) {
	o.PowerPorts = powerPorts
}

// WithQ adds the q to the dcim devices list params
func (o *DcimDevicesListParams) WithQ(q *string) *DcimDevicesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim devices list params
func (o *DcimDevicesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRackGroupID adds the rackGroupID to the dcim devices list params
func (o *DcimDevicesListParams) WithRackGroupID(rackGroupID *string) *DcimDevicesListParams {
	o.SetRackGroupID(rackGroupID)
	return o
}

// SetRackGroupID adds the rackGroupId to the dcim devices list params
func (o *DcimDevicesListParams) SetRackGroupID(rackGroupID *string) {
	o.RackGroupID = rackGroupID
}

// WithRackGroupIDn adds the rackGroupIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithRackGroupIDn(rackGroupIDn *string) *DcimDevicesListParams {
	o.SetRackGroupIDn(rackGroupIDn)
	return o
}

// SetRackGroupIDn adds the rackGroupIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetRackGroupIDn(rackGroupIDn *string) {
	o.RackGroupIDn = rackGroupIDn
}

// WithRackID adds the rackID to the dcim devices list params
func (o *DcimDevicesListParams) WithRackID(rackID *string) *DcimDevicesListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim devices list params
func (o *DcimDevicesListParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithRackIDn adds the rackIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithRackIDn(rackIDn *string) *DcimDevicesListParams {
	o.SetRackIDn(rackIDn)
	return o
}

// SetRackIDn adds the rackIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetRackIDn(rackIDn *string) {
	o.RackIDn = rackIDn
}

// WithRegion adds the region to the dcim devices list params
func (o *DcimDevicesListParams) WithRegion(region *string) *DcimDevicesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim devices list params
func (o *DcimDevicesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim devices list params
func (o *DcimDevicesListParams) WithRegionn(regionn *string) *DcimDevicesListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim devices list params
func (o *DcimDevicesListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim devices list params
func (o *DcimDevicesListParams) WithRegionID(regionID *string) *DcimDevicesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim devices list params
func (o *DcimDevicesListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithRegionIDn(regionIDn *string) *DcimDevicesListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithRole adds the role to the dcim devices list params
func (o *DcimDevicesListParams) WithRole(role *string) *DcimDevicesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the dcim devices list params
func (o *DcimDevicesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the dcim devices list params
func (o *DcimDevicesListParams) WithRolen(rolen *string) *DcimDevicesListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the dcim devices list params
func (o *DcimDevicesListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the dcim devices list params
func (o *DcimDevicesListParams) WithRoleID(roleID *string) *DcimDevicesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the dcim devices list params
func (o *DcimDevicesListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithRoleIDn(roleIDn *string) *DcimDevicesListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithSerial adds the serial to the dcim devices list params
func (o *DcimDevicesListParams) WithSerial(serial *string) *DcimDevicesListParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the dcim devices list params
func (o *DcimDevicesListParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithSite adds the site to the dcim devices list params
func (o *DcimDevicesListParams) WithSite(site *string) *DcimDevicesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim devices list params
func (o *DcimDevicesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim devices list params
func (o *DcimDevicesListParams) WithSiten(siten *string) *DcimDevicesListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim devices list params
func (o *DcimDevicesListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim devices list params
func (o *DcimDevicesListParams) WithSiteID(siteID *string) *DcimDevicesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim devices list params
func (o *DcimDevicesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithSiteIDn(siteIDn *string) *DcimDevicesListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithStatus adds the status to the dcim devices list params
func (o *DcimDevicesListParams) WithStatus(status *string) *DcimDevicesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim devices list params
func (o *DcimDevicesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the dcim devices list params
func (o *DcimDevicesListParams) WithStatusn(statusn *string) *DcimDevicesListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the dcim devices list params
func (o *DcimDevicesListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the dcim devices list params
func (o *DcimDevicesListParams) WithTag(tag *string) *DcimDevicesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim devices list params
func (o *DcimDevicesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim devices list params
func (o *DcimDevicesListParams) WithTagn(tagn *string) *DcimDevicesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim devices list params
func (o *DcimDevicesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the dcim devices list params
func (o *DcimDevicesListParams) WithTenant(tenant *string) *DcimDevicesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim devices list params
func (o *DcimDevicesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantn(tenantn *string) *DcimDevicesListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantGroup(tenantGroup *string) *DcimDevicesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantGroupn(tenantGroupn *string) *DcimDevicesListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantGroupID(tenantGroupID *string) *DcimDevicesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantGroupIDn(tenantGroupIDn *string) *DcimDevicesListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantID(tenantID *string) *DcimDevicesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithTenantIDn(tenantIDn *string) *DcimDevicesListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithVcPosition adds the vcPosition to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPosition(vcPosition *string) *DcimDevicesListParams {
	o.SetVcPosition(vcPosition)
	return o
}

// SetVcPosition adds the vcPosition to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPosition(vcPosition *string) {
	o.VcPosition = vcPosition
}

// WithVcPositionGt adds the vcPositionGt to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPositionGt(vcPositionGt *string) *DcimDevicesListParams {
	o.SetVcPositionGt(vcPositionGt)
	return o
}

// SetVcPositionGt adds the vcPositionGt to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPositionGt(vcPositionGt *string) {
	o.VcPositionGt = vcPositionGt
}

// WithVcPositionGte adds the vcPositionGte to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPositionGte(vcPositionGte *string) *DcimDevicesListParams {
	o.SetVcPositionGte(vcPositionGte)
	return o
}

// SetVcPositionGte adds the vcPositionGte to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPositionGte(vcPositionGte *string) {
	o.VcPositionGte = vcPositionGte
}

// WithVcPositionLt adds the vcPositionLt to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPositionLt(vcPositionLt *string) *DcimDevicesListParams {
	o.SetVcPositionLt(vcPositionLt)
	return o
}

// SetVcPositionLt adds the vcPositionLt to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPositionLt(vcPositionLt *string) {
	o.VcPositionLt = vcPositionLt
}

// WithVcPositionLte adds the vcPositionLte to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPositionLte(vcPositionLte *string) *DcimDevicesListParams {
	o.SetVcPositionLte(vcPositionLte)
	return o
}

// SetVcPositionLte adds the vcPositionLte to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPositionLte(vcPositionLte *string) {
	o.VcPositionLte = vcPositionLte
}

// WithVcPositionn adds the vcPositionn to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPositionn(vcPositionn *string) *DcimDevicesListParams {
	o.SetVcPositionn(vcPositionn)
	return o
}

// SetVcPositionn adds the vcPositionN to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPositionn(vcPositionn *string) {
	o.VcPositionn = vcPositionn
}

// WithVcPriority adds the vcPriority to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPriority(vcPriority *string) *DcimDevicesListParams {
	o.SetVcPriority(vcPriority)
	return o
}

// SetVcPriority adds the vcPriority to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPriority(vcPriority *string) {
	o.VcPriority = vcPriority
}

// WithVcPriorityGt adds the vcPriorityGt to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPriorityGt(vcPriorityGt *string) *DcimDevicesListParams {
	o.SetVcPriorityGt(vcPriorityGt)
	return o
}

// SetVcPriorityGt adds the vcPriorityGt to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPriorityGt(vcPriorityGt *string) {
	o.VcPriorityGt = vcPriorityGt
}

// WithVcPriorityGte adds the vcPriorityGte to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPriorityGte(vcPriorityGte *string) *DcimDevicesListParams {
	o.SetVcPriorityGte(vcPriorityGte)
	return o
}

// SetVcPriorityGte adds the vcPriorityGte to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPriorityGte(vcPriorityGte *string) {
	o.VcPriorityGte = vcPriorityGte
}

// WithVcPriorityLt adds the vcPriorityLt to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPriorityLt(vcPriorityLt *string) *DcimDevicesListParams {
	o.SetVcPriorityLt(vcPriorityLt)
	return o
}

// SetVcPriorityLt adds the vcPriorityLt to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPriorityLt(vcPriorityLt *string) {
	o.VcPriorityLt = vcPriorityLt
}

// WithVcPriorityLte adds the vcPriorityLte to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPriorityLte(vcPriorityLte *string) *DcimDevicesListParams {
	o.SetVcPriorityLte(vcPriorityLte)
	return o
}

// SetVcPriorityLte adds the vcPriorityLte to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPriorityLte(vcPriorityLte *string) {
	o.VcPriorityLte = vcPriorityLte
}

// WithVcPriorityn adds the vcPriorityn to the dcim devices list params
func (o *DcimDevicesListParams) WithVcPriorityn(vcPriorityn *string) *DcimDevicesListParams {
	o.SetVcPriorityn(vcPriorityn)
	return o
}

// SetVcPriorityn adds the vcPriorityN to the dcim devices list params
func (o *DcimDevicesListParams) SetVcPriorityn(vcPriorityn *string) {
	o.VcPriorityn = vcPriorityn
}

// WithVirtualChassisID adds the virtualChassisID to the dcim devices list params
func (o *DcimDevicesListParams) WithVirtualChassisID(virtualChassisID *string) *DcimDevicesListParams {
	o.SetVirtualChassisID(virtualChassisID)
	return o
}

// SetVirtualChassisID adds the virtualChassisId to the dcim devices list params
func (o *DcimDevicesListParams) SetVirtualChassisID(virtualChassisID *string) {
	o.VirtualChassisID = virtualChassisID
}

// WithVirtualChassisIDn adds the virtualChassisIDn to the dcim devices list params
func (o *DcimDevicesListParams) WithVirtualChassisIDn(virtualChassisIDn *string) *DcimDevicesListParams {
	o.SetVirtualChassisIDn(virtualChassisIDn)
	return o
}

// SetVirtualChassisIDn adds the virtualChassisIdN to the dcim devices list params
func (o *DcimDevicesListParams) SetVirtualChassisIDn(virtualChassisIDn *string) {
	o.VirtualChassisIDn = virtualChassisIDn
}

// WithVirtualChassisMember adds the virtualChassisMember to the dcim devices list params
func (o *DcimDevicesListParams) WithVirtualChassisMember(virtualChassisMember *string) *DcimDevicesListParams {
	o.SetVirtualChassisMember(virtualChassisMember)
	return o
}

// SetVirtualChassisMember adds the virtualChassisMember to the dcim devices list params
func (o *DcimDevicesListParams) SetVirtualChassisMember(virtualChassisMember *string) {
	o.VirtualChassisMember = virtualChassisMember
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDevicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ClusterID != nil {

		// query param cluster_id
		var qrClusterID string

		if o.ClusterID != nil {
			qrClusterID = *o.ClusterID
		}
		qClusterID := qrClusterID
		if qClusterID != "" {

			if err := r.SetQueryParam("cluster_id", qClusterID); err != nil {
				return err
			}
		}
	}

	if o.ClusterIDn != nil {

		// query param cluster_id__n
		var qrClusterIDn string

		if o.ClusterIDn != nil {
			qrClusterIDn = *o.ClusterIDn
		}
		qClusterIDn := qrClusterIDn
		if qClusterIDn != "" {

			if err := r.SetQueryParam("cluster_id__n", qClusterIDn); err != nil {
				return err
			}
		}
	}

	if o.ConsolePorts != nil {

		// query param console_ports
		var qrConsolePorts string

		if o.ConsolePorts != nil {
			qrConsolePorts = *o.ConsolePorts
		}
		qConsolePorts := qrConsolePorts
		if qConsolePorts != "" {

			if err := r.SetQueryParam("console_ports", qConsolePorts); err != nil {
				return err
			}
		}
	}

	if o.ConsoleServerPorts != nil {

		// query param console_server_ports
		var qrConsoleServerPorts string

		if o.ConsoleServerPorts != nil {
			qrConsoleServerPorts = *o.ConsoleServerPorts
		}
		qConsoleServerPorts := qrConsoleServerPorts
		if qConsoleServerPorts != "" {

			if err := r.SetQueryParam("console_server_ports", qConsoleServerPorts); err != nil {
				return err
			}
		}
	}

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.DeviceBays != nil {

		// query param device_bays
		var qrDeviceBays string

		if o.DeviceBays != nil {
			qrDeviceBays = *o.DeviceBays
		}
		qDeviceBays := qrDeviceBays
		if qDeviceBays != "" {

			if err := r.SetQueryParam("device_bays", qDeviceBays); err != nil {
				return err
			}
		}
	}

	if o.DeviceTypeID != nil {

		// query param device_type_id
		var qrDeviceTypeID string

		if o.DeviceTypeID != nil {
			qrDeviceTypeID = *o.DeviceTypeID
		}
		qDeviceTypeID := qrDeviceTypeID
		if qDeviceTypeID != "" {

			if err := r.SetQueryParam("device_type_id", qDeviceTypeID); err != nil {
				return err
			}
		}
	}

	if o.DeviceTypeIDn != nil {

		// query param device_type_id__n
		var qrDeviceTypeIDn string

		if o.DeviceTypeIDn != nil {
			qrDeviceTypeIDn = *o.DeviceTypeIDn
		}
		qDeviceTypeIDn := qrDeviceTypeIDn
		if qDeviceTypeIDn != "" {

			if err := r.SetQueryParam("device_type_id__n", qDeviceTypeIDn); err != nil {
				return err
			}
		}
	}

	if o.Face != nil {

		// query param face
		var qrFace string

		if o.Face != nil {
			qrFace = *o.Face
		}
		qFace := qrFace
		if qFace != "" {

			if err := r.SetQueryParam("face", qFace); err != nil {
				return err
			}
		}
	}

	if o.Facen != nil {

		// query param face__n
		var qrFacen string

		if o.Facen != nil {
			qrFacen = *o.Facen
		}
		qFacen := qrFacen
		if qFacen != "" {

			if err := r.SetQueryParam("face__n", qFacen); err != nil {
				return err
			}
		}
	}

	if o.HasPrimaryIP != nil {

		// query param has_primary_ip
		var qrHasPrimaryIP string

		if o.HasPrimaryIP != nil {
			qrHasPrimaryIP = *o.HasPrimaryIP
		}
		qHasPrimaryIP := qrHasPrimaryIP
		if qHasPrimaryIP != "" {

			if err := r.SetQueryParam("has_primary_ip", qHasPrimaryIP); err != nil {
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

	if o.Interfaces != nil {

		// query param interfaces
		var qrInterfaces string

		if o.Interfaces != nil {
			qrInterfaces = *o.Interfaces
		}
		qInterfaces := qrInterfaces
		if qInterfaces != "" {

			if err := r.SetQueryParam("interfaces", qInterfaces); err != nil {
				return err
			}
		}
	}

	if o.IsFullDepth != nil {

		// query param is_full_depth
		var qrIsFullDepth string

		if o.IsFullDepth != nil {
			qrIsFullDepth = *o.IsFullDepth
		}
		qIsFullDepth := qrIsFullDepth
		if qIsFullDepth != "" {

			if err := r.SetQueryParam("is_full_depth", qIsFullDepth); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
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

	if o.LocalContextData != nil {

		// query param local_context_data
		var qrLocalContextData string

		if o.LocalContextData != nil {
			qrLocalContextData = *o.LocalContextData
		}
		qLocalContextData := qrLocalContextData
		if qLocalContextData != "" {

			if err := r.SetQueryParam("local_context_data", qLocalContextData); err != nil {
				return err
			}
		}
	}

	if o.LocalContextSchema != nil {

		// query param local_context_schema
		var qrLocalContextSchema string

		if o.LocalContextSchema != nil {
			qrLocalContextSchema = *o.LocalContextSchema
		}
		qLocalContextSchema := qrLocalContextSchema
		if qLocalContextSchema != "" {

			if err := r.SetQueryParam("local_context_schema", qLocalContextSchema); err != nil {
				return err
			}
		}
	}

	if o.LocalContextScheman != nil {

		// query param local_context_schema__n
		var qrLocalContextScheman string

		if o.LocalContextScheman != nil {
			qrLocalContextScheman = *o.LocalContextScheman
		}
		qLocalContextScheman := qrLocalContextScheman
		if qLocalContextScheman != "" {

			if err := r.SetQueryParam("local_context_schema__n", qLocalContextScheman); err != nil {
				return err
			}
		}
	}

	if o.LocalContextSchemaID != nil {

		// query param local_context_schema_id
		var qrLocalContextSchemaID string

		if o.LocalContextSchemaID != nil {
			qrLocalContextSchemaID = *o.LocalContextSchemaID
		}
		qLocalContextSchemaID := qrLocalContextSchemaID
		if qLocalContextSchemaID != "" {

			if err := r.SetQueryParam("local_context_schema_id", qLocalContextSchemaID); err != nil {
				return err
			}
		}
	}

	if o.LocalContextSchemaIDn != nil {

		// query param local_context_schema_id__n
		var qrLocalContextSchemaIDn string

		if o.LocalContextSchemaIDn != nil {
			qrLocalContextSchemaIDn = *o.LocalContextSchemaIDn
		}
		qLocalContextSchemaIDn := qrLocalContextSchemaIDn
		if qLocalContextSchemaIDn != "" {

			if err := r.SetQueryParam("local_context_schema_id__n", qLocalContextSchemaIDn); err != nil {
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

	if o.Model != nil {

		// query param model
		var qrModel string

		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {

			if err := r.SetQueryParam("model", qModel); err != nil {
				return err
			}
		}
	}

	if o.Modeln != nil {

		// query param model__n
		var qrModeln string

		if o.Modeln != nil {
			qrModeln = *o.Modeln
		}
		qModeln := qrModeln
		if qModeln != "" {

			if err := r.SetQueryParam("model__n", qModeln); err != nil {
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

	if o.PassThroughPorts != nil {

		// query param pass_through_ports
		var qrPassThroughPorts string

		if o.PassThroughPorts != nil {
			qrPassThroughPorts = *o.PassThroughPorts
		}
		qPassThroughPorts := qrPassThroughPorts
		if qPassThroughPorts != "" {

			if err := r.SetQueryParam("pass_through_ports", qPassThroughPorts); err != nil {
				return err
			}
		}
	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string

		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {

			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}
	}

	if o.Platformn != nil {

		// query param platform__n
		var qrPlatformn string

		if o.Platformn != nil {
			qrPlatformn = *o.Platformn
		}
		qPlatformn := qrPlatformn
		if qPlatformn != "" {

			if err := r.SetQueryParam("platform__n", qPlatformn); err != nil {
				return err
			}
		}
	}

	if o.PlatformID != nil {

		// query param platform_id
		var qrPlatformID string

		if o.PlatformID != nil {
			qrPlatformID = *o.PlatformID
		}
		qPlatformID := qrPlatformID
		if qPlatformID != "" {

			if err := r.SetQueryParam("platform_id", qPlatformID); err != nil {
				return err
			}
		}
	}

	if o.PlatformIDn != nil {

		// query param platform_id__n
		var qrPlatformIDn string

		if o.PlatformIDn != nil {
			qrPlatformIDn = *o.PlatformIDn
		}
		qPlatformIDn := qrPlatformIDn
		if qPlatformIDn != "" {

			if err := r.SetQueryParam("platform_id__n", qPlatformIDn); err != nil {
				return err
			}
		}
	}

	if o.Position != nil {

		// query param position
		var qrPosition string

		if o.Position != nil {
			qrPosition = *o.Position
		}
		qPosition := qrPosition
		if qPosition != "" {

			if err := r.SetQueryParam("position", qPosition); err != nil {
				return err
			}
		}
	}

	if o.PositionGt != nil {

		// query param position__gt
		var qrPositionGt string

		if o.PositionGt != nil {
			qrPositionGt = *o.PositionGt
		}
		qPositionGt := qrPositionGt
		if qPositionGt != "" {

			if err := r.SetQueryParam("position__gt", qPositionGt); err != nil {
				return err
			}
		}
	}

	if o.PositionGte != nil {

		// query param position__gte
		var qrPositionGte string

		if o.PositionGte != nil {
			qrPositionGte = *o.PositionGte
		}
		qPositionGte := qrPositionGte
		if qPositionGte != "" {

			if err := r.SetQueryParam("position__gte", qPositionGte); err != nil {
				return err
			}
		}
	}

	if o.PositionLt != nil {

		// query param position__lt
		var qrPositionLt string

		if o.PositionLt != nil {
			qrPositionLt = *o.PositionLt
		}
		qPositionLt := qrPositionLt
		if qPositionLt != "" {

			if err := r.SetQueryParam("position__lt", qPositionLt); err != nil {
				return err
			}
		}
	}

	if o.PositionLte != nil {

		// query param position__lte
		var qrPositionLte string

		if o.PositionLte != nil {
			qrPositionLte = *o.PositionLte
		}
		qPositionLte := qrPositionLte
		if qPositionLte != "" {

			if err := r.SetQueryParam("position__lte", qPositionLte); err != nil {
				return err
			}
		}
	}

	if o.Positionn != nil {

		// query param position__n
		var qrPositionn string

		if o.Positionn != nil {
			qrPositionn = *o.Positionn
		}
		qPositionn := qrPositionn
		if qPositionn != "" {

			if err := r.SetQueryParam("position__n", qPositionn); err != nil {
				return err
			}
		}
	}

	if o.PowerOutlets != nil {

		// query param power_outlets
		var qrPowerOutlets string

		if o.PowerOutlets != nil {
			qrPowerOutlets = *o.PowerOutlets
		}
		qPowerOutlets := qrPowerOutlets
		if qPowerOutlets != "" {

			if err := r.SetQueryParam("power_outlets", qPowerOutlets); err != nil {
				return err
			}
		}
	}

	if o.PowerPorts != nil {

		// query param power_ports
		var qrPowerPorts string

		if o.PowerPorts != nil {
			qrPowerPorts = *o.PowerPorts
		}
		qPowerPorts := qrPowerPorts
		if qPowerPorts != "" {

			if err := r.SetQueryParam("power_ports", qPowerPorts); err != nil {
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

	if o.RackGroupID != nil {

		// query param rack_group_id
		var qrRackGroupID string

		if o.RackGroupID != nil {
			qrRackGroupID = *o.RackGroupID
		}
		qRackGroupID := qrRackGroupID
		if qRackGroupID != "" {

			if err := r.SetQueryParam("rack_group_id", qRackGroupID); err != nil {
				return err
			}
		}
	}

	if o.RackGroupIDn != nil {

		// query param rack_group_id__n
		var qrRackGroupIDn string

		if o.RackGroupIDn != nil {
			qrRackGroupIDn = *o.RackGroupIDn
		}
		qRackGroupIDn := qrRackGroupIDn
		if qRackGroupIDn != "" {

			if err := r.SetQueryParam("rack_group_id__n", qRackGroupIDn); err != nil {
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

	if o.RackIDn != nil {

		// query param rack_id__n
		var qrRackIDn string

		if o.RackIDn != nil {
			qrRackIDn = *o.RackIDn
		}
		qRackIDn := qrRackIDn
		if qRackIDn != "" {

			if err := r.SetQueryParam("rack_id__n", qRackIDn); err != nil {
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

	if o.Role != nil {

		// query param role
		var qrRole string

		if o.Role != nil {
			qrRole = *o.Role
		}
		qRole := qrRole
		if qRole != "" {

			if err := r.SetQueryParam("role", qRole); err != nil {
				return err
			}
		}
	}

	if o.Rolen != nil {

		// query param role__n
		var qrRolen string

		if o.Rolen != nil {
			qrRolen = *o.Rolen
		}
		qRolen := qrRolen
		if qRolen != "" {

			if err := r.SetQueryParam("role__n", qRolen); err != nil {
				return err
			}
		}
	}

	if o.RoleID != nil {

		// query param role_id
		var qrRoleID string

		if o.RoleID != nil {
			qrRoleID = *o.RoleID
		}
		qRoleID := qrRoleID
		if qRoleID != "" {

			if err := r.SetQueryParam("role_id", qRoleID); err != nil {
				return err
			}
		}
	}

	if o.RoleIDn != nil {

		// query param role_id__n
		var qrRoleIDn string

		if o.RoleIDn != nil {
			qrRoleIDn = *o.RoleIDn
		}
		qRoleIDn := qrRoleIDn
		if qRoleIDn != "" {

			if err := r.SetQueryParam("role_id__n", qRoleIDn); err != nil {
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

	if o.Tenantn != nil {

		// query param tenant__n
		var qrTenantn string

		if o.Tenantn != nil {
			qrTenantn = *o.Tenantn
		}
		qTenantn := qrTenantn
		if qTenantn != "" {

			if err := r.SetQueryParam("tenant__n", qTenantn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroup != nil {

		// query param tenant_group
		var qrTenantGroup string

		if o.TenantGroup != nil {
			qrTenantGroup = *o.TenantGroup
		}
		qTenantGroup := qrTenantGroup
		if qTenantGroup != "" {

			if err := r.SetQueryParam("tenant_group", qTenantGroup); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupn != nil {

		// query param tenant_group__n
		var qrTenantGroupn string

		if o.TenantGroupn != nil {
			qrTenantGroupn = *o.TenantGroupn
		}
		qTenantGroupn := qrTenantGroupn
		if qTenantGroupn != "" {

			if err := r.SetQueryParam("tenant_group__n", qTenantGroupn); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupID != nil {

		// query param tenant_group_id
		var qrTenantGroupID string

		if o.TenantGroupID != nil {
			qrTenantGroupID = *o.TenantGroupID
		}
		qTenantGroupID := qrTenantGroupID
		if qTenantGroupID != "" {

			if err := r.SetQueryParam("tenant_group_id", qTenantGroupID); err != nil {
				return err
			}
		}
	}

	if o.TenantGroupIDn != nil {

		// query param tenant_group_id__n
		var qrTenantGroupIDn string

		if o.TenantGroupIDn != nil {
			qrTenantGroupIDn = *o.TenantGroupIDn
		}
		qTenantGroupIDn := qrTenantGroupIDn
		if qTenantGroupIDn != "" {

			if err := r.SetQueryParam("tenant_group_id__n", qTenantGroupIDn); err != nil {
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

	if o.TenantIDn != nil {

		// query param tenant_id__n
		var qrTenantIDn string

		if o.TenantIDn != nil {
			qrTenantIDn = *o.TenantIDn
		}
		qTenantIDn := qrTenantIDn
		if qTenantIDn != "" {

			if err := r.SetQueryParam("tenant_id__n", qTenantIDn); err != nil {
				return err
			}
		}
	}

	if o.VcPosition != nil {

		// query param vc_position
		var qrVcPosition string

		if o.VcPosition != nil {
			qrVcPosition = *o.VcPosition
		}
		qVcPosition := qrVcPosition
		if qVcPosition != "" {

			if err := r.SetQueryParam("vc_position", qVcPosition); err != nil {
				return err
			}
		}
	}

	if o.VcPositionGt != nil {

		// query param vc_position__gt
		var qrVcPositionGt string

		if o.VcPositionGt != nil {
			qrVcPositionGt = *o.VcPositionGt
		}
		qVcPositionGt := qrVcPositionGt
		if qVcPositionGt != "" {

			if err := r.SetQueryParam("vc_position__gt", qVcPositionGt); err != nil {
				return err
			}
		}
	}

	if o.VcPositionGte != nil {

		// query param vc_position__gte
		var qrVcPositionGte string

		if o.VcPositionGte != nil {
			qrVcPositionGte = *o.VcPositionGte
		}
		qVcPositionGte := qrVcPositionGte
		if qVcPositionGte != "" {

			if err := r.SetQueryParam("vc_position__gte", qVcPositionGte); err != nil {
				return err
			}
		}
	}

	if o.VcPositionLt != nil {

		// query param vc_position__lt
		var qrVcPositionLt string

		if o.VcPositionLt != nil {
			qrVcPositionLt = *o.VcPositionLt
		}
		qVcPositionLt := qrVcPositionLt
		if qVcPositionLt != "" {

			if err := r.SetQueryParam("vc_position__lt", qVcPositionLt); err != nil {
				return err
			}
		}
	}

	if o.VcPositionLte != nil {

		// query param vc_position__lte
		var qrVcPositionLte string

		if o.VcPositionLte != nil {
			qrVcPositionLte = *o.VcPositionLte
		}
		qVcPositionLte := qrVcPositionLte
		if qVcPositionLte != "" {

			if err := r.SetQueryParam("vc_position__lte", qVcPositionLte); err != nil {
				return err
			}
		}
	}

	if o.VcPositionn != nil {

		// query param vc_position__n
		var qrVcPositionn string

		if o.VcPositionn != nil {
			qrVcPositionn = *o.VcPositionn
		}
		qVcPositionn := qrVcPositionn
		if qVcPositionn != "" {

			if err := r.SetQueryParam("vc_position__n", qVcPositionn); err != nil {
				return err
			}
		}
	}

	if o.VcPriority != nil {

		// query param vc_priority
		var qrVcPriority string

		if o.VcPriority != nil {
			qrVcPriority = *o.VcPriority
		}
		qVcPriority := qrVcPriority
		if qVcPriority != "" {

			if err := r.SetQueryParam("vc_priority", qVcPriority); err != nil {
				return err
			}
		}
	}

	if o.VcPriorityGt != nil {

		// query param vc_priority__gt
		var qrVcPriorityGt string

		if o.VcPriorityGt != nil {
			qrVcPriorityGt = *o.VcPriorityGt
		}
		qVcPriorityGt := qrVcPriorityGt
		if qVcPriorityGt != "" {

			if err := r.SetQueryParam("vc_priority__gt", qVcPriorityGt); err != nil {
				return err
			}
		}
	}

	if o.VcPriorityGte != nil {

		// query param vc_priority__gte
		var qrVcPriorityGte string

		if o.VcPriorityGte != nil {
			qrVcPriorityGte = *o.VcPriorityGte
		}
		qVcPriorityGte := qrVcPriorityGte
		if qVcPriorityGte != "" {

			if err := r.SetQueryParam("vc_priority__gte", qVcPriorityGte); err != nil {
				return err
			}
		}
	}

	if o.VcPriorityLt != nil {

		// query param vc_priority__lt
		var qrVcPriorityLt string

		if o.VcPriorityLt != nil {
			qrVcPriorityLt = *o.VcPriorityLt
		}
		qVcPriorityLt := qrVcPriorityLt
		if qVcPriorityLt != "" {

			if err := r.SetQueryParam("vc_priority__lt", qVcPriorityLt); err != nil {
				return err
			}
		}
	}

	if o.VcPriorityLte != nil {

		// query param vc_priority__lte
		var qrVcPriorityLte string

		if o.VcPriorityLte != nil {
			qrVcPriorityLte = *o.VcPriorityLte
		}
		qVcPriorityLte := qrVcPriorityLte
		if qVcPriorityLte != "" {

			if err := r.SetQueryParam("vc_priority__lte", qVcPriorityLte); err != nil {
				return err
			}
		}
	}

	if o.VcPriorityn != nil {

		// query param vc_priority__n
		var qrVcPriorityn string

		if o.VcPriorityn != nil {
			qrVcPriorityn = *o.VcPriorityn
		}
		qVcPriorityn := qrVcPriorityn
		if qVcPriorityn != "" {

			if err := r.SetQueryParam("vc_priority__n", qVcPriorityn); err != nil {
				return err
			}
		}
	}

	if o.VirtualChassisID != nil {

		// query param virtual_chassis_id
		var qrVirtualChassisID string

		if o.VirtualChassisID != nil {
			qrVirtualChassisID = *o.VirtualChassisID
		}
		qVirtualChassisID := qrVirtualChassisID
		if qVirtualChassisID != "" {

			if err := r.SetQueryParam("virtual_chassis_id", qVirtualChassisID); err != nil {
				return err
			}
		}
	}

	if o.VirtualChassisIDn != nil {

		// query param virtual_chassis_id__n
		var qrVirtualChassisIDn string

		if o.VirtualChassisIDn != nil {
			qrVirtualChassisIDn = *o.VirtualChassisIDn
		}
		qVirtualChassisIDn := qrVirtualChassisIDn
		if qVirtualChassisIDn != "" {

			if err := r.SetQueryParam("virtual_chassis_id__n", qVirtualChassisIDn); err != nil {
				return err
			}
		}
	}

	if o.VirtualChassisMember != nil {

		// query param virtual_chassis_member
		var qrVirtualChassisMember string

		if o.VirtualChassisMember != nil {
			qrVirtualChassisMember = *o.VirtualChassisMember
		}
		qVirtualChassisMember := qrVirtualChassisMember
		if qVirtualChassisMember != "" {

			if err := r.SetQueryParam("virtual_chassis_member", qVirtualChassisMember); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
