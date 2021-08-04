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

// NewDcimDeviceTypesListParams creates a new DcimDeviceTypesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimDeviceTypesListParams() *DcimDeviceTypesListParams {
	return &DcimDeviceTypesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceTypesListParamsWithTimeout creates a new DcimDeviceTypesListParams object
// with the ability to set a timeout on a request.
func NewDcimDeviceTypesListParamsWithTimeout(timeout time.Duration) *DcimDeviceTypesListParams {
	return &DcimDeviceTypesListParams{
		timeout: timeout,
	}
}

// NewDcimDeviceTypesListParamsWithContext creates a new DcimDeviceTypesListParams object
// with the ability to set a context for a request.
func NewDcimDeviceTypesListParamsWithContext(ctx context.Context) *DcimDeviceTypesListParams {
	return &DcimDeviceTypesListParams{
		Context: ctx,
	}
}

// NewDcimDeviceTypesListParamsWithHTTPClient creates a new DcimDeviceTypesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimDeviceTypesListParamsWithHTTPClient(client *http.Client) *DcimDeviceTypesListParams {
	return &DcimDeviceTypesListParams{
		HTTPClient: client,
	}
}

/* DcimDeviceTypesListParams contains all the parameters to send to the API endpoint
   for the dcim device types list operation.

   Typically these are written to a http.Request.
*/
type DcimDeviceTypesListParams struct {

	// ConsolePorts.
	ConsolePorts *string

	// ConsoleServerPorts.
	ConsoleServerPorts *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DeviceBays.
	DeviceBays *string

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

	// Interfaces.
	Interfaces *string

	// IsFullDepth.
	IsFullDepth *string

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

	// Model.
	Model *string

	// ModelIc.
	ModelIc *string

	// ModelIe.
	ModelIe *string

	// ModelIew.
	ModelIew *string

	// ModelIsw.
	ModelIsw *string

	// Modeln.
	Modeln *string

	// ModelNic.
	ModelNic *string

	// ModelNie.
	ModelNie *string

	// ModelNiew.
	ModelNiew *string

	// ModelNisw.
	ModelNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// PartNumber.
	PartNumber *string

	// PartNumberIc.
	PartNumberIc *string

	// PartNumberIe.
	PartNumberIe *string

	// PartNumberIew.
	PartNumberIew *string

	// PartNumberIsw.
	PartNumberIsw *string

	// PartNumbern.
	PartNumbern *string

	// PartNumberNic.
	PartNumberNic *string

	// PartNumberNie.
	PartNumberNie *string

	// PartNumberNiew.
	PartNumberNiew *string

	// PartNumberNisw.
	PartNumberNisw *string

	// PassThroughPorts.
	PassThroughPorts *string

	// PowerOutlets.
	PowerOutlets *string

	// PowerPorts.
	PowerPorts *string

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

	// SubdeviceRole.
	SubdeviceRole *string

	// SubdeviceRolen.
	SubdeviceRolen *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim device types list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesListParams) WithDefaults() *DcimDeviceTypesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim device types list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimDeviceTypesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithTimeout(timeout time.Duration) *DcimDeviceTypesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithContext(ctx context.Context) *DcimDeviceTypesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithHTTPClient(client *http.Client) *DcimDeviceTypesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsolePorts adds the consolePorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithConsolePorts(consolePorts *string) *DcimDeviceTypesListParams {
	o.SetConsolePorts(consolePorts)
	return o
}

// SetConsolePorts adds the consolePorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetConsolePorts(consolePorts *string) {
	o.ConsolePorts = consolePorts
}

// WithConsoleServerPorts adds the consoleServerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithConsoleServerPorts(consoleServerPorts *string) *DcimDeviceTypesListParams {
	o.SetConsoleServerPorts(consoleServerPorts)
	return o
}

// SetConsoleServerPorts adds the consoleServerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetConsoleServerPorts(consoleServerPorts *string) {
	o.ConsoleServerPorts = consoleServerPorts
}

// WithCreated adds the created to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithCreated(created *string) *DcimDeviceTypesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithCreatedGte(createdGte *string) *DcimDeviceTypesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithCreatedLte(createdLte *string) *DcimDeviceTypesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDeviceBays adds the deviceBays to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithDeviceBays(deviceBays *string) *DcimDeviceTypesListParams {
	o.SetDeviceBays(deviceBays)
	return o
}

// SetDeviceBays adds the deviceBays to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetDeviceBays(deviceBays *string) {
	o.DeviceBays = deviceBays
}

// WithID adds the id to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithID(id *string) *DcimDeviceTypesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDIc(iDIc *string) *DcimDeviceTypesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDIe(iDIe *string) *DcimDeviceTypesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDIew(iDIew *string) *DcimDeviceTypesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDIsw(iDIsw *string) *DcimDeviceTypesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDn(iDn *string) *DcimDeviceTypesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDNic(iDNic *string) *DcimDeviceTypesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDNie(iDNie *string) *DcimDeviceTypesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDNiew(iDNiew *string) *DcimDeviceTypesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIDNisw(iDNisw *string) *DcimDeviceTypesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithInterfaces adds the interfaces to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithInterfaces(interfaces *string) *DcimDeviceTypesListParams {
	o.SetInterfaces(interfaces)
	return o
}

// SetInterfaces adds the interfaces to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetInterfaces(interfaces *string) {
	o.Interfaces = interfaces
}

// WithIsFullDepth adds the isFullDepth to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithIsFullDepth(isFullDepth *string) *DcimDeviceTypesListParams {
	o.SetIsFullDepth(isFullDepth)
	return o
}

// SetIsFullDepth adds the isFullDepth to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetIsFullDepth(isFullDepth *string) {
	o.IsFullDepth = isFullDepth
}

// WithLastUpdated adds the lastUpdated to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithLastUpdated(lastUpdated *string) *DcimDeviceTypesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimDeviceTypesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimDeviceTypesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithLimit(limit *int64) *DcimDeviceTypesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithManufacturer adds the manufacturer to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturer(manufacturer *string) *DcimDeviceTypesListParams {
	o.SetManufacturer(manufacturer)
	return o
}

// SetManufacturer adds the manufacturer to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturer(manufacturer *string) {
	o.Manufacturer = manufacturer
}

// WithManufacturern adds the manufacturern to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturern(manufacturern *string) *DcimDeviceTypesListParams {
	o.SetManufacturern(manufacturern)
	return o
}

// SetManufacturern adds the manufacturerN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturern(manufacturern *string) {
	o.Manufacturern = manufacturern
}

// WithManufacturerID adds the manufacturerID to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturerID(manufacturerID *string) *DcimDeviceTypesListParams {
	o.SetManufacturerID(manufacturerID)
	return o
}

// SetManufacturerID adds the manufacturerId to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturerID(manufacturerID *string) {
	o.ManufacturerID = manufacturerID
}

// WithManufacturerIDn adds the manufacturerIDn to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithManufacturerIDn(manufacturerIDn *string) *DcimDeviceTypesListParams {
	o.SetManufacturerIDn(manufacturerIDn)
	return o
}

// SetManufacturerIDn adds the manufacturerIdN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetManufacturerIDn(manufacturerIDn *string) {
	o.ManufacturerIDn = manufacturerIDn
}

// WithModel adds the model to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModel(model *string) *DcimDeviceTypesListParams {
	o.SetModel(model)
	return o
}

// SetModel adds the model to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModel(model *string) {
	o.Model = model
}

// WithModelIc adds the modelIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelIc(modelIc *string) *DcimDeviceTypesListParams {
	o.SetModelIc(modelIc)
	return o
}

// SetModelIc adds the modelIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelIc(modelIc *string) {
	o.ModelIc = modelIc
}

// WithModelIe adds the modelIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelIe(modelIe *string) *DcimDeviceTypesListParams {
	o.SetModelIe(modelIe)
	return o
}

// SetModelIe adds the modelIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelIe(modelIe *string) {
	o.ModelIe = modelIe
}

// WithModelIew adds the modelIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelIew(modelIew *string) *DcimDeviceTypesListParams {
	o.SetModelIew(modelIew)
	return o
}

// SetModelIew adds the modelIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelIew(modelIew *string) {
	o.ModelIew = modelIew
}

// WithModelIsw adds the modelIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelIsw(modelIsw *string) *DcimDeviceTypesListParams {
	o.SetModelIsw(modelIsw)
	return o
}

// SetModelIsw adds the modelIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelIsw(modelIsw *string) {
	o.ModelIsw = modelIsw
}

// WithModeln adds the modeln to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModeln(modeln *string) *DcimDeviceTypesListParams {
	o.SetModeln(modeln)
	return o
}

// SetModeln adds the modelN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModeln(modeln *string) {
	o.Modeln = modeln
}

// WithModelNic adds the modelNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelNic(modelNic *string) *DcimDeviceTypesListParams {
	o.SetModelNic(modelNic)
	return o
}

// SetModelNic adds the modelNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelNic(modelNic *string) {
	o.ModelNic = modelNic
}

// WithModelNie adds the modelNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelNie(modelNie *string) *DcimDeviceTypesListParams {
	o.SetModelNie(modelNie)
	return o
}

// SetModelNie adds the modelNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelNie(modelNie *string) {
	o.ModelNie = modelNie
}

// WithModelNiew adds the modelNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelNiew(modelNiew *string) *DcimDeviceTypesListParams {
	o.SetModelNiew(modelNiew)
	return o
}

// SetModelNiew adds the modelNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelNiew(modelNiew *string) {
	o.ModelNiew = modelNiew
}

// WithModelNisw adds the modelNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithModelNisw(modelNisw *string) *DcimDeviceTypesListParams {
	o.SetModelNisw(modelNisw)
	return o
}

// SetModelNisw adds the modelNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetModelNisw(modelNisw *string) {
	o.ModelNisw = modelNisw
}

// WithOffset adds the offset to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithOffset(offset *int64) *DcimDeviceTypesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPartNumber adds the partNumber to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumber(partNumber *string) *DcimDeviceTypesListParams {
	o.SetPartNumber(partNumber)
	return o
}

// SetPartNumber adds the partNumber to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumber(partNumber *string) {
	o.PartNumber = partNumber
}

// WithPartNumberIc adds the partNumberIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberIc(partNumberIc *string) *DcimDeviceTypesListParams {
	o.SetPartNumberIc(partNumberIc)
	return o
}

// SetPartNumberIc adds the partNumberIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberIc(partNumberIc *string) {
	o.PartNumberIc = partNumberIc
}

// WithPartNumberIe adds the partNumberIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberIe(partNumberIe *string) *DcimDeviceTypesListParams {
	o.SetPartNumberIe(partNumberIe)
	return o
}

// SetPartNumberIe adds the partNumberIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberIe(partNumberIe *string) {
	o.PartNumberIe = partNumberIe
}

// WithPartNumberIew adds the partNumberIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberIew(partNumberIew *string) *DcimDeviceTypesListParams {
	o.SetPartNumberIew(partNumberIew)
	return o
}

// SetPartNumberIew adds the partNumberIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberIew(partNumberIew *string) {
	o.PartNumberIew = partNumberIew
}

// WithPartNumberIsw adds the partNumberIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberIsw(partNumberIsw *string) *DcimDeviceTypesListParams {
	o.SetPartNumberIsw(partNumberIsw)
	return o
}

// SetPartNumberIsw adds the partNumberIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberIsw(partNumberIsw *string) {
	o.PartNumberIsw = partNumberIsw
}

// WithPartNumbern adds the partNumbern to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumbern(partNumbern *string) *DcimDeviceTypesListParams {
	o.SetPartNumbern(partNumbern)
	return o
}

// SetPartNumbern adds the partNumberN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumbern(partNumbern *string) {
	o.PartNumbern = partNumbern
}

// WithPartNumberNic adds the partNumberNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberNic(partNumberNic *string) *DcimDeviceTypesListParams {
	o.SetPartNumberNic(partNumberNic)
	return o
}

// SetPartNumberNic adds the partNumberNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberNic(partNumberNic *string) {
	o.PartNumberNic = partNumberNic
}

// WithPartNumberNie adds the partNumberNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberNie(partNumberNie *string) *DcimDeviceTypesListParams {
	o.SetPartNumberNie(partNumberNie)
	return o
}

// SetPartNumberNie adds the partNumberNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberNie(partNumberNie *string) {
	o.PartNumberNie = partNumberNie
}

// WithPartNumberNiew adds the partNumberNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberNiew(partNumberNiew *string) *DcimDeviceTypesListParams {
	o.SetPartNumberNiew(partNumberNiew)
	return o
}

// SetPartNumberNiew adds the partNumberNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberNiew(partNumberNiew *string) {
	o.PartNumberNiew = partNumberNiew
}

// WithPartNumberNisw adds the partNumberNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPartNumberNisw(partNumberNisw *string) *DcimDeviceTypesListParams {
	o.SetPartNumberNisw(partNumberNisw)
	return o
}

// SetPartNumberNisw adds the partNumberNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPartNumberNisw(partNumberNisw *string) {
	o.PartNumberNisw = partNumberNisw
}

// WithPassThroughPorts adds the passThroughPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPassThroughPorts(passThroughPorts *string) *DcimDeviceTypesListParams {
	o.SetPassThroughPorts(passThroughPorts)
	return o
}

// SetPassThroughPorts adds the passThroughPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPassThroughPorts(passThroughPorts *string) {
	o.PassThroughPorts = passThroughPorts
}

// WithPowerOutlets adds the powerOutlets to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPowerOutlets(powerOutlets *string) *DcimDeviceTypesListParams {
	o.SetPowerOutlets(powerOutlets)
	return o
}

// SetPowerOutlets adds the powerOutlets to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPowerOutlets(powerOutlets *string) {
	o.PowerOutlets = powerOutlets
}

// WithPowerPorts adds the powerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithPowerPorts(powerPorts *string) *DcimDeviceTypesListParams {
	o.SetPowerPorts(powerPorts)
	return o
}

// SetPowerPorts adds the powerPorts to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetPowerPorts(powerPorts *string) {
	o.PowerPorts = powerPorts
}

// WithQ adds the q to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithQ(q *string) *DcimDeviceTypesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetQ(q *string) {
	o.Q = q
}

// WithSlug adds the slug to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlug(slug *string) *DcimDeviceTypesListParams {
	o.SetSlug(slug)
	return o
}

// SetSlug adds the slug to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlug(slug *string) {
	o.Slug = slug
}

// WithSlugIc adds the slugIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugIc(slugIc *string) *DcimDeviceTypesListParams {
	o.SetSlugIc(slugIc)
	return o
}

// SetSlugIc adds the slugIc to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugIc(slugIc *string) {
	o.SlugIc = slugIc
}

// WithSlugIe adds the slugIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugIe(slugIe *string) *DcimDeviceTypesListParams {
	o.SetSlugIe(slugIe)
	return o
}

// SetSlugIe adds the slugIe to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugIe(slugIe *string) {
	o.SlugIe = slugIe
}

// WithSlugIew adds the slugIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugIew(slugIew *string) *DcimDeviceTypesListParams {
	o.SetSlugIew(slugIew)
	return o
}

// SetSlugIew adds the slugIew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugIew(slugIew *string) {
	o.SlugIew = slugIew
}

// WithSlugIsw adds the slugIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugIsw(slugIsw *string) *DcimDeviceTypesListParams {
	o.SetSlugIsw(slugIsw)
	return o
}

// SetSlugIsw adds the slugIsw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugIsw(slugIsw *string) {
	o.SlugIsw = slugIsw
}

// WithSlugn adds the slugn to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugn(slugn *string) *DcimDeviceTypesListParams {
	o.SetSlugn(slugn)
	return o
}

// SetSlugn adds the slugN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugn(slugn *string) {
	o.Slugn = slugn
}

// WithSlugNic adds the slugNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugNic(slugNic *string) *DcimDeviceTypesListParams {
	o.SetSlugNic(slugNic)
	return o
}

// SetSlugNic adds the slugNic to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugNic(slugNic *string) {
	o.SlugNic = slugNic
}

// WithSlugNie adds the slugNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugNie(slugNie *string) *DcimDeviceTypesListParams {
	o.SetSlugNie(slugNie)
	return o
}

// SetSlugNie adds the slugNie to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugNie(slugNie *string) {
	o.SlugNie = slugNie
}

// WithSlugNiew adds the slugNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugNiew(slugNiew *string) *DcimDeviceTypesListParams {
	o.SetSlugNiew(slugNiew)
	return o
}

// SetSlugNiew adds the slugNiew to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugNiew(slugNiew *string) {
	o.SlugNiew = slugNiew
}

// WithSlugNisw adds the slugNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSlugNisw(slugNisw *string) *DcimDeviceTypesListParams {
	o.SetSlugNisw(slugNisw)
	return o
}

// SetSlugNisw adds the slugNisw to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSlugNisw(slugNisw *string) {
	o.SlugNisw = slugNisw
}

// WithSubdeviceRole adds the subdeviceRole to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSubdeviceRole(subdeviceRole *string) *DcimDeviceTypesListParams {
	o.SetSubdeviceRole(subdeviceRole)
	return o
}

// SetSubdeviceRole adds the subdeviceRole to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSubdeviceRole(subdeviceRole *string) {
	o.SubdeviceRole = subdeviceRole
}

// WithSubdeviceRolen adds the subdeviceRolen to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithSubdeviceRolen(subdeviceRolen *string) *DcimDeviceTypesListParams {
	o.SetSubdeviceRolen(subdeviceRolen)
	return o
}

// SetSubdeviceRolen adds the subdeviceRoleN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetSubdeviceRolen(subdeviceRolen *string) {
	o.SubdeviceRolen = subdeviceRolen
}

// WithTag adds the tag to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithTag(tag *string) *DcimDeviceTypesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithTagn(tagn *string) *DcimDeviceTypesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithUHeight adds the uHeight to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeight(uHeight *string) *DcimDeviceTypesListParams {
	o.SetUHeight(uHeight)
	return o
}

// SetUHeight adds the uHeight to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeight(uHeight *string) {
	o.UHeight = uHeight
}

// WithUHeightGt adds the uHeightGt to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeightGt(uHeightGt *string) *DcimDeviceTypesListParams {
	o.SetUHeightGt(uHeightGt)
	return o
}

// SetUHeightGt adds the uHeightGt to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeightGt(uHeightGt *string) {
	o.UHeightGt = uHeightGt
}

// WithUHeightGte adds the uHeightGte to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeightGte(uHeightGte *string) *DcimDeviceTypesListParams {
	o.SetUHeightGte(uHeightGte)
	return o
}

// SetUHeightGte adds the uHeightGte to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeightGte(uHeightGte *string) {
	o.UHeightGte = uHeightGte
}

// WithUHeightLt adds the uHeightLt to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeightLt(uHeightLt *string) *DcimDeviceTypesListParams {
	o.SetUHeightLt(uHeightLt)
	return o
}

// SetUHeightLt adds the uHeightLt to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeightLt(uHeightLt *string) {
	o.UHeightLt = uHeightLt
}

// WithUHeightLte adds the uHeightLte to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeightLte(uHeightLte *string) *DcimDeviceTypesListParams {
	o.SetUHeightLte(uHeightLte)
	return o
}

// SetUHeightLte adds the uHeightLte to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeightLte(uHeightLte *string) {
	o.UHeightLte = uHeightLte
}

// WithUHeightn adds the uHeightn to the dcim device types list params
func (o *DcimDeviceTypesListParams) WithUHeightn(uHeightn *string) *DcimDeviceTypesListParams {
	o.SetUHeightn(uHeightn)
	return o
}

// SetUHeightn adds the uHeightN to the dcim device types list params
func (o *DcimDeviceTypesListParams) SetUHeightn(uHeightn *string) {
	o.UHeightn = uHeightn
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceTypesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConsolePorts != nil {

		// query param console_ports
		var qrConsolePorts string

		if o.ConsolePorts != nil {
			qrConsolePorts = *o.ConsolePorts
		}
		qConsolePorts := qrConsolePorts
		if qConsolePorts != "" {

			if err := r.SetQueryParam("console_ports", qConsolePorts); err != nil {
				return err
			}
		}
	}

	if o.ConsoleServerPorts != nil {

		// query param console_server_ports
		var qrConsoleServerPorts string

		if o.ConsoleServerPorts != nil {
			qrConsoleServerPorts = *o.ConsoleServerPorts
		}
		qConsoleServerPorts := qrConsoleServerPorts
		if qConsoleServerPorts != "" {

			if err := r.SetQueryParam("console_server_ports", qConsoleServerPorts); err != nil {
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

	if o.DeviceBays != nil {

		// query param device_bays
		var qrDeviceBays string

		if o.DeviceBays != nil {
			qrDeviceBays = *o.DeviceBays
		}
		qDeviceBays := qrDeviceBays
		if qDeviceBays != "" {

			if err := r.SetQueryParam("device_bays", qDeviceBays); err != nil {
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

	if o.Interfaces != nil {

		// query param interfaces
		var qrInterfaces string

		if o.Interfaces != nil {
			qrInterfaces = *o.Interfaces
		}
		qInterfaces := qrInterfaces
		if qInterfaces != "" {

			if err := r.SetQueryParam("interfaces", qInterfaces); err != nil {
				return err
			}
		}
	}

	if o.IsFullDepth != nil {

		// query param is_full_depth
		var qrIsFullDepth string

		if o.IsFullDepth != nil {
			qrIsFullDepth = *o.IsFullDepth
		}
		qIsFullDepth := qrIsFullDepth
		if qIsFullDepth != "" {

			if err := r.SetQueryParam("is_full_depth", qIsFullDepth); err != nil {
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

	if o.Model != nil {

		// query param model
		var qrModel string

		if o.Model != nil {
			qrModel = *o.Model
		}
		qModel := qrModel
		if qModel != "" {

			if err := r.SetQueryParam("model", qModel); err != nil {
				return err
			}
		}
	}

	if o.ModelIc != nil {

		// query param model__ic
		var qrModelIc string

		if o.ModelIc != nil {
			qrModelIc = *o.ModelIc
		}
		qModelIc := qrModelIc
		if qModelIc != "" {

			if err := r.SetQueryParam("model__ic", qModelIc); err != nil {
				return err
			}
		}
	}

	if o.ModelIe != nil {

		// query param model__ie
		var qrModelIe string

		if o.ModelIe != nil {
			qrModelIe = *o.ModelIe
		}
		qModelIe := qrModelIe
		if qModelIe != "" {

			if err := r.SetQueryParam("model__ie", qModelIe); err != nil {
				return err
			}
		}
	}

	if o.ModelIew != nil {

		// query param model__iew
		var qrModelIew string

		if o.ModelIew != nil {
			qrModelIew = *o.ModelIew
		}
		qModelIew := qrModelIew
		if qModelIew != "" {

			if err := r.SetQueryParam("model__iew", qModelIew); err != nil {
				return err
			}
		}
	}

	if o.ModelIsw != nil {

		// query param model__isw
		var qrModelIsw string

		if o.ModelIsw != nil {
			qrModelIsw = *o.ModelIsw
		}
		qModelIsw := qrModelIsw
		if qModelIsw != "" {

			if err := r.SetQueryParam("model__isw", qModelIsw); err != nil {
				return err
			}
		}
	}

	if o.Modeln != nil {

		// query param model__n
		var qrModeln string

		if o.Modeln != nil {
			qrModeln = *o.Modeln
		}
		qModeln := qrModeln
		if qModeln != "" {

			if err := r.SetQueryParam("model__n", qModeln); err != nil {
				return err
			}
		}
	}

	if o.ModelNic != nil {

		// query param model__nic
		var qrModelNic string

		if o.ModelNic != nil {
			qrModelNic = *o.ModelNic
		}
		qModelNic := qrModelNic
		if qModelNic != "" {

			if err := r.SetQueryParam("model__nic", qModelNic); err != nil {
				return err
			}
		}
	}

	if o.ModelNie != nil {

		// query param model__nie
		var qrModelNie string

		if o.ModelNie != nil {
			qrModelNie = *o.ModelNie
		}
		qModelNie := qrModelNie
		if qModelNie != "" {

			if err := r.SetQueryParam("model__nie", qModelNie); err != nil {
				return err
			}
		}
	}

	if o.ModelNiew != nil {

		// query param model__niew
		var qrModelNiew string

		if o.ModelNiew != nil {
			qrModelNiew = *o.ModelNiew
		}
		qModelNiew := qrModelNiew
		if qModelNiew != "" {

			if err := r.SetQueryParam("model__niew", qModelNiew); err != nil {
				return err
			}
		}
	}

	if o.ModelNisw != nil {

		// query param model__nisw
		var qrModelNisw string

		if o.ModelNisw != nil {
			qrModelNisw = *o.ModelNisw
		}
		qModelNisw := qrModelNisw
		if qModelNisw != "" {

			if err := r.SetQueryParam("model__nisw", qModelNisw); err != nil {
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

	if o.PartNumber != nil {

		// query param part_number
		var qrPartNumber string

		if o.PartNumber != nil {
			qrPartNumber = *o.PartNumber
		}
		qPartNumber := qrPartNumber
		if qPartNumber != "" {

			if err := r.SetQueryParam("part_number", qPartNumber); err != nil {
				return err
			}
		}
	}

	if o.PartNumberIc != nil {

		// query param part_number__ic
		var qrPartNumberIc string

		if o.PartNumberIc != nil {
			qrPartNumberIc = *o.PartNumberIc
		}
		qPartNumberIc := qrPartNumberIc
		if qPartNumberIc != "" {

			if err := r.SetQueryParam("part_number__ic", qPartNumberIc); err != nil {
				return err
			}
		}
	}

	if o.PartNumberIe != nil {

		// query param part_number__ie
		var qrPartNumberIe string

		if o.PartNumberIe != nil {
			qrPartNumberIe = *o.PartNumberIe
		}
		qPartNumberIe := qrPartNumberIe
		if qPartNumberIe != "" {

			if err := r.SetQueryParam("part_number__ie", qPartNumberIe); err != nil {
				return err
			}
		}
	}

	if o.PartNumberIew != nil {

		// query param part_number__iew
		var qrPartNumberIew string

		if o.PartNumberIew != nil {
			qrPartNumberIew = *o.PartNumberIew
		}
		qPartNumberIew := qrPartNumberIew
		if qPartNumberIew != "" {

			if err := r.SetQueryParam("part_number__iew", qPartNumberIew); err != nil {
				return err
			}
		}
	}

	if o.PartNumberIsw != nil {

		// query param part_number__isw
		var qrPartNumberIsw string

		if o.PartNumberIsw != nil {
			qrPartNumberIsw = *o.PartNumberIsw
		}
		qPartNumberIsw := qrPartNumberIsw
		if qPartNumberIsw != "" {

			if err := r.SetQueryParam("part_number__isw", qPartNumberIsw); err != nil {
				return err
			}
		}
	}

	if o.PartNumbern != nil {

		// query param part_number__n
		var qrPartNumbern string

		if o.PartNumbern != nil {
			qrPartNumbern = *o.PartNumbern
		}
		qPartNumbern := qrPartNumbern
		if qPartNumbern != "" {

			if err := r.SetQueryParam("part_number__n", qPartNumbern); err != nil {
				return err
			}
		}
	}

	if o.PartNumberNic != nil {

		// query param part_number__nic
		var qrPartNumberNic string

		if o.PartNumberNic != nil {
			qrPartNumberNic = *o.PartNumberNic
		}
		qPartNumberNic := qrPartNumberNic
		if qPartNumberNic != "" {

			if err := r.SetQueryParam("part_number__nic", qPartNumberNic); err != nil {
				return err
			}
		}
	}

	if o.PartNumberNie != nil {

		// query param part_number__nie
		var qrPartNumberNie string

		if o.PartNumberNie != nil {
			qrPartNumberNie = *o.PartNumberNie
		}
		qPartNumberNie := qrPartNumberNie
		if qPartNumberNie != "" {

			if err := r.SetQueryParam("part_number__nie", qPartNumberNie); err != nil {
				return err
			}
		}
	}

	if o.PartNumberNiew != nil {

		// query param part_number__niew
		var qrPartNumberNiew string

		if o.PartNumberNiew != nil {
			qrPartNumberNiew = *o.PartNumberNiew
		}
		qPartNumberNiew := qrPartNumberNiew
		if qPartNumberNiew != "" {

			if err := r.SetQueryParam("part_number__niew", qPartNumberNiew); err != nil {
				return err
			}
		}
	}

	if o.PartNumberNisw != nil {

		// query param part_number__nisw
		var qrPartNumberNisw string

		if o.PartNumberNisw != nil {
			qrPartNumberNisw = *o.PartNumberNisw
		}
		qPartNumberNisw := qrPartNumberNisw
		if qPartNumberNisw != "" {

			if err := r.SetQueryParam("part_number__nisw", qPartNumberNisw); err != nil {
				return err
			}
		}
	}

	if o.PassThroughPorts != nil {

		// query param pass_through_ports
		var qrPassThroughPorts string

		if o.PassThroughPorts != nil {
			qrPassThroughPorts = *o.PassThroughPorts
		}
		qPassThroughPorts := qrPassThroughPorts
		if qPassThroughPorts != "" {

			if err := r.SetQueryParam("pass_through_ports", qPassThroughPorts); err != nil {
				return err
			}
		}
	}

	if o.PowerOutlets != nil {

		// query param power_outlets
		var qrPowerOutlets string

		if o.PowerOutlets != nil {
			qrPowerOutlets = *o.PowerOutlets
		}
		qPowerOutlets := qrPowerOutlets
		if qPowerOutlets != "" {

			if err := r.SetQueryParam("power_outlets", qPowerOutlets); err != nil {
				return err
			}
		}
	}

	if o.PowerPorts != nil {

		// query param power_ports
		var qrPowerPorts string

		if o.PowerPorts != nil {
			qrPowerPorts = *o.PowerPorts
		}
		qPowerPorts := qrPowerPorts
		if qPowerPorts != "" {

			if err := r.SetQueryParam("power_ports", qPowerPorts); err != nil {
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

	if o.SubdeviceRole != nil {

		// query param subdevice_role
		var qrSubdeviceRole string

		if o.SubdeviceRole != nil {
			qrSubdeviceRole = *o.SubdeviceRole
		}
		qSubdeviceRole := qrSubdeviceRole
		if qSubdeviceRole != "" {

			if err := r.SetQueryParam("subdevice_role", qSubdeviceRole); err != nil {
				return err
			}
		}
	}

	if o.SubdeviceRolen != nil {

		// query param subdevice_role__n
		var qrSubdeviceRolen string

		if o.SubdeviceRolen != nil {
			qrSubdeviceRolen = *o.SubdeviceRolen
		}
		qSubdeviceRolen := qrSubdeviceRolen
		if qSubdeviceRolen != "" {

			if err := r.SetQueryParam("subdevice_role__n", qSubdeviceRolen); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
