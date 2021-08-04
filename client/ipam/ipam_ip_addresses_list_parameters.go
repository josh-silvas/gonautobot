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

// NewIpamIPAddressesListParams creates a new IpamIPAddressesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamIPAddressesListParams() *IpamIPAddressesListParams {
	return &IpamIPAddressesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamIPAddressesListParamsWithTimeout creates a new IpamIPAddressesListParams object
// with the ability to set a timeout on a request.
func NewIpamIPAddressesListParamsWithTimeout(timeout time.Duration) *IpamIPAddressesListParams {
	return &IpamIPAddressesListParams{
		timeout: timeout,
	}
}

// NewIpamIPAddressesListParamsWithContext creates a new IpamIPAddressesListParams object
// with the ability to set a context for a request.
func NewIpamIPAddressesListParamsWithContext(ctx context.Context) *IpamIPAddressesListParams {
	return &IpamIPAddressesListParams{
		Context: ctx,
	}
}

// NewIpamIPAddressesListParamsWithHTTPClient creates a new IpamIPAddressesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamIPAddressesListParamsWithHTTPClient(client *http.Client) *IpamIPAddressesListParams {
	return &IpamIPAddressesListParams{
		HTTPClient: client,
	}
}

/* IpamIPAddressesListParams contains all the parameters to send to the API endpoint
   for the ipam ip addresses list operation.

   Typically these are written to a http.Request.
*/
type IpamIPAddressesListParams struct {

	// Address.
	Address *string

	// AssignedToInterface.
	AssignedToInterface *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Device.
	Device *string

	// DeviceID.
	DeviceID *string

	// DNSName.
	DNSName *string

	// DNSNameIc.
	DNSNameIc *string

	// DNSNameIe.
	DNSNameIe *string

	// DNSNameIew.
	DNSNameIew *string

	// DNSNameIsw.
	DNSNameIsw *string

	// DNSNamen.
	DNSNamen *string

	// DNSNameNic.
	DNSNameNic *string

	// DNSNameNie.
	DNSNameNie *string

	// DNSNameNiew.
	DNSNameNiew *string

	// DNSNameNisw.
	DNSNameNisw *string

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

	// Interface.
	Interface *string

	// Interfacen.
	Interfacen *string

	// InterfaceID.
	InterfaceID *string

	// InterfaceIDn.
	InterfaceIDn *string

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

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Parent.
	Parent *string

	// PresentInVrf.
	PresentInVrf *string

	// PresentInVrfID.
	PresentInVrfID *string

	// Q.
	Q *string

	// Role.
	Role *string

	// Rolen.
	Rolen *string

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

	// VirtualMachine.
	VirtualMachine *string

	// VirtualMachineID.
	VirtualMachineID *string

	// Vminterface.
	Vminterface *string

	// Vminterfacen.
	Vminterfacen *string

	// VminterfaceID.
	VminterfaceID *string

	// VminterfaceIDn.
	VminterfaceIDn *string

	// Vrf.
	Vrf *string

	// Vrfn.
	Vrfn *string

	// VrfID.
	VrfID *string

	// VrfIDn.
	VrfIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam ip addresses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesListParams) WithDefaults() *IpamIPAddressesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam ip addresses list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamIPAddressesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTimeout(timeout time.Duration) *IpamIPAddressesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithContext(ctx context.Context) *IpamIPAddressesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithHTTPClient(client *http.Client) *IpamIPAddressesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithAddress(address *string) *IpamIPAddressesListParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetAddress(address *string) {
	o.Address = address
}

// WithAssignedToInterface adds the assignedToInterface to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithAssignedToInterface(assignedToInterface *string) *IpamIPAddressesListParams {
	o.SetAssignedToInterface(assignedToInterface)
	return o
}

// SetAssignedToInterface adds the assignedToInterface to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetAssignedToInterface(assignedToInterface *string) {
	o.AssignedToInterface = assignedToInterface
}

// WithCreated adds the created to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithCreated(created *string) *IpamIPAddressesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithCreatedGte(createdGte *string) *IpamIPAddressesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithCreatedLte(createdLte *string) *IpamIPAddressesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevice adds the device to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDevice(device *string) *IpamIPAddressesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDeviceID adds the deviceID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDeviceID(deviceID *string) *IpamIPAddressesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDNSName adds the dNSName to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSName(dNSName *string) *IpamIPAddressesListParams {
	o.SetDNSName(dNSName)
	return o
}

// SetDNSName adds the dnsName to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSName(dNSName *string) {
	o.DNSName = dNSName
}

// WithDNSNameIc adds the dNSNameIc to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameIc(dNSNameIc *string) *IpamIPAddressesListParams {
	o.SetDNSNameIc(dNSNameIc)
	return o
}

// SetDNSNameIc adds the dnsNameIc to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameIc(dNSNameIc *string) {
	o.DNSNameIc = dNSNameIc
}

// WithDNSNameIe adds the dNSNameIe to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameIe(dNSNameIe *string) *IpamIPAddressesListParams {
	o.SetDNSNameIe(dNSNameIe)
	return o
}

// SetDNSNameIe adds the dnsNameIe to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameIe(dNSNameIe *string) {
	o.DNSNameIe = dNSNameIe
}

// WithDNSNameIew adds the dNSNameIew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameIew(dNSNameIew *string) *IpamIPAddressesListParams {
	o.SetDNSNameIew(dNSNameIew)
	return o
}

// SetDNSNameIew adds the dnsNameIew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameIew(dNSNameIew *string) {
	o.DNSNameIew = dNSNameIew
}

// WithDNSNameIsw adds the dNSNameIsw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameIsw(dNSNameIsw *string) *IpamIPAddressesListParams {
	o.SetDNSNameIsw(dNSNameIsw)
	return o
}

// SetDNSNameIsw adds the dnsNameIsw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameIsw(dNSNameIsw *string) {
	o.DNSNameIsw = dNSNameIsw
}

// WithDNSNamen adds the dNSNamen to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNamen(dNSNamen *string) *IpamIPAddressesListParams {
	o.SetDNSNamen(dNSNamen)
	return o
}

// SetDNSNamen adds the dnsNameN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNamen(dNSNamen *string) {
	o.DNSNamen = dNSNamen
}

// WithDNSNameNic adds the dNSNameNic to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameNic(dNSNameNic *string) *IpamIPAddressesListParams {
	o.SetDNSNameNic(dNSNameNic)
	return o
}

// SetDNSNameNic adds the dnsNameNic to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameNic(dNSNameNic *string) {
	o.DNSNameNic = dNSNameNic
}

// WithDNSNameNie adds the dNSNameNie to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameNie(dNSNameNie *string) *IpamIPAddressesListParams {
	o.SetDNSNameNie(dNSNameNie)
	return o
}

// SetDNSNameNie adds the dnsNameNie to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameNie(dNSNameNie *string) {
	o.DNSNameNie = dNSNameNie
}

// WithDNSNameNiew adds the dNSNameNiew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameNiew(dNSNameNiew *string) *IpamIPAddressesListParams {
	o.SetDNSNameNiew(dNSNameNiew)
	return o
}

// SetDNSNameNiew adds the dnsNameNiew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameNiew(dNSNameNiew *string) {
	o.DNSNameNiew = dNSNameNiew
}

// WithDNSNameNisw adds the dNSNameNisw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithDNSNameNisw(dNSNameNisw *string) *IpamIPAddressesListParams {
	o.SetDNSNameNisw(dNSNameNisw)
	return o
}

// SetDNSNameNisw adds the dnsNameNisw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetDNSNameNisw(dNSNameNisw *string) {
	o.DNSNameNisw = dNSNameNisw
}

// WithFamily adds the family to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithFamily(family *float64) *IpamIPAddressesListParams {
	o.SetFamily(family)
	return o
}

// SetFamily adds the family to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetFamily(family *float64) {
	o.Family = family
}

// WithID adds the id to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithID(id *string) *IpamIPAddressesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDIc(iDIc *string) *IpamIPAddressesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDIe(iDIe *string) *IpamIPAddressesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDIew(iDIew *string) *IpamIPAddressesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDIsw(iDIsw *string) *IpamIPAddressesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDn(iDn *string) *IpamIPAddressesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDNic(iDNic *string) *IpamIPAddressesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDNie(iDNie *string) *IpamIPAddressesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDNiew(iDNiew *string) *IpamIPAddressesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithIDNisw(iDNisw *string) *IpamIPAddressesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithInterface adds the interfaceVar to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithInterface(interfaceVar *string) *IpamIPAddressesListParams {
	o.SetInterface(interfaceVar)
	return o
}

// SetInterface adds the interface to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetInterface(interfaceVar *string) {
	o.Interface = interfaceVar
}

// WithInterfacen adds the interfacen to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithInterfacen(interfacen *string) *IpamIPAddressesListParams {
	o.SetInterfacen(interfacen)
	return o
}

// SetInterfacen adds the interfaceN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetInterfacen(interfacen *string) {
	o.Interfacen = interfacen
}

// WithInterfaceID adds the interfaceID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithInterfaceID(interfaceID *string) *IpamIPAddressesListParams {
	o.SetInterfaceID(interfaceID)
	return o
}

// SetInterfaceID adds the interfaceId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetInterfaceID(interfaceID *string) {
	o.InterfaceID = interfaceID
}

// WithInterfaceIDn adds the interfaceIDn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithInterfaceIDn(interfaceIDn *string) *IpamIPAddressesListParams {
	o.SetInterfaceIDn(interfaceIDn)
	return o
}

// SetInterfaceIDn adds the interfaceIdN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetInterfaceIDn(interfaceIDn *string) {
	o.InterfaceIDn = interfaceIDn
}

// WithLastUpdated adds the lastUpdated to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithLastUpdated(lastUpdated *string) *IpamIPAddressesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamIPAddressesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamIPAddressesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithLimit(limit *int64) *IpamIPAddressesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithMaskLength adds the maskLength to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithMaskLength(maskLength *float64) *IpamIPAddressesListParams {
	o.SetMaskLength(maskLength)
	return o
}

// SetMaskLength adds the maskLength to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetMaskLength(maskLength *float64) {
	o.MaskLength = maskLength
}

// WithOffset adds the offset to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithOffset(offset *int64) *IpamIPAddressesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithParent adds the parent to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithParent(parent *string) *IpamIPAddressesListParams {
	o.SetParent(parent)
	return o
}

// SetParent adds the parent to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetParent(parent *string) {
	o.Parent = parent
}

// WithPresentInVrf adds the presentInVrf to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithPresentInVrf(presentInVrf *string) *IpamIPAddressesListParams {
	o.SetPresentInVrf(presentInVrf)
	return o
}

// SetPresentInVrf adds the presentInVrf to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetPresentInVrf(presentInVrf *string) {
	o.PresentInVrf = presentInVrf
}

// WithPresentInVrfID adds the presentInVrfID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithPresentInVrfID(presentInVrfID *string) *IpamIPAddressesListParams {
	o.SetPresentInVrfID(presentInVrfID)
	return o
}

// SetPresentInVrfID adds the presentInVrfId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetPresentInVrfID(presentInVrfID *string) {
	o.PresentInVrfID = presentInVrfID
}

// WithQ adds the q to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithQ(q *string) *IpamIPAddressesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetQ(q *string) {
	o.Q = q
}

// WithRole adds the role to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithRole(role *string) *IpamIPAddressesListParams {
	o.SetRole(role)
	return o
}

// SetRole adds the role to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetRole(role *string) {
	o.Role = role
}

// WithRolen adds the rolen to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithRolen(rolen *string) *IpamIPAddressesListParams {
	o.SetRolen(rolen)
	return o
}

// SetRolen adds the roleN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetRolen(rolen *string) {
	o.Rolen = rolen
}

// WithStatus adds the status to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithStatus(status *string) *IpamIPAddressesListParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetStatus(status *string) {
	o.Status = status
}

// WithStatusn adds the statusn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithStatusn(statusn *string) *IpamIPAddressesListParams {
	o.SetStatusn(statusn)
	return o
}

// SetStatusn adds the statusN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetStatusn(statusn *string) {
	o.Statusn = statusn
}

// WithTag adds the tag to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTag(tag *string) *IpamIPAddressesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTagn(tagn *string) *IpamIPAddressesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithTenant adds the tenant to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenant(tenant *string) *IpamIPAddressesListParams {
	o.SetTenant(tenant)
	return o
}

// SetTenant adds the tenant to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenant(tenant *string) {
	o.Tenant = tenant
}

// WithTenantn adds the tenantn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantn(tenantn *string) *IpamIPAddressesListParams {
	o.SetTenantn(tenantn)
	return o
}

// SetTenantn adds the tenantN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantn(tenantn *string) {
	o.Tenantn = tenantn
}

// WithTenantGroup adds the tenantGroup to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantGroup(tenantGroup *string) *IpamIPAddressesListParams {
	o.SetTenantGroup(tenantGroup)
	return o
}

// SetTenantGroup adds the tenantGroup to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantGroup(tenantGroup *string) {
	o.TenantGroup = tenantGroup
}

// WithTenantGroupn adds the tenantGroupn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantGroupn(tenantGroupn *string) *IpamIPAddressesListParams {
	o.SetTenantGroupn(tenantGroupn)
	return o
}

// SetTenantGroupn adds the tenantGroupN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantGroupn(tenantGroupn *string) {
	o.TenantGroupn = tenantGroupn
}

// WithTenantGroupID adds the tenantGroupID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantGroupID(tenantGroupID *string) *IpamIPAddressesListParams {
	o.SetTenantGroupID(tenantGroupID)
	return o
}

// SetTenantGroupID adds the tenantGroupId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantGroupID(tenantGroupID *string) {
	o.TenantGroupID = tenantGroupID
}

// WithTenantGroupIDn adds the tenantGroupIDn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantGroupIDn(tenantGroupIDn *string) *IpamIPAddressesListParams {
	o.SetTenantGroupIDn(tenantGroupIDn)
	return o
}

// SetTenantGroupIDn adds the tenantGroupIdN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantGroupIDn(tenantGroupIDn *string) {
	o.TenantGroupIDn = tenantGroupIDn
}

// WithTenantID adds the tenantID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantID(tenantID *string) *IpamIPAddressesListParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantID(tenantID *string) {
	o.TenantID = tenantID
}

// WithTenantIDn adds the tenantIDn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithTenantIDn(tenantIDn *string) *IpamIPAddressesListParams {
	o.SetTenantIDn(tenantIDn)
	return o
}

// SetTenantIDn adds the tenantIdN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetTenantIDn(tenantIDn *string) {
	o.TenantIDn = tenantIDn
}

// WithVirtualMachine adds the virtualMachine to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVirtualMachine(virtualMachine *string) *IpamIPAddressesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachineID adds the virtualMachineID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVirtualMachineID(virtualMachineID *string) *IpamIPAddressesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVirtualMachineID(virtualMachineID *string) {
	o.VirtualMachineID = virtualMachineID
}

// WithVminterface adds the vminterface to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVminterface(vminterface *string) *IpamIPAddressesListParams {
	o.SetVminterface(vminterface)
	return o
}

// SetVminterface adds the vminterface to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVminterface(vminterface *string) {
	o.Vminterface = vminterface
}

// WithVminterfacen adds the vminterfacen to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVminterfacen(vminterfacen *string) *IpamIPAddressesListParams {
	o.SetVminterfacen(vminterfacen)
	return o
}

// SetVminterfacen adds the vminterfaceN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVminterfacen(vminterfacen *string) {
	o.Vminterfacen = vminterfacen
}

// WithVminterfaceID adds the vminterfaceID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVminterfaceID(vminterfaceID *string) *IpamIPAddressesListParams {
	o.SetVminterfaceID(vminterfaceID)
	return o
}

// SetVminterfaceID adds the vminterfaceId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVminterfaceID(vminterfaceID *string) {
	o.VminterfaceID = vminterfaceID
}

// WithVminterfaceIDn adds the vminterfaceIDn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVminterfaceIDn(vminterfaceIDn *string) *IpamIPAddressesListParams {
	o.SetVminterfaceIDn(vminterfaceIDn)
	return o
}

// SetVminterfaceIDn adds the vminterfaceIdN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVminterfaceIDn(vminterfaceIDn *string) {
	o.VminterfaceIDn = vminterfaceIDn
}

// WithVrf adds the vrf to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVrf(vrf *string) *IpamIPAddressesListParams {
	o.SetVrf(vrf)
	return o
}

// SetVrf adds the vrf to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVrf(vrf *string) {
	o.Vrf = vrf
}

// WithVrfn adds the vrfn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVrfn(vrfn *string) *IpamIPAddressesListParams {
	o.SetVrfn(vrfn)
	return o
}

// SetVrfn adds the vrfN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVrfn(vrfn *string) {
	o.Vrfn = vrfn
}

// WithVrfID adds the vrfID to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVrfID(vrfID *string) *IpamIPAddressesListParams {
	o.SetVrfID(vrfID)
	return o
}

// SetVrfID adds the vrfId to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVrfID(vrfID *string) {
	o.VrfID = vrfID
}

// WithVrfIDn adds the vrfIDn to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) WithVrfIDn(vrfIDn *string) *IpamIPAddressesListParams {
	o.SetVrfIDn(vrfIDn)
	return o
}

// SetVrfIDn adds the vrfIdN to the ipam ip addresses list params
func (o *IpamIPAddressesListParams) SetVrfIDn(vrfIDn *string) {
	o.VrfIDn = vrfIDn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamIPAddressesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {

		// query param address
		var qrAddress string

		if o.Address != nil {
			qrAddress = *o.Address
		}
		qAddress := qrAddress
		if qAddress != "" {

			if err := r.SetQueryParam("address", qAddress); err != nil {
				return err
			}
		}
	}

	if o.AssignedToInterface != nil {

		// query param assigned_to_interface
		var qrAssignedToInterface string

		if o.AssignedToInterface != nil {
			qrAssignedToInterface = *o.AssignedToInterface
		}
		qAssignedToInterface := qrAssignedToInterface
		if qAssignedToInterface != "" {

			if err := r.SetQueryParam("assigned_to_interface", qAssignedToInterface); err != nil {
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

	if o.Device != nil {

		// query param device
		var qrDevice string

		if o.Device != nil {
			qrDevice = *o.Device
		}
		qDevice := qrDevice
		if qDevice != "" {

			if err := r.SetQueryParam("device", qDevice); err != nil {
				return err
			}
		}
	}

	if o.DeviceID != nil {

		// query param device_id
		var qrDeviceID string

		if o.DeviceID != nil {
			qrDeviceID = *o.DeviceID
		}
		qDeviceID := qrDeviceID
		if qDeviceID != "" {

			if err := r.SetQueryParam("device_id", qDeviceID); err != nil {
				return err
			}
		}
	}

	if o.DNSName != nil {

		// query param dns_name
		var qrDNSName string

		if o.DNSName != nil {
			qrDNSName = *o.DNSName
		}
		qDNSName := qrDNSName
		if qDNSName != "" {

			if err := r.SetQueryParam("dns_name", qDNSName); err != nil {
				return err
			}
		}
	}

	if o.DNSNameIc != nil {

		// query param dns_name__ic
		var qrDNSNameIc string

		if o.DNSNameIc != nil {
			qrDNSNameIc = *o.DNSNameIc
		}
		qDNSNameIc := qrDNSNameIc
		if qDNSNameIc != "" {

			if err := r.SetQueryParam("dns_name__ic", qDNSNameIc); err != nil {
				return err
			}
		}
	}

	if o.DNSNameIe != nil {

		// query param dns_name__ie
		var qrDNSNameIe string

		if o.DNSNameIe != nil {
			qrDNSNameIe = *o.DNSNameIe
		}
		qDNSNameIe := qrDNSNameIe
		if qDNSNameIe != "" {

			if err := r.SetQueryParam("dns_name__ie", qDNSNameIe); err != nil {
				return err
			}
		}
	}

	if o.DNSNameIew != nil {

		// query param dns_name__iew
		var qrDNSNameIew string

		if o.DNSNameIew != nil {
			qrDNSNameIew = *o.DNSNameIew
		}
		qDNSNameIew := qrDNSNameIew
		if qDNSNameIew != "" {

			if err := r.SetQueryParam("dns_name__iew", qDNSNameIew); err != nil {
				return err
			}
		}
	}

	if o.DNSNameIsw != nil {

		// query param dns_name__isw
		var qrDNSNameIsw string

		if o.DNSNameIsw != nil {
			qrDNSNameIsw = *o.DNSNameIsw
		}
		qDNSNameIsw := qrDNSNameIsw
		if qDNSNameIsw != "" {

			if err := r.SetQueryParam("dns_name__isw", qDNSNameIsw); err != nil {
				return err
			}
		}
	}

	if o.DNSNamen != nil {

		// query param dns_name__n
		var qrDNSNamen string

		if o.DNSNamen != nil {
			qrDNSNamen = *o.DNSNamen
		}
		qDNSNamen := qrDNSNamen
		if qDNSNamen != "" {

			if err := r.SetQueryParam("dns_name__n", qDNSNamen); err != nil {
				return err
			}
		}
	}

	if o.DNSNameNic != nil {

		// query param dns_name__nic
		var qrDNSNameNic string

		if o.DNSNameNic != nil {
			qrDNSNameNic = *o.DNSNameNic
		}
		qDNSNameNic := qrDNSNameNic
		if qDNSNameNic != "" {

			if err := r.SetQueryParam("dns_name__nic", qDNSNameNic); err != nil {
				return err
			}
		}
	}

	if o.DNSNameNie != nil {

		// query param dns_name__nie
		var qrDNSNameNie string

		if o.DNSNameNie != nil {
			qrDNSNameNie = *o.DNSNameNie
		}
		qDNSNameNie := qrDNSNameNie
		if qDNSNameNie != "" {

			if err := r.SetQueryParam("dns_name__nie", qDNSNameNie); err != nil {
				return err
			}
		}
	}

	if o.DNSNameNiew != nil {

		// query param dns_name__niew
		var qrDNSNameNiew string

		if o.DNSNameNiew != nil {
			qrDNSNameNiew = *o.DNSNameNiew
		}
		qDNSNameNiew := qrDNSNameNiew
		if qDNSNameNiew != "" {

			if err := r.SetQueryParam("dns_name__niew", qDNSNameNiew); err != nil {
				return err
			}
		}
	}

	if o.DNSNameNisw != nil {

		// query param dns_name__nisw
		var qrDNSNameNisw string

		if o.DNSNameNisw != nil {
			qrDNSNameNisw = *o.DNSNameNisw
		}
		qDNSNameNisw := qrDNSNameNisw
		if qDNSNameNisw != "" {

			if err := r.SetQueryParam("dns_name__nisw", qDNSNameNisw); err != nil {
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

	if o.Interface != nil {

		// query param interface
		var qrInterface string

		if o.Interface != nil {
			qrInterface = *o.Interface
		}
		qInterface := qrInterface
		if qInterface != "" {

			if err := r.SetQueryParam("interface", qInterface); err != nil {
				return err
			}
		}
	}

	if o.Interfacen != nil {

		// query param interface__n
		var qrInterfacen string

		if o.Interfacen != nil {
			qrInterfacen = *o.Interfacen
		}
		qInterfacen := qrInterfacen
		if qInterfacen != "" {

			if err := r.SetQueryParam("interface__n", qInterfacen); err != nil {
				return err
			}
		}
	}

	if o.InterfaceID != nil {

		// query param interface_id
		var qrInterfaceID string

		if o.InterfaceID != nil {
			qrInterfaceID = *o.InterfaceID
		}
		qInterfaceID := qrInterfaceID
		if qInterfaceID != "" {

			if err := r.SetQueryParam("interface_id", qInterfaceID); err != nil {
				return err
			}
		}
	}

	if o.InterfaceIDn != nil {

		// query param interface_id__n
		var qrInterfaceIDn string

		if o.InterfaceIDn != nil {
			qrInterfaceIDn = *o.InterfaceIDn
		}
		qInterfaceIDn := qrInterfaceIDn
		if qInterfaceIDn != "" {

			if err := r.SetQueryParam("interface_id__n", qInterfaceIDn); err != nil {
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

	if o.Parent != nil {

		// query param parent
		var qrParent string

		if o.Parent != nil {
			qrParent = *o.Parent
		}
		qParent := qrParent
		if qParent != "" {

			if err := r.SetQueryParam("parent", qParent); err != nil {
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

	if o.VirtualMachine != nil {

		// query param virtual_machine
		var qrVirtualMachine string

		if o.VirtualMachine != nil {
			qrVirtualMachine = *o.VirtualMachine
		}
		qVirtualMachine := qrVirtualMachine
		if qVirtualMachine != "" {

			if err := r.SetQueryParam("virtual_machine", qVirtualMachine); err != nil {
				return err
			}
		}
	}

	if o.VirtualMachineID != nil {

		// query param virtual_machine_id
		var qrVirtualMachineID string

		if o.VirtualMachineID != nil {
			qrVirtualMachineID = *o.VirtualMachineID
		}
		qVirtualMachineID := qrVirtualMachineID
		if qVirtualMachineID != "" {

			if err := r.SetQueryParam("virtual_machine_id", qVirtualMachineID); err != nil {
				return err
			}
		}
	}

	if o.Vminterface != nil {

		// query param vminterface
		var qrVminterface string

		if o.Vminterface != nil {
			qrVminterface = *o.Vminterface
		}
		qVminterface := qrVminterface
		if qVminterface != "" {

			if err := r.SetQueryParam("vminterface", qVminterface); err != nil {
				return err
			}
		}
	}

	if o.Vminterfacen != nil {

		// query param vminterface__n
		var qrVminterfacen string

		if o.Vminterfacen != nil {
			qrVminterfacen = *o.Vminterfacen
		}
		qVminterfacen := qrVminterfacen
		if qVminterfacen != "" {

			if err := r.SetQueryParam("vminterface__n", qVminterfacen); err != nil {
				return err
			}
		}
	}

	if o.VminterfaceID != nil {

		// query param vminterface_id
		var qrVminterfaceID string

		if o.VminterfaceID != nil {
			qrVminterfaceID = *o.VminterfaceID
		}
		qVminterfaceID := qrVminterfaceID
		if qVminterfaceID != "" {

			if err := r.SetQueryParam("vminterface_id", qVminterfaceID); err != nil {
				return err
			}
		}
	}

	if o.VminterfaceIDn != nil {

		// query param vminterface_id__n
		var qrVminterfaceIDn string

		if o.VminterfaceIDn != nil {
			qrVminterfaceIDn = *o.VminterfaceIDn
		}
		qVminterfaceIDn := qrVminterfaceIDn
		if qVminterfaceIDn != "" {

			if err := r.SetQueryParam("vminterface_id__n", qVminterfaceIDn); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
