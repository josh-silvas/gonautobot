package ipam

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

// NewIpamPrefixesListParams creates a new IpamPrefixesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamPrefixesListParams() *IpamPrefixesListParams {
	return &IpamPrefixesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamPrefixesListParamsWithTimeout creates a new IpamPrefixesListParams object
// with the ability to set a timeout on a request.
func NewIpamPrefixesListParamsWithTimeout(timeout time.Duration) *IpamPrefixesListParams {
	return &IpamPrefixesListParams{
		timeout: timeout,
	}
}

// NewIpamPrefixesListParamsWithContext creates a new IpamPrefixesListParams object
// with the ability to set a context for a request.
func NewIpamPrefixesListParamsWithContext(ctx context.Context) *IpamPrefixesListParams {
	return &IpamPrefixesListParams{
		Context: ctx,
	}
}

// NewIpamPrefixesListParamsWithHTTPClient creates a new IpamPrefixesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamPrefixesListParamsWithHTTPClient(client *http.Client) *IpamPrefixesListParams {
	return &IpamPrefixesListParams{
		HTTPClient: client,
	}
}

/* IpamPrefixesListParams contains all the parameters to send to the API endpoint
   for the ipam prefixes list operation.

   Typically these are written to a http.Request.
*/
type IpamPrefixesListParams struct {

	// Contains.
	Contains *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Family.
	Family *float64

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

	// IsPool.
	IsPool *string

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

	// MaskLength.
	MaskLength *float64

	// MaskLengthGte.
	MaskLengthGte *float64

	// MaskLengthLte.
	MaskLengthLte *float64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Prefix.
	Prefix *string

	// PresentInVrf.
	PresentInVrf *string

	// PresentInVrfID.
	PresentInVrfID *string

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

	// VlanID.
	VlanID *string

	// VlanIDn.
	VlanIDn *string

	// VlanVid.
	VlanVid *float64

	// Vrf.
	Vrf *string

	// Vrfn.
	Vrfn *string

	// VrfID.
	VrfID *string

	// VrfIDn.
	VrfIDn *string

	// Within.
	Within *string

	// WithinInclude.
	WithinInclude *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam prefixes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesListParams) WithDefaults() *IpamPrefixesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam prefixes list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamPrefixesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTimeout(timeout time.Duration) *IpamPrefixesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithContext(ctx context.Context) *IpamPrefixesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithHTTPClient(client *http.Client) *IpamPrefixesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContains adds the contains to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithContains(contains *string) *IpamPrefixesListParams {
	o.SetContains(contains)
	return o
}

// SetContains adds the contains to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetContains(contains *string) {
	o.Contains = contains
}

// WithCreated adds the created to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithCreated(created *string) *IpamPrefixesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithCreatedGte(createdGte *string) *IpamPrefixesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithCreatedLte(createdLte *string) *IpamPrefixesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithFamily adds the family to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithFamily(family *float64) *IpamPrefixesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetFamily(family *float64) {
	o.Family = family
}

// WithID adds the id to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithID(id *string) *IpamPrefixesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDIc(iDIc *string) *IpamPrefixesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDIe(iDIe *string) *IpamPrefixesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDIew(iDIew *string) *IpamPrefixesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDIsw(iDIsw *string) *IpamPrefixesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDn(iDn *string) *IpamPrefixesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDNic(iDNic *string) *IpamPrefixesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDNie(iDNie *string) *IpamPrefixesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDNiew(iDNiew *string) *IpamPrefixesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIDNisw(iDNisw *string) *IpamPrefixesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithIsPool adds the isPool to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithIsPool(isPool *string) *IpamPrefixesListParams {
	o.SetIsPool(isPool)
	return o
}

// SetIsPool adds the isPool to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetIsPool(isPool *string) {
	o.IsPool = isPool
}

// WithLastUpdated adds the lastUpdated to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithLastUpdated(lastUpdated *string) *IpamPrefixesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamPrefixesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamPrefixesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithLimit(limit *int64) *IpamPrefixesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaskLength adds the maskLength to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithMaskLength(maskLength *float64) *IpamPrefixesListParams {
	o.SetMaskLength(maskLength)
	return o
}

// SetMaskLength adds the maskLength to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetMaskLength(maskLength *float64) {
	o.MaskLength = maskLength
}

// WithMaskLengthGte adds the maskLengthGte to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithMaskLengthGte(maskLengthGte *float64) *IpamPrefixesListParams {
	o.SetMaskLengthGte(maskLengthGte)
	return o
}

// SetMaskLengthGte adds the maskLengthGte to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetMaskLengthGte(maskLengthGte *float64) {
	o.MaskLengthGte = maskLengthGte
}

// WithMaskLengthLte adds the maskLengthLte to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithMaskLengthLte(maskLengthLte *float64) *IpamPrefixesListParams {
	o.SetMaskLengthLte(maskLengthLte)
	return o
}

// SetMaskLengthLte adds the maskLengthLte to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetMaskLengthLte(maskLengthLte *float64) {
	o.MaskLengthLte = maskLengthLte
}

// WithOffset adds the offset to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithOffset(offset *int64) *IpamPrefixesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPrefix adds the prefix to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithPrefix(prefix *string) *IpamPrefixesListParams {
	o.SetPrefix(prefix)
	return o
}

// SetPrefix adds the prefix to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetPrefix(prefix *string) {
	o.Prefix = prefix
}

// WithPresentInVrf adds the presentInVrf to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithPresentInVrf(presentInVrf *string) *IpamPrefixesListParams {
	o.SetPresentInVrf(presentInVrf)
	return o
}

// SetPresentInVrf adds the presentInVrf to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetPresentInVrf(presentInVrf *string) {
	o.PresentInVrf = presentInVrf
}

// WithPresentInVrfID adds the presentInVrfID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithPresentInVrfID(presentInVrfID *string) *IpamPrefixesListParams {
	o.SetPresentInVrfID(presentInVrfID)
	return o
}

// SetPresentInVrfID adds the presentInVrfId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetPresentInVrfID(presentInVrfID *string) {
	o.PresentInVrfID = presentInVrfID
}

// WithQ adds the q to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithQ(q *string) *IpamPrefixesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRegion(region *string) *IpamPrefixesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRegionn(regionn *string) *IpamPrefixesListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRegionID(regionID *string) *IpamPrefixesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRegionIDn(regionIDn *string) *IpamPrefixesListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithRole adds the role to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRole(role *string) *IpamPrefixesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRolen(rolen *string) *IpamPrefixesListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRoleID(roleID *string) *IpamPrefixesListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithRoleIDn(roleIDn *string) *IpamPrefixesListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithSite adds the site to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithSite(site *string) *IpamPrefixesListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithSiten(siten *string) *IpamPrefixesListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithSiteID(siteID *string) *IpamPrefixesListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithSiteIDn(siteIDn *string) *IpamPrefixesListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithStatus adds the status to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithStatus(status *string) *IpamPrefixesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithStatusn(statusn *string) *IpamPrefixesListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTag(tag *string) *IpamPrefixesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTagn(tagn *string) *IpamPrefixesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenant(tenant *string) *IpamPrefixesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenantn(tenantn *string) *IpamPrefixesListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenantGroup(tenantGroup *string) *IpamPrefixesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenantGroupn(tenantGroupn *string) *IpamPrefixesListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenantGroupID(tenantGroupID *string) *IpamPrefixesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenantGroupIDn(tenantGroupIDn *string) *IpamPrefixesListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenantID(tenantID *string) *IpamPrefixesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithTenantIDn(tenantIDn *string) *IpamPrefixesListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithVlanID adds the vlanID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithVlanID(vlanID *string) *IpamPrefixesListParams {
	o.SetVlanID(vlanID)
	return o
}

// SetVlanID adds the vlanId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetVlanID(vlanID *string) {
	o.VlanID = vlanID
}

// WithVlanIDn adds the vlanIDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithVlanIDn(vlanIDn *string) *IpamPrefixesListParams {
	o.SetVlanIDn(vlanIDn)
	return o
}

// SetVlanIDn adds the vlanIdN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetVlanIDn(vlanIDn *string) {
	o.VlanIDn = vlanIDn
}

// WithVlanVid adds the vlanVid to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithVlanVid(vlanVid *float64) *IpamPrefixesListParams {
	o.SetVlanVid(vlanVid)
	return o
}

// SetVlanVid adds the vlanVid to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetVlanVid(vlanVid *float64) {
	o.VlanVid = vlanVid
}

// WithVrf adds the vrf to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithVrf(vrf *string) *IpamPrefixesListParams {
	o.SetVrf(vrf)
	return o
}

// SetVrf adds the vrf to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetVrf(vrf *string) {
	o.Vrf = vrf
}

// WithVrfn adds the vrfn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithVrfn(vrfn *string) *IpamPrefixesListParams {
	o.SetVrfn(vrfn)
	return o
}

// SetVrfn adds the vrfN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetVrfn(vrfn *string) {
	o.Vrfn = vrfn
}

// WithVrfID adds the vrfID to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithVrfID(vrfID *string) *IpamPrefixesListParams {
	o.SetVrfID(vrfID)
	return o
}

// SetVrfID adds the vrfId to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetVrfID(vrfID *string) {
	o.VrfID = vrfID
}

// WithVrfIDn adds the vrfIDn to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithVrfIDn(vrfIDn *string) *IpamPrefixesListParams {
	o.SetVrfIDn(vrfIDn)
	return o
}

// SetVrfIDn adds the vrfIdN to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetVrfIDn(vrfIDn *string) {
	o.VrfIDn = vrfIDn
}

// WithWithin adds the within to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithWithin(within *string) *IpamPrefixesListParams {
	o.SetWithin(within)
	return o
}

// SetWithin adds the within to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetWithin(within *string) {
	o.Within = within
}

// WithWithinInclude adds the withinInclude to the ipam prefixes list params
func (o *IpamPrefixesListParams) WithWithinInclude(withinInclude *string) *IpamPrefixesListParams {
	o.SetWithinInclude(withinInclude)
	return o
}

// SetWithinInclude adds the withinInclude to the ipam prefixes list params
func (o *IpamPrefixesListParams) SetWithinInclude(withinInclude *string) {
	o.WithinInclude = withinInclude
}

// WriteToRequest writes these params to a swagger request
func (o *IpamPrefixesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Contains != nil {

		// query param contains
		var qrContains string

		if o.Contains != nil {
			qrContains = *o.Contains
		}
		qContains := qrContains
		if qContains != "" {

			if err := r.SetQueryParam("contains", qContains); err != nil {
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

	if o.Family != nil {

		// query param family
		var qrFamily float64

		if o.Family != nil {
			qrFamily = *o.Family
		}
		qFamily := swag.FormatFloat64(qrFamily)
		if qFamily != "" {

			if err := r.SetQueryParam("family", qFamily); err != nil {
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

	if o.IsPool != nil {

		// query param is_pool
		var qrIsPool string

		if o.IsPool != nil {
			qrIsPool = *o.IsPool
		}
		qIsPool := qrIsPool
		if qIsPool != "" {

			if err := r.SetQueryParam("is_pool", qIsPool); err != nil {
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

	if o.MaskLength != nil {

		// query param mask_length
		var qrMaskLength float64

		if o.MaskLength != nil {
			qrMaskLength = *o.MaskLength
		}
		qMaskLength := swag.FormatFloat64(qrMaskLength)
		if qMaskLength != "" {

			if err := r.SetQueryParam("mask_length", qMaskLength); err != nil {
				return err
			}
		}
	}

	if o.MaskLengthGte != nil {

		// query param mask_length__gte
		var qrMaskLengthGte float64

		if o.MaskLengthGte != nil {
			qrMaskLengthGte = *o.MaskLengthGte
		}
		qMaskLengthGte := swag.FormatFloat64(qrMaskLengthGte)
		if qMaskLengthGte != "" {

			if err := r.SetQueryParam("mask_length__gte", qMaskLengthGte); err != nil {
				return err
			}
		}
	}

	if o.MaskLengthLte != nil {

		// query param mask_length__lte
		var qrMaskLengthLte float64

		if o.MaskLengthLte != nil {
			qrMaskLengthLte = *o.MaskLengthLte
		}
		qMaskLengthLte := swag.FormatFloat64(qrMaskLengthLte)
		if qMaskLengthLte != "" {

			if err := r.SetQueryParam("mask_length__lte", qMaskLengthLte); err != nil {
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

	if o.Prefix != nil {

		// query param prefix
		var qrPrefix string

		if o.Prefix != nil {
			qrPrefix = *o.Prefix
		}
		qPrefix := qrPrefix
		if qPrefix != "" {

			if err := r.SetQueryParam("prefix", qPrefix); err != nil {
				return err
			}
		}
	}

	if o.PresentInVrf != nil {

		// query param present_in_vrf
		var qrPresentInVrf string

		if o.PresentInVrf != nil {
			qrPresentInVrf = *o.PresentInVrf
		}
		qPresentInVrf := qrPresentInVrf
		if qPresentInVrf != "" {

			if err := r.SetQueryParam("present_in_vrf", qPresentInVrf); err != nil {
				return err
			}
		}
	}

	if o.PresentInVrfID != nil {

		// query param present_in_vrf_id
		var qrPresentInVrfID string

		if o.PresentInVrfID != nil {
			qrPresentInVrfID = *o.PresentInVrfID
		}
		qPresentInVrfID := qrPresentInVrfID
		if qPresentInVrfID != "" {

			if err := r.SetQueryParam("present_in_vrf_id", qPresentInVrfID); err != nil {
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

	if o.VlanIDn != nil {

		// query param vlan_id__n
		var qrVlanIDn string

		if o.VlanIDn != nil {
			qrVlanIDn = *o.VlanIDn
		}
		qVlanIDn := qrVlanIDn
		if qVlanIDn != "" {

			if err := r.SetQueryParam("vlan_id__n", qVlanIDn); err != nil {
				return err
			}
		}
	}

	if o.VlanVid != nil {

		// query param vlan_vid
		var qrVlanVid float64

		if o.VlanVid != nil {
			qrVlanVid = *o.VlanVid
		}
		qVlanVid := swag.FormatFloat64(qrVlanVid)
		if qVlanVid != "" {

			if err := r.SetQueryParam("vlan_vid", qVlanVid); err != nil {
				return err
			}
		}
	}

	if o.Vrf != nil {

		// query param vrf
		var qrVrf string

		if o.Vrf != nil {
			qrVrf = *o.Vrf
		}
		qVrf := qrVrf
		if qVrf != "" {

			if err := r.SetQueryParam("vrf", qVrf); err != nil {
				return err
			}
		}
	}

	if o.Vrfn != nil {

		// query param vrf__n
		var qrVrfn string

		if o.Vrfn != nil {
			qrVrfn = *o.Vrfn
		}
		qVrfn := qrVrfn
		if qVrfn != "" {

			if err := r.SetQueryParam("vrf__n", qVrfn); err != nil {
				return err
			}
		}
	}

	if o.VrfID != nil {

		// query param vrf_id
		var qrVrfID string

		if o.VrfID != nil {
			qrVrfID = *o.VrfID
		}
		qVrfID := qrVrfID
		if qVrfID != "" {

			if err := r.SetQueryParam("vrf_id", qVrfID); err != nil {
				return err
			}
		}
	}

	if o.VrfIDn != nil {

		// query param vrf_id__n
		var qrVrfIDn string

		if o.VrfIDn != nil {
			qrVrfIDn = *o.VrfIDn
		}
		qVrfIDn := qrVrfIDn
		if qVrfIDn != "" {

			if err := r.SetQueryParam("vrf_id__n", qVrfIDn); err != nil {
				return err
			}
		}
	}

	if o.Within != nil {

		// query param within
		var qrWithin string

		if o.Within != nil {
			qrWithin = *o.Within
		}
		qWithin := qrWithin
		if qWithin != "" {

			if err := r.SetQueryParam("within", qWithin); err != nil {
				return err
			}
		}
	}

	if o.WithinInclude != nil {

		// query param within_include
		var qrWithinInclude string

		if o.WithinInclude != nil {
			qrWithinInclude = *o.WithinInclude
		}
		qWithinInclude := qrWithinInclude
		if qWithinInclude != "" {

			if err := r.SetQueryParam("within_include", qWithinInclude); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
