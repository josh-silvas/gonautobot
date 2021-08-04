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

// NewDcimRackReservationsListParams creates a new DcimRackReservationsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimRackReservationsListParams() *DcimRackReservationsListParams {
	return &DcimRackReservationsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimRackReservationsListParamsWithTimeout creates a new DcimRackReservationsListParams object
// with the ability to set a timeout on a request.
func NewDcimRackReservationsListParamsWithTimeout(timeout time.Duration) *DcimRackReservationsListParams {
	return &DcimRackReservationsListParams{
		timeout: timeout,
	}
}

// NewDcimRackReservationsListParamsWithContext creates a new DcimRackReservationsListParams object
// with the ability to set a context for a request.
func NewDcimRackReservationsListParamsWithContext(ctx context.Context) *DcimRackReservationsListParams {
	return &DcimRackReservationsListParams{
		Context: ctx,
	}
}

// NewDcimRackReservationsListParamsWithHTTPClient creates a new DcimRackReservationsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimRackReservationsListParamsWithHTTPClient(client *http.Client) *DcimRackReservationsListParams {
	return &DcimRackReservationsListParams{
		HTTPClient: client,
	}
}

/* DcimRackReservationsListParams contains all the parameters to send to the API endpoint
   for the dcim rack reservations list operation.

   Typically these are written to a http.Request.
*/
type DcimRackReservationsListParams struct {

	// Created.
	Created *string

	// CreatedGt.
	CreatedGt *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLt.
	CreatedLt *string

	// CreatedLte.
	CreatedLte *string

	// Createdn.
	Createdn *string

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

	// RackID.
	RackID *string

	// RackIDn.
	RackIDn *string

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

	// User.
	User *string

	// Usern.
	Usern *string

	// UserID.
	UserID *string

	// UserIDn.
	UserIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim rack reservations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsListParams) WithDefaults() *DcimRackReservationsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim rack reservations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimRackReservationsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTimeout(timeout time.Duration) *DcimRackReservationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithContext(ctx context.Context) *DcimRackReservationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithHTTPClient(client *http.Client) *DcimRackReservationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithCreated(created *string) *DcimRackReservationsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGt adds the createdGt to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithCreatedGt(createdGt *string) *DcimRackReservationsListParams {
	o.SetCreatedGt(createdGt)
	return o
}

// SetCreatedGt adds the createdGt to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetCreatedGt(createdGt *string) {
	o.CreatedGt = createdGt
}

// WithCreatedGte adds the createdGte to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithCreatedGte(createdGte *string) *DcimRackReservationsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLt adds the createdLt to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithCreatedLt(createdLt *string) *DcimRackReservationsListParams {
	o.SetCreatedLt(createdLt)
	return o
}

// SetCreatedLt adds the createdLt to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetCreatedLt(createdLt *string) {
	o.CreatedLt = createdLt
}

// WithCreatedLte adds the createdLte to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithCreatedLte(createdLte *string) *DcimRackReservationsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithCreatedn adds the createdn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithCreatedn(createdn *string) *DcimRackReservationsListParams {
	o.SetCreatedn(createdn)
	return o
}

// SetCreatedn adds the createdN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetCreatedn(createdn *string) {
	o.Createdn = createdn
}

// WithGroup adds the group to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithGroup(group *string) *DcimRackReservationsListParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetGroup(group *string) {
	o.Group = group
}

// WithGroupn adds the groupn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithGroupn(groupn *string) *DcimRackReservationsListParams {
	o.SetGroupn(groupn)
	return o
}

// SetGroupn adds the groupN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetGroupn(groupn *string) {
	o.Groupn = groupn
}

// WithGroupID adds the groupID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithGroupID(groupID *string) *DcimRackReservationsListParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetGroupID(groupID *string) {
	o.GroupID = groupID
}

// WithGroupIDn adds the groupIDn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithGroupIDn(groupIDn *string) *DcimRackReservationsListParams {
	o.SetGroupIDn(groupIDn)
	return o
}

// SetGroupIDn adds the groupIdN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetGroupIDn(groupIDn *string) {
	o.GroupIDn = groupIDn
}

// WithID adds the id to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithID(id *string) *DcimRackReservationsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDIc(iDIc *string) *DcimRackReservationsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDIe(iDIe *string) *DcimRackReservationsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDIew(iDIew *string) *DcimRackReservationsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDIsw(iDIsw *string) *DcimRackReservationsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDn(iDn *string) *DcimRackReservationsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDNic(iDNic *string) *DcimRackReservationsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDNie(iDNie *string) *DcimRackReservationsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDNiew(iDNiew *string) *DcimRackReservationsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithIDNisw(iDNisw *string) *DcimRackReservationsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLimit adds the limit to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithLimit(limit *int64) *DcimRackReservationsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithOffset(offset *int64) *DcimRackReservationsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithQ(q *string) *DcimRackReservationsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRackID adds the rackID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithRackID(rackID *string) *DcimRackReservationsListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithRackIDn adds the rackIDn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithRackIDn(rackIDn *string) *DcimRackReservationsListParams {
	o.SetRackIDn(rackIDn)
	return o
}

// SetRackIDn adds the rackIdN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetRackIDn(rackIDn *string) {
	o.RackIDn = rackIDn
}

// WithSite adds the site to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithSite(site *string) *DcimRackReservationsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithSiten(siten *string) *DcimRackReservationsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithSiteID(siteID *string) *DcimRackReservationsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithSiteIDn(siteIDn *string) *DcimRackReservationsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithTag adds the tag to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTag(tag *string) *DcimRackReservationsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTagn(tagn *string) *DcimRackReservationsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenant(tenant *string) *DcimRackReservationsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantn(tenantn *string) *DcimRackReservationsListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantGroup(tenantGroup *string) *DcimRackReservationsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantGroupn(tenantGroupn *string) *DcimRackReservationsListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantGroupID(tenantGroupID *string) *DcimRackReservationsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantGroupIDn(tenantGroupIDn *string) *DcimRackReservationsListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantID(tenantID *string) *DcimRackReservationsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithTenantIDn(tenantIDn *string) *DcimRackReservationsListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithUser adds the user to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithUser(user *string) *DcimRackReservationsListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetUser(user *string) {
	o.User = user
}

// WithUsern adds the usern to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithUsern(usern *string) *DcimRackReservationsListParams {
	o.SetUsern(usern)
	return o
}

// SetUsern adds the userN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetUsern(usern *string) {
	o.Usern = usern
}

// WithUserID adds the userID to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithUserID(userID *string) *DcimRackReservationsListParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithUserIDn adds the userIDn to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) WithUserIDn(userIDn *string) *DcimRackReservationsListParams {
	o.SetUserIDn(userIDn)
	return o
}

// SetUserIDn adds the userIdN to the dcim rack reservations list params
func (o *DcimRackReservationsListParams) SetUserIDn(userIDn *string) {
	o.UserIDn = userIDn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimRackReservationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CreatedGt != nil {

		// query param created__gt
		var qrCreatedGt string

		if o.CreatedGt != nil {
			qrCreatedGt = *o.CreatedGt
		}
		qCreatedGt := qrCreatedGt
		if qCreatedGt != "" {

			if err := r.SetQueryParam("created__gt", qCreatedGt); err != nil {
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

	if o.CreatedLt != nil {

		// query param created__lt
		var qrCreatedLt string

		if o.CreatedLt != nil {
			qrCreatedLt = *o.CreatedLt
		}
		qCreatedLt := qrCreatedLt
		if qCreatedLt != "" {

			if err := r.SetQueryParam("created__lt", qCreatedLt); err != nil {
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

	if o.Createdn != nil {

		// query param created__n
		var qrCreatedn string

		if o.Createdn != nil {
			qrCreatedn = *o.Createdn
		}
		qCreatedn := qrCreatedn
		if qCreatedn != "" {

			if err := r.SetQueryParam("created__n", qCreatedn); err != nil {
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

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if o.Usern != nil {

		// query param user__n
		var qrUsern string

		if o.Usern != nil {
			qrUsern = *o.Usern
		}
		qUsern := qrUsern
		if qUsern != "" {

			if err := r.SetQueryParam("user__n", qUsern); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if o.UserIDn != nil {

		// query param user_id__n
		var qrUserIDn string

		if o.UserIDn != nil {
			qrUserIDn = *o.UserIDn
		}
		qUserIDn := qrUserIDn
		if qUserIDn != "" {

			if err := r.SetQueryParam("user_id__n", qUserIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
