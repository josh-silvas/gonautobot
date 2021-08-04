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

// NewDcimSitesListParams creates a new DcimSitesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimSitesListParams() *DcimSitesListParams {
	return &DcimSitesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimSitesListParamsWithTimeout creates a new DcimSitesListParams object
// with the ability to set a timeout on a request.
func NewDcimSitesListParamsWithTimeout(timeout time.Duration) *DcimSitesListParams {
	return &DcimSitesListParams{
		timeout: timeout,
	}
}

// NewDcimSitesListParamsWithContext creates a new DcimSitesListParams object
// with the ability to set a context for a request.
func NewDcimSitesListParamsWithContext(ctx context.Context) *DcimSitesListParams {
	return &DcimSitesListParams{
		Context: ctx,
	}
}

// NewDcimSitesListParamsWithHTTPClient creates a new DcimSitesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimSitesListParamsWithHTTPClient(client *http.Client) *DcimSitesListParams {
	return &DcimSitesListParams{
		HTTPClient: client,
	}
}

/* DcimSitesListParams contains all the parameters to send to the API endpoint
   for the dcim sites list operation.

   Typically these are written to a http.Request.
*/
type DcimSitesListParams struct {

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

	// ContactEmail.
	ContactEmail *string

	// ContactEmailIc.
	ContactEmailIc *string

	// ContactEmailIe.
	ContactEmailIe *string

	// ContactEmailIew.
	ContactEmailIew *string

	// ContactEmailIsw.
	ContactEmailIsw *string

	// ContactEmailn.
	ContactEmailn *string

	// ContactEmailNic.
	ContactEmailNic *string

	// ContactEmailNie.
	ContactEmailNie *string

	// ContactEmailNiew.
	ContactEmailNiew *string

	// ContactEmailNisw.
	ContactEmailNisw *string

	// ContactName.
	ContactName *string

	// ContactNameIc.
	ContactNameIc *string

	// ContactNameIe.
	ContactNameIe *string

	// ContactNameIew.
	ContactNameIew *string

	// ContactNameIsw.
	ContactNameIsw *string

	// ContactNamen.
	ContactNamen *string

	// ContactNameNic.
	ContactNameNic *string

	// ContactNameNie.
	ContactNameNie *string

	// ContactNameNiew.
	ContactNameNiew *string

	// ContactNameNisw.
	ContactNameNisw *string

	// ContactPhone.
	ContactPhone *string

	// ContactPhoneIc.
	ContactPhoneIc *string

	// ContactPhoneIe.
	ContactPhoneIe *string

	// ContactPhoneIew.
	ContactPhoneIew *string

	// ContactPhoneIsw.
	ContactPhoneIsw *string

	// ContactPhonen.
	ContactPhonen *string

	// ContactPhoneNic.
	ContactPhoneNic *string

	// ContactPhoneNie.
	ContactPhoneNie *string

	// ContactPhoneNiew.
	ContactPhoneNiew *string

	// ContactPhoneNisw.
	ContactPhoneNisw *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Facility.
	Facility *string

	// FacilityIc.
	FacilityIc *string

	// FacilityIe.
	FacilityIe *string

	// FacilityIew.
	FacilityIew *string

	// FacilityIsw.
	FacilityIsw *string

	// Facilityn.
	Facilityn *string

	// FacilityNic.
	FacilityNic *string

	// FacilityNie.
	FacilityNie *string

	// FacilityNiew.
	FacilityNiew *string

	// FacilityNisw.
	FacilityNisw *string

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

	// Latitude.
	Latitude *string

	// LatitudeGt.
	LatitudeGt *string

	// LatitudeGte.
	LatitudeGte *string

	// LatitudeLt.
	LatitudeLt *string

	// LatitudeLte.
	LatitudeLte *string

	// Latituden.
	Latituden *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Longitude.
	Longitude *string

	// LongitudeGt.
	LongitudeGt *string

	// LongitudeGte.
	LongitudeGte *string

	// LongitudeLt.
	LongitudeLt *string

	// LongitudeLte.
	LongitudeLte *string

	// Longituden.
	Longituden *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim sites list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesListParams) WithDefaults() *DcimSitesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim sites list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimSitesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim sites list params
func (o *DcimSitesListParams) WithTimeout(timeout time.Duration) *DcimSitesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim sites list params
func (o *DcimSitesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim sites list params
func (o *DcimSitesListParams) WithContext(ctx context.Context) *DcimSitesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim sites list params
func (o *DcimSitesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim sites list params
func (o *DcimSitesListParams) WithHTTPClient(client *http.Client) *DcimSitesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim sites list params
func (o *DcimSitesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAsn adds the asn to the dcim sites list params
func (o *DcimSitesListParams) WithAsn(asn *string) *DcimSitesListParams {
	o.SetAsn(asn)
	return o
}

// SetAsn adds the asn to the dcim sites list params
func (o *DcimSitesListParams) SetAsn(asn *string) {
	o.Asn = asn
}

// WithAsnGt adds the asnGt to the dcim sites list params
func (o *DcimSitesListParams) WithAsnGt(asnGt *string) *DcimSitesListParams {
	o.SetAsnGt(asnGt)
	return o
}

// SetAsnGt adds the asnGt to the dcim sites list params
func (o *DcimSitesListParams) SetAsnGt(asnGt *string) {
	o.AsnGt = asnGt
}

// WithAsnGte adds the asnGte to the dcim sites list params
func (o *DcimSitesListParams) WithAsnGte(asnGte *string) *DcimSitesListParams {
	o.SetAsnGte(asnGte)
	return o
}

// SetAsnGte adds the asnGte to the dcim sites list params
func (o *DcimSitesListParams) SetAsnGte(asnGte *string) {
	o.AsnGte = asnGte
}

// WithAsnLt adds the asnLt to the dcim sites list params
func (o *DcimSitesListParams) WithAsnLt(asnLt *string) *DcimSitesListParams {
	o.SetAsnLt(asnLt)
	return o
}

// SetAsnLt adds the asnLt to the dcim sites list params
func (o *DcimSitesListParams) SetAsnLt(asnLt *string) {
	o.AsnLt = asnLt
}

// WithAsnLte adds the asnLte to the dcim sites list params
func (o *DcimSitesListParams) WithAsnLte(asnLte *string) *DcimSitesListParams {
	o.SetAsnLte(asnLte)
	return o
}

// SetAsnLte adds the asnLte to the dcim sites list params
func (o *DcimSitesListParams) SetAsnLte(asnLte *string) {
	o.AsnLte = asnLte
}

// WithAsnn adds the asnn to the dcim sites list params
func (o *DcimSitesListParams) WithAsnn(asnn *string) *DcimSitesListParams {
	o.SetAsnn(asnn)
	return o
}

// SetAsnn adds the asnN to the dcim sites list params
func (o *DcimSitesListParams) SetAsnn(asnn *string) {
	o.Asnn = asnn
}

// WithContactEmail adds the contactEmail to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmail(contactEmail *string) *DcimSitesListParams {
	o.SetContactEmail(contactEmail)
	return o
}

// SetContactEmail adds the contactEmail to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmail(contactEmail *string) {
	o.ContactEmail = contactEmail
}

// WithContactEmailIc adds the contactEmailIc to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailIc(contactEmailIc *string) *DcimSitesListParams {
	o.SetContactEmailIc(contactEmailIc)
	return o
}

// SetContactEmailIc adds the contactEmailIc to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailIc(contactEmailIc *string) {
	o.ContactEmailIc = contactEmailIc
}

// WithContactEmailIe adds the contactEmailIe to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailIe(contactEmailIe *string) *DcimSitesListParams {
	o.SetContactEmailIe(contactEmailIe)
	return o
}

// SetContactEmailIe adds the contactEmailIe to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailIe(contactEmailIe *string) {
	o.ContactEmailIe = contactEmailIe
}

// WithContactEmailIew adds the contactEmailIew to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailIew(contactEmailIew *string) *DcimSitesListParams {
	o.SetContactEmailIew(contactEmailIew)
	return o
}

// SetContactEmailIew adds the contactEmailIew to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailIew(contactEmailIew *string) {
	o.ContactEmailIew = contactEmailIew
}

// WithContactEmailIsw adds the contactEmailIsw to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailIsw(contactEmailIsw *string) *DcimSitesListParams {
	o.SetContactEmailIsw(contactEmailIsw)
	return o
}

// SetContactEmailIsw adds the contactEmailIsw to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailIsw(contactEmailIsw *string) {
	o.ContactEmailIsw = contactEmailIsw
}

// WithContactEmailn adds the contactEmailn to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailn(contactEmailn *string) *DcimSitesListParams {
	o.SetContactEmailn(contactEmailn)
	return o
}

// SetContactEmailn adds the contactEmailN to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailn(contactEmailn *string) {
	o.ContactEmailn = contactEmailn
}

// WithContactEmailNic adds the contactEmailNic to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailNic(contactEmailNic *string) *DcimSitesListParams {
	o.SetContactEmailNic(contactEmailNic)
	return o
}

// SetContactEmailNic adds the contactEmailNic to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailNic(contactEmailNic *string) {
	o.ContactEmailNic = contactEmailNic
}

// WithContactEmailNie adds the contactEmailNie to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailNie(contactEmailNie *string) *DcimSitesListParams {
	o.SetContactEmailNie(contactEmailNie)
	return o
}

// SetContactEmailNie adds the contactEmailNie to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailNie(contactEmailNie *string) {
	o.ContactEmailNie = contactEmailNie
}

// WithContactEmailNiew adds the contactEmailNiew to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailNiew(contactEmailNiew *string) *DcimSitesListParams {
	o.SetContactEmailNiew(contactEmailNiew)
	return o
}

// SetContactEmailNiew adds the contactEmailNiew to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailNiew(contactEmailNiew *string) {
	o.ContactEmailNiew = contactEmailNiew
}

// WithContactEmailNisw adds the contactEmailNisw to the dcim sites list params
func (o *DcimSitesListParams) WithContactEmailNisw(contactEmailNisw *string) *DcimSitesListParams {
	o.SetContactEmailNisw(contactEmailNisw)
	return o
}

// SetContactEmailNisw adds the contactEmailNisw to the dcim sites list params
func (o *DcimSitesListParams) SetContactEmailNisw(contactEmailNisw *string) {
	o.ContactEmailNisw = contactEmailNisw
}

// WithContactName adds the contactName to the dcim sites list params
func (o *DcimSitesListParams) WithContactName(contactName *string) *DcimSitesListParams {
	o.SetContactName(contactName)
	return o
}

// SetContactName adds the contactName to the dcim sites list params
func (o *DcimSitesListParams) SetContactName(contactName *string) {
	o.ContactName = contactName
}

// WithContactNameIc adds the contactNameIc to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameIc(contactNameIc *string) *DcimSitesListParams {
	o.SetContactNameIc(contactNameIc)
	return o
}

// SetContactNameIc adds the contactNameIc to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameIc(contactNameIc *string) {
	o.ContactNameIc = contactNameIc
}

// WithContactNameIe adds the contactNameIe to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameIe(contactNameIe *string) *DcimSitesListParams {
	o.SetContactNameIe(contactNameIe)
	return o
}

// SetContactNameIe adds the contactNameIe to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameIe(contactNameIe *string) {
	o.ContactNameIe = contactNameIe
}

// WithContactNameIew adds the contactNameIew to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameIew(contactNameIew *string) *DcimSitesListParams {
	o.SetContactNameIew(contactNameIew)
	return o
}

// SetContactNameIew adds the contactNameIew to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameIew(contactNameIew *string) {
	o.ContactNameIew = contactNameIew
}

// WithContactNameIsw adds the contactNameIsw to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameIsw(contactNameIsw *string) *DcimSitesListParams {
	o.SetContactNameIsw(contactNameIsw)
	return o
}

// SetContactNameIsw adds the contactNameIsw to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameIsw(contactNameIsw *string) {
	o.ContactNameIsw = contactNameIsw
}

// WithContactNamen adds the contactNamen to the dcim sites list params
func (o *DcimSitesListParams) WithContactNamen(contactNamen *string) *DcimSitesListParams {
	o.SetContactNamen(contactNamen)
	return o
}

// SetContactNamen adds the contactNameN to the dcim sites list params
func (o *DcimSitesListParams) SetContactNamen(contactNamen *string) {
	o.ContactNamen = contactNamen
}

// WithContactNameNic adds the contactNameNic to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameNic(contactNameNic *string) *DcimSitesListParams {
	o.SetContactNameNic(contactNameNic)
	return o
}

// SetContactNameNic adds the contactNameNic to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameNic(contactNameNic *string) {
	o.ContactNameNic = contactNameNic
}

// WithContactNameNie adds the contactNameNie to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameNie(contactNameNie *string) *DcimSitesListParams {
	o.SetContactNameNie(contactNameNie)
	return o
}

// SetContactNameNie adds the contactNameNie to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameNie(contactNameNie *string) {
	o.ContactNameNie = contactNameNie
}

// WithContactNameNiew adds the contactNameNiew to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameNiew(contactNameNiew *string) *DcimSitesListParams {
	o.SetContactNameNiew(contactNameNiew)
	return o
}

// SetContactNameNiew adds the contactNameNiew to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameNiew(contactNameNiew *string) {
	o.ContactNameNiew = contactNameNiew
}

// WithContactNameNisw adds the contactNameNisw to the dcim sites list params
func (o *DcimSitesListParams) WithContactNameNisw(contactNameNisw *string) *DcimSitesListParams {
	o.SetContactNameNisw(contactNameNisw)
	return o
}

// SetContactNameNisw adds the contactNameNisw to the dcim sites list params
func (o *DcimSitesListParams) SetContactNameNisw(contactNameNisw *string) {
	o.ContactNameNisw = contactNameNisw
}

// WithContactPhone adds the contactPhone to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhone(contactPhone *string) *DcimSitesListParams {
	o.SetContactPhone(contactPhone)
	return o
}

// SetContactPhone adds the contactPhone to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhone(contactPhone *string) {
	o.ContactPhone = contactPhone
}

// WithContactPhoneIc adds the contactPhoneIc to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneIc(contactPhoneIc *string) *DcimSitesListParams {
	o.SetContactPhoneIc(contactPhoneIc)
	return o
}

// SetContactPhoneIc adds the contactPhoneIc to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneIc(contactPhoneIc *string) {
	o.ContactPhoneIc = contactPhoneIc
}

// WithContactPhoneIe adds the contactPhoneIe to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneIe(contactPhoneIe *string) *DcimSitesListParams {
	o.SetContactPhoneIe(contactPhoneIe)
	return o
}

// SetContactPhoneIe adds the contactPhoneIe to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneIe(contactPhoneIe *string) {
	o.ContactPhoneIe = contactPhoneIe
}

// WithContactPhoneIew adds the contactPhoneIew to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneIew(contactPhoneIew *string) *DcimSitesListParams {
	o.SetContactPhoneIew(contactPhoneIew)
	return o
}

// SetContactPhoneIew adds the contactPhoneIew to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneIew(contactPhoneIew *string) {
	o.ContactPhoneIew = contactPhoneIew
}

// WithContactPhoneIsw adds the contactPhoneIsw to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneIsw(contactPhoneIsw *string) *DcimSitesListParams {
	o.SetContactPhoneIsw(contactPhoneIsw)
	return o
}

// SetContactPhoneIsw adds the contactPhoneIsw to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneIsw(contactPhoneIsw *string) {
	o.ContactPhoneIsw = contactPhoneIsw
}

// WithContactPhonen adds the contactPhonen to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhonen(contactPhonen *string) *DcimSitesListParams {
	o.SetContactPhonen(contactPhonen)
	return o
}

// SetContactPhonen adds the contactPhoneN to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhonen(contactPhonen *string) {
	o.ContactPhonen = contactPhonen
}

// WithContactPhoneNic adds the contactPhoneNic to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneNic(contactPhoneNic *string) *DcimSitesListParams {
	o.SetContactPhoneNic(contactPhoneNic)
	return o
}

// SetContactPhoneNic adds the contactPhoneNic to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneNic(contactPhoneNic *string) {
	o.ContactPhoneNic = contactPhoneNic
}

// WithContactPhoneNie adds the contactPhoneNie to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneNie(contactPhoneNie *string) *DcimSitesListParams {
	o.SetContactPhoneNie(contactPhoneNie)
	return o
}

// SetContactPhoneNie adds the contactPhoneNie to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneNie(contactPhoneNie *string) {
	o.ContactPhoneNie = contactPhoneNie
}

// WithContactPhoneNiew adds the contactPhoneNiew to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneNiew(contactPhoneNiew *string) *DcimSitesListParams {
	o.SetContactPhoneNiew(contactPhoneNiew)
	return o
}

// SetContactPhoneNiew adds the contactPhoneNiew to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneNiew(contactPhoneNiew *string) {
	o.ContactPhoneNiew = contactPhoneNiew
}

// WithContactPhoneNisw adds the contactPhoneNisw to the dcim sites list params
func (o *DcimSitesListParams) WithContactPhoneNisw(contactPhoneNisw *string) *DcimSitesListParams {
	o.SetContactPhoneNisw(contactPhoneNisw)
	return o
}

// SetContactPhoneNisw adds the contactPhoneNisw to the dcim sites list params
func (o *DcimSitesListParams) SetContactPhoneNisw(contactPhoneNisw *string) {
	o.ContactPhoneNisw = contactPhoneNisw
}

// WithCreated adds the created to the dcim sites list params
func (o *DcimSitesListParams) WithCreated(created *string) *DcimSitesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim sites list params
func (o *DcimSitesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim sites list params
func (o *DcimSitesListParams) WithCreatedGte(createdGte *string) *DcimSitesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim sites list params
func (o *DcimSitesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim sites list params
func (o *DcimSitesListParams) WithCreatedLte(createdLte *string) *DcimSitesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim sites list params
func (o *DcimSitesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithFacility adds the facility to the dcim sites list params
func (o *DcimSitesListParams) WithFacility(facility *string) *DcimSitesListParams {
	o.SetFacility(facility)
	return o
}

// SetFacility adds the facility to the dcim sites list params
func (o *DcimSitesListParams) SetFacility(facility *string) {
	o.Facility = facility
}

// WithFacilityIc adds the facilityIc to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityIc(facilityIc *string) *DcimSitesListParams {
	o.SetFacilityIc(facilityIc)
	return o
}

// SetFacilityIc adds the facilityIc to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityIc(facilityIc *string) {
	o.FacilityIc = facilityIc
}

// WithFacilityIe adds the facilityIe to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityIe(facilityIe *string) *DcimSitesListParams {
	o.SetFacilityIe(facilityIe)
	return o
}

// SetFacilityIe adds the facilityIe to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityIe(facilityIe *string) {
	o.FacilityIe = facilityIe
}

// WithFacilityIew adds the facilityIew to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityIew(facilityIew *string) *DcimSitesListParams {
	o.SetFacilityIew(facilityIew)
	return o
}

// SetFacilityIew adds the facilityIew to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityIew(facilityIew *string) {
	o.FacilityIew = facilityIew
}

// WithFacilityIsw adds the facilityIsw to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityIsw(facilityIsw *string) *DcimSitesListParams {
	o.SetFacilityIsw(facilityIsw)
	return o
}

// SetFacilityIsw adds the facilityIsw to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityIsw(facilityIsw *string) {
	o.FacilityIsw = facilityIsw
}

// WithFacilityn adds the facilityn to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityn(facilityn *string) *DcimSitesListParams {
	o.SetFacilityn(facilityn)
	return o
}

// SetFacilityn adds the facilityN to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityn(facilityn *string) {
	o.Facilityn = facilityn
}

// WithFacilityNic adds the facilityNic to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityNic(facilityNic *string) *DcimSitesListParams {
	o.SetFacilityNic(facilityNic)
	return o
}

// SetFacilityNic adds the facilityNic to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityNic(facilityNic *string) {
	o.FacilityNic = facilityNic
}

// WithFacilityNie adds the facilityNie to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityNie(facilityNie *string) *DcimSitesListParams {
	o.SetFacilityNie(facilityNie)
	return o
}

// SetFacilityNie adds the facilityNie to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityNie(facilityNie *string) {
	o.FacilityNie = facilityNie
}

// WithFacilityNiew adds the facilityNiew to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityNiew(facilityNiew *string) *DcimSitesListParams {
	o.SetFacilityNiew(facilityNiew)
	return o
}

// SetFacilityNiew adds the facilityNiew to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityNiew(facilityNiew *string) {
	o.FacilityNiew = facilityNiew
}

// WithFacilityNisw adds the facilityNisw to the dcim sites list params
func (o *DcimSitesListParams) WithFacilityNisw(facilityNisw *string) *DcimSitesListParams {
	o.SetFacilityNisw(facilityNisw)
	return o
}

// SetFacilityNisw adds the facilityNisw to the dcim sites list params
func (o *DcimSitesListParams) SetFacilityNisw(facilityNisw *string) {
	o.FacilityNisw = facilityNisw
}

// WithID adds the id to the dcim sites list params
func (o *DcimSitesListParams) WithID(id *string) *DcimSitesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim sites list params
func (o *DcimSitesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim sites list params
func (o *DcimSitesListParams) WithIDIc(iDIc *string) *DcimSitesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim sites list params
func (o *DcimSitesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim sites list params
func (o *DcimSitesListParams) WithIDIe(iDIe *string) *DcimSitesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim sites list params
func (o *DcimSitesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim sites list params
func (o *DcimSitesListParams) WithIDIew(iDIew *string) *DcimSitesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim sites list params
func (o *DcimSitesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim sites list params
func (o *DcimSitesListParams) WithIDIsw(iDIsw *string) *DcimSitesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim sites list params
func (o *DcimSitesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim sites list params
func (o *DcimSitesListParams) WithIDn(iDn *string) *DcimSitesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim sites list params
func (o *DcimSitesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim sites list params
func (o *DcimSitesListParams) WithIDNic(iDNic *string) *DcimSitesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim sites list params
func (o *DcimSitesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim sites list params
func (o *DcimSitesListParams) WithIDNie(iDNie *string) *DcimSitesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim sites list params
func (o *DcimSitesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim sites list params
func (o *DcimSitesListParams) WithIDNiew(iDNiew *string) *DcimSitesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim sites list params
func (o *DcimSitesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim sites list params
func (o *DcimSitesListParams) WithIDNisw(iDNisw *string) *DcimSitesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim sites list params
func (o *DcimSitesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the dcim sites list params
func (o *DcimSitesListParams) WithLastUpdated(lastUpdated *string) *DcimSitesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim sites list params
func (o *DcimSitesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim sites list params
func (o *DcimSitesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimSitesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim sites list params
func (o *DcimSitesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim sites list params
func (o *DcimSitesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimSitesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim sites list params
func (o *DcimSitesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLatitude adds the latitude to the dcim sites list params
func (o *DcimSitesListParams) WithLatitude(latitude *string) *DcimSitesListParams {
	o.SetLatitude(latitude)
	return o
}

// SetLatitude adds the latitude to the dcim sites list params
func (o *DcimSitesListParams) SetLatitude(latitude *string) {
	o.Latitude = latitude
}

// WithLatitudeGt adds the latitudeGt to the dcim sites list params
func (o *DcimSitesListParams) WithLatitudeGt(latitudeGt *string) *DcimSitesListParams {
	o.SetLatitudeGt(latitudeGt)
	return o
}

// SetLatitudeGt adds the latitudeGt to the dcim sites list params
func (o *DcimSitesListParams) SetLatitudeGt(latitudeGt *string) {
	o.LatitudeGt = latitudeGt
}

// WithLatitudeGte adds the latitudeGte to the dcim sites list params
func (o *DcimSitesListParams) WithLatitudeGte(latitudeGte *string) *DcimSitesListParams {
	o.SetLatitudeGte(latitudeGte)
	return o
}

// SetLatitudeGte adds the latitudeGte to the dcim sites list params
func (o *DcimSitesListParams) SetLatitudeGte(latitudeGte *string) {
	o.LatitudeGte = latitudeGte
}

// WithLatitudeLt adds the latitudeLt to the dcim sites list params
func (o *DcimSitesListParams) WithLatitudeLt(latitudeLt *string) *DcimSitesListParams {
	o.SetLatitudeLt(latitudeLt)
	return o
}

// SetLatitudeLt adds the latitudeLt to the dcim sites list params
func (o *DcimSitesListParams) SetLatitudeLt(latitudeLt *string) {
	o.LatitudeLt = latitudeLt
}

// WithLatitudeLte adds the latitudeLte to the dcim sites list params
func (o *DcimSitesListParams) WithLatitudeLte(latitudeLte *string) *DcimSitesListParams {
	o.SetLatitudeLte(latitudeLte)
	return o
}

// SetLatitudeLte adds the latitudeLte to the dcim sites list params
func (o *DcimSitesListParams) SetLatitudeLte(latitudeLte *string) {
	o.LatitudeLte = latitudeLte
}

// WithLatituden adds the latituden to the dcim sites list params
func (o *DcimSitesListParams) WithLatituden(latituden *string) *DcimSitesListParams {
	o.SetLatituden(latituden)
	return o
}

// SetLatituden adds the latitudeN to the dcim sites list params
func (o *DcimSitesListParams) SetLatituden(latituden *string) {
	o.Latituden = latituden
}

// WithLimit adds the limit to the dcim sites list params
func (o *DcimSitesListParams) WithLimit(limit *int64) *DcimSitesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim sites list params
func (o *DcimSitesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithLongitude adds the longitude to the dcim sites list params
func (o *DcimSitesListParams) WithLongitude(longitude *string) *DcimSitesListParams {
	o.SetLongitude(longitude)
	return o
}

// SetLongitude adds the longitude to the dcim sites list params
func (o *DcimSitesListParams) SetLongitude(longitude *string) {
	o.Longitude = longitude
}

// WithLongitudeGt adds the longitudeGt to the dcim sites list params
func (o *DcimSitesListParams) WithLongitudeGt(longitudeGt *string) *DcimSitesListParams {
	o.SetLongitudeGt(longitudeGt)
	return o
}

// SetLongitudeGt adds the longitudeGt to the dcim sites list params
func (o *DcimSitesListParams) SetLongitudeGt(longitudeGt *string) {
	o.LongitudeGt = longitudeGt
}

// WithLongitudeGte adds the longitudeGte to the dcim sites list params
func (o *DcimSitesListParams) WithLongitudeGte(longitudeGte *string) *DcimSitesListParams {
	o.SetLongitudeGte(longitudeGte)
	return o
}

// SetLongitudeGte adds the longitudeGte to the dcim sites list params
func (o *DcimSitesListParams) SetLongitudeGte(longitudeGte *string) {
	o.LongitudeGte = longitudeGte
}

// WithLongitudeLt adds the longitudeLt to the dcim sites list params
func (o *DcimSitesListParams) WithLongitudeLt(longitudeLt *string) *DcimSitesListParams {
	o.SetLongitudeLt(longitudeLt)
	return o
}

// SetLongitudeLt adds the longitudeLt to the dcim sites list params
func (o *DcimSitesListParams) SetLongitudeLt(longitudeLt *string) {
	o.LongitudeLt = longitudeLt
}

// WithLongitudeLte adds the longitudeLte to the dcim sites list params
func (o *DcimSitesListParams) WithLongitudeLte(longitudeLte *string) *DcimSitesListParams {
	o.SetLongitudeLte(longitudeLte)
	return o
}

// SetLongitudeLte adds the longitudeLte to the dcim sites list params
func (o *DcimSitesListParams) SetLongitudeLte(longitudeLte *string) {
	o.LongitudeLte = longitudeLte
}

// WithLongituden adds the longituden to the dcim sites list params
func (o *DcimSitesListParams) WithLongituden(longituden *string) *DcimSitesListParams {
	o.SetLongituden(longituden)
	return o
}

// SetLongituden adds the longitudeN to the dcim sites list params
func (o *DcimSitesListParams) SetLongituden(longituden *string) {
	o.Longituden = longituden
}

// WithName adds the name to the dcim sites list params
func (o *DcimSitesListParams) WithName(name *string) *DcimSitesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim sites list params
func (o *DcimSitesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim sites list params
func (o *DcimSitesListParams) WithNameIc(nameIc *string) *DcimSitesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim sites list params
func (o *DcimSitesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim sites list params
func (o *DcimSitesListParams) WithNameIe(nameIe *string) *DcimSitesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim sites list params
func (o *DcimSitesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim sites list params
func (o *DcimSitesListParams) WithNameIew(nameIew *string) *DcimSitesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim sites list params
func (o *DcimSitesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim sites list params
func (o *DcimSitesListParams) WithNameIsw(nameIsw *string) *DcimSitesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim sites list params
func (o *DcimSitesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim sites list params
func (o *DcimSitesListParams) WithNamen(namen *string) *DcimSitesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim sites list params
func (o *DcimSitesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim sites list params
func (o *DcimSitesListParams) WithNameNic(nameNic *string) *DcimSitesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim sites list params
func (o *DcimSitesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim sites list params
func (o *DcimSitesListParams) WithNameNie(nameNie *string) *DcimSitesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim sites list params
func (o *DcimSitesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim sites list params
func (o *DcimSitesListParams) WithNameNiew(nameNiew *string) *DcimSitesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim sites list params
func (o *DcimSitesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim sites list params
func (o *DcimSitesListParams) WithNameNisw(nameNisw *string) *DcimSitesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim sites list params
func (o *DcimSitesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim sites list params
func (o *DcimSitesListParams) WithOffset(offset *int64) *DcimSitesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim sites list params
func (o *DcimSitesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim sites list params
func (o *DcimSitesListParams) WithQ(q *string) *DcimSitesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim sites list params
func (o *DcimSitesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRegion adds the region to the dcim sites list params
func (o *DcimSitesListParams) WithRegion(region *string) *DcimSitesListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim sites list params
func (o *DcimSitesListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim sites list params
func (o *DcimSitesListParams) WithRegionn(regionn *string) *DcimSitesListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim sites list params
func (o *DcimSitesListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim sites list params
func (o *DcimSitesListParams) WithRegionID(regionID *string) *DcimSitesListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim sites list params
func (o *DcimSitesListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim sites list params
func (o *DcimSitesListParams) WithRegionIDn(regionIDn *string) *DcimSitesListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim sites list params
func (o *DcimSitesListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSlug adds the slug to the dcim sites list params
func (o *DcimSitesListParams) WithSlug(slug *string) *DcimSitesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim sites list params
func (o *DcimSitesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the dcim sites list params
func (o *DcimSitesListParams) WithSlugIc(slugIc *string) *DcimSitesListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the dcim sites list params
func (o *DcimSitesListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the dcim sites list params
func (o *DcimSitesListParams) WithSlugIe(slugIe *string) *DcimSitesListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the dcim sites list params
func (o *DcimSitesListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the dcim sites list params
func (o *DcimSitesListParams) WithSlugIew(slugIew *string) *DcimSitesListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the dcim sites list params
func (o *DcimSitesListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the dcim sites list params
func (o *DcimSitesListParams) WithSlugIsw(slugIsw *string) *DcimSitesListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the dcim sites list params
func (o *DcimSitesListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the dcim sites list params
func (o *DcimSitesListParams) WithSlugn(slugn *string) *DcimSitesListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the dcim sites list params
func (o *DcimSitesListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the dcim sites list params
func (o *DcimSitesListParams) WithSlugNic(slugNic *string) *DcimSitesListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the dcim sites list params
func (o *DcimSitesListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the dcim sites list params
func (o *DcimSitesListParams) WithSlugNie(slugNie *string) *DcimSitesListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the dcim sites list params
func (o *DcimSitesListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the dcim sites list params
func (o *DcimSitesListParams) WithSlugNiew(slugNiew *string) *DcimSitesListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the dcim sites list params
func (o *DcimSitesListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the dcim sites list params
func (o *DcimSitesListParams) WithSlugNisw(slugNisw *string) *DcimSitesListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the dcim sites list params
func (o *DcimSitesListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WithStatus adds the status to the dcim sites list params
func (o *DcimSitesListParams) WithStatus(status *string) *DcimSitesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim sites list params
func (o *DcimSitesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the dcim sites list params
func (o *DcimSitesListParams) WithStatusn(statusn *string) *DcimSitesListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the dcim sites list params
func (o *DcimSitesListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the dcim sites list params
func (o *DcimSitesListParams) WithTag(tag *string) *DcimSitesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim sites list params
func (o *DcimSitesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim sites list params
func (o *DcimSitesListParams) WithTagn(tagn *string) *DcimSitesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim sites list params
func (o *DcimSitesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the dcim sites list params
func (o *DcimSitesListParams) WithTenant(tenant *string) *DcimSitesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the dcim sites list params
func (o *DcimSitesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the dcim sites list params
func (o *DcimSitesListParams) WithTenantn(tenantn *string) *DcimSitesListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the dcim sites list params
func (o *DcimSitesListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the dcim sites list params
func (o *DcimSitesListParams) WithTenantGroup(tenantGroup *string) *DcimSitesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the dcim sites list params
func (o *DcimSitesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the dcim sites list params
func (o *DcimSitesListParams) WithTenantGroupn(tenantGroupn *string) *DcimSitesListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the dcim sites list params
func (o *DcimSitesListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the dcim sites list params
func (o *DcimSitesListParams) WithTenantGroupID(tenantGroupID *string) *DcimSitesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the dcim sites list params
func (o *DcimSitesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the dcim sites list params
func (o *DcimSitesListParams) WithTenantGroupIDn(tenantGroupIDn *string) *DcimSitesListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the dcim sites list params
func (o *DcimSitesListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the dcim sites list params
func (o *DcimSitesListParams) WithTenantID(tenantID *string) *DcimSitesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the dcim sites list params
func (o *DcimSitesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the dcim sites list params
func (o *DcimSitesListParams) WithTenantIDn(tenantIDn *string) *DcimSitesListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the dcim sites list params
func (o *DcimSitesListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimSitesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ContactEmail != nil {

		// query param contact_email
		var qrContactEmail string

		if o.ContactEmail != nil {
			qrContactEmail = *o.ContactEmail
		}
		qContactEmail := qrContactEmail
		if qContactEmail != "" {

			if err := r.SetQueryParam("contact_email", qContactEmail); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailIc != nil {

		// query param contact_email__ic
		var qrContactEmailIc string

		if o.ContactEmailIc != nil {
			qrContactEmailIc = *o.ContactEmailIc
		}
		qContactEmailIc := qrContactEmailIc
		if qContactEmailIc != "" {

			if err := r.SetQueryParam("contact_email__ic", qContactEmailIc); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailIe != nil {

		// query param contact_email__ie
		var qrContactEmailIe string

		if o.ContactEmailIe != nil {
			qrContactEmailIe = *o.ContactEmailIe
		}
		qContactEmailIe := qrContactEmailIe
		if qContactEmailIe != "" {

			if err := r.SetQueryParam("contact_email__ie", qContactEmailIe); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailIew != nil {

		// query param contact_email__iew
		var qrContactEmailIew string

		if o.ContactEmailIew != nil {
			qrContactEmailIew = *o.ContactEmailIew
		}
		qContactEmailIew := qrContactEmailIew
		if qContactEmailIew != "" {

			if err := r.SetQueryParam("contact_email__iew", qContactEmailIew); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailIsw != nil {

		// query param contact_email__isw
		var qrContactEmailIsw string

		if o.ContactEmailIsw != nil {
			qrContactEmailIsw = *o.ContactEmailIsw
		}
		qContactEmailIsw := qrContactEmailIsw
		if qContactEmailIsw != "" {

			if err := r.SetQueryParam("contact_email__isw", qContactEmailIsw); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailn != nil {

		// query param contact_email__n
		var qrContactEmailn string

		if o.ContactEmailn != nil {
			qrContactEmailn = *o.ContactEmailn
		}
		qContactEmailn := qrContactEmailn
		if qContactEmailn != "" {

			if err := r.SetQueryParam("contact_email__n", qContactEmailn); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailNic != nil {

		// query param contact_email__nic
		var qrContactEmailNic string

		if o.ContactEmailNic != nil {
			qrContactEmailNic = *o.ContactEmailNic
		}
		qContactEmailNic := qrContactEmailNic
		if qContactEmailNic != "" {

			if err := r.SetQueryParam("contact_email__nic", qContactEmailNic); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailNie != nil {

		// query param contact_email__nie
		var qrContactEmailNie string

		if o.ContactEmailNie != nil {
			qrContactEmailNie = *o.ContactEmailNie
		}
		qContactEmailNie := qrContactEmailNie
		if qContactEmailNie != "" {

			if err := r.SetQueryParam("contact_email__nie", qContactEmailNie); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailNiew != nil {

		// query param contact_email__niew
		var qrContactEmailNiew string

		if o.ContactEmailNiew != nil {
			qrContactEmailNiew = *o.ContactEmailNiew
		}
		qContactEmailNiew := qrContactEmailNiew
		if qContactEmailNiew != "" {

			if err := r.SetQueryParam("contact_email__niew", qContactEmailNiew); err != nil {
				return err
			}
		}
	}

	if o.ContactEmailNisw != nil {

		// query param contact_email__nisw
		var qrContactEmailNisw string

		if o.ContactEmailNisw != nil {
			qrContactEmailNisw = *o.ContactEmailNisw
		}
		qContactEmailNisw := qrContactEmailNisw
		if qContactEmailNisw != "" {

			if err := r.SetQueryParam("contact_email__nisw", qContactEmailNisw); err != nil {
				return err
			}
		}
	}

	if o.ContactName != nil {

		// query param contact_name
		var qrContactName string

		if o.ContactName != nil {
			qrContactName = *o.ContactName
		}
		qContactName := qrContactName
		if qContactName != "" {

			if err := r.SetQueryParam("contact_name", qContactName); err != nil {
				return err
			}
		}
	}

	if o.ContactNameIc != nil {

		// query param contact_name__ic
		var qrContactNameIc string

		if o.ContactNameIc != nil {
			qrContactNameIc = *o.ContactNameIc
		}
		qContactNameIc := qrContactNameIc
		if qContactNameIc != "" {

			if err := r.SetQueryParam("contact_name__ic", qContactNameIc); err != nil {
				return err
			}
		}
	}

	if o.ContactNameIe != nil {

		// query param contact_name__ie
		var qrContactNameIe string

		if o.ContactNameIe != nil {
			qrContactNameIe = *o.ContactNameIe
		}
		qContactNameIe := qrContactNameIe
		if qContactNameIe != "" {

			if err := r.SetQueryParam("contact_name__ie", qContactNameIe); err != nil {
				return err
			}
		}
	}

	if o.ContactNameIew != nil {

		// query param contact_name__iew
		var qrContactNameIew string

		if o.ContactNameIew != nil {
			qrContactNameIew = *o.ContactNameIew
		}
		qContactNameIew := qrContactNameIew
		if qContactNameIew != "" {

			if err := r.SetQueryParam("contact_name__iew", qContactNameIew); err != nil {
				return err
			}
		}
	}

	if o.ContactNameIsw != nil {

		// query param contact_name__isw
		var qrContactNameIsw string

		if o.ContactNameIsw != nil {
			qrContactNameIsw = *o.ContactNameIsw
		}
		qContactNameIsw := qrContactNameIsw
		if qContactNameIsw != "" {

			if err := r.SetQueryParam("contact_name__isw", qContactNameIsw); err != nil {
				return err
			}
		}
	}

	if o.ContactNamen != nil {

		// query param contact_name__n
		var qrContactNamen string

		if o.ContactNamen != nil {
			qrContactNamen = *o.ContactNamen
		}
		qContactNamen := qrContactNamen
		if qContactNamen != "" {

			if err := r.SetQueryParam("contact_name__n", qContactNamen); err != nil {
				return err
			}
		}
	}

	if o.ContactNameNic != nil {

		// query param contact_name__nic
		var qrContactNameNic string

		if o.ContactNameNic != nil {
			qrContactNameNic = *o.ContactNameNic
		}
		qContactNameNic := qrContactNameNic
		if qContactNameNic != "" {

			if err := r.SetQueryParam("contact_name__nic", qContactNameNic); err != nil {
				return err
			}
		}
	}

	if o.ContactNameNie != nil {

		// query param contact_name__nie
		var qrContactNameNie string

		if o.ContactNameNie != nil {
			qrContactNameNie = *o.ContactNameNie
		}
		qContactNameNie := qrContactNameNie
		if qContactNameNie != "" {

			if err := r.SetQueryParam("contact_name__nie", qContactNameNie); err != nil {
				return err
			}
		}
	}

	if o.ContactNameNiew != nil {

		// query param contact_name__niew
		var qrContactNameNiew string

		if o.ContactNameNiew != nil {
			qrContactNameNiew = *o.ContactNameNiew
		}
		qContactNameNiew := qrContactNameNiew
		if qContactNameNiew != "" {

			if err := r.SetQueryParam("contact_name__niew", qContactNameNiew); err != nil {
				return err
			}
		}
	}

	if o.ContactNameNisw != nil {

		// query param contact_name__nisw
		var qrContactNameNisw string

		if o.ContactNameNisw != nil {
			qrContactNameNisw = *o.ContactNameNisw
		}
		qContactNameNisw := qrContactNameNisw
		if qContactNameNisw != "" {

			if err := r.SetQueryParam("contact_name__nisw", qContactNameNisw); err != nil {
				return err
			}
		}
	}

	if o.ContactPhone != nil {

		// query param contact_phone
		var qrContactPhone string

		if o.ContactPhone != nil {
			qrContactPhone = *o.ContactPhone
		}
		qContactPhone := qrContactPhone
		if qContactPhone != "" {

			if err := r.SetQueryParam("contact_phone", qContactPhone); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneIc != nil {

		// query param contact_phone__ic
		var qrContactPhoneIc string

		if o.ContactPhoneIc != nil {
			qrContactPhoneIc = *o.ContactPhoneIc
		}
		qContactPhoneIc := qrContactPhoneIc
		if qContactPhoneIc != "" {

			if err := r.SetQueryParam("contact_phone__ic", qContactPhoneIc); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneIe != nil {

		// query param contact_phone__ie
		var qrContactPhoneIe string

		if o.ContactPhoneIe != nil {
			qrContactPhoneIe = *o.ContactPhoneIe
		}
		qContactPhoneIe := qrContactPhoneIe
		if qContactPhoneIe != "" {

			if err := r.SetQueryParam("contact_phone__ie", qContactPhoneIe); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneIew != nil {

		// query param contact_phone__iew
		var qrContactPhoneIew string

		if o.ContactPhoneIew != nil {
			qrContactPhoneIew = *o.ContactPhoneIew
		}
		qContactPhoneIew := qrContactPhoneIew
		if qContactPhoneIew != "" {

			if err := r.SetQueryParam("contact_phone__iew", qContactPhoneIew); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneIsw != nil {

		// query param contact_phone__isw
		var qrContactPhoneIsw string

		if o.ContactPhoneIsw != nil {
			qrContactPhoneIsw = *o.ContactPhoneIsw
		}
		qContactPhoneIsw := qrContactPhoneIsw
		if qContactPhoneIsw != "" {

			if err := r.SetQueryParam("contact_phone__isw", qContactPhoneIsw); err != nil {
				return err
			}
		}
	}

	if o.ContactPhonen != nil {

		// query param contact_phone__n
		var qrContactPhonen string

		if o.ContactPhonen != nil {
			qrContactPhonen = *o.ContactPhonen
		}
		qContactPhonen := qrContactPhonen
		if qContactPhonen != "" {

			if err := r.SetQueryParam("contact_phone__n", qContactPhonen); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneNic != nil {

		// query param contact_phone__nic
		var qrContactPhoneNic string

		if o.ContactPhoneNic != nil {
			qrContactPhoneNic = *o.ContactPhoneNic
		}
		qContactPhoneNic := qrContactPhoneNic
		if qContactPhoneNic != "" {

			if err := r.SetQueryParam("contact_phone__nic", qContactPhoneNic); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneNie != nil {

		// query param contact_phone__nie
		var qrContactPhoneNie string

		if o.ContactPhoneNie != nil {
			qrContactPhoneNie = *o.ContactPhoneNie
		}
		qContactPhoneNie := qrContactPhoneNie
		if qContactPhoneNie != "" {

			if err := r.SetQueryParam("contact_phone__nie", qContactPhoneNie); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneNiew != nil {

		// query param contact_phone__niew
		var qrContactPhoneNiew string

		if o.ContactPhoneNiew != nil {
			qrContactPhoneNiew = *o.ContactPhoneNiew
		}
		qContactPhoneNiew := qrContactPhoneNiew
		if qContactPhoneNiew != "" {

			if err := r.SetQueryParam("contact_phone__niew", qContactPhoneNiew); err != nil {
				return err
			}
		}
	}

	if o.ContactPhoneNisw != nil {

		// query param contact_phone__nisw
		var qrContactPhoneNisw string

		if o.ContactPhoneNisw != nil {
			qrContactPhoneNisw = *o.ContactPhoneNisw
		}
		qContactPhoneNisw := qrContactPhoneNisw
		if qContactPhoneNisw != "" {

			if err := r.SetQueryParam("contact_phone__nisw", qContactPhoneNisw); err != nil {
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

	if o.Facility != nil {

		// query param facility
		var qrFacility string

		if o.Facility != nil {
			qrFacility = *o.Facility
		}
		qFacility := qrFacility
		if qFacility != "" {

			if err := r.SetQueryParam("facility", qFacility); err != nil {
				return err
			}
		}
	}

	if o.FacilityIc != nil {

		// query param facility__ic
		var qrFacilityIc string

		if o.FacilityIc != nil {
			qrFacilityIc = *o.FacilityIc
		}
		qFacilityIc := qrFacilityIc
		if qFacilityIc != "" {

			if err := r.SetQueryParam("facility__ic", qFacilityIc); err != nil {
				return err
			}
		}
	}

	if o.FacilityIe != nil {

		// query param facility__ie
		var qrFacilityIe string

		if o.FacilityIe != nil {
			qrFacilityIe = *o.FacilityIe
		}
		qFacilityIe := qrFacilityIe
		if qFacilityIe != "" {

			if err := r.SetQueryParam("facility__ie", qFacilityIe); err != nil {
				return err
			}
		}
	}

	if o.FacilityIew != nil {

		// query param facility__iew
		var qrFacilityIew string

		if o.FacilityIew != nil {
			qrFacilityIew = *o.FacilityIew
		}
		qFacilityIew := qrFacilityIew
		if qFacilityIew != "" {

			if err := r.SetQueryParam("facility__iew", qFacilityIew); err != nil {
				return err
			}
		}
	}

	if o.FacilityIsw != nil {

		// query param facility__isw
		var qrFacilityIsw string

		if o.FacilityIsw != nil {
			qrFacilityIsw = *o.FacilityIsw
		}
		qFacilityIsw := qrFacilityIsw
		if qFacilityIsw != "" {

			if err := r.SetQueryParam("facility__isw", qFacilityIsw); err != nil {
				return err
			}
		}
	}

	if o.Facilityn != nil {

		// query param facility__n
		var qrFacilityn string

		if o.Facilityn != nil {
			qrFacilityn = *o.Facilityn
		}
		qFacilityn := qrFacilityn
		if qFacilityn != "" {

			if err := r.SetQueryParam("facility__n", qFacilityn); err != nil {
				return err
			}
		}
	}

	if o.FacilityNic != nil {

		// query param facility__nic
		var qrFacilityNic string

		if o.FacilityNic != nil {
			qrFacilityNic = *o.FacilityNic
		}
		qFacilityNic := qrFacilityNic
		if qFacilityNic != "" {

			if err := r.SetQueryParam("facility__nic", qFacilityNic); err != nil {
				return err
			}
		}
	}

	if o.FacilityNie != nil {

		// query param facility__nie
		var qrFacilityNie string

		if o.FacilityNie != nil {
			qrFacilityNie = *o.FacilityNie
		}
		qFacilityNie := qrFacilityNie
		if qFacilityNie != "" {

			if err := r.SetQueryParam("facility__nie", qFacilityNie); err != nil {
				return err
			}
		}
	}

	if o.FacilityNiew != nil {

		// query param facility__niew
		var qrFacilityNiew string

		if o.FacilityNiew != nil {
			qrFacilityNiew = *o.FacilityNiew
		}
		qFacilityNiew := qrFacilityNiew
		if qFacilityNiew != "" {

			if err := r.SetQueryParam("facility__niew", qFacilityNiew); err != nil {
				return err
			}
		}
	}

	if o.FacilityNisw != nil {

		// query param facility__nisw
		var qrFacilityNisw string

		if o.FacilityNisw != nil {
			qrFacilityNisw = *o.FacilityNisw
		}
		qFacilityNisw := qrFacilityNisw
		if qFacilityNisw != "" {

			if err := r.SetQueryParam("facility__nisw", qFacilityNisw); err != nil {
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

	if o.Latitude != nil {

		// query param latitude
		var qrLatitude string

		if o.Latitude != nil {
			qrLatitude = *o.Latitude
		}
		qLatitude := qrLatitude
		if qLatitude != "" {

			if err := r.SetQueryParam("latitude", qLatitude); err != nil {
				return err
			}
		}
	}

	if o.LatitudeGt != nil {

		// query param latitude__gt
		var qrLatitudeGt string

		if o.LatitudeGt != nil {
			qrLatitudeGt = *o.LatitudeGt
		}
		qLatitudeGt := qrLatitudeGt
		if qLatitudeGt != "" {

			if err := r.SetQueryParam("latitude__gt", qLatitudeGt); err != nil {
				return err
			}
		}
	}

	if o.LatitudeGte != nil {

		// query param latitude__gte
		var qrLatitudeGte string

		if o.LatitudeGte != nil {
			qrLatitudeGte = *o.LatitudeGte
		}
		qLatitudeGte := qrLatitudeGte
		if qLatitudeGte != "" {

			if err := r.SetQueryParam("latitude__gte", qLatitudeGte); err != nil {
				return err
			}
		}
	}

	if o.LatitudeLt != nil {

		// query param latitude__lt
		var qrLatitudeLt string

		if o.LatitudeLt != nil {
			qrLatitudeLt = *o.LatitudeLt
		}
		qLatitudeLt := qrLatitudeLt
		if qLatitudeLt != "" {

			if err := r.SetQueryParam("latitude__lt", qLatitudeLt); err != nil {
				return err
			}
		}
	}

	if o.LatitudeLte != nil {

		// query param latitude__lte
		var qrLatitudeLte string

		if o.LatitudeLte != nil {
			qrLatitudeLte = *o.LatitudeLte
		}
		qLatitudeLte := qrLatitudeLte
		if qLatitudeLte != "" {

			if err := r.SetQueryParam("latitude__lte", qLatitudeLte); err != nil {
				return err
			}
		}
	}

	if o.Latituden != nil {

		// query param latitude__n
		var qrLatituden string

		if o.Latituden != nil {
			qrLatituden = *o.Latituden
		}
		qLatituden := qrLatituden
		if qLatituden != "" {

			if err := r.SetQueryParam("latitude__n", qLatituden); err != nil {
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

	if o.Longitude != nil {

		// query param longitude
		var qrLongitude string

		if o.Longitude != nil {
			qrLongitude = *o.Longitude
		}
		qLongitude := qrLongitude
		if qLongitude != "" {

			if err := r.SetQueryParam("longitude", qLongitude); err != nil {
				return err
			}
		}
	}

	if o.LongitudeGt != nil {

		// query param longitude__gt
		var qrLongitudeGt string

		if o.LongitudeGt != nil {
			qrLongitudeGt = *o.LongitudeGt
		}
		qLongitudeGt := qrLongitudeGt
		if qLongitudeGt != "" {

			if err := r.SetQueryParam("longitude__gt", qLongitudeGt); err != nil {
				return err
			}
		}
	}

	if o.LongitudeGte != nil {

		// query param longitude__gte
		var qrLongitudeGte string

		if o.LongitudeGte != nil {
			qrLongitudeGte = *o.LongitudeGte
		}
		qLongitudeGte := qrLongitudeGte
		if qLongitudeGte != "" {

			if err := r.SetQueryParam("longitude__gte", qLongitudeGte); err != nil {
				return err
			}
		}
	}

	if o.LongitudeLt != nil {

		// query param longitude__lt
		var qrLongitudeLt string

		if o.LongitudeLt != nil {
			qrLongitudeLt = *o.LongitudeLt
		}
		qLongitudeLt := qrLongitudeLt
		if qLongitudeLt != "" {

			if err := r.SetQueryParam("longitude__lt", qLongitudeLt); err != nil {
				return err
			}
		}
	}

	if o.LongitudeLte != nil {

		// query param longitude__lte
		var qrLongitudeLte string

		if o.LongitudeLte != nil {
			qrLongitudeLte = *o.LongitudeLte
		}
		qLongitudeLte := qrLongitudeLte
		if qLongitudeLte != "" {

			if err := r.SetQueryParam("longitude__lte", qLongitudeLte); err != nil {
				return err
			}
		}
	}

	if o.Longituden != nil {

		// query param longitude__n
		var qrLongituden string

		if o.Longituden != nil {
			qrLongituden = *o.Longituden
		}
		qLongituden := qrLongituden
		if qLongituden != "" {

			if err := r.SetQueryParam("longitude__n", qLongituden); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
