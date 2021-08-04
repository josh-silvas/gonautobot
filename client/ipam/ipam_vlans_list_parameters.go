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

// NewIpamVlansListParams creates a new IpamVlansListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVlansListParams() *IpamVlansListParams {
	return &IpamVlansListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVlansListParamsWithTimeout creates a new IpamVlansListParams object
// with the ability to set a timeout on a request.
func NewIpamVlansListParamsWithTimeout(timeout time.Duration) *IpamVlansListParams {
	return &IpamVlansListParams{
		timeout: timeout,
	}
}

// NewIpamVlansListParamsWithContext creates a new IpamVlansListParams object
// with the ability to set a context for a request.
func NewIpamVlansListParamsWithContext(ctx context.Context) *IpamVlansListParams {
	return &IpamVlansListParams{
		Context: ctx,
	}
}

// NewIpamVlansListParamsWithHTTPClient creates a new IpamVlansListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVlansListParamsWithHTTPClient(client *http.Client) *IpamVlansListParams {
	return &IpamVlansListParams{
		HTTPClient: client,
	}
}

/* IpamVlansListParams contains all the parameters to send to the API endpoint
   for the ipam vlans list operation.

   Typically these are written to a http.Request.
*/
type IpamVlansListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

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

	// Vid.
	Vid *string

	// VidGt.
	VidGt *string

	// VidGte.
	VidGte *string

	// VidLt.
	VidLt *string

	// VidLte.
	VidLte *string

	// Vidn.
	Vidn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam vlans list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansListParams) WithDefaults() *IpamVlansListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vlans list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVlansListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vlans list params
func (o *IpamVlansListParams) WithTimeout(timeout time.Duration) *IpamVlansListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans list params
func (o *IpamVlansListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans list params
func (o *IpamVlansListParams) WithContext(ctx context.Context) *IpamVlansListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans list params
func (o *IpamVlansListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans list params
func (o *IpamVlansListParams) WithHTTPClient(client *http.Client) *IpamVlansListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans list params
func (o *IpamVlansListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam vlans list params
func (o *IpamVlansListParams) WithCreated(created *string) *IpamVlansListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam vlans list params
func (o *IpamVlansListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam vlans list params
func (o *IpamVlansListParams) WithCreatedGte(createdGte *string) *IpamVlansListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam vlans list params
func (o *IpamVlansListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam vlans list params
func (o *IpamVlansListParams) WithCreatedLte(createdLte *string) *IpamVlansListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam vlans list params
func (o *IpamVlansListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithGroup adds the group to the ipam vlans list params
func (o *IpamVlansListParams) WithGroup(group *string) *IpamVlansListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the ipam vlans list params
func (o *IpamVlansListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupn adds the groupn to the ipam vlans list params
func (o *IpamVlansListParams) WithGroupn(groupn *string) *IpamVlansListParams {
	o.SetGroupn(groupn)
	return o
}

// SetGroupn adds the groupN to the ipam vlans list params
func (o *IpamVlansListParams) SetGroupn(groupn *string) {
	o.Groupn = groupn
}

// WithGroupID adds the groupID to the ipam vlans list params
func (o *IpamVlansListParams) WithGroupID(groupID *string) *IpamVlansListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the ipam vlans list params
func (o *IpamVlansListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithGroupIDn adds the groupIDn to the ipam vlans list params
func (o *IpamVlansListParams) WithGroupIDn(groupIDn *string) *IpamVlansListParams {
	o.SetGroupIDn(groupIDn)
	return o
}

// SetGroupIDn adds the groupIdN to the ipam vlans list params
func (o *IpamVlansListParams) SetGroupIDn(groupIDn *string) {
	o.GroupIDn = groupIDn
}

// WithID adds the id to the ipam vlans list params
func (o *IpamVlansListParams) WithID(id *string) *IpamVlansListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlans list params
func (o *IpamVlansListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the ipam vlans list params
func (o *IpamVlansListParams) WithIDIc(iDIc *string) *IpamVlansListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the ipam vlans list params
func (o *IpamVlansListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the ipam vlans list params
func (o *IpamVlansListParams) WithIDIe(iDIe *string) *IpamVlansListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the ipam vlans list params
func (o *IpamVlansListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the ipam vlans list params
func (o *IpamVlansListParams) WithIDIew(iDIew *string) *IpamVlansListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the ipam vlans list params
func (o *IpamVlansListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the ipam vlans list params
func (o *IpamVlansListParams) WithIDIsw(iDIsw *string) *IpamVlansListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the ipam vlans list params
func (o *IpamVlansListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the ipam vlans list params
func (o *IpamVlansListParams) WithIDn(iDn *string) *IpamVlansListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam vlans list params
func (o *IpamVlansListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the ipam vlans list params
func (o *IpamVlansListParams) WithIDNic(iDNic *string) *IpamVlansListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the ipam vlans list params
func (o *IpamVlansListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the ipam vlans list params
func (o *IpamVlansListParams) WithIDNie(iDNie *string) *IpamVlansListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the ipam vlans list params
func (o *IpamVlansListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the ipam vlans list params
func (o *IpamVlansListParams) WithIDNiew(iDNiew *string) *IpamVlansListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the ipam vlans list params
func (o *IpamVlansListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the ipam vlans list params
func (o *IpamVlansListParams) WithIDNisw(iDNisw *string) *IpamVlansListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the ipam vlans list params
func (o *IpamVlansListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the ipam vlans list params
func (o *IpamVlansListParams) WithLastUpdated(lastUpdated *string) *IpamVlansListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam vlans list params
func (o *IpamVlansListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam vlans list params
func (o *IpamVlansListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamVlansListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam vlans list params
func (o *IpamVlansListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam vlans list params
func (o *IpamVlansListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamVlansListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam vlans list params
func (o *IpamVlansListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam vlans list params
func (o *IpamVlansListParams) WithLimit(limit *int64) *IpamVlansListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vlans list params
func (o *IpamVlansListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vlans list params
func (o *IpamVlansListParams) WithName(name *string) *IpamVlansListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vlans list params
func (o *IpamVlansListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the ipam vlans list params
func (o *IpamVlansListParams) WithNameIc(nameIc *string) *IpamVlansListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam vlans list params
func (o *IpamVlansListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam vlans list params
func (o *IpamVlansListParams) WithNameIe(nameIe *string) *IpamVlansListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam vlans list params
func (o *IpamVlansListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam vlans list params
func (o *IpamVlansListParams) WithNameIew(nameIew *string) *IpamVlansListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam vlans list params
func (o *IpamVlansListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam vlans list params
func (o *IpamVlansListParams) WithNameIsw(nameIsw *string) *IpamVlansListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam vlans list params
func (o *IpamVlansListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam vlans list params
func (o *IpamVlansListParams) WithNamen(namen *string) *IpamVlansListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam vlans list params
func (o *IpamVlansListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam vlans list params
func (o *IpamVlansListParams) WithNameNic(nameNic *string) *IpamVlansListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam vlans list params
func (o *IpamVlansListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam vlans list params
func (o *IpamVlansListParams) WithNameNie(nameNie *string) *IpamVlansListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam vlans list params
func (o *IpamVlansListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam vlans list params
func (o *IpamVlansListParams) WithNameNiew(nameNiew *string) *IpamVlansListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam vlans list params
func (o *IpamVlansListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam vlans list params
func (o *IpamVlansListParams) WithNameNisw(nameNisw *string) *IpamVlansListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam vlans list params
func (o *IpamVlansListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam vlans list params
func (o *IpamVlansListParams) WithOffset(offset *int64) *IpamVlansListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vlans list params
func (o *IpamVlansListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vlans list params
func (o *IpamVlansListParams) WithQ(q *string) *IpamVlansListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vlans list params
func (o *IpamVlansListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the ipam vlans list params
func (o *IpamVlansListParams) WithRegion(region *string) *IpamVlansListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the ipam vlans list params
func (o *IpamVlansListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the ipam vlans list params
func (o *IpamVlansListParams) WithRegionn(regionn *string) *IpamVlansListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the ipam vlans list params
func (o *IpamVlansListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the ipam vlans list params
func (o *IpamVlansListParams) WithRegionID(regionID *string) *IpamVlansListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the ipam vlans list params
func (o *IpamVlansListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the ipam vlans list params
func (o *IpamVlansListParams) WithRegionIDn(regionIDn *string) *IpamVlansListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the ipam vlans list params
func (o *IpamVlansListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithRole adds the role to the ipam vlans list params
func (o *IpamVlansListParams) WithRole(role *string) *IpamVlansListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam vlans list params
func (o *IpamVlansListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the ipam vlans list params
func (o *IpamVlansListParams) WithRolen(rolen *string) *IpamVlansListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the ipam vlans list params
func (o *IpamVlansListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithRoleID adds the roleID to the ipam vlans list params
func (o *IpamVlansListParams) WithRoleID(roleID *string) *IpamVlansListParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the ipam vlans list params
func (o *IpamVlansListParams) SetRoleID(roleID *string) {
	o.RoleID = roleID
}

// WithRoleIDn adds the roleIDn to the ipam vlans list params
func (o *IpamVlansListParams) WithRoleIDn(roleIDn *string) *IpamVlansListParams {
	o.SetRoleIDn(roleIDn)
	return o
}

// SetRoleIDn adds the roleIdN to the ipam vlans list params
func (o *IpamVlansListParams) SetRoleIDn(roleIDn *string) {
	o.RoleIDn = roleIDn
}

// WithSite adds the site to the ipam vlans list params
func (o *IpamVlansListParams) WithSite(site *string) *IpamVlansListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the ipam vlans list params
func (o *IpamVlansListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the ipam vlans list params
func (o *IpamVlansListParams) WithSiten(siten *string) *IpamVlansListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the ipam vlans list params
func (o *IpamVlansListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the ipam vlans list params
func (o *IpamVlansListParams) WithSiteID(siteID *string) *IpamVlansListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the ipam vlans list params
func (o *IpamVlansListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the ipam vlans list params
func (o *IpamVlansListParams) WithSiteIDn(siteIDn *string) *IpamVlansListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the ipam vlans list params
func (o *IpamVlansListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithStatus adds the status to the ipam vlans list params
func (o *IpamVlansListParams) WithStatus(status *string) *IpamVlansListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam vlans list params
func (o *IpamVlansListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the ipam vlans list params
func (o *IpamVlansListParams) WithStatusn(statusn *string) *IpamVlansListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the ipam vlans list params
func (o *IpamVlansListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the ipam vlans list params
func (o *IpamVlansListParams) WithTag(tag *string) *IpamVlansListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam vlans list params
func (o *IpamVlansListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam vlans list params
func (o *IpamVlansListParams) WithTagn(tagn *string) *IpamVlansListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam vlans list params
func (o *IpamVlansListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the ipam vlans list params
func (o *IpamVlansListParams) WithTenant(tenant *string) *IpamVlansListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam vlans list params
func (o *IpamVlansListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantn(tenantn *string) *IpamVlansListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantGroup(tenantGroup *string) *IpamVlansListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantGroupn(tenantGroupn *string) *IpamVlansListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantGroupID(tenantGroupID *string) *IpamVlansListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantGroupIDn(tenantGroupIDn *string) *IpamVlansListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantID(tenantID *string) *IpamVlansListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the ipam vlans list params
func (o *IpamVlansListParams) WithTenantIDn(tenantIDn *string) *IpamVlansListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the ipam vlans list params
func (o *IpamVlansListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithVid adds the vid to the ipam vlans list params
func (o *IpamVlansListParams) WithVid(vid *string) *IpamVlansListParams {
	o.SetVid(vid)
	return o
}

// SetVid adds the vid to the ipam vlans list params
func (o *IpamVlansListParams) SetVid(vid *string) {
	o.Vid = vid
}

// WithVidGt adds the vidGt to the ipam vlans list params
func (o *IpamVlansListParams) WithVidGt(vidGt *string) *IpamVlansListParams {
	o.SetVidGt(vidGt)
	return o
}

// SetVidGt adds the vidGt to the ipam vlans list params
func (o *IpamVlansListParams) SetVidGt(vidGt *string) {
	o.VidGt = vidGt
}

// WithVidGte adds the vidGte to the ipam vlans list params
func (o *IpamVlansListParams) WithVidGte(vidGte *string) *IpamVlansListParams {
	o.SetVidGte(vidGte)
	return o
}

// SetVidGte adds the vidGte to the ipam vlans list params
func (o *IpamVlansListParams) SetVidGte(vidGte *string) {
	o.VidGte = vidGte
}

// WithVidLt adds the vidLt to the ipam vlans list params
func (o *IpamVlansListParams) WithVidLt(vidLt *string) *IpamVlansListParams {
	o.SetVidLt(vidLt)
	return o
}

// SetVidLt adds the vidLt to the ipam vlans list params
func (o *IpamVlansListParams) SetVidLt(vidLt *string) {
	o.VidLt = vidLt
}

// WithVidLte adds the vidLte to the ipam vlans list params
func (o *IpamVlansListParams) WithVidLte(vidLte *string) *IpamVlansListParams {
	o.SetVidLte(vidLte)
	return o
}

// SetVidLte adds the vidLte to the ipam vlans list params
func (o *IpamVlansListParams) SetVidLte(vidLte *string) {
	o.VidLte = vidLte
}

// WithVidn adds the vidn to the ipam vlans list params
func (o *IpamVlansListParams) WithVidn(vidn *string) *IpamVlansListParams {
	o.SetVidn(vidn)
	return o
}

// SetVidn adds the vidN to the ipam vlans list params
func (o *IpamVlansListParams) SetVidn(vidn *string) {
	o.Vidn = vidn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVlansListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.Vid != nil {

		// query param vid
		var qrVid string

		if o.Vid != nil {
			qrVid = *o.Vid
		}
		qVid := qrVid
		if qVid != "" {

			if err := r.SetQueryParam("vid", qVid); err != nil {
				return err
			}
		}
	}

	if o.VidGt != nil {

		// query param vid__gt
		var qrVidGt string

		if o.VidGt != nil {
			qrVidGt = *o.VidGt
		}
		qVidGt := qrVidGt
		if qVidGt != "" {

			if err := r.SetQueryParam("vid__gt", qVidGt); err != nil {
				return err
			}
		}
	}

	if o.VidGte != nil {

		// query param vid__gte
		var qrVidGte string

		if o.VidGte != nil {
			qrVidGte = *o.VidGte
		}
		qVidGte := qrVidGte
		if qVidGte != "" {

			if err := r.SetQueryParam("vid__gte", qVidGte); err != nil {
				return err
			}
		}
	}

	if o.VidLt != nil {

		// query param vid__lt
		var qrVidLt string

		if o.VidLt != nil {
			qrVidLt = *o.VidLt
		}
		qVidLt := qrVidLt
		if qVidLt != "" {

			if err := r.SetQueryParam("vid__lt", qVidLt); err != nil {
				return err
			}
		}
	}

	if o.VidLte != nil {

		// query param vid__lte
		var qrVidLte string

		if o.VidLte != nil {
			qrVidLte = *o.VidLte
		}
		qVidLte := qrVidLte
		if qVidLte != "" {

			if err := r.SetQueryParam("vid__lte", qVidLte); err != nil {
				return err
			}
		}
	}

	if o.Vidn != nil {

		// query param vid__n
		var qrVidn string

		if o.Vidn != nil {
			qrVidn = *o.Vidn
		}
		qVidn := qrVidn
		if qVidn != "" {

			if err := r.SetQueryParam("vid__n", qVidn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
