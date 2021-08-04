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

// NewDcimRacksListParams creates a new DcimRacksListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRacksListParams() *DcimRacksListParams {
	return &DcimRacksListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRacksListParamsWithTimeout creates a new DcimRacksListParams object
// with the ability to set a timeout on a request.
func NewDcimRacksListParamsWithTimeout(timeout time.Duration) *DcimRacksListParams {
	return &DcimRacksListParams{
		timeout: timeout,
	}
}

// NewDcimRacksListParamsWithContext creates a new DcimRacksListParams object
// with the ability to set a context for a request.
func NewDcimRacksListParamsWithContext(ctx context.Context) *DcimRacksListParams {
	return &DcimRacksListParams{
		Context: ctx,
	}
}

// NewDcimRacksListParamsWithHTTPClient creates a new DcimRacksListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRacksListParamsWithHTTPClient(client *http.Client) *DcimRacksListParams {
	return &DcimRacksListParams{
		HTTPClient: client,
	}
}

/* DcimRacksListParams contains all the parameters to send to the API endpoint
   for the dcim racks list operation.

   Typically these are written to a http.Request.
*/
type DcimRacksListParams struct {

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

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DescUnits.
	DescUnits *string

	// FacilityID.
	FacilityID *string

	// FacilityIDIc.
	FacilityIDIc *string

	// FacilityIDIe.
	FacilityIDIe *string

	// FacilityIDIew.
	FacilityIDIew *string

	// FacilityIDIsw.
	FacilityIDIsw *string

	// FacilityIDn.
	FacilityIDn *string

	// FacilityIDNic.
	FacilityIDNic *string

	// FacilityIDNie.
	FacilityIDNie *string

	// FacilityIDNiew.
	FacilityIDNiew *string

	// FacilityIDNisw.
	FacilityIDNisw *string

	// Group.
	Group *string

	// Groupn.
	Groupn *string

	// GroupID.
	GroupID *string

	// GroupIDn.
	GroupIDn *string

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

	// OuterDepth.
	OuterDepth *string

	// OuterDepthGt.
	OuterDepthGt *string

	// OuterDepthGte.
	OuterDepthGte *string

	// OuterDepthLt.
	OuterDepthLt *string

	// OuterDepthLte.
	OuterDepthLte *string

	// OuterDepthn.
	OuterDepthn *string

	// OuterUnit.
	OuterUnit *string

	// OuterUnitn.
	OuterUnitn *string

	// OuterWidth.
	OuterWidth *string

	// OuterWidthGt.
	OuterWidthGt *string

	// OuterWidthGte.
	OuterWidthGte *string

	// OuterWidthLt.
	OuterWidthLt *string

	// OuterWidthLte.
	OuterWidthLte *string

	// OuterWidthn.
	OuterWidthn *string

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

	// Type.
	Type *string

	// Typen.
	Typen *string

	// UHeight.
	UHeight *string

	// UHeightGt.
	UHeightGt *string

	// UHeightGte.
	UHeightGte *string

	// UHeightLt.
	UHeightLt *string

	// UHeightLte.
	UHeightLte *string

	// UHeightn.
	UHeightn *string

	// Width.
	Width *string

	// Widthn.
	Widthn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim racks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksListParams) WithDefaults() *DcimRacksListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim racks list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRacksListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim racks list params
func (o *DcimRacksListParams) WithTimeout(timeout time.Duration) *DcimRacksListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim racks list params
func (o *DcimRacksListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim racks list params
func (o *DcimRacksListParams) WithContext(ctx context.Context) *DcimRacksListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim racks list params
func (o *DcimRacksListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim racks list params
func (o *DcimRacksListParams) WithHTTPClient(client *http.Client) *DcimRacksListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim racks list params
func (o *DcimRacksListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssetTag adds the assetTag to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTag(assetTag *string) *DcimRacksListParams {
	o.SetAssetTag(assetTag)
	return o
}

// SetAssetTag adds the assetTag to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTag(assetTag *string) {
	o.AssetTag = assetTag
}

// WithAssetTagIc adds the assetTagIc to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagIc(assetTagIc *string) *DcimRacksListParams {
	o.SetAssetTagIc(assetTagIc)
	return o
}

// SetAssetTagIc adds the assetTagIc to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagIc(assetTagIc *string) {
	o.AssetTagIc = assetTagIc
}

// WithAssetTagIe adds the assetTagIe to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagIe(assetTagIe *string) *DcimRacksListParams {
	o.SetAssetTagIe(assetTagIe)
	return o
}

// SetAssetTagIe adds the assetTagIe to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagIe(assetTagIe *string) {
	o.AssetTagIe = assetTagIe
}

// WithAssetTagIew adds the assetTagIew to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagIew(assetTagIew *string) *DcimRacksListParams {
	o.SetAssetTagIew(assetTagIew)
	return o
}

// SetAssetTagIew adds the assetTagIew to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagIew(assetTagIew *string) {
	o.AssetTagIew = assetTagIew
}

// WithAssetTagIsw adds the assetTagIsw to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagIsw(assetTagIsw *string) *DcimRacksListParams {
	o.SetAssetTagIsw(assetTagIsw)
	return o
}

// SetAssetTagIsw adds the assetTagIsw to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagIsw(assetTagIsw *string) {
	o.AssetTagIsw = assetTagIsw
}

// WithAssetTagn adds the assetTagn to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagn(assetTagn *string) *DcimRacksListParams {
	o.SetAssetTagn(assetTagn)
	return o
}

// SetAssetTagn adds the assetTagN to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagn(assetTagn *string) {
	o.AssetTagn = assetTagn
}

// WithAssetTagNic adds the assetTagNic to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagNic(assetTagNic *string) *DcimRacksListParams {
	o.SetAssetTagNic(assetTagNic)
	return o
}

// SetAssetTagNic adds the assetTagNic to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagNic(assetTagNic *string) {
	o.AssetTagNic = assetTagNic
}

// WithAssetTagNie adds the assetTagNie to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagNie(assetTagNie *string) *DcimRacksListParams {
	o.SetAssetTagNie(assetTagNie)
	return o
}

// SetAssetTagNie adds the assetTagNie to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagNie(assetTagNie *string) {
	o.AssetTagNie = assetTagNie
}

// WithAssetTagNiew adds the assetTagNiew to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagNiew(assetTagNiew *string) *DcimRacksListParams {
	o.SetAssetTagNiew(assetTagNiew)
	return o
}

// SetAssetTagNiew adds the assetTagNiew to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagNiew(assetTagNiew *string) {
	o.AssetTagNiew = assetTagNiew
}

// WithAssetTagNisw adds the assetTagNisw to the dcim racks list params
func (o *DcimRacksListParams) WithAssetTagNisw(assetTagNisw *string) *DcimRacksListParams {
	o.SetAssetTagNisw(assetTagNisw)
	return o
}

// SetAssetTagNisw adds the assetTagNisw to the dcim racks list params
func (o *DcimRacksListParams) SetAssetTagNisw(assetTagNisw *string) {
	o.AssetTagNisw = assetTagNisw
}

// WithCreated adds the created to the dcim racks list params
func (o *DcimRacksListParams) WithCreated(created *string) *DcimRacksListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim racks list params
func (o *DcimRacksListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim racks list params
func (o *DcimRacksListParams) WithCreatedGte(createdGte *string) *DcimRacksListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim racks list params
func (o *DcimRacksListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim racks list params
func (o *DcimRacksListParams) WithCreatedLte(createdLte *string) *DcimRacksListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim racks list params
func (o *DcimRacksListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescUnits adds the descUnits to the dcim racks list params
func (o *DcimRacksListParams) WithDescUnits(descUnits *string) *DcimRacksListParams {
	o.SetDescUnits(descUnits)
	return o
}

// SetDescUnits adds the descUnits to the dcim racks list params
func (o *DcimRacksListParams) SetDescUnits(descUnits *string) {
	o.DescUnits = descUnits
}

// WithFacilityID adds the facilityID to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityID(facilityID *string) *DcimRacksListParams {
	o.SetFacilityID(facilityID)
	return o
}

// SetFacilityID adds the facilityId to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityID(facilityID *string) {
	o.FacilityID = facilityID
}

// WithFacilityIDIc adds the facilityIDIc to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDIc(facilityIDIc *string) *DcimRacksListParams {
	o.SetFacilityIDIc(facilityIDIc)
	return o
}

// SetFacilityIDIc adds the facilityIdIc to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDIc(facilityIDIc *string) {
	o.FacilityIDIc = facilityIDIc
}

// WithFacilityIDIe adds the facilityIDIe to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDIe(facilityIDIe *string) *DcimRacksListParams {
	o.SetFacilityIDIe(facilityIDIe)
	return o
}

// SetFacilityIDIe adds the facilityIdIe to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDIe(facilityIDIe *string) {
	o.FacilityIDIe = facilityIDIe
}

// WithFacilityIDIew adds the facilityIDIew to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDIew(facilityIDIew *string) *DcimRacksListParams {
	o.SetFacilityIDIew(facilityIDIew)
	return o
}

// SetFacilityIDIew adds the facilityIdIew to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDIew(facilityIDIew *string) {
	o.FacilityIDIew = facilityIDIew
}

// WithFacilityIDIsw adds the facilityIDIsw to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDIsw(facilityIDIsw *string) *DcimRacksListParams {
	o.SetFacilityIDIsw(facilityIDIsw)
	return o
}

// SetFacilityIDIsw adds the facilityIdIsw to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDIsw(facilityIDIsw *string) {
	o.FacilityIDIsw = facilityIDIsw
}

// WithFacilityIDn adds the facilityIDn to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDn(facilityIDn *string) *DcimRacksListParams {
	o.SetFacilityIDn(facilityIDn)
	return o
}

// SetFacilityIDn adds the facilityIdN to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDn(facilityIDn *string) {
	o.FacilityIDn = facilityIDn
}

// WithFacilityIDNic adds the facilityIDNic to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDNic(facilityIDNic *string) *DcimRacksListParams {
	o.SetFacilityIDNic(facilityIDNic)
	return o
}

// SetFacilityIDNic adds the facilityIdNic to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDNic(facilityIDNic *string) {
	o.FacilityIDNic = facilityIDNic
}

// WithFacilityIDNie adds the facilityIDNie to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDNie(facilityIDNie *string) *DcimRacksListParams {
	o.SetFacilityIDNie(facilityIDNie)
	return o
}

// SetFacilityIDNie adds the facilityIdNie to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDNie(facilityIDNie *string) {
	o.FacilityIDNie = facilityIDNie
}

// WithFacilityIDNiew adds the facilityIDNiew to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDNiew(facilityIDNiew *string) *DcimRacksListParams {
	o.SetFacilityIDNiew(facilityIDNiew)
	return o
}

// SetFacilityIDNiew adds the facilityIdNiew to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDNiew(facilityIDNiew *string) {
	o.FacilityIDNiew = facilityIDNiew
}

// WithFacilityIDNisw adds the facilityIDNisw to the dcim racks list params
func (o *DcimRacksListParams) WithFacilityIDNisw(facilityIDNisw *string) *DcimRacksListParams {
	o.SetFacilityIDNisw(facilityIDNisw)
	return o
}

// SetFacilityIDNisw adds the facilityIdNisw to the dcim racks list params
func (o *DcimRacksListParams) SetFacilityIDNisw(facilityIDNisw *string) {
	o.FacilityIDNisw = facilityIDNisw
}

// WithGroup adds the group to the dcim racks list params
func (o *DcimRacksListParams) WithGroup(group *string) *DcimRacksListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the dcim racks list params
func (o *DcimRacksListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupn adds the groupn to the dcim racks list params
func (o *DcimRacksListParams) WithGroupn(groupn *string) *DcimRacksListParams {
	o.SetGroupn(groupn)
	return o
}

// SetGroupn adds the groupN to the dcim racks list params
func (o *DcimRacksListParams) SetGroupn(groupn *string) {
	o.Groupn = groupn
}

// WithGroupID adds the groupID to the dcim racks list params
func (o *DcimRacksListParams) WithGroupID(groupID *string) *DcimRacksListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the dcim racks list params
func (o *DcimRacksListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithGroupIDn adds the groupIDn to the dcim racks list params
func (o *DcimRacksListParams) WithGroupIDn(groupIDn *string) *DcimRacksListParams {
	o.SetGroupIDn(groupIDn)
	return o
}

// SetGroupIDn adds the groupIdN to the dcim racks list params
func (o *DcimRacksListParams) SetGroupIDn(groupIDn *string) {
	o.GroupIDn = groupIDn
}

// WithID adds the id to the dcim racks list params
func (o *DcimRacksListParams) WithID(id *string) *DcimRacksListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim racks list params
func (o *DcimRacksListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim racks list params
func (o *DcimRacksListParams) WithIDIc(iDIc *string) *DcimRacksListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim racks list params
func (o *DcimRacksListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim racks list params
func (o *DcimRacksListParams) WithIDIe(iDIe *string) *DcimRacksListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim racks list params
func (o *DcimRacksListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim racks list params
func (o *DcimRacksListParams) WithIDIew(iDIew *string) *DcimRacksListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim racks list params
func (o *DcimRacksListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim racks list params
func (o *DcimRacksListParams) WithIDIsw(iDIsw *string) *DcimRacksListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim racks list params
func (o *DcimRacksListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim racks list params
func (o *DcimRacksListParams) WithIDn(iDn *string) *DcimRacksListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim racks list params
func (o *DcimRacksListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim racks list params
func (o *DcimRacksListParams) WithIDNic(iDNic *string) *DcimRacksListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim racks list params
func (o *DcimRacksListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim racks list params
func (o *DcimRacksListParams) WithIDNie(iDNie *string) *DcimRacksListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim racks list params
func (o *DcimRacksListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim racks list params
func (o *DcimRacksListParams) WithIDNiew(iDNiew *string) *DcimRacksListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim racks list params
func (o *DcimRacksListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim racks list params
func (o *DcimRacksListParams) WithIDNisw(iDNisw *string) *DcimRacksListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim racks list params
func (o *DcimRacksListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the dcim racks list params
func (o *DcimRacksListParams) WithLastUpdated(lastUpdated *string) *DcimRacksListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim racks list params
func (o *DcimRacksListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim racks list params
func (o *DcimRacksListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimRacksListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim racks list params
func (o *DcimRacksListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim racks list params
func (o *DcimRacksListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimRacksListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim racks list params
func (o *DcimRacksListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim racks list params
func (o *DcimRacksListParams) WithLimit(limit *int64) *DcimRacksListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim racks list params
func (o *DcimRacksListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the dcim racks list params
func (o *DcimRacksListParams) WithName(name *string) *DcimRacksListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim racks list params
func (o *DcimRacksListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim racks list params
func (o *DcimRacksListParams) WithNameIc(nameIc *string) *DcimRacksListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim racks list params
func (o *DcimRacksListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim racks list params
func (o *DcimRacksListParams) WithNameIe(nameIe *string) *DcimRacksListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim racks list params
func (o *DcimRacksListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim racks list params
func (o *DcimRacksListParams) WithNameIew(nameIew *string) *DcimRacksListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim racks list params
func (o *DcimRacksListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim racks list params
func (o *DcimRacksListParams) WithNameIsw(nameIsw *string) *DcimRacksListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim racks list params
func (o *DcimRacksListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim racks list params
func (o *DcimRacksListParams) WithNamen(namen *string) *DcimRacksListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim racks list params
func (o *DcimRacksListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim racks list params
func (o *DcimRacksListParams) WithNameNic(nameNic *string) *DcimRacksListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim racks list params
func (o *DcimRacksListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim racks list params
func (o *DcimRacksListParams) WithNameNie(nameNie *string) *DcimRacksListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim racks list params
func (o *DcimRacksListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim racks list params
func (o *DcimRacksListParams) WithNameNiew(nameNiew *string) *DcimRacksListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim racks list params
func (o *DcimRacksListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim racks list params
func (o *DcimRacksListParams) WithNameNisw(nameNisw *string) *DcimRacksListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim racks list params
func (o *DcimRacksListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim racks list params
func (o *DcimRacksListParams) WithOffset(offset *int64) *DcimRacksListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim racks list params
func (o *DcimRacksListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOuterDepth adds the outerDepth to the dcim racks list params
func (o *DcimRacksListParams) WithOuterDepth(outerDepth *string) *DcimRacksListParams {
	o.SetOuterDepth(outerDepth)
	return o
}

// SetOuterDepth adds the outerDepth to the dcim racks list params
func (o *DcimRacksListParams) SetOuterDepth(outerDepth *string) {
	o.OuterDepth = outerDepth
}

// WithOuterDepthGt adds the outerDepthGt to the dcim racks list params
func (o *DcimRacksListParams) WithOuterDepthGt(outerDepthGt *string) *DcimRacksListParams {
	o.SetOuterDepthGt(outerDepthGt)
	return o
}

// SetOuterDepthGt adds the outerDepthGt to the dcim racks list params
func (o *DcimRacksListParams) SetOuterDepthGt(outerDepthGt *string) {
	o.OuterDepthGt = outerDepthGt
}

// WithOuterDepthGte adds the outerDepthGte to the dcim racks list params
func (o *DcimRacksListParams) WithOuterDepthGte(outerDepthGte *string) *DcimRacksListParams {
	o.SetOuterDepthGte(outerDepthGte)
	return o
}

// SetOuterDepthGte adds the outerDepthGte to the dcim racks list params
func (o *DcimRacksListParams) SetOuterDepthGte(outerDepthGte *string) {
	o.OuterDepthGte = outerDepthGte
}

// WithOuterDepthLt adds the outerDepthLt to the dcim racks list params
func (o *DcimRacksListParams) WithOuterDepthLt(outerDepthLt *string) *DcimRacksListParams {
	o.SetOuterDepthLt(outerDepthLt)
	return o
}

// SetOuterDepthLt adds the outerDepthLt to the dcim racks list params
func (o *DcimRacksListParams) SetOuterDepthLt(outerDepthLt *string) {
	o.OuterDepthLt = outerDepthLt
}

// WithOuterDepthLte adds the outerDepthLte to the dcim racks list params
func (o *DcimRacksListParams) WithOuterDepthLte(outerDepthLte *string) *DcimRacksListParams {
	o.SetOuterDepthLte(outerDepthLte)
	return o
}

// SetOuterDepthLte adds the outerDepthLte to the dcim racks list params
func (o *DcimRacksListParams) SetOuterDepthLte(outerDepthLte *string) {
	o.OuterDepthLte = outerDepthLte
}

// WithOuterDepthn adds the outerDepthn to the dcim racks list params
func (o *DcimRacksListParams) WithOuterDepthn(outerDepthn *string) *DcimRacksListParams {
	o.SetOuterDepthn(outerDepthn)
	return o
}

// SetOuterDepthn adds the outerDepthN to the dcim racks list params
func (o *DcimRacksListParams) SetOuterDepthn(outerDepthn *string) {
	o.OuterDepthn = outerDepthn
}

// WithOuterUnit adds the outerUnit to the dcim racks list params
func (o *DcimRacksListParams) WithOuterUnit(outerUnit *string) *DcimRacksListParams {
	o.SetOuterUnit(outerUnit)
	return o
}

// SetOuterUnit adds the outerUnit to the dcim racks list params
func (o *DcimRacksListParams) SetOuterUnit(outerUnit *string) {
	o.OuterUnit = outerUnit
}

// WithOuterUnitn adds the outerUnitn to the dcim racks list params
func (o *DcimRacksListParams) WithOuterUnitn(outerUnitn *string) *DcimRacksListParams {
	o.SetOuterUnitn(outerUnitn)
	return o
}

// SetOuterUnitn adds the outerUnitN to the dcim racks list params
func (o *DcimRacksListParams) SetOuterUnitn(outerUnitn *string) {
	o.OuterUnitn = outerUnitn
}

// WithOuterWidth adds the outerWidth to the dcim racks list params
func (o *DcimRacksListParams) WithOuterWidth(outerWidth *string) *DcimRacksListParams {
	o.SetOuterWidth(outerWidth)
	return o
}

// SetOuterWidth adds the outerWidth to the dcim racks list params
func (o *DcimRacksListParams) SetOuterWidth(outerWidth *string) {
	o.OuterWidth = outerWidth
}

// WithOuterWidthGt adds the outerWidthGt to the dcim racks list params
func (o *DcimRacksListParams) WithOuterWidthGt(outerWidthGt *string) *DcimRacksListParams {
	o.SetOuterWidthGt(outerWidthGt)
	return o
}

// SetOuterWidthGt adds the outerWidthGt to the dcim racks list params
func (o *DcimRacksListParams) SetOuterWidthGt(outerWidthGt *string) {
	o.OuterWidthGt = outerWidthGt
}

// WithOuterWidthGte adds the outerWidthGte to the dcim racks list params
func (o *DcimRacksListParams) WithOuterWidthGte(outerWidthGte *string) *DcimRacksListParams {
	o.SetOuterWidthGte(outerWidthGte)
	return o
}

// SetOuterWidthGte adds the outerWidthGte to the dcim racks list params
func (o *DcimRacksListParams) SetOuterWidthGte(outerWidthGte *string) {
	o.OuterWidthGte = outerWidthGte
}

// WithOuterWidthLt adds the outerWidthLt to the dcim racks list params
func (o *DcimRacksListParams) WithOuterWidthLt(outerWidthLt *string) *DcimRacksListParams {
	o.SetOuterWidthLt(outerWidthLt)
	return o
}

// SetOuterWidthLt adds the outerWidthLt to the dcim racks list params
func (o *DcimRacksListParams) SetOuterWidthLt(outerWidthLt *string) {
	o.OuterWidthLt = outerWidthLt
}

// WithOuterWidthLte adds the outerWidthLte to the dcim racks list params
func (o *DcimRacksListParams) WithOuterWidthLte(outerWidthLte *string) *DcimRacksListParams {
	o.SetOuterWidthLte(outerWidthLte)
	return o
}

// SetOuterWidthLte adds the outerWidthLte to the dcim racks list params
func (o *DcimRacksListParams) SetOuterWidthLte(outerWidthLte *string) {
	o.OuterWidthLte = outerWidthLte
}

// WithOuterWidthn adds the outerWidthn to the dcim racks list params
func (o *DcimRacksListParams) WithOuterWidthn(outerWidthn *string) *DcimRacksListParams {
	o.SetOuterWidthn(outerWidthn)
	return o
}

// SetOuterWidthn adds the outerWidthN to the dcim racks list params
func (o *DcimRacksListParams) SetOuterWidthn(outerWidthn *string) {
	o.OuterWidthn = outerWidthn
}

// WithQ adds the q to the dcim racks list params
func (o *DcimRacksListParams) WithQ(q *string) *DcimRacksListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim racks list params
func (o *DcimRacksListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim racks list params
func (o *DcimRacksListParams) WithRegion(region *string) *DcimRacksListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim racks list params
func (o *DcimRacksListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim racks list params
func (o *DcimRacksListParams) WithRegionn(regionn *string) *DcimRacksListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim racks list params
func (o *DcimRacksListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim racks list params
func (o *DcimRacksListParams) WithRegionID(regionID *string) *DcimRacksListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim racks list params
func (o *DcimRacksListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim racks list params
func (o *DcimRacksListParams) WithRegionIDn(regionIDn *string) *DcimRacksListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim racks list params
func (o *DcimRacksListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithRole adds the role to the dcim racks list params
func (o *DcimRacksListParams) WithRole(role *string) *DcimRacksListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the dcim racks list params
func (o *DcimRacksListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the dcim racks list params
func (o *DcimRacksListParams) WithRolen(rolen *string) *DcimRacksListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the dcim racks list params
func (o *DcimRacksListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the dcim racks list params
func (o *DcimRacksListParams) WithRoleID(roleID *string) *DcimRacksListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the dcim racks list params
func (o *DcimRacksListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the dcim racks list params
func (o *DcimRacksListParams) WithRoleIDn(roleIDn *string) *DcimRacksListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the dcim racks list params
func (o *DcimRacksListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithSerial adds the serial to the dcim racks list params
func (o *DcimRacksListParams) WithSerial(serial *string) *DcimRacksListParams {
	o.SetSerial(serial)
	return o
}

// SetSerial adds the serial to the dcim racks list params
func (o *DcimRacksListParams) SetSerial(serial *string) {
	o.Serial = serial
}

// WithSite adds the site to the dcim racks list params
func (o *DcimRacksListParams) WithSite(site *string) *DcimRacksListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim racks list params
func (o *DcimRacksListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim racks list params
func (o *DcimRacksListParams) WithSiten(siten *string) *DcimRacksListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim racks list params
func (o *DcimRacksListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim racks list params
func (o *DcimRacksListParams) WithSiteID(siteID *string) *DcimRacksListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim racks list params
func (o *DcimRacksListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim racks list params
func (o *DcimRacksListParams) WithSiteIDn(siteIDn *string) *DcimRacksListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim racks list params
func (o *DcimRacksListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithStatus adds the status to the dcim racks list params
func (o *DcimRacksListParams) WithStatus(status *string) *DcimRacksListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim racks list params
func (o *DcimRacksListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the dcim racks list params
func (o *DcimRacksListParams) WithStatusn(statusn *string) *DcimRacksListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the dcim racks list params
func (o *DcimRacksListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the dcim racks list params
func (o *DcimRacksListParams) WithTag(tag *string) *DcimRacksListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim racks list params
func (o *DcimRacksListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim racks list params
func (o *DcimRacksListParams) WithTagn(tagn *string) *DcimRacksListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim racks list params
func (o *DcimRacksListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the dcim racks list params
func (o *DcimRacksListParams) WithTenant(tenant *string) *DcimRacksListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim racks list params
func (o *DcimRacksListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the dcim racks list params
func (o *DcimRacksListParams) WithTenantn(tenantn *string) *DcimRacksListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the dcim racks list params
func (o *DcimRacksListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the dcim racks list params
func (o *DcimRacksListParams) WithTenantGroup(tenantGroup *string) *DcimRacksListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the dcim racks list params
func (o *DcimRacksListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the dcim racks list params
func (o *DcimRacksListParams) WithTenantGroupn(tenantGroupn *string) *DcimRacksListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the dcim racks list params
func (o *DcimRacksListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the dcim racks list params
func (o *DcimRacksListParams) WithTenantGroupID(tenantGroupID *string) *DcimRacksListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the dcim racks list params
func (o *DcimRacksListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the dcim racks list params
func (o *DcimRacksListParams) WithTenantGroupIDn(tenantGroupIDn *string) *DcimRacksListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the dcim racks list params
func (o *DcimRacksListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the dcim racks list params
func (o *DcimRacksListParams) WithTenantID(tenantID *string) *DcimRacksListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim racks list params
func (o *DcimRacksListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the dcim racks list params
func (o *DcimRacksListParams) WithTenantIDn(tenantIDn *string) *DcimRacksListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the dcim racks list params
func (o *DcimRacksListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithType adds the typeVar to the dcim racks list params
func (o *DcimRacksListParams) WithType(typeVar *string) *DcimRacksListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim racks list params
func (o *DcimRacksListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim racks list params
func (o *DcimRacksListParams) WithTypen(typen *string) *DcimRacksListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim racks list params
func (o *DcimRacksListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WithUHeight adds the uHeight to the dcim racks list params
func (o *DcimRacksListParams) WithUHeight(uHeight *string) *DcimRacksListParams {
	o.SetUHeight(uHeight)
	return o
}

// SetUHeight adds the uHeight to the dcim racks list params
func (o *DcimRacksListParams) SetUHeight(uHeight *string) {
	o.UHeight = uHeight
}

// WithUHeightGt adds the uHeightGt to the dcim racks list params
func (o *DcimRacksListParams) WithUHeightGt(uHeightGt *string) *DcimRacksListParams {
	o.SetUHeightGt(uHeightGt)
	return o
}

// SetUHeightGt adds the uHeightGt to the dcim racks list params
func (o *DcimRacksListParams) SetUHeightGt(uHeightGt *string) {
	o.UHeightGt = uHeightGt
}

// WithUHeightGte adds the uHeightGte to the dcim racks list params
func (o *DcimRacksListParams) WithUHeightGte(uHeightGte *string) *DcimRacksListParams {
	o.SetUHeightGte(uHeightGte)
	return o
}

// SetUHeightGte adds the uHeightGte to the dcim racks list params
func (o *DcimRacksListParams) SetUHeightGte(uHeightGte *string) {
	o.UHeightGte = uHeightGte
}

// WithUHeightLt adds the uHeightLt to the dcim racks list params
func (o *DcimRacksListParams) WithUHeightLt(uHeightLt *string) *DcimRacksListParams {
	o.SetUHeightLt(uHeightLt)
	return o
}

// SetUHeightLt adds the uHeightLt to the dcim racks list params
func (o *DcimRacksListParams) SetUHeightLt(uHeightLt *string) {
	o.UHeightLt = uHeightLt
}

// WithUHeightLte adds the uHeightLte to the dcim racks list params
func (o *DcimRacksListParams) WithUHeightLte(uHeightLte *string) *DcimRacksListParams {
	o.SetUHeightLte(uHeightLte)
	return o
}

// SetUHeightLte adds the uHeightLte to the dcim racks list params
func (o *DcimRacksListParams) SetUHeightLte(uHeightLte *string) {
	o.UHeightLte = uHeightLte
}

// WithUHeightn adds the uHeightn to the dcim racks list params
func (o *DcimRacksListParams) WithUHeightn(uHeightn *string) *DcimRacksListParams {
	o.SetUHeightn(uHeightn)
	return o
}

// SetUHeightn adds the uHeightN to the dcim racks list params
func (o *DcimRacksListParams) SetUHeightn(uHeightn *string) {
	o.UHeightn = uHeightn
}

// WithWidth adds the width to the dcim racks list params
func (o *DcimRacksListParams) WithWidth(width *string) *DcimRacksListParams {
	o.SetWidth(width)
	return o
}

// SetWidth adds the width to the dcim racks list params
func (o *DcimRacksListParams) SetWidth(width *string) {
	o.Width = width
}

// WithWidthn adds the widthn to the dcim racks list params
func (o *DcimRacksListParams) WithWidthn(widthn *string) *DcimRacksListParams {
	o.SetWidthn(widthn)
	return o
}

// SetWidthn adds the widthN to the dcim racks list params
func (o *DcimRacksListParams) SetWidthn(widthn *string) {
	o.Widthn = widthn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRacksListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DescUnits != nil {

		// query param desc_units
		var qrDescUnits string

		if o.DescUnits != nil {
			qrDescUnits = *o.DescUnits
		}
		qDescUnits := qrDescUnits
		if qDescUnits != "" {

			if err := r.SetQueryParam("desc_units", qDescUnits); err != nil {
				return err
			}
		}
	}

	if o.FacilityID != nil {

		// query param facility_id
		var qrFacilityID string

		if o.FacilityID != nil {
			qrFacilityID = *o.FacilityID
		}
		qFacilityID := qrFacilityID
		if qFacilityID != "" {

			if err := r.SetQueryParam("facility_id", qFacilityID); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDIc != nil {

		// query param facility_id__ic
		var qrFacilityIDIc string

		if o.FacilityIDIc != nil {
			qrFacilityIDIc = *o.FacilityIDIc
		}
		qFacilityIDIc := qrFacilityIDIc
		if qFacilityIDIc != "" {

			if err := r.SetQueryParam("facility_id__ic", qFacilityIDIc); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDIe != nil {

		// query param facility_id__ie
		var qrFacilityIDIe string

		if o.FacilityIDIe != nil {
			qrFacilityIDIe = *o.FacilityIDIe
		}
		qFacilityIDIe := qrFacilityIDIe
		if qFacilityIDIe != "" {

			if err := r.SetQueryParam("facility_id__ie", qFacilityIDIe); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDIew != nil {

		// query param facility_id__iew
		var qrFacilityIDIew string

		if o.FacilityIDIew != nil {
			qrFacilityIDIew = *o.FacilityIDIew
		}
		qFacilityIDIew := qrFacilityIDIew
		if qFacilityIDIew != "" {

			if err := r.SetQueryParam("facility_id__iew", qFacilityIDIew); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDIsw != nil {

		// query param facility_id__isw
		var qrFacilityIDIsw string

		if o.FacilityIDIsw != nil {
			qrFacilityIDIsw = *o.FacilityIDIsw
		}
		qFacilityIDIsw := qrFacilityIDIsw
		if qFacilityIDIsw != "" {

			if err := r.SetQueryParam("facility_id__isw", qFacilityIDIsw); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDn != nil {

		// query param facility_id__n
		var qrFacilityIDn string

		if o.FacilityIDn != nil {
			qrFacilityIDn = *o.FacilityIDn
		}
		qFacilityIDn := qrFacilityIDn
		if qFacilityIDn != "" {

			if err := r.SetQueryParam("facility_id__n", qFacilityIDn); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDNic != nil {

		// query param facility_id__nic
		var qrFacilityIDNic string

		if o.FacilityIDNic != nil {
			qrFacilityIDNic = *o.FacilityIDNic
		}
		qFacilityIDNic := qrFacilityIDNic
		if qFacilityIDNic != "" {

			if err := r.SetQueryParam("facility_id__nic", qFacilityIDNic); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDNie != nil {

		// query param facility_id__nie
		var qrFacilityIDNie string

		if o.FacilityIDNie != nil {
			qrFacilityIDNie = *o.FacilityIDNie
		}
		qFacilityIDNie := qrFacilityIDNie
		if qFacilityIDNie != "" {

			if err := r.SetQueryParam("facility_id__nie", qFacilityIDNie); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDNiew != nil {

		// query param facility_id__niew
		var qrFacilityIDNiew string

		if o.FacilityIDNiew != nil {
			qrFacilityIDNiew = *o.FacilityIDNiew
		}
		qFacilityIDNiew := qrFacilityIDNiew
		if qFacilityIDNiew != "" {

			if err := r.SetQueryParam("facility_id__niew", qFacilityIDNiew); err != nil {
				return err
			}
		}
	}

	if o.FacilityIDNisw != nil {

		// query param facility_id__nisw
		var qrFacilityIDNisw string

		if o.FacilityIDNisw != nil {
			qrFacilityIDNisw = *o.FacilityIDNisw
		}
		qFacilityIDNisw := qrFacilityIDNisw
		if qFacilityIDNisw != "" {

			if err := r.SetQueryParam("facility_id__nisw", qFacilityIDNisw); err != nil {
				return err
			}
		}
	}

	if o.Group != nil {

		// query param group
		var qrGroup string

		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := qrGroup
		if qGroup != "" {

			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}
	}

	if o.Groupn != nil {

		// query param group__n
		var qrGroupn string

		if o.Groupn != nil {
			qrGroupn = *o.Groupn
		}
		qGroupn := qrGroupn
		if qGroupn != "" {

			if err := r.SetQueryParam("group__n", qGroupn); err != nil {
				return err
			}
		}
	}

	if o.GroupID != nil {

		// query param group_id
		var qrGroupID string

		if o.GroupID != nil {
			qrGroupID = *o.GroupID
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("group_id", qGroupID); err != nil {
				return err
			}
		}
	}

	if o.GroupIDn != nil {

		// query param group_id__n
		var qrGroupIDn string

		if o.GroupIDn != nil {
			qrGroupIDn = *o.GroupIDn
		}
		qGroupIDn := qrGroupIDn
		if qGroupIDn != "" {

			if err := r.SetQueryParam("group_id__n", qGroupIDn); err != nil {
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

	if o.OuterDepth != nil {

		// query param outer_depth
		var qrOuterDepth string

		if o.OuterDepth != nil {
			qrOuterDepth = *o.OuterDepth
		}
		qOuterDepth := qrOuterDepth
		if qOuterDepth != "" {

			if err := r.SetQueryParam("outer_depth", qOuterDepth); err != nil {
				return err
			}
		}
	}

	if o.OuterDepthGt != nil {

		// query param outer_depth__gt
		var qrOuterDepthGt string

		if o.OuterDepthGt != nil {
			qrOuterDepthGt = *o.OuterDepthGt
		}
		qOuterDepthGt := qrOuterDepthGt
		if qOuterDepthGt != "" {

			if err := r.SetQueryParam("outer_depth__gt", qOuterDepthGt); err != nil {
				return err
			}
		}
	}

	if o.OuterDepthGte != nil {

		// query param outer_depth__gte
		var qrOuterDepthGte string

		if o.OuterDepthGte != nil {
			qrOuterDepthGte = *o.OuterDepthGte
		}
		qOuterDepthGte := qrOuterDepthGte
		if qOuterDepthGte != "" {

			if err := r.SetQueryParam("outer_depth__gte", qOuterDepthGte); err != nil {
				return err
			}
		}
	}

	if o.OuterDepthLt != nil {

		// query param outer_depth__lt
		var qrOuterDepthLt string

		if o.OuterDepthLt != nil {
			qrOuterDepthLt = *o.OuterDepthLt
		}
		qOuterDepthLt := qrOuterDepthLt
		if qOuterDepthLt != "" {

			if err := r.SetQueryParam("outer_depth__lt", qOuterDepthLt); err != nil {
				return err
			}
		}
	}

	if o.OuterDepthLte != nil {

		// query param outer_depth__lte
		var qrOuterDepthLte string

		if o.OuterDepthLte != nil {
			qrOuterDepthLte = *o.OuterDepthLte
		}
		qOuterDepthLte := qrOuterDepthLte
		if qOuterDepthLte != "" {

			if err := r.SetQueryParam("outer_depth__lte", qOuterDepthLte); err != nil {
				return err
			}
		}
	}

	if o.OuterDepthn != nil {

		// query param outer_depth__n
		var qrOuterDepthn string

		if o.OuterDepthn != nil {
			qrOuterDepthn = *o.OuterDepthn
		}
		qOuterDepthn := qrOuterDepthn
		if qOuterDepthn != "" {

			if err := r.SetQueryParam("outer_depth__n", qOuterDepthn); err != nil {
				return err
			}
		}
	}

	if o.OuterUnit != nil {

		// query param outer_unit
		var qrOuterUnit string

		if o.OuterUnit != nil {
			qrOuterUnit = *o.OuterUnit
		}
		qOuterUnit := qrOuterUnit
		if qOuterUnit != "" {

			if err := r.SetQueryParam("outer_unit", qOuterUnit); err != nil {
				return err
			}
		}
	}

	if o.OuterUnitn != nil {

		// query param outer_unit__n
		var qrOuterUnitn string

		if o.OuterUnitn != nil {
			qrOuterUnitn = *o.OuterUnitn
		}
		qOuterUnitn := qrOuterUnitn
		if qOuterUnitn != "" {

			if err := r.SetQueryParam("outer_unit__n", qOuterUnitn); err != nil {
				return err
			}
		}
	}

	if o.OuterWidth != nil {

		// query param outer_width
		var qrOuterWidth string

		if o.OuterWidth != nil {
			qrOuterWidth = *o.OuterWidth
		}
		qOuterWidth := qrOuterWidth
		if qOuterWidth != "" {

			if err := r.SetQueryParam("outer_width", qOuterWidth); err != nil {
				return err
			}
		}
	}

	if o.OuterWidthGt != nil {

		// query param outer_width__gt
		var qrOuterWidthGt string

		if o.OuterWidthGt != nil {
			qrOuterWidthGt = *o.OuterWidthGt
		}
		qOuterWidthGt := qrOuterWidthGt
		if qOuterWidthGt != "" {

			if err := r.SetQueryParam("outer_width__gt", qOuterWidthGt); err != nil {
				return err
			}
		}
	}

	if o.OuterWidthGte != nil {

		// query param outer_width__gte
		var qrOuterWidthGte string

		if o.OuterWidthGte != nil {
			qrOuterWidthGte = *o.OuterWidthGte
		}
		qOuterWidthGte := qrOuterWidthGte
		if qOuterWidthGte != "" {

			if err := r.SetQueryParam("outer_width__gte", qOuterWidthGte); err != nil {
				return err
			}
		}
	}

	if o.OuterWidthLt != nil {

		// query param outer_width__lt
		var qrOuterWidthLt string

		if o.OuterWidthLt != nil {
			qrOuterWidthLt = *o.OuterWidthLt
		}
		qOuterWidthLt := qrOuterWidthLt
		if qOuterWidthLt != "" {

			if err := r.SetQueryParam("outer_width__lt", qOuterWidthLt); err != nil {
				return err
			}
		}
	}

	if o.OuterWidthLte != nil {

		// query param outer_width__lte
		var qrOuterWidthLte string

		if o.OuterWidthLte != nil {
			qrOuterWidthLte = *o.OuterWidthLte
		}
		qOuterWidthLte := qrOuterWidthLte
		if qOuterWidthLte != "" {

			if err := r.SetQueryParam("outer_width__lte", qOuterWidthLte); err != nil {
				return err
			}
		}
	}

	if o.OuterWidthn != nil {

		// query param outer_width__n
		var qrOuterWidthn string

		if o.OuterWidthn != nil {
			qrOuterWidthn = *o.OuterWidthn
		}
		qOuterWidthn := qrOuterWidthn
		if qOuterWidthn != "" {

			if err := r.SetQueryParam("outer_width__n", qOuterWidthn); err != nil {
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

	if o.UHeight != nil {

		// query param u_height
		var qrUHeight string

		if o.UHeight != nil {
			qrUHeight = *o.UHeight
		}
		qUHeight := qrUHeight
		if qUHeight != "" {

			if err := r.SetQueryParam("u_height", qUHeight); err != nil {
				return err
			}
		}
	}

	if o.UHeightGt != nil {

		// query param u_height__gt
		var qrUHeightGt string

		if o.UHeightGt != nil {
			qrUHeightGt = *o.UHeightGt
		}
		qUHeightGt := qrUHeightGt
		if qUHeightGt != "" {

			if err := r.SetQueryParam("u_height__gt", qUHeightGt); err != nil {
				return err
			}
		}
	}

	if o.UHeightGte != nil {

		// query param u_height__gte
		var qrUHeightGte string

		if o.UHeightGte != nil {
			qrUHeightGte = *o.UHeightGte
		}
		qUHeightGte := qrUHeightGte
		if qUHeightGte != "" {

			if err := r.SetQueryParam("u_height__gte", qUHeightGte); err != nil {
				return err
			}
		}
	}

	if o.UHeightLt != nil {

		// query param u_height__lt
		var qrUHeightLt string

		if o.UHeightLt != nil {
			qrUHeightLt = *o.UHeightLt
		}
		qUHeightLt := qrUHeightLt
		if qUHeightLt != "" {

			if err := r.SetQueryParam("u_height__lt", qUHeightLt); err != nil {
				return err
			}
		}
	}

	if o.UHeightLte != nil {

		// query param u_height__lte
		var qrUHeightLte string

		if o.UHeightLte != nil {
			qrUHeightLte = *o.UHeightLte
		}
		qUHeightLte := qrUHeightLte
		if qUHeightLte != "" {

			if err := r.SetQueryParam("u_height__lte", qUHeightLte); err != nil {
				return err
			}
		}
	}

	if o.UHeightn != nil {

		// query param u_height__n
		var qrUHeightn string

		if o.UHeightn != nil {
			qrUHeightn = *o.UHeightn
		}
		qUHeightn := qrUHeightn
		if qUHeightn != "" {

			if err := r.SetQueryParam("u_height__n", qUHeightn); err != nil {
				return err
			}
		}
	}

	if o.Width != nil {

		// query param width
		var qrWidth string

		if o.Width != nil {
			qrWidth = *o.Width
		}
		qWidth := qrWidth
		if qWidth != "" {

			if err := r.SetQueryParam("width", qWidth); err != nil {
				return err
			}
		}
	}

	if o.Widthn != nil {

		// query param width__n
		var qrWidthn string

		if o.Widthn != nil {
			qrWidthn = *o.Widthn
		}
		qWidthn := qrWidthn
		if qWidthn != "" {

			if err := r.SetQueryParam("width__n", qWidthn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
