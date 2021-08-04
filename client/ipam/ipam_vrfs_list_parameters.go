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

// NewIpamVrfsListParams creates a new IpamVrfsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamVrfsListParams() *IpamVrfsListParams {
	return &IpamVrfsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamVrfsListParamsWithTimeout creates a new IpamVrfsListParams object
// with the ability to set a timeout on a request.
func NewIpamVrfsListParamsWithTimeout(timeout time.Duration) *IpamVrfsListParams {
	return &IpamVrfsListParams{
		timeout: timeout,
	}
}

// NewIpamVrfsListParamsWithContext creates a new IpamVrfsListParams object
// with the ability to set a context for a request.
func NewIpamVrfsListParamsWithContext(ctx context.Context) *IpamVrfsListParams {
	return &IpamVrfsListParams{
		Context: ctx,
	}
}

// NewIpamVrfsListParamsWithHTTPClient creates a new IpamVrfsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamVrfsListParamsWithHTTPClient(client *http.Client) *IpamVrfsListParams {
	return &IpamVrfsListParams{
		HTTPClient: client,
	}
}

/* IpamVrfsListParams contains all the parameters to send to the API endpoint
   for the ipam vrfs list operation.

   Typically these are written to a http.Request.
*/
type IpamVrfsListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// EnforceUnique.
	EnforceUnique *string

	// ExportTarget.
	ExportTarget *string

	// ExportTargetn.
	ExportTargetn *string

	// ExportTargetID.
	ExportTargetID *string

	// ExportTargetIDn.
	ExportTargetIDn *string

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

	// ImportTarget.
	ImportTarget *string

	// ImportTargetn.
	ImportTargetn *string

	// ImportTargetID.
	ImportTargetID *string

	// ImportTargetIDn.
	ImportTargetIDn *string

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

	// Rd.
	Rd *string

	// RdIc.
	RdIc *string

	// RdIe.
	RdIe *string

	// RdIew.
	RdIew *string

	// RdIsw.
	RdIsw *string

	// Rdn.
	Rdn *string

	// RdNic.
	RdNic *string

	// RdNie.
	RdNie *string

	// RdNiew.
	RdNiew *string

	// RdNisw.
	RdNisw *string

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

// WithDefaults hydrates default values in the ipam vrfs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsListParams) WithDefaults() *IpamVrfsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam vrfs list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamVrfsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTimeout(timeout time.Duration) *IpamVrfsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vrfs list params
func (o *IpamVrfsListParams) WithContext(ctx context.Context) *IpamVrfsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vrfs list params
func (o *IpamVrfsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vrfs list params
func (o *IpamVrfsListParams) WithHTTPClient(client *http.Client) *IpamVrfsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vrfs list params
func (o *IpamVrfsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam vrfs list params
func (o *IpamVrfsListParams) WithCreated(created *string) *IpamVrfsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam vrfs list params
func (o *IpamVrfsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam vrfs list params
func (o *IpamVrfsListParams) WithCreatedGte(createdGte *string) *IpamVrfsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam vrfs list params
func (o *IpamVrfsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam vrfs list params
func (o *IpamVrfsListParams) WithCreatedLte(createdLte *string) *IpamVrfsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam vrfs list params
func (o *IpamVrfsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithEnforceUnique adds the enforceUnique to the ipam vrfs list params
func (o *IpamVrfsListParams) WithEnforceUnique(enforceUnique *string) *IpamVrfsListParams {
	o.SetEnforceUnique(enforceUnique)
	return o
}

// SetEnforceUnique adds the enforceUnique to the ipam vrfs list params
func (o *IpamVrfsListParams) SetEnforceUnique(enforceUnique *string) {
	o.EnforceUnique = enforceUnique
}

// WithExportTarget adds the exportTarget to the ipam vrfs list params
func (o *IpamVrfsListParams) WithExportTarget(exportTarget *string) *IpamVrfsListParams {
	o.SetExportTarget(exportTarget)
	return o
}

// SetExportTarget adds the exportTarget to the ipam vrfs list params
func (o *IpamVrfsListParams) SetExportTarget(exportTarget *string) {
	o.ExportTarget = exportTarget
}

// WithExportTargetn adds the exportTargetn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithExportTargetn(exportTargetn *string) *IpamVrfsListParams {
	o.SetExportTargetn(exportTargetn)
	return o
}

// SetExportTargetn adds the exportTargetN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetExportTargetn(exportTargetn *string) {
	o.ExportTargetn = exportTargetn
}

// WithExportTargetID adds the exportTargetID to the ipam vrfs list params
func (o *IpamVrfsListParams) WithExportTargetID(exportTargetID *string) *IpamVrfsListParams {
	o.SetExportTargetID(exportTargetID)
	return o
}

// SetExportTargetID adds the exportTargetId to the ipam vrfs list params
func (o *IpamVrfsListParams) SetExportTargetID(exportTargetID *string) {
	o.ExportTargetID = exportTargetID
}

// WithExportTargetIDn adds the exportTargetIDn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithExportTargetIDn(exportTargetIDn *string) *IpamVrfsListParams {
	o.SetExportTargetIDn(exportTargetIDn)
	return o
}

// SetExportTargetIDn adds the exportTargetIdN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetExportTargetIDn(exportTargetIDn *string) {
	o.ExportTargetIDn = exportTargetIDn
}

// WithID adds the id to the ipam vrfs list params
func (o *IpamVrfsListParams) WithID(id *string) *IpamVrfsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vrfs list params
func (o *IpamVrfsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDIc(iDIc *string) *IpamVrfsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDIe(iDIe *string) *IpamVrfsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDIew(iDIew *string) *IpamVrfsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDIsw(iDIsw *string) *IpamVrfsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDn(iDn *string) *IpamVrfsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDNic(iDNic *string) *IpamVrfsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDNie(iDNie *string) *IpamVrfsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDNiew(iDNiew *string) *IpamVrfsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the ipam vrfs list params
func (o *IpamVrfsListParams) WithIDNisw(iDNisw *string) *IpamVrfsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the ipam vrfs list params
func (o *IpamVrfsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithImportTarget adds the importTarget to the ipam vrfs list params
func (o *IpamVrfsListParams) WithImportTarget(importTarget *string) *IpamVrfsListParams {
	o.SetImportTarget(importTarget)
	return o
}

// SetImportTarget adds the importTarget to the ipam vrfs list params
func (o *IpamVrfsListParams) SetImportTarget(importTarget *string) {
	o.ImportTarget = importTarget
}

// WithImportTargetn adds the importTargetn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithImportTargetn(importTargetn *string) *IpamVrfsListParams {
	o.SetImportTargetn(importTargetn)
	return o
}

// SetImportTargetn adds the importTargetN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetImportTargetn(importTargetn *string) {
	o.ImportTargetn = importTargetn
}

// WithImportTargetID adds the importTargetID to the ipam vrfs list params
func (o *IpamVrfsListParams) WithImportTargetID(importTargetID *string) *IpamVrfsListParams {
	o.SetImportTargetID(importTargetID)
	return o
}

// SetImportTargetID adds the importTargetId to the ipam vrfs list params
func (o *IpamVrfsListParams) SetImportTargetID(importTargetID *string) {
	o.ImportTargetID = importTargetID
}

// WithImportTargetIDn adds the importTargetIDn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithImportTargetIDn(importTargetIDn *string) *IpamVrfsListParams {
	o.SetImportTargetIDn(importTargetIDn)
	return o
}

// SetImportTargetIDn adds the importTargetIdN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetImportTargetIDn(importTargetIDn *string) {
	o.ImportTargetIDn = importTargetIDn
}

// WithLastUpdated adds the lastUpdated to the ipam vrfs list params
func (o *IpamVrfsListParams) WithLastUpdated(lastUpdated *string) *IpamVrfsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam vrfs list params
func (o *IpamVrfsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam vrfs list params
func (o *IpamVrfsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamVrfsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam vrfs list params
func (o *IpamVrfsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam vrfs list params
func (o *IpamVrfsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamVrfsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam vrfs list params
func (o *IpamVrfsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam vrfs list params
func (o *IpamVrfsListParams) WithLimit(limit *int64) *IpamVrfsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam vrfs list params
func (o *IpamVrfsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam vrfs list params
func (o *IpamVrfsListParams) WithName(name *string) *IpamVrfsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam vrfs list params
func (o *IpamVrfsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameIc(nameIc *string) *IpamVrfsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameIe(nameIe *string) *IpamVrfsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameIew(nameIew *string) *IpamVrfsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameIsw(nameIsw *string) *IpamVrfsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNamen(namen *string) *IpamVrfsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameNic(nameNic *string) *IpamVrfsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameNie(nameNie *string) *IpamVrfsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameNiew(nameNiew *string) *IpamVrfsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam vrfs list params
func (o *IpamVrfsListParams) WithNameNisw(nameNisw *string) *IpamVrfsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam vrfs list params
func (o *IpamVrfsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam vrfs list params
func (o *IpamVrfsListParams) WithOffset(offset *int64) *IpamVrfsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam vrfs list params
func (o *IpamVrfsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the ipam vrfs list params
func (o *IpamVrfsListParams) WithQ(q *string) *IpamVrfsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam vrfs list params
func (o *IpamVrfsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRd adds the rd to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRd(rd *string) *IpamVrfsListParams {
	o.SetRd(rd)
	return o
}

// SetRd adds the rd to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRd(rd *string) {
	o.Rd = rd
}

// WithRdIc adds the rdIc to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdIc(rdIc *string) *IpamVrfsListParams {
	o.SetRdIc(rdIc)
	return o
}

// SetRdIc adds the rdIc to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdIc(rdIc *string) {
	o.RdIc = rdIc
}

// WithRdIe adds the rdIe to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdIe(rdIe *string) *IpamVrfsListParams {
	o.SetRdIe(rdIe)
	return o
}

// SetRdIe adds the rdIe to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdIe(rdIe *string) {
	o.RdIe = rdIe
}

// WithRdIew adds the rdIew to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdIew(rdIew *string) *IpamVrfsListParams {
	o.SetRdIew(rdIew)
	return o
}

// SetRdIew adds the rdIew to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdIew(rdIew *string) {
	o.RdIew = rdIew
}

// WithRdIsw adds the rdIsw to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdIsw(rdIsw *string) *IpamVrfsListParams {
	o.SetRdIsw(rdIsw)
	return o
}

// SetRdIsw adds the rdIsw to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdIsw(rdIsw *string) {
	o.RdIsw = rdIsw
}

// WithRdn adds the rdn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdn(rdn *string) *IpamVrfsListParams {
	o.SetRdn(rdn)
	return o
}

// SetRdn adds the rdN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdn(rdn *string) {
	o.Rdn = rdn
}

// WithRdNic adds the rdNic to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdNic(rdNic *string) *IpamVrfsListParams {
	o.SetRdNic(rdNic)
	return o
}

// SetRdNic adds the rdNic to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdNic(rdNic *string) {
	o.RdNic = rdNic
}

// WithRdNie adds the rdNie to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdNie(rdNie *string) *IpamVrfsListParams {
	o.SetRdNie(rdNie)
	return o
}

// SetRdNie adds the rdNie to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdNie(rdNie *string) {
	o.RdNie = rdNie
}

// WithRdNiew adds the rdNiew to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdNiew(rdNiew *string) *IpamVrfsListParams {
	o.SetRdNiew(rdNiew)
	return o
}

// SetRdNiew adds the rdNiew to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdNiew(rdNiew *string) {
	o.RdNiew = rdNiew
}

// WithRdNisw adds the rdNisw to the ipam vrfs list params
func (o *IpamVrfsListParams) WithRdNisw(rdNisw *string) *IpamVrfsListParams {
	o.SetRdNisw(rdNisw)
	return o
}

// SetRdNisw adds the rdNisw to the ipam vrfs list params
func (o *IpamVrfsListParams) SetRdNisw(rdNisw *string) {
	o.RdNisw = rdNisw
}

// WithTag adds the tag to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTag(tag *string) *IpamVrfsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTagn(tagn *string) *IpamVrfsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenant(tenant *string) *IpamVrfsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantn(tenantn *string) *IpamVrfsListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantGroup(tenantGroup *string) *IpamVrfsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantGroupn(tenantGroupn *string) *IpamVrfsListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantGroupID(tenantGroupID *string) *IpamVrfsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantGroupIDn(tenantGroupIDn *string) *IpamVrfsListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantID(tenantID *string) *IpamVrfsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the ipam vrfs list params
func (o *IpamVrfsListParams) WithTenantIDn(tenantIDn *string) *IpamVrfsListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the ipam vrfs list params
func (o *IpamVrfsListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamVrfsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnforceUnique != nil {

		// query param enforce_unique
		var qrEnforceUnique string

		if o.EnforceUnique != nil {
			qrEnforceUnique = *o.EnforceUnique
		}
		qEnforceUnique := qrEnforceUnique
		if qEnforceUnique != "" {

			if err := r.SetQueryParam("enforce_unique", qEnforceUnique); err != nil {
				return err
			}
		}
	}

	if o.ExportTarget != nil {

		// query param export_target
		var qrExportTarget string

		if o.ExportTarget != nil {
			qrExportTarget = *o.ExportTarget
		}
		qExportTarget := qrExportTarget
		if qExportTarget != "" {

			if err := r.SetQueryParam("export_target", qExportTarget); err != nil {
				return err
			}
		}
	}

	if o.ExportTargetn != nil {

		// query param export_target__n
		var qrExportTargetn string

		if o.ExportTargetn != nil {
			qrExportTargetn = *o.ExportTargetn
		}
		qExportTargetn := qrExportTargetn
		if qExportTargetn != "" {

			if err := r.SetQueryParam("export_target__n", qExportTargetn); err != nil {
				return err
			}
		}
	}

	if o.ExportTargetID != nil {

		// query param export_target_id
		var qrExportTargetID string

		if o.ExportTargetID != nil {
			qrExportTargetID = *o.ExportTargetID
		}
		qExportTargetID := qrExportTargetID
		if qExportTargetID != "" {

			if err := r.SetQueryParam("export_target_id", qExportTargetID); err != nil {
				return err
			}
		}
	}

	if o.ExportTargetIDn != nil {

		// query param export_target_id__n
		var qrExportTargetIDn string

		if o.ExportTargetIDn != nil {
			qrExportTargetIDn = *o.ExportTargetIDn
		}
		qExportTargetIDn := qrExportTargetIDn
		if qExportTargetIDn != "" {

			if err := r.SetQueryParam("export_target_id__n", qExportTargetIDn); err != nil {
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

	if o.ImportTarget != nil {

		// query param import_target
		var qrImportTarget string

		if o.ImportTarget != nil {
			qrImportTarget = *o.ImportTarget
		}
		qImportTarget := qrImportTarget
		if qImportTarget != "" {

			if err := r.SetQueryParam("import_target", qImportTarget); err != nil {
				return err
			}
		}
	}

	if o.ImportTargetn != nil {

		// query param import_target__n
		var qrImportTargetn string

		if o.ImportTargetn != nil {
			qrImportTargetn = *o.ImportTargetn
		}
		qImportTargetn := qrImportTargetn
		if qImportTargetn != "" {

			if err := r.SetQueryParam("import_target__n", qImportTargetn); err != nil {
				return err
			}
		}
	}

	if o.ImportTargetID != nil {

		// query param import_target_id
		var qrImportTargetID string

		if o.ImportTargetID != nil {
			qrImportTargetID = *o.ImportTargetID
		}
		qImportTargetID := qrImportTargetID
		if qImportTargetID != "" {

			if err := r.SetQueryParam("import_target_id", qImportTargetID); err != nil {
				return err
			}
		}
	}

	if o.ImportTargetIDn != nil {

		// query param import_target_id__n
		var qrImportTargetIDn string

		if o.ImportTargetIDn != nil {
			qrImportTargetIDn = *o.ImportTargetIDn
		}
		qImportTargetIDn := qrImportTargetIDn
		if qImportTargetIDn != "" {

			if err := r.SetQueryParam("import_target_id__n", qImportTargetIDn); err != nil {
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

	if o.Rd != nil {

		// query param rd
		var qrRd string

		if o.Rd != nil {
			qrRd = *o.Rd
		}
		qRd := qrRd
		if qRd != "" {

			if err := r.SetQueryParam("rd", qRd); err != nil {
				return err
			}
		}
	}

	if o.RdIc != nil {

		// query param rd__ic
		var qrRdIc string

		if o.RdIc != nil {
			qrRdIc = *o.RdIc
		}
		qRdIc := qrRdIc
		if qRdIc != "" {

			if err := r.SetQueryParam("rd__ic", qRdIc); err != nil {
				return err
			}
		}
	}

	if o.RdIe != nil {

		// query param rd__ie
		var qrRdIe string

		if o.RdIe != nil {
			qrRdIe = *o.RdIe
		}
		qRdIe := qrRdIe
		if qRdIe != "" {

			if err := r.SetQueryParam("rd__ie", qRdIe); err != nil {
				return err
			}
		}
	}

	if o.RdIew != nil {

		// query param rd__iew
		var qrRdIew string

		if o.RdIew != nil {
			qrRdIew = *o.RdIew
		}
		qRdIew := qrRdIew
		if qRdIew != "" {

			if err := r.SetQueryParam("rd__iew", qRdIew); err != nil {
				return err
			}
		}
	}

	if o.RdIsw != nil {

		// query param rd__isw
		var qrRdIsw string

		if o.RdIsw != nil {
			qrRdIsw = *o.RdIsw
		}
		qRdIsw := qrRdIsw
		if qRdIsw != "" {

			if err := r.SetQueryParam("rd__isw", qRdIsw); err != nil {
				return err
			}
		}
	}

	if o.Rdn != nil {

		// query param rd__n
		var qrRdn string

		if o.Rdn != nil {
			qrRdn = *o.Rdn
		}
		qRdn := qrRdn
		if qRdn != "" {

			if err := r.SetQueryParam("rd__n", qRdn); err != nil {
				return err
			}
		}
	}

	if o.RdNic != nil {

		// query param rd__nic
		var qrRdNic string

		if o.RdNic != nil {
			qrRdNic = *o.RdNic
		}
		qRdNic := qrRdNic
		if qRdNic != "" {

			if err := r.SetQueryParam("rd__nic", qRdNic); err != nil {
				return err
			}
		}
	}

	if o.RdNie != nil {

		// query param rd__nie
		var qrRdNie string

		if o.RdNie != nil {
			qrRdNie = *o.RdNie
		}
		qRdNie := qrRdNie
		if qRdNie != "" {

			if err := r.SetQueryParam("rd__nie", qRdNie); err != nil {
				return err
			}
		}
	}

	if o.RdNiew != nil {

		// query param rd__niew
		var qrRdNiew string

		if o.RdNiew != nil {
			qrRdNiew = *o.RdNiew
		}
		qRdNiew := qrRdNiew
		if qRdNiew != "" {

			if err := r.SetQueryParam("rd__niew", qRdNiew); err != nil {
				return err
			}
		}
	}

	if o.RdNisw != nil {

		// query param rd__nisw
		var qrRdNisw string

		if o.RdNisw != nil {
			qrRdNisw = *o.RdNisw
		}
		qRdNisw := qrRdNisw
		if qRdNisw != "" {

			if err := r.SetQueryParam("rd__nisw", qRdNisw); err != nil {
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
