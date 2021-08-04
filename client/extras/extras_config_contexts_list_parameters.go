package extras

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

// NewExtrasConfigContextsListParams creates a new ExtrasConfigContextsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasConfigContextsListParams() *ExtrasConfigContextsListParams {
	return &ExtrasConfigContextsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasConfigContextsListParamsWithTimeout creates a new ExtrasConfigContextsListParams object
// with the ability to set a timeout on a request.
func NewExtrasConfigContextsListParamsWithTimeout(timeout time.Duration) *ExtrasConfigContextsListParams {
	return &ExtrasConfigContextsListParams{
		timeout: timeout,
	}
}

// NewExtrasConfigContextsListParamsWithContext creates a new ExtrasConfigContextsListParams object
// with the ability to set a context for a request.
func NewExtrasConfigContextsListParamsWithContext(ctx context.Context) *ExtrasConfigContextsListParams {
	return &ExtrasConfigContextsListParams{
		Context: ctx,
	}
}

// NewExtrasConfigContextsListParamsWithHTTPClient creates a new ExtrasConfigContextsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasConfigContextsListParamsWithHTTPClient(client *http.Client) *ExtrasConfigContextsListParams {
	return &ExtrasConfigContextsListParams{
		HTTPClient: client,
	}
}

/* ExtrasConfigContextsListParams contains all the parameters to send to the API endpoint
   for the extras config contexts list operation.

   Typically these are written to a http.Request.
*/
type ExtrasConfigContextsListParams struct {

	// ClusterGroup.
	ClusterGroup *string

	// ClusterGroupn.
	ClusterGroupn *string

	// ClusterGroupID.
	ClusterGroupID *string

	// ClusterGroupIDn.
	ClusterGroupIDn *string

	// ClusterID.
	ClusterID *string

	// ClusterIDn.
	ClusterIDn *string

	// DeviceType.
	DeviceType *string

	// DeviceTypen.
	DeviceTypen *string

	// DeviceTypeID.
	DeviceTypeID *string

	// DeviceTypeIDn.
	DeviceTypeIDn *string

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

	// IsActive.
	IsActive *string

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

	// OwnerContentType.
	OwnerContentType *string

	// OwnerContentTypen.
	OwnerContentTypen *string

	// OwnerObjectID.
	OwnerObjectID *string

	// OwnerObjectIDIc.
	OwnerObjectIDIc *string

	// OwnerObjectIDIe.
	OwnerObjectIDIe *string

	// OwnerObjectIDIew.
	OwnerObjectIDIew *string

	// OwnerObjectIDIsw.
	OwnerObjectIDIsw *string

	// OwnerObjectIDn.
	OwnerObjectIDn *string

	// OwnerObjectIDNic.
	OwnerObjectIDNic *string

	// OwnerObjectIDNie.
	OwnerObjectIDNie *string

	// OwnerObjectIDNiew.
	OwnerObjectIDNiew *string

	// OwnerObjectIDNisw.
	OwnerObjectIDNisw *string

	// Platform.
	Platform *string

	// Platformn.
	Platformn *string

	// PlatformID.
	PlatformID *string

	// PlatformIDn.
	PlatformIDn *string

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

	// Role.
	Role *string

	// Rolen.
	Rolen *string

	// RoleID.
	RoleID *string

	// RoleIDn.
	RoleIDn *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras config contexts list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsListParams) WithDefaults() *ExtrasConfigContextsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras config contexts list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasConfigContextsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTimeout(timeout time.Duration) *ExtrasConfigContextsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithContext(ctx context.Context) *ExtrasConfigContextsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithHTTPClient(client *http.Client) *ExtrasConfigContextsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterGroup adds the clusterGroup to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithClusterGroup(clusterGroup *string) *ExtrasConfigContextsListParams {
	o.SetClusterGroup(clusterGroup)
	return o
}

// SetClusterGroup adds the clusterGroup to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetClusterGroup(clusterGroup *string) {
	o.ClusterGroup = clusterGroup
}

// WithClusterGroupn adds the clusterGroupn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithClusterGroupn(clusterGroupn *string) *ExtrasConfigContextsListParams {
	o.SetClusterGroupn(clusterGroupn)
	return o
}

// SetClusterGroupn adds the clusterGroupN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetClusterGroupn(clusterGroupn *string) {
	o.ClusterGroupn = clusterGroupn
}

// WithClusterGroupID adds the clusterGroupID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithClusterGroupID(clusterGroupID *string) *ExtrasConfigContextsListParams {
	o.SetClusterGroupID(clusterGroupID)
	return o
}

// SetClusterGroupID adds the clusterGroupId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetClusterGroupID(clusterGroupID *string) {
	o.ClusterGroupID = clusterGroupID
}

// WithClusterGroupIDn adds the clusterGroupIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithClusterGroupIDn(clusterGroupIDn *string) *ExtrasConfigContextsListParams {
	o.SetClusterGroupIDn(clusterGroupIDn)
	return o
}

// SetClusterGroupIDn adds the clusterGroupIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetClusterGroupIDn(clusterGroupIDn *string) {
	o.ClusterGroupIDn = clusterGroupIDn
}

// WithClusterID adds the clusterID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithClusterID(clusterID *string) *ExtrasConfigContextsListParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetClusterID(clusterID *string) {
	o.ClusterID = clusterID
}

// WithClusterIDn adds the clusterIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithClusterIDn(clusterIDn *string) *ExtrasConfigContextsListParams {
	o.SetClusterIDn(clusterIDn)
	return o
}

// SetClusterIDn adds the clusterIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetClusterIDn(clusterIDn *string) {
	o.ClusterIDn = clusterIDn
}

// WithDeviceType adds the deviceType to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithDeviceType(deviceType *string) *ExtrasConfigContextsListParams {
	o.SetDeviceType(deviceType)
	return o
}

// SetDeviceType adds the deviceType to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetDeviceType(deviceType *string) {
	o.DeviceType = deviceType
}

// WithDeviceTypen adds the deviceTypen to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithDeviceTypen(deviceTypen *string) *ExtrasConfigContextsListParams {
	o.SetDeviceTypen(deviceTypen)
	return o
}

// SetDeviceTypen adds the deviceTypeN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetDeviceTypen(deviceTypen *string) {
	o.DeviceTypen = deviceTypen
}

// WithDeviceTypeID adds the deviceTypeID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithDeviceTypeID(deviceTypeID *string) *ExtrasConfigContextsListParams {
	o.SetDeviceTypeID(deviceTypeID)
	return o
}

// SetDeviceTypeID adds the deviceTypeId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetDeviceTypeID(deviceTypeID *string) {
	o.DeviceTypeID = deviceTypeID
}

// WithDeviceTypeIDn adds the deviceTypeIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithDeviceTypeIDn(deviceTypeIDn *string) *ExtrasConfigContextsListParams {
	o.SetDeviceTypeIDn(deviceTypeIDn)
	return o
}

// SetDeviceTypeIDn adds the deviceTypeIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetDeviceTypeIDn(deviceTypeIDn *string) {
	o.DeviceTypeIDn = deviceTypeIDn
}

// WithID adds the id to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithID(id *string) *ExtrasConfigContextsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDIc(iDIc *string) *ExtrasConfigContextsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDIe(iDIe *string) *ExtrasConfigContextsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDIew(iDIew *string) *ExtrasConfigContextsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDIsw(iDIsw *string) *ExtrasConfigContextsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDn(iDn *string) *ExtrasConfigContextsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDNic(iDNic *string) *ExtrasConfigContextsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDNie(iDNie *string) *ExtrasConfigContextsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDNiew(iDNiew *string) *ExtrasConfigContextsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIDNisw(iDNisw *string) *ExtrasConfigContextsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithIsActive adds the isActive to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithIsActive(isActive *string) *ExtrasConfigContextsListParams {
	o.SetIsActive(isActive)
	return o
}

// SetIsActive adds the isActive to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetIsActive(isActive *string) {
	o.IsActive = isActive
}

// WithLimit adds the limit to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithLimit(limit *int64) *ExtrasConfigContextsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithName(name *string) *ExtrasConfigContextsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameIc(nameIc *string) *ExtrasConfigContextsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameIe(nameIe *string) *ExtrasConfigContextsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameIew(nameIew *string) *ExtrasConfigContextsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameIsw(nameIsw *string) *ExtrasConfigContextsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNamen(namen *string) *ExtrasConfigContextsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameNic(nameNic *string) *ExtrasConfigContextsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameNie(nameNie *string) *ExtrasConfigContextsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameNiew(nameNiew *string) *ExtrasConfigContextsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithNameNisw(nameNisw *string) *ExtrasConfigContextsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOffset(offset *int64) *ExtrasConfigContextsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwnerContentType adds the ownerContentType to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerContentType(ownerContentType *string) *ExtrasConfigContextsListParams {
	o.SetOwnerContentType(ownerContentType)
	return o
}

// SetOwnerContentType adds the ownerContentType to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerContentType(ownerContentType *string) {
	o.OwnerContentType = ownerContentType
}

// WithOwnerContentTypen adds the ownerContentTypen to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerContentTypen(ownerContentTypen *string) *ExtrasConfigContextsListParams {
	o.SetOwnerContentTypen(ownerContentTypen)
	return o
}

// SetOwnerContentTypen adds the ownerContentTypeN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerContentTypen(ownerContentTypen *string) {
	o.OwnerContentTypen = ownerContentTypen
}

// WithOwnerObjectID adds the ownerObjectID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectID(ownerObjectID *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectID(ownerObjectID)
	return o
}

// SetOwnerObjectID adds the ownerObjectId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectID(ownerObjectID *string) {
	o.OwnerObjectID = ownerObjectID
}

// WithOwnerObjectIDIc adds the ownerObjectIDIc to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDIc(ownerObjectIDIc *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDIc(ownerObjectIDIc)
	return o
}

// SetOwnerObjectIDIc adds the ownerObjectIdIc to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDIc(ownerObjectIDIc *string) {
	o.OwnerObjectIDIc = ownerObjectIDIc
}

// WithOwnerObjectIDIe adds the ownerObjectIDIe to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDIe(ownerObjectIDIe *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDIe(ownerObjectIDIe)
	return o
}

// SetOwnerObjectIDIe adds the ownerObjectIdIe to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDIe(ownerObjectIDIe *string) {
	o.OwnerObjectIDIe = ownerObjectIDIe
}

// WithOwnerObjectIDIew adds the ownerObjectIDIew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDIew(ownerObjectIDIew *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDIew(ownerObjectIDIew)
	return o
}

// SetOwnerObjectIDIew adds the ownerObjectIdIew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDIew(ownerObjectIDIew *string) {
	o.OwnerObjectIDIew = ownerObjectIDIew
}

// WithOwnerObjectIDIsw adds the ownerObjectIDIsw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDIsw(ownerObjectIDIsw *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDIsw(ownerObjectIDIsw)
	return o
}

// SetOwnerObjectIDIsw adds the ownerObjectIdIsw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDIsw(ownerObjectIDIsw *string) {
	o.OwnerObjectIDIsw = ownerObjectIDIsw
}

// WithOwnerObjectIDn adds the ownerObjectIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDn(ownerObjectIDn *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDn(ownerObjectIDn)
	return o
}

// SetOwnerObjectIDn adds the ownerObjectIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDn(ownerObjectIDn *string) {
	o.OwnerObjectIDn = ownerObjectIDn
}

// WithOwnerObjectIDNic adds the ownerObjectIDNic to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDNic(ownerObjectIDNic *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDNic(ownerObjectIDNic)
	return o
}

// SetOwnerObjectIDNic adds the ownerObjectIdNic to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDNic(ownerObjectIDNic *string) {
	o.OwnerObjectIDNic = ownerObjectIDNic
}

// WithOwnerObjectIDNie adds the ownerObjectIDNie to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDNie(ownerObjectIDNie *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDNie(ownerObjectIDNie)
	return o
}

// SetOwnerObjectIDNie adds the ownerObjectIdNie to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDNie(ownerObjectIDNie *string) {
	o.OwnerObjectIDNie = ownerObjectIDNie
}

// WithOwnerObjectIDNiew adds the ownerObjectIDNiew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDNiew(ownerObjectIDNiew *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDNiew(ownerObjectIDNiew)
	return o
}

// SetOwnerObjectIDNiew adds the ownerObjectIdNiew to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDNiew(ownerObjectIDNiew *string) {
	o.OwnerObjectIDNiew = ownerObjectIDNiew
}

// WithOwnerObjectIDNisw adds the ownerObjectIDNisw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithOwnerObjectIDNisw(ownerObjectIDNisw *string) *ExtrasConfigContextsListParams {
	o.SetOwnerObjectIDNisw(ownerObjectIDNisw)
	return o
}

// SetOwnerObjectIDNisw adds the ownerObjectIdNisw to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetOwnerObjectIDNisw(ownerObjectIDNisw *string) {
	o.OwnerObjectIDNisw = ownerObjectIDNisw
}

// WithPlatform adds the platform to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithPlatform(platform *string) *ExtrasConfigContextsListParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPlatformn adds the platformn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithPlatformn(platformn *string) *ExtrasConfigContextsListParams {
	o.SetPlatformn(platformn)
	return o
}

// SetPlatformn adds the platformN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetPlatformn(platformn *string) {
	o.Platformn = platformn
}

// WithPlatformID adds the platformID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithPlatformID(platformID *string) *ExtrasConfigContextsListParams {
	o.SetPlatformID(platformID)
	return o
}

// SetPlatformID adds the platformId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetPlatformID(platformID *string) {
	o.PlatformID = platformID
}

// WithPlatformIDn adds the platformIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithPlatformIDn(platformIDn *string) *ExtrasConfigContextsListParams {
	o.SetPlatformIDn(platformIDn)
	return o
}

// SetPlatformIDn adds the platformIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetPlatformIDn(platformIDn *string) {
	o.PlatformIDn = platformIDn
}

// WithQ adds the q to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithQ(q *string) *ExtrasConfigContextsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRegion(region *string) *ExtrasConfigContextsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRegionn(regionn *string) *ExtrasConfigContextsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRegionID(regionID *string) *ExtrasConfigContextsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRegionIDn(regionIDn *string) *ExtrasConfigContextsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithRole adds the role to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRole(role *string) *ExtrasConfigContextsListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRolen(rolen *string) *ExtrasConfigContextsListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRoleID(roleID *string) *ExtrasConfigContextsListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithRoleIDn(roleIDn *string) *ExtrasConfigContextsListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithSite adds the site to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithSite(site *string) *ExtrasConfigContextsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithSiten(siten *string) *ExtrasConfigContextsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithSiteID(siteID *string) *ExtrasConfigContextsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithSiteIDn(siteIDn *string) *ExtrasConfigContextsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTag(tag *string) *ExtrasConfigContextsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTagn(tagn *string) *ExtrasConfigContextsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenant(tenant *string) *ExtrasConfigContextsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenantn(tenantn *string) *ExtrasConfigContextsListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenantGroup(tenantGroup *string) *ExtrasConfigContextsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenantGroupn(tenantGroupn *string) *ExtrasConfigContextsListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenantGroupID(tenantGroupID *string) *ExtrasConfigContextsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenantGroupIDn(tenantGroupIDn *string) *ExtrasConfigContextsListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenantID(tenantID *string) *ExtrasConfigContextsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) WithTenantIDn(tenantIDn *string) *ExtrasConfigContextsListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the extras config contexts list params
func (o *ExtrasConfigContextsListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasConfigContextsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterGroup != nil {

		// query param cluster_group
		var qrClusterGroup string

		if o.ClusterGroup != nil {
			qrClusterGroup = *o.ClusterGroup
		}
		qClusterGroup := qrClusterGroup
		if qClusterGroup != "" {

			if err := r.SetQueryParam("cluster_group", qClusterGroup); err != nil {
				return err
			}
		}
	}

	if o.ClusterGroupn != nil {

		// query param cluster_group__n
		var qrClusterGroupn string

		if o.ClusterGroupn != nil {
			qrClusterGroupn = *o.ClusterGroupn
		}
		qClusterGroupn := qrClusterGroupn
		if qClusterGroupn != "" {

			if err := r.SetQueryParam("cluster_group__n", qClusterGroupn); err != nil {
				return err
			}
		}
	}

	if o.ClusterGroupID != nil {

		// query param cluster_group_id
		var qrClusterGroupID string

		if o.ClusterGroupID != nil {
			qrClusterGroupID = *o.ClusterGroupID
		}
		qClusterGroupID := qrClusterGroupID
		if qClusterGroupID != "" {

			if err := r.SetQueryParam("cluster_group_id", qClusterGroupID); err != nil {
				return err
			}
		}
	}

	if o.ClusterGroupIDn != nil {

		// query param cluster_group_id__n
		var qrClusterGroupIDn string

		if o.ClusterGroupIDn != nil {
			qrClusterGroupIDn = *o.ClusterGroupIDn
		}
		qClusterGroupIDn := qrClusterGroupIDn
		if qClusterGroupIDn != "" {

			if err := r.SetQueryParam("cluster_group_id__n", qClusterGroupIDn); err != nil {
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

	if o.DeviceType != nil {

		// query param device_type
		var qrDeviceType string

		if o.DeviceType != nil {
			qrDeviceType = *o.DeviceType
		}
		qDeviceType := qrDeviceType
		if qDeviceType != "" {

			if err := r.SetQueryParam("device_type", qDeviceType); err != nil {
				return err
			}
		}
	}

	if o.DeviceTypen != nil {

		// query param device_type__n
		var qrDeviceTypen string

		if o.DeviceTypen != nil {
			qrDeviceTypen = *o.DeviceTypen
		}
		qDeviceTypen := qrDeviceTypen
		if qDeviceTypen != "" {

			if err := r.SetQueryParam("device_type__n", qDeviceTypen); err != nil {
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

	if o.IsActive != nil {

		// query param is_active
		var qrIsActive string

		if o.IsActive != nil {
			qrIsActive = *o.IsActive
		}
		qIsActive := qrIsActive
		if qIsActive != "" {

			if err := r.SetQueryParam("is_active", qIsActive); err != nil {
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

	if o.OwnerContentType != nil {

		// query param owner_content_type
		var qrOwnerContentType string

		if o.OwnerContentType != nil {
			qrOwnerContentType = *o.OwnerContentType
		}
		qOwnerContentType := qrOwnerContentType
		if qOwnerContentType != "" {

			if err := r.SetQueryParam("owner_content_type", qOwnerContentType); err != nil {
				return err
			}
		}
	}

	if o.OwnerContentTypen != nil {

		// query param owner_content_type__n
		var qrOwnerContentTypen string

		if o.OwnerContentTypen != nil {
			qrOwnerContentTypen = *o.OwnerContentTypen
		}
		qOwnerContentTypen := qrOwnerContentTypen
		if qOwnerContentTypen != "" {

			if err := r.SetQueryParam("owner_content_type__n", qOwnerContentTypen); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectID != nil {

		// query param owner_object_id
		var qrOwnerObjectID string

		if o.OwnerObjectID != nil {
			qrOwnerObjectID = *o.OwnerObjectID
		}
		qOwnerObjectID := qrOwnerObjectID
		if qOwnerObjectID != "" {

			if err := r.SetQueryParam("owner_object_id", qOwnerObjectID); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIc != nil {

		// query param owner_object_id__ic
		var qrOwnerObjectIDIc string

		if o.OwnerObjectIDIc != nil {
			qrOwnerObjectIDIc = *o.OwnerObjectIDIc
		}
		qOwnerObjectIDIc := qrOwnerObjectIDIc
		if qOwnerObjectIDIc != "" {

			if err := r.SetQueryParam("owner_object_id__ic", qOwnerObjectIDIc); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIe != nil {

		// query param owner_object_id__ie
		var qrOwnerObjectIDIe string

		if o.OwnerObjectIDIe != nil {
			qrOwnerObjectIDIe = *o.OwnerObjectIDIe
		}
		qOwnerObjectIDIe := qrOwnerObjectIDIe
		if qOwnerObjectIDIe != "" {

			if err := r.SetQueryParam("owner_object_id__ie", qOwnerObjectIDIe); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIew != nil {

		// query param owner_object_id__iew
		var qrOwnerObjectIDIew string

		if o.OwnerObjectIDIew != nil {
			qrOwnerObjectIDIew = *o.OwnerObjectIDIew
		}
		qOwnerObjectIDIew := qrOwnerObjectIDIew
		if qOwnerObjectIDIew != "" {

			if err := r.SetQueryParam("owner_object_id__iew", qOwnerObjectIDIew); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDIsw != nil {

		// query param owner_object_id__isw
		var qrOwnerObjectIDIsw string

		if o.OwnerObjectIDIsw != nil {
			qrOwnerObjectIDIsw = *o.OwnerObjectIDIsw
		}
		qOwnerObjectIDIsw := qrOwnerObjectIDIsw
		if qOwnerObjectIDIsw != "" {

			if err := r.SetQueryParam("owner_object_id__isw", qOwnerObjectIDIsw); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDn != nil {

		// query param owner_object_id__n
		var qrOwnerObjectIDn string

		if o.OwnerObjectIDn != nil {
			qrOwnerObjectIDn = *o.OwnerObjectIDn
		}
		qOwnerObjectIDn := qrOwnerObjectIDn
		if qOwnerObjectIDn != "" {

			if err := r.SetQueryParam("owner_object_id__n", qOwnerObjectIDn); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNic != nil {

		// query param owner_object_id__nic
		var qrOwnerObjectIDNic string

		if o.OwnerObjectIDNic != nil {
			qrOwnerObjectIDNic = *o.OwnerObjectIDNic
		}
		qOwnerObjectIDNic := qrOwnerObjectIDNic
		if qOwnerObjectIDNic != "" {

			if err := r.SetQueryParam("owner_object_id__nic", qOwnerObjectIDNic); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNie != nil {

		// query param owner_object_id__nie
		var qrOwnerObjectIDNie string

		if o.OwnerObjectIDNie != nil {
			qrOwnerObjectIDNie = *o.OwnerObjectIDNie
		}
		qOwnerObjectIDNie := qrOwnerObjectIDNie
		if qOwnerObjectIDNie != "" {

			if err := r.SetQueryParam("owner_object_id__nie", qOwnerObjectIDNie); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNiew != nil {

		// query param owner_object_id__niew
		var qrOwnerObjectIDNiew string

		if o.OwnerObjectIDNiew != nil {
			qrOwnerObjectIDNiew = *o.OwnerObjectIDNiew
		}
		qOwnerObjectIDNiew := qrOwnerObjectIDNiew
		if qOwnerObjectIDNiew != "" {

			if err := r.SetQueryParam("owner_object_id__niew", qOwnerObjectIDNiew); err != nil {
				return err
			}
		}
	}

	if o.OwnerObjectIDNisw != nil {

		// query param owner_object_id__nisw
		var qrOwnerObjectIDNisw string

		if o.OwnerObjectIDNisw != nil {
			qrOwnerObjectIDNisw = *o.OwnerObjectIDNisw
		}
		qOwnerObjectIDNisw := qrOwnerObjectIDNisw
		if qOwnerObjectIDNisw != "" {

			if err := r.SetQueryParam("owner_object_id__nisw", qOwnerObjectIDNisw); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
