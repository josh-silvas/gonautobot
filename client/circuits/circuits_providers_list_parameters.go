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

// NewCircuitsProvidersListParams creates a new CircuitsProvidersListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCircuitsProvidersListParams() *CircuitsProvidersListParams {
	return &CircuitsProvidersListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsProvidersListParamsWithTimeout creates a new CircuitsProvidersListParams object
// with the ability to set a timeout on a request.
func NewCircuitsProvidersListParamsWithTimeout(timeout time.Duration) *CircuitsProvidersListParams {
	return &CircuitsProvidersListParams{
		timeout: timeout,
	}
}

// NewCircuitsProvidersListParamsWithContext creates a new CircuitsProvidersListParams object
// with the ability to set a context for a request.
func NewCircuitsProvidersListParamsWithContext(ctx context.Context) *CircuitsProvidersListParams {
	return &CircuitsProvidersListParams{
		Context: ctx,
	}
}

// NewCircuitsProvidersListParamsWithHTTPClient creates a new CircuitsProvidersListParams object
// with the ability to set a custom HTTPClient for a request.
func NewCircuitsProvidersListParamsWithHTTPClient(client *http.Client) *CircuitsProvidersListParams {
	return &CircuitsProvidersListParams{
		HTTPClient: client,
	}
}

/* CircuitsProvidersListParams contains all the parameters to send to the API endpoint
   for the circuits providers list operation.

   Typically these are written to a http.Request.
*/
type CircuitsProvidersListParams struct {

	// Account.
	Account *string

	// AccountIc.
	AccountIc *string

	// AccountIe.
	AccountIe *string

	// AccountIew.
	AccountIew *string

	// AccountIsw.
	AccountIsw *string

	// Accountn.
	Accountn *string

	// AccountNic.
	AccountNic *string

	// AccountNie.
	AccountNie *string

	// AccountNiew.
	AccountNiew *string

	// AccountNisw.
	AccountNisw *string

	// Asn.
	Asn *string

	// AsnGt.
	AsnGt *string

	// AsnGte.
	AsnGte *string

	// AsnLt.
	AsnLt *string

	// AsnLte.
	AsnLte *string

	// Asnn.
	Asnn *string

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

	// Site.
	Site *string

	// Siten.
	Siten *string

	// SiteID.
	SiteID *string

	// SiteIDn.
	SiteIDn *string

	// Slug.
	Slug *string

	// SlugIc.
	SlugIc *string

	// SlugIe.
	SlugIe *string

	// SlugIew.
	SlugIew *string

	// SlugIsw.
	SlugIsw *string

	// Slugn.
	Slugn *string

	// SlugNic.
	SlugNic *string

	// SlugNie.
	SlugNie *string

	// SlugNiew.
	SlugNiew *string

	// SlugNisw.
	SlugNisw *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the circuits providers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersListParams) WithDefaults() *CircuitsProvidersListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the circuits providers list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CircuitsProvidersListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the circuits providers list params
func (o *CircuitsProvidersListParams) WithTimeout(timeout time.Duration) *CircuitsProvidersListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits providers list params
func (o *CircuitsProvidersListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits providers list params
func (o *CircuitsProvidersListParams) WithContext(ctx context.Context) *CircuitsProvidersListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits providers list params
func (o *CircuitsProvidersListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits providers list params
func (o *CircuitsProvidersListParams) WithHTTPClient(client *http.Client) *CircuitsProvidersListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits providers list params
func (o *CircuitsProvidersListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccount adds the account to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccount(account *string) *CircuitsProvidersListParams {
	o.SetAccount(account)
	return o
}

// SetAccount adds the account to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccount(account *string) {
	o.Account = account
}

// WithAccountIc adds the accountIc to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountIc(accountIc *string) *CircuitsProvidersListParams {
	o.SetAccountIc(accountIc)
	return o
}

// SetAccountIc adds the accountIc to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountIc(accountIc *string) {
	o.AccountIc = accountIc
}

// WithAccountIe adds the accountIe to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountIe(accountIe *string) *CircuitsProvidersListParams {
	o.SetAccountIe(accountIe)
	return o
}

// SetAccountIe adds the accountIe to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountIe(accountIe *string) {
	o.AccountIe = accountIe
}

// WithAccountIew adds the accountIew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountIew(accountIew *string) *CircuitsProvidersListParams {
	o.SetAccountIew(accountIew)
	return o
}

// SetAccountIew adds the accountIew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountIew(accountIew *string) {
	o.AccountIew = accountIew
}

// WithAccountIsw adds the accountIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountIsw(accountIsw *string) *CircuitsProvidersListParams {
	o.SetAccountIsw(accountIsw)
	return o
}

// SetAccountIsw adds the accountIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountIsw(accountIsw *string) {
	o.AccountIsw = accountIsw
}

// WithAccountn adds the accountn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountn(accountn *string) *CircuitsProvidersListParams {
	o.SetAccountn(accountn)
	return o
}

// SetAccountn adds the accountN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountn(accountn *string) {
	o.Accountn = accountn
}

// WithAccountNic adds the accountNic to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountNic(accountNic *string) *CircuitsProvidersListParams {
	o.SetAccountNic(accountNic)
	return o
}

// SetAccountNic adds the accountNic to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountNic(accountNic *string) {
	o.AccountNic = accountNic
}

// WithAccountNie adds the accountNie to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountNie(accountNie *string) *CircuitsProvidersListParams {
	o.SetAccountNie(accountNie)
	return o
}

// SetAccountNie adds the accountNie to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountNie(accountNie *string) {
	o.AccountNie = accountNie
}

// WithAccountNiew adds the accountNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountNiew(accountNiew *string) *CircuitsProvidersListParams {
	o.SetAccountNiew(accountNiew)
	return o
}

// SetAccountNiew adds the accountNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountNiew(accountNiew *string) {
	o.AccountNiew = accountNiew
}

// WithAccountNisw adds the accountNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAccountNisw(accountNisw *string) *CircuitsProvidersListParams {
	o.SetAccountNisw(accountNisw)
	return o
}

// SetAccountNisw adds the accountNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAccountNisw(accountNisw *string) {
	o.AccountNisw = accountNisw
}

// WithAsn adds the asn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAsn(asn *string) *CircuitsProvidersListParams {
	o.SetAsn(asn)
	return o
}

// SetAsn adds the asn to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAsn(asn *string) {
	o.Asn = asn
}

// WithAsnGt adds the asnGt to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAsnGt(asnGt *string) *CircuitsProvidersListParams {
	o.SetAsnGt(asnGt)
	return o
}

// SetAsnGt adds the asnGt to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAsnGt(asnGt *string) {
	o.AsnGt = asnGt
}

// WithAsnGte adds the asnGte to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAsnGte(asnGte *string) *CircuitsProvidersListParams {
	o.SetAsnGte(asnGte)
	return o
}

// SetAsnGte adds the asnGte to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAsnGte(asnGte *string) {
	o.AsnGte = asnGte
}

// WithAsnLt adds the asnLt to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAsnLt(asnLt *string) *CircuitsProvidersListParams {
	o.SetAsnLt(asnLt)
	return o
}

// SetAsnLt adds the asnLt to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAsnLt(asnLt *string) {
	o.AsnLt = asnLt
}

// WithAsnLte adds the asnLte to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAsnLte(asnLte *string) *CircuitsProvidersListParams {
	o.SetAsnLte(asnLte)
	return o
}

// SetAsnLte adds the asnLte to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAsnLte(asnLte *string) {
	o.AsnLte = asnLte
}

// WithAsnn adds the asnn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithAsnn(asnn *string) *CircuitsProvidersListParams {
	o.SetAsnn(asnn)
	return o
}

// SetAsnn adds the asnN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetAsnn(asnn *string) {
	o.Asnn = asnn
}

// WithCreated adds the created to the circuits providers list params
func (o *CircuitsProvidersListParams) WithCreated(created *string) *CircuitsProvidersListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the circuits providers list params
func (o *CircuitsProvidersListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the circuits providers list params
func (o *CircuitsProvidersListParams) WithCreatedGte(createdGte *string) *CircuitsProvidersListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the circuits providers list params
func (o *CircuitsProvidersListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the circuits providers list params
func (o *CircuitsProvidersListParams) WithCreatedLte(createdLte *string) *CircuitsProvidersListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the circuits providers list params
func (o *CircuitsProvidersListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithID adds the id to the circuits providers list params
func (o *CircuitsProvidersListParams) WithID(id *string) *CircuitsProvidersListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits providers list params
func (o *CircuitsProvidersListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDIc(iDIc *string) *CircuitsProvidersListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDIe(iDIe *string) *CircuitsProvidersListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDIew(iDIew *string) *CircuitsProvidersListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDIsw(iDIsw *string) *CircuitsProvidersListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDn(iDn *string) *CircuitsProvidersListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDNic(iDNic *string) *CircuitsProvidersListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDNie(iDNie *string) *CircuitsProvidersListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDNiew(iDNiew *string) *CircuitsProvidersListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithIDNisw(iDNisw *string) *CircuitsProvidersListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the circuits providers list params
func (o *CircuitsProvidersListParams) WithLastUpdated(lastUpdated *string) *CircuitsProvidersListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the circuits providers list params
func (o *CircuitsProvidersListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the circuits providers list params
func (o *CircuitsProvidersListParams) WithLastUpdatedGte(lastUpdatedGte *string) *CircuitsProvidersListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the circuits providers list params
func (o *CircuitsProvidersListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the circuits providers list params
func (o *CircuitsProvidersListParams) WithLastUpdatedLte(lastUpdatedLte *string) *CircuitsProvidersListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the circuits providers list params
func (o *CircuitsProvidersListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the circuits providers list params
func (o *CircuitsProvidersListParams) WithLimit(limit *int64) *CircuitsProvidersListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the circuits providers list params
func (o *CircuitsProvidersListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the circuits providers list params
func (o *CircuitsProvidersListParams) WithName(name *string) *CircuitsProvidersListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the circuits providers list params
func (o *CircuitsProvidersListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameIc(nameIc *string) *CircuitsProvidersListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameIe(nameIe *string) *CircuitsProvidersListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameIew(nameIew *string) *CircuitsProvidersListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameIsw(nameIsw *string) *CircuitsProvidersListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNamen(namen *string) *CircuitsProvidersListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameNic(nameNic *string) *CircuitsProvidersListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameNie(nameNie *string) *CircuitsProvidersListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameNiew(nameNiew *string) *CircuitsProvidersListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithNameNisw(nameNisw *string) *CircuitsProvidersListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the circuits providers list params
func (o *CircuitsProvidersListParams) WithOffset(offset *int64) *CircuitsProvidersListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the circuits providers list params
func (o *CircuitsProvidersListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the circuits providers list params
func (o *CircuitsProvidersListParams) WithQ(q *string) *CircuitsProvidersListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the circuits providers list params
func (o *CircuitsProvidersListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the circuits providers list params
func (o *CircuitsProvidersListParams) WithRegion(region *string) *CircuitsProvidersListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the circuits providers list params
func (o *CircuitsProvidersListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithRegionn(regionn *string) *CircuitsProvidersListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the circuits providers list params
func (o *CircuitsProvidersListParams) WithRegionID(regionID *string) *CircuitsProvidersListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the circuits providers list params
func (o *CircuitsProvidersListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithRegionIDn(regionIDn *string) *CircuitsProvidersListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSite(site *string) *CircuitsProvidersListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSiten(siten *string) *CircuitsProvidersListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSiteID(siteID *string) *CircuitsProvidersListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSiteIDn(siteIDn *string) *CircuitsProvidersListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithSlug adds the slug to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlug(slug *string) *CircuitsProvidersListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugIc(slugIc *string) *CircuitsProvidersListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugIe(slugIe *string) *CircuitsProvidersListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugIew(slugIew *string) *CircuitsProvidersListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugIsw(slugIsw *string) *CircuitsProvidersListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugn(slugn *string) *CircuitsProvidersListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugNic(slugNic *string) *CircuitsProvidersListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugNie(slugNie *string) *CircuitsProvidersListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugNiew(slugNiew *string) *CircuitsProvidersListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) WithSlugNisw(slugNisw *string) *CircuitsProvidersListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the circuits providers list params
func (o *CircuitsProvidersListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WithTag adds the tag to the circuits providers list params
func (o *CircuitsProvidersListParams) WithTag(tag *string) *CircuitsProvidersListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the circuits providers list params
func (o *CircuitsProvidersListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the circuits providers list params
func (o *CircuitsProvidersListParams) WithTagn(tagn *string) *CircuitsProvidersListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the circuits providers list params
func (o *CircuitsProvidersListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsProvidersListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Account != nil {

		// query param account
		var qrAccount string

		if o.Account != nil {
			qrAccount = *o.Account
		}
		qAccount := qrAccount
		if qAccount != "" {

			if err := r.SetQueryParam("account", qAccount); err != nil {
				return err
			}
		}
	}

	if o.AccountIc != nil {

		// query param account__ic
		var qrAccountIc string

		if o.AccountIc != nil {
			qrAccountIc = *o.AccountIc
		}
		qAccountIc := qrAccountIc
		if qAccountIc != "" {

			if err := r.SetQueryParam("account__ic", qAccountIc); err != nil {
				return err
			}
		}
	}

	if o.AccountIe != nil {

		// query param account__ie
		var qrAccountIe string

		if o.AccountIe != nil {
			qrAccountIe = *o.AccountIe
		}
		qAccountIe := qrAccountIe
		if qAccountIe != "" {

			if err := r.SetQueryParam("account__ie", qAccountIe); err != nil {
				return err
			}
		}
	}

	if o.AccountIew != nil {

		// query param account__iew
		var qrAccountIew string

		if o.AccountIew != nil {
			qrAccountIew = *o.AccountIew
		}
		qAccountIew := qrAccountIew
		if qAccountIew != "" {

			if err := r.SetQueryParam("account__iew", qAccountIew); err != nil {
				return err
			}
		}
	}

	if o.AccountIsw != nil {

		// query param account__isw
		var qrAccountIsw string

		if o.AccountIsw != nil {
			qrAccountIsw = *o.AccountIsw
		}
		qAccountIsw := qrAccountIsw
		if qAccountIsw != "" {

			if err := r.SetQueryParam("account__isw", qAccountIsw); err != nil {
				return err
			}
		}
	}

	if o.Accountn != nil {

		// query param account__n
		var qrAccountn string

		if o.Accountn != nil {
			qrAccountn = *o.Accountn
		}
		qAccountn := qrAccountn
		if qAccountn != "" {

			if err := r.SetQueryParam("account__n", qAccountn); err != nil {
				return err
			}
		}
	}

	if o.AccountNic != nil {

		// query param account__nic
		var qrAccountNic string

		if o.AccountNic != nil {
			qrAccountNic = *o.AccountNic
		}
		qAccountNic := qrAccountNic
		if qAccountNic != "" {

			if err := r.SetQueryParam("account__nic", qAccountNic); err != nil {
				return err
			}
		}
	}

	if o.AccountNie != nil {

		// query param account__nie
		var qrAccountNie string

		if o.AccountNie != nil {
			qrAccountNie = *o.AccountNie
		}
		qAccountNie := qrAccountNie
		if qAccountNie != "" {

			if err := r.SetQueryParam("account__nie", qAccountNie); err != nil {
				return err
			}
		}
	}

	if o.AccountNiew != nil {

		// query param account__niew
		var qrAccountNiew string

		if o.AccountNiew != nil {
			qrAccountNiew = *o.AccountNiew
		}
		qAccountNiew := qrAccountNiew
		if qAccountNiew != "" {

			if err := r.SetQueryParam("account__niew", qAccountNiew); err != nil {
				return err
			}
		}
	}

	if o.AccountNisw != nil {

		// query param account__nisw
		var qrAccountNisw string

		if o.AccountNisw != nil {
			qrAccountNisw = *o.AccountNisw
		}
		qAccountNisw := qrAccountNisw
		if qAccountNisw != "" {

			if err := r.SetQueryParam("account__nisw", qAccountNisw); err != nil {
				return err
			}
		}
	}

	if o.Asn != nil {

		// query param asn
		var qrAsn string

		if o.Asn != nil {
			qrAsn = *o.Asn
		}
		qAsn := qrAsn
		if qAsn != "" {

			if err := r.SetQueryParam("asn", qAsn); err != nil {
				return err
			}
		}
	}

	if o.AsnGt != nil {

		// query param asn__gt
		var qrAsnGt string

		if o.AsnGt != nil {
			qrAsnGt = *o.AsnGt
		}
		qAsnGt := qrAsnGt
		if qAsnGt != "" {

			if err := r.SetQueryParam("asn__gt", qAsnGt); err != nil {
				return err
			}
		}
	}

	if o.AsnGte != nil {

		// query param asn__gte
		var qrAsnGte string

		if o.AsnGte != nil {
			qrAsnGte = *o.AsnGte
		}
		qAsnGte := qrAsnGte
		if qAsnGte != "" {

			if err := r.SetQueryParam("asn__gte", qAsnGte); err != nil {
				return err
			}
		}
	}

	if o.AsnLt != nil {

		// query param asn__lt
		var qrAsnLt string

		if o.AsnLt != nil {
			qrAsnLt = *o.AsnLt
		}
		qAsnLt := qrAsnLt
		if qAsnLt != "" {

			if err := r.SetQueryParam("asn__lt", qAsnLt); err != nil {
				return err
			}
		}
	}

	if o.AsnLte != nil {

		// query param asn__lte
		var qrAsnLte string

		if o.AsnLte != nil {
			qrAsnLte = *o.AsnLte
		}
		qAsnLte := qrAsnLte
		if qAsnLte != "" {

			if err := r.SetQueryParam("asn__lte", qAsnLte); err != nil {
				return err
			}
		}
	}

	if o.Asnn != nil {

		// query param asn__n
		var qrAsnn string

		if o.Asnn != nil {
			qrAsnn = *o.Asnn
		}
		qAsnn := qrAsnn
		if qAsnn != "" {

			if err := r.SetQueryParam("asn__n", qAsnn); err != nil {
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

	if o.Slug != nil {

		// query param slug
		var qrSlug string

		if o.Slug != nil {
			qrSlug = *o.Slug
		}
		qSlug := qrSlug
		if qSlug != "" {

			if err := r.SetQueryParam("slug", qSlug); err != nil {
				return err
			}
		}
	}

	if o.SlugIc != nil {

		// query param slug__ic
		var qrSlugIc string

		if o.SlugIc != nil {
			qrSlugIc = *o.SlugIc
		}
		qSlugIc := qrSlugIc
		if qSlugIc != "" {

			if err := r.SetQueryParam("slug__ic", qSlugIc); err != nil {
				return err
			}
		}
	}

	if o.SlugIe != nil {

		// query param slug__ie
		var qrSlugIe string

		if o.SlugIe != nil {
			qrSlugIe = *o.SlugIe
		}
		qSlugIe := qrSlugIe
		if qSlugIe != "" {

			if err := r.SetQueryParam("slug__ie", qSlugIe); err != nil {
				return err
			}
		}
	}

	if o.SlugIew != nil {

		// query param slug__iew
		var qrSlugIew string

		if o.SlugIew != nil {
			qrSlugIew = *o.SlugIew
		}
		qSlugIew := qrSlugIew
		if qSlugIew != "" {

			if err := r.SetQueryParam("slug__iew", qSlugIew); err != nil {
				return err
			}
		}
	}

	if o.SlugIsw != nil {

		// query param slug__isw
		var qrSlugIsw string

		if o.SlugIsw != nil {
			qrSlugIsw = *o.SlugIsw
		}
		qSlugIsw := qrSlugIsw
		if qSlugIsw != "" {

			if err := r.SetQueryParam("slug__isw", qSlugIsw); err != nil {
				return err
			}
		}
	}

	if o.Slugn != nil {

		// query param slug__n
		var qrSlugn string

		if o.Slugn != nil {
			qrSlugn = *o.Slugn
		}
		qSlugn := qrSlugn
		if qSlugn != "" {

			if err := r.SetQueryParam("slug__n", qSlugn); err != nil {
				return err
			}
		}
	}

	if o.SlugNic != nil {

		// query param slug__nic
		var qrSlugNic string

		if o.SlugNic != nil {
			qrSlugNic = *o.SlugNic
		}
		qSlugNic := qrSlugNic
		if qSlugNic != "" {

			if err := r.SetQueryParam("slug__nic", qSlugNic); err != nil {
				return err
			}
		}
	}

	if o.SlugNie != nil {

		// query param slug__nie
		var qrSlugNie string

		if o.SlugNie != nil {
			qrSlugNie = *o.SlugNie
		}
		qSlugNie := qrSlugNie
		if qSlugNie != "" {

			if err := r.SetQueryParam("slug__nie", qSlugNie); err != nil {
				return err
			}
		}
	}

	if o.SlugNiew != nil {

		// query param slug__niew
		var qrSlugNiew string

		if o.SlugNiew != nil {
			qrSlugNiew = *o.SlugNiew
		}
		qSlugNiew := qrSlugNiew
		if qSlugNiew != "" {

			if err := r.SetQueryParam("slug__niew", qSlugNiew); err != nil {
				return err
			}
		}
	}

	if o.SlugNisw != nil {

		// query param slug__nisw
		var qrSlugNisw string

		if o.SlugNisw != nil {
			qrSlugNisw = *o.SlugNisw
		}
		qSlugNisw := qrSlugNisw
		if qSlugNisw != "" {

			if err := r.SetQueryParam("slug__nisw", qSlugNisw); err != nil {
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
