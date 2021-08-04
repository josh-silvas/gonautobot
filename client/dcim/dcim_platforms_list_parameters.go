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

// NewDcimPlatformsListParams creates a new DcimPlatformsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPlatformsListParams() *DcimPlatformsListParams {
	return &DcimPlatformsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPlatformsListParamsWithTimeout creates a new DcimPlatformsListParams object
// with the ability to set a timeout on a request.
func NewDcimPlatformsListParamsWithTimeout(timeout time.Duration) *DcimPlatformsListParams {
	return &DcimPlatformsListParams{
		timeout: timeout,
	}
}

// NewDcimPlatformsListParamsWithContext creates a new DcimPlatformsListParams object
// with the ability to set a context for a request.
func NewDcimPlatformsListParamsWithContext(ctx context.Context) *DcimPlatformsListParams {
	return &DcimPlatformsListParams{
		Context: ctx,
	}
}

// NewDcimPlatformsListParamsWithHTTPClient creates a new DcimPlatformsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPlatformsListParamsWithHTTPClient(client *http.Client) *DcimPlatformsListParams {
	return &DcimPlatformsListParams{
		HTTPClient: client,
	}
}

/* DcimPlatformsListParams contains all the parameters to send to the API endpoint
   for the dcim platforms list operation.

   Typically these are written to a http.Request.
*/
type DcimPlatformsListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Description.
	Description *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

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

	// Manufacturer.
	Manufacturer *string

	// Manufacturern.
	Manufacturern *string

	// ManufacturerID.
	ManufacturerID *string

	// ManufacturerIDn.
	ManufacturerIDn *string

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

	// NapalmDriver.
	NapalmDriver *string

	// NapalmDriverIc.
	NapalmDriverIc *string

	// NapalmDriverIe.
	NapalmDriverIe *string

	// NapalmDriverIew.
	NapalmDriverIew *string

	// NapalmDriverIsw.
	NapalmDriverIsw *string

	// NapalmDrivern.
	NapalmDrivern *string

	// NapalmDriverNic.
	NapalmDriverNic *string

	// NapalmDriverNie.
	NapalmDriverNie *string

	// NapalmDriverNiew.
	NapalmDriverNiew *string

	// NapalmDriverNisw.
	NapalmDriverNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim platforms list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsListParams) WithDefaults() *DcimPlatformsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim platforms list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPlatformsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim platforms list params
func (o *DcimPlatformsListParams) WithTimeout(timeout time.Duration) *DcimPlatformsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim platforms list params
func (o *DcimPlatformsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim platforms list params
func (o *DcimPlatformsListParams) WithContext(ctx context.Context) *DcimPlatformsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim platforms list params
func (o *DcimPlatformsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim platforms list params
func (o *DcimPlatformsListParams) WithHTTPClient(client *http.Client) *DcimPlatformsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim platforms list params
func (o *DcimPlatformsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the dcim platforms list params
func (o *DcimPlatformsListParams) WithCreated(created *string) *DcimPlatformsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim platforms list params
func (o *DcimPlatformsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim platforms list params
func (o *DcimPlatformsListParams) WithCreatedGte(createdGte *string) *DcimPlatformsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim platforms list params
func (o *DcimPlatformsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim platforms list params
func (o *DcimPlatformsListParams) WithCreatedLte(createdLte *string) *DcimPlatformsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim platforms list params
func (o *DcimPlatformsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescription(description *string) *DcimPlatformsListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionIc adds the descriptionIc to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionIc(descriptionIc *string) *DcimPlatformsListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionIe(descriptionIe *string) *DcimPlatformsListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionIew(descriptionIew *string) *DcimPlatformsListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionIsw(descriptionIsw *string) *DcimPlatformsListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionn(descriptionn *string) *DcimPlatformsListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionNic(descriptionNic *string) *DcimPlatformsListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionNie(descriptionNie *string) *DcimPlatformsListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionNiew(descriptionNiew *string) *DcimPlatformsListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithDescriptionNisw(descriptionNisw *string) *DcimPlatformsListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithID adds the id to the dcim platforms list params
func (o *DcimPlatformsListParams) WithID(id *string) *DcimPlatformsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim platforms list params
func (o *DcimPlatformsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDIc(iDIc *string) *DcimPlatformsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDIe(iDIe *string) *DcimPlatformsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDIew(iDIew *string) *DcimPlatformsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDIsw(iDIsw *string) *DcimPlatformsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDn(iDn *string) *DcimPlatformsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDNic(iDNic *string) *DcimPlatformsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDNie(iDNie *string) *DcimPlatformsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDNiew(iDNiew *string) *DcimPlatformsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithIDNisw(iDNisw *string) *DcimPlatformsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the dcim platforms list params
func (o *DcimPlatformsListParams) WithLastUpdated(lastUpdated *string) *DcimPlatformsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim platforms list params
func (o *DcimPlatformsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim platforms list params
func (o *DcimPlatformsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimPlatformsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim platforms list params
func (o *DcimPlatformsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim platforms list params
func (o *DcimPlatformsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimPlatformsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim platforms list params
func (o *DcimPlatformsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim platforms list params
func (o *DcimPlatformsListParams) WithLimit(limit *int64) *DcimPlatformsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim platforms list params
func (o *DcimPlatformsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the dcim platforms list params
func (o *DcimPlatformsListParams) WithManufacturer(manufacturer *string) *DcimPlatformsListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim platforms list params
func (o *DcimPlatformsListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturern adds the manufacturern to the dcim platforms list params
func (o *DcimPlatformsListParams) WithManufacturern(manufacturern *string) *DcimPlatformsListParams {
	o.SetManufacturern(manufacturern)
	return o
}

// SetManufacturern adds the manufacturerN to the dcim platforms list params
func (o *DcimPlatformsListParams) SetManufacturern(manufacturern *string) {
	o.Manufacturern = manufacturern
}

// WithManufacturerID adds the manufacturerID to the dcim platforms list params
func (o *DcimPlatformsListParams) WithManufacturerID(manufacturerID *string) *DcimPlatformsListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim platforms list params
func (o *DcimPlatformsListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithManufacturerIDn adds the manufacturerIDn to the dcim platforms list params
func (o *DcimPlatformsListParams) WithManufacturerIDn(manufacturerIDn *string) *DcimPlatformsListParams {
	o.SetManufacturerIDn(manufacturerIDn)
	return o
}

// SetManufacturerIDn adds the manufacturerIdN to the dcim platforms list params
func (o *DcimPlatformsListParams) SetManufacturerIDn(manufacturerIDn *string) {
	o.ManufacturerIDn = manufacturerIDn
}

// WithName adds the name to the dcim platforms list params
func (o *DcimPlatformsListParams) WithName(name *string) *DcimPlatformsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim platforms list params
func (o *DcimPlatformsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameIc(nameIc *string) *DcimPlatformsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameIe(nameIe *string) *DcimPlatformsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameIew(nameIew *string) *DcimPlatformsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameIsw(nameIsw *string) *DcimPlatformsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNamen(namen *string) *DcimPlatformsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameNic(nameNic *string) *DcimPlatformsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameNie(nameNie *string) *DcimPlatformsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameNiew(nameNiew *string) *DcimPlatformsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNameNisw(nameNisw *string) *DcimPlatformsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithNapalmDriver adds the napalmDriver to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriver(napalmDriver *string) *DcimPlatformsListParams {
	o.SetNapalmDriver(napalmDriver)
	return o
}

// SetNapalmDriver adds the napalmDriver to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriver(napalmDriver *string) {
	o.NapalmDriver = napalmDriver
}

// WithNapalmDriverIc adds the napalmDriverIc to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverIc(napalmDriverIc *string) *DcimPlatformsListParams {
	o.SetNapalmDriverIc(napalmDriverIc)
	return o
}

// SetNapalmDriverIc adds the napalmDriverIc to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverIc(napalmDriverIc *string) {
	o.NapalmDriverIc = napalmDriverIc
}

// WithNapalmDriverIe adds the napalmDriverIe to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverIe(napalmDriverIe *string) *DcimPlatformsListParams {
	o.SetNapalmDriverIe(napalmDriverIe)
	return o
}

// SetNapalmDriverIe adds the napalmDriverIe to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverIe(napalmDriverIe *string) {
	o.NapalmDriverIe = napalmDriverIe
}

// WithNapalmDriverIew adds the napalmDriverIew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverIew(napalmDriverIew *string) *DcimPlatformsListParams {
	o.SetNapalmDriverIew(napalmDriverIew)
	return o
}

// SetNapalmDriverIew adds the napalmDriverIew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverIew(napalmDriverIew *string) {
	o.NapalmDriverIew = napalmDriverIew
}

// WithNapalmDriverIsw adds the napalmDriverIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverIsw(napalmDriverIsw *string) *DcimPlatformsListParams {
	o.SetNapalmDriverIsw(napalmDriverIsw)
	return o
}

// SetNapalmDriverIsw adds the napalmDriverIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverIsw(napalmDriverIsw *string) {
	o.NapalmDriverIsw = napalmDriverIsw
}

// WithNapalmDrivern adds the napalmDrivern to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDrivern(napalmDrivern *string) *DcimPlatformsListParams {
	o.SetNapalmDrivern(napalmDrivern)
	return o
}

// SetNapalmDrivern adds the napalmDriverN to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDrivern(napalmDrivern *string) {
	o.NapalmDrivern = napalmDrivern
}

// WithNapalmDriverNic adds the napalmDriverNic to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverNic(napalmDriverNic *string) *DcimPlatformsListParams {
	o.SetNapalmDriverNic(napalmDriverNic)
	return o
}

// SetNapalmDriverNic adds the napalmDriverNic to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverNic(napalmDriverNic *string) {
	o.NapalmDriverNic = napalmDriverNic
}

// WithNapalmDriverNie adds the napalmDriverNie to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverNie(napalmDriverNie *string) *DcimPlatformsListParams {
	o.SetNapalmDriverNie(napalmDriverNie)
	return o
}

// SetNapalmDriverNie adds the napalmDriverNie to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverNie(napalmDriverNie *string) {
	o.NapalmDriverNie = napalmDriverNie
}

// WithNapalmDriverNiew adds the napalmDriverNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverNiew(napalmDriverNiew *string) *DcimPlatformsListParams {
	o.SetNapalmDriverNiew(napalmDriverNiew)
	return o
}

// SetNapalmDriverNiew adds the napalmDriverNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverNiew(napalmDriverNiew *string) {
	o.NapalmDriverNiew = napalmDriverNiew
}

// WithNapalmDriverNisw adds the napalmDriverNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithNapalmDriverNisw(napalmDriverNisw *string) *DcimPlatformsListParams {
	o.SetNapalmDriverNisw(napalmDriverNisw)
	return o
}

// SetNapalmDriverNisw adds the napalmDriverNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetNapalmDriverNisw(napalmDriverNisw *string) {
	o.NapalmDriverNisw = napalmDriverNisw
}

// WithOffset adds the offset to the dcim platforms list params
func (o *DcimPlatformsListParams) WithOffset(offset *int64) *DcimPlatformsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim platforms list params
func (o *DcimPlatformsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim platforms list params
func (o *DcimPlatformsListParams) WithQ(q *string) *DcimPlatformsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim platforms list params
func (o *DcimPlatformsListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlug(slug *string) *DcimPlatformsListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugIc(slugIc *string) *DcimPlatformsListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugIe(slugIe *string) *DcimPlatformsListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugIew(slugIew *string) *DcimPlatformsListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugIsw(slugIsw *string) *DcimPlatformsListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugn(slugn *string) *DcimPlatformsListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugNic(slugNic *string) *DcimPlatformsListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugNie(slugNie *string) *DcimPlatformsListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugNiew(slugNiew *string) *DcimPlatformsListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) WithSlugNisw(slugNisw *string) *DcimPlatformsListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the dcim platforms list params
func (o *DcimPlatformsListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPlatformsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
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

	if o.NapalmDriver != nil {

		// query param napalm_driver
		var qrNapalmDriver string

		if o.NapalmDriver != nil {
			qrNapalmDriver = *o.NapalmDriver
		}
		qNapalmDriver := qrNapalmDriver
		if qNapalmDriver != "" {

			if err := r.SetQueryParam("napalm_driver", qNapalmDriver); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverIc != nil {

		// query param napalm_driver__ic
		var qrNapalmDriverIc string

		if o.NapalmDriverIc != nil {
			qrNapalmDriverIc = *o.NapalmDriverIc
		}
		qNapalmDriverIc := qrNapalmDriverIc
		if qNapalmDriverIc != "" {

			if err := r.SetQueryParam("napalm_driver__ic", qNapalmDriverIc); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverIe != nil {

		// query param napalm_driver__ie
		var qrNapalmDriverIe string

		if o.NapalmDriverIe != nil {
			qrNapalmDriverIe = *o.NapalmDriverIe
		}
		qNapalmDriverIe := qrNapalmDriverIe
		if qNapalmDriverIe != "" {

			if err := r.SetQueryParam("napalm_driver__ie", qNapalmDriverIe); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverIew != nil {

		// query param napalm_driver__iew
		var qrNapalmDriverIew string

		if o.NapalmDriverIew != nil {
			qrNapalmDriverIew = *o.NapalmDriverIew
		}
		qNapalmDriverIew := qrNapalmDriverIew
		if qNapalmDriverIew != "" {

			if err := r.SetQueryParam("napalm_driver__iew", qNapalmDriverIew); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverIsw != nil {

		// query param napalm_driver__isw
		var qrNapalmDriverIsw string

		if o.NapalmDriverIsw != nil {
			qrNapalmDriverIsw = *o.NapalmDriverIsw
		}
		qNapalmDriverIsw := qrNapalmDriverIsw
		if qNapalmDriverIsw != "" {

			if err := r.SetQueryParam("napalm_driver__isw", qNapalmDriverIsw); err != nil {
				return err
			}
		}
	}

	if o.NapalmDrivern != nil {

		// query param napalm_driver__n
		var qrNapalmDrivern string

		if o.NapalmDrivern != nil {
			qrNapalmDrivern = *o.NapalmDrivern
		}
		qNapalmDrivern := qrNapalmDrivern
		if qNapalmDrivern != "" {

			if err := r.SetQueryParam("napalm_driver__n", qNapalmDrivern); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverNic != nil {

		// query param napalm_driver__nic
		var qrNapalmDriverNic string

		if o.NapalmDriverNic != nil {
			qrNapalmDriverNic = *o.NapalmDriverNic
		}
		qNapalmDriverNic := qrNapalmDriverNic
		if qNapalmDriverNic != "" {

			if err := r.SetQueryParam("napalm_driver__nic", qNapalmDriverNic); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverNie != nil {

		// query param napalm_driver__nie
		var qrNapalmDriverNie string

		if o.NapalmDriverNie != nil {
			qrNapalmDriverNie = *o.NapalmDriverNie
		}
		qNapalmDriverNie := qrNapalmDriverNie
		if qNapalmDriverNie != "" {

			if err := r.SetQueryParam("napalm_driver__nie", qNapalmDriverNie); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverNiew != nil {

		// query param napalm_driver__niew
		var qrNapalmDriverNiew string

		if o.NapalmDriverNiew != nil {
			qrNapalmDriverNiew = *o.NapalmDriverNiew
		}
		qNapalmDriverNiew := qrNapalmDriverNiew
		if qNapalmDriverNiew != "" {

			if err := r.SetQueryParam("napalm_driver__niew", qNapalmDriverNiew); err != nil {
				return err
			}
		}
	}

	if o.NapalmDriverNisw != nil {

		// query param napalm_driver__nisw
		var qrNapalmDriverNisw string

		if o.NapalmDriverNisw != nil {
			qrNapalmDriverNisw = *o.NapalmDriverNisw
		}
		qNapalmDriverNisw := qrNapalmDriverNisw
		if qNapalmDriverNisw != "" {

			if err := r.SetQueryParam("napalm_driver__nisw", qNapalmDriverNisw); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
