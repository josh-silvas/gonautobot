package circuits

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

// NewCircuitsCircuitsListParams creates a new CircuitsCircuitsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsCircuitsListParams() *CircuitsCircuitsListParams {
	return &CircuitsCircuitsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitsListParamsWithTimeout creates a new CircuitsCircuitsListParams object
// with the ability to set a timeout on a request.
func NewCircuitsCircuitsListParamsWithTimeout(timeout time.Duration) *CircuitsCircuitsListParams {
	return &CircuitsCircuitsListParams{
		timeout: timeout,
	}
}

// NewCircuitsCircuitsListParamsWithContext creates a new CircuitsCircuitsListParams object
// with the ability to set a context for a request.
func NewCircuitsCircuitsListParamsWithContext(ctx context.Context) *CircuitsCircuitsListParams {
	return &CircuitsCircuitsListParams{
		Context: ctx,
	}
}

// NewCircuitsCircuitsListParamsWithHTTPClient creates a new CircuitsCircuitsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsCircuitsListParamsWithHTTPClient(client *http.Client) *CircuitsCircuitsListParams {
	return &CircuitsCircuitsListParams{
		HTTPClient: client,
	}
}

/* CircuitsCircuitsListParams contains all the parameters to send to the API endpoint
   for the circuits circuits list operation.

   Typically these are written to a http.Request.
*/
type CircuitsCircuitsListParams struct {

	// Cid.
	Cid *string

	// CidIc.
	CidIc *string

	// CidIe.
	CidIe *string

	// CidIew.
	CidIew *string

	// CidIsw.
	CidIsw *string

	// Cidn.
	Cidn *string

	// CidNic.
	CidNic *string

	// CidNie.
	CidNie *string

	// CidNiew.
	CidNiew *string

	// CidNisw.
	CidNisw *string

	// CommitRate.
	CommitRate *string

	// CommitRateGt.
	CommitRateGt *string

	// CommitRateGte.
	CommitRateGte *string

	// CommitRateLt.
	CommitRateLt *string

	// CommitRateLte.
	CommitRateLte *string

	// CommitRaten.
	CommitRaten *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

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

	// InstallDate.
	InstallDate *string

	// InstallDateGt.
	InstallDateGt *string

	// InstallDateGte.
	InstallDateGte *string

	// InstallDateLt.
	InstallDateLt *string

	// InstallDateLte.
	InstallDateLte *string

	// InstallDaten.
	InstallDaten *string

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

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Provider.
	Provider *string

	// Providern.
	Providern *string

	// ProviderID.
	ProviderID *string

	// ProviderIDn.
	ProviderIDn *string

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

	// TypeID.
	TypeID *string

	// TypeIDn.
	TypeIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits circuits list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitsListParams) WithDefaults() *CircuitsCircuitsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits circuits list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsCircuitsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTimeout(timeout time.Duration) *CircuitsCircuitsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithContext(ctx context.Context) *CircuitsCircuitsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithHTTPClient(client *http.Client) *CircuitsCircuitsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCid(cid *string) *CircuitsCircuitsListParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCid(cid *string) {
	o.Cid = cid
}

// WithCidIc adds the cidIc to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidIc(cidIc *string) *CircuitsCircuitsListParams {
	o.SetCidIc(cidIc)
	return o
}

// SetCidIc adds the cidIc to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidIc(cidIc *string) {
	o.CidIc = cidIc
}

// WithCidIe adds the cidIe to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidIe(cidIe *string) *CircuitsCircuitsListParams {
	o.SetCidIe(cidIe)
	return o
}

// SetCidIe adds the cidIe to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidIe(cidIe *string) {
	o.CidIe = cidIe
}

// WithCidIew adds the cidIew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidIew(cidIew *string) *CircuitsCircuitsListParams {
	o.SetCidIew(cidIew)
	return o
}

// SetCidIew adds the cidIew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidIew(cidIew *string) {
	o.CidIew = cidIew
}

// WithCidIsw adds the cidIsw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidIsw(cidIsw *string) *CircuitsCircuitsListParams {
	o.SetCidIsw(cidIsw)
	return o
}

// SetCidIsw adds the cidIsw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidIsw(cidIsw *string) {
	o.CidIsw = cidIsw
}

// WithCidn adds the cidn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidn(cidn *string) *CircuitsCircuitsListParams {
	o.SetCidn(cidn)
	return o
}

// SetCidn adds the cidN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidn(cidn *string) {
	o.Cidn = cidn
}

// WithCidNic adds the cidNic to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidNic(cidNic *string) *CircuitsCircuitsListParams {
	o.SetCidNic(cidNic)
	return o
}

// SetCidNic adds the cidNic to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidNic(cidNic *string) {
	o.CidNic = cidNic
}

// WithCidNie adds the cidNie to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidNie(cidNie *string) *CircuitsCircuitsListParams {
	o.SetCidNie(cidNie)
	return o
}

// SetCidNie adds the cidNie to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidNie(cidNie *string) {
	o.CidNie = cidNie
}

// WithCidNiew adds the cidNiew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidNiew(cidNiew *string) *CircuitsCircuitsListParams {
	o.SetCidNiew(cidNiew)
	return o
}

// SetCidNiew adds the cidNiew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidNiew(cidNiew *string) {
	o.CidNiew = cidNiew
}

// WithCidNisw adds the cidNisw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCidNisw(cidNisw *string) *CircuitsCircuitsListParams {
	o.SetCidNisw(cidNisw)
	return o
}

// SetCidNisw adds the cidNisw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCidNisw(cidNisw *string) {
	o.CidNisw = cidNisw
}

// WithCommitRate adds the commitRate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCommitRate(commitRate *string) *CircuitsCircuitsListParams {
	o.SetCommitRate(commitRate)
	return o
}

// SetCommitRate adds the commitRate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCommitRate(commitRate *string) {
	o.CommitRate = commitRate
}

// WithCommitRateGt adds the commitRateGt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCommitRateGt(commitRateGt *string) *CircuitsCircuitsListParams {
	o.SetCommitRateGt(commitRateGt)
	return o
}

// SetCommitRateGt adds the commitRateGt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCommitRateGt(commitRateGt *string) {
	o.CommitRateGt = commitRateGt
}

// WithCommitRateGte adds the commitRateGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCommitRateGte(commitRateGte *string) *CircuitsCircuitsListParams {
	o.SetCommitRateGte(commitRateGte)
	return o
}

// SetCommitRateGte adds the commitRateGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCommitRateGte(commitRateGte *string) {
	o.CommitRateGte = commitRateGte
}

// WithCommitRateLt adds the commitRateLt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCommitRateLt(commitRateLt *string) *CircuitsCircuitsListParams {
	o.SetCommitRateLt(commitRateLt)
	return o
}

// SetCommitRateLt adds the commitRateLt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCommitRateLt(commitRateLt *string) {
	o.CommitRateLt = commitRateLt
}

// WithCommitRateLte adds the commitRateLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCommitRateLte(commitRateLte *string) *CircuitsCircuitsListParams {
	o.SetCommitRateLte(commitRateLte)
	return o
}

// SetCommitRateLte adds the commitRateLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCommitRateLte(commitRateLte *string) {
	o.CommitRateLte = commitRateLte
}

// WithCommitRaten adds the commitRaten to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCommitRaten(commitRaten *string) *CircuitsCircuitsListParams {
	o.SetCommitRaten(commitRaten)
	return o
}

// SetCommitRaten adds the commitRateN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCommitRaten(commitRaten *string) {
	o.CommitRaten = commitRaten
}

// WithCreated adds the created to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCreated(created *string) *CircuitsCircuitsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCreatedGte(createdGte *string) *CircuitsCircuitsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithCreatedLte(createdLte *string) *CircuitsCircuitsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithID adds the id to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithID(id *string) *CircuitsCircuitsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDIc(iDIc *string) *CircuitsCircuitsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDIe(iDIe *string) *CircuitsCircuitsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDIew(iDIew *string) *CircuitsCircuitsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDIsw(iDIsw *string) *CircuitsCircuitsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDn(iDn *string) *CircuitsCircuitsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDNic(iDNic *string) *CircuitsCircuitsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDNie(iDNie *string) *CircuitsCircuitsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDNiew(iDNiew *string) *CircuitsCircuitsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithIDNisw(iDNisw *string) *CircuitsCircuitsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithInstallDate adds the installDate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithInstallDate(installDate *string) *CircuitsCircuitsListParams {
	o.SetInstallDate(installDate)
	return o
}

// SetInstallDate adds the installDate to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetInstallDate(installDate *string) {
	o.InstallDate = installDate
}

// WithInstallDateGt adds the installDateGt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithInstallDateGt(installDateGt *string) *CircuitsCircuitsListParams {
	o.SetInstallDateGt(installDateGt)
	return o
}

// SetInstallDateGt adds the installDateGt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetInstallDateGt(installDateGt *string) {
	o.InstallDateGt = installDateGt
}

// WithInstallDateGte adds the installDateGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithInstallDateGte(installDateGte *string) *CircuitsCircuitsListParams {
	o.SetInstallDateGte(installDateGte)
	return o
}

// SetInstallDateGte adds the installDateGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetInstallDateGte(installDateGte *string) {
	o.InstallDateGte = installDateGte
}

// WithInstallDateLt adds the installDateLt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithInstallDateLt(installDateLt *string) *CircuitsCircuitsListParams {
	o.SetInstallDateLt(installDateLt)
	return o
}

// SetInstallDateLt adds the installDateLt to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetInstallDateLt(installDateLt *string) {
	o.InstallDateLt = installDateLt
}

// WithInstallDateLte adds the installDateLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithInstallDateLte(installDateLte *string) *CircuitsCircuitsListParams {
	o.SetInstallDateLte(installDateLte)
	return o
}

// SetInstallDateLte adds the installDateLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetInstallDateLte(installDateLte *string) {
	o.InstallDateLte = installDateLte
}

// WithInstallDaten adds the installDaten to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithInstallDaten(installDaten *string) *CircuitsCircuitsListParams {
	o.SetInstallDaten(installDaten)
	return o
}

// SetInstallDaten adds the installDateN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetInstallDaten(installDaten *string) {
	o.InstallDaten = installDaten
}

// WithLastUpdated adds the lastUpdated to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLastUpdated(lastUpdated *string) *CircuitsCircuitsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *CircuitsCircuitsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *CircuitsCircuitsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithLimit(limit *int64) *CircuitsCircuitsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithOffset(offset *int64) *CircuitsCircuitsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithProvider adds the provider to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithProvider(provider *string) *CircuitsCircuitsListParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetProvider(provider *string) {
	o.Provider = provider
}

// WithProvidern adds the providern to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithProvidern(providern *string) *CircuitsCircuitsListParams {
	o.SetProvidern(providern)
	return o
}

// SetProvidern adds the providerN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetProvidern(providern *string) {
	o.Providern = providern
}

// WithProviderID adds the providerID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithProviderID(providerID *string) *CircuitsCircuitsListParams {
	o.SetProviderID(providerID)
	return o
}

// SetProviderID adds the providerId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetProviderID(providerID *string) {
	o.ProviderID = providerID
}

// WithProviderIDn adds the providerIDn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithProviderIDn(providerIDn *string) *CircuitsCircuitsListParams {
	o.SetProviderIDn(providerIDn)
	return o
}

// SetProviderIDn adds the providerIdN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetProviderIDn(providerIDn *string) {
	o.ProviderIDn = providerIDn
}

// WithQ adds the q to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithQ(q *string) *CircuitsCircuitsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithRegion(region *string) *CircuitsCircuitsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithRegionn(regionn *string) *CircuitsCircuitsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithRegionID(regionID *string) *CircuitsCircuitsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithRegionIDn(regionIDn *string) *CircuitsCircuitsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithSite(site *string) *CircuitsCircuitsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithSiten(siten *string) *CircuitsCircuitsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithSiteID(siteID *string) *CircuitsCircuitsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithSiteIDn(siteIDn *string) *CircuitsCircuitsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithStatus adds the status to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithStatus(status *string) *CircuitsCircuitsListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithStatusn(statusn *string) *CircuitsCircuitsListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTag(tag *string) *CircuitsCircuitsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTagn(tagn *string) *CircuitsCircuitsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenant(tenant *string) *CircuitsCircuitsListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantn(tenantn *string) *CircuitsCircuitsListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantGroup(tenantGroup *string) *CircuitsCircuitsListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantGroupn(tenantGroupn *string) *CircuitsCircuitsListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantGroupID(tenantGroupID *string) *CircuitsCircuitsListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantGroupIDn(tenantGroupIDn *string) *CircuitsCircuitsListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantID(tenantID *string) *CircuitsCircuitsListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTenantIDn(tenantIDn *string) *CircuitsCircuitsListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithType adds the typeVar to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithType(typeVar *string) *CircuitsCircuitsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTypen(typen *string) *CircuitsCircuitsListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WithTypeID adds the typeID to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTypeID(typeID *string) *CircuitsCircuitsListParams {
	o.SetTypeID(typeID)
	return o
}

// SetTypeID adds the typeId to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTypeID(typeID *string) {
	o.TypeID = typeID
}

// WithTypeIDn adds the typeIDn to the circuits circuits list params
func (o *CircuitsCircuitsListParams) WithTypeIDn(typeIDn *string) *CircuitsCircuitsListParams {
	o.SetTypeIDn(typeIDn)
	return o
}

// SetTypeIDn adds the typeIdN to the circuits circuits list params
func (o *CircuitsCircuitsListParams) SetTypeIDn(typeIDn *string) {
	o.TypeIDn = typeIDn
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cid != nil {

		// query param cid
		var qrCid string

		if o.Cid != nil {
			qrCid = *o.Cid
		}
		qCid := qrCid
		if qCid != "" {

			if err := r.SetQueryParam("cid", qCid); err != nil {
				return err
			}
		}
	}

	if o.CidIc != nil {

		// query param cid__ic
		var qrCidIc string

		if o.CidIc != nil {
			qrCidIc = *o.CidIc
		}
		qCidIc := qrCidIc
		if qCidIc != "" {

			if err := r.SetQueryParam("cid__ic", qCidIc); err != nil {
				return err
			}
		}
	}

	if o.CidIe != nil {

		// query param cid__ie
		var qrCidIe string

		if o.CidIe != nil {
			qrCidIe = *o.CidIe
		}
		qCidIe := qrCidIe
		if qCidIe != "" {

			if err := r.SetQueryParam("cid__ie", qCidIe); err != nil {
				return err
			}
		}
	}

	if o.CidIew != nil {

		// query param cid__iew
		var qrCidIew string

		if o.CidIew != nil {
			qrCidIew = *o.CidIew
		}
		qCidIew := qrCidIew
		if qCidIew != "" {

			if err := r.SetQueryParam("cid__iew", qCidIew); err != nil {
				return err
			}
		}
	}

	if o.CidIsw != nil {

		// query param cid__isw
		var qrCidIsw string

		if o.CidIsw != nil {
			qrCidIsw = *o.CidIsw
		}
		qCidIsw := qrCidIsw
		if qCidIsw != "" {

			if err := r.SetQueryParam("cid__isw", qCidIsw); err != nil {
				return err
			}
		}
	}

	if o.Cidn != nil {

		// query param cid__n
		var qrCidn string

		if o.Cidn != nil {
			qrCidn = *o.Cidn
		}
		qCidn := qrCidn
		if qCidn != "" {

			if err := r.SetQueryParam("cid__n", qCidn); err != nil {
				return err
			}
		}
	}

	if o.CidNic != nil {

		// query param cid__nic
		var qrCidNic string

		if o.CidNic != nil {
			qrCidNic = *o.CidNic
		}
		qCidNic := qrCidNic
		if qCidNic != "" {

			if err := r.SetQueryParam("cid__nic", qCidNic); err != nil {
				return err
			}
		}
	}

	if o.CidNie != nil {

		// query param cid__nie
		var qrCidNie string

		if o.CidNie != nil {
			qrCidNie = *o.CidNie
		}
		qCidNie := qrCidNie
		if qCidNie != "" {

			if err := r.SetQueryParam("cid__nie", qCidNie); err != nil {
				return err
			}
		}
	}

	if o.CidNiew != nil {

		// query param cid__niew
		var qrCidNiew string

		if o.CidNiew != nil {
			qrCidNiew = *o.CidNiew
		}
		qCidNiew := qrCidNiew
		if qCidNiew != "" {

			if err := r.SetQueryParam("cid__niew", qCidNiew); err != nil {
				return err
			}
		}
	}

	if o.CidNisw != nil {

		// query param cid__nisw
		var qrCidNisw string

		if o.CidNisw != nil {
			qrCidNisw = *o.CidNisw
		}
		qCidNisw := qrCidNisw
		if qCidNisw != "" {

			if err := r.SetQueryParam("cid__nisw", qCidNisw); err != nil {
				return err
			}
		}
	}

	if o.CommitRate != nil {

		// query param commit_rate
		var qrCommitRate string

		if o.CommitRate != nil {
			qrCommitRate = *o.CommitRate
		}
		qCommitRate := qrCommitRate
		if qCommitRate != "" {

			if err := r.SetQueryParam("commit_rate", qCommitRate); err != nil {
				return err
			}
		}
	}

	if o.CommitRateGt != nil {

		// query param commit_rate__gt
		var qrCommitRateGt string

		if o.CommitRateGt != nil {
			qrCommitRateGt = *o.CommitRateGt
		}
		qCommitRateGt := qrCommitRateGt
		if qCommitRateGt != "" {

			if err := r.SetQueryParam("commit_rate__gt", qCommitRateGt); err != nil {
				return err
			}
		}
	}

	if o.CommitRateGte != nil {

		// query param commit_rate__gte
		var qrCommitRateGte string

		if o.CommitRateGte != nil {
			qrCommitRateGte = *o.CommitRateGte
		}
		qCommitRateGte := qrCommitRateGte
		if qCommitRateGte != "" {

			if err := r.SetQueryParam("commit_rate__gte", qCommitRateGte); err != nil {
				return err
			}
		}
	}

	if o.CommitRateLt != nil {

		// query param commit_rate__lt
		var qrCommitRateLt string

		if o.CommitRateLt != nil {
			qrCommitRateLt = *o.CommitRateLt
		}
		qCommitRateLt := qrCommitRateLt
		if qCommitRateLt != "" {

			if err := r.SetQueryParam("commit_rate__lt", qCommitRateLt); err != nil {
				return err
			}
		}
	}

	if o.CommitRateLte != nil {

		// query param commit_rate__lte
		var qrCommitRateLte string

		if o.CommitRateLte != nil {
			qrCommitRateLte = *o.CommitRateLte
		}
		qCommitRateLte := qrCommitRateLte
		if qCommitRateLte != "" {

			if err := r.SetQueryParam("commit_rate__lte", qCommitRateLte); err != nil {
				return err
			}
		}
	}

	if o.CommitRaten != nil {

		// query param commit_rate__n
		var qrCommitRaten string

		if o.CommitRaten != nil {
			qrCommitRaten = *o.CommitRaten
		}
		qCommitRaten := qrCommitRaten
		if qCommitRaten != "" {

			if err := r.SetQueryParam("commit_rate__n", qCommitRaten); err != nil {
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

	if o.InstallDate != nil {

		// query param install_date
		var qrInstallDate string

		if o.InstallDate != nil {
			qrInstallDate = *o.InstallDate
		}
		qInstallDate := qrInstallDate
		if qInstallDate != "" {

			if err := r.SetQueryParam("install_date", qInstallDate); err != nil {
				return err
			}
		}
	}

	if o.InstallDateGt != nil {

		// query param install_date__gt
		var qrInstallDateGt string

		if o.InstallDateGt != nil {
			qrInstallDateGt = *o.InstallDateGt
		}
		qInstallDateGt := qrInstallDateGt
		if qInstallDateGt != "" {

			if err := r.SetQueryParam("install_date__gt", qInstallDateGt); err != nil {
				return err
			}
		}
	}

	if o.InstallDateGte != nil {

		// query param install_date__gte
		var qrInstallDateGte string

		if o.InstallDateGte != nil {
			qrInstallDateGte = *o.InstallDateGte
		}
		qInstallDateGte := qrInstallDateGte
		if qInstallDateGte != "" {

			if err := r.SetQueryParam("install_date__gte", qInstallDateGte); err != nil {
				return err
			}
		}
	}

	if o.InstallDateLt != nil {

		// query param install_date__lt
		var qrInstallDateLt string

		if o.InstallDateLt != nil {
			qrInstallDateLt = *o.InstallDateLt
		}
		qInstallDateLt := qrInstallDateLt
		if qInstallDateLt != "" {

			if err := r.SetQueryParam("install_date__lt", qInstallDateLt); err != nil {
				return err
			}
		}
	}

	if o.InstallDateLte != nil {

		// query param install_date__lte
		var qrInstallDateLte string

		if o.InstallDateLte != nil {
			qrInstallDateLte = *o.InstallDateLte
		}
		qInstallDateLte := qrInstallDateLte
		if qInstallDateLte != "" {

			if err := r.SetQueryParam("install_date__lte", qInstallDateLte); err != nil {
				return err
			}
		}
	}

	if o.InstallDaten != nil {

		// query param install_date__n
		var qrInstallDaten string

		if o.InstallDaten != nil {
			qrInstallDaten = *o.InstallDaten
		}
		qInstallDaten := qrInstallDaten
		if qInstallDaten != "" {

			if err := r.SetQueryParam("install_date__n", qInstallDaten); err != nil {
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

	if o.Provider != nil {

		// query param provider
		var qrProvider string

		if o.Provider != nil {
			qrProvider = *o.Provider
		}
		qProvider := qrProvider
		if qProvider != "" {

			if err := r.SetQueryParam("provider", qProvider); err != nil {
				return err
			}
		}
	}

	if o.Providern != nil {

		// query param provider__n
		var qrProvidern string

		if o.Providern != nil {
			qrProvidern = *o.Providern
		}
		qProvidern := qrProvidern
		if qProvidern != "" {

			if err := r.SetQueryParam("provider__n", qProvidern); err != nil {
				return err
			}
		}
	}

	if o.ProviderID != nil {

		// query param provider_id
		var qrProviderID string

		if o.ProviderID != nil {
			qrProviderID = *o.ProviderID
		}
		qProviderID := qrProviderID
		if qProviderID != "" {

			if err := r.SetQueryParam("provider_id", qProviderID); err != nil {
				return err
			}
		}
	}

	if o.ProviderIDn != nil {

		// query param provider_id__n
		var qrProviderIDn string

		if o.ProviderIDn != nil {
			qrProviderIDn = *o.ProviderIDn
		}
		qProviderIDn := qrProviderIDn
		if qProviderIDn != "" {

			if err := r.SetQueryParam("provider_id__n", qProviderIDn); err != nil {
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

	if o.TypeID != nil {

		// query param type_id
		var qrTypeID string

		if o.TypeID != nil {
			qrTypeID = *o.TypeID
		}
		qTypeID := qrTypeID
		if qTypeID != "" {

			if err := r.SetQueryParam("type_id", qTypeID); err != nil {
				return err
			}
		}
	}

	if o.TypeIDn != nil {

		// query param type_id__n
		var qrTypeIDn string

		if o.TypeIDn != nil {
			qrTypeIDn = *o.TypeIDn
		}
		qTypeIDn := qrTypeIDn
		if qTypeIDn != "" {

			if err := r.SetQueryParam("type_id__n", qTypeIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
