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

// NewIpamServicesListParams creates a new IpamServicesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesListParams() *IpamServicesListParams {
	return &IpamServicesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesListParamsWithTimeout creates a new IpamServicesListParams object
// with the ability to set a timeout on a request.
func NewIpamServicesListParamsWithTimeout(timeout time.Duration) *IpamServicesListParams {
	return &IpamServicesListParams{
		timeout: timeout,
	}
}

// NewIpamServicesListParamsWithContext creates a new IpamServicesListParams object
// with the ability to set a context for a request.
func NewIpamServicesListParamsWithContext(ctx context.Context) *IpamServicesListParams {
	return &IpamServicesListParams{
		Context: ctx,
	}
}

// NewIpamServicesListParamsWithHTTPClient creates a new IpamServicesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesListParamsWithHTTPClient(client *http.Client) *IpamServicesListParams {
	return &IpamServicesListParams{
		HTTPClient: client,
	}
}

/* IpamServicesListParams contains all the parameters to send to the API endpoint
   for the ipam services list operation.

   Typically these are written to a http.Request.
*/
type IpamServicesListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Device.
	Device *string

	// Devicen.
	Devicen *string

	// DeviceID.
	DeviceID *string

	// DeviceIDn.
	DeviceIDn *string

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

	// Port.
	Port *float64

	// Protocol.
	Protocol *string

	// Protocoln.
	Protocoln *string

	// Q.
	Q *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	// VirtualMachine.
	VirtualMachine *string

	// VirtualMachinen.
	VirtualMachinen *string

	// VirtualMachineID.
	VirtualMachineID *string

	// VirtualMachineIDn.
	VirtualMachineIDn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesListParams) WithDefaults() *IpamServicesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services list params
func (o *IpamServicesListParams) WithTimeout(timeout time.Duration) *IpamServicesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services list params
func (o *IpamServicesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services list params
func (o *IpamServicesListParams) WithContext(ctx context.Context) *IpamServicesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services list params
func (o *IpamServicesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services list params
func (o *IpamServicesListParams) WithHTTPClient(client *http.Client) *IpamServicesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services list params
func (o *IpamServicesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam services list params
func (o *IpamServicesListParams) WithCreated(created *string) *IpamServicesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam services list params
func (o *IpamServicesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the ipam services list params
func (o *IpamServicesListParams) WithCreatedGte(createdGte *string) *IpamServicesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam services list params
func (o *IpamServicesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the ipam services list params
func (o *IpamServicesListParams) WithCreatedLte(createdLte *string) *IpamServicesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam services list params
func (o *IpamServicesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevice adds the device to the ipam services list params
func (o *IpamServicesListParams) WithDevice(device *string) *IpamServicesListParams {
	o.SetDevice(device)
	return o
}

// SetDevice adds the device to the ipam services list params
func (o *IpamServicesListParams) SetDevice(device *string) {
	o.Device = device
}

// WithDevicen adds the devicen to the ipam services list params
func (o *IpamServicesListParams) WithDevicen(devicen *string) *IpamServicesListParams {
	o.SetDevicen(devicen)
	return o
}

// SetDevicen adds the deviceN to the ipam services list params
func (o *IpamServicesListParams) SetDevicen(devicen *string) {
	o.Devicen = devicen
}

// WithDeviceID adds the deviceID to the ipam services list params
func (o *IpamServicesListParams) WithDeviceID(deviceID *string) *IpamServicesListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the ipam services list params
func (o *IpamServicesListParams) SetDeviceID(deviceID *string) {
	o.DeviceID = deviceID
}

// WithDeviceIDn adds the deviceIDn to the ipam services list params
func (o *IpamServicesListParams) WithDeviceIDn(deviceIDn *string) *IpamServicesListParams {
	o.SetDeviceIDn(deviceIDn)
	return o
}

// SetDeviceIDn adds the deviceIdN to the ipam services list params
func (o *IpamServicesListParams) SetDeviceIDn(deviceIDn *string) {
	o.DeviceIDn = deviceIDn
}

// WithID adds the id to the ipam services list params
func (o *IpamServicesListParams) WithID(id *string) *IpamServicesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services list params
func (o *IpamServicesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDIc adds the iDIc to the ipam services list params
func (o *IpamServicesListParams) WithIDIc(iDIc *string) *IpamServicesListParams {
	o.SetIDIc(iDIc)
	return o
}

// SetIDIc adds the idIc to the ipam services list params
func (o *IpamServicesListParams) SetIDIc(iDIc *string) {
	o.IDIc = iDIc
}

// WithIDIe adds the iDIe to the ipam services list params
func (o *IpamServicesListParams) WithIDIe(iDIe *string) *IpamServicesListParams {
	o.SetIDIe(iDIe)
	return o
}

// SetIDIe adds the idIe to the ipam services list params
func (o *IpamServicesListParams) SetIDIe(iDIe *string) {
	o.IDIe = iDIe
}

// WithIDIew adds the iDIew to the ipam services list params
func (o *IpamServicesListParams) WithIDIew(iDIew *string) *IpamServicesListParams {
	o.SetIDIew(iDIew)
	return o
}

// SetIDIew adds the idIew to the ipam services list params
func (o *IpamServicesListParams) SetIDIew(iDIew *string) {
	o.IDIew = iDIew
}

// WithIDIsw adds the iDIsw to the ipam services list params
func (o *IpamServicesListParams) WithIDIsw(iDIsw *string) *IpamServicesListParams {
	o.SetIDIsw(iDIsw)
	return o
}

// SetIDIsw adds the idIsw to the ipam services list params
func (o *IpamServicesListParams) SetIDIsw(iDIsw *string) {
	o.IDIsw = iDIsw
}

// WithIDn adds the iDn to the ipam services list params
func (o *IpamServicesListParams) WithIDn(iDn *string) *IpamServicesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam services list params
func (o *IpamServicesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithIDNic adds the iDNic to the ipam services list params
func (o *IpamServicesListParams) WithIDNic(iDNic *string) *IpamServicesListParams {
	o.SetIDNic(iDNic)
	return o
}

// SetIDNic adds the idNic to the ipam services list params
func (o *IpamServicesListParams) SetIDNic(iDNic *string) {
	o.IDNic = iDNic
}

// WithIDNie adds the iDNie to the ipam services list params
func (o *IpamServicesListParams) WithIDNie(iDNie *string) *IpamServicesListParams {
	o.SetIDNie(iDNie)
	return o
}

// SetIDNie adds the idNie to the ipam services list params
func (o *IpamServicesListParams) SetIDNie(iDNie *string) {
	o.IDNie = iDNie
}

// WithIDNiew adds the iDNiew to the ipam services list params
func (o *IpamServicesListParams) WithIDNiew(iDNiew *string) *IpamServicesListParams {
	o.SetIDNiew(iDNiew)
	return o
}

// SetIDNiew adds the idNiew to the ipam services list params
func (o *IpamServicesListParams) SetIDNiew(iDNiew *string) {
	o.IDNiew = iDNiew
}

// WithIDNisw adds the iDNisw to the ipam services list params
func (o *IpamServicesListParams) WithIDNisw(iDNisw *string) *IpamServicesListParams {
	o.SetIDNisw(iDNisw)
	return o
}

// SetIDNisw adds the idNisw to the ipam services list params
func (o *IpamServicesListParams) SetIDNisw(iDNisw *string) {
	o.IDNisw = iDNisw
}

// WithLastUpdated adds the lastUpdated to the ipam services list params
func (o *IpamServicesListParams) WithLastUpdated(lastUpdated *string) *IpamServicesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam services list params
func (o *IpamServicesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam services list params
func (o *IpamServicesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamServicesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam services list params
func (o *IpamServicesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam services list params
func (o *IpamServicesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamServicesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam services list params
func (o *IpamServicesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the ipam services list params
func (o *IpamServicesListParams) WithLimit(limit *int64) *IpamServicesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam services list params
func (o *IpamServicesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam services list params
func (o *IpamServicesListParams) WithName(name *string) *IpamServicesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam services list params
func (o *IpamServicesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the ipam services list params
func (o *IpamServicesListParams) WithNameIc(nameIc *string) *IpamServicesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam services list params
func (o *IpamServicesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam services list params
func (o *IpamServicesListParams) WithNameIe(nameIe *string) *IpamServicesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam services list params
func (o *IpamServicesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam services list params
func (o *IpamServicesListParams) WithNameIew(nameIew *string) *IpamServicesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam services list params
func (o *IpamServicesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam services list params
func (o *IpamServicesListParams) WithNameIsw(nameIsw *string) *IpamServicesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam services list params
func (o *IpamServicesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam services list params
func (o *IpamServicesListParams) WithNamen(namen *string) *IpamServicesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam services list params
func (o *IpamServicesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam services list params
func (o *IpamServicesListParams) WithNameNic(nameNic *string) *IpamServicesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam services list params
func (o *IpamServicesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam services list params
func (o *IpamServicesListParams) WithNameNie(nameNie *string) *IpamServicesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam services list params
func (o *IpamServicesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam services list params
func (o *IpamServicesListParams) WithNameNiew(nameNiew *string) *IpamServicesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam services list params
func (o *IpamServicesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam services list params
func (o *IpamServicesListParams) WithNameNisw(nameNisw *string) *IpamServicesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam services list params
func (o *IpamServicesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam services list params
func (o *IpamServicesListParams) WithOffset(offset *int64) *IpamServicesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam services list params
func (o *IpamServicesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPort adds the port to the ipam services list params
func (o *IpamServicesListParams) WithPort(port *float64) *IpamServicesListParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the ipam services list params
func (o *IpamServicesListParams) SetPort(port *float64) {
	o.Port = port
}

// WithProtocol adds the protocol to the ipam services list params
func (o *IpamServicesListParams) WithProtocol(protocol *string) *IpamServicesListParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the ipam services list params
func (o *IpamServicesListParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithProtocoln adds the protocoln to the ipam services list params
func (o *IpamServicesListParams) WithProtocoln(protocoln *string) *IpamServicesListParams {
	o.SetProtocoln(protocoln)
	return o
}

// SetProtocoln adds the protocolN to the ipam services list params
func (o *IpamServicesListParams) SetProtocoln(protocoln *string) {
	o.Protocoln = protocoln
}

// WithQ adds the q to the ipam services list params
func (o *IpamServicesListParams) WithQ(q *string) *IpamServicesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam services list params
func (o *IpamServicesListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the ipam services list params
func (o *IpamServicesListParams) WithTag(tag *string) *IpamServicesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam services list params
func (o *IpamServicesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam services list params
func (o *IpamServicesListParams) WithTagn(tagn *string) *IpamServicesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam services list params
func (o *IpamServicesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WithVirtualMachine adds the virtualMachine to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachine(virtualMachine *string) *IpamServicesListParams {
	o.SetVirtualMachine(virtualMachine)
	return o
}

// SetVirtualMachine adds the virtualMachine to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachine(virtualMachine *string) {
	o.VirtualMachine = virtualMachine
}

// WithVirtualMachinen adds the virtualMachinen to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachinen(virtualMachinen *string) *IpamServicesListParams {
	o.SetVirtualMachinen(virtualMachinen)
	return o
}

// SetVirtualMachinen adds the virtualMachineN to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachinen(virtualMachinen *string) {
	o.VirtualMachinen = virtualMachinen
}

// WithVirtualMachineID adds the virtualMachineID to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachineID(virtualMachineID *string) *IpamServicesListParams {
	o.SetVirtualMachineID(virtualMachineID)
	return o
}

// SetVirtualMachineID adds the virtualMachineId to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachineID(virtualMachineID *string) {
	o.VirtualMachineID = virtualMachineID
}

// WithVirtualMachineIDn adds the virtualMachineIDn to the ipam services list params
func (o *IpamServicesListParams) WithVirtualMachineIDn(virtualMachineIDn *string) *IpamServicesListParams {
	o.SetVirtualMachineIDn(virtualMachineIDn)
	return o
}

// SetVirtualMachineIDn adds the virtualMachineIdN to the ipam services list params
func (o *IpamServicesListParams) SetVirtualMachineIDn(virtualMachineIDn *string) {
	o.VirtualMachineIDn = virtualMachineIDn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Devicen != nil {

		// query param device__n
		var qrDevicen string

		if o.Devicen != nil {
			qrDevicen = *o.Devicen
		}
		qDevicen := qrDevicen
		if qDevicen != "" {

			if err := r.SetQueryParam("device__n", qDevicen); err != nil {
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

	if o.DeviceIDn != nil {

		// query param device_id__n
		var qrDeviceIDn string

		if o.DeviceIDn != nil {
			qrDeviceIDn = *o.DeviceIDn
		}
		qDeviceIDn := qrDeviceIDn
		if qDeviceIDn != "" {

			if err := r.SetQueryParam("device_id__n", qDeviceIDn); err != nil {
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

	if o.Port != nil {

		// query param port
		var qrPort float64

		if o.Port != nil {
			qrPort = *o.Port
		}
		qPort := swag.FormatFloat64(qrPort)
		if qPort != "" {

			if err := r.SetQueryParam("port", qPort); err != nil {
				return err
			}
		}
	}

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string

		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
				return err
			}
		}
	}

	if o.Protocoln != nil {

		// query param protocol__n
		var qrProtocoln string

		if o.Protocoln != nil {
			qrProtocoln = *o.Protocoln
		}
		qProtocoln := qrProtocoln
		if qProtocoln != "" {

			if err := r.SetQueryParam("protocol__n", qProtocoln); err != nil {
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

	if o.VirtualMachinen != nil {

		// query param virtual_machine__n
		var qrVirtualMachinen string

		if o.VirtualMachinen != nil {
			qrVirtualMachinen = *o.VirtualMachinen
		}
		qVirtualMachinen := qrVirtualMachinen
		if qVirtualMachinen != "" {

			if err := r.SetQueryParam("virtual_machine__n", qVirtualMachinen); err != nil {
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

	if o.VirtualMachineIDn != nil {

		// query param virtual_machine_id__n
		var qrVirtualMachineIDn string

		if o.VirtualMachineIDn != nil {
			qrVirtualMachineIDn = *o.VirtualMachineIDn
		}
		qVirtualMachineIDn := qrVirtualMachineIDn
		if qVirtualMachineIDn != "" {

			if err := r.SetQueryParam("virtual_machine_id__n", qVirtualMachineIDn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
