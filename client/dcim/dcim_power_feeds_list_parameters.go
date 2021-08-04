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

// NewDcimPowerFeedsListParams creates a new DcimPowerFeedsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimPowerFeedsListParams() *DcimPowerFeedsListParams {
	return &DcimPowerFeedsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerFeedsListParamsWithTimeout creates a new DcimPowerFeedsListParams object
// with the ability to set a timeout on a request.
func NewDcimPowerFeedsListParamsWithTimeout(timeout time.Duration) *DcimPowerFeedsListParams {
	return &DcimPowerFeedsListParams{
		timeout: timeout,
	}
}

// NewDcimPowerFeedsListParamsWithContext creates a new DcimPowerFeedsListParams object
// with the ability to set a context for a request.
func NewDcimPowerFeedsListParamsWithContext(ctx context.Context) *DcimPowerFeedsListParams {
	return &DcimPowerFeedsListParams{
		Context: ctx,
	}
}

// NewDcimPowerFeedsListParamsWithHTTPClient creates a new DcimPowerFeedsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimPowerFeedsListParamsWithHTTPClient(client *http.Client) *DcimPowerFeedsListParams {
	return &DcimPowerFeedsListParams{
		HTTPClient: client,
	}
}

/* DcimPowerFeedsListParams contains all the parameters to send to the API endpoint
   for the dcim power feeds list operation.

   Typically these are written to a http.Request.
*/
type DcimPowerFeedsListParams struct {

	// Amperage.
	Amperage *string

	// AmperageGt.
	AmperageGt *string

	// AmperageGte.
	AmperageGte *string

	// AmperageLt.
	AmperageLt *string

	// AmperageLte.
	AmperageLte *string

	// Amperagen.
	Amperagen *string

	// Cabled.
	Cabled *string

	// Connected.
	Connected *string

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

	// MaxUtilization.
	MaxUtilization *string

	// MaxUtilizationGt.
	MaxUtilizationGt *string

	// MaxUtilizationGte.
	MaxUtilizationGte *string

	// MaxUtilizationLt.
	MaxUtilizationLt *string

	// MaxUtilizationLte.
	MaxUtilizationLte *string

	// MaxUtilizationn.
	MaxUtilizationn *string

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

	// Phase.
	Phase *string

	// Phasen.
	Phasen *string

	// PowerPanelID.
	PowerPanelID *string

	// PowerPanelIDn.
	PowerPanelIDn *string

	// Q.
	Q *string

	// RackID.
	RackID *string

	// RackIDn.
	RackIDn *string

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

	// Supply.
	Supply *string

	// Supplyn.
	Supplyn *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// Type.
	Type *string

	// Typen.
	Typen *string

	// Voltage.
	Voltage *string

	// VoltageGt.
	VoltageGt *string

	// VoltageGte.
	VoltageGte *string

	// VoltageLt.
	VoltageLt *string

	// VoltageLte.
	VoltageLte *string

	// Voltagen.
	Voltagen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim power feeds list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsListParams) WithDefaults() *DcimPowerFeedsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim power feeds list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimPowerFeedsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithTimeout(timeout time.Duration) *DcimPowerFeedsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithContext(ctx context.Context) *DcimPowerFeedsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithHTTPClient(client *http.Client) *DcimPowerFeedsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAmperage adds the amperage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithAmperage(amperage *string) *DcimPowerFeedsListParams {
	o.SetAmperage(amperage)
	return o
}

// SetAmperage adds the amperage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetAmperage(amperage *string) {
	o.Amperage = amperage
}

// WithAmperageGt adds the amperageGt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithAmperageGt(amperageGt *string) *DcimPowerFeedsListParams {
	o.SetAmperageGt(amperageGt)
	return o
}

// SetAmperageGt adds the amperageGt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetAmperageGt(amperageGt *string) {
	o.AmperageGt = amperageGt
}

// WithAmperageGte adds the amperageGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithAmperageGte(amperageGte *string) *DcimPowerFeedsListParams {
	o.SetAmperageGte(amperageGte)
	return o
}

// SetAmperageGte adds the amperageGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetAmperageGte(amperageGte *string) {
	o.AmperageGte = amperageGte
}

// WithAmperageLt adds the amperageLt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithAmperageLt(amperageLt *string) *DcimPowerFeedsListParams {
	o.SetAmperageLt(amperageLt)
	return o
}

// SetAmperageLt adds the amperageLt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetAmperageLt(amperageLt *string) {
	o.AmperageLt = amperageLt
}

// WithAmperageLte adds the amperageLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithAmperageLte(amperageLte *string) *DcimPowerFeedsListParams {
	o.SetAmperageLte(amperageLte)
	return o
}

// SetAmperageLte adds the amperageLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetAmperageLte(amperageLte *string) {
	o.AmperageLte = amperageLte
}

// WithAmperagen adds the amperagen to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithAmperagen(amperagen *string) *DcimPowerFeedsListParams {
	o.SetAmperagen(amperagen)
	return o
}

// SetAmperagen adds the amperageN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetAmperagen(amperagen *string) {
	o.Amperagen = amperagen
}

// WithCabled adds the cabled to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithCabled(cabled *string) *DcimPowerFeedsListParams {
	o.SetCabled(cabled)
	return o
}

// SetCabled adds the cabled to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetCabled(cabled *string) {
	o.Cabled = cabled
}

// WithConnected adds the connected to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithConnected(connected *string) *DcimPowerFeedsListParams {
	o.SetConnected(connected)
	return o
}

// SetConnected adds the connected to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetConnected(connected *string) {
	o.Connected = connected
}

// WithCreated adds the created to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithCreated(created *string) *DcimPowerFeedsListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithCreatedGte(createdGte *string) *DcimPowerFeedsListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithCreatedLte(createdLte *string) *DcimPowerFeedsListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithID adds the id to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithID(id *string) *DcimPowerFeedsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDIc(iDIc *string) *DcimPowerFeedsListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDIe(iDIe *string) *DcimPowerFeedsListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDIew(iDIew *string) *DcimPowerFeedsListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDIsw(iDIsw *string) *DcimPowerFeedsListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDn(iDn *string) *DcimPowerFeedsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDNic(iDNic *string) *DcimPowerFeedsListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDNie(iDNie *string) *DcimPowerFeedsListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDNiew(iDNiew *string) *DcimPowerFeedsListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithIDNisw(iDNisw *string) *DcimPowerFeedsListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithLastUpdated(lastUpdated *string) *DcimPowerFeedsListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimPowerFeedsListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimPowerFeedsListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithLimit(limit *int64) *DcimPowerFeedsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaxUtilization adds the maxUtilization to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithMaxUtilization(maxUtilization *string) *DcimPowerFeedsListParams {
	o.SetMaxUtilization(maxUtilization)
	return o
}

// SetMaxUtilization adds the maxUtilization to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetMaxUtilization(maxUtilization *string) {
	o.MaxUtilization = maxUtilization
}

// WithMaxUtilizationGt adds the maxUtilizationGt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithMaxUtilizationGt(maxUtilizationGt *string) *DcimPowerFeedsListParams {
	o.SetMaxUtilizationGt(maxUtilizationGt)
	return o
}

// SetMaxUtilizationGt adds the maxUtilizationGt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetMaxUtilizationGt(maxUtilizationGt *string) {
	o.MaxUtilizationGt = maxUtilizationGt
}

// WithMaxUtilizationGte adds the maxUtilizationGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithMaxUtilizationGte(maxUtilizationGte *string) *DcimPowerFeedsListParams {
	o.SetMaxUtilizationGte(maxUtilizationGte)
	return o
}

// SetMaxUtilizationGte adds the maxUtilizationGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetMaxUtilizationGte(maxUtilizationGte *string) {
	o.MaxUtilizationGte = maxUtilizationGte
}

// WithMaxUtilizationLt adds the maxUtilizationLt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithMaxUtilizationLt(maxUtilizationLt *string) *DcimPowerFeedsListParams {
	o.SetMaxUtilizationLt(maxUtilizationLt)
	return o
}

// SetMaxUtilizationLt adds the maxUtilizationLt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetMaxUtilizationLt(maxUtilizationLt *string) {
	o.MaxUtilizationLt = maxUtilizationLt
}

// WithMaxUtilizationLte adds the maxUtilizationLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithMaxUtilizationLte(maxUtilizationLte *string) *DcimPowerFeedsListParams {
	o.SetMaxUtilizationLte(maxUtilizationLte)
	return o
}

// SetMaxUtilizationLte adds the maxUtilizationLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetMaxUtilizationLte(maxUtilizationLte *string) {
	o.MaxUtilizationLte = maxUtilizationLte
}

// WithMaxUtilizationn adds the maxUtilizationn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithMaxUtilizationn(maxUtilizationn *string) *DcimPowerFeedsListParams {
	o.SetMaxUtilizationn(maxUtilizationn)
	return o
}

// SetMaxUtilizationn adds the maxUtilizationN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetMaxUtilizationn(maxUtilizationn *string) {
	o.MaxUtilizationn = maxUtilizationn
}

// WithName adds the name to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithName(name *string) *DcimPowerFeedsListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameIc(nameIc *string) *DcimPowerFeedsListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameIe(nameIe *string) *DcimPowerFeedsListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameIew(nameIew *string) *DcimPowerFeedsListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameIsw(nameIsw *string) *DcimPowerFeedsListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNamen(namen *string) *DcimPowerFeedsListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameNic(nameNic *string) *DcimPowerFeedsListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameNie(nameNie *string) *DcimPowerFeedsListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameNiew(nameNiew *string) *DcimPowerFeedsListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithNameNisw(nameNisw *string) *DcimPowerFeedsListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithOffset(offset *int64) *DcimPowerFeedsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPhase adds the phase to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithPhase(phase *string) *DcimPowerFeedsListParams {
	o.SetPhase(phase)
	return o
}

// SetPhase adds the phase to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetPhase(phase *string) {
	o.Phase = phase
}

// WithPhasen adds the phasen to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithPhasen(phasen *string) *DcimPowerFeedsListParams {
	o.SetPhasen(phasen)
	return o
}

// SetPhasen adds the phaseN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetPhasen(phasen *string) {
	o.Phasen = phasen
}

// WithPowerPanelID adds the powerPanelID to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithPowerPanelID(powerPanelID *string) *DcimPowerFeedsListParams {
	o.SetPowerPanelID(powerPanelID)
	return o
}

// SetPowerPanelID adds the powerPanelId to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetPowerPanelID(powerPanelID *string) {
	o.PowerPanelID = powerPanelID
}

// WithPowerPanelIDn adds the powerPanelIDn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithPowerPanelIDn(powerPanelIDn *string) *DcimPowerFeedsListParams {
	o.SetPowerPanelIDn(powerPanelIDn)
	return o
}

// SetPowerPanelIDn adds the powerPanelIdN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetPowerPanelIDn(powerPanelIDn *string) {
	o.PowerPanelIDn = powerPanelIDn
}

// WithQ adds the q to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithQ(q *string) *DcimPowerFeedsListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetQ(q *string) {
	o.Q = q
}

// WithRackID adds the rackID to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithRackID(rackID *string) *DcimPowerFeedsListParams {
	o.SetRackID(rackID)
	return o
}

// SetRackID adds the rackId to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetRackID(rackID *string) {
	o.RackID = rackID
}

// WithRackIDn adds the rackIDn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithRackIDn(rackIDn *string) *DcimPowerFeedsListParams {
	o.SetRackIDn(rackIDn)
	return o
}

// SetRackIDn adds the rackIdN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetRackIDn(rackIDn *string) {
	o.RackIDn = rackIDn
}

// WithRegion adds the region to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithRegion(region *string) *DcimPowerFeedsListParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetRegion(region *string) {
	o.Region = region
}

// WithRegionn adds the regionn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithRegionn(regionn *string) *DcimPowerFeedsListParams {
	o.SetRegionn(regionn)
	return o
}

// SetRegionn adds the regionN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetRegionn(regionn *string) {
	o.Regionn = regionn
}

// WithRegionID adds the regionID to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithRegionID(regionID *string) *DcimPowerFeedsListParams {
	o.SetRegionID(regionID)
	return o
}

// SetRegionID adds the regionId to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetRegionID(regionID *string) {
	o.RegionID = regionID
}

// WithRegionIDn adds the regionIDn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithRegionIDn(regionIDn *string) *DcimPowerFeedsListParams {
	o.SetRegionIDn(regionIDn)
	return o
}

// SetRegionIDn adds the regionIdN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetRegionIDn(regionIDn *string) {
	o.RegionIDn = regionIDn
}

// WithSite adds the site to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSite(site *string) *DcimPowerFeedsListParams {
	o.SetSite(site)
	return o
}

// SetSite adds the site to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSite(site *string) {
	o.Site = site
}

// WithSiten adds the siten to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSiten(siten *string) *DcimPowerFeedsListParams {
	o.SetSiten(siten)
	return o
}

// SetSiten adds the siteN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSiten(siten *string) {
	o.Siten = siten
}

// WithSiteID adds the siteID to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSiteID(siteID *string) *DcimPowerFeedsListParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSiteIDn adds the siteIDn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSiteIDn(siteIDn *string) *DcimPowerFeedsListParams {
	o.SetSiteIDn(siteIDn)
	return o
}

// SetSiteIDn adds the siteIdN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSiteIDn(siteIDn *string) {
	o.SiteIDn = siteIDn
}

// WithStatus adds the status to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithStatus(status *string) *DcimPowerFeedsListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithStatusn(statusn *string) *DcimPowerFeedsListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithSupply adds the supply to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSupply(supply *string) *DcimPowerFeedsListParams {
	o.SetSupply(supply)
	return o
}

// SetSupply adds the supply to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSupply(supply *string) {
	o.Supply = supply
}

// WithSupplyn adds the supplyn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithSupplyn(supplyn *string) *DcimPowerFeedsListParams {
	o.SetSupplyn(supplyn)
	return o
}

// SetSupplyn adds the supplyN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetSupplyn(supplyn *string) {
	o.Supplyn = supplyn
}

// WithTag adds the tag to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithTag(tag *string) *DcimPowerFeedsListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithTagn(tagn *string) *DcimPowerFeedsListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithType adds the typeVar to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithType(typeVar *string) *DcimPowerFeedsListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithTypen(typen *string) *DcimPowerFeedsListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WithVoltage adds the voltage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithVoltage(voltage *string) *DcimPowerFeedsListParams {
	o.SetVoltage(voltage)
	return o
}

// SetVoltage adds the voltage to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetVoltage(voltage *string) {
	o.Voltage = voltage
}

// WithVoltageGt adds the voltageGt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithVoltageGt(voltageGt *string) *DcimPowerFeedsListParams {
	o.SetVoltageGt(voltageGt)
	return o
}

// SetVoltageGt adds the voltageGt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetVoltageGt(voltageGt *string) {
	o.VoltageGt = voltageGt
}

// WithVoltageGte adds the voltageGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithVoltageGte(voltageGte *string) *DcimPowerFeedsListParams {
	o.SetVoltageGte(voltageGte)
	return o
}

// SetVoltageGte adds the voltageGte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetVoltageGte(voltageGte *string) {
	o.VoltageGte = voltageGte
}

// WithVoltageLt adds the voltageLt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithVoltageLt(voltageLt *string) *DcimPowerFeedsListParams {
	o.SetVoltageLt(voltageLt)
	return o
}

// SetVoltageLt adds the voltageLt to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetVoltageLt(voltageLt *string) {
	o.VoltageLt = voltageLt
}

// WithVoltageLte adds the voltageLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithVoltageLte(voltageLte *string) *DcimPowerFeedsListParams {
	o.SetVoltageLte(voltageLte)
	return o
}

// SetVoltageLte adds the voltageLte to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetVoltageLte(voltageLte *string) {
	o.VoltageLte = voltageLte
}

// WithVoltagen adds the voltagen to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) WithVoltagen(voltagen *string) *DcimPowerFeedsListParams {
	o.SetVoltagen(voltagen)
	return o
}

// SetVoltagen adds the voltageN to the dcim power feeds list params
func (o *DcimPowerFeedsListParams) SetVoltagen(voltagen *string) {
	o.Voltagen = voltagen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerFeedsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Amperage != nil {

		// query param amperage
		var qrAmperage string

		if o.Amperage != nil {
			qrAmperage = *o.Amperage
		}
		qAmperage := qrAmperage
		if qAmperage != "" {

			if err := r.SetQueryParam("amperage", qAmperage); err != nil {
				return err
			}
		}
	}

	if o.AmperageGt != nil {

		// query param amperage__gt
		var qrAmperageGt string

		if o.AmperageGt != nil {
			qrAmperageGt = *o.AmperageGt
		}
		qAmperageGt := qrAmperageGt
		if qAmperageGt != "" {

			if err := r.SetQueryParam("amperage__gt", qAmperageGt); err != nil {
				return err
			}
		}
	}

	if o.AmperageGte != nil {

		// query param amperage__gte
		var qrAmperageGte string

		if o.AmperageGte != nil {
			qrAmperageGte = *o.AmperageGte
		}
		qAmperageGte := qrAmperageGte
		if qAmperageGte != "" {

			if err := r.SetQueryParam("amperage__gte", qAmperageGte); err != nil {
				return err
			}
		}
	}

	if o.AmperageLt != nil {

		// query param amperage__lt
		var qrAmperageLt string

		if o.AmperageLt != nil {
			qrAmperageLt = *o.AmperageLt
		}
		qAmperageLt := qrAmperageLt
		if qAmperageLt != "" {

			if err := r.SetQueryParam("amperage__lt", qAmperageLt); err != nil {
				return err
			}
		}
	}

	if o.AmperageLte != nil {

		// query param amperage__lte
		var qrAmperageLte string

		if o.AmperageLte != nil {
			qrAmperageLte = *o.AmperageLte
		}
		qAmperageLte := qrAmperageLte
		if qAmperageLte != "" {

			if err := r.SetQueryParam("amperage__lte", qAmperageLte); err != nil {
				return err
			}
		}
	}

	if o.Amperagen != nil {

		// query param amperage__n
		var qrAmperagen string

		if o.Amperagen != nil {
			qrAmperagen = *o.Amperagen
		}
		qAmperagen := qrAmperagen
		if qAmperagen != "" {

			if err := r.SetQueryParam("amperage__n", qAmperagen); err != nil {
				return err
			}
		}
	}

	if o.Cabled != nil {

		// query param cabled
		var qrCabled string

		if o.Cabled != nil {
			qrCabled = *o.Cabled
		}
		qCabled := qrCabled
		if qCabled != "" {

			if err := r.SetQueryParam("cabled", qCabled); err != nil {
				return err
			}
		}
	}

	if o.Connected != nil {

		// query param connected
		var qrConnected string

		if o.Connected != nil {
			qrConnected = *o.Connected
		}
		qConnected := qrConnected
		if qConnected != "" {

			if err := r.SetQueryParam("connected", qConnected); err != nil {
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

	if o.MaxUtilization != nil {

		// query param max_utilization
		var qrMaxUtilization string

		if o.MaxUtilization != nil {
			qrMaxUtilization = *o.MaxUtilization
		}
		qMaxUtilization := qrMaxUtilization
		if qMaxUtilization != "" {

			if err := r.SetQueryParam("max_utilization", qMaxUtilization); err != nil {
				return err
			}
		}
	}

	if o.MaxUtilizationGt != nil {

		// query param max_utilization__gt
		var qrMaxUtilizationGt string

		if o.MaxUtilizationGt != nil {
			qrMaxUtilizationGt = *o.MaxUtilizationGt
		}
		qMaxUtilizationGt := qrMaxUtilizationGt
		if qMaxUtilizationGt != "" {

			if err := r.SetQueryParam("max_utilization__gt", qMaxUtilizationGt); err != nil {
				return err
			}
		}
	}

	if o.MaxUtilizationGte != nil {

		// query param max_utilization__gte
		var qrMaxUtilizationGte string

		if o.MaxUtilizationGte != nil {
			qrMaxUtilizationGte = *o.MaxUtilizationGte
		}
		qMaxUtilizationGte := qrMaxUtilizationGte
		if qMaxUtilizationGte != "" {

			if err := r.SetQueryParam("max_utilization__gte", qMaxUtilizationGte); err != nil {
				return err
			}
		}
	}

	if o.MaxUtilizationLt != nil {

		// query param max_utilization__lt
		var qrMaxUtilizationLt string

		if o.MaxUtilizationLt != nil {
			qrMaxUtilizationLt = *o.MaxUtilizationLt
		}
		qMaxUtilizationLt := qrMaxUtilizationLt
		if qMaxUtilizationLt != "" {

			if err := r.SetQueryParam("max_utilization__lt", qMaxUtilizationLt); err != nil {
				return err
			}
		}
	}

	if o.MaxUtilizationLte != nil {

		// query param max_utilization__lte
		var qrMaxUtilizationLte string

		if o.MaxUtilizationLte != nil {
			qrMaxUtilizationLte = *o.MaxUtilizationLte
		}
		qMaxUtilizationLte := qrMaxUtilizationLte
		if qMaxUtilizationLte != "" {

			if err := r.SetQueryParam("max_utilization__lte", qMaxUtilizationLte); err != nil {
				return err
			}
		}
	}

	if o.MaxUtilizationn != nil {

		// query param max_utilization__n
		var qrMaxUtilizationn string

		if o.MaxUtilizationn != nil {
			qrMaxUtilizationn = *o.MaxUtilizationn
		}
		qMaxUtilizationn := qrMaxUtilizationn
		if qMaxUtilizationn != "" {

			if err := r.SetQueryParam("max_utilization__n", qMaxUtilizationn); err != nil {
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

	if o.Phase != nil {

		// query param phase
		var qrPhase string

		if o.Phase != nil {
			qrPhase = *o.Phase
		}
		qPhase := qrPhase
		if qPhase != "" {

			if err := r.SetQueryParam("phase", qPhase); err != nil {
				return err
			}
		}
	}

	if o.Phasen != nil {

		// query param phase__n
		var qrPhasen string

		if o.Phasen != nil {
			qrPhasen = *o.Phasen
		}
		qPhasen := qrPhasen
		if qPhasen != "" {

			if err := r.SetQueryParam("phase__n", qPhasen); err != nil {
				return err
			}
		}
	}

	if o.PowerPanelID != nil {

		// query param power_panel_id
		var qrPowerPanelID string

		if o.PowerPanelID != nil {
			qrPowerPanelID = *o.PowerPanelID
		}
		qPowerPanelID := qrPowerPanelID
		if qPowerPanelID != "" {

			if err := r.SetQueryParam("power_panel_id", qPowerPanelID); err != nil {
				return err
			}
		}
	}

	if o.PowerPanelIDn != nil {

		// query param power_panel_id__n
		var qrPowerPanelIDn string

		if o.PowerPanelIDn != nil {
			qrPowerPanelIDn = *o.PowerPanelIDn
		}
		qPowerPanelIDn := qrPowerPanelIDn
		if qPowerPanelIDn != "" {

			if err := r.SetQueryParam("power_panel_id__n", qPowerPanelIDn); err != nil {
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

	if o.Supply != nil {

		// query param supply
		var qrSupply string

		if o.Supply != nil {
			qrSupply = *o.Supply
		}
		qSupply := qrSupply
		if qSupply != "" {

			if err := r.SetQueryParam("supply", qSupply); err != nil {
				return err
			}
		}
	}

	if o.Supplyn != nil {

		// query param supply__n
		var qrSupplyn string

		if o.Supplyn != nil {
			qrSupplyn = *o.Supplyn
		}
		qSupplyn := qrSupplyn
		if qSupplyn != "" {

			if err := r.SetQueryParam("supply__n", qSupplyn); err != nil {
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

	if o.Voltage != nil {

		// query param voltage
		var qrVoltage string

		if o.Voltage != nil {
			qrVoltage = *o.Voltage
		}
		qVoltage := qrVoltage
		if qVoltage != "" {

			if err := r.SetQueryParam("voltage", qVoltage); err != nil {
				return err
			}
		}
	}

	if o.VoltageGt != nil {

		// query param voltage__gt
		var qrVoltageGt string

		if o.VoltageGt != nil {
			qrVoltageGt = *o.VoltageGt
		}
		qVoltageGt := qrVoltageGt
		if qVoltageGt != "" {

			if err := r.SetQueryParam("voltage__gt", qVoltageGt); err != nil {
				return err
			}
		}
	}

	if o.VoltageGte != nil {

		// query param voltage__gte
		var qrVoltageGte string

		if o.VoltageGte != nil {
			qrVoltageGte = *o.VoltageGte
		}
		qVoltageGte := qrVoltageGte
		if qVoltageGte != "" {

			if err := r.SetQueryParam("voltage__gte", qVoltageGte); err != nil {
				return err
			}
		}
	}

	if o.VoltageLt != nil {

		// query param voltage__lt
		var qrVoltageLt string

		if o.VoltageLt != nil {
			qrVoltageLt = *o.VoltageLt
		}
		qVoltageLt := qrVoltageLt
		if qVoltageLt != "" {

			if err := r.SetQueryParam("voltage__lt", qVoltageLt); err != nil {
				return err
			}
		}
	}

	if o.VoltageLte != nil {

		// query param voltage__lte
		var qrVoltageLte string

		if o.VoltageLte != nil {
			qrVoltageLte = *o.VoltageLte
		}
		qVoltageLte := qrVoltageLte
		if qVoltageLte != "" {

			if err := r.SetQueryParam("voltage__lte", qVoltageLte); err != nil {
				return err
			}
		}
	}

	if o.Voltagen != nil {

		// query param voltage__n
		var qrVoltagen string

		if o.Voltagen != nil {
			qrVoltagen = *o.Voltagen
		}
		qVoltagen := qrVoltagen
		if qVoltagen != "" {

			if err := r.SetQueryParam("voltage__n", qVoltagen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
